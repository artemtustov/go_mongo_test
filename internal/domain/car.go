package domain

type CarInternal struct {
	Model  string `bson:"model"`
	Engine string `bson:"engine"`
	Year   int    `bson:"year"`
}

type Car struct {
	Model  string `json:"model"`
	Engine string `json:"engine"`
	Year   int    `json:"year"`
}
