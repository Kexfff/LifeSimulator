package main

import (
	"fmt"
	"math/rand"
	"time"
)

type human struct {
	age     int
	sex     int
	power   int
	HasBred bool
}

func (s human) PrintHuman() {
	fmt.Print("Age: ", s.age, " Sex: ", s.sex, " Has bred: ", s.HasBred)
	fmt.Println("")
}

func remove(slice []human, s int) []human {
	return append(slice[:s], slice[s+1:]...)
}

func (self *human) check() {
	self.HasBred = false
	self.age = self.age + 10
	if self.age >= 10 && self.age <= 50 {
		rand.Seed(time.Now().UnixNano())
		self.power = self.power + rand.Intn(10)
	} else if self.age < 10 {
		rand.Seed(time.Now().UnixNano())
		self.power = self.power + rand.Intn(5)
	} else {
		rand.Seed(time.Now().UnixNano())
		self.power = self.power - rand.Intn(5)
	}

}

func breed(humans *[]human) {
	time.Sleep(20 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	var r = rand.Intn(2)
	time.Sleep(20 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	var p = rand.Intn(5)
	*humans = append(*humans, human{0, r, p, false})
}

func (s *human) CheckForBreeding(h *human, humans *[]human) {
	if (s.age > 10 && s.age < 50) && (h.age > 10 && h.age < 50) && (s.sex != h.sex) && (s.HasBred == false && h.HasBred == false) {
		s.HasBred = true
		h.HasBred = true
		breed(humans)
	}
}

func main() {
	var Humans []human
	breed(&Humans)
	breed(&Humans)
	breed(&Humans)
	breed(&Humans)
	breed(&Humans)
	breed(&Humans)
	breed(&Humans)

	for i := 0; i < 4; i++ {
		for j := 0; j < len(Humans); j++ {
			Humans[j].check()
		}
		for j := 0; j < len(Humans); j++ {
			for k := 0; k < len(Humans); k++ {
				Humans[j].CheckForBreeding(&Humans[k], &Humans)
			}
		}
	}
	for j := 0; j < len(Humans); j++ {
		Humans[j].PrintHuman()
	}

}
