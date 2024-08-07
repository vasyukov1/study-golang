package parallelism

func task(c chan int, N int) {
	c <- N + 1
}

func task2(c chan string, s string) {
	for i := 0; i < 5; i++ {
		c <- s + " "
	}
}

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	defer close(outputStream)
	var prev string
	var isFirst bool = true
	for value := range inputStream {
		if isFirst || prev != value {
			outputStream <- value
			prev = value
			isFirst = false
		}
	}
}
