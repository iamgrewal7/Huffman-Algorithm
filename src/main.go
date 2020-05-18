package main

import "huffman"
import "io/ioutil"
import "sync"
import "time"
import "fmt"



func main() {
	start := time.Now()
	data, err := ioutil.ReadFile("/home/iamgrewal/projects/Huffman/example.txt")
	if err != nil{
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	huffman.Compress(string(data), 1, &wg)
	// wg.Add(1)
	// huffman.Compress(string(data), 2, &wg)
	wg.Wait()
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
	
}