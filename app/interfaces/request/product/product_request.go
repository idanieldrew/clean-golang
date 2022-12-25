package product

type (
	Product struct {
		Name   string            `bson:"name" json:"name"`
		Price  string            `bson:"price" json:"price"`
		Fields map[string]string `bson:"fields" json:"fields"`
	}
)
