package entity

type User interface {
	ToUserEntity(token, url string) User
	Auth(token, url string) (bool, error)
	Allow() bool
}

type UserRedis struct {
}

func (u *UserRedis) ToUserEntity(token, url string) User {
	//TODO implement me
	panic("implement me")
}

func (u *UserRedis) Auth(token, url string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRedis) Allow() bool {
	//TODO implement me
	panic("implement me")
}

type UserStr string

func (u *UserStr) ToUserEntity(token, url string) User {
	//TODO implement me
	panic("implement me")
}

func (u *UserStr) Auth(token, url string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserStr) Allow() bool {
	//TODO implement me
	panic("implement me")
}

type UserEntity struct {
}

func (u *UserEntity) ToUserEntity(token, url string) User {
	//TODO implement me
	panic("implement me")
}

func (u *UserEntity) Auth(token, url string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserEntity) Allow() bool {
	//TODO implement me
	panic("implement me")
}

func NewUserEntity() {

}
