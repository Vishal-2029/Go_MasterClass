package main

import "fmt"

var Count int
type Node struct {
	Key   int
	Left  *Node
	Right *Node 
}

func (n *Node) Insert (K int){
	if n == nil{
		return 
	}

	if K < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: K}
		}
		n.Left.Insert(K)
	} else if K > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: K}
		}
		n.Right.Insert(K)
	}
}

func (n *Node) Search(K int) bool {

	Count++

	if n == nil {
		return false
	}

	if K < n.Key {
		return n.Left.Search(K)
	}else if K > n.Key {
		return n.Right.Search(K)
	}
		return true
}

func (n *Node) Delete(K int) *Node{

	if n == nil {
		return nil
	}

	if K < n.Key {
		n.Left = n.Left.Delete(K)
	} else if K > n.Key {
		n.Right = n.Right.Delete(K)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}
		minNode := n.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		n.Key = minNode.Key
		n.Right = n.Right.Delete(minNode.Key)
	}
	return n
}

func main() {
	tree := Node{Key: 100}
	tree.Insert(50)
	tree.Insert(40)
	tree.Insert(70)
	tree.Insert(110)
	tree.Insert(150)

	fmt.Println(tree)
	fmt.Println(tree.Search(50)) 
	fmt.Println(tree.Delete(50))
	fmt.Println(tree.Search(50)) 
	fmt.Println(Count)
}