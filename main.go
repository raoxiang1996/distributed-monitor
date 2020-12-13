package main

import (
	"distributed-monitor/sampled_data"
	"distributed-monitor/utils"
	"fmt"
)

func main() {
	fmt.Println(utils.Ip)
	fmt.Println(utils.Port)
	fmt.Println(sampled_data.GetState('Z'))
}
