package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var place_kinds = []string {
	"Tavern", "Inn", "General Store", "Fletchery Store", "Butcher Shop", "Bath House", "Open-air Market",
	"Temple", "Fishing Goods", "Post Office", "Owlry", "Falconer", "Herb Shop", "Arboretum", "Petting Zoo",
	"Botanical Garden", "Cured Meats Store", "Guildhall", "Stable", "Docks", "Apothecary", "Blacksmith",
	"Leatherworks", "Book Store", "Magick Shoppe", "Jewelery", "Candlestick Maker", "Bakery", "Jail",
	"Oddities", "Cobbler", "Pier", "Cartographers Guildhall", "Softies", "Fortuitous Froths", "Trinkets",
	"Rock Candy", "Fur-ever Homes", "Mountaineers Bandoleer", "Bloodletting Shop", "Morgue", "Taxidermists",
	"Tinkers' Construct", "Hatter", "Theater", "Restaurant", "Curiosities", "Stone Carvings", "Statuary",
	"Menders", "Whispers and Secrets", "Library", "Spirits", "Liquor", "Hunting Goods", "Habidashery",
	"Mercenary Guild", "Weapons Shop", "Architect", "City Hall", "Engineers", "Textiles", "Armory",
}

var place_decorations = []string {
	"Hanging Fishing Tools", "Paintings", "Rotted Wood", "Cobwebs", "Hanging Ribbons", "Live Music",
	"Plants", "Bookshelves", "Rug", "Candles", "Chandelier", "Hat Stand", "Coat Rack", "Lanterns",
	"Drapes", "Flowers",
}

var place_feels = []string {
	"Warm", "Smelly", "Damp", "Cold", "Homey", "Distrusting", "Quiet", "Awkward", "Busy",
	"Dark", "Well-lit", "Welcoming", "Clean", "Dirty", "Dingy", "Bright", "Cramped", "Spacious",
	"Eerie", "Fabulous", "Fair", "Fancy", "Natural", "Unnatural", "Empty", "Hot", "Kooky",
	"Poor", "Wealthy", "Luxurious", "Bare-bones", "Antique", "Rowdy", "Ugly", "Unlit", "Vivid",
}

var place_intrigue = []string {
	"Cellar", "Outhouse", "Basement", "Pathway", "Altar", "Chamber", "Vault behind a painting",
	"Portal behind a fake wall", "Trapdoor in the floor", "Tripwire across back doorway",
	"Pressure plate by the side", "Illusory fireplace", "Hidden blades", "Living suit of armor",
	"Door swings open when knocked upon", "Room of unliving statues warriors",
}

var place_pop_min = 1
var place_pop_max = 10
var place_intrigue_chance = 10 // rand 0 - this value, if 0 then intrigue generates

type Place struct {
	name       string
	kind       string    // tavern, inn, store, etc...
	decoration string
	feel       string    // warm, cold, smelly, ...
	people     []*Person // the first person is what the place is named after:
						 // [Bob, Gracie] -> Bob's Tavern, Gracie is customer
	size       int       // number of people determined
	intrigue   string
}

func (self *Place) String() string {
	builder := fmt.Sprintf("%s | Feels %s, decorated with %s, and %d inside",
		self.name, self.feel, self.decoration, self.size)

	if self.intrigue != "" {
		builder = fmt.Sprintf("%s | A hidden %s", builder, self.intrigue)
	}
	builder += "\n"

	for _, person := range self.people {
		builder = fmt.Sprintf("%s\t%s\n", builder, person.String())
	}

	return builder
}

func Place_new() *Place {
	return Place_new_kind(pick_one(place_kinds))
}

func Place_new_kind(kind string) *Place {
	self := Place{}

	self.size = rand.Intn(place_pop_max) + place_pop_min
	self.people = make([]*Person, self.size)

	var wg sync.WaitGroup
	wg.Add(self.size)
	for i := 0; i < self.size; i++ {
		go func(idx int) {
			self.people[idx] = Person_new()
			wg.Done()
		}(i)
	}

	self.kind       = kind
	self.decoration = pick_one(place_decorations)
	self.feel       = pick_one(place_feels)

	if rand.Int() % place_intrigue_chance != 0 {
		self.intrigue = ""
	} else {
		self.intrigue = pick_one(place_intrigue)
	}

	wg.Wait()
	self.name = self.people[0].name + "'s " + self.kind

	return &self
}
