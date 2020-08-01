package domain

type Course struct {
	Id        string `bson:"id"`
	Nome      string `bson:"nome"`
	Descricao string `bson:"descricao"`
	Preco     int    `bson:"preco"`
}
