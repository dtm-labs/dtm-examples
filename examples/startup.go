package examples

import (
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
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

type PostRoute struct {
	Route   string
	Handler func(*gin.Context) interface{}
}

var routes = []PostRoute{}

func AddRoutes(app *gin.Engine) {
	for _, r := range routes {
		app.POST(r.Route, dtmutil.WrapHandler2(r.Handler))
	}
}

var DtmServer = dtmutil.DefaultHTTPServer
