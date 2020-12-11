package logger

import (
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log              *zap.Logger
	customTimeFormat string
	onceInit         sync.Once
)

func customeTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(customTimeFormat))
}

func Init(lvl int, timeFormat string) error {
	var err error
	//Do calls the function f if and only if Do is being called for the first time for this instance of Once
	onceInit.Do(func() {
		globalLevel := zapcore.Level(lvl)
		//high => stderr, low => standard out
		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})
		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= globalLevel && lvl < zapcore.ErrorLevel
		})

		consoleInfos := zapcore.Lock(os.Stdout)
		consoleErrors := zapcore.Lock(os.Stderr)

		//Configure console output
		var useCustomTimeFormat bool
		ecfg := zap.NewProductionEncoderConfig()
		if len(timeFormat) > 0 {
			customTimeFormat = timeFormat
			ecfg.EncodeTime = customeTimeEncoder
			useCustomTimeFormat = true
		}

		//Join outputs, encoders, and levelhandling functions into zapcore
		//NewTee creates a Core that duplicates log entries into two or more underlying Cores.
		consoleEncoder := zapcore.NewJSONEncoder(ecfg)
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
			zapcore.NewCore(consoleEncoder, consoleInfos, lowPriority),
		)

		//Construct a Logger
		Log = zap.New(core)
		zap.RedirectStdLog(Log)
		if !useCustomTimeFormat {
			Log.Warn("time format for logger is not provided - use zap default")
		}
	})
	return err
}
