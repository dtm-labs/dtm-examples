#!/bin/bash

set -x

rm -rf busi dtmutil
cp -rf ../dtm/test/busi ./
cp -rf ../dtm/dtmutil ./
sed -i '' -e 's/dtm-labs\/dtm\/dtmutil/dtm-labs\/dtm-examples\/dtmutil/g' *.go */**.go
sed -i '' -e 's/dtm-labs\/dtm\/dtmcli/dtm-labs\/dtmcli/g' *.go */**.go
sed -i '' -e 's/dtm-labs\/dtm\/dtmgrpc/dtm-labs\/dtmgrpc/g' *.go */**.go
