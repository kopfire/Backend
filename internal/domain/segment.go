package domain

type Segment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewSegment(name string) *Segment {
	return &Segment{
		Name: name,
	}
}
