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
	AddCommand("http_xa", func() string {
		gid := shortuuid.New()
		err := dtmcli.XaGlobalTransaction(dtmutil.DefaultHTTPServer, gid, func(xa *dtmcli.Xa) (*resty.Response, error) {
			resp, err := xa.CallBranch(&busi.ReqHTTP{Amount: 30}, busi.Busi+"/TransOutXa")
			if err != nil {
				return resp, err
			}
			return xa.CallBranch(&busi.ReqHTTP{Amount: 30}, busi.Busi+"/TransInXa")
		})
		logger.FatalIfError(err)
		return gid
	})
}
