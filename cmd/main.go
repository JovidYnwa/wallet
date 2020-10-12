package main

import (
	"github.com/JovidYnwa/wallet/pkg/wallet"
)

func main() {

	//Writing files
	/* 	file1, err := os.Create("../data/message.txt")
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}
	   	defer func() {
	   		if cerr := file1.Close(); cerr != nil {
	   			log.Print(err)
	   		}
	   	}()

	   	_, err = file1.Write([]byte("Hey you!"))
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	} */

	/* 	Relates to homeworks 15
	   	svc := &wallet.Service{}

	   	wallet.RegisterAccount(svc, "+992907013487")
	   	svc.RegisterAccount("+992907013487")
	   	svc.Deposit(1, 10)
	   	account, err := svc.RegisterAccount("+992907013487")

	   	if err != nil {
	   		fmt.Println(*svc)
	   		return
	   	}

	   	err = svc.Deposit(account.ID, 10)
	   	if err != nil {
	   		fmt.Println(err)
	   		return
	   	}

	   	payment, err := svc.Pay(account.ID, 88, "some_category")
	   	findAccount, err := svc.FindAccountByID(2)
	   	findPayment, err := svc.FindPaymentByID("88")
	   	fmt.Println(account)
	   	fmt.Println(account.Balance, account.ID)
	   	fmt.Println(payment.Amount)
	   	fmt.Println(findAccount)
	   	fmt.Println(findPayment)
	*/

	//HOME WORK 15
	//Getting working direcotry

	/* 	wd, err := os.Getwd()
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}
		   log.Print(wd) */

	//Opening file readme.txt
	/* 	file, err := os.Open("../data/readme.txt")
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}
	   	log.Printf("%#v", file) */

	/* 	err = file.Close() //First way of closing the file
	   	if err != nil {
	   		log.Print(err)
	   	} */

	//defer closeFile(file)  //Second way by created function

	//Third way anonymous functions
	/* 	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Print(err)
		}
	}(file) */

	/* 	defer func() {
		err := file.Close()
		if err != nil {
			log.Print(err)
		}
	}() */

	//Reading way one
	/* 	buf := make([]byte, 4096)
	   	read, err := file.Read(buf)
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}
	   	data := string(buf[:read])
	   	log.Print(data) */
	/*
		content := make([]byte, 4)
		buf := make([]byte, 4)
		for {
			read, err := file.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Print(err)
				return
			}
			content = append(content, buf[:read]...)
		} */
	//data := string(content)
	//log.Print(data)

	svc := &wallet.Service{}

	svc.ExportToFile("data/export.txt")

	svc.ImportFromFile("data/import.txt")

}

/* func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Print(err)
	}
}
*/
