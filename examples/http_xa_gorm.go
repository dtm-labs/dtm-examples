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
	addSample("http_xa_gorm", func() string {
		gid := dtmcli.MustGenGid(dtmutil.DefaultHttpServer)
		err := busi.XaClient.XaGlobalTransaction(gid, func(xa *dtmcli.Xa) (*resty.Response, error) {
			resp, err := xa.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TransOutXaGorm")
			if err != nil {
				return resp, err
			}
			return xa.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TransInXa")
		})
		logger.FatalIfError(err)
		return gid
	})

}
