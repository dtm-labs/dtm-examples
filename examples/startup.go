package examples

import (
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/dtm-labs/dtmcli/logger"
)

type commandInfo struct {
	Arg    string
	Action func() string
	Desc   string
}

// Commands 所有的示例都会注册到这里
var Commands = []commandInfo{}

func AddCommand(name string, fn func() string) {
	logger.FatalfIf(IsExists(name), "%s already exists", name)
	Commands = append(Commands, commandInfo{Arg: name, Action: fn})
}

func IsExists(name string) bool {
	for _, c := range Commands {
		if c.Arg == name {
			return true
		}
	}
	return false
}

func Call(name string) {
	for _, c := range Commands {
		if c.Arg == name {
			c.Action()
		}
	}
}

var DtmServer = dtmutil.DefaultHttpServer
