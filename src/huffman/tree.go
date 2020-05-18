package huffman

// import "fmt"

// Tree Struct
type Tree struct {
	Char      rune
	Frequency int
	LeftNode  *Tree
	RightNode *Tree
}


func newTree(char rune, frequency int, left, right *Tree) *Tree {
	return &Tree{char, frequency, left, right}
}


// PreOrder function
func getTree(hmt *Tree, finalTree *string) {
	if hmt == nil {
		return
	}
	
	if hmt.LeftNode == nil && hmt.RightNode == nil{
		*finalTree += "1" + string(hmt.Char)
	} else {
		*finalTree += "0"
	}
	getTree(hmt.LeftNode, finalTree)
	getTree(hmt.RightNode, finalTree) 
}


func recoverTree(text string) *Tree{
	if len(text) == 0 {
		return nil
	}
	//fmt.Println(text)

	if text[0] == '1' {
		return &Tree{rune(text[1]), -1, nil, nil}
	}

	root := &Tree{'$', -1, nil, nil}
	root.LeftNode = recoverTree(text[1:])
	text = text[2:]
	root.RightNode = recoverTree(text[1:])
	text = text[2:]
	return root
	
}