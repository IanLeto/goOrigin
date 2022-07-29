package config_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

// ViperConfigurationSuite :
type ViperConfigurationSuite struct {
	suite.Suite
}

func (s *ViperConfigurationSuite) SetupTest() {
}

// TestMarshal :
func (s *ViperConfigurationSuite) TestConfig() {

}

func (s ViperConfigurationSuite) TestMySqlBackendConfig() {
	var t = map[string]interface{}{
		"apiVersion": "apps/v1",
		"kind":       "Deployment",
		"metadata": map[string]interface{}{
			"name": "demo-deployment",
		},
		"spec": map[string]interface{}{
			"replicas": 2,
			"selector": map[string]interface{}{
				"matchLabels": map[string]interface{}{
					"app": "demo",
				},
			},
			"template": map[string]interface{}{
				"metadata": map[string]interface{}{
					"labels": map[string]interface{}{
						"app": "demo",
					},
				},

				"spec": map[string]interface{}{
					"containers": []map[string]interface{}{
						{
							"name":  "web",
							"image": "nginx:1.12",
							"ports": []map[string]interface{}{
								{
									"name":          "http",
									"protocol":      "TCP",
									"containerPort": 80,
								},
							},
						},
					},
				},
			},
		},
	}
	res, err := json.Marshal(t)
	s.NoError(err)
	fmt.Println(string(res))

	//s.Equal("localhost:3306", config..Backend.MySqlBackendConfig.Address)

}

// TestViperConfiguration :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(ViperConfigurationSuite))
}
