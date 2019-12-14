package main

import gdc "github.com/ozankasikci/go-dynamic-conditions"

func main() {
	res := make(map[string]interface{})

	conditions := []gdc.Conditioner{
		gdc.ConditionGTE{
			Condition: gdc.Condition{
				LeftOperand: 2,
				RightOperand: 1,
				Action: &gdc.Action{
					Effect: func() {
						println(1111)
						res["test"] = "asdf"
					},
				},
			},
		},
		gdc.ConditionEQ{
			Condition: gdc.Condition{
				LeftOperand:  0,
				RightOperand: 0,
				Action: &gdc.Action{
					Effect: func() {
						println(1)
					},
				},
			},
		},
	}

	gdc.Execute(conditions)
	s, _ := res["test"].(string)
	println(s)
}
