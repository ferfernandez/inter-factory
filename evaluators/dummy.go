package evaluators

import (
	"github.com/ferfernandez/inter-factory/factory"
)

type DummyEvaluator struct {
}

func (d *DummyEvaluator) Evaluate(context *factory.EvaluationContext) (bool, error) {
	return (context.Thing != 0), nil
}

func init() {
	var cse factory.ConditionalScopeEvaluator
	cse = &DummyEvaluator{}

	factory.RegisterEvaluator("dummy", &cse)
}
