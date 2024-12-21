package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	Text    string
	Answer  int
	Options []string
}

type GameState struct {
	Name      string
	Points    int
	Questions []Question
}

func (g *GameState) Init() {
	fmt.Println("Iniciando o jogo...")
	fmt.Println("Como você se chama?")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Não foi possível ler o nome")
	}

	g.Name = name

	fmt.Println("Olá" + "," + g.Name + "Vamos começar o jogo!")
}

func (g *GameState) ProcessCSV() {
	file, err := os.Open("quizgo.csv")

	if err != nil {
		panic("Não foi possível abrir o arquivo")
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		panic("Não foi possível ler o arquivo")
	}

	for index, record := range records {
		if index > 0 {
			correctAnswer, _ := toInt(record[5])
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func (g *GameState) Run() {
	for index, question := range g.Questions {
		fmt.Printf("Pergunta %d: %s\n", index+1, question.Text)
		for index, option := range question.Options {
			fmt.Printf("[%d]: %s\n", index+1, option)
		}

		fmt.Println("Digite uma alternativa:")
		var answer int
		var err error

		for {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')

			answer, err = toInt(read[:len(read)-1])

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break
		}
		if answer == question.Answer {
			fmt.Println("Resposta correta!")
			g.Points += 10
		} else {
			fmt.Println("Resposta incorreta!")
		}
	}
}

func main() {
	game := &GameState{}
	go game.ProcessCSV()
	game.Init()
	game.Run()

	fmt.Printf("Parabéns %s, você fez %d pontos\n", game.Name, game.Points)
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(strings.TrimSpace(s))

	if err != nil {
		return 0, errors.New("não é permitido caractere diferente de número. por favor digite um número")
	}

	return i, nil
}
