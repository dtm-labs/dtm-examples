#!/bin/bash

set -x

rm -rf busi dtmutil
cp -rf ../dtm/test/busi ./
cp -rf ../dtm/dtmutil ./
sed -i '' -e 's/dtm-labs\/dtm\/dtmutil/dtm-labs\/dtm-examples\/dtmutil/g' *.go */**.go
rm -rf dtmutil/*_test.go
sed -i '' -e 's/dtm-labs\/dtm\/client\/dtmcli/dtm-labs\/client\/dtmcli/g' *.go */**.go
sed -i '' -e 's/dtm-labs\/dtm\/client\/dtmgrpc/dtm-labs\/client\/dtmgrpc/g' *.go */**.go
sed -i '' -e 's/dtm-labs\/dtm\/client\/workflow/dtm-labs\/client\/workflow/g' *.go */**.go
