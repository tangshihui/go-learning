package checker

import "fmt"

//package run steps
//1 variables in package level
//2 init function

var NAME string = "CHECKER"

func init() {
	fmt.Println("init package checker")
}

func CheckArgs(args ...interface{}) {
	for _, v := range args {
		switch v.(type) {
		case int:
			fmt.Println(v, "type is int.")
		case float32:
			fmt.Println(v, "type is float32.")
		case float64:
			fmt.Println(v, "type is float64.")
		case string:
			fmt.Println(v, "type is string.")
		default:
			fmt.Println(v, "type unknown.")
		}
	}
}

func CheckArg(arg interface{}) string {
	switch v := arg.(type) {
	case int:
		return fmt.Sprintf("%v:%v", v, "type is int.")
	case float32:
		return fmt.Sprintf("%v:%v", v, "type is float32.")
	case float64:
		return fmt.Sprintf("%v:%v", v, "type is float64.")
	case string:
		return fmt.Sprintf("%v:%v", v, "type is string.")
	default:
		return fmt.Sprintf("%v:%v", v, "type unknown.")
	}
}
