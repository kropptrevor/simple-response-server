# Simple Response Server
The purpose of this project is to provide simple scripts which can be used to provide a fixed response (or lack of) by running them as a script.
The scripts will start up an HTTP server and respond in the way specified in the script.

This project is only intended to be run locally for testing.
It is not indented to be run on a centralized server anywhere.

## Requirements
- Go 1.21.0 or higher

## Running

Execute the command 
```
go run ./cmd/200-with-body/
```
to run the script `cmd/200-with-body/main.go`.

You can test that it is running by executing
```
curl http://localhost:3333
```

## Port
By default, it uses 3333 throughout. If you want to change it, just change it in the script.

## Scripts

| Name              | Description                                                                            | Response Status |
| ----------------- | -------------------------------------------------------------------------------------- | --------------- |
| 200-with-body     | This returns 200 OK with a fixed body                                                  | 200 OK          |
| connection-failed | This doesn't return and merely fails, which can invoke errors on upstream applications | N/A             | 