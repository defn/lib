package main

import (
	"fmt"
	"time"

	"github.com/defn/cloud/pkg/meh"
)

func main() {
	fmt.Println(meh.Hello("pants 9"))
	time.Sleep(86400 * time.Second)
}
