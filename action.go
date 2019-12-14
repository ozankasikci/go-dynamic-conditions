package godynamicconditions

type Action struct {
	Value interface{}
	Values []interface{}
	Condition Conditioner
	Action *Action
	Effect func()
}

func (t Action) ExecuteAction()  {
	if t.Action == nil && t.Effect != nil {
		t.ExecuteEffect()
	} else if t.Action != nil && t.Effect != nil {
		t.Action.ExecuteAction()
		t.ExecuteEffect()
	}
}

func (t Action) ExecuteEffect()  {
	if t.Effect != nil {
		t.Effect()
	}
}
