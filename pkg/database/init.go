package database

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"

	_ "github.com/lib/pq"
)

type DB struct {
	Master                *sqlx.DB
	Slave                 *sqlx.DB
	Logger                *logger.Logger
	Config                DatabaseConfig
	Debug                 bool
	IsStopCheckConnection chan bool
	Scope                 string
}

type DatabaseConfig struct {
	DriverName              string
	SourceMaster            string
	SourceSlave             string
	MaxOpenConns            int
	MaxIdleConns            int
	ConnMaxLifetime         int
	IntervalCheckConnection int
}

const (
	Postgres = "postgres"
	Mysql    = "mysql"
)

func NewDatabase(config DatabaseConfig, logger *logger.Logger, scope string, debug bool) (*DB, error) {
	// store the config in the database
	database := &DB{Config: config, Logger: logger, Scope: scope, Debug: debug}
	err := database.connect()
	if err != nil {
		return nil, err
	}

	return database, nil
}

func (db *DB) connect() error {
	var err error
	db.Master, err = connectDB(db.Config, db.Config.SourceMaster)
	if err != nil {
		return err
	}

	// check if there is slave connection
	if db.Config.SourceSlave != "" {
		db.Slave, err = connectDB(db.Config, db.Config.SourceSlave)
		if err != nil {
			return err
		}
	}
	return nil
}

func connectDB(config DatabaseConfig, source string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(config.DriverName, source)
	if err != nil {
		return nil, err
	}

	db.Mapper = reflectx.NewMapper("json")
	return db, err
}

func (db *DB) CheckConnection() {
	db.IsStopCheckConnection = make(chan bool)

	go func() {
		for {
			select {
			case <-db.IsStopCheckConnection:
				log.Println("postgres stop checking connection")
				return
			case <-time.After(time.Duration(db.Config.IntervalCheckConnection) * time.Second):
				db.checkingConnection()
			}
		}
	}()
}

func (db *DB) checkingConnection() {
	err := db.Master.Ping()
	if err != nil {
		log.Println("postgres master connection lost. Reconnecting...")

		masterConn, err := connectDB(db.Config, db.Config.SourceMaster)
		if err != nil {
			log.Println("Postgres master connection is error: ", err)
		} else {
			db.Master = masterConn
			log.Println("Reconnected to postgres master")
		}
	}

	if db.Slave != nil {
		err := db.Slave.Ping()
		if err != nil {
			log.Println("postgres slave connection lost. Reconnecting...")

			slaveConn, err := connectDB(db.Config, db.Config.SourceSlave)
			if err != nil {
				log.Println("Postgres slave connection is error: ", err)
			} else {
				db.Slave = slaveConn
				log.Println("Reconnected to postgres slave")
			}
		}
	}
}

func (db *DB) Close() {
	if db.IsStopCheckConnection != nil {
		// stop check connection
		close(db.IsStopCheckConnection)
	}

	if db.Master != nil {
		db.Master.Close()
	}

	if db.Slave != nil {
		db.Slave.Close()
	}
}

func (db *DB) GetMasterConnection() (*sqlx.DB, error) {
	var err error
	if db.Master != nil {
		return db.Master, nil
	}
	db.Master, err = connectDB(db.Config, db.Config.SourceMaster)
	if err != nil {
		return nil, err
	}
	return db.Master, nil
}
