package models

type Ranking struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Pontuacao int    `json:"pontuacao"`
}
