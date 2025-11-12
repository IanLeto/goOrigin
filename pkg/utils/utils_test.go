package utils_test

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"

	"github.com/stretchr/testify/suite"
)

// TemplateSuite :
type TemplateSuite struct {
	suite.Suite
	templateStr string
	tmpl        *template.Template
}

func (s *TemplateSuite) SetupTest() {
	s.templateStr = `{{$pod := (print .name ".*")}}{{if eq .kind "DaemonSet"}}{{$pod = (print .name "-[a-z0-9]{5}")}}{{end}}{{if eq .kind "StatefulSet"}}{{$pod = (print .name "-[0-9]{1,3}")}}{{end}}{{if eq .kind "Deployment"}}{{$pod = (print .name "-[a-z0-9]{7,10}-[a-z0-9]{5}")}}{{end}}sum by (pod, pod_ip, host_ip) (last_over_time(kube_pod_info{namespace="{{.namespace}}", pod=~"{{$pod}}", created_by_kind="{{.kind}}"}[7d]) unless kube_pod_container_info)`

	var err error
	s.tmpl, err = template.New("prometheus").Parse(s.templateStr)
	s.Require().NoError(err, "Template should parse without error")
}

// TestDeploymentTemplate : 测试 Deployment 类型的模板渲染
func (s *TemplateSuite) TestDeploymentTemplate() {
	data := map[string]string{
		"namespace": "default",
		"name":      "nginx",
		"kind":      "Deployment",
	}

	var buf bytes.Buffer
	err := s.tmpl.Execute(&buf, data)
	s.NoError(err, "Template execution should not error")

	expected := `sum by (pod, pod_ip, host_ip) (last_over_time(kube_pod_info{namespace="default", pod=~"nginx-[a-z0-9]{7,10}-[a-z0-9]{5}", created_by_kind="Deployment"}[7d]) unless kube_pod_container_info)`
	s.Equal(expected, buf.String(), "Deployment template should render correctly")
}

// TestDaemonSetTemplate : 测试 DaemonSet 类型的模板渲染
func (s *TemplateSuite) TestDaemonSetTemplate() {
	data := map[string]string{
		"namespace": "kube-system",
		"name":      "fluentd",
		"kind":      "DaemonSet",
	}

	var buf bytes.Buffer
	err := s.tmpl.Execute(&buf, data)
	s.NoError(err, "Template execution should not error")

	expected := `sum by (pod, pod_ip, host_ip) (last_over_time(kube_pod_info{namespace="kube-system", pod=~"fluentd-[a-z0-9]{5}", created_by_kind="DaemonSet"}[7d]) unless kube_pod_container_info)`
	s.Equal(expected, buf.String(), "DaemonSet template should render correctly")
}

// TestStatefulSetTemplate : 测试 StatefulSet 类型的模板渲染
func (s *TemplateSuite) TestStatefulSetTemplate() {
	data := map[string]string{
		"namespace": "database",
		"name":      "mysql",
		"kind":      "StatefulSet",
	}

	var buf bytes.Buffer
	err := s.tmpl.Execute(&buf, data)
	s.NoError(err, "Template execution should not error")

	expected := `sum by (pod, pod_ip, host_ip) (last_over_time(kube_pod_info{namespace="database", pod=~"mysql-[0-9]{1,3}", created_by_kind="StatefulSet"}[7d]) unless kube_pod_container_info)`
	s.Equal(expected, buf.String(), "StatefulSet template should render correctly")
}

// TestUnknownKindTemplate : 测试未知类型的模板渲染（应该使用默认的 .* 匹配）
func (s *TemplateSuite) TestUnknownKindTemplate() {
	data := map[string]string{
		"namespace": "default",
		"name":      "custom-app",
		"kind":      "Job",
	}

	var buf bytes.Buffer
	err := s.tmpl.Execute(&buf, data)
	s.NoError(err, "Template execution should not error")

	expected := `sum by (pod, pod_ip, host_ip) (last_over_time(kube_pod_info{namespace="default", pod=~"custom-app.*", created_by_kind="Job"}[7d]) unless kube_pod_container_info)`
	s.Equal(expected, buf.String(), "Unknown kind template should render with default pattern")
}

