# GO-Common library

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

A small GO library, which may come handy for recurring tasks or simple use cases.

## Features

A short briefing of the single modules:

### Environment
This provides a simple way to fetch environment variables with defined default values, if they can't be found.
There are three types:
- Get the raw string
- Get an int
- Get a bool

### Filesystem
This provides one function at the moment:
- Check, if a files exists and if it is really a file (not a directory)
- 
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

