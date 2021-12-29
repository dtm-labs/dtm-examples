package busi

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/gin-gonic/gin"
)

func dbGet() *dtmutil.DB {
	return dtmutil.DbGet(BusiConf)
}

func sdbGet() *sql.DB {
	db, err := dtmimp.PooledDB(BusiConf)
	logger.FatalIfError(err)
	return db
}

func txGet() *sql.Tx {
	db := sdbGet()
	tx, err := db.Begin()
	logger.FatalIfError(err)
	return tx
}

func resetXaData() {
	if BusiConf.Driver != "mysql" {
		return
	}

	db := dbGet()
	type XaRow struct {
		Data string
	}
	xas := []XaRow{}
	db.Must().Raw("xa recover").Scan(&xas)
	for _, xa := range xas {
		db.Must().Exec(fmt.Sprintf("xa rollback '%s'", xa.Data))
	}
}

// MustBarrierFromGin 1
func MustBarrierFromGin(c *gin.Context) *dtmcli.BranchBarrier {
	ti, err := dtmcli.BarrierFromQuery(c.Request.URL.Query())
	logger.FatalIfError(err)
	return ti
}

// MustBarrierFromGrpc 1
func MustBarrierFromGrpc(ctx context.Context) *dtmcli.BranchBarrier {
	ti, err := dtmgrpc.BarrierFromGrpc(ctx)
	logger.FatalIfError(err)
	return ti
}
