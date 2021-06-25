package main

import (
	"fmt"
	"os"
	"strings"
)

func sendToAddress(address_inp string,amount_inp string){
	//TODO:check amount from json_file,then check address of receiver(true\false) and update amount of users
	fmt.Println("Transaction done! Do you want create new operation?(Yes/No)")
	var answer string
	fmt.Fscan(os.Stdin,&answer)
	if(answer=="No" || answer=="no"){
		os.Exit(1)
	}else{
		fmt.Println("Input amount = ")
		var amount,address string
		fmt.Fscan(os.Stdin,&amount)

		fmt.Println("Input address: ")
		fmt.Fscan(os.Stdin,&address)
		sendToAddress(amount,address)
	}

}

func main() {
	fmt.Println("Welcome,input your command(for more information input 'help'):")

	var input=""
	var amount,address string
	for strings.TrimRight(input, "\n")!="exit"{
		fmt.Fscan(os.Stdin, &input)
		if(input=="help") {
			fmt.Println("Input sendToAddress if you want do transaction,else input exit to close programm.")
		} else if(input=="exit"){
			os.Exit(1)
		} else if(input=="sendToAddress"){
			fmt.Println("Input amount = ")
			fmt.Fscan(os.Stdin,&amount)

			fmt.Println("Input address: ")
			fmt.Fscan(os.Stdin,&input)
			sendToAddress(amount,address)

		} else{
			fmt.Println("Invalid command! Input 'help' and try again.")
		}
	}

}
