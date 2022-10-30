package main

import (
	"fmt"
	"time"

	"github.com/defn/cloud/pkg/meh"

	"k8s.io/client-go/kubernetes/fake"
)

func main() {
	fmt.Println(fake.NewSimpleClientset().CoreV1().Events(""))
	fmt.Println(meh.Hello("pants 11"))
	time.Sleep(86400 * time.Second)
}
