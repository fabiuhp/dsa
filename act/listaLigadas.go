package main

type Node struct {
	data string
	next *Node
}

func main() {
	nodes := Node{
		data: "Fabio",
		next: &Node{
			data: "Fabio",
			next: &Node{
				data: "Pereira",
				next: nil,
			},
		},
	}

	_ = nodes
}
