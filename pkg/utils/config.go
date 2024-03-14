package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

const (
	// default values by convention
	DefaultType     = "json"
	DefaultFilename = "config"
	// environment variable key names
	EnvConsulHostKey = "GOCONF_CONSUL"
	EnvTypeKey       = "GOCONF_TYPE"
	EnvFileNameKey   = "GOCONF_FILENAME"
	EnvPrefixKey     = "GOCONF_ENV_PREFIX"
)

var (
	prefix     string
	configType = DefaultType
	configName = DefaultFilename
	dirs       = []string{".", "$HOME", "/usr/local/etc", "/etc"}
)

func LoadConfiguration() (v *viper.Viper, err error) {
	if err = godotenv.Load(); err != nil {
		return
	}
	if v := os.Getenv(EnvTypeKey); len(v) > 0 {
		configType = v
	}
	if v := os.Getenv(EnvFileNameKey); len(v) > 0 {
		configName = v
	}
	if v := os.Getenv(EnvPrefixKey); len(v) > 0 {
		prefix = v
	}

	v = viper.New()
	v.SetConfigType(configType)
	v.SetConfigName(configName)
	if len(prefix) > 0 {
		v.SetEnvPrefix(prefix)
	}
	v.AutomaticEnv()

	// load remote consul
	if ch := os.Getenv(EnvConsulHostKey); ch != "" {
		if err = v.AddRemoteProvider("consul", ch, configName); err != nil {
			return
		} else {
			connect := func() error { return v.ReadRemoteConfig() }
			notify := func(err error, t time.Duration) { log.Println("[consul]", err.Error(), t) }
			b := backoff.NewExponentialBackOff()
			b.MaxElapsedTime = 2 * time.Minute

			err = backoff.RetryNotify(connect, b, notify)
			if err != nil {
				// log.Printf("[consul] giving up connecting to remote config ")
				return
			}
		}
	} else {
		err = errors.New("failed loading remote source; ENV not defined")
		return
	}
	// load from local directory
	for _, d := range dirs {
		v.AddConfigPath(d)
	}
	return
}
