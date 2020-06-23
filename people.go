package main

import (
	"fmt"
	"math/rand"
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

var race_names = []string {
	"Human", "Elf", "Orc", "Dwarf", "Gnome", "Halfling", "Tiefling",
}
/*var race_names = []string {
    "English Human", "Irish Elf", "Scottish Dwarf",
}*/

var race_defs = map[string]Race {
    "English Human": {
		male_names: []string{
            "Tawnie", "Holden", "Edison", "Winthorpe", "Lyndon", "Sawyer",
            "Melba", "Haywood", "Alfred", "Alden", "Linton", "Huxley", "Winston",
            "Knox", "Sutton", "Edvin", "Harlan", "Oswald", "Milton", "Cyneric",
            "Wyot", "Auden", "Edric", "Hedley", "Tatum", "Fulton", "Egbert",
            "Woodrow", "Osbert", "Egbert", "Aiken", "Graeme", "Landyn", "Hale",
            "Riley", "Edward", "Elton", "Hereward", "Becket", "Edmund", "Wulfnod",
            "Eardwulf", "Dane", "Cynebald", "Gibson", "Aesc", "Swithin", "Harold",
		},
		female_names: []string {
            "Kelsey", "Kaelyn", "Ash", "Tinble", "Kimberley", "Edeva", "Gerarda",
            "Beverly", "Odella", "Velma", "Brynlee", "Brooke", "Rue", "Lena", "Lana",
            "Willow", "Bathilde", "Raylee", "Leanne", "Levina", "Godiva", "Deana",
            "Eden", "Ainslee", "Dell", "Audrey", "Ealdgyo", "Willa", "Eoforhild",
            "Trilby", "Zanna", "Marjorie", "Edgarda", "Holly", "Marigold", "Iria",
            "Lindley", "Louella", "Haralda", "Piper", "Flyta", "Imogene", "Dale",
            "Earna", "Delwyn", "Zelene", "Paige", "Kelcie", "Daisy", "Chancey",
		},
		surnames: []string {
            "Wilmere", "Lintone", "Hardene", "Thorntone", "Glasse", "Waller", "Hayes",
            "Bardsleye", "Easome", "Underwoode", "Burkee", "Blakee", "Kimballe",
            "Froste", "Hawkinge", "Tashe", "Hathawaye", "Buckley", "Hambletone",
            "Weaver", "Woodcocke", "Presleye", "Boatwrighte", "Alvine", "Cookee",
            "Weekese", "Harrise", "Warwicke", "Aytone", "Whitee", "Everlye", "Hanleye",
            "Harrelsone", "Uggerii", "Beasleye", "Brewstere", "Greye", "Alfredssons",
            "Fenne", "Lyndone", "Truemane", "Blackbournee", "Alveye", "Caldwelle",
		},
		skin_colors: []string {
			"Light", "Pale", "Tan",
		},
		hair_colors: []string {
			"Black", "Brown", "Blonde", "Red", "White", "Gray", "Bald",
		},
		eye_colors: []string {
			"Brown", "Blue", "Green", "Hazel", "Dark", "Gray",
		},
		height_min: 5.0,
		height_max: 6.5,
		weight_min: 120,
		weight_max: 200,
		age_min: 20,
		age_max: 70,
    },
    "Irish Elf": {
		male_names: []string {
            "Abban", "Ailill", "Ailbe", "Ainmire", "Amhlaoibh", "Anlon", "Aodh",
            "Aodghan", "Aongus", "Banan", "Bradan", "Brandan", "Brin", "Cairbre",
            "Kenneth", "Kevin", "Carroll", "Kane", "Killian", "Kieran", "Colm",
            "Colman", "Conan", "Conn", "Cormac", "Christopher", "Dara", "Dermot",
            "Donald", "Edmund", "Heber", "Enda", "Owen", "Fergal", "Fergus", "Barry",
            "Finbar", "Finn", "Garvan", "Garret", "Christian", "Jarlath", "Liam",
            "Lawrence", "Loo", "Mannix", "Manus", "Mel", "Oran", "Patrick", "Peter",
            "Ryan", "Ronan", "Ross", "Seamus", "James", "Sean", "Tiernan", "Thomas",
            "Ultan",
		},
		female_names: []string {
            "Afric", "April", "Aideen", "Aileen", "Ashling", "Alannah", "Alma", "Eefa",
            "Betha", "Bridget", "Brona", "Caitlin", "Cait", "Keelin", "Keva", "Kathleen",
            "Keera", "Kiara", "Colleen", "Datan", "Darerca", "Darina", "Dervla", "Duveesa",
            "Ealga", "Eileen", "Elizabeth", "Emer", "Enya", "Helen", "Erin", "Etain",
            "Fiona", "Grania", "Agnes", "Iona", "Isleen", "Macha", "Maeve", "Mary", "Maura",
            "Moira", "Melissa", "Mona", "Nessa", "Neve", "Noelle", "Nora", "Una", "Orla",
            "Rianne", "Riona", "Rosaleen", "Sive", "Saoirse", "Shannon", "Shona", "Sheena",
            "Shivaun", "Joan", "Sarah",
		},
		surnames: []string {
            "Murphy", "Kelly", "O'Sullivan", "Walsh", "Smith", "O'Brien", "Byrne", "Ryan",
            "O'Connor", "O'Niell", "O'Reilly", "Doyle", "McCarthy", "Gallagher", "O'Doherty",
            "Kennedy", "Lynch", "Murray", "Quinn", "Moore", "McLoughlin", "O'Carroll",
            "Connolly", "Daly", "O'Connell", "Wilson", "Dunne", "Brennan", "Burke", "Collins",
            "Campbell", "Clarke", "Johnston", "Hughes", "O'Farrell", "Fitzgerald", "Brown",
            "Martin", "Maguire", "Nolan", "Flynn", "Thompson", "O'Callaghan", "O'Donnell",
            "Duffy", "O'Mahony", "Boyle", "Healy", "O'Shea", "White", "Sweeney", "Hayes",
            "Kavanagh", "Power", "McGrath", "Moran", "Brady", "Stewart", "Casey", "Foley",
            "Fitzpatrick", "O'Leary", "McDonnell", "MacMahon", "Donnelly", "Regan", "Donovan",
            "Burns", "Flanagan", "Mullan", "Barry", "Kane", "Robinson", "Cunningham", "Griffin",
            "Kenny", "Sheehan", "Ward", "Whelan", "LYons", "Reid", "Graham", "Higgins",
            "Cullen", "O'Keeffe", "Magee", "MacNamara", "MacDonald", "MacDermott", "Molony",
            "O'Rourke", "Buckley", "O'Dwyer",
		},
		skin_colors: []string {
			"Light", "Pale", "Tan",
		},
		hair_colors: []string {
			"Black", "Silver", "Brown", "Gold", "Blonde", "White", "Gray", "Bald",
		},
		eye_colors: []string {
			"Brown", "Light", "Green", "Blue", "Yellow", "Hazel", "Gray", "Gold", "Silver",
		},
		height_min: 5.3,
		height_max: 6.2,
		weight_min: 110,
		weight_max: 180,
		age_min: 20,
		age_max: 700,
    },
    "Scottish Dwarf": {
		male_names: []string{
            "Adie", "Addie", "Adam", "Albert", "Alan", "Alec", "Alex", "Alpin",
            "Andrew", "Alastair", "Aulay", "Angie", "Hugh", "Aidan", "Angus",
            "Harold", "Arthur", "Aaron", "Walter", "Benedict", "Brian", "Benjamin",
            "Colin", "Kevin", "Charles", "Kieran", "Clement", "Kenneth", "Constantine",
            "Coll", "Columba", "Connal", "Cormac", "Christopher", "Christian", "David",
            "George", "Dermid", "Duff", "Danny", "Douglas", "Hector", "Archie", "Ellar",
            "Edmund", "Ewan", "John", "Jonathan", "Joseph", "Fillan", "Farquhar",
            "Philip", "Fingal", "Franklin", "Gilbert", "Gordon", "Godfrey", "Gregory",
            "Grant", "Archibald", "Jack", "Jonah", "Lawrence", "Lewis", "Luke", "Myles",
            "Malcolm", "Miles", "Mark", "Martin", "Michael", "Morgan", "Magnus",
            "Oliver", "Nicholas", "Neil", "Francis", "Paul", "Peter", "Rob",
            "Derrick", "Roderick", "Rory", "Zachary", "Samuel", "Seth", "Simon",
            "Thomas", "Norman", "Willie", "Eugene",
		},
		female_names: []string{
            "Alice", "Amelia", "Emily", "Angelica", "Annabella", "Ann", "Anne", "Annie",
            "Ava", "Barbara", "Bessy", "Becky", "Beatrice", "Victoria", "Bridget",
            "Christine", "Christina", "Catherine", "Catrina", "Kenna", "Kate", "Katie",
            "Dorcas", "Dorothy", "Dolly", "Doreen", "Elizabeth", "Ailie", "Ellen", "Helen",
            "Eve", "Eva", "Evelyn", "Fiona", "Flora", "Florence", "Frances", "Grace",
            "Isabel", "Isabella", "Lexie", "Lily", "Lucretia", "Lisa", "Maggie",
            "Marjory", "Martha", "Mildred", "Mabel", "Mary", "Marion", "Olivia", "Rachel",
            "Ruth", "Jessie", "Joan", "Shona", "Janet", "Susan", "Susanna", "Claire",
            "Clara", "Sarah", "Sorche", "Cecilia", "Julia", "Jane", "Jean", "Jenny", "Una",
		},
		surnames: []string{
            "Abercrombie", "Smith", "Brown", "Wilson", "Robertson", "Thomson", "Campbell",
            "Stewart", "Anderson", "Scott", "Murray", "MacDonald", "Reid", "Taylor",
            "Clark", "Ross", "Young", "Mitchell", "Watson", "Paterson", "Morrison",
            "Murphy", "Kelly", "Walsh", "Byrne", "Ryan", "Doyle", "McCarthy", "Gallagher",
            "Kennedy", "Lynch", "Murray", "Quinn", "Moore", "McLoughlin", "Connolly",
            "Daly", "Wilson", "Dunne", "Brennan", "Burke", "Collins", "Campbell", "Clarke",
            "Johnston", "Hughes", "Fitzgerald", "Brown", "Martin", "Maguire", "Nolan",
            "Flynn", "Thompson", "Duffy", "Boyle", "Healy", "White", "Sweeney", "Hayes",
            "Kavanagh", "Power", "McGrath", "Moran", "Brady", "Stewart", "Casey", "Foley",
            "McDonnell", "MacMahon", "Donnelly", "Regan", "Donovan", "Burns", "Flanagan",
            "Mullan", "Barry", "Kane", "Robinson", "Cunningham", "Griffin", "Kenny",
            "Sheehan", "Ward", "Whelan", "LYons", "Reid", "Graham", "Higgins", "Cullen",
            "Magee", "MacNamara", "MacDonald", "MacDermott", "Molony", "Buckley",
		},
		skin_colors: []string{
			"Light", "Pale", "Tan",
		},
		hair_colors: []string{
			"Red", "Brown", "Black", "Bald",
		},
		eye_colors: []string{
			"Golden", "Hazel", "Brown", "Green", "Gray",
		},
		height_min: 3.0,
		height_max: 5.0,
		weight_min: 50,
		weight_max: 200,
		age_min: 20,
		age_max: 450,
    },
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
		female_names: []string {
			"Atala", "Ceidil", "Hama", "Jasmal", "Meilil", "Seipora", "Yasheira", "Zasheida",
			"Arveene", "Esvele", "Jhessail", "Kerri", "Lureene", "Miri", "Rowan", "Shandri", "Tessele",
			"Alethra", "Kara", "Katernin", "Mara", "Natali", "Olma", "Tana", "Zora",
			"Amafrey", "Betha", "Cefrey", "Kethra", "Mara", "Olga", "Silifrey", "Westra",
			"Arizima", "Chathi", "Nephis", "Nulara", "Murithi", "Sefris", "Thola", "Umara", "Zolis",
			"Fyevarra", "Hulmarra", "Immith", "Imzel", "Navarra", "Shevarra", "Tammith", "Yuldra",
			"Bai", "Chao", "Jia", "Lei", "Mei", "Qiao", "Shui", "Tai",
			"Balama", "Dona", "Faila", "Jalana", "Luisa", "Marta", "Quara", "Selise", "Vonda",
		},
		surnames: []string {
			"Basha", "Dumein", "Jassan", "Khalid", "Mostana", "Pashar", "Rein",
			"Amblecrown", "Buckman", "Dundragon", "Evenwood", "Greycastle", "Tallstag",
			"Bersk", "Chernin", "Dotsk", "Kulenov", "Marsk", "Nemetsk", "Shemov", "Starag",
			"Brightwood", "Helder", "Hornraven", "Lackman", "Stormwind", "Windrivver",
			"Ankhalab", "Anskuld", "Fezim", "Hahpet", "Nathandem", "Sepret", "Uuthrakt",
			"Chergoba", "Dyernina", "Iltazyara", "Murnyethara", "Stayanoga", "Ulmokina",
			"Chien", "Huang", "Kao", "Kung", "Lao", "Ling", "Mei", "Pin", "Shin", "Sum", "Tan", "Wan",
			"Agosto", "Astorio", "Calabra", "Domine", "Falone", "Marivaldi", "Pisacar", "Ramondo",
		},
		skin_colors: []string {
			"Light", "Dark", "Tan",
		},
		hair_colors: []string {
			"Black", "Brown", "Blonde", "Red", "White", "Gray", "Bald",
		},
		eye_colors: []string {
			"Brown", "Blue", "Green", "Hazel", "Dark", "Gray",
		},
		height_min: 5.0,
		height_max: 6.5,
		weight_min: 120,
		weight_max: 200,
		age_min: 20,
		age_max: 70,
	},
	"Elf": {
		male_names: []string {
			"Adran", "Aelar", "Aramil", "Arannis", "Aust", "Beiro", "Berrian", "Carric", "Enialis",
			"Erdan", "Erevan", "Galinndan", "Hadarai", "Heian", "Himo", "Immeral", "Ivellios", "Laucian",
			"Mindartis", "Paelias", "Peren", "Quarion", "Riardon", "Rolen", "Soveliss", "Thamior",
			"Tharivol", "Theren", "Varis",
		},
		female_names: []string {
			"Adrie", "Althaea", "Anastrianna", "Andraste", "Antinua", "Bethrynna", "Birel", "Caelynn",
			"Drusilia", "Enna", "Felosial", "Ielenia", "Jelenneth", "Keyleth", "Leshanna", "Lia", "Meriele",
			"Mialee", "Naivara", "Quelenna", "Quillathe", "Sariel", "Shanairra", "Shava", "Silaqui",
			"Theirastra", "Thia", "Vadania", "Valanthe", "Xanaphia",
		},
		surnames: []string {
			"Amakiir", "Amastacia", "Galanodel", "Holimion",
			"Ilphelkiir", "Liadon", "Meliamne", "NaNightbreeze",
			"Siannodel", "Xiloscient",
			"Gemflower", "Starflower", "Moonwhisper", "Diamonddew", "Gemblossom", "Silverfrond", "Oakenheel",
			"Nightbreeze", "Moonbrook", "Goldpetal",
		},
		skin_colors: []string {
			"Light", "Dark", "Hazel", "Tan",
		},
		hair_colors: []string {
			"Black", "Silver", "Brown", "Gold", "Blonde", "White", "Gray", "Bald",
		},
		eye_colors: []string {
			"Brown", "Light", "Green", "Blue", "Yellow", "Hazel", "Gray", "Gold", "Silver",
		},
		height_min: 5.3,
		height_max: 6.2,
		weight_min: 110,
		weight_max: 180,
		age_min: 20,
		age_max: 700,
	},
	"Orc": {
		male_names: []string {
			"Dench", "Feng", "Gell", "Henk", "Holg", "Imsh", "Keth", "Krusk",
			"Mhurren", "Ront", "Shump", "Thokk",
		},
		female_names: []string {
			"Baggi", "Emen", "Engong", "Kansif", "Myev", "Neega", "Ovak", "Ownka",
			"Shautha", "Sutha", "Vola", "Volen", "Yevelda",
		},
		surnames: []string {
			"",
		},
		skin_colors: []string {
			"Green", "Gray", "Dark", "Purple", "Light Blue", "Pale",
		},
		hair_colors: []string {
			"Black", "Bald", "Brown", "Maroon", "Bald",
		},
		eye_colors: []string {
			"Black", "Dark", "Brown", "Green", "Yellow", "Orange", "Purple",
		},
		height_min: 5.8,
		height_max: 7.5,
		weight_min: 170,
		weight_max: 240,
		age_min: 20,
		age_max: 70,
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
			"Red", "Brown", "Black", "Bald",
		},
		eye_colors: []string{
			"Golden", "Hazel", "Brown", "Green", "Gray",
		},
		height_min: 3.0,
		height_max: 5.0,
		weight_min: 50,
		weight_max: 200,
		age_min: 20,
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
			"Blonde", "Brown", "Black", "Bald",
		},
		eye_colors: []string{
			"Turquoise", "Blue", "Green", "Hazel", "Brown", "Gray", "Gold", "Silver",
		},
		height_min: 2.0,
		height_max: 3.5,
		weight_min: 30,
		weight_max: 50,
		age_min: 20,
		age_max: 400,
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
			"Sandy", "Brown", "Blonde", "Black", "Bald",
		},
		eye_colors: []string{
			"Brown", "Blue", "Green", "Hazel", "Gray",
		},
		height_min: 2.0,
		height_max: 3.5,
		weight_min: 30,
		weight_max: 50,
		age_min: 20,
		age_max: 75,
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
		height_min: 5.0,
		height_max: 6.5,
		weight_min: 90,
		weight_max: 200,
		age_min: 20,
		age_max: 75,
	},
}

