package main

import (
	"encoding/json"
	"fmt"
	"myrrh/config"
	"myrrh/schedule"
)

func main() {
	req := schedule.Requirement{}
	_ = json.Unmarshal([]byte(config.SimulationActivity), &req)
	fmt.Println(">>> Start executing requirement......")

	desc := schedule.ExpressionDescription{}
	_ = json.Unmarshal([]byte(config.SimulationExpr), &desc)
	simplifyExpr := schedule.SimplifyExpression(&desc)
	fmt.Printf(">>> Simplify Expression: %+v \n", simplifyExpr)

	executeResult := req.Execute()
	fmt.Printf(">>> The requirement execute result is: %t \n", executeResult)
	fmt.Println(">>> End of requirement execution......")
}
