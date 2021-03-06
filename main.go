package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"math"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"time"
)

type recievePack struct {
	InputOp [] string
	OperatingMode int64
	EquationText string
	Extra interface{}
}

type opeRand struct {
	operation string
	opLevel int64
}

type returnPack struct {
	Answer string
	OperationMode int64
	ErrorMsg string
	Date string
	EquantionText string
}

//----------------------------------tcp html recieve and return---------------------------------
func index_handler(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	} else{
		r.ParseForm();
		fmt.Println("------Start Print Request------")
		fmt.Println("Host: ", r.Host)
		fmt.Println("Method ", r.Method)
		fmt.Println("URL: ", r.URL)
		fmt.Println("Parameter: ", r.URL.Query())
		fmt.Println("Encode: ", r.Form.Encode())
	}
}

func htmlRepson(w http.ResponseWriter,r *http.Request, backinfo returnPack){
	fmt.Println("------Start Print Response------")

	//encode to byte[]
	stringInfoInByte, err := json.Marshal(backinfo)
	//convert byte[] to string
	strConverted := string(stringInfoInByte)

	json.NewEncoder(w).Encode(string(strConverted))
	//check error
	if err != nil{
		fmt.Println("ERROR")
	}
	fmt.Println(stringInfoInByte)
	fmt.Println(strConverted)
}

func recieveData(w http.ResponseWriter,r *http.Request) recievePack {
	//decode package string->byte->struct
	strToByte := []byte(r.FormValue("first"));
	fmt.Println(string(strToByte));
	//convert to struct
	var calPackage recievePack;
	err :=json.Unmarshal(strToByte, &calPackage);
	if err != nil{
		fmt.Println("ERROR", err);
	}

	return calPackage;
}
//-----------------------------------finish tcp-------------------------------------------------



//-----------------------------------finish database-----------------------------------------------

//storting out the parenthesis leave operation
//inside call the stack calculation
func basicCalStack(operation [] string) (float64, string) {
	//filter out the exp
	var err string;
	operation, err = filterExp(operation);
	//if there are problems in the exp return it
	if err != "Good"{
		return 0, err;
	}

	//go through the string to filter the level
	var numSta [] float64;
	var opSta [] opeRand;

	var numOfParenthesis int64 = 0
	for i := 0; i < len(operation); i++{
		numtemp, err := strconv.ParseFloat(operation[i], 64)
		//fmt.Println(numtemp, err)
		//then it is a operataion
		if err != nil {
			//create a new node for oprand
			var newOp opeRand;
			var curOpLevel int64;
			//sort the level
			switch operation[i] {
			case "a":
				curOpLevel = 1 + numOfParenthesis * 3
				//create new node
				newOp.operation = operation[i]
				newOp.opLevel = curOpLevel
				opSta = append(opSta, newOp)
				break;
			case "-":
				curOpLevel = 1 + numOfParenthesis * 3
				//create new node
				newOp.operation = operation[i]
				newOp.opLevel = curOpLevel
				opSta = append(opSta, newOp)
				break;
			case "*":
				curOpLevel = 2 + numOfParenthesis * 3;
				//create new node
				newOp.operation = operation[i];
				newOp.opLevel = curOpLevel;
				opSta = append(opSta, newOp);
				break;
			case "/":
				curOpLevel = 2 + numOfParenthesis * 3;
				//create new node
				newOp.operation = operation[i];
				newOp.opLevel = curOpLevel;

				//error divide by 0
				if operation[i] == "0"{
					return 0, "Cannot Divide By Zero!"
				}

				opSta = append(opSta, newOp)
				break;
			case "(":
				numOfParenthesis++
				break
			case ")":
				numOfParenthesis--
				break
			case "^":
				curOpLevel = 3 + numOfParenthesis * 3
				//create new node
				newOp.operation = operation[i];
				newOp.opLevel = curOpLevel;
				opSta = append(opSta, newOp);
			}
		}else{
			numSta = append(numSta, numtemp);
		}
	}

	//--------------------------------------------error checking----------------------------------------------
	//too much ()
	if numOfParenthesis > 0{
		return 0, "Missing One of Closed Parenthesis!";
	}else if numOfParenthesis < 0{
		return 0, "Missing One of Open Parenthesis!";
	}

	//too much operation
	if len(numSta) != len(opSta) + 1{
		return 0, "Too Much Operations";
	}



	var temp  = opeRand{"eof", -32767};
	opSta = append(opSta, temp);
	numSta = append(numSta, -32767);

	//fmt.Println("OpRand:", opSta);
	//fmt.Println("Number", numSta);

	//then use the stack calculation to get answer
	answer, err := stackCalculation(numSta, opSta);

	if math.IsInf(answer, 0){
		return 0, "Yoooo Congratulation You Got Infinity!";
	}

	return answer, err;
}

