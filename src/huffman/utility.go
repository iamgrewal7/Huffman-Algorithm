package huffman

import (
	"fmt"
	"io/ioutil"
	"bytes"
	"encoding/binary"
	"strconv"
)

// This function act as iterator
func getNext(text string) (string, rune){
	text = text[1:]	
	return text, rune(text[0])
}

// This function is used for converting
// string data (0, 1 from huffman)to int8 and writing to file
func writeToFile(data string, id int) {
	var binaryArray []int8
	i := 0
	for i + 8 <= len(data) {
		num, err := strconv.ParseInt(data[i:i+8], 2, 8)
		if err == nil{
			binaryArray = append(binaryArray, int8(num))
		}
		i+=8
	}

	if i < len(data){
		num, err := strconv.ParseInt(data[i:], 2, 8)
		if err == nil{
			binaryArray = append(binaryArray, int8(num))
		}
	}

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, binaryArray)
	if err != nil {
		panic(err)
	}

	e := ioutil.WriteFile(fmt.Sprintf("/home/iamgrewal/projects/Huffman/output_%d.bin", id), buf.Bytes(), 0644)
	if e != nil {
		panic(err)
	} 
}