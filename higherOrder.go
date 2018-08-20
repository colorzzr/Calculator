package main

import (
	"fmt"
	"math"
	"strings"
	"strconv"
)

type variableInfo struct {
	index float64
	order int64
}

// for currently, it only solve the 2nd and 3rd order
func higherOrderCalc(inputOp []string){

	paraInOrder, highestOrder := flteringOrder(inputOp);

	//then using formulas to solve them
	if highestOrder == 1{
		fmt.Println("-----Order 1------")
		ans := order1func(paraInOrder[:highestOrder + 1])
		fmt.Println(ans)
	}else if highestOrder == 2{
		fmt.Println("-----Order 2------")
		ans := order2func(paraInOrder[:highestOrder + 1])
		fmt.Println(ans)
	}else if highestOrder == 3{
		fmt.Println("-----Order 3------")

	}else {
		fmt.Println("-----Order 4------")

	}

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
		if (opArr[i] == "+"){
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
	fmt.Println(posOfx);

	if posOfx == -1{
		indexPara,_ := strconv.ParseFloat(varible, 64)
		return indexPara, 0
	}

	//then the thing before is index, after is order
	indexPara,_ := strconv.ParseFloat(varible[:posOfx], 64)
	order, _ := strconv.ParseInt(varible[posOfx + 1:], 10, 64)
	fmt.Println(indexPara, order);

	return indexPara, order;
}


//assume now is ax + b = 0
func order1func(paraInOrder []float64)float64{
	fmt.Println(paraInOrder)
	return (paraInOrder[0]/ paraInOrder[1]);
}

func spiltParaOfOrder2(opArr []string, numArr []float64)(int64, int64, int64){

	return 1, 2, 3
}

func order2func(paraInOrder []float64)[]float64{
	a, b, c := paraInOrder[0], paraInOrder[1], paraInOrder[2]
	bqrm4ac := math.Pow(b, 2) - 4*a*c
	var ans1, ans2 float64
	if bqrm4ac >= 0{
		ans1 = (-b+bqrm4ac)/(2*a)
		ans2 = (-b-bqrm4ac)/(2*a)
	}

	return []float64{ans1, ans2}
}