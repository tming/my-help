package tools

import (
	"fmt"
	"my-help/src/common"

	"github.com/apaxa-go/eval"
)

func Calc(str string) error {
	if str == "" {
		return common.ErrorTarget
	}

	// parse str to expression
	expr, err := eval.ParseString(str, "")
	if err != nil {
		fmt.Println("Invalid expression:", str)
		return err
	}

	// calc result for expression
	result, err := expr.EvalToInterface(nil)
	if err != nil {
		fmt.Println("Evaluation error:", err)
		return err
	}

	// output
	fmt.Println("result: ", result)

	return nil
}
