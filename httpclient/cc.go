package httpclient

import (
	"github.com/dghubble/sling"
	"goOrigin/config"
	"goOrigin/logging"
	"goOrigin/utils"
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
	var address = config.GlobalConfig.Client.CC.Address
	if ccConf != nil {
		address = ccConf.Address
	}

	return &CCClient{
		agent: sling.New().Base(address),
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

func PingCCClient(c *CCClient) func(args ...interface{}) error {
	return func(args ...interface{}) error {
		var (
			logger = logging.GetStdLogger()
			ticker = time.NewTicker(time.Duration(utils.ConvOrDefaultInt(config.GlobalConfig.Client.CC.HeartBeat, 100)) * time.Second)
		)

		result := struct {
		}{}
		for {
			select {
			case <-ticker.C:
				response, err := c.Agent().Get("/ping").Receive(result, result)
				if err != nil || response.StatusCode != 200 {
					logger.Errorf("cc 启动失败 %v", err)
				}
				return err
			}
		}

	}

}
