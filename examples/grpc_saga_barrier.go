/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/dtm/dtmcli/logger"
	"github.com/dtm-labs/dtm/dtmgrpc"
	"github.com/dtm-labs/dtm/dtmsvr"
	"github.com/dtm-labs/dtm/test/busi"
)

func init() {
	addSample("grpc_saga_barrier", func() string {
		req := &busi.BusiReq{Amount: 30}
		gid := dtmgrpc.MustGenGid(dtmsvr.DefaultGrpcServer)
		saga := dtmgrpc.NewSagaGrpc(dtmsvr.DefaultGrpcServer, gid).
			Add(busi.BusiGrpc+"/examples.Busi/TransOutBSaga", busi.BusiGrpc+"/examples.Busi/TransOutRevertBSaga", req).
			Add(busi.BusiGrpc+"/examples.Busi/TransInBSaga", busi.BusiGrpc+"/examples.Busi/TransInRevertBSaga", req)
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
}
