package main

import (
	"fmt"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

func main() {
	bps := load.Instances([]string{"."}, nil)
	ctx := cuecontext.New()
	for _, b := range bps {
		v := ctx.BuildInstance(b)
		fmt.Printf("%v\n", v)
	}
}
