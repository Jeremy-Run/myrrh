package main

import (
	"fmt"
	"myrrh/schedule"
)

func main() {
	req := schedule.Requirement{}
	executeResult := req.Execute()
	fmt.Printf("The execute result is: %t \n", executeResult)
}
