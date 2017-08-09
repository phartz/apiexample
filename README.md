# api example

1. [Description](#description)
2. [Install, Run, Usage](#install-run-usage)
3. [Features](#features)
4. [Dependencies](#dependencies)
5. [Architecture Structure and Decisions](#architecture-structure-and-decisions)
6. [API Endpoints](#api-endpoints)

## Desctiption
`apiexample` is a simple web server of how to implement a lightweight API written in golang.
It is designed to either give back it's own state (always ok) or the state of a similar application on another server.

## Install, Run, Usage
Clone the repo:
`go get github.com/phartz/apiexample`

To compile simply run `make`.

## Features
`./apiexmaple -v` shows the version information

## Dependencies

## Architecture Structure and Decisions

`curl http://<ADDRESS_APP1>/remotestate/<ADDRESS_APP2>` 

requests APP1 to ask APP2 for its  status
```
                       *------*          *------*
  /remotestate/<APP2>  | APP1 |  /state  | APP2 |
 --------------------> |      | -------> |      |
                       |      |   {OK}   |      |
                       *------*          *------*
```

## API Endpoints

| route | parameters | example | response |
|-------|------------|---------|----------|
| state | | curl http://localhost:3000/state | {"state":"ok"} |
| remotestate | address | curl http://localhost:3000/remotestate/10.0.10.6:3000 | {"state":"ok"} |

In case of an error `state` returns the error.

