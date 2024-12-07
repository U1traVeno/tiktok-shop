package utils

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitHlog() {

	// Customizable output directory.
	var logFilePath string
	dir := "./hlog"
	logFilePath = dir + "/logs/"
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		log.Println(err.Error())
		return
	}

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return
		}
	}

	// For logrus detailed settings, please refer to https://github.com/hertz-contrib/logger/tree/main/logrus and https://github.com/sirupsen/logrus
	logger := hertzlogrus.NewLogger()
	logger.Logger().SetReportCaller(true)
	// hlog will warp a layer of logrus, so you need to calculate the depth of the caller file separately.
	logger.Logger().AddHook(NewCustomHook(10))
	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}

	logger.SetOutput(lumberjackLogger)
	logger.SetLevel(hlog.LevelDebug)
	// if you want to output the log to the file and the stdout at the same time, you can use the following codes

	// fileWriter := io.MultiWriter(lumberjackLogger, os.Stdout)
	// logger.SetOutput(fileWriter)
	hlog.SetLogger(logger)
}

// CustomHook Custom Hook for processing logs
type CustomHook struct {
	CallerDepth int
}

func NewCustomHook(depth int) *CustomHook {
	return &CustomHook{
		CallerDepth: depth,
	}
}

func (hook *CustomHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *CustomHook) Fire(entry *logrus.Entry) error {
	// Get caller information and specify depth
	pc, file, line, ok := runtime.Caller(hook.CallerDepth)
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		entry.Data["caller"] = fmt.Sprintf("%s:%d %s", file, line, funcName)
	}

	return nil
}
