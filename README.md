# Golang Boilerplate

This repository provides a boilerplate setup for building applications in Golang using the Fiber web framework, Cobra CLI for command-line interface tools, and Makefile for managing build and development commands.

## Features

- **Fiber**: A fast and lightweight HTTP framework for Golang, providing an expressive API and middleware support.
- **Cobra**: A powerful CLI library for creating command-line applications in Go.
- **Makefile**: Automates common tasks such as building, testing, and cleaning the project.

## Getting Started

### Prerequisites

- Go 1.18 or later
- Make

### Installation

1. Clone the repository:
```bash
git clone https://github.com/stnss/dealls-interview.git
cd go-boilerplate
```
2. Install Dependency
```bash
make install
```
3. Running Tests
```bash
make test
```
4. Generate Coverage Report
```bash
make test-cover
```

## Usage
To create commands using Cobra, define your command structures and logic in the cmd directory. Use the main.go file to initialize and run your CLI application.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request.