package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

var (
	GinLogPath = "D:/GoProject/src/github.com/Fanny-wen/gin_demo/logs"
	GinLogFile = "gin.log"
)

var GinLogInstance *logrus.Logger

func InitLog(logPath, logFile string) {
	var LogInstance = logrus.New()
	logFileName := filepath.Join(logPath, logFile)

	// 使用滚动压缩方式记录日志
	rolling(logFileName, LogInstance)

	var f *os.File
	var err error
	//判断日志文件是否存在，不存在则创建，否则就直接打开
	if _, err := os.Stat(logFileName); os.IsNotExist(err) {
		f, err = os.Create(logFileName)
	} else {
		f, err = os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}
	//f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("open file %s faild, err: %v\n", logFileName, err))
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 设置日志输出到文件
	LogInstance.SetOutput(f)
	// 设置日志输出格式
	LogInstance.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	//LogInstance.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	// 设置日志记录级别
	LogInstance.SetLevel(logrus.DebugLevel)
	// 将调用的函数名添加为字段
	//LogInstance.SetReportCaller(true)
	GinLogInstance = LogInstance
}

// 日志滚动设置
func rolling(logFile string, log *logrus.Logger) {
	// 设置输出
	log.SetOutput(&lumberjack.Logger{
		Filename: logFile, //日志文件位置
		MaxSize:  1,       // 单文件最大容量,单位是MB
		//MaxBackups: 3,// 最大保留过期文件个数
		//MaxAge: 1 ,// 保留过期文件的最大时间间隔,单位是天
		Compress: true, // 是否需要压缩滚动日志, 使用的 gzip 压缩
	})
}
