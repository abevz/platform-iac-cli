# ADR-0002: Logging Framework Choice

## Status
Accepted

## Date
2026-01-07

## Context
The `platform-iac-cli` tool requires a structured logging system to:
1. Provide human-readable output for debugging during manual execution.
2. Provide machine-readable (JSON) output for integration with the Model Context Protocol (MCP) and external AI agents.
3. Minimize external dependencies to keep the binary lightweight and secure.
4. Measure command execution latency (Lead Time) as per KMP requirements.

## Decision
We will use the **`slog`** package from the Go Standard Library (introduced in Go 1.21).

## Rationale
- **Zero Dependencies:** Using `slog` eliminates the need for third-party libraries like `zap` or `logrus`, reducing the supply chain attack surface.
- **Performance:** `slog` is designed with performance in mind, offering near-zero overhead for structured logging.
- **Standardization:** As a part of the standard library, it ensures long-term compatibility and follows Go's idiomatic patterns.
- **Agentic Compatibility:** `slog.JSONHandler` provides native JSON formatting, which is essential for indexing logs by LLM agents and MCP-compatible systems.
- **Observability:** `slog` allows easy implementation of custom handlers for capturing telemetry (e.g., Lead Time metrics for Kanban analysis).

## Implementation Details
- **Default Handler:** `slog.NewTextHandler` for standard terminal output.
- **Machine Mode:** Use a global flag `--json` to switch to `slog.NewJSONHandler`.
- **Level Management:** Default level set to `Info`. `Debug` level to be activated via `--verbose` flag.

## Consequences
- Requires Go 1.21 or higher (Project uses 1.23+, so no impact).
- Developers must use structured key-value pairs instead of formatted strings (e.g., `slog.Info("msg", "key", "value")` instead of `fmt.Printf`).
