# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Build

- *(deps)* Bump golang.org/x/net from 0.37.0 to 0.38.0

### ⚙️ Miscellaneous

- Update go version
- Update deps
- Update deps
- Update go get command
- Update deps
- Update deps
- Update deps
- Update deps
- Update deps
- *(deps)* Update dependency ubuntu to v24
- *(deps)* Update actions/checkout action to v5
- *(deps)* Update actions/setup-go action to v6
- Update deps
- Update deps
- *(deps)* Update actions/checkout action to v6
- Add dependency audit workflow for vulnerability checks and go.mod verification
- Update Go version to 1.26.0 and task version to 3.48.0 in configuration files
- *(deps)* Update golangci/golangci-lint-action action to v9

### 🐛 Bug Fixes

- Typo
- Liter errors
- Liter errors
- Typo
- Non-constant format string
- Clean file paths and improve error messages in config loading

### 📚 Documentation

- Add CONTRIBUTING.md to guide new contributors on project participation
- Update README.md with command package details and examples

### 🔧 Refactoring

- Update HeartBeat to use context for cancellation and atomic operations
- Improve MQTT client implementation and error handling
- Simplify staticcheck command in Taskfile.yaml for improved maintainability
- Improve error handling and resource cleanup in file operations

### 🚀 Features

- Add GH Action test pipeline
- Implement CLI parameter manager with support for various types
- Add config package for loading environment variables and file-based configurations
- Add environment package for reading environment variables with type support
- Add filesystem package with FileExists function and tests
- Add filter package with generic filtering functionality and tests
- Add heartbeat package for interval-based function execution
- Enhance logging package with structured logging and severity levels
- Enhance MQTT client with TLS configuration and error handling
- Add process package with helper function to get executable name
- Add shutdown package for graceful application exit handling
- Add strings package with PrettyPrintJson and PrettyPrintYaml functions
- Add templates package for managing and rendering templates with a data structure
- Add yamlconfig package for loading YAML configurations with deprecation notice
- Add ResetForTesting function to enable isolated testing of shutdown hooks
- Implement command execution with logging and dry-run support

### 🧪 Testing

- Refactor tests to use testify assertions for improved readability
- Refactor logging tests to use testify assertions for consistency and clarity
- Add additional unit tests for all modules

## [0.10.0] - 2024-07-11

### ⚙️ Miscellaneous

- Update deps

### 🔧 Refactoring

- Remove export of internal variables

### 🚀 Features

- Add func to get float env values
- Add a basic cli manager
- Update README.md

## [0.9.1] - 2024-07-10

### ⚙️ Miscellaneous

- Update deps
- Update deps

### 🐛 Bug Fixes

- Nil error when shutdown is not properly initialized

## [0.9.0] - 2024-05-22

### ⚙️ Miscellaneous

- Add deprecation warning
- Add staticcheck linter ignore comment
- Remove unneeded "break" statement

### 🚀 Features

- Add staticcheck task
- Add configuration loader (file and environment) + update README.md

## [0.8.1] - 2024-05-14

### ⚙️ Miscellaneous

- Change task name

### 🐛 Bug Fixes

- Attach "_" to prefix to avoid issues with non prefix strings

### 📚 Documentation

- Adjust README.md

### 🚀 Features

- Load environment variables via prefix as map
- Add date and new line to log entry

## [0.8.0] - 2024-05-03

### ⚙️ Miscellaneous

- Replace deprecated method
- Add license
- Update deps
- Remove refactored stuff

### 📚 Documentation

- Add the first batch of code documentation
- Update go docs

### 🔧 Refactoring

- Move WithTlsCertificates to BrokerBuilder

### 🚀 Features

- Add exchangeable log writer
- Add struct logger and tests
- Add PrettyPrint for yaml and tests
- Add manager for go-templating engine (as generic impl)
- Add MQTT TLS configuration
- Add BrokerBuilder to set up broker host, credentials and tls

### 🧪 Testing

- Add logger tests
- Add tests for mqtt

## [0.7.0] - 2024-03-21

### ⚙️ Miscellaneous

- Update deps
- Add go test as task
- Adjust README.md

### 🐛 Bug Fixes

- *(shutdown)* Close channel when used

### 🚀 Features

- *(strings)* Add PrettyPrint function and according test
- *(filter)* Add generic filter and according test
- *(filesystem)* Add FileExists function and according test
- *(environment)* Add function to get defined environment values and according test
- *(logging)* Add "warning" level and rename "dryRun" to "extend"

### 🧪 Testing

- Add simple test

## [0.6.2-2] - 2024-03-14

### 🐛 Bug Fixes

- *(heartbeat)* Wrong "first time execution" logic part 2

## [0.6.2-1] - 2024-03-14

### 🐛 Bug Fixes

- *(heartbeat)* Wrong "first time execution" logic

## [0.6.2] - 2024-03-11

### 🐛 Bug Fixes

- *(heartbeat)* Deadlock when "noWait" set: when error on first loop

### 🚀 Features

- *(shutdown)* Reformat output

## [0.6.1-1] - 2024-03-11

### 🚀 Features

- *(logging)* Execute shutdown.Exit ob logger.Fatal call
- *(shutdown)* Add Exit functions and refactor command summary output

## [0.6.0] - 2024-03-07

### 🚀 Features

- *(heartbeat)* Add "runForever" function
- *(process)* Add "GetExecutableName" function
- *(logging)* Add simple logging

## [0.5.1] - 2024-02-28

### 🔧 Refactoring

- *(mqtt)* Prepare for module testing
- *(shutdown)* Replace native func definition with type
- *(heartbeat)* Add "noWait" option to run beat without waiting for "interval" first

### 🚀 Features

- Create README.md

## [0.5.0] - 2024-02-26

### 🚀 Features

- *(mqtt)* Add new mqtt constructor

## [0.4.0] - 2024-02-26

### ⚙️ Miscellaneous

- Update dependencies
- Update dependencies
- Update dependencies

### 🚀 Features

- Add Taskfile with update-deps task
- Run pub / sub async

## [0.3.0] - 2023-07-16

### Build

- *(deps)* Bump golang.org/x/net

### ⚙️ Miscellaneous

- Remove gitlab ci config file
- Update dependencies
- Update package name

### 🐛 Bug Fixes

- *(deps)* Update module github.com/eclipse/paho.mqtt.golang to v1.4.3

## [0.2.0] - 2023-03-19

### 🚀 Features

- Add yaml config loader

## [0.1.1] - 2022-12-16

### 🐛 Bug Fixes

- Fix imports after module refactor

## [0.0.1] - 2022-12-16

### 🚀 Features

- Initial commit

<!-- generated by git-cliff -->
