package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Invite struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Invitee string             `bson:"invitee"`
	Inviter string             `bson:"inviter"`
}

type InviterScoreResponse struct {
	Inviter string `json:"inviter"`
	Score   int    `json:"score"`
}
