package rabbitmq

import (
	"fmt"
	"log"

	"github.com/cronusqiu90/golib/amqp/logger"
)

// Logger is describes a logging structure. It can be set using
// WithPublisherOptionsLogger() or WithConsumerOptionsLogger().
type Logger logger.Logger

const loggingPrefix = "[amqp]"

type stdDebugLogger struct{}

// Fatalf -
func (l stdDebugLogger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(fmt.Sprintf("%s: %s", loggingPrefix, format), v...)
}

// Errorf -
func (l stdDebugLogger) Errorf(format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("[E] %s: %s", loggingPrefix, format), v...)
}

// Warnf -
func (l stdDebugLogger) Warnf(format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("[W] %s: %s", loggingPrefix, format), v...)
}

// Infof -
func (l stdDebugLogger) Infof(format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("[I] %s: %s", loggingPrefix, format), v...)
}

// Debugf -
func (l stdDebugLogger) Debugf(format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("[D] %s: %s", loggingPrefix, format), v...)
}
