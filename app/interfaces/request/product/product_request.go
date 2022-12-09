package product

type (
	Refrigerator struct {
		Name      string  `bson:"name"`
		Price     float32 `bson:"price"`
		Color     string  `bson:"color"`
		CountDoor int     `bson:"count_door"`
		Types     string  `bson:"types"`
	}

	VacuumCleaner struct {
		Name    string  `bson:"name"`
		Price   float32 `bson:"price"`
		Color   string  `bson:"color"`
		Suction int     `bson:"Suction"`
		Types   string  `bson:"types"`
	}
)
