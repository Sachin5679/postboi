# HTTP Client CLI Tool

This project is a Command Line Interface (CLI) tool written in Go, designed to send HTTP requests with specified methods, headers, data, and authentication. It outputs the response status and body, making it useful for testing HTTP APIs from the command line.

## Features

- Supports `GET`, `POST`, `PUT`, and other HTTP methods.
- Custom headers can be added to requests.
- Sends data payload for `POST` and `PUT` methods.
- Supports `bearer` and `basic` authentication.
- Outputs the HTTP response status and body to the console.

## Input format

go run main.go -method <HTTP_METHOD> -header "<HEADER>" -data "<DATA_PAYLOAD>" -auth-type "<AUTH_TYPE>" -token "<TOKEN>" <URL>

## Command-Line Flags

- `-method`: HTTP method (GET, POST, PUT, etc.)
- `-header`: Custom headers in the format `Key: Value` (comma-separated for multiple headers)
- `-data`: Payload data (required for POST and PUT methods)
- `-auth-type`: Authentication type (`bearer` or `basic`)
- `-token`: Token for `bearer` or `username:password` for `basic` authentication



