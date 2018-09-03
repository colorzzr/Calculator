package main

import (	"strconv"
	clc "github.com/TheDemx27/calculus"
	"fmt"
	"math"
)

func singleVarIntegral(pack recievePack) (float64, string){
	//forming the function and limit
	f := clc.NewFunc(pack.EquationText)
	upperBound := pack.Extra.(map[string]interface {})["upperBound"]
	lowerBound := pack.Extra.(map[string]interface {})["lowerBound"]

	return f.AntiDiff(interfaceToFloat(lowerBound),interfaceToFloat(upperBound)), "Good"
}

func interfaceToFloat(input interface{}) float64 {
	switch i := input.(type) {
	case string:
		temp, _ := strconv.ParseFloat(i, 64)
		return temp
	default:
		return -32767.32767
	}
}

//----------------------------------------------------------Above is using library------------------------------------------------
func RiemannSumIntegral(inputFun [] string, upperBound float64, lowerBound float64, customLen ...int)(float64, string){
	fmt.Println(inputFun, upperBound, lowerBound, customLen);

	//if user doesnot specify the length for Riemann Sum then use 10000
	var RiemannLen int
	if len(customLen) > 0{
		RiemannLen = customLen[0]
	}else{
		RiemannLen = 50000
	}

	interval := (upperBound - lowerBound) / float64(RiemannLen)

	startPoint := lowerBound
	sum := 0.0
	//looping to get area
	for i := 0; i < RiemannLen; i++{
		//using rectangle Riemann Sum
		//replace the x with startPoint
		//fmt.Println("inputFun:",inputFun)
		curPoint := strconv.FormatFloat(startPoint, 'f', -1, 64)
		replacedFunc := replaceEleInArr(inputFun, "x", curPoint)
		ans, err := basicCalStack(replacedFunc)
		if err != "Good"{
			return 0, err
		}


		//fmt.Println("startPoint:",startPoint)
		//fmt.Println("curPoint:",curPoint)
		//fmt.Println("replacedFunc:",replacedFunc)
		//fmt.Println("ans",ans)

		sum = sum + ans * interval
		startPoint = startPoint + interval
	}

	//round up with 2 decimals
	sum = math.Floor(sum*100 + 0.5) / 100.0
	return sum, "Good"
}


func doubleRiemannSumIntegral(inputFun [] string, xUpperBound float64, xLowerBound float64,
	yUpperBound float64, yLowerBound float64, customLen ...int)(float64, string){
	//if user doesnot specify the length for Riemann Sum then use 10000
	var RiemannLen int
	if len(customLen) > 0{
		RiemannLen = customLen[0]
	}else{
		RiemannLen = 500
	}

	dx := (xUpperBound - xLowerBound) / float64(RiemannLen)
	dy := (yUpperBound - yLowerBound) / float64(RiemannLen)

	fmt.Println("dx:", dx, ",dy:", dy)

	xStartPoint := xLowerBound
	yStartPoint := yLowerBound
	sum := 0.0

	//same looping throught the two bounds
	for i := 0; i < RiemannLen; i++{
		curPoint := strconv.FormatFloat(xStartPoint, 'f', -1, 64)
		replacedFuncX := replaceEleInArr(inputFun, "x", curPoint)

		//ybound
		for j := 0; j < RiemannLen; j++{
			curPoint =	strconv.FormatFloat(yStartPoint, 'f', -1, 64)
			replacedFuncY := replaceEleInArr(replacedFuncX, "y", curPoint)

			ans, err := basicCalStack(replacedFuncY)
			if err != "Good"{
				return 0, err
			}
			sum = sum + ans * dx * dy
			//increament y
			yStartPoint = yStartPoint + dy
		}

		//reset y and increa the x
		yStartPoint = yLowerBound
		xStartPoint = xStartPoint + dx
	}


	//round up with 2 decimals
	sum = math.Floor(sum*100 + 0.5) / 100.0
	return sum, "Good"
}

func tripleRiemannSumIntegral(inputFun [] string,
	xUpperBound float64, xLowerBound float64,
	yUpperBound float64, yLowerBound float64,
	zUpperBound float64, zLowerBound float64,
	customLen ...int)(float64, string){
	//if user doesnot specify the length for Riemann Sum then use 10000
	var RiemannLen int
	if len(customLen) > 0{
		RiemannLen = customLen[0]
	}else{
		RiemannLen = 150
	}

	dx := (xUpperBound - xLowerBound) / float64(RiemannLen)
	dy := (yUpperBound - yLowerBound) / float64(RiemannLen)
	dz := (zUpperBound - zLowerBound) / float64(RiemannLen)

	xStartPoint := xLowerBound
	yStartPoint := yLowerBound
	zStartPoint := zLowerBound
	sum := 0.0

	//same looping throught the two bounds
	for i := 0; i < RiemannLen; i++{
		curPoint := strconv.FormatFloat(xStartPoint, 'f', -1, 64)
		replacedFuncX := replaceEleInArr(inputFun, "x", curPoint)

		//ybound
		for j := 0; j < RiemannLen; j++{
			curPoint =	strconv.FormatFloat(yStartPoint, 'f', -1, 64)
			replacedFuncY := replaceEleInArr(replacedFuncX, "y", curPoint)

			for k := 0; k < RiemannLen; k++{
				curPoint =	strconv.FormatFloat(zStartPoint, 'f', -1, 64)
				replacedFuncZ := replaceEleInArr(replacedFuncY, "z", curPoint)
				ans, err := basicCalStack(replacedFuncZ)
				if err != "Good"{
					return 0, err
				}
				sum = sum + ans * dx * dy * dz
				zStartPoint = zStartPoint + dz
			}
			zStartPoint = zLowerBound
			//increament y
			yStartPoint = yStartPoint + dy
		}

		//reset y && z and increa the x
		zStartPoint = zLowerBound
		yStartPoint = yLowerBound
		xStartPoint = xStartPoint + dx
	}

	//round up with 2 decimals
	sum = math.Floor(sum*100 + 0.5) / 100.0
	return sum, "Good"
}

func replaceEleInArr(input []string, old string, new string) []string{
	//deep copy with array
	temp := make([]string, len(input))
	copy(temp, input)
	for i := 0; i < len(temp); i++{
		if temp[i] == old{
			temp[i] = new
		}
	}

	return temp
}