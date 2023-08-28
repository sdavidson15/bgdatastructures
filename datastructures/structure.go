package datastructures

import (
	"github.com/google/uuid"
)

type Structure struct {
	UUID string  `json:"UUID"`
	Name *string `json:"Name,omitempty"`
}

func (s *Structure) GetUUID() string {
	return s.UUID
}

func (s *Structure) GetName() string {
	return strPtrToStr(s.Name)
}

func NewStructure(name *string) *Structure {
	return &Structure{
		UUID: uuid.NewString(),
		Name: name,
	}
}
