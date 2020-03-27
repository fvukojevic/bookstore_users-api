/**
We separate our logging system in it's own package. If tomorrow we want to change
our logging system, we only have to look at this file.
*/

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			/**
			Here we define how our log will look. For this Encoder config it should look something like
			{
				"level":"info",
				"time": "2020-01-01T23:00:23..."
				"msg": "logging message"
			}
			*/
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,    // Time format for Time Key
			EncodeLevel:  zapcore.LowercaseLevelEncoder, // Format for Level Key
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	return log
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	_ = log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	_ = log.Sync()
}
