package product

type (
	Product struct {
		Name   string            `bson:"name" json:"name"`
		Slug   string            `bson:"slug" json:"slug"`
		Price  string            `bson:"price" json:"price"`
		Fields map[string]string `bson:"fields" json:"fields"`
	}
)
