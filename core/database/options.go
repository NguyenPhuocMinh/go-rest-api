package database

type FindOptionModel struct {
	Skip  int
	Limit int
	Sort  []string
}

type FindOneOptionModel struct {
	Sort       interface{}
	Projection interface{}
}
