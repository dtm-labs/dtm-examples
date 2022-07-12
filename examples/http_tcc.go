/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmcli/logger"
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/go-resty/resty/v2"
	"github.com/lithammer/shortuuid/v3"
)

func init() {
	AddCommand("http_tcc_nested", func() string {
		gid := shortuuid.New()
		err := dtmcli.TccGlobalTransaction(dtmutil.DefaultHTTPServer, gid, func(tcc *dtmcli.Tcc) (*resty.Response, error) {
			resp, err := tcc.CallBranch(&busi.ReqHTTP{Amount: 30}, busi.Busi+"/TransOut", busi.Busi+"/TransOutConfirm", busi.Busi+"/TransOutRevert")
			if err != nil {
				return resp, err
			}
			return tcc.CallBranch(&busi.ReqHTTP{Amount: 30}, busi.Busi+"/TransInTccNested", busi.Busi+"/TransInConfirm", busi.Busi+"/TransInRevert")
		})
		logger.FatalIfError(err)
		return gid
	})
	AddCommand("http_tcc", func() string {
		logger.Debugf("tcc simple transaction begin")
		gid := shortuuid.New()
		err := dtmcli.TccGlobalTransaction(dtmutil.DefaultHTTPServer, gid, func(tcc *dtmcli.Tcc) (*resty.Response, error) {
			resp, err := tcc.CallBranch(&busi.ReqHTTP{Amount: 30}, busi.Busi+"/TransOut", busi.Busi+"/TransOutConfirm", busi.Busi+"/TransOutRevert")
			if err != nil {
				return resp, err
			}
			return tcc.CallBranch(&busi.ReqHTTP{Amount: 30}, busi.Busi+"/TransIn", busi.Busi+"/TransInConfirm", busi.Busi+"/TransInRevert")
		})
		logger.FatalIfError(err)
		return gid
	})
	AddCommand("http_tcc_rollback", func() string {
		logger.Debugf("tcc simple transaction begin")
		gid := shortuuid.New()
		err := dtmcli.TccGlobalTransaction(dtmutil.DefaultHTTPServer, gid, func(tcc *dtmcli.Tcc) (*resty.Response, error) {
			req := &busi.ReqHTTP{Amount: 30, TransInResult: "FAILURE"}
			resp, err := tcc.CallBranch(req, busi.Busi+"/TransOut", busi.Busi+"/TransOutConfirm", busi.Busi+"/TransOutRevert")
			if err != nil {
				return resp, err
			}
			return tcc.CallBranch(req, busi.Busi+"/TransIn", busi.Busi+"/TransInConfirm", busi.Busi+"/TransInRevert")
		})
		logger.Errorf("error is: %s", err.Error())
		return gid
	})
}
