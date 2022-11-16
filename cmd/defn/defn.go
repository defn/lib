package main

import (
	"fmt"

	_ "embed"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

//go:embed feh.cue
var user_schema_cue string

func main() {
	ctx := cuecontext.New()

	user_input_instance := load.Instances([]string{"."}, nil)[0]

	user_schema := ctx.CompileString(user_schema_cue)

	user_input := ctx.BuildInstance(user_input_instance)

	valid := user_schema.Unify(user_input)

	fmt.Printf("%v\n", valid)
}
