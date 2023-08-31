package domain

type Segment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SegmentReqJSON struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	UserPercent int    `json:"user_percent"`
}

func NewSegment(name string) *Segment {
	return &Segment{
		Name: name,
	}
}
