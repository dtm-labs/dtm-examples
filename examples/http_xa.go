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
	addSample("xa", func() string {
		gid := dtmcli.MustGenGid(dtmsvr.DefaultHttpServer)
		err := busi.XaClient.XaGlobalTransaction(gid, func(xa *dtmcli.Xa) (*resty.Response, error) {
			resp, err := xa.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TransOutXa")
			if err != nil {
				return resp, err
			}
			return xa.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TransInXa")
		})
		logger.FatalIfError(err)
		return gid
	})
}
