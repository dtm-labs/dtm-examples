/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/client/dtmcli/logger"
	dtmgrpc "github.com/dtm-labs/client/dtmgrpc"
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/lithammer/shortuuid/v3"
)

func init() {
	AddCommand("grpc_msg", func() string {
		req := &busi.ReqGrpc{Amount: 30}
		gid := shortuuid.New()
		msg := dtmgrpc.NewMsgGrpc(dtmutil.DefaultGrpcServer, gid).
			Add(busi.BusiGrpc+"/busi.Busi/TransOut", req).
			Add(busi.BusiGrpc+"/busi.Busi/TransIn", req)
		err := msg.Submit()
		logger.FatalIfError(err)
		return msg.Gid
	})
}
