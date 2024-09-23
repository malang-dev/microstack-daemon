package logs

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type loggerCtxKeyType string

// LoggerKey is the key to use for the logger in the context.
const LoggerKey loggerCtxKeyType = "microstackd.logger"

const (
	// LogFormatText specifies a textual log format.
	LogFormatText = "text"
	// LogFormatJSON specifies a JSON log format.
	LogFormatJSON = "json"
)

const (
	// LogVerbosityInfo is the verbosity level for info logging.
	LogVerbosityInfo = 0
	// LogVerbosityDebug is the verbosity level for debug logging.
	LogVerbosityDebug = 2
	// LogVerbosityTrace is the verbosity level for trace logging.
	LogVerbosityTrace = 9
)

type Config struct {
	// Verbosity specifies the logging verbosity level.
	Verbosity int
	// Format specifies the logging output format.
	Format string
	// Output specifies the destination for logging. You can specify the special
	// values of 'stderr' or 'stdout' or a file path.
	Output string
}

var CommandContext context.Context = nil

// Configure will configure the logger from the supplied config.
func Configure(logConfig *Config) error {
	configureVerbosity(logConfig)

	if err := configureFormatter(logConfig); err != nil {
		return fmt.Errorf("configuring log formatter: %w", err)
	}

	if err := configureOutput(logConfig); err != nil {
		return fmt.Errorf("configuring log output: %w", err)
	}

	return nil
}

func configureFormatter(logConfig *Config) error {
	switch logConfig.Format {
	case LogFormatJSON:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case LogFormatText:
		logrus.SetFormatter(&logrus.TextFormatter{})
	default:
		return invalidLogFormatError{format: logConfig.Format}
	}

	return nil
}

func configureVerbosity(logConfig *Config) {
	logrus.SetLevel(logrus.InfoLevel)

	if logConfig.Verbosity >= LogVerbosityDebug && logConfig.Verbosity < LogVerbosityTrace {
		logrus.SetLevel(logrus.DebugLevel)
	} else if logConfig.Verbosity >= LogVerbosityTrace {
		logrus.SetLevel(logrus.TraceLevel)
	}
}

func configureOutput(logConfig *Config) error {
	output := strings.ToLower(logConfig.Output)
	switch output {
	case "stdout":
		logrus.SetOutput(os.Stdout)
	case "stderr":
		logrus.SetOutput(os.Stderr)
	case "":
		return ErrLogOutputRequired
	default:
		file, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			return fmt.Errorf("opening log file %s: %w", output, err)
		}

		logrus.SetOutput(file)
	}

	return nil
}

// WithLogger is used to attached a logger to a specific context.
func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(CommandContext, LoggerKey, logger)
}

// GetLogger will get a logger from the supplied context for create a new logger.
func GetLogger() *logrus.Entry {
	logger := CommandContext.Value(LoggerKey)

	if logger == nil {
		return logrus.NewEntry(logrus.StandardLogger())
	}

	return logger.(*logrus.Entry)
}

// SetContext will set the context for the logger.
func SetContext(ctx context.Context) {
	CommandContext = ctx
}
