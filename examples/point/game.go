package main

import (
	"fmt"
	"sync"
)

type coordinate struct {
	i, j int
}

type direction byte

const (
	UP    direction = 'U'
	DOWN  direction = 'D'
	LEFT  direction = 'L'
	RIGHT direction = 'R'
)

type Game struct {
	point     coordinate
	direction direction
	lock      sync.Mutex
}

func newGame() Game {
	return Game{direction: LEFT,
		point: coordinate{i: 1, j: 2}}
}

func (g *Game) Print() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			if g.point.i == i && g.point.j == j {
				fmt.Print("X")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Game) Step(add func(i int, j int), remove func(i int, j int)) {
	g.lock.Lock()
	defer g.lock.Unlock()

	newPoint := coordinate{i: g.point.i, j: g.point.j}
	switch g.direction {
	case UP:
		newPoint.i = abs(g.point.i+3-1) % 3
	case DOWN:
		newPoint.i = abs(g.point.i+1) % 3
	case LEFT:
		newPoint.j = abs(g.point.j+5-1) % 5
	case RIGHT:
		newPoint.j = abs(g.point.j+1) % 5
	}

	remove(g.point.i, g.point.j)
	g.point = newPoint
	add(g.point.i, g.point.j)
}

func (g *Game) ChangeDirection(newDirection direction) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.direction = newDirection
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
