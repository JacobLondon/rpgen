package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sync"
)

var town_pop_min = 50
var town_pop_max = 12000

var town_names = []string {
	"Yaafmouth", "Yhirora", "Puebury", "Zhuypus", "Freurith", "Xens", "Wrico", "Crale", "Illetin",
	"Ensmore", "Trocester", "Glefrough", "Goiplens", "Kocland", "Pheatol", "Iclance", "Pling", "Vreley",
	"Osaburgh", "Ingpus", "Droungan", "Zhefield", "Glushire", "Yohwell", "Zruuyham", "Strane", "Ezodon",
	"Jico", "Inemond", "Ingdon", "Ikhuiyland", "Chiham", "Juclens", "Xoysea", "Pautgend", "Yhence", "Iheigh",
	"Izhouis", "Eighport", "Olisdale", "Jewood", "Prievale", "Lunby", "Breefield", "Glanard", "Vouis", "Ceka",
	"Ogrens", "Amsas", "Inasburg", "Xutpolis", "Kreuhburgh", "Kifbridge", "Wotin", "Detsas", "Qront", "Haso",
	"Galo", "Inasshire", "Adabury",
}

var town_neara = []string {
	"Ocean", "Mountain Range", "River", "Valley", "Stream", "Lake", "Quarry", "Cave", "Dungeon", "Ruins",
}

var town_sizes = []int {
	200, 1000, 5000, 12000, 25000,
}

var town_size_names = []string {
	"Hamlet", "Village", "Town", "City", "Large City",
}

type Town struct {
	pop    int
	name   string
	places []*Place
	neara  string   // ocean, mountain range, valley, river, ...
	size   string   // village, town, ...
}

func Town_new() *Town {
	self := Town{}

	self.pop = rand.Intn(town_pop_max) + town_pop_min
	for i, size := range town_sizes {
		if self.pop < size {
			self.size = town_size_names[i]
			break
		}
	}

	self.name = pick_one(town_names)
	self.neara = pick_one(town_neara)

	fit := math.Log10(float64(self.pop))
	fit *= fit
	self.places = make([]*Place, int(fit))
	var wg sync.WaitGroup

	wg.Add(len(self.places))
	for i := 0; i < len(self.places); i++ {
		go func(idx int) {
			self.places[idx] = Place_new()
			wg.Done()
		}(i)
	}

	wg.Wait()
	return &self
}

func (self *Town) String() string {
	builder := fmt.Sprintf("%s of %s, population %d | %d places, %s nearby\n",
		self.size, self.name, self.pop, len(self.places), self.neara)
	
	for _, p := range self.places {
		builder += p.String()
	}

	return builder
}

func (self *Town) Write() {
	self.Write_path("out.txt")
}

func (self *Town) Write_path(path string) {
	file, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(self.String())
	if err != nil {
		panic(err)
	}
}
