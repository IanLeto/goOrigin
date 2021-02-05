package client

import (
	"github.com/dghubble/sling"
	"goOrigin/conf"
)

type APIClient interface {
	Mkc() CCResponseInfo
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
func NewCCClient() *CCClient {
	return &CCClient{
		agent: sling.New().Base(conf.Conf.Client.CC.Address),
	}
}

func (c CCClient) Mkc(product, configVersion, serverVersion, serverName, configTemplate,
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
