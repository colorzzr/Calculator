package main

import "testing"

func Test_order1(t *testing.T){
	inputOp := []string{"23x1", "+" ,"1"}
	higherOrderCalc(inputOp)
	t.Log("Finish Test_order1");
}

func Test_order2(t *testing.T){
	inputOp := []string{"1x2", "+" ,"1", "+", "2x1"}
	higherOrderCalc(inputOp)
	t.Log("Finish Test_order2");
}