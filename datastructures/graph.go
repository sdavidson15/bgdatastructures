package datastructures

import "fmt"

type Graph interface {
	AddEdge(Edge) error
	AddNode(Node) error
	GetEdges() Edges
	GetNodes() Nodes
	GetUUID() string
	GetName() string
	IsDirectional() bool
	IsWeighted() bool
}

type BaseGraph struct {
	Structure
	Nodes       Nodes `json:"Nodes,omitempty"`
	Edges       Edges `json:"Edges,omitempty"`
	Weighted    bool  `json:"Weighted,omitempty"`
	Directional bool  `json:"Directional,omitempty"`
	nodeIDs     map[string]struct{}
	edgeIDs     map[string]struct{}
	// Acyclic     bool
}

type Graphs []*Graph

func NewGraph(name *string, weighted bool, directional bool) Graph {
	return &BaseGraph{
		Structure:   *NewStructure(name),
		Nodes:       Nodes{},
		Edges:       Edges{},
		Weighted:    weighted,
		Directional: directional,
	}
}

func (b *BaseGraph) AddEdge(e Edge) error {
	if _, ok := b.edgeIDs[e.GetUUID()]; ok {
		return fmt.Errorf(`cannot add edge to graph, ids must be unique`)
	}
	startNodeID := e.GetStartNodeID()
	if _, ok := b.nodeIDs[startNodeID]; !ok {
		return fmt.Errorf(`cannot add edge to graph, start node does not exist in the graph`)
	}
	endNodeID := e.GetEndNodeID()
	if _, ok := b.nodeIDs[endNodeID]; !ok {
		return fmt.Errorf(`cannot add edge to graph, end node does not exist in the graph`)
	}
	if b.Directional && e.GetDirection() == NoEdgeDirection {
		return fmt.Errorf(`cannot add non-directional edge to a directional graph`)
	}
	if !b.Directional && e.GetDirection() != NoEdgeDirection {
		return fmt.Errorf(`cannot add directional edge to a non-directional graph`)
	}
	if b.Weighted && e.GetWeight() == nil {
		return fmt.Errorf(`cannot add unweighted edge to a weighted graph`)
	}
	if !b.Weighted && e.GetWeight() != nil {
		return fmt.Errorf(`cannot add weighted edge to an unweighted graph`)
	}
	b.Edges = append(b.Edges, e)
	b.edgeIDs[e.GetUUID()] = struct{}{}
	return nil
}

func (b *BaseGraph) AddNode(n Node) error {
	if _, ok := b.nodeIDs[n.GetUUID()]; ok {
		return fmt.Errorf(`cannot add node to graph, ids must be unique`)
	}
	b.Nodes = append(b.Nodes, n)
	b.nodeIDs[n.GetUUID()] = struct{}{}
	return nil
}

func (b *BaseGraph) GetEdges() Edges {
	return b.Edges
}

func (b *BaseGraph) GetNodes() Nodes {
	return b.Nodes
}

func (b *BaseGraph) IsDirectional() bool {
	return b.Directional
}

func (b *BaseGraph) IsWeighted() bool {
	return b.Weighted
}
