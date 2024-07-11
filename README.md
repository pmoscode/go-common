# GO-Common library

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

A small GO library, which may come handy for recurring tasks or simple use cases.

## Features

A short briefing of the single modules:

### Cli

The cli manager can help to get parameters, which are passed as cli arguments and also are provided as environment variables.

A simple example:

```go 
func LoadCliParameters() {
    file := NewParameter[string]("file", "file.yaml", "description", "ENV_VAR_FILE")
    dryRun := NewParameter[bool]("dryRun", false, "description", "ENV_VAR_DRY_RUN")
    
    mgr := New()
    mgr.AddStringParameter(file)
    mgr.AddBoolParameter(dryRun)
    mgr.Parse()
    
    fmt.Println(*file.GetValue())
    fmt.Println(*dryRun.GetValue())
}
```

### Config

Loads configuration into a struct. Config can come from a file or the os environment or both.
The "env" tag must be used to configure the environment loading.

Example:
```go
type ConfigData struct {
	Data1   string             `env:"name=DATA_ONE_STR"`
	Data2   string             `env:""`
	Data3   string             `env:"self"`
    Data4   int                `env:"name=DATA_ONE_INT"`
    Data5   map[string]string  `env:"prefix=GMD,cutoff=false"`
    Data6   map[string]float32 `env:"prefix=DAT,cutoff=true"`
    Data7   map[string]string  `env:"prefix=CUT"`
}
```

There are two main categories:

- single (named) values -> string, int, bool, ...
- multiple named values -> map[string]string, map[string]int, map[string]bool, ...

#### Single values

Tag format: 
- `env:"name=DATA_ONE_STR,default=default string"`
- `env:"name=DATA_ONE_STR"`
- `env:"self,default=default string"`

Explanation:

- "name": defines the name of the environment variable to be loaded into this field
- "self": if "name" is not provided or the "env" tag is empty, the snake_cased field name is used
- "default": defines a default value, if the environment variable is not found

#### Multiple named values

Tag format:
- `env:"prefix=GMD,cutoff=false"`
- `env:"prefix=DAT,cutoff=true"`
- `env:"prefix=CUT"`

Explanation:

- "prefix": defines the prefix of the environment variables, which are going to be grouped. NOTE: There is always an "_" at the end of the prefix. If it's missing here, it will be attached automatically.
- "cutoff": defines, if the prefix itself will be cut off. Ex.: if true, "PRE_DATA_1" => "DATA_1". If just "cutoff" is provided, it's treated as "cutoff=true" internally.

#### Order

The tags "name", "self" and "prefix" are ordered:

1. name
2. self
3. prefix

So, if "name" is found, everything else (self, prefix) is ignored.

NOTE: "default" is only considered when "name" or "self" is provided.
"cutoff" is only considered when "prefix" is provided.

### Environment
This provides a simple way to fetch environment variables with defined default values, if they can't be found.
There are three types:
- Get the raw string
- Get an int
- Get an float
- Get a bool

### Filesystem
This provides one function at the moment:
- Check, if a files exists and if it is really a file (not a directory)

### Filter
This is a helper package to filter a slice. You provide one or more functions to set the matching boundaries.

### Heartbeat
The heartbeat package is some kind of timer, which executes a given function at an interval.
It can be configured, if the first execution should start immediately or after the first interval.

### Logging
The logging package is a simple logger, where you can configure the appearance of the log entry.
You can also define the io.Writer for it.

### MQTT
The mqtt package is a wrapper for the Paho Mqtt client.
It simplifies the configuration and usage of the publish and subscribe "actions".

### Process
Here at the moment only one function is available:
- Get the name of the current executable. And only the name, without path.

### Shutdown
The shutdown package executes defined function when the app is exiting.
Either when killed (code 15) or exits normally. Depending on the configuration.

### String
In this package you will find two functions:
- One to pretty format JSON structs
- One to pretty format YAML structs

### Templates
This package contains a TemplateManager. The purpose is to make GO templates (via the GO templating engine)
accessible via name. So called "named templates".
Every template in this manager can be populated with defined options. Currently only with custom template functions.

### Yamlconfig
The yamlconfig package loads YAML config files into a given struct.

## Installation

Import the module:

```bash
 go get github.com/pmoscode/go-common
```
If you want to get a specific version:

```bash
 go get github.com/pmoscode/go-common@v0.7.0
```
    
## Usage/Examples

Most packages have tests, so you can see the usage there.
But for most packages, even that is not necessary.

## Documentation

All packages are populated with go docs. So you should take a look here:
https://pkg.go.dev/github.com/pmoscode/go-common


## License

[MIT](https://choosealicense.com/licenses/mit/)

