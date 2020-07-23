package main

import (
	"fmt"
	"os"

	_ "github.com/ferfernandez/inter-factory/evaluators"
	"github.com/ferfernandez/inter-factory/factory"
)

func main() {
	evname := "none"
	if len(os.Args) > 1 {
		evname = os.Args[1]
	}

	if e, exists := factory.Get(evname); !exists {
		fmt.Println("No Evaluator found called", evname)
	} else {
		eval, error := (*e).Evaluate(&factory.EvaluationContext{Thing: 1, Some: "algo"})

		if error != nil {
			fmt.Println("Evaluation error ", error)
		} else {
			fmt.Println(eval)
		}
	}
}
