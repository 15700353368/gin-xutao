package initialize

import (
	"gin-xutao/global"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)


//func Zap() *zap.Logger {
//	writeSyncer, _ := os.Create("./log/info.log")                           //日志文件存放目录
//	encoderConfig := zap.NewProductionEncoderConfig()                    //指定时间格式
//	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
//	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
//	encoder := zapcore.NewConsoleEncoder(encoderConfig)                 //获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
//	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)    //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
//	return zap.New(core,zap.AddCaller())                                //AddCaller()为显示文件名和行号
//
//}

var Logger *zap.Logger

func Zap() *zap.Logger {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	// 实现两个判断日志等级的interface (其实 zapcore.*Level 自身就是 interface)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现

	infoWriter := getWriter(global.GVA_CONFIG.Zap.Director)
	warnWriter := getWriter(global.GVA_CONFIG.Zap.Director)

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)

	return zap.New(core, zap.AddCaller())  // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
}

func getWriter(filename string) *rotatelogs.RotateLogs {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+"%Y%m%d"+".log", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}


