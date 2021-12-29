package busi

import (
	"fmt"

	common "github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/gin-gonic/gin"
)

// Startup startup the busi's grpc and http service
func Startup() *gin.Engine {
	GrpcStartup()
	return BaseAppStartup()
}

// PopulateDB populate example mysql data
func PopulateDB(skipDrop bool) {
	resetXaData()
	file := fmt.Sprintf("%s/busi.%s.sql", common.GetSqlDir(), BusiConf.Driver)
	common.RunSQLScript(BusiConf, file, skipDrop)
	file = fmt.Sprintf("%s/dtmcli.barrier.%s.sql", common.GetSqlDir(), BusiConf.Driver)
	common.RunSQLScript(BusiConf, file, skipDrop)
}
