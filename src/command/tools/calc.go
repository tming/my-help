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

	// 解析字符串为表达式
	expr, err := eval.ParseString(str, "")
	if err != nil {
		fmt.Println("Invalid expression:", str)
		return err
	}

	// 计算表达式的值
	result, err := expr.EvalToInterface(nil)
	if err != nil {
		fmt.Println("Evaluation error:", err)
		return err
	}

	// 打印计算结果
	fmt.Println("result: ", result)

	return nil
}
