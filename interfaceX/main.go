package main

import "fmt"

type Student struct {
	name  string
	age   int
	score int
}

type Study interface {
	learn()
	getScore() int
}

type Play interface {
	listenMusic()
}

type Life interface {
	Study
	Play
}

func (s *Student) listenMusic() {
	fmt.Println(s.name, " is listening music.")
}

func (s *Student) learn() {
	fmt.Printf("%s is learnig...\n", s.name)
	s.score = s.score + 1
	if s.score > 100 {
		s.score = 100
	}
}

func (s *Student) getScore() int {
	return s.score
}

func main() {
	tom := Student{
		name:  "Tom",
		age:   20,
		score: 90,
	}

	tom.learn()
	fmt.Println("Tom's score:", tom.getScore())

	var sp Study
	sp = &tom
	sp.learn()
	fmt.Println("Tom's score:", sp.getScore())

	tom.listenMusic()

	var lifeP Life
	lifeP = &tom
	lifeP.learn()
	lifeP.listenMusic()
}
