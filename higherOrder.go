package main

import (
	"fmt"
	"math"
	"strings"
	"strconv"
	"math/cmplx"
)

type variableInfo struct {
	index float64
	order int64
}

// for currently, it only solve the 2nd and 3rd order
func higherOrderCalc(inputOp []string) []interface{}{

	paraInOrder, highestOrder := flteringOrder(inputOp);

	var ans []interface{}
	//then using formulas to solve them
	if highestOrder == 1{
		fmt.Println("-----Order 1------")
		ans = order1func(paraInOrder[:highestOrder + 1])
	}else if highestOrder == 2{
		fmt.Println("-----Order 2------")
		ans = order2func(paraInOrder[:highestOrder + 1])
	}else if highestOrder == 3{
		fmt.Println("-----Order 3------")
		ans = order3func(paraInOrder[:highestOrder + 1])
	}else {
		fmt.Println("-----Order 4------")

	}

	return ans
}

// input op supposes to be like ["2x2", "+" ,"1"] -> 2x^2 + 1 = 0
func flteringOrder(inputOp []string)([5]float64, int64){
	numArr := []string{inputOp[0]};
	var opArr []string

	//odd position is operation, even is variable
	for i := 1; i < len(inputOp); i++{
		if i % 2 == 1{
			opArr = append(opArr, inputOp[i])
		}else{
			numArr = append(numArr, inputOp[i])
		}
	}

	fmt.Println("opArr:", opArr)
	fmt.Println("numArr:", numArr)


	var indexParaArr = [5]float64{0}
	var highestOrder = int64(0);

	//tackle the first one
	indexPara, order  := splitindexAndOrder(numArr[0]);
	indexParaArr[order] = indexParaArr[order] + indexPara
	if order > highestOrder{
		highestOrder = order;
	}

	//now combine all term together!
	for i := 0; i < len(opArr); i++{
		indexPara, order  := splitindexAndOrder(numArr[i + 1])

		//check sign
		if opArr[i] == "a"{
			indexParaArr[order] = indexParaArr[order] + indexPara
		}else {
			indexParaArr[order] = indexParaArr[order] - indexPara
		}

		if order > highestOrder{
			highestOrder = order;
		}

	}

	fmt.Println(indexParaArr);


	return indexParaArr, highestOrder
}

//the function would return the index and the order of a varible "2x2" -> {index:2, order:2}
func splitindexAndOrder(varible string)(float64, int64){
	//find the place of x
	posOfx := strings.Index(varible, "x");
	//fmt.Println(posOfx);

	if posOfx == -1{
		indexPara,_ := strconv.ParseFloat(varible, 64)
		return indexPara, 0
	}

	//then the thing before is index, after is order
	indexPara,_ := strconv.ParseFloat(varible[:posOfx], 64)
	order, _ := strconv.ParseInt(varible[posOfx + 1:], 10, 64)
	//fmt.Println(indexPara, order);

	return indexPara, order;
}


//assume now is ax + b = 0
func order1func(paraInOrder []float64)[]interface{}{
	fmt.Println(paraInOrder)
	return []interface{}{(-paraInOrder[0]/ paraInOrder[1])}
}

func order2func(paraInOrder []float64)[]interface{}{
	a, b, c := paraInOrder[2], paraInOrder[1], paraInOrder[0]
	bqrm4ac := math.Pow(b, 2) - 4*a*c
	var ans1, ans2 interface{}
	if bqrm4ac >= 0{
		ans1 = (- b + math.Sqrt(bqrm4ac))/(2*a)
		ans2 = (- b - math.Sqrt(bqrm4ac))/(2*a)
	}else{
		ans1 = complex(- b/(2*a) , math.Sqrt(-bqrm4ac)/(2*a))
		ans2 = complex(- b/(2*a) , math.Sqrt(-bqrm4ac)/(2*a))
	}

	return []interface{}{ans1, ans2}
}

func order3func(paraInOrder []float64)[]interface{}{
	a, b, c, d := paraInOrder[3], paraInOrder[2], paraInOrder[1], paraInOrder[0]
	equ1 := 36*a*b*c - 8*math.Pow(b, 3) - 108*math.Pow(a, 2)*d
	equ2 := 12*a*c - 4*math.Pow(b, 2)
	delta := math.Pow(equ1, 2) + math.Pow(equ2, 3)

	//split by delta cos the cmplx.cbrt(-real) may not give the real answer it would take img answer
	var equ3, equ4 complex128
	if delta >= 0{
		equ3 = complex(math.Cbrt(equ1 + math.Sqrt(delta)), 0)
		equ4 = complex(math.Cbrt(equ1 - math.Sqrt(delta)), 0)
	}else{
		equ3 = cmplx.Pow(complex(equ1, 0) + cmplx.Sqrt(complex(delta, 0)), 1.0/3.0)
		equ4 = cmplx.Pow(complex(equ1, 0) - cmplx.Sqrt(complex(delta, 0)), 1.0/3.0)
	}

	x1 := (complex(-2*b, 0) + equ3 + equ4)/complex(6*a, 0)

	para1 := complex(-1, math.Sqrt(3))/complex(12*a, 0)
	para2 := complex(-1, -math.Sqrt(3))/complex(12*a, 0)

	x2 := complex(-b/(3*a), 0) + para1*equ3 + para2*equ4
	x3 := complex(-b/(3*a), 0) + para2*equ3 + para1*equ4

	return []interface{}{x1, x2, x3}
}