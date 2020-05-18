package huffman
import "fmt"

// Decompress Function
func Decompress(text string) {
	root := recoverTree(&text)
	fmt.Println(root)
}