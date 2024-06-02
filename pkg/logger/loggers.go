package loggers

import (
	"os"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"
	logrusadapter "logur.dev/adapter/logrus"
	"logur.dev/logur"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	return log
}

func NewLoggerTemporal(l *logrus.Logger) logur.KVLoggerFacade {
	return logur.LoggerToKV(logrusadapter.New(l))

}

func HasSensitiveData(data interface{}) (hasSensitiveData bool) {
	checkData := reflect.ValueOf(data)
	if checkData.Kind() == reflect.Ptr {
		checkData = checkData.Elem()
	}

	switch checkData.Kind() {
	case reflect.Struct:
		for i := 0; i < checkData.NumField(); i++ {
			typeField := checkData.Type().Field(i)
			tag := typeField.Tag.Get("json")

			if tag == "card_number" {
				hasSensitiveData = true
				break
			}
		}
	case reflect.Map:
		for _, k := range checkData.MapKeys() {
			if k.String() == "card_number" {
				hasSensitiveData = true
				break
			}
		}

	}

	return
}
