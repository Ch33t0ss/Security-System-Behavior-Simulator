package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SentinelGuard struct {
	Name    string
	Stamina int
}

func NewSentinelGuard(name string, stamina int) *SentinelGuard {
	return &SentinelGuard{
		Name:    name,
		Stamina: stamina,
	}
}

func (g *SentinelGuard) Patrol(night, intruder bool) string {
	if g.Stamina <= 0 {
		return fmt.Sprintf("%s устал и не может патрулировать", g.Name)
	}
	if intruder {
		g.Stamina -= 15
		if night {
			return fmt.Sprintf("%s обнаружил и поймал нарушителя ночью (вынссливость %d)", g.Name, g.Stamina)
		}
		return fmt.Sprintf("%s обнаружил и поймал нарушителя (выносливость %d)", g.Name, g.Stamina)
	}
	g.Stamina -= 1
	return fmt.Sprintf("%s сообщения: всё чисто (выносливость %d)", g.Name, g.Stamina)
}

func runSentinel() {
	guard := NewSentinelGuard("Мохаммад", 75)
	night := time.Now().Hour() >= 20 || time.Now().Hour() < 6
	intruder := rand.Intn(100) < 30

	status := guard.Patrol(night, intruder)
	fmt.Println(status)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	runSentinel()
}