//using a stack of numbers and operations to get the answer
func stackCalculation(numSta [] float64, opSta[] opeRand)(float64, string){
	var numNewSta [] float64;
	var opNewSta [] opeRand;
	//push first
	numNewSta = append(numNewSta, numSta[0]);
	var topLevel int64 = 0;
	//looping to calculate
	for i := 0; i < len(opSta); i++{
		curLevel := opSta[i].opLevel;
		//if it is higher level we need check later one
		if curLevel > topLevel{
			numNewSta = append(numNewSta, numSta[i + 1])
			opNewSta = append(opNewSta, opSta[i])
			topLevel = curLevel
			//else do operation until there is no higher
		} else{
			for ; len(numNewSta) != 1 && curLevel <= topLevel;{
				//refresh new level

				//fmt.Println("topLevel:",topLevel, "curLevel:",curLevel);
				//fmt.Println("OpRand:", numNewSta);
				//fmt.Println("Number", opNewSta);

				temp := simpleCal(numNewSta[len(numNewSta)- 2], numNewSta[len(numNewSta) -  1], opNewSta[len(opNewSta)- 1].operation)
				//fmt.Println(numNewSta[len(numNewSta)- 2], numNewSta[len(numNewSta) -  1], opNewSta[len(opNewSta)- 1].operation, "Answer:", temp)

				//pop the two number and a oprand
				numNewSta = numNewSta[:len(numNewSta)- 2]
				opNewSta = opNewSta[:len(opNewSta) - 1]

				//push answer and new things
				numNewSta = append(numNewSta, temp)
				//fmt.Println("New OpRand:", opNewSta);
				//fmt.Println("New Number", numNewSta);
				//fmt.Println("");
				if len(opNewSta) <= 0{
					break
				}
				topLevel = opNewSta[len(opNewSta)- 1].opLevel

			}
			if opSta[i].operation != "eof"{
				//then push the new one
				numNewSta = append(numNewSta, numSta[i + 1])
				opNewSta = append(opNewSta, opSta[i])
				topLevel = curLevel
			}
		}
	}
	//fmt.Println("Now OpRand:", opNewSta);
	//fmt.Println("Now Number", numNewSta);

	return numNewSta[0], "Good";
}

//deal with the + - * / ^ return a number
func simpleCal(num1 float64, num2 float64, opRnad string) float64{
	switch opRnad {
	case "a":
		return num1 + num2;
	case "-":
		return num1 - num2;
	case "*":
		return num1 * num2;
	case "/":
		return num1 / num2;
	case "^":
		return math.Pow(num1, num2)
	}
	return 0;
}

//--------------------------------------------Start Imaginary function---------------------------------------------
/********************************************************************
 *	function split the real number, imaginary number and operation	*
 * 	it return a number and code 0, if it is the real number			*
 *	it return a number and code 1, if it is a imaginary numebr		*
 *	it return 0 and code 2, if it is an operation					*
 ********************************************************************/
func splitRealImaOp(input string)(float64, int64){
	//first check number or operation
	numtemp, err := strconv.ParseFloat(input, 64)
	//if no erro it is number
	if err == nil{
		return numtemp, 0
	} else {
		//conner case input i
		if(input == "i"){
			return 1, 1
		}

		//then remove the end of string
		input = input[:len(input)-1]
		numtemp1, err1 := strconv.ParseFloat(input, 64)
		//if it is number , then it is imaginary
		if err1 == nil{
			return numtemp1, 1
		} else{
			return 0, 2;
		}
	}
}

