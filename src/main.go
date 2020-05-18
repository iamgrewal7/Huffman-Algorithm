package main

import "huffman"
import "io/ioutil"
import "sync"
import "time"
import "fmt"


func main() {
	// Timer used for timing run time
	start := time.Now()
	
	// Read FIle
	data, err := ioutil.ReadFile("/home/iamgrewal/projects/Huffman/example.txt")
	if err != nil{
		panic(err)
	}

	// To make sure all coroutines are completed before program terminates add wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go huffman.Compress(string(data), 1, &wg)
	wg.Wait()

	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Printf("Finished in: ")
	fmt.Println(elapsed)
	
}