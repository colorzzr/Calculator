package main

import (
	"fmt"
	"github.com/kylemcc/parse"
)

func test (){
	fmt.Println("cccc");
}
//"hdswitch", "UpdKbelU7zvtsCCW", "jQAr0Xqhbkw45mSW", "http://localhost:8080/v1"
func ConnectToServer(){
	fmt.Println("------Connect to Sever------");
	//initial the parse sever with appid masterkey
	parse.Initialize("Calculator", "UpdKbelU7zvtsCCW", "jQAr0Xqhbkw45mSW");
	//give the target server
	parse.ServerURL("http://127.0.0.1:8080/v1");
}

func pushAns(ans returnPack){
	err := parse.Create(&ans, false);
	fmt.Println(err);
}