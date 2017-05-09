package log

import (
	"github.com/astaxie/beego/logs"
)

func Debug(s string) {
	logs.Debug(s)
}
