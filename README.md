# Keyboard-auto-typer

## Summary

The Keyboard Auto Typer project is a Go-based server that utilizes code from the "send-keys" project to programmatically generate and send keyboard events. These events simulate human-like typing behavior, enabling the server to automate the input of text as if a person were physically typing on the keyboard.

## Features

- HTTP Server: The project creates an HTTP server in Go, capable of receiving POST requests.

- Keyboard Simulation: Upon receiving a POST request, the server leverages the [send-keys](https://github.com/yunginnanet/sendkeys) code to simulate keyboard events, effectively typing out the specified text.

- Customizable: Users can define the text to be typed and, if needed, customize the typing speed or intervals between keypresses.

- Automation: This tool is ideal for automating repetitive data entry tasks, form filling, or any application where simulated keyboard input is required.

## Getting Started

- Clone the project repository to your local machine.
- Ensure you have Go installed.
- Start the Go server using `go run main.go`.
- Make a POST request to the server with the text you want to simulate typing.

## Usage

To use the Keyboard Auto Typer, follow these steps:

1. Start the server using go run main.go.

1. Send a POST request to the server's endpoint at http://localhost:51212/send. The request body should contain a JSON schema with the following fields:

    - text (array of strings): An array of strings that represents the text you want to simulate typing.

    - startAfter (time in milliseconds): An optional delay, in milliseconds, before typing begins. This allows you to wait for a specific event or condition before initiating the typing simulation.

    - delay (time in milliseconds): The time interval, in milliseconds, between keypresses. This controls the typing speed.

Here's an example of a JSON schema:

```json
{
    "text": ["Hello, World!", "This is a test."],
    "startAfter": 3000,
    "delay": 10
}
```

Using curl in the command line:
```shell
curl -X POST -d '{
    "text": ["Hello, World!", "This is a test."],
    "startAfter": 3000,
    "delay": 10
}' http://localhost:51212/send
```
The server will receive the POST request, parse the JSON schema, and simulate the keyboard typing based on the provided parameters.
