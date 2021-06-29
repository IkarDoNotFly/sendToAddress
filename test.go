package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
type Owner struct{
	my_amount float64 `json:"my_amount"`

}
type Accounts struct{
	Accounts 	[]Account 	`json:"Accounts"`
}

type Account struct{
	amount float64 `json:"amount"`
	owner_address string `json:"owner_address"`
}



func sendToAddress(address_inp string,amount_inp float64){
	jsonFile,err:=os.Open("accounts.json")

	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println("Successfully Opened accounts.json")
	}

	var accounts Accounts
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue,&accounts)

	for i:=0; i<len(accounts.Accounts);i++{
		if Owner.my_amount<amount_inp && address_inp==accounts.Accounts.owner_address[i] {
			//I don`t know how to work with it correctly(may be we will discuss it on interview)
			Owner.my_amount-=amount_inp
			accounts.Accounts.amount[i]+=amount_inp
		}else{
			fmt.Println("Sorry,we can`t do this.\n Your amount isn`t enough or address of recipient wrong!")
		}
	}


	defer jsonFile.Close()


	fmt.Println("Transaction done! Do you want create new operation?(Yes/No)")


	var answer string
	fmt.Fscan(os.Stdin,&answer)
	if(answer=="No" || answer=="no"){
		os.Exit(1)
	}else{
		inputData()
	}

}

func inputData(){
	var input=""
	var address string
	var amount float64
	fmt.Println("Input amount = ")
	fmt.Fscan(os.Stdin,&amount)

	fmt.Println("Input address: ")
	fmt.Fscan(os.Stdin,&input)
	sendToAddress(address,amount)
}

func main() {
	fmt.Println("Welcome,input your command(for more information input 'help'):")

	var input=""

	for strings.TrimRight(input, "\n")!="exit"{
		fmt.Fscan(os.Stdin, &input)
		if(input=="help") {
			fmt.Println("Input sendToAddress if you want do transaction,else input exit to close programm.")
		} else if(input=="exit"){
			os.Exit(1)
		} else if(input=="sendToAddress"){
			inputData()

		} else{
			fmt.Println("Invalid command! Input 'help' and try again.")
		}
	}

}
