package datastructures

type Node interface {
	GetUUID() string
	GetName() string
}

type Nodes []Node

type BaseNode struct {
	Structure
}

func NewNode(name *string) Node {
	return &BaseNode{
		Structure: *NewStructure(name),
	}
}
