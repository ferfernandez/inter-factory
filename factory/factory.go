package factory

import (
	"sync"
)

type EvaluationContext struct {
	Some  string
	Thing int
}

type ConditionalScopeEvaluator interface {
	Evaluate(context *EvaluationContext) (bool, error)
}

var (
	mu         sync.RWMutex
	evaluators = make(map[string](*ConditionalScopeEvaluator))
)

func RegisterEvaluator(name string, e *ConditionalScopeEvaluator) {
	mu.Lock()
	defer mu.Unlock()

	if e == nil {
		panic("Nil Evaluator not allowed")
	}
	if _, exist := evaluators[name]; exist {
		panic("Evaluator is already registered")
	}

	evaluators[name] = e
}

func Get(evaluatorName string) (*ConditionalScopeEvaluator, bool) {
	mu.RLock()
	defer mu.RUnlock()
	if e, exist := evaluators[evaluatorName]; exist {
		return e, true
	}

	return nil, false
}