var person_goals = []string{
	"Explore a different land",
	"Provide for family",
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
	"Build a temple to their god",
	"Survive",
	"Seek a position of power",
	"Explore ruins",
	"Become a cleric",
	"Become a priest",
	"Become an explorer",
	"Become a historian",
	"Become a luthier",
	"Become a musician",
	"Become a warrior",
	"Become a miner",
	"To fall in love",
	"Find their mother",
	"Find their father",
	"Find their son",
	"Find their daughter",
	"Take their revenge",
	"To help their guild",
	"Rise the ranks of their guild",
	"Rise the ranks of the military",
	"Discover a new, rare animal",
	"Find and protect a hidden natural wonderland",
	"Write a comprehensive bestiary of the continent",
	"Learn to fight",
	"Open a great zoo",
	"To create a masterpiece painting",
    "To build houses",
    "To trade between distant cities",
}

var person_clothing_adjectives = []string {
	"Red", "Orange", "Yellow", "Green", "Blue", "Purple", "Indigo", "Violet",
	"Emerald", "Jade", "Beige", "Black", "Gray", "Maroon", "Colorful",
	"Fancy", "Worn", "Sleaveless", "Turquoise", "Cyan", "Sapphire", "Caped",
	"Silk", "Intricate", "Hooded", "Mantled",
}

