package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Race struct {
	male_names   []string
	female_names []string
	surnames     []string
	skin_colors  []string
	hair_colors  []string
	eye_colors   []string
	height_min   float32
	height_max   float32
	weight_min   int
	weight_max   int
	age_min      int
	age_max      int
}

var race_defs = map[string]Race{
	"Human": {
		male_names: []string{
			"Aseir", "Bardeid", "Haseid", "Khemed", "Mehmen", "Sudeiman", "Zasheir",
			"Darvin", "Dorn", "Evendur", "Gorstag", "Grim", "Helm", "Malark", "Morn", "Randal", "Stedd",
			"Bor", "Fodel", "Glar", "Grigor", "Igan", "Ivor", "Kosef", "Mival", "Orel", "Pavel", "Sergor",
			"Ander", "Blath", "Bran", "Frath", "Geth", "Lander", "Luth", "Malcer", "Stor", "Taman", "Urth",
			"Aoth", "Bareris", "Ehput-Ki", "Kethoth", "Mumed", "Ramas", "So-Kehur", "Thazar-De", "Urhur",
			"Borivik", "Faurgar", "Jandar", "Kanithar", "Madislak", "Ralmevik", "Shaumar", "Vladislak",
			"An", "Chen", "Chi", "Fai", "Jiang", "Jun", "Lian", "Long", "Meng", "On", "Shan", "Shui", "Wen",
			"Anton", "Diero", "Marcon", "Pieron", "Rimardo", "Romero", "Salazar", "Umbero",
		},
		female_names: []string{
			"Atala", "Ceidil", "Hama", "Jasmal", "Meilil", "Seipora", "Yasheira", "Zasheida",
			"Arveene", "Esvele", "Jhessail", "Kerri", "Lureene", "Miri", "Rowan", "Shandri", "Tessele",
			"Alethra", "Kara", "Katernin", "Mara", "Natali", "Olma", "Tana", "Zora",
			"Amafrey", "Betha", "Cefrey", "Kethra", "Mara", "Olga", "Silifrey", "Westra",
			"Arizima", "Chathi", "Nephis", "Nulara", "Murithi", "Sefris", "Thola", "Umara", "Zolis",
			"Fyevarra", "Hulmarra", "Immith", "Imzel", "Navarra", "Shevarra", "Tammith", "Yuldra",
			"Bai", "Chao", "Jia", "Lei", "Mei", "Qiao", "Shui", "Tai",
			"Balama", "Dona", "Faila", "Jalana", "Luisa", "Marta", "Quara", "Selise", "Vonda",
		},
		surnames: []string{
			"Basha", "Dumein", "Jassan", "Khalid", "Mostana", "Pashar", "Rein",
			"Amblecrown", "Buckman", "Dundragon", "Evenwood", "Greycastle", "Tallstag",
			"Bersk", "Chernin", "Dotsk", "Kulenov", "Marsk", "Nemetsk", "Shemov", "Starag",
			"Brightwood", "Helder", "Hornraven", "Lackman", "Stormwind", "Windrivver",
			"Ankhalab", "Anskuld", "Fezim", "Hahpet", "Nathandem", "Sepret", "Uuthrakt",
			"Chergoba", "Dyernina", "Iltazyara", "Murnyethara", "Stayanoga", "Ulmokina",
			"Chien", "Huang", "Kao", "Kung", "Lao", "Ling", "Mei", "Pin", "Shin", "Sum", "Tan", "Wan",
			"Agosto", "Astorio", "Calabra", "Domine", "Falone", "Marivaldi", "Pisacar", "Ramondo",
		},
		skin_colors: []string{
			"Light", "Dark", "Tan", "Freckled",
		},
		hair_colors: []string{
			"Black", "Brown", "Blonde", "Red", "White", "Gray", "Bald",
		},
		eye_colors: []string{
			"Brown", "Blue", "Green", "Hazel", "Dark",
		},
		height_min: 4.0,
		height_max: 7.0,
		weight_min: 90,
		weight_max: 300,
		age_min: 15,
		age_max: 90,
	},
	"Elf": {
		male_names: []string{
			"Adran", "Aelar", "Aramil", "Arannis", "Aust", "Beiro", "Berrian", "Carric", "Enialis",
			"Erdan", "Erevan", "Galinndan", "Hadarai", "Heian", "Himo", "Immeral", "Ivellios", "Laucian",
			"Mindartis", "Paelias", "Peren", "Quarion", "Riardon", "Rolen", "Soveliss", "Thamior",
			"Tharivol", "Theren", "Varis",
		},
		female_names: []string{
			"Adrie", "Althaea", "Anastrianna", "Andraste", "Antinua", "Bethrynna", "Birel", "Caelynn",
			"Drusilia", "Enna", "Felosial", "Ielenia", "Jelenneth", "Keyleth", "Leshanna", "Lia", "Meriele",
			"Mialee", "Naivara", "Quelenna", "Quillathe", "Sariel", "Shanairra", "Shava", "Silaqui",
			"Theirastra", "Thia", "Vadania", "Valanthe", "Xanaphia",
		},
		surnames: []string{
			"Amakiir (Gemflower)", "Amastacia (Starflower)", "Galanodel (Moonwhisper)", "Holimion (Diamonddew)",
			"Ilphelkiir (Gemblossom)", "Liadon (Silverfrond)", "Meliamne (Oakenheel)", "Naïlo (Nightbreeze)",
			"Siannodel (Moonbrook)", "Xiloscient (Goldpetal)",
		},
		skin_colors: []string{
			"Light", "Dark",
		},
		hair_colors: []string{
			"Black", "Silver", "Brown", "Gold",
		},
		eye_colors: []string{
			"Brown", "Light", "Green", "Blue",
		},
		height_min: 4.5,
		height_max: 6.2,
		weight_min: 70,
		weight_max: 200,
		age_min: 15,
		age_max: 800,
	},
	"Orc": {
		male_names: []string{
			"Dench", "Feng", "Gell", "Henk", "Holg", "Imsh", "Keth", "Krusk",
			"Mhurren", "Ront", "Shump", "Thokk",
		},
		female_names: []string{
			"Baggi", "Emen", "Engong", "Kansif", "Myev", "Neega", "Ovak", "Ownka",
			"Shautha", "Sutha", "Vola", "Volen", "Yevelda",
		},
		surnames: []string{
			"",
		},
		skin_colors: []string{
			"Green", "Gray", "Dark",
		},
		hair_colors: []string{
			"Black", "Bald", "Brown",
		},
		eye_colors: []string{
			"Black", "Dark", "Brown", "Green",
		},
		height_min: 5.8,
		height_max: 8.0,
		weight_min: 170,
		weight_max: 300,
		age_min: 15,
		age_max: 90,
	},
	"Dwarf": {
		male_names: []string{
			"Adrik", "Alberich", "Baern", "Barendd", "Brottor", "Bruenor", "Dain", "Darrak", "Delg", "Eberk",
			"Einkil", "Fargrim", "Flint", "Gardain", "Harbek", "Kildrak", "Morgran", "Orsik", "Oskar", "Rangrim",
			"Rurik", "Taklinn", "Thoradin", "Thorin", "Tordek", "Traubon", "Travok", "Ulfgar", "Veit", "Vondal",
		},
		female_names: []string{
			"Amber", "Artin", "Audhild", "Bardryn", "Dagnal", "Diesa", "Eldeth",
			"Falkrunn", "Finellen", "Gunnloda", "Gurdis", "Helja", "Hlin", "Kathra",
			"Kristryd", "Ilde", "Liftrasa", "Mardred", "Riswynn", "Sannl", "Torbera",
			"Torgga", "Vistra",
		},
		// clan names
		surnames: []string{
			"Balderk", "Battlehammer", "Brawnanvil", "Dankil", "Fireforge", "Frostbeard",
			"Gorunn", "Holderhek","Ironfist", "Loderr", "Lutgehr", "Rumnaheim",
			"Strakeln", "Torunn", "Ungart",
		},
		skin_colors: []string{
			"Light", "Dark", "Tan",
		},
		hair_colors: []string{
			"Red", "Brown", "Black",
		},
		eye_colors: []string{
			"Golden", "Hazel", "Brown",
		},
		height_min: 3.0,
		height_max: 5.0,
		weight_min: 50,
		weight_max: 200,
		age_min: 10,
		age_max: 450,
	},
	"Gnome": {
		male_names: []string{
			"Alston", "Alvyn", "Boddynock", "Brocc", "Burgell", "Dimble", "Eldon",
			"Erky", "Fonkin", "Frug", "Gerbo", "Gimble", "Glim", "Jebeddo", "Kellen",
			"Namfoodle", "Orryn", "Roondar", "Seebo", "Sindri", "Warryn", "Wrenn", "Zook",
		},
		female_names: []string{
			"Bimpnottin", "Breena", "Caramip", "Carlin", "Donella", "Duvamil", "Ella",
			"Ellyjobell", "Ellywick", "Lilli", "Loopmottin", "Lorilla", "Mardnab",
			"Nissa", "Nyx", "Oda", "Orla", "Roywyn", "Shamil", "Tana", "Waywocket", "Zanna",
		},
		surnames: []string{
			"Beren", "Daergel", "Folkor", "Garrick", "Nackle", "Murnig", "Ningel",
			"Raulnor", "Scheppen", "Timbers", "Turen",
		},
		skin_colors: []string{
			"Walnut", "Light", "Dark", "Tan",
		},
		hair_colors: []string{
			"Blonde", "Brown", "Black",
		},
		eye_colors: []string{
			"Turquoise", "Blue", "Green", "Hazel", "Brown",
		},
		height_min: 2.0,
		height_max: 3.5,
		weight_min: 30,
		weight_max: 50,
		age_min: 14,
		age_max: 500,
	},
	"Halfling": {
		male_names: []string{
			"Alton", "Ander", "Cade", "Corrin", "Eldon", "Errich", "Finnan",
			"Garret", "Lindal", "Lyle", "Merric", "Milo", "Osborn", "Perrin",
			"Reed", "Roscoe", "Wellby",
		},
		female_names: []string{
			"Andry", "Bree", "Callie", "Cora", "Euphemia", "Jillian", "Kithri",
			"Lavinia", "Lidda", "Merla", "Nedda", "Paela", "Portia", "Seraphina",
			"Shaena", "Trym", "Vani", "Verna",
		},
		surnames: []string{
			"Brushgather", "Goodbarrel", "Greenbottle", "High-hill", "Hilltopple",
			"Leagallow", "Tealeaf", "Thorngage", "Tosscobble", "Underbough",
		},
		skin_colors: []string{
			"Tan", "Light", "Pale",
		},
		hair_colors: []string{
			"Sandy", "Brown", "Blonde",
		},
		eye_colors: []string{
			"Brown", "Blue", "Green",
		},
		height_min: 2.0,
		height_max: 3.5,
		weight_min: 30,
		weight_max: 50,
		age_min: 10,
		age_max: 90,
	},
	"Tiefling": {
		male_names: []string{
			"Akmenos", "Amnon", "Barakas", "Damakos", "Ekemon", "Iados", "Kairon",
			"Leucis", "Melech", "Mordai", "Morthos", "Pelaios", "Skamos", "Therai",
			"Art", "Carrion", "Chant", "Creed", "Despair", "Excellence", "Fear",
			"Glory", "Hope", "Ideal", "Music", "Nowhere", "Open", "Poetry", "Quest",
			"Random", "Reverence", "Sorrow", "Temerity", "Torment", "Weary",
		},
		female_names: []string{
			"Akta", "Anakis", "Bryseis", "Criella", "Damaia", "Ea", "Kallista",
			"Lerissa", "Makaria", "Nemeia", "Orianna", "Phelaia", "Rieta",
			"Art", "Carrion", "Chant", "Creed", "Despair", "Excellence", "Fear",
			"Glory", "Hope", "Ideal", "Music", "Nowhere", "Open", "Poetry", "Quest",
			"Random", "Reverence", "Sorrow", "Temerity", "Torment", "Weary",
		},
		surnames: []string{
			"",
		},
		skin_colors: []string{
			"Red", "Purple", "Blue",
		},
		hair_colors: []string{
			"Black", "Brown", "Red", "Blue", "Purple",
		},
		eye_colors: []string{
			"Black", "Red", "White", "Silver", "Gold",
		},
		height_min: 4.0,
		height_max: 7.0,
		weight_min: 90,
		weight_max: 240,
		age_min: 15,
		age_max: 110,
	},
}

