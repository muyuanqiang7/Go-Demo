package main

import (
	"errors"
	"fmt"
)

func sqrt(param float64) (float64, error) {
	if param < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return param, nil
}

func main() {
	result, err := sqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	// 正常情况
	if result, errorMsg := divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}

type DivideError struct {
	dividee int
	divider int
}

func (divideError *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
	`
	return fmt.Sprintf(strFormat, divideError.dividee)
}

// define `int` division function
func divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}