var person_clothings_male = []string {
	"Shirt", "Robe", "Smock", "Tunic", "Chestplate", "Trenchcoat", "Religious Armor",
	"Artisan's Clothing", "Cleric's vestments", "Cold weather outfit", "Religious Robes",
	"Cloak", "Cape", "Coat", "Fur-lined Cloak", "Woolen Cloak", "Jacket", "Woolen Tunic",
    "Collared Shirt", "Shirt with waistcoat", "Shirt with cummerbund",
}

var person_clothings_female = []string {
	"Dress", "Laced Dress", "Blouse", "Bodice", "Dress with flared sleeves", "Dress with petticoat",
    "Polonaise", "Witzchoura", "Overgarment", "Gown", "Long skirt", "Surcoat", "Bliaut", "Cotehardie",
}

var person_distinctions = []string {
	"Grim Smile", "Calm Demeanor", "Firey Eyes", "Lined Forehead", "Smile Marks",
	"Scar across eye", "Missing an eye", "Muscley", "Missing a tooth", "Long hair",
	"Lazy Eye", "Scar across nose", "Missing piece of ear", "Walk with limp",
	"Always wears armor", "Perpetually drunk", "Walks briskly", "Mumbles to self",
	"Paces back and forth", "Walks gracefully", "Great Smile", "Sparkling eyes",
	"Skips when walking", "Talks to self", "Speaks in third person", "Wears Glasses",
	"Wears a beret", "Wears a scarf", "Wears a top hat", "Carries a purse",
	"Carries a backpack", "Travels with a bird", "Travels with a dog", "Travels with a cat",
	"Travels with an owl", "Travels with a bat", "Always carries sword", "Carries longsword",
	"Carries greatsword", "Carries shortsword", "Carries longbow", "Carries shortbow",
	"Carries crossbow", "Carries hammer", "Carries axe", "Carries a quiver", "Walks with a cane",
	"Sharp chin", "Oval face", "Uses wheelchair", "Carrier hiking gear", "Carries fishing gear",
	"Has arm tattoo", "Has neck tattoo", "Has chest tattoo", "Jade Necklace", "Long Necklace",
	"Necklace religious symbol", "Necklace with latch", "Earrings", "Hooked Earring",
}

