package godynamicconditions

import "reflect"

type ConditionLTE struct {
	Condition
}

func (t ConditionLTE) Evaluate() bool {
	leftOpVal := reflect.ValueOf(t.LeftOperand)
	rightOpVal := reflect.ValueOf(t.RightOperand)

	if !isLTE(leftOpVal, rightOpVal) {
		return false
	}

	if t.Action != nil {
		if t.Action.Effect != nil {
			t.Action.Effect()
		}
	}

	return true
}

