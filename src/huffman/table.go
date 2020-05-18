package huffman


func tableToString(table map[byte]string) string{
	res := ""
	for key, value := range table {
		res += string(key) + value + " "
	}
	return res
}


func buildTable(hmt *Tree, table map[rune]string, bits string) {
	if hmt == nil {
		return
	}
	buildTable(hmt.LeftNode, table, bits + "0")
	if hmt.LeftNode == nil && hmt.RightNode == nil {
		table[hmt.Char] = bits
	}
	buildTable(hmt.RightNode, table, bits + "1")
}


func getFrequencyMap(text string) map[rune]int {
	frequencyMap := make(map[rune]int)
	for _, char := range text {
		frequencyMap[char]++
	}
	return frequencyMap
}
