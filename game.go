package main

type GameState struct {
	MaxSteps      int64  `json:"max_steps"`
	CurrentStep   int64  `json:"current_step"`
	Boxer1        Boxer  `json:"boxer_1"`
	Boxer2        Boxer  `json:"boxer_2"`
	GameOver      bool   `json:"game_over"`
	CurrentWinner *Boxer `json:"current_winner"`
}

func CreateNewGame(MaxSteps int64, Boxer1 *Boxer, Boxer2 *Boxer) GameState {
	newGame := GameState{
		MaxSteps:    MaxSteps,
		CurrentStep: 1,
		Boxer1:      *Boxer1,
		Boxer2:      *Boxer2,
		GameOver:    false,
	}
	return newGame
}

func (gameState *GameState) Next() {
	gameState.CurrentStep++
	gameState.Boxer1.AttackedBodyPart, gameState.Boxer2.BlockedBodyPart = gameState.Boxer1.Attack(&gameState.Boxer2)
	gameState.Boxer2.AttackedBodyPart, gameState.Boxer1.BlockedBodyPart = gameState.Boxer2.Attack(&gameState.Boxer1)
	gameState.GameOver = !(gameState.Boxer1.Health > 0 && gameState.Boxer2.Health > 0 && gameState.CurrentStep < gameState.MaxSteps)
	if gameState.Boxer1.Health > gameState.Boxer2.Health {
		gameState.CurrentWinner = &gameState.Boxer1
	} else if gameState.Boxer2.Health > gameState.Boxer1.Health {
		gameState.CurrentWinner = &gameState.Boxer2
	} else {
		gameState.CurrentWinner = nil
	}
}
