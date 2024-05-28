package ml

import (
	"io"
	"nfthook/config"
	"os"
	"path"

	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	os.MkdirAll(config.Get().Env.Log, 0755)
	Log.SetReportCaller(false)
	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	// Log.SetLevel(logrus.DebugLevel)

	logFileName := path.Join(config.Get().Env.Log, "nfthook") + ".%Y%m%d.log"
	logFileCut := LogFileCut(logFileName)
	writers := []io.Writer{logFileCut, os.Stdout}

	fileAndStdoutWriter := io.MultiWriter(writers...)
	gin.DefaultWriter = fileAndStdoutWriter
	Log.SetOutput(fileAndStdoutWriter)
}

func LogFileCut(fileName string) *rotatelogs.RotateLogs {
	logier, err := rotatelogs.New(
		fileName,
		//rotatelogs.WithLinkName(Current.LogDir),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		//rotatelogs.WithRotationCount(3),
		//rotatelogs.WithRotationTime(time.Minute),
	)
	if err != nil {
		panic(err)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  logier,
		logrus.FatalLevel: logier,
		logrus.DebugLevel: logier,
		logrus.WarnLevel:  logier,
		logrus.ErrorLevel: logier,
		logrus.PanicLevel: logier,
	},
		// 设置分割日志样式
		&logrus.TextFormatter{})
	logrus.AddHook(lfHook)
	return logier
}
