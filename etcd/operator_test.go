package etcd

import (
	"fmt"
	"testing"
)

func TestOperator(t *testing.T) {
	Init()
	//Insert(Clictl, "rao", "xiang")
	value, _ := Search(Clictl, "rao")
	fmt.Printf("value: %s\n", value)
}
