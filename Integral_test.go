package main

import (
	"testing"
)

func Test_Integral_1(t *testing.T){

	fun := []string{"x","^","2"}

	RiemannSumIntegral(fun, 1.0, 2.0)
}

func Test_Integral_2(t *testing.T){

	fun := []string{"9", "*","x","^","3"}

	RiemannSumIntegral(fun, 1.0, 2.0)
}

func Test_Integral_3(t *testing.T){

	fun := []string{"9", "/", "(","x","^","3", ")"}

	RiemannSumIntegral(fun, 1.0, 2.0)
}

func Test_Double_Integral_1(t *testing.T){

	fun := []string{"x", "a", "y"}

	doubleRiemannSumIntegral(fun, 2.0, 1.0, 2.0, 1.0)
}

func Test_Double_Integral_2(t *testing.T){

	fun := []string{"x", "*", "y"}

	doubleRiemannSumIntegral(fun, 2.0, 1.0, 2.0, 1.0)
}

func Test_Triple_Integral_1(t *testing.T){

	fun := []string{"x", "a", "y"}

	tripleRiemannSumIntegral(fun, 2.0, 1.0, 2.0, 1.0, 2.0, 1.0)
}


func Test_Triple_Integral_2(t *testing.T){

	fun := []string{"x", "a", "y", "a", "z"}

	tripleRiemannSumIntegral(fun, 2.0, 1.0, 2.0, 1.0, 2.0, 1.0)
}