// TestEmptyNameTemplate : 测试空名称的情况
func (s *TemplateSuite) TestEmptyNameTemplate() {
	data := map[string]string{
		"namespace": "default",
		"name":      "",
		"kind":      "Deployment",
	}

	var buf bytes.Buffer
	err := s.tmpl.Execute(&buf, data)
	s.NoError(err, "Template execution should not error")

	expected := `sum by (pod, pod_ip, host_ip) (last_over_time(kube_pod_info{namespace="default", pod=~"-[a-z0-9]{7,10}-[a-z0-9]{5}", created_by_kind="Deployment"}[7d]) unless kube_pod_container_info)`
	s.Equal(expected, buf.String(), "Empty name template should still render")
	fmt.Println(buf.String())
}

// TestSpecialCharactersInName : 测试名称中包含特殊字符
func (s *TemplateSuite) TestSpecialCharactersInName() {
	data := map[string]string{
		"namespace": "test-ns",
		"name":      "app-with-dash",
		"kind":      "StatefulSet",
	}

	var buf bytes.Buffer
	err := s.tmpl.Execute(&buf, data)
	s.NoError(err, "Template execution should not error")

	expected := `sum by (pod, pod_ip, host_ip) (last_over_time(kube_pod_info{namespace="test-ns", pod=~"app-with-dash-[0-9]{1,3}", created_by_kind="StatefulSet"}[7d]) unless kube_pod_container_info)`
	s.Equal(expected, buf.String(), "Name with special characters should render correctly")
}

// TestAllKindsTable : 表格驱动测试，测试所有 kind 类型
func (s *TemplateSuite) TestAllKindsTable() {
	testCases := []struct {
		name     string
		input    map[string]string
		expected string
	}{
		{
			name: "Deployment with nginx",
			input: map[string]string{
				"namespace": "default",
				"name":      "nginx",
				"kind":      "Deployment",
			},
			expected: `sum by (pod, pod_ip, host_ip) (last_over_time(kube_pod_info{namespace="default", pod=~"nginx-[a-z0-9]{7,10}-[a-z0-9]{5}", created_by_kind="Deployment"}[7d]) unless kube_pod_container_info)`,
		},
		{
			name: "DaemonSet with fluentd",
			input: map[string]string{
				"namespace": "logging",
				"name":      "fluentd",
				"kind":      "DaemonSet",
			},
			expected: `sum by (pod, pod_ip, host_ip) (last_over_time(kube_pod_info{namespace="logging", pod=~"fluentd-[a-z0-9]{5}", created_by_kind="DaemonSet"}[7d]) unless kube_pod_container_info)`,
		},
		{
			name: "StatefulSet with postgres",
			input: map[string]string{
				"namespace": "database",
				"name":      "postgres",
				"kind":      "StatefulSet",
			},
			expected: `sum by (pod, pod_ip, host_ip) (last_over_time(kube_pod_info{namespace="database", pod=~"postgres-[0-9]{1,3}", created_by_kind="StatefulSet"}[7d]) unless kube_pod_container_info)`,
		},
		{
			name: "CronJob as unknown kind",
			input: map[string]string{
				"namespace": "batch",
				"name":      "backup",
				"kind":      "CronJob",
			},
			expected: `sum by (pod, pod_ip, host_ip) (last_over_time(kube_pod_info{namespace="batch", pod=~"backup.*", created_by_kind="CronJob"}[7d]) unless kube_pod_container_info)`,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			var buf bytes.Buffer
			err := s.tmpl.Execute(&buf, tc.input)
			s.NoError(err, "Template execution should not error for %s", tc.name)
			s.Equal(tc.expected, buf.String(), "Template output mismatch for %s", tc.name)
		})
	}
}

// TestTemplateRun : 运行测试套件
func TestTemplateRun(t *testing.T) {
	suite.Run(t, new(TemplateSuite))
}
