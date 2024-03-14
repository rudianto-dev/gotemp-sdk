package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
)

func RequestLogger(logger *log.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&logrusRequestLogger{Logger: logger})
}

func RequestLoggerWithBody(logger *log.Logger, blacklistKeys []string) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&logrusRequestLogger{Logger: logger, BlacklistKey: blacklistKeys})
}

type logrusRequestLogger struct {
	Logger       *log.Logger
	BlacklistKey []string
}

func (l *logrusRequestLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &logrusRequestLoggerEntry{Logger: log.NewEntry(l.Logger)}
	logFields := log.Fields{}

	logFields["ts"] = time.Now().UTC().Format(time.RFC1123)
	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["req_id"] = reqID
	}
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	logFields["http_scheme"] = scheme
	logFields["http_proto"] = r.Proto
	logFields["http_method"] = r.Method
	logFields["remote_addr"] = r.RemoteAddr
	logFields["user_agent"] = r.UserAgent()
	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	if len(l.BlacklistKey) > 0 {
		logFields["request_body"] = requestBody(r, l.BlacklistKey)
	}

	entry.Logger = entry.Logger.WithFields(logFields)

	return entry
}

type logrusRequestLoggerEntry struct {
	Logger log.FieldLogger
}

func (l *logrusRequestLoggerEntry) Write(status, bytes int, _ http.Header, elapsed time.Duration, _ interface{}) {
	l.Logger = l.Logger.WithFields(log.Fields{
		"resp_status": status, "resp_bytes_length": bytes,
		"resp_elapsed_ms": float64(elapsed.Nanoseconds()) / 1000000.0,
	})
	l.Logger.Infoln("request complete")
}

func (l *logrusRequestLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.WithFields(log.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
}

func MakeLogEntry(action string) (l *log.Entry) {
	l = log.WithFields(log.Fields{
		"scope":  "middleware",
		"action": action,
	})
	return
}

func GetLogEntry(r *http.Request) log.FieldLogger {
	entry := middleware.GetLogEntry(r).(*logrusRequestLoggerEntry)
	return entry.Logger
}

func GetLogtrailByCtx(ctx context.Context) log.FieldLogger {
	entry, _ := ctx.Value(middleware.LogEntryCtxKey).(*logrusRequestLoggerEntry)
	entry.Logger = entry.Logger.WithFields(log.Fields{
		"log_type": "audit",
	})
	return entry.Logger
}

func LogEntrySetField(r *http.Request, key string, value interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*logrusRequestLoggerEntry); ok {
		entry.Logger = entry.Logger.WithField(key, value)
	}
}

func LogEntrySetFields(r *http.Request, fields map[string]interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*logrusRequestLoggerEntry); ok {
		entry.Logger = entry.Logger.WithFields(fields)
	}
}

func Logtrail(ctx context.Context, key string, userId string, oldValue, newValue interface{}) {
	if entry, ok := ctx.Value(middleware.LogEntryCtxKey).(*logrusRequestLoggerEntry); ok && oldValue != newValue {
		data := map[string]interface{}{
			"id":       userId,
			"old_data": oldValue,
			"new_data": newValue,
		}
		entry.Logger.WithFields(log.Fields{
			"logtrail": key,
			"log_type": "audit",
		}).Println(data)
	}
}

func requestBody(r *http.Request, blacklistKeys []string) string {
	buf, _ := ioutil.ReadAll(r.Body)
	readerLog := ioutil.NopCloser(bytes.NewBuffer(buf))
	readerBody := ioutil.NopCloser(bytes.NewBuffer(buf))
	r.Body = readerBody

	requestBody := map[string]interface{}{}

	err := json.NewDecoder(readerLog).Decode(&requestBody)
	if err != nil {
		return ""
	}

	filterSensitiveAtrribute(requestBody, blacklistKeys)

	bodyJSON, _ := json.Marshal(requestBody)

	return string(bodyJSON)
}
func filterSensitiveAtrribute(body map[string]interface{}, blacklistKeys []string) {
	for key, value := range body {
		if valBody, ok := value.(map[string]interface{}); ok {
			filterSensitiveAtrribute(valBody, blacklistKeys)
			continue
		}

		if valBody, ok := value.([]map[string]interface{}); ok {
			for _, valBodySlice := range valBody {
				filterSensitiveAtrribute(valBodySlice, blacklistKeys)
			}
			continue
		}

		for _, blacklistKey := range blacklistKeys {
			if strings.Contains(key, blacklistKey) {
				body[key] = "*REMOVED*"
				continue
			}
		}

	}
}
