package httpclient

import (
	"fmt"
	"github.com/dghubble/sling"
	"goOrigin/config"
	"sync"
	"time"
)

type APIClient interface {
	Mkc() (CCResponseInfo, error)
}

type CCClient struct {
	agent *sling.Sling
}

func (c *CCClient) Agent() *sling.Sling {
	return c.agent.New()
}

// Post :
func (c *CCClient) Post(path string) *sling.Sling {
	return c.Agent().Post(path)
}

// Get :
func (c *CCClient) Get(path string) *sling.Sling {
	return c.Agent().Get(path)
}

// NewCCClient :
func NewCCClient(ccConf *config.CCConf) *CCClient {
	var address = config.Conf.Client.CC.Address
	if ccConf != nil {
		address = ccConf.Address
	}

	return &CCClient{
		agent: sling.New().Base(fmt.Sprintf("http://" + address)),
	}
}

func (c *CCClient) Mkc(product, configVersion, serverVersion, serverName, configTemplate,
	configTemplatePath, businessTemplate, businessTemplatePath, commonFile, commonFilePath string, output bool) (*CCResponseInfo, error) {

	result := struct {
		Data *CCResponseInfo
	}{}
	response, err := c.Agent().Post("mkv").BodyJSON(CCRequestInfo{
		Product:              product,
		ConfigVersion:        configVersion,
		ServerVersion:        serverVersion,
		ServerName:           serverName,
		Output:               false,
		ConfigTemplate:       configTemplate,
		ConfigTemplatePath:   configTemplatePath,
		BusinessTemplate:     businessTemplate,
		BusinessTemplatePath: businessTemplatePath,
		CommonFile:           commonFile,
		CommonFilePath:       commonFilePath,
	}).Receive(result, result)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, err
	}
	return result.Data, nil
}

func PingCCClient(c *CCClient, ch chan struct{}) func(args ...interface{}) error {

	return func(args ...interface{}) error {
		var (
			//ticker = time.NewTicker(time.Duration(utils.ConvOrDefaultInt(config.GlobalConfig.ClientSet.CC.HeartBeat, 100)) * time.Second)
			ticker = time.NewTicker(2000 * time.Second)
			once   = &sync.Once{}
		)

		result := struct {
			Msg string `json:"msg"`
		}{}
		go func() {
			for {
				select {
				case <-ticker.C:
					response, err := c.Agent().Get("ping").Receive(&result, &result)
					if err != nil || response.StatusCode != 200 {
					}
					once.Do(func() {
						ch <- struct{}{}
					})

				}
			}
		}()
		return nil
	}

}
