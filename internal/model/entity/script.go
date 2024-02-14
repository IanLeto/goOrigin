package entity

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
)

type Script interface {
	Do() error
}

type BaseScript struct {
	ID         string
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

// BoolQueryScript  注意query 顺序
func BoolQueryScript(ctx context.Context, client *elastic.Client, query ...elastic.Query) (scripts []*BaseScript, err error) {
	searchSvc := client.Search().Index("script")
	for _, e := range query {
		searchSvc = searchSvc.Query(e)
	}
	result, err := searchSvc.Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, hit := range result.Hits.Hits {
		var ephemeralSc BaseScript
		err = json.Unmarshal(hit.Source, &ephemeralSc)
		scripts = append(scripts, &ephemeralSc)
	}
	return scripts, err
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
