package service

import (
	"context"
	"globetrotter/game/db"
	"globetrotter/game/models"
	"globetrotter/user/service"
	"math/rand"
)

type GameServiceImpl struct {
	destDB      db.DestinationDB
	userService service.ServerService
}

func NewGameServiceImpl(destDB db.DestinationDB, userService service.ServerService) *GameServiceImpl {
	return &GameServiceImpl{destDB: destDB, userService: userService}
}

func (service *GameServiceImpl) GetRandomQuestion(ctx *context.Context) (*models.Question, error) {
	destination, err := service.destDB.GetRandomDestination(ctx)
	if err != nil {
		return nil, err
	}

	clue := getRandomClue(destination.Clues)
	return &models.Question{
		QuestionId: destination.ID.Hex(),
		Clue:       clue,
		Options:    destination.Options,
	}, nil
}

func getRandomClue(clues []string) string {
	return clues[rand.Intn(len(clues))]
}

func (service *GameServiceImpl) SubmitAnswer(ctx *context.Context, userId string, req *models.AnswerRequest) (*models.Verdict, error) {
	destination, err := service.destDB.GetDestinationByID(ctx, req.QuestionId)
	if err != nil {
		return nil, err
	}

	isCorrect := req.Answer == destination.City
	feedback := "😢"
	if isCorrect {
		feedback = "🎉"
		service.updateScore(ctx, userId, 20)
	}

	funFact := getRandomFunFact(destination.FunFacts)

	return &models.Verdict{
		Feedback: feedback,
		FunFact:  funFact,
	}, nil
}

func (service *GameServiceImpl) updateScore(ctx *context.Context, userId string, score int) {
	go func() {
		ctx := context.Background()
		service.userService.UpdateScore(&ctx, userId, score)
	}()
}

func getRandomFunFact(facts []string) string {
	return facts[rand.Intn(len(facts))]
}
