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
	"github.com/go-resty/resty/v2"
)

func init() {
	addSample("tcc_nested", func() string {
		gid := dtmcli.MustGenGid(dtmsvr.DefaultHttpServer)
		err := dtmcli.TccGlobalTransaction(dtmsvr.DefaultHttpServer, gid, func(tcc *dtmcli.Tcc) (*resty.Response, error) {
			resp, err := tcc.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TransOut", busi.Busi+"/TransOutConfirm", busi.Busi+"/TransOutRevert")
			if err != nil {
				return resp, err
			}
			return tcc.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TransInTccParent", busi.Busi+"/TransInConfirm", busi.Busi+"/TransInRevert")
		})
		logger.FatalIfError(err)
		return gid
	})
	addSample("tcc", func() string {
		logger.Debugf("tcc simple transaction begin")
		gid := dtmcli.MustGenGid(dtmsvr.DefaultHttpServer)
		err := dtmcli.TccGlobalTransaction(dtmsvr.DefaultHttpServer, gid, func(tcc *dtmcli.Tcc) (*resty.Response, error) {
			resp, err := tcc.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TransOut", busi.Busi+"/TransOutConfirm", busi.Busi+"/TransOutRevert")
			if err != nil {
				return resp, err
			}
			return tcc.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TransIn", busi.Busi+"/TransInConfirm", busi.Busi+"/TransInRevert")
		})
		logger.FatalIfError(err)
		return gid
	})
}
