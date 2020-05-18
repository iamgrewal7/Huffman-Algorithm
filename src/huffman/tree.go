package huffman

// Tree Struct
type Tree struct {
	Char      rune
	Frequency int
	LeftNode  *Tree
	RightNode *Tree
}

// Tree Constructor
func newTree(char rune, frequency int, left, right *Tree) *Tree {
	return &Tree{char, frequency, left, right}
}

// This functions does a postorder traversal for saving all the paths
func encodeTree(hmt *Tree, finalTree *string) {
	if hmt == nil {
		return
	}
	
	if hmt.LeftNode == nil && hmt.RightNode == nil{
		*finalTree += "1" + string(hmt.Char)
	} else {
		*finalTree += "0"
	}
	encodeTree(hmt.LeftNode, finalTree)
	encodeTree(hmt.RightNode, finalTree) 
}

// This function will be used for recreating the string
// using from the huffman tree that is saved in file along data
func recoverTree(text string) *Tree{
	if len(text) == 0 {
		return nil
	}

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