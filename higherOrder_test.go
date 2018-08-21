package main

import (
	"testing"
	"fmt"
)

func Test_order1(t *testing.T){
	inputOp := []string{"23x1", "a" ,"1"}
	higherOrderCalc(inputOp)
	t.Log("Finish Test_order1");
}

func Test_order2(t *testing.T){
	inputOp := []string{"1x2", "a" ,"1", "a", "2x1"}
	higherOrderCalc(inputOp)
	t.Log("Finish Test_order2");
}

func Test_order2_2(t *testing.T){
	inputOp := []string{"1x2", "a" ,"4", "a", "1x1"}
	ans := higherOrderCalc(inputOp)
	switch ans[0].(type) {
	case float64:
		fmt.Println("float64!")
		break
	case complex64:
		fmt.Println("complex64!")
		break
	case complex128:
		fmt.Println("complex128!")
		break
	default:
		fmt.Println("??!")
		break
	}
	t.Log("Finish Test_order2_2");
}


func Test_order3_1(t *testing.T){
	inputOp := []string{"1x3", "a" ,"8"}
	ans := higherOrderCalc(inputOp)
	fmt.Println(ans)
}


func Test_order3_2(t *testing.T){
	inputOp := []string{"1x3", "a" ,"2x2", "a", "3x1", "a", "8"}
	ans := higherOrderCalc(inputOp)
	fmt.Println(ans)
}

func Test_order3_3(t *testing.T){
	inputOp := []string{"1x2", "-" ,"1x3", "a", "3x1", "a", "1"}
	ans := higherOrderCalc(inputOp)
	fmt.Println(ans)
}

func Test_order3_4(t *testing.T){
	inputOp := []string{"1231x2", "-" ,"444x3", "-", "9876x1", "a", "5555"}
	ans := higherOrderCalc(inputOp)
	fmt.Println(ans)
}