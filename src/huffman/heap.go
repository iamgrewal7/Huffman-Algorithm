package huffman

// NodeList Type
type NodeList struct {
	Nodes []*Tree
}


// Len Function
func (nodeList NodeList) Len() int {
	return len(nodeList.Nodes)
}


// Less Function
func (nodeList NodeList) Less(i, j int) bool {
	return nodeList.Nodes[i].Frequency < nodeList.Nodes[j].Frequency
}

// Swap Function
func (nodeList NodeList) Swap(i, j int) {
	nodeList.Nodes[i], nodeList.Nodes[j] = nodeList.Nodes[j], nodeList.Nodes[i]
}

// Push Function
func (nodeList *NodeList) Push(x interface{}) {
	nodeList.Nodes = append(nodeList.Nodes, x.(*Tree))
}


// Pop Function
func (nodeList *NodeList) Pop() interface{} {
	old := nodeList.Nodes
	n := len(old)
	x := old[n-1]
	nodeList.Nodes = old[0:n-1]
	return x
}

