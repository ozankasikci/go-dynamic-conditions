package godynamicconditions

import "reflect"

type ConditionGTE struct {
	Condition
}

func (t ConditionGTE) Evaluate() bool {
	leftOpVal := reflect.ValueOf(t.LeftOperand)
	rightOpVal := reflect.ValueOf(t.RightOperand)

	if !isGTE(leftOpVal, rightOpVal) {
		return false
	}

	if t.Action != nil {
		if t.Action.Effect != nil {
			t.Action.Effect()
		}
	}

	return true
}
