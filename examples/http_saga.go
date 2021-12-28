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
	addSample("saga", func() string {
		logger.Debugf("a saga busi transaction begin")
		req := &busi.TransReq{Amount: 30}
		saga := dtmcli.NewSaga(dtmsvr.DefaultHttpServer, dtmcli.MustGenGid(dtmsvr.DefaultHttpServer)).
			Add(busi.Busi+"/TransOut", busi.Busi+"/TransOutRevert", req).
			Add(busi.Busi+"/TransIn", busi.Busi+"/TransInRevert", req)
		logger.Debugf("saga busi trans submit")
		err := saga.Submit()
		logger.Debugf("result gid is: %s", saga.Gid)
		logger.FatalIfError(err)
		return saga.Gid
	})
	addSample("saga_wait", func() string {
		logger.Debugf("a saga busi transaction begin")
		req := &busi.TransReq{Amount: 30}
		saga := dtmcli.NewSaga(dtmsvr.DefaultHttpServer, dtmcli.MustGenGid(dtmsvr.DefaultHttpServer)).
			Add(busi.Busi+"/TransOut", busi.Busi+"/TransOutRevert", req).
			Add(busi.Busi+"/TransIn", busi.Busi+"/TransInRevert", req)
		saga.SetOptions(&dtmcli.TransOptions{WaitResult: true})
		err := saga.Submit()
		logger.Debugf("result gid is: %s", saga.Gid)
		logger.FatalIfError(err)
		return saga.Gid
	})
	addSample("concurrent_saga", func() string {
		logger.Debugf("a concurrent saga busi transaction begin")
		req := &busi.TransReq{Amount: 30}
		csaga := dtmcli.NewSaga(dtmsvr.DefaultHttpServer, dtmcli.MustGenGid(dtmsvr.DefaultHttpServer)).
			Add(busi.Busi+"/TransOut", busi.Busi+"/TransOutRevert", req).
			Add(busi.Busi+"/TransOut", busi.Busi+"/TransOutRevert", req).
			Add(busi.Busi+"/TransIn", busi.Busi+"/TransInRevert", req).
			Add(busi.Busi+"/TransIn", busi.Busi+"/TransInRevert", req).
			EnableConcurrent().
			AddBranchOrder(2, []int{0, 1}).
			AddBranchOrder(3, []int{0, 1})
		logger.Debugf("concurrent saga busi trans submit")
		err := csaga.Submit()
		logger.Debugf("result gid is: %s", csaga.Gid)
		logger.FatalIfError(err)
		return csaga.Gid
	})
}
