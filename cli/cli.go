// Package cli provides a basic approach to manage parameters, which come either as part of the command line or/and via environment variable.
package cli

import (
	"flag"
	"strconv"

	"github.com/pmoscode/go-common/environment"
)

// Manager holds the main entities.
type Manager struct {
	parameters []identity
	usedFlags  map[string]struct{}
}

// addParameter wires up a Parameter with its flag- and environment-lookup
// functions and registers it with the manager. Shared by all Add*Parameter
// methods to avoid duplicating that logic per type.
func addParameter[T any](m *Manager, param *Parameter[T], flagFn func(p *T, name string, value T, usage string), envFn func(key string, defaultVal T) T) {
	param.setFlagFunc(flagFn)
	param.setEnvironmentFunc(envFn)
	m.parameters = append(m.parameters, param)
}

// AddStringParameter adds a string parameter to the manager. It sets the flagFunc and the envFunc.
func (m *Manager) AddStringParameter(param *Parameter[string]) {
	addParameter(m, param, flag.StringVar, environment.GetEnv)
}

// AddIntParameter adds a int parameter to the manager. It sets the flagFunc and the envFunc.
func (m *Manager) AddIntParameter(param *Parameter[int]) {
	addParameter(m, param, flag.IntVar, environment.GetEnvInt)
}

// AddFloat64Parameter adds a float64 parameter to the manager. It sets the flagFunc and the envFunc.
func (m *Manager) AddFloat64Parameter(param *Parameter[float64]) {
	addParameter(m, param, flag.Float64Var, environment.GetEnvFloat64)
}

// AddFloat32Parameter adds a float32 parameter to the manager. It sets the flagFunc and the envFunc.
//
// The standard library's "flag" package has no native Float32Var, so this uses
// flag.Var with a small internal adapter type.
func (m *Manager) AddFloat32Parameter(param *Parameter[float32]) {
	addParameter(m, param, float32Var, environment.GetEnvFloat32)
}

// AddBoolParameter adds a bool parameter to the manager. It sets the flagFunc and the envFunc.
func (m *Manager) AddBoolParameter(param *Parameter[bool]) {
	addParameter(m, param, flag.BoolVar, environment.GetEnvBool)
}

// float32Value adapts a float32 to the flag.Value interface, since the standard
// library's "flag" package only provides a Float64Var.
type float32Value float32

func (f *float32Value) String() string {
	return strconv.FormatFloat(float64(*f), 'g', -1, 32)
}

func (f *float32Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	*f = float32Value(v)
	return nil
}

// float32Var registers a float32 flag, matching the signature expected by
// Parameter.setFlagFunc (analogous to flag.StringVar/flag.IntVar/...).
func float32Var(p *float32, name string, value float32, usage string) {
	*p = value
	flag.Var((*float32Value)(p), name, usage)
}

// Parse does the whole execution. It handles the right order of the parameters origin. (cli arf or env var)
func (m *Manager) Parse() {
	flag.Parse()
	flag.Visit(m.visit())

	for _, item := range m.parameters {
		if _, used := m.usedFlags[item.getName()]; !used {
			item.loadEnvValue()
		}
	}
}

// visit is passed to the flag.Visit func to mark the parameters, which were set via command line.
func (m *Manager) visit() func(f *flag.Flag) {
	return func(flag *flag.Flag) {
		m.usedFlags[flag.Name] = struct{}{}
	}
}

// New returns a new cli manager instance
func New() *Manager {
	return &Manager{
		usedFlags: make(map[string]struct{}),
	}
}
