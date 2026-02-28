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
    file := cli.NewParameter[string]("file", "file.yaml", "description", "ENV_VAR_FILE")
    dryRun := cli.NewParameter[bool]("dryRun", false, "description", "ENV_VAR_DRY_RUN")
    
    mgr := cli.New()
    mgr.AddStringParameter(file)
    mgr.AddBoolParameter(dryRun)
    mgr.Parse()
    
    fmt.Println(*file.GetValue())
    fmt.Println(*dryRun.GetValue())
}
```

### Command

The command package provides a way to execute external commands with logging and dry-run support.
Parameters can be masked so that sensitive values (e.g. passwords) are hidden in log output.

> **WARNING:** This package passes commands directly to `exec.Command` without sanitization.
> The caller **must** ensure all inputs are trusted. Do not pass unsanitized user input.

A simple example:

```go
logger := logging.NewLogger()

// Direct usage
cmd := command.NewCommand(logger, false)
params := command.NewParameters(
    command.WithCommand("echo"),
    command.WithValue("hello"),
    command.WithParam("--verbose"),
    command.WithValueMasked("s3cret"), // will be shown as *** in logs
)
err := cmd.Execute(params)

// Or via the Manager (builder pattern)
mgr := command.NewCommandManager(logger, true) // true = dry-run
mgr.AddParameter(command.WithCommand("echo"))
mgr.AddParameter(command.WithValue("hello"))
err = mgr.ExecuteCommand()
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
There are functions for the following types:
- `GetEnv` – raw string
- `GetEnvInt` – int
- `GetEnvFloat32` – float32
- `GetEnvFloat64` – float64
- `GetEnvBool` – bool
- `GetEnvMap` – all variables with a given prefix as `map[string]string`

### Filesystem
This provides one function at the moment:
- Check, if a file exists and if it is really a file (not a directory)

### Filter
This is a generic helper package to filter a slice. You provide one or more filter functions to set the matching boundaries.

```go
items := []*MyStruct{ /* ... */ }
f := filter.NewFilter(items)
result := f.Filter(
    func(val MyStruct) bool { return val.Age > 18 },
    func(val MyStruct) bool { return val.Active },
)
```

### Heartbeat
The heartbeat package is some kind of timer, which executes a given function at an interval.
It supports `context.Context` for clean cancellation.
It can be configured, if the first execution should start immediately or after the first interval.

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

hb, err := heartbeat.New(5*time.Second, myFunc, heartbeat.WithNoWait())
if err != nil {
    log.Fatal(err)
}

// Non-blocking:
hb.Run(ctx)

// Or blocking (until context is cancelled):
// hb.RunForever(ctx)
```

### Logging
The logging package is a simple logger, where you can configure the appearance of the log entry.
You can also define the `io.Writer` for it.

Available options:
- `WithName` – set the logger name
- `WithSeverity` – set the log level (`INFO`, `DEBUG`, `WARNING`, `TRACE`)
- `WithLogWriter` – set a custom `io.Writer` (default: `os.Stderr`)
- `WithExtend` – add an extension string to the log header
- `WithNameSpacing` / `WithSeveritySpacing` – control column widths

Structs can be logged in JSON or YAML format via `InfoStruct`.

### MQTT
The mqtt package is a wrapper for the [Paho MQTT client](https://github.com/eclipse/paho.mqtt.golang).
It simplifies the configuration and usage of publish and subscribe actions.

The client supports TLS configuration via `NewBrokerBuilder` and clean shutdown via `context.Context`:

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

brokerOpt, err := mqtt.NewBrokerBuilder().
    WithHost("localhost").
    WithPort(1883).
    WithProtocol(mqtt.MqttTcp).
    Build()

client := mqtt.NewClient(mqtt.WithClientId("my-app"), brokerOpt)
client.Connect()
client.LoopForever(ctx) // blocks until context is cancelled
```

### Process
Here at the moment only one function is available:
- `GetExecutableName()` – returns the name of the current executable (without path). Returns an error if the name cannot be determined.

### Shutdown
The shutdown package executes registered functions when the app is exiting.
It hooks on SIGTERM (code 15) for graceful shutdown and also supports explicit exit via `Exit()`.

```go
shutdown.GetObserver().AddCommand(func() error {
    fmt.Println("Cleaning up...")
    return nil
})

// For panic recovery:
defer shutdown.ExitOnPanic()
```

### Strings
In this package you will find two functions:
- `PrettyPrintJson` – pretty format any struct as JSON
- `PrettyPrintYaml` – pretty format any struct as YAML

### Templates
This package contains a generic `TemplateManager`. The purpose is to make GO templates (via the GO templating engine)
accessible via name. So called "named templates".
Every template in this manager can be populated with defined options. Currently only with custom template functions.

> **Note:** This package uses `text/template` which does **not** perform HTML escaping.
> Do not use it for HTML output. Use `html/template` from the standard library instead.

### Yamlconfig

> **Deprecated:** Use the `config` package with `formats.ParseYamlConfig` instead.

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

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Documentation

All packages are populated with go docs. So you should take a look here:
https://pkg.go.dev/github.com/pmoscode/go-common


## License

[MIT](https://choosealicense.com/licenses/mit/)

