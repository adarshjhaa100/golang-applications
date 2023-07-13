package concurrency

import (
	"fmt"
	"sync"
)


func unbufferedChan(){
	wg := sync.WaitGroup{}
	// An unbuffered channel should have a receive running in parallel 
	// once value is send to it. Else it will reach a deadlock
	fmt.Println("Unbuffered channel")

	ch := make(chan int)
	ch1 := make(chan int)

	wg.Add(1)

	// This would lead to a deadlock (there's no receiver to channel)
	go func(){
		ch1 <- 12
		wg.Done()
	}()
	

	for i := 0; i < 100; i++ {
		wg.Add(2)

		go func(){
			fmt.Println(<-ch) // receive from channel
			wg.Done()
		}()

		go func(){
			ch <- i // send to a channel
			wg.Done()
		}()

	} 
	wg.Wait()
}

func sendReceiveOnlyChan() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(2)

		// Receive only channel(<-chan). Can only receive from this channel
		go func(ch <-chan int){
			fmt.Println(<-ch)
			wg.Done()
		}(ch)
		
		// Send only channel(<-). Can only receive from this channel
		go func(ch chan<- int){
			ch <- i
			wg.Done()
		}(ch)
		
		wg.Wait() // wait for all channels to signal back, continue post that

	}

	wg.Wait()
}

func buffChan(){
	// Use case could be when sender and receiver operate at 
	// different frequencies (e.g sender is a fast sensor beaming data, receiver slow)
	ch_buff := make(chan int, 50)
	// ch_buff := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	
	go func(ch <-chan int){
		// for range reads until channel is closed
		// for i := range ch {
		// 	// time.Sleep(time.Millisecond * 1000)
		// 	fmt.Printf("Next element %#v, len:%#v \n", i, len(ch))
		// }
		for i := 0; i <= 20; i++ {
			fmt.Printf("Next element %#v, len:%#v \n", i, len(ch))
		}
		wg.Done()
	}(ch_buff)


	go func(ch chan<- int){
			for i := 0; i <= 40; i++ {
				ch <- i
			}
			defer close(ch_buff) // close channel from sender's end
			wg.Done()
	}(ch_buff)
	
	wg.Wait()
	
}


func selectChan() {

	wg := sync.WaitGroup{}

	valChan := make(chan int)
	sigChan := make(chan struct{}) // 0 byte channel, useful for signals


	wg.Add(2)

	go func(){
		for {
			// will run until either of signals receive values
			select {
				case vl, ok := <-valChan :
					// only work on this until channel open	
					if(ok){
						fmt.Println("Received Val: ", vl, ok)
					}
				// Break out of loop if signal channel is able to receive
				case <-sigChan:
					fmt.Println("Sig Chan")
					wg.Done()
					return
			}
		}
	}()

	go func(){
		
		// At end, close channel of value, and signal sigChan to close
		for i := 0; i < 10; i++ {
			valChan <- i
		}
		
		close(valChan)
		sigChan<-struct{}{}


		wg.Done()

	}()

	wg.Wait()

}


func BuffChannelTst(){
	wg := sync.WaitGroup{}

	ch := make(chan int, 10000)

	wg.Add(2)
	go func(ch <-chan int){
		for i := 0; i < 10000; i++ {
			if(i % 23 == 0){
				val, ok := <-ch
				if(ok){
					fmt.Println(val)
				} else {
					fmt.Println("Not OK")
				}
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int){
		for i := 0; i< 10000; i++ {
			ch <- i
		}
		wg.Done()
	}(ch)

	wg.Wait()
}



func RunChannels(){
	// unbufferedChan()
	// sendReceiveOnlyChan()
	// buffChan()
	// selectChan()
	BuffChannelTst()

}