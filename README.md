# platform-iac-cli

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Architecture](https://img.shields.io/badge/Architecture-Standard_Go_Layout-orange.svg)](docs/adr/0001-directory-structure.md)

## Overview

**platform-iac-cli** is a high-performance command-line interface designed for Senior DevSecOps and Platform Engineers. It serves as an orchestration layer for secure cloud infrastructure, primarily targeting AWS environments managed via Terragrunt and OpenTofu.

The project is built with **Agentic Integration** in mind, supporting the **Model Context Protocol (MCP)** to allow AI agents to interact with and reason about infrastructure states.

## Key Features

- **IaC Orchestration:** Seamless interaction with Terragrunt/OpenTofu stacks.
- **Security by Design:** Automated pre-apply checks using `tfsec` and `checkov`.
- **Advanced Networking:** Native support for AWS VPC Lattice service network orchestration.
- **Agentic Readiness:** JSON-structured output for indexing by LLM-based agents (MCP compatible).
- **Compliance:** Strict enforcement of FQCN for configuration management (e.g., using `ansible.build.systemd` in generated manifests).

## Tech Stack

- **Language:** Go 1.23+
- **CLI Framework:** Cobra
- **Configuration:** Viper
- **Logging:** Structured logging (`slog`)
- **Infrastructure:** AWS, Terragrunt, OpenTofu, VPC Lattice

## Architecture & Development

This project follows the **Standard Go Project Layout**.

### Architectural Decision Records (ADR)
All major technical decisions are documented in the `docs/adr/` directory.
- [ADR-0001: Choice of Directory Structure](docs/adr/0001-directory-structure.md)
- [ADR-0002: Logging Framework Choice](docs/adr/0002-logging-framework.md)

### Git Worktree Workflow
To maintain high velocity and clean context switching, this repository is optimized for `git worktree`:
- `main/` - Stable development branch.
- `v2-alpha/` - Experimental MCP features.

## Roadmap (Kanban / WIP Limited)

- [ ] **Phase 1: Foundation**
    - [x] CLI Core (Cobra/Viper)
    - [ ] `init` command: Project scaffolding for AWS/Terragrunt
    - [ ] `check` command: Static analysis integration
- [ ] **Phase 2: Agentic & Networking**
    - [ ] MCP Server implementation
    - [ ] VPC Lattice orchestration modules
    - [ ] Karpenter optimization for GPU nodes

## Getting Started

### Prerequisites
- Go 1.23+
- Terragrunt & OpenTofu
- AWS CLI configured

### Installation
```bash
git clone --bare <repository-url> .bare
git worktree add main main
cd main
go build -o bin/platform-cli ./cmd/platform-cli/
```

## Compliance Note

All generated Ansible roles and playbooks must use Full Qualified Collection Names (FQCN). Example: Use `ansible.build.systemd` instead of systemd.

© 2026 Platform Engineering Solutions. All rights reserved.
