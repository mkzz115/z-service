package log

import (
	"os"
	"strings"
	"testing"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"

	"github.com/cihub/seelog"
)

var ff = `<seelog type="asynctimer" asyncinterval="5000000" minlevel="info" maxlevel="critical">
    <outputs formatid="main">
        <filter levels="trace,debug,info,warn,error,critical">
        <buffered size="10000" flushperiod="1000">
            <rollingfile type="date" filename="run.log" datepattern="20060102" maxrolls="30"/>
        </buffered>
        </filter>
    </outputs>
    <formats>
        <format id="main" format="%Date %Time [%LEV]%Line: %Msg%n"/>
    </formats>
</seelog>`

func TestInfo(t *testing.T) {
	pth := "./test.log"
	//err := Init(pth)
	fff := strings.Replace(ff, "$name$", pth, -1)
	ll, err := seelog.LoggerFromConfigAsString(fff)
	if err != nil {
		t.Logf("errr:%v\n", err.Error())
		return
	}
	seelog.ReplaceLogger(ll)
	defer seelog.Flush()
	ll.Infof("testINfo,%d", 2)

}

func TestWarn(t *testing.T) {
	testLogrotate(t)
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func testLogrotate(t *testing.T) {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "zaplog.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			w),
		zap.InfoLevel,
	)
	logger := zap.New(core, zap.AddCaller())
	logger.Info()
}