var race_names = []string{
	"Human", "Elf", "Orc", "Dwarf", "Gnome", "Halfling", "Tiefling",
}

var person_goals = []string{
	"Explore a different land",
	"Provide for my family",
	"Conduct business all across the country",
	"Craft something exquisite",
	"Hike to the top of a great mountain",
	"Discover new magic",
	"Study magic at the college",
	"Purchase a boat to explore the lake",
	"Find a hidden treasure",
	"Travel to another plane",
	"Take back the city from the man",
	"Rebel against the government",
	"Be rewarded with great riches",
	"To rebuild the city",
	"Help the poor",
	"Study lore",
	"Build a temple to my god",
}

var person_clothings_male = []string{
	"Fabric Shirt", "Leather Shirt", "Sleaveless Shirt", "Worn Robe", "Fancy Robe", "Silk Robe",
	"Smock", "Tunic", "Caped Robe", "Chestplate",
}

var person_clothings_female = []string{
	"Dress", "Colorful Dress", "Worn Dress", "Robe", "Magical Robe", "Caped Robe", "Sleaveless Shirt",
	"Smock", "Blouse",
}

var person_distinctions = []string{
	"Grim Smile", "Calm Demeanor", "Firey Eyes", "Lined Forehead", "Smile Marks",
	"Scar across eye", "Missing an eye", "Muscley", "Missing a tooth",
}

