/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/dtm/dtmcli/logger"
	dtmgrpc "github.com/dtm-labs/dtm/dtmgrpc"
	"github.com/dtm-labs/dtm/dtmsvr"
	"github.com/dtm-labs/dtm/test/busi"
)

func init() {
	addSample("grpc_saga", func() string {
		req := &busi.BusiReq{Amount: 30}
		gid := dtmgrpc.MustGenGid(dtmsvr.DefaultGrpcServer)
		saga := dtmgrpc.NewSagaGrpc(dtmsvr.DefaultGrpcServer, gid).
			Add(busi.BusiGrpc+"/examples.Busi/TransOut", busi.BusiGrpc+"/examples.Busi/TransOutRevert", req).
			Add(busi.BusiGrpc+"/examples.Busi/TransIn", busi.BusiGrpc+"/examples.Busi/TransInRevert", req)
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
	addSample("grpc_saga_wait", func() string {
		req := &busi.BusiReq{Amount: 30}
		gid := dtmgrpc.MustGenGid(dtmsvr.DefaultGrpcServer)
		saga := dtmgrpc.NewSagaGrpc(dtmsvr.DefaultGrpcServer, gid).
			Add(busi.BusiGrpc+"/examples.Busi/TransOut", busi.BusiGrpc+"/examples.Busi/TransOutRevert", req).
			Add(busi.BusiGrpc+"/examples.Busi/TransIn", busi.BusiGrpc+"/examples.Busi/TransInRevert", req)
		saga.WaitResult = true
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
}
