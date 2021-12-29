package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dtm-labs/dtm-examples/examples"
)

func main() {
	if len(os.Args) == 1 {
		for name := range examples.Samples {
			fmt.Printf("%4s%-18srun a sample includes %s\n", "", name, strings.ReplaceAll(name, "_", " "))
		}
	}
	examples.QsStartSvr()
	examples.QsFireRequest()
	time.Sleep(1 * time.Second)
}
