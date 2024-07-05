package entity

import v1 "k8s.io/api/authorization/v1"

type User interface {
	ToUserEntity(token, url string) VersionUserEntity
	Auth(token, url string) (bool, error)
}

type UserRedis struct {
}

func (u *UserRedis) ToUserEntity(token, url string) VersionUserEntity {
	//TODO implement me
	panic("implement me")
}

func (u *UserRedis) Auth(token, url string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

type UserStr string

func (u *UserStr) ToUserEntity(token, url string) VersionUserEntity {
	//TODO implement me
	panic("implement me")
}

func (u *UserStr) Auth(token, url string) (bool, error) {
	var (
		userEntity VersionUserEntity
	)
	userEntity = u.ToUserEntity(token, url)
	review := userEntity.SubjectReview()
	return review.Status.Allowed, nil

}

type VersionUserEntity interface {
	SubjectReview() v1.SelfSubjectAccessReview
}

type CpaasUserEntity struct{}

func (u *CpaasUserEntity) SubjectReview() v1.SelfSubjectAccessReview {
	//TODO implement me
	panic("implement me")
}

func (u *CpaasUserEntity) ToUserEntity(token, url string) VersionUserEntity {
	//TODO implement me
	panic("implement me")
}

func (u *CpaasUserEntity) Auth(token, url string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
