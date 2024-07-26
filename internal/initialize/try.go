package initialize

import (
	"fmt"
	"gproject/internal/initialize/global"
	"gproject/internal/initialize/po"
	"github.com/google/uuid"
	"math/rand/v2"
)

func TryDataSample() {
	tryRedis()
	tryMySQL()

	// kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	// defer kafkaProducer.Close()
}

func tryRedis() {
	global.Rdb.Set(ctx, "score", 100, 0)
	global.Rdb.IncrBy(ctx, "score", 10)
	value, _ := global.Rdb.Get(ctx, "score").Result()
	global.Logger.Info("test redis with score: " + value)
}

func tryMySQL() {
	// *************************************** remove table
	// global.Mdb.Migrator().DropTable(&po.Role{})
	// global.Mdb.Migrator().DropTable(&po.User{})


	// *************************************** add table
	// global.Mdb.Migrator().CreateTable(&po.Role{})
	// global.Mdb.Migrator().CreateTable(&po.User{})

	global.Mdb.AutoMigrate(&po.User{})
	global.Mdb.AutoMigrate(&po.Role{})

	// *************************************** insert 1
	// role := po.Role{
	// 	RoleName: "Admin",
	// 	RoleNote: "Administrator role",
	// }
	// global.Mdb.Create(&role)

	// *************************************** insert n
	inserts := generateRoles(1000)
	global.Mdb.Create(&inserts)

	// *************************************** get all
	// var roles []po.Role
	// global.Mdb.Find(&roles)
	// for _, role := range roles {
	// 	fmt.Printf("ID: %d \t Name: %s \t Note: %s\n", role.ID, role.RoleName, role.RoleNote)
	// }

	// *************************************** count
	// var count int64
	// global.Mdb.Model(&po.Role{}).Count(&count)
	// fmt.Printf("Total rows in the table: %d\n", count)

	// *************************************** Retrieve roles with note starting with "A"
	// var roles []po.Role
	// global.Mdb.Where("role_note LIKE ?", "C%").Find(&roles)
	// for _, role := range roles {
	// 	fmt.Printf("ID: %d \t Name: %s \t Note: %s\n", role.ID, role.RoleName, role.RoleNote)
	// }


	user := po.User{
		UUID:     uuid.New(),
		UserName: "john_doe",
		IsActive: true,
		Roles: []po.Role{
			{ID: rand.Int64N(1000)},
			{ID: rand.Int64N(1000)},
			{ID: rand.Int64N(1000)},
			{ID: rand.Int64N(1000)},
		},
	}
	global.Mdb.Create(&user)

	var users []po.User
	global.Mdb.Preload("Roles").Find(&users)


	for _, user := range users {
		fmt.Printf("User: %s\n", user.UserName)
		fmt.Println("Roles:")
		for _, role := range user.Roles {
			fmt.Printf("- %s\n", role.RoleName)
		}
		fmt.Println()
	}
}

func generateRoles(count int) []po.Role {
	var roles []po.Role

	for i := 0; i < count; i++ {
		role := po.Role{
			RoleName: fmt.Sprintf("%s_%d",  generateRandomNote(), i+1),
			RoleNote: generateRandomNote(),
		}
		roles = append(roles, role)
	}

	return roles
}