var person_feels = []string {
	"Angry", "Stubborn", "Quiet", "Crotchety", "Calm", "Focused", "Intelligent", "Dumb",
	"Asshole", "Fake", "Truthful", "Helpful", "Kind", "Thankful", "Encouraging", "Naive",
	"Purposeful", "Meaningless", "Haggard", "Non-conformist", "Conformist", "Boring",
	"Entrepenurial", "Handy", "Happy", "Joyful", "Healthy", "Fat", "Skinny",
	"Happy-go-lucky", "Hard-ass", "Hard-working", "Sickly", "Harsh", "Hesitating",
	"Disgruntled", "Jaded", "Hasty", "Slow", "Fast", "Clumsy", "Graceful", "Hip",
	"Horrible", "Annoying", "Grumpy", "Depressed", "Immaculate", "Sarcastic", "Jovial",
	"Lax", "Merry", "Middle-aged", "Morbid", "Ne'er-do-well", "Obese", "Obtuse", "Acute",
	"Offensive", "Bold", "Promiscuous", "Pervasive", "Pious", "Sharp", "Resourceful",
	"Rude", "Witty", "Shy", "Happy", "Cheerful", "Bright", "Wonderful", "Kind", "Thrifty",
	"Trustworthy", "Brave", "Curteous", "Clean", "Useful", "Present",
}

