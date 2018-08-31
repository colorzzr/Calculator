package main

import (	"strconv"
	clc "github.com/TheDemx27/calculus"
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