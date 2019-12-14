package godynamicconditions

import "reflect"

type ConditionEQ struct {
	Condition
}

func (t ConditionEQ) Evaluate() bool {
	leftOpVal := reflect.ValueOf(t.LeftOperand)
	rightOpVal := reflect.ValueOf(t.RightOperand)

	if !isEqualVal(leftOpVal, rightOpVal) {
		return false
	}

	t.ExecuteAction()
	return true
}

