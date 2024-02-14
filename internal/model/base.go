package model

type Entity interface {
	ToDao() Dao
}

type Dao interface {
	ToEntity() Entity
}
