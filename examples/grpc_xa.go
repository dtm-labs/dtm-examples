/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/dtm-labs/dtmgrpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func init() {
	AddCommand("grpc_xa", func() string {
		gid := dtmgrpc.MustGenGid(dtmutil.DefaultGrpcServer)
		req := &busi.BusiReq{Amount: 30}
		err := busi.XaGrpcClient.XaGlobalTransaction(gid, func(xa *dtmgrpc.XaGrpc) error {
			r := &emptypb.Empty{}
			err := xa.CallBranch(req, busi.BusiGrpc+"/busi.Busi/TransOutXa", r)
			if err != nil {
				return err
			}
			err = xa.CallBranch(req, busi.BusiGrpc+"/busi.Busi/TransInXa", r)
			return err
		})
		logger.FatalIfError(err)
		return gid
	})
}
