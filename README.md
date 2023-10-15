# Keyboard-auto-typer

## Summary

The Keyboard Auto Typer project is a Go-based server that utilizes code from the "send-keys" project to programmatically generate and send keyboard events. These events simulate human-like typing behavior, enabling the server to automate the input of text as if a person were physically typing on the keyboard.

## Features

- HTTP Server: The project creates an HTTP server in Go, capable of receiving POST requests.

- Keyboard Simulation: Upon receiving a POST request, the server leverages the "send-keys" code to simulate keyboard events, effectively typing out the specified text.

- Customizable: Users can define the text to be typed and, if needed, customize the typing speed or intervals between keypresses.

- Automation: This tool is ideal for automating repetitive data entry tasks, form filling, or any application where simulated keyboard input is required.

## Getting Started

- Clone the project repository to your local machine.
- Ensure you have Go installed.
- Start the Go server using `go run main.go`.
- Make a POST request to the server with the text you want to simulate typing.
