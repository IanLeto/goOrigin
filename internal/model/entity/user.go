package entity

type User interface {
	ToUserEntity(input interface{}) User
	Auth() error
	Allow() bool
}

type UserRedis struct {
}

func (u *UserRedis) ToUserEntity(input interface{}) User {
	//TODO implement me
	panic("implement me")
}

func (u *UserRedis) Auth() error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRedis) Allow() bool {
	//TODO implement me
	panic("implement me")
}

type UserStr string

func (u *UserStr) ToUserEntity(input interface{}) User {
	//TODO implement me
	panic("implement me")
}

func (u *UserStr) Auth() error {
	//TODO implement me
	panic("implement me")
}

func (u *UserStr) Allow() bool {
	//TODO implement me
	panic("implement me")
}
