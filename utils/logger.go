package utils

import (
	zap2 "aisecurity/utils/log/zap"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"strings"
	"time"
)

type LogConfig struct {
	Level     string `json:"level"`     // Level 最低日志等级，DEBUG<INFO<WARN<ERROR<FATAL 例如：info-->收集info等级以上的日志
	FileName  string `json:"file_name"` // FileName 日志文件位置
	LocalTime bool   // 是否使用本地时间，默认 UTC 时间
	// 按大小轮转配置
	MaxSize    int  `json:"max_size"`    // MaxSize 进行切割之前，日志文件的最大大小(MB为单位)，默认为100MB
	MaxAge     int  `json:"max_age"`     // MaxAge 是根据文件名中编码的时间戳保留旧日志文件的最大天数。
	MaxBackups int  `json:"max_backups"` // MaxBackups 是要保留的旧日志文件的最大数量。默认是保留所有旧的日志文件（尽管 MaxAge 可能仍会导致它们被删除。）
	Compress   bool // 是否对日志文件进行压缩归档
	// 按时间轮转配置
	RotationTime time.Duration // 日志文件轮转时间
}

// var Logger *zap.Logger
var Logger *zap2.Logger

// NewProductionRotateByTime 创建按时间轮转的 io.Writer
func NewProductionRotateByTime(filename string) io.Writer {
	return NewRotateByTime(NewProductionRotateConfig(filename))
}

// NewProductionRotateBySize 创建按大小轮转的 io.Writer
func NewProductionRotateBySize(filename string) io.Writer {
	return NewRotateBySize(NewProductionRotateConfig(filename))
}

func NewProductionRotateConfig(filename string) *zap2.RotateConfig {
	return &zap2.RotateConfig{
		Filename: filename,
		MaxAge:   30, // 日志保留 30 天

		RotationTime: time.Hour * 24, // 24 小时轮转一次

		MaxSize:    100, // 100M
		MaxBackups: 100,
		Compress:   true,
		LocalTime:  false,
	}
}

func NewRotateByTime(cfg *zap2.RotateConfig) io.Writer {
	opts := []rotatelogs.Option{
		rotatelogs.WithMaxAge(time.Duration(cfg.MaxAge) * time.Hour * 24),
		rotatelogs.WithRotationTime(cfg.RotationTime),
		rotatelogs.WithLinkName(cfg.Filename),
	}
	if !cfg.LocalTime {
		rotatelogs.WithClock(rotatelogs.UTC)
	}
	filename := strings.SplitN(cfg.Filename, ".", 2)
	l, _ := rotatelogs.New(
		filename[0]+".%Y-%m-%d-%H-%M-%S."+filename[1],
		opts...,
	)
	return l
}

func NewRotateBySize(cfg *zap2.RotateConfig) io.Writer {
	return &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		LocalTime:  cfg.LocalTime,
		Compress:   cfg.Compress,
	}
}

// 负责设置 encoding 的日志格式
func getEncoder() zapcore.Encoder {
	// 获取一个指定的的EncoderConfig，进行自定义
	encodeConfig := zap.NewProductionEncoderConfig()
	// 设置每个日志条目使用的键。如果有任何键为空，则省略该条目的部分。
	// 打印格式: {"level":"info","ts":1662032576.6267354,"msg":"test","line":1}

	// 序列化时间。eg: 2022-09-01T19:11:35.921+0800
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 打印格式：{"level":"info","ts":"2022-09-01T19:43:07.178+0800","msg":"test","line":1}

	// "time":"2022-09-01T19:11:35.921+0800"
	encodeConfig.TimeKey = "time"
	// 打印格式：{"level":"info","time":"2022-09-01T19:43:20.558+0800","msg":"test","line":1}

	// 将Level序列化为全大写字符串。例如，将info level序列化为INFO。
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 打印格式：{"level":"INFO","time":"2022-09-01T19:43:41.192+0800","msg":"test","line":1}

	// 以 package/file:行 的格式 序列化调用程序，从完整路径中删除除最后一个目录外的所有目录。
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 打印格式：{"level":"INFO","time":"2022-09-01T19:41:39.819+0800","caller":"test/test.go:20","msg":"test","line":1}
	// 这个需要注意，是要结合 logger := zap.New(core, zap.AddCaller())，一起使用的

	//syncFile := zapcore.AddSync(lumberJackLogger) // 打印到文件
	//syncConsole := zapcore.AddSync(os.Stderr)     // 打印到控制台
	//return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
	return zapcore.NewJSONEncoder(encodeConfig)
}

// 负责日志写入的位置
func getLogWriter(filename string, maxsize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  // 文件位置
		MaxSize:    maxsize,   // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxAge:     maxAge,    // 保留旧文件的最大天数
		MaxBackups: maxBackup, // 保留旧文件的最大个数
		Compress:   false,     // 是否压缩/归档旧文件
	}
	// AddSync 将 io.Writer 转换为 WriteSyncer。
	// 它试图变得智能：如果 io.Writer 的具体类型实现了 WriteSyncer，我们将使用现有的 Sync 方法。
	// 如果没有，我们将添加一个无操作同步。

	return zapcore.AddSync(lumberJackLogger)
}

// InitLogger 初始化Logger
func InitLogger() {
	conf := LogConfig{
		Level:    "debug",
		FileName: fmt.Sprintf("./logs/general/%v.log", time.Now().Unix()),
		MaxSize:  1,
		MaxAge:   100,
	}
	// 获取日志写入位置
	writeSyncer := getLogWriter(conf.FileName, conf.MaxSize, conf.MaxBackups, conf.MaxAge)
	// 获取日志编码格式
	encoder := getEncoder()

	// 获取日志最低等级，即>=该等级，才会被写入。
	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(conf.Level)); err != nil {
		log.Printf("failed to unmarshal logger level: %v", err)
		return
	}

	// 创建一个将日志写入 WriteSyncer 的核心。
	core := zapcore.NewCore(encoder, writeSyncer, l)

	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(zap.New(core, zap.AddCaller()))
	//Logger = zap.L()
}

func InitLogger2() {
	tees := []zap2.TeeOption{
		{
			Out: zap2.NewProductionRotateBySize("log/pile.log"),
			LevelEnablerFunc: zap2.LevelEnablerFunc(func(level zap2.Level) bool {
				//return level < zap2.WarnLevel
				return true
			}),
		},
		{
			Out: zap2.NewProductionRotateByTime("log/date/date.log"),
			LevelEnablerFunc: zap2.LevelEnablerFunc(func(level zap2.Level) bool {
				//return level >= zap2.WarnLevel
				return true
			}),
		},
	}
	opts := []zap2.Option{
		// 附加日志调用信息
		zap2.WithCaller(true),
		zap2.AddCallerSkip(1),
		// Warn 级别日志 Hook
		zap2.Hooks(func(entry zapcore.Entry) error {
			if entry.Level == zap2.WarnLevel {
				fmt.Printf("Warn Hook: msg=%s\n", entry.Message)
			}
			return nil
		}),
		// Fatal 级别日志 Hook
		zap.WithFatalHook(FatalHook{}),
	}
	Logger = zap2.NewTee(tees, opts...)
}

type FatalHook struct{}

func (h FatalHook) OnWrite(ce *zapcore.CheckedEntry, field []zapcore.Field) {
	fmt.Printf("Fatal Hook: msg=%s, field=%+v\n", ce.Message, field)
}
