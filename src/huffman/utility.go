package huffman

import (
	"fmt"
	"io/ioutil"
	"bytes"
	"encoding/binary"
	"strconv"
)

func getNext(text string) (string, rune){
	text = text[1:]	
	return text, rune(text[0])
}


func writeToFile(data string, id int) {
	var ns []int8
	i := 0
	for i + 8 <= len(data) {
		n, e := strconv.ParseInt(data[i:i+8], 2, 8)
		if e == nil{
			ns = append(ns, int8(n))
		}
		i+=8
	}

	if i < len(data){
		n, e := strconv.ParseInt(data[i:], 2, 8)
		if e == nil{
			ns = append(ns, int8(n))
		}
	}


	buf := new(bytes.Buffer)
	e2 := binary.Write(buf, binary.LittleEndian, ns)
	if e2 != nil {
		panic(e2)
	}

	err := ioutil.WriteFile(fmt.Sprintf("/home/iamgrewal/projects/Huffman/output_%d.bin", id), buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	} 
}