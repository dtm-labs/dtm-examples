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
	addSample("tcc_barrier", func() string {
		logger.Debugf("tcc transaction begin")
		gid := dtmcli.MustGenGid(dtmsvr.DefaultHttpServer)
		err := dtmcli.TccGlobalTransaction(dtmsvr.DefaultHttpServer, gid, func(tcc *dtmcli.Tcc) (*resty.Response, error) {
			resp, err := tcc.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TccBTransOutTry",
				busi.Busi+"/TccBTransOutConfirm", busi.Busi+"/TccBTransOutCancel")
			if err != nil {
				return resp, err
			}
			return tcc.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TccBTransInTry", busi.Busi+"/TccBTransInConfirm", busi.Busi+"/TccBTransInCancel")
		})
		logger.FatalIfError(err)
		return gid
	})
}
