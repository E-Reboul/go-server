
# Go Server API

This repository is an experiment in building a clean, idiomatic API server in Go, following best practices and conventions of the language.

## Features

- Modular logging system (per category, file-based, using zap)
- Organized project structure (middlewares, routes, types, server)
- Example of idiomatic Go code and comments

## Why Go?

- Modern, statically typed, and high-performance language
- Clear conventions and integrated tooling (formatting, testing, etc.)
- Well-suited for APIs and web services

## Getting Started

1. Clone the repository:
	```sh
	git clone https://github.com/E-Reboul/go-server.git
	```
2. Install dependencies:
	```sh
	go mod tidy
	```
3. Run the server:
	```sh
	go run main.go
	```

## Project Structure

- `server/` — Main server launch and router configuration
- `middlewares/logs/` — Logging system (initialization, writing, constants)
- `routes/healthcheck/` — Healthcheck route and handler
- `types/logs/` — Log category types

## About

This project is for experimentation and learning purposes only. It is not intended for production use.