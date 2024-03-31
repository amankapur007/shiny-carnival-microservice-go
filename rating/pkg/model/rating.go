package model

type RecordID string

type RecordType string

const (
	RecordTypeMovie = RecordType("movie")
)

type UserID string

type RatingValue int

type Rating struct {
	RecordID   string      `json:"recordId"`
	UserID     UserID      `json:"userId"`
	RecordType string      `json:"recordType"`
	Value      RatingValue `json:"value"`
}
