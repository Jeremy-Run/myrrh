package main

import (
	"flag"

	"encoding/json"
	"fmt"
	"myrrh/config"
	"myrrh/schedule"
)

func main() {

	num := flag.String("num", "1", "Test sequence number")
	flag.Parse()

	caseActivity := config.CaseActivity1
	caseExpr := config.CaseExpr1
	switch *num {
	case "1":
		caseActivity = config.CaseActivity1
		caseExpr = config.CaseExpr1
	case "2":
		caseActivity = config.CaseActivity2
		caseExpr = config.CaseExpr2
	case "3":
		caseActivity = config.CaseActivity3
		caseExpr = config.CaseExpr3
	}

	req := schedule.Requirement{}
	_ = json.Unmarshal([]byte(caseActivity), &req)
	fmt.Println(">>> Start......")

	desc := schedule.ExpressionDescription{}
	_ = json.Unmarshal([]byte(caseExpr), &desc)
	simpleExpr := schedule.SimpleExpression(&desc)
	fmt.Printf(">>> Simple Expression: %+v \n", simpleExpr)

	executeResult := req.Execute()
	fmt.Printf(">>> The requirement execute result is: %t \n", executeResult)
	fmt.Println(">>> End......")
}
