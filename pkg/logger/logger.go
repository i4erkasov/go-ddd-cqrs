package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(cnf *viper.Viper) (*zap.Logger, error) {
	var cfg zap.Config

	// Загрузка общей конфигурации логгера
	if err := cnf.UnmarshalKey("app.logger", &cfg); err != nil {
		return nil, err
	}

	// Проверяем окружение
	var encoderConfig zapcore.EncoderConfig
	if cnf.GetString("app.environment") == "production" {
		encoderConfig = zap.NewProductionEncoderConfig()
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}

	// Настройка энкодера
	if cnf.GetString("app.logger.encoderConfig.timeEncoder") == "iso8601" {
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}
	cfg.EncoderConfig = encoderConfig

	// Настройка уровня логирования
	if level, err := zapcore.ParseLevel(cnf.GetString("app.logger.lvl")); err == nil {
		cfg.Level = zap.NewAtomicLevelAt(level)
	}

	// Настройка путей вывода логов
	if paths := cnf.GetStringSlice("app.logger.outputPaths"); len(paths) > 0 {
		cfg.OutputPaths = paths
	}
	if errorPaths := cnf.GetStringSlice("app.logger.errorOutputPaths"); len(errorPaths) > 0 {
		cfg.ErrorOutputPaths = errorPaths
	}

	// Создание и возвращение логгера
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	// Настройки для вызывающего и трассировки стека
	if cnf.GetBool("app.logger.caller") {
		logger = logger.WithOptions(zap.AddCaller())
	}
	if stacktraceLevel := cnf.GetString("app.logger.stacktrace"); stacktraceLevel != "" {
		var level zapcore.Level
		if err := level.UnmarshalText([]byte(stacktraceLevel)); err == nil {
			logger = logger.WithOptions(zap.AddStacktrace(level))
		}
	}

	return logger, nil
}
