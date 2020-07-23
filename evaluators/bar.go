package evaluators

import (
	"github.com/ferfernandez/inter-factory/factory"
)

type BarEvaluator struct {
}

func (d *BarEvaluator) Evaluate(context *factory.EvaluationContext) (bool, error) {
	return false, nil
}

func init() {
	var cse factory.ConditionalScopeEvaluator
	cse = &BarEvaluator{}

	factory.RegisterEvaluator("bar", &cse)
}
