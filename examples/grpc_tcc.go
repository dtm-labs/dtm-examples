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
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func init() {
	addSample("grpc_tcc", func() string {
		logger.Debugf("tcc simple transaction begin")
		gid := dtmgrpc.MustGenGid(dtmsvr.DefaultGrpcServer)
		err := dtmgrpc.TccGlobalTransaction(dtmsvr.DefaultGrpcServer, gid, func(tcc *dtmgrpc.TccGrpc) error {
			data := &busi.BusiReq{Amount: 30}
			r := &emptypb.Empty{}
			err := tcc.CallBranch(data, busi.BusiGrpc+"/examples.Busi/TransOutTcc", busi.BusiGrpc+"/examples.Busi/TransOutConfirm", busi.BusiGrpc+"/examples.Busi/TransOutRevert", r)
			if err != nil {
				return err
			}
			err = tcc.CallBranch(data, busi.BusiGrpc+"/examples.Busi/TransInTcc", busi.BusiGrpc+"/examples.Busi/TransInConfirm", busi.BusiGrpc+"/examples.Busi/TransInRevert", r)
			return err
		})
		logger.FatalIfError(err)
		return gid
	})
}
