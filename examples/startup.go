package examples

import "github.com/dtm-labs/dtmcli/logger"

type sampleInfo struct {
	Arg    string
	Action func() string
	Desc   string
}

// Samples 所有的示例都会注册到这里
var Samples = map[string]*sampleInfo{}

func addSample(name string, fn func() string) {
	logger.FatalfIf(Samples[name] != nil, "%s already exists", name)
	Samples[name] = &sampleInfo{Arg: name, Action: fn}
}
