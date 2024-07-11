package cli

// identity provide unified methods for the parse func.
type identity interface {
	getName() string
	getEnv() string
	loadEnvValue()
}

// Parameter holds the parameter configuration and the final value.
type Parameter[T any] struct {
	Name         string
	Usage        string
	DefaultValue T
	EnvVarName   string
	setCli       bool
	setEnv       bool
	value        *T
	envGetFunc   func(key string, defaultVal T) T
}

// setFlagFunc adds the parameter to the flag context.
func (p *Parameter[T]) setFlagFunc(fn func(p *T, name string, value T, usage string)) {
	fn(p.value, p.Name, p.DefaultValue, p.Usage)
}

// setEnvironmentFunc sets the func, which fetches the value from the environment.
func (p *Parameter[T]) setEnvironmentFunc(fn func(key string, defaultVal T) T) {
	p.envGetFunc = fn
}

// GetValue return the final value for the parameter.
func (p *Parameter[T]) GetValue() *T {
	return p.value
}

// getName return the parameters name as part of the identity interface.
func (p *Parameter[T]) getName() string {
	return p.Name
}

// getEnv return the parameters env name as part of the identity interface.
func (p *Parameter[T]) getEnv() string {
	return p.EnvVarName
}

// loadEnvValue fetches the environment variable as part of the identity interface.
func (p *Parameter[T]) loadEnvValue() {
	value := p.envGetFunc(p.EnvVarName, p.DefaultValue)
	p.value = &value
}

// NewParameter return a new parameter instance.
func NewParameter[T any](name string, defaultValue T, usage, envVarName string) *Parameter[T] {
	return &Parameter[T]{
		Name:         name,
		DefaultValue: defaultValue,
		Usage:        usage,
		EnvVarName:   envVarName,
		setCli:       false,
		setEnv:       false,
		value:        new(T),
	}
}
