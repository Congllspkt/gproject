package cli

import (
	"os"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LogMain() {
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("info log", zap.Int("line", 1))
	logger.Error("error log", zap.Int("line", 2))

}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {

	file, _ := os.OpenFile("./track_log/logFile.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncComsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncComsole, syncFile)
}
