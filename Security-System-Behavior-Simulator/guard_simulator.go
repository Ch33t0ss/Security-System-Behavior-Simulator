package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Guard struct {
	name    string
	energy  int
	alert   bool
	weather string
}

func (g *Guard) react(threat int) string {
	switch {
	case g.energy < 30:
		if threat > 50 {
			return "Слишком устал, запрашивает подмогу..."
		}
		return "Устал, но наблюдает..."
	case g.alert && threat > 70:
		return "Активирует тревогу и бежит к цели!"
	case g.weather == "rain" && g.energy > 60:
		return "Проверяет камеры из укрытия..."
	default:
		if rand.Intn(100) > 60 {
			return "Двигается вдоль периметра."
		}
		return "Стоит на посту и ждет сигнала."
	}
}

func simulate(g *Guard, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		threat := rand.Intn(100)
		result := g.react(threat)
		if threat > 80 && g.energy > 50 {
			g.energy -= 10
			result += " Энергия падает."
		} else if g.energy < 20 {
			g.energy += 15
			result += " Восстанавливает силы."
		}
		ch <- fmt.Sprintf("[%s] Угроза: %d | %s", g.name, threat, result)
		time.Sleep(300 * time.Millisecond)
	}
}
func runSimulation() {
	guards := []*Guard{
		{"Амир", 70, true, "clear"},
		{"Жалгас", 45, false, "rain"},
		{"Мурад", 25, true, "clear"},
	}
	ch := make(chan string)
	var wg sync.WaitGroup
	for _, g := range guards {
		wg.Add(1)
		go simulate(g, &wg, ch)
	}
	go func() { wg.Wait(); close(ch) }()
	for msg := range ch {
		fmt.Println(msg)
	}
}
