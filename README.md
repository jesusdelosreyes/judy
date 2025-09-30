
<div align="center">
  <br />
  <h1>Judy 🌻</h1>
  <p><b>A friendly command-line assistant for QA and SDET teams to streamline testing workflows.</b></p>
  <br />
</div>

<p align="center">
  <a href="https://github.com/jesusdelosreyes/judy/actions/workflows/build.yml">
    <img src="https://github.com/jesusdelosreyes/judy/actions/workflows/build.yml/badge.svg" alt="build status">
  </a>
  <a href="https://goreportcard.com/report/github.com/jesusdelosreyes/judy">
    <img src="https://goreportcard.com/badge/github.com/jesusdelosreyes/judy" alt="go report card">
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="license">
  </a>
</p>

## Overview

Judy is a CLI tool built with Go that automates the repetitive tasks in the testing lifecycle. It helps you scaffold new projects, manage Dockerized test environments, and run tests from a single, consistent interface, optimizing the developer's "inner loop".

## Table of Contents

- [✨ Features](#-features)
- [🚀 Installation](#-installation)
- [🚀 Quick Start](#-quick-start)
- [🎮 Usage](#-usage)
- [⚙️ Configuration](#️-configuration)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)

---

## ✨ Features

- **Project Scaffolding**: Create standardized Go API test project structures in seconds.
- **Environment Management**: Spin up and tear down Dockerized environments with simple, reliable commands.
- **Test Orchestration**: Run k6 performance tests directly from the CLI.
- **Unified Workflow**: Provides a single, easy-to-remember interface for common QA tasks.
- **Cross-Platform**: Built as a single binary that runs on Windows, macOS, and Linux with no dependencies.

---

## 🚀 Installation

### Using `go install` (Recommended)
With [Go](https://go.dev/doc/install) (v1.21+) configured on your system:
```bash
go install [github.com/jesusdelosreyes/judy@latest](https://github.com/jesusdelosreyes/judy@latest)
````

### From GitHub Releases

Download the pre-compiled binary for your operating system from the [Releases Page](https://www.google.com/search?q=https://github.com/jesusdelosreyes/judy/releases).

### From Source

```bash
git clone [https://github.com/jesusdelosreyes/judy.git](https://github.com/jesusdelosreyes/judy.git)
cd judy
go build -o judy .
# Move the 'judy' binary to a directory in your PATH
```

-----

## 🚀 Quick Start

Get your first test project up and running in under 60 seconds.

```bash
# 1. Create a new test project structure
judy setup-tests "My First Judy Project" --type go-api

# 2. Navigate into your new project
cd "My First Judy Project"

# 3. Download dependencies defined in the template
make tidy

# 4. Run the example test to verify everything works!
make test
```

-----

## 🎮 Usage

### General Commands

| Command | Description |
| :--- | :--- |
| `judy version` | Checks the installed version of Judy. |
| `judy help` | Shows help for all available commands. |

### Scaffolding a New Project (`setup-tests`)

Initializes the full structure for a new Go API test project.

**Command:**

```bash
judy setup-tests "My Payments API Tests" --type go-api
```

| Parameter | Type | Required | Description |
| :--- | :--- | :--- | :--- |
| `projectName` | Argument | Yes | The name of the project. Use quotes for spaces. |
| `--type`, `-t` | Flag | Yes | Specifies the project type (currently supports `go-api`). |

### Environment Management (`deploy-env` / `destroy-env`)

These commands manage Dockerized environments and assume a `docker-compose.yml` file exists in the current directory.

| Command | Description |
| :--- | :--- |
| `judy deploy-env` | Spins up the environment (`docker compose up -d`). |
| `judy destroy-env`| Tears down the environment (`docker compose down -v`). |

-----

## ⚙️ Configuration

Judy can be configured via a `judy.yml` file in your project's root directory to coordinate multi-repository workflows.

**Example `judy.yml`:**

```yaml
# judy.yml
# This file defines all the components of our testing system.
project_name: "E-Commerce Promotions Feature"

# Define components and their local paths
components:
  - name: promotions-api
    path: ../promotions-api # Path to the microservice repo
    type: service
  - name: e2e-tests
    path: ../e2e-automation-suite # Path to the E2E test repo
    type: test-suite

# Configure commands to target specific components
commands:
  deploy-env:
    compose_file: ../promotions-api/docker-compose.yml
  run-e2e:
    target: e2e-tests
    test_command: "npm test"
```

-----

## 🤝 Contributing

Contributions are welcome\! For major changes, please open an issue first to discuss what you would like to change. Please make sure to update tests as appropriate.

## 📄 License

This project is licensed under the MIT License. See the `LICENSE` file for details.

```
```