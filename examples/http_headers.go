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
	AddCommand("http_saga_customHeaders", func() string {
		gid := shortuuid.New()
		req := &busi.ReqHTTP{Amount: 30}
		saga := dtmcli.NewSaga(dtmutil.DefaultHTTPServer, gid).
			Add(busi.Busi+"/TransOutHeaderYes", "", req) // /TransOutHeaderYes will check header exists
		saga.BranchHeaders = map[string]string{
			"test_header": "test",
		}
		saga.WaitResult = true
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
	AddCommand("http_tcc_customHeaders", func() string {
		gid := shortuuid.New()
		err := dtmcli.TccGlobalTransaction2(dtmutil.DefaultHTTPServer, gid, func(t *dtmcli.Tcc) {
			t.BranchHeaders = map[string]string{
				"test_header": "test",
			}
			t.WaitResult = true
		}, func(tcc *dtmcli.Tcc) (*resty.Response, error) {
			req := &busi.ReqHTTP{Amount: 30}
			return tcc.CallBranch(req, busi.Busi+"/TransOutHeaderYes", "", "")
		})
		logger.FatalIfError(err)
		return gid
	})
}
