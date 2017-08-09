# api example

1. [Description](#description)
2. [Install, Run, Usage](#install-run-usage)
3. [Features](#features)
4. [Dependencies](#dependencies)
5. [Architecture Structure and Decisions](#architecture-structure-and-decisions)
6. [API Endpoints](#api-endpoints)
7. [Debugging Hints](#debugging-hints)
8. [Manifest Examples](#manifest-examples)
9. [Links to external Documentation](#links-to-external-documentation)
10. [TODOs](#todos)
11. [Nice to have](#nice-to-have)
12. [Known Issues](#known-issues)

## Desctiption
`apiexample` is a simple web server of how to implement a lightweight API written in golang.
It is designed to either give back it's own state (always ok) or the state of a similar application on another server.

## Install, Run, Usage
Clone the repo:
'go get github.com/phartz/apiexample'

To compile simply run `got get` after that `go build main.go`

## Features

## Dependencies

## Architecture Structure and Decisions

## API Endpoints

| route | parameters | example | response |
|-------|------------|---------|----------|
| state | | curl http://localhost:3000/state | {"state":"ok"} |
| remotestate | address | curl http://localhost:3000/remotestate/10.0.10.6:3000 | {"state":"ok"} |

In case of an error the state returns the error.

## Debugging Hints

## Manifest Examples

## Links to External Documentation

## TODOs

## Nice to have

## Known Issues
