/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/go-resty/resty/v2"
)

func init() {
	AddCommand("http_saga_customHeaders", func() string {
		gid := dtmcli.MustGenGid(dtmutil.DefaultHttpServer)
		req := &busi.TransReq{Amount: 30}
		saga := dtmcli.NewSaga(dtmutil.DefaultHttpServer, gid).
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
		gid := dtmcli.MustGenGid(dtmutil.DefaultHttpServer)
		err := dtmcli.TccGlobalTransaction2(dtmutil.DefaultHttpServer, gid, func(t *dtmcli.Tcc) {
			t.BranchHeaders = map[string]string{
				"test_header": "test",
			}
			t.WaitResult = true
		}, func(tcc *dtmcli.Tcc) (*resty.Response, error) {
			req := &busi.TransReq{Amount: 30}
			return tcc.CallBranch(req, busi.Busi+"/TransOutHeaderYes", "", "")
		})
		logger.FatalIfError(err)
		return gid
	})
	AddCommand("http_saga_passthroughHeaders", func() string {
		dtmcli.SetPassthroughHeaders([]string{"test_header"}) // set passthrough headers. dtm will
		gid := dtmcli.MustGenGid(dtmutil.DefaultHttpServer) + "HeadersYes"
		dtmcli.OnBeforeRequest(busi.SetHttpHeaderForHeadersYes) // will set header in this middleware
		req := &busi.TransReq{Amount: 30}
		saga := dtmcli.NewSaga(dtmutil.DefaultHttpServer, gid).
			Add(busi.Busi+"/TransOutHeaderYes", "", req) // /TransOutHeaderYes will check header exists
		saga.WaitResult = true
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
}
