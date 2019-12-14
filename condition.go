package godynamicconditions

import (
	"reflect"
)

type Conditioner interface {
	Evaluate() bool
}

type Condition struct {
	LeftOperand interface{}
	RightOperand interface{}
	Action *Action
}

func (t Condition) ExecuteAction() {
	if t.Action != nil {
		t.Action.ExecuteEffect()
	}
}

func isSameKind(leftOpVal, rightOpVal reflect.Value) bool {
	return leftOpVal.Kind() == rightOpVal.Kind()
}

func isEqualVal(leftOpVal, rightOpVal reflect.Value) bool {
	if !isSameKind(leftOpVal, rightOpVal) {
		return false
	}

	switch leftOpVal.Kind() {
	case reflect.Int:
		if leftOpVal.Int() != rightOpVal.Int() {
			return false
		}
	case reflect.String:
		if leftOpVal.String() != rightOpVal.String() {
			return false
		}
	case reflect.Bool:
		if leftOpVal.Bool() != rightOpVal.Bool() {
			return false
		}
	}

	return true
}

func isGTE(leftOpVal, rightOpVal reflect.Value) bool {
	if !isSameKind(leftOpVal, rightOpVal)	{
		return false
	}

	switch leftOpVal.Kind() {
	case reflect.Int:
		if !(leftOpVal.Int() >= rightOpVal.Int()) {
			return false
		}
	default:
		return false
	}

	return true
}

func isLTE(leftOpVal, rightOpVal reflect.Value) bool {
	if !isSameKind(leftOpVal, rightOpVal)	{
		return false
	}

	switch leftOpVal.Kind() {
	case reflect.Int:
		if !(leftOpVal.Int() <= rightOpVal.Int()) {
			return false
		}
	default:
		return false
	}

	return true
}

func Execute(conditions []Conditioner) {
	for _, cond := range conditions {
		cond.Evaluate()
	}
}
