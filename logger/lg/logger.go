package lg

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

var L = logrus.New()

//全局logger配置
func init() {
	L.SetReportCaller(true)
	L.SetOutput(os.Stdout)
	L.SetLevel(logrus.DebugLevel)
	L.SetFormatter(&logrus.TextFormatter{ForceColors: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			function = ""
			//处理文件名
			file = fmt.Sprintf(" %s:%d ", frame.File, frame.Line)
			return
		}},
	)
}
