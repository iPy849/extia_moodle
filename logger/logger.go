package logger

import (
	"errors"
	"io"
	"log"
	"os"
)

var (
	Logger                             *log.Logger
	errCannotConnectLoggerToOutputFile = errors.New("cannot find or create a log file")
)

const (
	STDOUT_LOGGER    = "stdout"
	DEFAULT_LOG_FILE = "./logs"
)

// Writer
type externalWriter struct {
	file *os.File
}

func (e externalWriter) Write(p []byte) (n int, err error) {
	n, err = e.file.Write(p)
	if err != nil {
		panic(err)
	}
	return
}

func newExternalWriter(path string) (e externalWriter) {
	var err error
	e.file, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(errCannotConnectLoggerToOutputFile)
	}
	return
}

// Logger constructor
func New(loggerOutput ...string) error {
	writers := make([]io.Writer, 0)
	for _, path := range loggerOutput {
		switch path {
		case STDOUT_LOGGER:
			writers = append(writers, os.Stdout)
			break
		case DEFAULT_LOG_FILE:
			writers = append(writers, newExternalWriter(DEFAULT_LOG_FILE))
			break
		default:
			writers = append(writers, newExternalWriter(path))
		}
	}

	Logger = log.New(io.MultiWriter(writers...), log.Prefix(), log.Llongfile|log.LstdFlags)
	return nil
}
