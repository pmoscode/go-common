// Package cli provides a basic approach to manage parameters, which come either as part of the command line or/and via environment variable.
package cli

import (
	"flag"
	"github.com/pmoscode/go-common/environment"
	"slices"
)

// Manager holds the main entities.
type Manager struct {
	parameters []identity
	usedFlags  []string
}

// AddStringParameter adds a string parameter to the manager. It sets the flagFunc and the envFunc.
func (m *Manager) AddStringParameter(param *Parameter[string]) {
	param.setFlagFunc(flag.StringVar)
	param.setEnvironmentFunc(environment.GetEnv)
	m.parameters = append(m.parameters, param)
}

// AddIntParameter adds a int parameter to the manager. It sets the flagFunc and the envFunc.
func (m *Manager) AddIntParameter(param *Parameter[int]) {
	param.setFlagFunc(flag.IntVar)
	param.setEnvironmentFunc(environment.GetEnvInt)
	m.parameters = append(m.parameters, param)
}

// AddFloat64Parameter adds a float64 parameter to the manager. It sets the flagFunc and the envFunc.
func (m *Manager) AddFloat64Parameter(param *Parameter[float64]) {
	param.setFlagFunc(flag.Float64Var)
	param.setEnvironmentFunc(environment.GetEnvFloat64)
	m.parameters = append(m.parameters, param)
}

// AddBoolParameter adds a bool parameter to the manager. It sets the flagFunc and the envFunc.
func (m *Manager) AddBoolParameter(param *Parameter[bool]) {
	param.setFlagFunc(flag.BoolVar)
	param.setEnvironmentFunc(environment.GetEnvBool)
	m.parameters = append(m.parameters, param)
}

// Parse does the whole execution. It handles the right order of the parameters origin. (cli arf or env var)
func (m *Manager) Parse() {
	flag.Parse()
	flag.Visit(m.visit())

	for _, item := range m.parameters {
		if !slices.Contains(m.usedFlags, item.getName()) {
			item.loadEnvValue()
		}
	}
}

// visit is passed to the flag.Visit func to mark the parameters, which were set via command line.
func (m *Manager) visit() func(f *flag.Flag) {
	return func(flag *flag.Flag) {
		m.usedFlags = append(m.usedFlags, flag.Name)
	}
}

// New returns a new cli manager instance
func New() *Manager {
	return &Manager{
		parameters: make([]identity, 0),
		usedFlags:  make([]string, 0),
	}
}