var person_feels = []string{
	"Angry", "Stubborn", "Quiet", "Crotchety", "Calm", "Focused", "Intelligent", "Dumb",
	"Asshole", "Fake", "Truthful", "Helpful", "Kind", "Thankful", "Encouraging", "Naive",
	"Purposeful", "Meaningless",
}

type Person struct {
	name        string
	race        string
	gender      string
	skin_color  string
	age         int
	hair_color  string
	eye_color   string
	distinction string  // grim smile, calm demeanor, firey eyes, ...
	feel        string  // quiet, angry, stubborn, ...
	height      float32 // feet
	weight      int     // lbs
	goal        string
	gold        int
	clothing    string
}

func (self *Person) String() string {
	return fmt.Sprintf("%+v ", *self)
}

func person() *Person {
	self := Person{}

	self.race = race_names[rand.Intn(len(race_names))]
	race := race_defs[self.race]

	age := rand.Int() % (race.age_max - race.age_min) + race.age_min
	self.age = age

	gender := rand.Int() % 2

	self.skin_color = race.skin_colors[rand.Intn(len(race.skin_colors))]
	self.hair_color = race.hair_colors[rand.Intn(len(race.hair_colors))]
	self.eye_color = race.eye_colors[rand.Intn(len(race.eye_colors))]
	self.distinction = person_distinctions[rand.Intn(len(person_distinctions))]
	self.feel = person_feels[rand.Intn(len(person_feels))]
	self.height = float32(rand.NormFloat64()) * float32(race.height_max - race.height_min)
	self.height = float32(math.Abs(float64(self.height))) + race.height_min
	self.weight = rand.Int() % (race.weight_max - race.weight_min) + race.weight_min
	self.goal = person_goals[rand.Intn(len(person_goals))]
	self.gold = rand.Intn(500)

	// male
	if gender == 0 {
		self.name = race.male_names[rand.Intn(len(race.male_names))] + " " + race.surnames[rand.Intn(len(race.surnames))]
		self.gender = "Male"
		self.clothing = person_clothings_male[rand.Intn(len(person_clothings_male))]
	// female
	} else {
		self.name = race.female_names[rand.Intn(len(race.female_names))] + " " + race.surnames[rand.Intn(len(race.surnames))]
		self.clothing = person_clothings_female[rand.Intn(len(person_clothings_female))]
		self.gender = "Female"
	}

	return &self
}

