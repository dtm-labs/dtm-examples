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
	"google.golang.org/protobuf/types/known/emptypb"
)

func init() {
	AddCommand("grpc_xa", func() string {
		gid := shortuuid.New()
		req := &busi.ReqGrpc{Amount: 30}
		err := dtmgrpc.XaGlobalTransaction(dtmutil.DefaultGrpcServer, gid, func(xa *dtmgrpc.XaGrpc) error {
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
