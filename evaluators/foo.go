package evaluators

import (
	"fmt"

	"github.com/ferfernandez/inter-factory/factory"
)

type FooEvaluator struct {
}

func (d *FooEvaluator) Evaluate(context *factory.EvaluationContext) (bool, error) {
	if context.Some == "algo" {
		return false, fmt.Errorf("I Cannot evaluate Some == Algo")
	}
	return (context.Thing != 0), nil
}

func init() {
	var cse factory.ConditionalScopeEvaluator
	cse = &FooEvaluator{}

	factory.RegisterEvaluator("foo", &cse)
}
