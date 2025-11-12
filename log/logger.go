package log

import (
    "os"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {

    if os.Getenv("ENV") != "prod" {
        Logger, _ = zap.NewDevelopment()
        return
    }
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.LowercaseLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	)

	Logger = zap.New(core, zap.AddCaller())
}
