package godynamicconditions

type And struct {
	LeftCondition Conditioner
	RightCondition Conditioner
	Action *Action
}

func (t And) Evaluate() bool {
	if t.LeftCondition.Evaluate() && t.RightCondition.Evaluate() {
		t.Action.ExecuteEffect()
		return true
	}

	return false
}
