/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/dtm/dtmcli"
	"github.com/dtm-labs/dtm/dtmcli/logger"
	"github.com/dtm-labs/dtm/dtmsvr"
	"github.com/dtm-labs/dtm/test/busi"
)

func init() {
	addSample("msg", func() string {
		logger.Debugf("a busi transaction begin")
		req := &busi.TransReq{Amount: 30}
		msg := dtmcli.NewMsg(dtmsvr.DefaultHttpServer, dtmcli.MustGenGid(dtmsvr.DefaultHttpServer)).
			Add(busi.Busi+"/TransOut", req).
			Add(busi.Busi+"/TransIn", req)
		err := msg.Prepare(busi.Busi + "/query")
		logger.FatalIfError(err)
		logger.Debugf("busi trans submit")
		err = msg.Submit()
		logger.FatalIfError(err)
		return msg.Gid
	})
}