func imageCalStack(operation [] string) (complex128, string){
	//go through the string to filter the level
	var numSta [] complex128;
	var opSta [] opeRand;

	var numOfParenthesis int64 = 0;
	for i := 0; i < len(operation); i++{
		numtemp, codeType := splitRealImaOp(operation[i]);
		//then it is a operataion
		if codeType == 2 {
			//create a new node for oprand
			var newOp opeRand;
			var curOpLevel int64;
			//sort the level
			switch operation[i] {
			case "a":
				curOpLevel = 1 + numOfParenthesis * 3;
				//create new node
				newOp.operation = operation[i];
				newOp.opLevel = curOpLevel;
				opSta = append(opSta, newOp);
				break;
			case "-":
				curOpLevel = 1 + numOfParenthesis * 3;
				//create new node
				newOp.operation = operation[i];
				newOp.opLevel = curOpLevel;
				opSta = append(opSta, newOp);
				break;
			case "*":
				curOpLevel = 2 + numOfParenthesis * 3;
				//create new node
				newOp.operation = operation[i];
				newOp.opLevel = curOpLevel;
				opSta = append(opSta, newOp);
				break;
			case "/":
				curOpLevel = 2 + numOfParenthesis * 3;
				//create new node
				newOp.operation = operation[i];
				newOp.opLevel = curOpLevel;

				//error divide by 0
				if numtemp == 0{
					return 0, "Cannot Divide By Zero!";
				}

				opSta = append(opSta, newOp);
				break;
			case "(":
				numOfParenthesis++;
				break;
			case ")":
				numOfParenthesis--;
				break;
			case "^":
				curOpLevel = 3 + numOfParenthesis * 3;
				//create new node
				newOp.operation = operation[i];
				newOp.opLevel = curOpLevel;
				opSta = append(opSta, newOp);
			}

			//then it is a imaginary
		} else if codeType == 1{
			numSta = append(numSta, complex(0, numtemp));
			//else it is a real number
		} else{
			numSta = append(numSta, complex(numtemp, 0));
		}
	}

	//--------------------------------------------error checking----------------------------------------------
	//too much ()
	if numOfParenthesis > 0{
		return 0, "Missing One of Closed Parenthesis!"
	}else if numOfParenthesis < 0{
		return 0, "Missing One of Open Parenthesis!"
	}
	//fmt.Println(len(numSta), len(opSta));
	//too much operation
	if len(numSta) != len(opSta) + 1{
		return 0, "Too Much Operations"
	}



	var temp  = opeRand{"eof", -32767}
	opSta = append(opSta, temp)
	numSta = append(numSta, -32767)

	fmt.Println("OpRand:", opSta)
	fmt.Println("Number", numSta)

	//then use the stack calculation to get answer
	answer, err := stackCalculationIm(numSta, opSta)

	return answer, err
}

func stackCalculationIm(numSta [] complex128, opSta[] opeRand)(complex128, string){
	var numNewSta [] complex128;
	var opNewSta [] opeRand;
	//push first
	numNewSta = append(numNewSta, numSta[0]);
	var topLevel int64 = 0;
	//looping to calculate
	for i := 0; i < len(opSta); i++{
		curLevel := opSta[i].opLevel;
		//if it is higher level we need check later one
		if curLevel > topLevel{
			numNewSta = append(numNewSta, numSta[i + 1]);
			opNewSta = append(opNewSta, opSta[i]);
			topLevel = curLevel;
			//else do operation until there is no higher
		} else{
			for ; len(numNewSta) != 1 && curLevel <= topLevel;{
				//refresh new level

				fmt.Println("topLevel:",topLevel, "curLevel:",curLevel);
				fmt.Println("OpRand:", numNewSta);
				fmt.Println("Number", opNewSta);
				temp := simpleCalIm(numNewSta[len(numNewSta)- 2], numNewSta[len(numNewSta) -  1], opNewSta[len(opNewSta)- 1].operation);
				//fmt.Println(numNewSta[len(numNewSta)- 2], numNewSta[len(numNewSta) -  1], opNewSta[len(opNewSta)- 1].operation, "Answer:", temp)

				//pop the two number and a oprand
				numNewSta = numNewSta[:len(numNewSta)- 2];
				opNewSta = opNewSta[:len(opNewSta) - 1];

				//push answer and new things
				numNewSta = append(numNewSta, temp);
				//fmt.Println("New OpRand:", opNewSta);
				//fmt.Println("New Number", numNewSta);
				//fmt.Println("");
				if len(opNewSta) <= 0{
					break;
				}
				topLevel = opNewSta[len(opNewSta)- 1].opLevel;

			}
			if opSta[i].operation != "eof"{
				//then push the new one
				numNewSta = append(numNewSta, numSta[i + 1]);
				opNewSta = append(opNewSta, opSta[i]);
				topLevel = curLevel;
			}
		}
	}
	fmt.Println("Now OpRand:", opNewSta);
	fmt.Println("Now Number", numNewSta);

	return numNewSta[0], "Good";
}

