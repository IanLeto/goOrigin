package logic

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
	"goOrigin/config"
)

type NodeEntity struct {
	ID    int
	Name  string
	Value int
}

type ScriptAPISuite struct {
	suite.Suite
	conf *config.Config
}

// Directly sets the value of the field.
func updateNodeDirectly(a1, a2 *NodeEntity) {
	if a1.Value != a2.Value {
		a1.Value = a2.Value
	}
}

// Uses reflection to set the value of the field.
func updateNodeWithReflection(a1, a2 *NodeEntity) {
	valA1 := reflect.ValueOf(a1).Elem()
	valA2 := reflect.ValueOf(a2).Elem()

	for i := 0; i < valA1.NumField(); i++ {
		if valA1.Field(i).Interface() != valA2.Field(i).Interface() {
			valA1.Field(i).Set(valA2.Field(i))
		}
	}
}

func (s *ScriptAPISuite) SetupTest() {
	// Setup logic if needed
}

func (s *ScriptAPISuite) TestConfig() {
	// Test configuration if needed
}

// Add performance tests here.
func (s *ScriptAPISuite) TestUpdateNodeDirectly() {
	a1 := &NodeEntity{Value: 1}
	a2 := &NodeEntity{Value: 2}

	s.Run("UpdateNodeDirectly", func() {
		updateNodeDirectly(a1, a2)
		s.Equal(2, a1.Value)
	})
}

func (s *ScriptAPISuite) TestUpdateNodeWithReflection() {
	a1 := &NodeEntity{Value: 1}
	a2 := &NodeEntity{Value: 2}

	s.Run("UpdateNodeWithReflection", func() {
		updateNodeWithReflection(a1, a2)
		s.Equal(2, a1.Value)
	})
}

// Benchmark for direct field access.
func BenchmarkUpdateNodeDirectly(b *testing.B) {
	a1 := &NodeEntity{Value: 1}
	a2 := &NodeEntity{Value: 2}

	for i := 0; i < b.N; i++ {
		updateNodeDirectly(a1, a2)
	}
}

// Benchmark for field access using reflection.
func BenchmarkUpdateNodeWithReflection(b *testing.B) {
	a1 := &NodeEntity{Value: 1}
	a2 := &NodeEntity{Value: 2}

	for i := 0; i < b.N; i++ {
		updateNodeWithReflection(a1, a2)
	}
}
func TestScriptConfiguration(t *testing.T) {
	suite.Run(t, new(ScriptAPISuite))
}