var person_gold_max int = 200

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
	return fmt.Sprintf("%s, %s %s %d, %.2f ft, %d lbs, %d gold | %s skin, %s hair, %s eyes, Wearing %s, %s | %s, wants to %s",
		self.name, self.gender, self.race, self.age, self.height, self.weight, self.gold,
		self.skin_color, self.hair_color, self.eye_color, self.clothing, self.distinction, self.feel, self.goal)
}

func Person_new() *Person {
	return Person_new_race(race_names[rand.Intn(len(race_names))])
}

func Person_new_race(race_name string) *Person {
	self := Person{}

	self.race = race_name
	race := race_defs[race_name]

	age := rand.Int() % (race.age_max - race.age_min) + race.age_min
	self.age = age

	self.skin_color  = pick_one(race.skin_colors)
	self.hair_color  = pick_one(race.hair_colors)
	self.eye_color   = pick_one(race.eye_colors)
	self.distinction = pick_one(person_distinctions)
	self.feel        = pick_one(person_feels)

	self.height = rand.Float32() * float32(race.height_max - race.height_min) + race.height_min
	height_norm := (self.height - race.height_min) / race.height_max
	self.weight = rand.Int() % (race.weight_max - race.weight_min) + race.weight_min
	self.weight += int(float32(self.weight) * height_norm)

	self.goal = pick_one(person_goals)
	self.gold = rand.Intn(person_gold_max)

	// male
	if rand.Int() % 2 == 0 {
		self.name = pick_one(race.male_names)
		surname := pick_one(race.surnames)
		if surname != "" {
			self.name += fmt.Sprintf(" %s", surname)
		}
		self.clothing = pick_one(person_clothing_adjectives) + " " + pick_one(person_clothings_male)
		self.gender = "Male"
	// female
	} else {
		self.name = pick_one(race.female_names)
		surname := pick_one(race.surnames)
		if surname != "" {
			self.name += fmt.Sprintf(" %s", surname)
		}

		if rand.Int() % 2 == 0 {
			self.clothing = pick_one(person_clothing_adjectives) + " " + pick_one(person_clothings_male)
		} else {
			self.clothing = pick_one(person_clothing_adjectives) + " " + pick_one(person_clothings_female)
		}
		self.gender = "Female"
	}

	return &self
}