//deal with the + - * / ^ in complex number return a number
func simpleCalIm(num1 complex128, num2 complex128, opRnad string) complex128{
	switch opRnad {
	case "a":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	}
	return 0;
}

//---------------------------------------------End of imaginary function---------------------------------------------

func filterExp(input [] string) ([]string, string){
	locOfExp := -1
	endOfExp := -1
	numOfParenthesis := 0
	numOfParenthesisBefor := 0
	for i := 0;  i < len(input); i++{
		if input[i] == "exp("{
			locOfExp = i + 1
			//record number of parenthesis before the exp
			numOfParenthesisBefor = numOfParenthesis
			numOfParenthesis++
		} else if input[i] == "("{
			numOfParenthesis++
		} else if input[i] == ")"{
			numOfParenthesis--
			//we found the right closed parenthesis
			if numOfParenthesisBefor == numOfParenthesis{
				endOfExp = i
				break
			}
		}
	}

	//if there is not enough or too much parenthesis
	if numOfParenthesis > 0{
		return nil, "Missing One of Closed Parenthesis!";
	}else if numOfParenthesis < 0{
		return nil, "Missing One of Open Parenthesis!";
	}

	//if not exp end recursion and return original
	if locOfExp == -1{
		return input, "Good"
	}

	fmt.Println(locOfExp, endOfExp)
	fmt.Println(input[locOfExp:endOfExp])
	answer, err := basicCalStack(input[locOfExp:endOfExp])
	//error checking
	if err != "Good"{
		return nil, err
	}
	strAns := strconv.FormatFloat(math.Exp(answer), 'f', 3, 64)

	//appending new array
	var newInput [] string

	for i:=0; i < locOfExp - 1; i++{
		newInput = append(newInput, input[i])
	}
	newInput = append(newInput, strAns)
	for i:=endOfExp + 1; i < len(input); i++{
		newInput = append(newInput, input[i])
	}

	fmt.Println(newInput);
	return newInput, "Good"
}

