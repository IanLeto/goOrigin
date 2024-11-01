package cron

import (
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"time"
)

type SyncJob struct {
	Address  string
	User     string
	Password string
	Alias    string // 别名
	Cluster  string
	AZ       string
	Topic    string
	Interval time.Duration
	Client   *elasticsearch7.Client
}

func RegSyncJobCron() error {
	//task := &SyncJob{}
	//var (
	//	err error
	//)
	//client, err := elasticsearch7.NewClient(elasticsearch7.Config{
	//	Addresses: []string{
	//		config.ConfV2.Es.Address,
	//	},
	//})

	return nil

}
