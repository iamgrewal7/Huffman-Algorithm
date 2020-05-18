package huffman

import (
	"sync"
	"container/heap"
)

// Compress Function
func Compress(text string, id int,  wg *sync.WaitGroup) {
	defer wg.Done()
	nodeList := &NodeList{}
	heap.Init(nodeList)

	frequencyMap := getFrequencyMap(text)
	for char, frequency := range frequencyMap {
		heap.Push(nodeList, newTree(char, frequency, nil, nil))
	}

	for nodeList.Len() > 1 {
		fs, ss := heap.Pop(nodeList).(*Tree), heap.Pop(nodeList).(*Tree)
		heap.Push(nodeList, newTree('$', fs.Frequency+ss.Frequency, fs, ss))
	}

	
	table := make(map[rune]string)
	buildTable(nodeList.Nodes[0], table, "")

	encodedTree := ""
	getTree(nodeList.Nodes[0], &encodedTree)
	
	encodedData := ""
	for _, char := range text {
		encodedData += table[char]
	}

	encodedData += "|" + encodedTree
	

	writeToFile(encodedData, id)
}