func calProcess(w http.ResponseWriter,r *http.Request) {
	fmt.Println("------calProcess Ack!------")

	//print request info
	w.Header().Set("Access-Control-Allow-Origin", "*")
	index_handler(w, r)

	//get the data in struct
	recievedData := recieveData(w, r)

	var answerPack returnPack

	//triple integraal calculation
	if recievedData.OperatingMode == 6 {
		XupperBound := interfaceToFloat(recievedData.Extra.(map[string]interface {})["XupperBound"])
		XlowerBound := interfaceToFloat(recievedData.Extra.(map[string]interface {})["XlowerBound"])
		YupperBound := interfaceToFloat(recievedData.Extra.(map[string]interface {})["YupperBound"])
		YlowerBound := interfaceToFloat(recievedData.Extra.(map[string]interface {})["YlowerBound"])
		ZupperBound := interfaceToFloat(recievedData.Extra.(map[string]interface {})["ZupperBound"])
		ZlowerBound := interfaceToFloat(recievedData.Extra.(map[string]interface {})["ZlowerBound"])


		answer, errorMsg := tripleRiemannSumIntegral(recievedData.InputOp, XupperBound, XlowerBound, YupperBound, YlowerBound, ZupperBound, ZlowerBound)
		answerPack.Answer = strconv.FormatFloat(answer, 'f', 2, 64)
		answerPack.ErrorMsg = errorMsg
	//double integraal calculation
	}else if recievedData.OperatingMode == 5 {
		XupperBound := interfaceToFloat(recievedData.Extra.(map[string]interface {})["XupperBound"])
		XlowerBound := interfaceToFloat(recievedData.Extra.(map[string]interface {})["XlowerBound"])
		YupperBound := interfaceToFloat(recievedData.Extra.(map[string]interface {})["YupperBound"])
		YlowerBound := interfaceToFloat(recievedData.Extra.(map[string]interface {})["YlowerBound"])

		answer, errorMsg := doubleRiemannSumIntegral(recievedData.InputOp, XupperBound, XlowerBound, YupperBound, YlowerBound)
		answerPack.Answer = strconv.FormatFloat(answer, 'f', 2, 64)
		answerPack.ErrorMsg = errorMsg


	//single integraal calculation
	}else if recievedData.OperatingMode == 4 {
		upperBound := interfaceToFloat(recievedData.Extra.(map[string]interface {})["XupperBound"])
		lowerBound := interfaceToFloat(recievedData.Extra.(map[string]interface {})["XlowerBound"])

		answer, errorMsg := RiemannSumIntegral(recievedData.InputOp, upperBound, lowerBound)
		answerPack.Answer = strconv.FormatFloat(answer, 'f', 2, 64)
		answerPack.ErrorMsg = errorMsg

	//higher order mode
	}else if recievedData.OperatingMode == 3{
		answer := higherOrderCalc(recievedData.InputOp)
		var ansStr string;
		//format all possible answer
		for i := 0; i < len(answer); i++{
			switch temp := answer[i].(type) {
			case float64:
				ansStr = ansStr + "x" +
					strconv.Itoa(i) + "=" +
					strconv.FormatFloat(temp, 'f', -1, 64) + " "
				break
			case complex128:
				ansStr = ansStr + "x" +
					strconv.Itoa(i) + "=" +
					strconv.FormatFloat(real(temp), 'f', 2, 64) + "+" +
					strconv.FormatFloat(imag(temp), 'f', 2, 64) + "i" + " "
				break
			}

		}

		answerPack.Answer = ansStr
		answerPack.ErrorMsg = "Good"

		//abs mode
	}else if recievedData.OperatingMode == 2{
		answer, errMsg := imageCalStack(recievedData.InputOp)

		//get the mag and phase for abs
		magnitude := math.Sqrt(math.Pow(real(answer), 2) + math.Pow(imag(answer), 2))
		phase := math.Atan(imag(answer)/real(answer))
		//forming the return package aaaas mag:a + phaaaase:b
		answerPack.Answer = "Magnitude:" + strconv.FormatFloat(magnitude, 'f', 2, 64) + "+" +
			"Magnitude:" + strconv.FormatFloat(phase, 'f', 2, 64)
		answerPack.ErrorMsg = errMsg

		//imaginery mode
	}else if recievedData.OperatingMode == 1{
		answer, errMsg := imageCalStack(recievedData.InputOp)
		//forming the return package a + bi
		answerPack.Answer = strconv.FormatFloat(real(answer), 'f', -1, 64) + "+" +
			strconv.FormatFloat(imag(answer), 'f', 2, 64) + "i"
		//if nothing wrong return Good
		answerPack.ErrorMsg = errMsg
		//normal mode
	}else {
		answer, errMsg := basicCalStack(recievedData.InputOp)
		//forming the return package
		answerPack.Answer = strconv.FormatFloat(answer, 'f', 2, 64)
		answerPack.ErrorMsg = errMsg
	}

	//assign the regular term
	answerPack.OperationMode = recievedData.OperatingMode
	y, m, d := time.Now().Date()
	answerPack.Date = strconv.Itoa(y) + "-" + m.String() + "-" + strconv.Itoa(d)
	answerPack.EquantionText = recievedData.EquationText

	//response to the web and store to database
	htmlRepson(w, r, answerPack)
	pushAns(answerPack)
	fmt.Println("-----Finish Data Exchange------")
}


func main() {
	fmt.Println("Welcome To Color's Calculator")
	ConnectToServer();

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/calProcess",calProcess)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}

}
