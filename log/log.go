package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type CustomFormatter struct {
	logrus.TextFormatter
}

var (
	defaultTimeFormatter = "2006/01/02 15:04:05"
	stringBufferPool     = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&CustomFormatter{})
}

func SetLevel(level logrus.Level) {
	logrus.SetLevel(level)
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	buffer := stringBufferPool.Get().(*bytes.Buffer)
	defer func() {
		buffer.Reset()
		stringBufferPool.Put(buffer)
	}()

	// 2006/01/02 15:04:05 [info] [main.go:400]
	buffer.WriteString(
		fmt.Sprintf("%s [%s] ",
			entry.Time.Format(defaultTimeFormatter),
			strings.ToUpper(entry.Level.String()),
		),
	)
	if entry.HasCaller() {
		fileBaseName := entry.Caller.File
		fileName := strings.TrimSuffix(fileBaseName, filepath.Ext(fileBaseName))
		buffer.WriteString(
			fmt.Sprintf("[%s:%d] ",
				fileName,
				entry.Caller.Line,
			),
		)
	}
	buffer.WriteString(fmt.Sprintf("%s\n", entry.Message))
	return buffer.Bytes(), nil
}
