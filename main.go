package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func main() {
	channel := make(chan int)
	go generateValue(channel)
	go generateValue(channel)
	go generateValue(channel)
	go generateValue(channel)

	// x := <-channel
	// y := <-channel

	// sum := x + y
	var sum int
	var index int

	for num := range channel {
		sum += num
		index++
		if index == 4 {
			close(channel)
		}
	}



	fmt.Println(sum)
}

func generateValue(channel chan int) {
	sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	channel <- randN.Intn(10)

}

// func main() {
// 	greet()
// 	storeData("This is some dummy data!", "dummy-data.txt")

// 	channel := make(chan int)

// 	go storeMoreData(50000, "50000_1.txt", channel)
// 	go storeMoreData(50000, "50000_2.txt", channel)

// 	<-channel
// 	<-channel
// }

// func greet() {
// 	fmt.Println("Hi there!")
// }

// func storeData(storableText string, fileName string) {
// 	file, err := os.OpenFile(fileName,
// 		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
// 		0666,
// 	)

// 	if err != nil {
// 		fmt.Println("Creating the file failed. Exiting.")
// 		return
// 	}

// 	defer file.Close()

// 	_, err = file.WriteString(storableText)

// 	if err != nil {
// 		fmt.Println("Writing to the file failed.")
// 	}
// }

// func storeMoreData(lines int, fileName string, c chan int) {
// 	for i := 0; i < lines; i++ {
// 		text := fmt.Sprintf("Line %v - Dummy Data\n", i)
// 		storeData(text, fileName)
// 	}

// 	fmt.Printf("-- Done storing %v lines of text --\n", lines)
// 	c <- 1
// }