func generateRandomNote() string {
	notes := []string{
		"Apple", "Banana", "Carrot", "Donut", "Eggplant", "Fig", "Grapefruit", "Honeydew", "Ice cream",
		"Jalapeno", "Kiwi", "Lemon", "Mango", "Nectarine", "Orange", "Pineapple", "Quinoa", "Raspberry",
		"Strawberry", "Tomato", "Udon", "Vermicelli", "Watermelon", "Xigua", "Yam", "Zucchini", "Almond",
		"Broccoli", "Cauliflower", "Doughnut", "Eclair", "Fettuccine", "Ginger", "Hazelnut", "Iceberg lettuce",
		"Jackfruit", "Kale", "Lettuce", "Mushroom", "Nutmeg", "Oatmeal", "Peach", "Quiche", "Raisin", "Spinach",
		"Tofu", "Ugli fruit", "Vanilla", "Walnut", "Xanthan gum", "Yogurt", "Ziti", "Avocado", "Blueberry", "Cabbage",
		"Date", "Egg", "Fries", "Gouda cheese", "Hamburger", "Italian bread", "Jicama", "Ketchup", "Lobster", "Milk",
		"Nacho", "Olive", "Pasta", "Quail", "Ramen", "Salmon", "Taco", "Ube", "Vegetable soup", "Watercress", "Xacuti",
		"Yakitori", "Zabaglione", "Artichoke", "Blackberry", "Caramel", "Dough", "Eggplant parmesan", "French fries",
		"Gingerbread", "Honey", "Ice pop", "Jelly", "Kale chips", "Lemonade", "Macadamia nut", "Noodle", "Oyster",
		"Pancake", "Quinoa salad", "Ratatouille", "Sausage", "Tiramisu", "Udon noodles", "Vegetable stir-fry",
		"Waffle", "Xmas cake", "Yakisoba", "Zucchini bread", "Archery", "Basketball", "Cycling", "Dancing",
		"Elliptical training", "Football", "Golfing", "Hiking", "Ice skating", "Jumping rope", "Kickboxing",
		"Lacrosse", "Mountain biking", "Nordic walking", "Orienteering", "Pilates", "Qigong", "Rock climbing",
		"Skiing", "Tennis", "Ultimate Frisbee", "Volleyball", "Walking", "X-country skiing", "Yoga", "Zumba",
		"Aerobics", "Badminton", "Canoeing", "Dodgeball", "Equestrianism", "Fencing", "Gymnastics", "Hockey",
		"Inline skating", "Jogging", "Kayaking", "Lawn bowling", "Martial arts", "Netball", "Obstacle course racing",
		"Paddleboarding", "Quidditch", "Rowing", "Swimming", "Table tennis", "Underwater hockey", "Vigorous cleaning",
		"Weightlifting", "X-fit training", "Yard work", "Zorbing", "Aqua aerobics", "Ballet", "Climbing stairs", "Darts",
		"Exercise classes", "Frisbee golf", "Gardening", "High-intensity interval training (HIIT)", "Ice hockey",
		"Juggling", "Kiteboarding", "Laser tag", "Meditation", "Ninja warrior training", "Outdoor circuit training",
		"Piloxing", "Qi gong", "Rugby", "Skateboarding", "Tai chi", "Unicycling", "Vigorous dancing", "Water polo",
		"Xtreme sports", "Yard games", "Ziplining", "Aerial yoga", "BollyX", "Circus arts", "Disc golf",
		"Endurance running", "Frisbee", "Gym workout", "Horseback riding", "Ice climbing", "Jet skiing",
		"Kangoo jumps", "LARPing", "Mauy Thai", "Nia", "Olympic lifting", "Paddle tennis", "Quidditch",
		"Racquetball", "Synchronized swimming", "Trail running", "Unicycling", "Vigorous swimming", "Water skiing",
		"Xtend barre", "Yin yoga", "Zigzag running", "Air conditioner", "Bluetooth speaker", "Camera", "Drone",
		"Electric toothbrush", "Fan", "Game console", "Hair dryer", "Iron", "Juicer", "Keyboard", "Laptop",
		"Microwave", "Nintendo Switch", "Oven", "Printer", "Quantum computer", "Refrigerator", "Smartphone",
		"Tablet", "Ultrasonic cleaner", "Vacuum cleaner", "Washing machine", "Xbox", "Yoga mat", "Zigbee hub",
		"Alarm clock", "Bluetooth headphones", "Coffee maker", "Digital scale", "E-book reader", "Freezer",
		"Gaming mouse", "Handheld vacuum", "iPad", "JBL speaker", "Kindle", "LED TV", "Monitor", "Nest thermostat",
		"Oculus Quest", "Portable charger", "Quadcopter", "Rice cooker", "Smartwatch", "Toaster", "USB flash drive",
		"Video doorbell", "Wireless mouse", "Xiaomi Mi Box", "Yamaha keyboard", "Z-Wave controller", "Air purifier",
		"Bluetooth earbuds", "Cordless phone", "Drone camera", "Electric kettle", "Fire TV Stick", "Garage door opener",
		"Hard drive", "Infrared thermometer", "Juice extractor", "Kitchen scale", "Laser printer", "Microwave oven",
		"Noise-canceling headphones", "Outdoor security camera", "Power bank", "Quadcopter drone", "Robot vacuum cleaner",
		"Smart lock", "TV remote control", "Ultrabook", "VR headset", "Water dispenser", "Xerox machine", "Yogurt maker",
		"Z-Wave dimmer", "Action camera", "Bluetooth keyboard", "CCTV camera", "Digital camera", "Electric shaver",
		"Fitness tracker", "Gaming console", "Hair straightener", "Induction cooktop", "Juice blender", "Keurig coffee maker",
		"Laptop charger", "Mobile hotspot", "Nest security camera", "OLED TV", "Power strip", "Quadcopter with camera",
		"Raspberry Pi", "Smart thermostat", "USB hub", "Video projector", "Wireless earphones", "Xbox controller",
		"Yamaha soundbar", "Zigbee light bulb",
	}

	randomIndex := rand.IntN(len(notes))
	return notes[randomIndex]
}
