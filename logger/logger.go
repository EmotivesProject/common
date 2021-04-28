package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger      *zap.Logger
	SugarLogger *zap.SugaredLogger
)

func InitLogger(name string) {
	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	Logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))

	defer Logger.Sync()

	SugarLogger = Logger.Sugar()

	atom.SetLevel(zap.InfoLevel)

	Logger = Logger.Named(name)

	Logger.Info("Created Logger")
}
