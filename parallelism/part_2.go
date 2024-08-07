package parallelism

import "sync"

func task_1() {
	done := make(chan struct{})
	go func() {
		//work()
		close(done)
	}()
	<-done
}

func task_2() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//work()
		}()
	}
	wg.Wait()
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		select {
		case x := <-firstChan:
			output <- x * x
		case x := <-secondChan:
			output <- x * 3
		case <-stopChan:
			return
		}
	}()
	return output
}

func calculator2(arguments <-chan int, done <-chan struct{}) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		sum := 0
		for {
			select {
			case x := <-arguments:
				sum += x
			case <-done:
				output <- sum
				return
			}
		}
	}()
	return output
}
