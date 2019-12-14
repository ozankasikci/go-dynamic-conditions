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
		gdc.And{
			LeftCondition: gdc.ConditionEQ{
				Condition: gdc.Condition{
					LeftOperand:  false,
					RightOperand: true,
					Action: &gdc.Action{
						Effect: func() {
							println(1)
						},
					},
				},
			},
			RightCondition: gdc.ConditionEQ{
				Condition: gdc.Condition{
					LeftOperand:  false,
					RightOperand: true,
					Action: &gdc.Action{
						Effect: func() {
							println(1)
						},
					},
				},
			},
		},
	}

	gdc.Execute(conditions)
	s, _ := res["test"].(string)
	println(s)
}