var place_kinds = []string{
	"Tavern", "Inn", "General Store", "Fletchery", "Butcher", "Bath House", "Open-air Market",
	"Temple", "Fishing Goods", "Mail Stop", "Owlry", "Falconer", "Herb Shop",
}

var place_decorations = []string{
	"Hanging Fishing Tools", "Paintings", "Rotted Wood", "Cobwebs", "Hanging Ribbons",
	"Live Music", "Plants", "Bookshelves",
}

var place_feels = []string{
	"Warm", "Smelly", "Damp", "Cold", "Homey", "Distrusting", "Quiet", "Awkward", "Busy",
	"Dark", "Well-lit", "Welcoming", "Clean", "Dirty",
}

var size_min = 1
var size_max = 10

type Place struct {
	name       string
	kind       string    // tavern, inn, store, etc...
	decoration string
	feel       string    // warm, cold, smelly, ...
	people     []*Person // the first person is what the place is named after:
	                     // [Bob, Gracie] -> Bob's Tavern, Gracie is customer
	size       int       // number of people determined
}

func (self *Place) String() string {
	builder := "{name:" + self.name + " decoration:" + self.decoration + " feel:" + self.feel + " "

	for _, p := range self.people {
		builder += p.String()
	}
	builder += "}\n"

	return builder
}

