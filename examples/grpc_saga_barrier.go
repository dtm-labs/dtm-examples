/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/client/dtmcli/logger"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/lithammer/shortuuid/v3"
)

func init() {
	AddCommand("grpc_saga_barrier", func() string {
		req := &busi.ReqGrpc{Amount: 30}
		gid := shortuuid.New()
		saga := dtmgrpc.NewSagaGrpc(dtmutil.DefaultGrpcServer, gid).
			Add(busi.BusiGrpc+"/busi.Busi/TransOutBSaga", busi.BusiGrpc+"/busi.Busi/TransOutRevertBSaga", req).
			Add(busi.BusiGrpc+"/busi.Busi/TransInBSaga", busi.BusiGrpc+"/busi.Busi/TransInRevertBSaga", req)
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
}
