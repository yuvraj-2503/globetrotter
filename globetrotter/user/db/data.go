package db

type UserDetails struct {
	UserId   string `bson:"userId"`
	Username string `bson:"username"`
	Score    int    `bson:"score"`
}
