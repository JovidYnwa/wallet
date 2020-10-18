package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

	/* 	svc := &wallet.Service{}

	   	svc.ExportToFile("data/export.txt")

	   	svc.ImportFromFile("data/import.txt") */

	/* 	wd, err := os.Getwd()
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}
	   	log.Print(wd)

	   	//changin deriction
	   	err = os.Chdir("..")
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}

	   	wd, err = os.Getwd()
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}
	   	log.Print(wd)

	   	abs, err := filepath.Abs(".")
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}
	   	log.Print(abs) */

	//writig a file
	/* 	err := ioutil.WriteFile("../data/export.txt", []byte("some data"), 0666)
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}
	   	//readig a file
	   	data, err := ioutil.ReadFile("../data/export.txt")
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}
	   	log.Print(data) */

	//Way one of copying
	/* 	src, err := os.Open("../data/export.txt")
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}

	   	defer func() {
	   		if cerr := src.Close(); cerr != nil {
	   			log.Print(cerr)
	   		}
	   	}()

	   	stats, err := src.Stat()
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}

	   	dst, err := os.Create("../data/copy.txt")
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}
	   	defer func() {
	   		if cerr := dst.Close(); cerr != nil {
	   			log.Print(cerr)
	   		}
	   	}()
	   	written, err := io.Copy(dst, src)
	   	if err != nil {
	   		log.Print(err)
	   		return
	   	}

	   	if written != stats.Size() {
	   		log.Print(fmt.Errorf("copied size: %d, original size: %d", written, stats.Size()))
	   		return
		   } */

	// calling fuction CopyFile
	/* 	err := CopyFile("../data/export.txt", "../data/copy.txt")
	   	if err != nil {
	   		log.Print(err)
		   } */

	//reading the file
	src, err := os.Open("../data/export.txt")
	if err != nil {
		log.Print(err)
		return
	}
	defer func() {
		if cerr := src.Close(); cerr != nil {
			log.Print(cerr)
		}
	}()
	reader := bufio.NewReader(src)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Print(line)
			break
		}
		if err != nil {
			log.Print(err)
			return
		}
		log.Print(line)
	}
	//[]Account{{Phone: "9010001000"}, {Phone: "9020002000"}}
	//inst1 := wallet.Service{}
	//fmt.Print(inst1.ExperMy())
	//inst2 := wallet.Service{accounts: []*types.Account{{Phone: "9010001000"}}}
	//inst3 := wallet.Service{Accounts: []*types.Account{{Phone: "9010001000"}, {Phone: "9020002000"}, {Phone: "9030003000"}}}
	//fmt.Print(inst3.ImportFromFile("../data/accounts.txt"))

}

/* func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Print(err)
	}
}
*/

//CopyFile копирует из  расположения from в to.
func CopyFile(from, to string) (err error) {
	src, err := os.Open(from)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := src.Close(); cerr != nil {
			if err == nil {
				err = cerr
			}
		}
	}()
	stats, err := src.Stat()
	if err != nil {
		return err
	}
	dst, err := os.Create(to)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := src.Close(); cerr != nil {
			log.Print(cerr)
		}
	}()

	written, err := io.Copy(dst, src)
	if err != nil {
		return err
	}

	if written != stats.Size() {
		return fmt.Errorf("copied size %d, original size %d", written, stats.Size())
	}
	return nil
}
