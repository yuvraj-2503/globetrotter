package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Destination struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	City     string             `bson:"city" json:"city"`
	Country  string             `bson:"country" json:"country"`
	Options  []string           `bson:"options" json:"options"`
	Clues    []string           `bson:"clues" json:"clues"`
	FunFacts []string           `bson:"fun_facts" json:"funFacts"`
	Trivia   []string           `bson:"trivia" json:"trivia"`
}

type Question struct {
	QuestionId string   `json:"questionId"`
	Clue       string   `json:"clue"`
	Options    []string `json:"options"`
}

type AnswerRequest struct {
	QuestionId string `json:"questionId"`
	Answer     string `json:"answer"`
}

type Verdict struct {
	Feedback string `json:"feedback"`
	FunFact  string `json:"funFact"`
}
