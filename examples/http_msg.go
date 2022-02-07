/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"database/sql"

	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
)

func init() {
	AddCommand("http_msg", func() string {
		logger.Debugf("a busi transaction begin")
		req := &busi.TransReq{Amount: 30}
		msg := dtmcli.NewMsg(dtmutil.DefaultHTTPServer, dtmcli.MustGenGid(dtmutil.DefaultHTTPServer)).
			Add(busi.Busi+"/TransOut", req).
			Add(busi.Busi+"/TransIn", req)
		err := msg.Prepare(busi.Busi + "/query")
		logger.FatalIfError(err)
		logger.Debugf("busi trans submit")
		err = msg.Submit()
		logger.FatalIfError(err)
		return msg.Gid
	})
	AddCommand("http_msg_doAndCommit", func() string {
		gid := dtmcli.MustGenGid(DtmServer)
		req := busi.GenTransReq(30, false, false)
		msg := dtmcli.NewMsg(DtmServer, gid).
			Add(busi.Busi+"/SagaBTransIn", req)
		err := msg.DoAndSubmitDB(busi.Busi+"/QueryPreparedB", dtmutil.DbGet(busi.BusiConf).ToSQLDB(), func(tx *sql.Tx) error {
			return busi.SagaAdjustBalance(tx, busi.TransOutUID, -req.Amount, "SUCCESS")
		})
		logger.FatalIfError(err)
		return gid
	})
}
