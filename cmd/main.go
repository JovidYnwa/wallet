package main

import (
	"log"
	"sync"
)

func main() {
	//lecture 18
	log.Print("main started")

	wg := sync.WaitGroup{}
	wg.Add(1) //сколько рутин ждем

	sum := 0
	go func() {
		defer wg.Done()
		for i := 0; i < 1_000; i++ {
			sum++
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1_00000; i++ {
			sum++
		}
	}()
	wg.Wait()
	log.Print("main finished")
	//time.Sleep(time.Second * 2)
	log.Print(sum)

}

func action(a int, b int) (int, int) {
	c := a + b
	d := a - b
	return c, d
}
