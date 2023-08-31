package domain

import "time"

type RefUsersSegments struct {
	Id        int        `json:"id"`
	UserId    int        `json:"user_id"`
	SegmentId int        `json:"segment_id"`
	Status    bool       `json:"status"`
	DateAdd   *time.Time `json:"date_add"`
	DateDel   *time.Time `json:"date_del"`
}

func NewRefUsersSegments(userId int, segmentId int) *RefUsersSegments {
	return &RefUsersSegments{
		UserId:    userId,
		SegmentId: segmentId,
	}
}
