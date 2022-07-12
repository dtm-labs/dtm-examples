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
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func init() {
	AddCommand("grpc_tcc", func() string {
		logger.Debugf("tcc simple transaction begin")
		gid := shortuuid.New()
		err := dtmgrpc.TccGlobalTransaction(dtmutil.DefaultGrpcServer, gid, func(tcc *dtmgrpc.TccGrpc) error {
			data := &busi.ReqGrpc{Amount: 30}
			r := &emptypb.Empty{}
			err := tcc.CallBranch(data, busi.BusiGrpc+"/busi.Busi/TransOutTcc", busi.BusiGrpc+"/busi.Busi/TransOutConfirm", busi.BusiGrpc+"/busi.Busi/TransOutRevert", r)
			if err != nil {
				return err
			}
			err = tcc.CallBranch(data, busi.BusiGrpc+"/busi.Busi/TransInTcc", busi.BusiGrpc+"/busi.Busi/TransInConfirm", busi.BusiGrpc+"/busi.Busi/TransInRevert", r)
			return err
		})
		logger.FatalIfError(err)
		return gid
	})
	AddCommand("grpc_tcc_rollback", func() string {
		logger.Debugf("tcc simple transaction begin")
		gid := shortuuid.New()
		err := dtmgrpc.TccGlobalTransaction(dtmutil.DefaultGrpcServer, gid, func(tcc *dtmgrpc.TccGrpc) error {
			data := &busi.ReqGrpc{Amount: 30, TransInResult: "FAILURE"}
			r := &emptypb.Empty{}
			err := tcc.CallBranch(data, busi.BusiGrpc+"/busi.Busi/TransOutTcc", busi.BusiGrpc+"/busi.Busi/TransOutConfirm", busi.BusiGrpc+"/busi.Busi/TransOutRevert", r)
			if err != nil {
				return err
			}
			err = tcc.CallBranch(data, busi.BusiGrpc+"/busi.Busi/TransInTcc", busi.BusiGrpc+"/busi.Busi/TransInConfirm", busi.BusiGrpc+"/busi.Busi/TransInRevert", r)
			return err
		})
		logger.Errorf("error is: %s", err.Error())
		return gid
	})
}
