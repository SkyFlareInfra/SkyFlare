package pkg

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogService interface {
	LogInfo(...interface{})
	LogInfoFormat(string, ...interface{})
	LogErrorFormat(string, ...interface{})
	LogError(string, error, ...zap.Field)
	LogFatalFormat(string, ...interface{})
	LogPanicFormat(string, ...interface{})
	WithContext(map[string]interface{}) LogService
}

type ZapWrapper struct {
	core *zap.Logger
}

func CreateLogService(configEnv DatabaseConfig) LogService {
	zapConfig := buildZapConfig()
	logCore, err := zapConfig.Build()
	if err != nil {
		panic("log system initialization failed: " + err.Error())
	}
	
	return &ZapWrapper{core: logCore}
}

func buildZapConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(parseLogLevel()),
		Encoding:    "json",
		OutputPaths: []string{determineOutput()},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "severity",
			TimeKey:      "timestamp",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
}

func parseLogLevel() zapcore.Level {
	level := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL")))
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info", "":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func determineOutput() string {
	output := strings.TrimSpace(os.Getenv("LOG_OUTPUT"))
	if output == "" {
		return "stdout"
	}
	return output
}

func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case error:
		return val.Error()
	default:
		return fmt.Sprintf("%v", val)
	}
}

func (z *ZapWrapper) LogInfo(args ...interface{}) {
	z.core.Info(formatArgs(args))
}

func (z *ZapWrapper) LogInfoFormat(template string, args ...interface{}) {
	z.core.Info(formatTemplate(template, args))
}

func (z *ZapWrapper) LogErrorFormat(template string, args ...interface{}) {
	z.core.Error(formatTemplate(template, args))
}

func (z *ZapWrapper) LogError(msg string, err error, fields ...zap.Field) {
	allFields := append(fields, zap.NamedError("root_cause", err))
	z.core.Error(msg, allFields...)
}

func (z *ZapWrapper) LogFatalFormat(template string, args ...interface{}) {
	z.core.Fatal(formatTemplate(template, args))
}

func (z *ZapWrapper) LogPanicFormat(template string, args ...interface{}) {
	z.core.Panic(formatTemplate(template, args))
}

func (z *ZapWrapper) WithContext(attributes map[string]interface{}) LogService {
	fields := convertToZapFields(attributes)
	return &ZapWrapper{core: z.core.With(fields...)}
}

func formatArgs(args []interface{}) string {
	if len(args) == 1 {
		return toString(args[0])
	}
	return fmt.Sprint(args...)
}

func formatTemplate(template string, args []interface{}) string {
	if len(args) == 0 {
		return template
	}
	return fmt.Sprintf(template, args...)
}

func convertToZapFields(attributes map[string]interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(attributes))
	for key, value := range attributes {
		fields = append(fields, zap.Any(key, value))
	}
	return fields
}
