package godynamicconditions

type And struct {
	LeftCondition Condition
	RightCondition Condition
}

func (t And) Evaluate()  {

}
