package datastructures

const (
	NoEdgeDirection      int32 = 0
	ForwardEdgeDirection int32 = 1
	ReverseEdgeDirection int32 = -1
)

type Edge interface {
	GetDirection() int32
	GetEndNodeID() string
	GetStartNodeID() string
	GetUUID() string
	GetName() string
	GetWeight() *float64
}

type Edges []Edge

type BaseEdge struct {
	Structure
	Weight      *float64 `json:"Weight,omitempty"`
	StartNodeID *string  `json:"StartNodeID,omitempty"`
	EndNodeID   *string  `json:"EndNodeID,omitempty"`
	Direction   int32    `json:"Direction,omitempty"`
}

func (b *BaseEdge) GetDirection() int32 {
	return b.Direction
}

func (b *BaseEdge) GetEndNodeID() string {
	return strPtrToStr(b.EndNodeID)
}

func (b *BaseEdge) GetStartNodeID() string {
	return strPtrToStr(b.StartNodeID)
}

func (b *BaseEdge) GetWeight() *float64 {
	return b.Weight
}

func NewEdge(
	name *string, weight *float64, startNodeID *string, endNodeID *string, direction int32,
) Edge {
	return &BaseEdge{
		Structure:   *NewStructure(name),
		Weight:      weight,
		StartNodeID: startNodeID,
		EndNodeID:   endNodeID,
		Direction:   direction,
	}
}
