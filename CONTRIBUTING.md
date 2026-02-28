# Contributing to go-common

Thank you for considering contributing to **go-common**! Every contribution is welcome, whether it's a bug report, a feature request, or a pull request.

## How to Report a Bug

1. Check the [existing issues](https://github.com/pmoscode/go-common/issues) to see if the bug has already been reported.
2. If not, open a new issue with:
   - A clear and descriptive title
   - Steps to reproduce the problem
   - Expected vs. actual behavior
   - Go version and OS

## How to Suggest a Feature

Open an issue with the **feature request** label. Describe the use case and why the feature would be useful.

## How to Submit a Pull Request

1. **Fork** the repository and create a new branch from `main`.
2. **Install dependencies:**
   ```bash
   go mod tidy
   ```
3. **Make your changes.** Follow the conventions below.
4. **Run the tests:**
   ```bash
   go test ./...
   ```
5. **Run static analysis** (if [staticcheck](https://staticcheck.dev/docs/getting-started/) is installed):
   ```bash
   staticcheck ./...
   ```
6. **Commit** with a clear message describing what you changed and why.
7. **Open a Pull Request** against `main`.

## Code Conventions

- **File naming:** Name the primary file in a package after the package itself (e.g. `filter/filter.go`), not `main.go`.
- **Error handling:** Return errors to the caller. Do not use `panic`, `log.Fatal`, or `fmt.Println` for error handling in library code.
- **Logging:** Use the `log` standard library for warnings. Do not use `fmt.Print*` for log output.
- **GoDoc:** All exported types, functions, and methods must have GoDoc comments.
- **Deprecation:** Use the standard Go `// Deprecated: ...` comment format.
- **Testing:** Use [testify](https://github.com/stretchr/testify) (`assert` / `require`) for assertions. Use `require` for setup errors and `assert` for value checks.

## Security

If you discover a security vulnerability, please **do not** open a public issue. Instead, refer to the [Security Policy](SECURITY.md) for instructions on responsible disclosure.

## License

By contributing, you agree that your contributions will be licensed under the same license as the project (see [LICENSE](LICENSE)).