func place() *Place {
	self := Place{}

	self.size = rand.Intn(size_max) + size_min
	self.kind = place_kinds[rand.Intn(len(place_kinds))]
	self.decoration = place_decorations[rand.Intn(len(place_decorations))]
	self.feel = place_feels[rand.Intn(len(place_feels))]

	self.people = make([]*Person, self.size)
	var wg sync.WaitGroup

	wg.Add(self.size)
	for i := 0; i < self.size; i++ {
		go func(idx int) {
			self.people[idx] = person()
			wg.Done()
		}(i)
	}

	wg.Wait()
	self.name = self.people[0].name + "'s " + self.kind

	return &self
}


var pop_min = 300
var pop_max = 12000
var pop_names = []string{
	"Yaafmouth", "Yhirora", "Puebury", "Zhuypus", "Freurith", "Xens", "Wrico", "Crale", "Illetin",
	"Ensmore", "Trocester", "Glefrough", "Goiplens", "Kocland", "Pheatol", "Iclance", "Pling", "Vreley",
	"Osaburgh", "Ingpus", "Droungan", "Zhefield", "Glushire", "Yohwell", "Zruuyham", "Strane", "Ezodon",
	"Jico", "Inemond", "Ingdon", "Ikhuiyland", "Chiham", "Juclens", "Xoysea", "Pautgend", "Yhence", "Iheigh",
	"Izhouis", "Eighport", "Olisdale", "Jewood", "Prievale", "Lunby", "Breefield", "Glanard", "Vouis", "Ceka",
	"Ogrens", "Amsas", "Inasburg", "Xutpolis", "Kreuhburgh", "Kifbridge", "Wotin", "Detsas", "Qront", "Haso",
	"Galo", "Inasshire", "Adabury",
}

var pop_neara = []string{
	"Ocean", "Mountain Range", "River", "Valley", "Stream", "Lake", "Quarry", "Cave", "Dungeon", "Ruins",
}

type Popcenter struct {
	pop    int
	name   string
	places []*Place
	neara  string   // ocean, mountain range, valley, river, ...
}

func popcenter() *Popcenter {
	rand.Seed(time.Now().Unix())
	self := Popcenter{}

	self.pop = rand.Intn(pop_max) + pop_min
	self.name = pop_names[rand.Intn(len(pop_names))]
	self.neara = pop_neara[rand.Intn(len(pop_neara))]

	self.places = make([]*Place, int(math.Log(float64(self.pop))) * 6)
	var wg sync.WaitGroup

	wg.Add(len(self.places))
	for i := 0; i < len(self.places); i++ {
		go func(idx int) {
			self.places[idx] = place()
			wg.Done()
		}(i)
	}

	wg.Wait()
	return &self
}

func (self *Popcenter) String() string {
	builder := "{name:" + self.name + " pop:" + fmt.Sprintf("%d", self.pop) + " neara:" + self.neara + " "
	
	for _, p := range self.places {
		builder += p.String()
	}

	builder += "}\n"

	return builder
}
