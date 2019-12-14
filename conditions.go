package godynamicconditions

import (
	"reflect"
)

type ConditionType int
type ActionType int

const (
	ConditionInBetween ConditionType = iota

	noopAction ActionType = iota
	customAction
)

type Conditioner interface {
	Evaluate()
}

type Condition struct {
	Type ConditionType
	LeftOperand interface{}
	RightOperand interface{}
	Action *Action
}

type ConditionGTE struct {
	Condition
}

type ConditionLTE struct {
	Condition
}

type ConditionEQ struct {
	Condition
}

func (t ConditionGTE) Evaluate() {
	leftOpVal := reflect.ValueOf(t.LeftOperand)
	rightOpVal := reflect.ValueOf(t.RightOperand)

	if !isGTE(leftOpVal, rightOpVal) {
		return
	}

	if t.Action != nil {
		if t.Action.Effect != nil {
			t.Action.Effect()
		}
	}
}


func (t ConditionLTE) Evaluate() {
}

func (t ConditionEQ) Evaluate() {
	leftOpVal := reflect.ValueOf(t.LeftOperand)
	rightOpVal := reflect.ValueOf(t.RightOperand)

	if !isEqualVal(leftOpVal, rightOpVal) {
		return
	}

	if t.Action != nil {
		if t.Action.Effect != nil {
			t.Action.Effect()
		}
	}
}

type Action struct {
	Type ActionType
	Value interface{}
	Values []interface{}
	Condition *Condition
	Action *Action
	Effect func()
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

func Execute(conditions []Conditioner) {
	for _, cond := range conditions {
		cond.Evaluate()


		//case ConditionInBetween:
		//	leftOpVal := reflect.ValueOf(root.LeftOperand)
		//	rightOpVal := reflect.ValueOf(root.RightOperand)
		//
		//	if !isSameKind(leftOpVal, rightOpVal) {
		//		continue
		//	}
		//
		//	if leftOpVal.Kind() != reflect.Int {
		//		continue
		//	}
		//
		//	if !(leftOpVal.Int() > first && value <= second) {
		//		continue
		//	}
		//
		//	if root.Action != nil {
		//		if root.Action.Effect != nil {
		//			root.Action.Effect()
		//		}
		//	}
		//}
	}
}
