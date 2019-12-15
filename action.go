package godynamicconditions

type Action struct {
	Value interface{}
	Values []interface{}
	Conditions []Conditioner
	Effect func()
}

func (t Action) Execute() {
	t.ExecuteEffect()
	t.EvaluateConditions()
}

func (t Action) ExecuteEffect()  {
	if t.Effect != nil {
		t.Effect()
	}
}

func (t Action) EvaluateConditions() {
	if t.Conditions != nil {
		for _, cond := range t.Conditions {
			cond.Evaluate()
		}
	}
}
