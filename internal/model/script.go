package model

type Script interface {
	Do() error
}

type BaseScript struct {
	ID         int
	Name       string
	Comment    string
	Type       string
	Content    string
	File       string
	Uploader   string
	CreateTime int
	UpdateTime int
	System     string
	IsFile     bool
	Timeout    int
	Tags       []string
	UsedTime   int
}

type PythonScript struct {
	*BaseScript
}

func (p *PythonScript) Do() error {
	//TODO implement me
	panic("implement me")
}

type SellScript struct {
	*BaseScript
}

func (s *SellScript) Do() error {
	panic(1)
}
