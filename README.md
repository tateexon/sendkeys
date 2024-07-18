# SendKeys

`sendkeys` is a Go program that simulates typing a string with random delays between each key press. The program can read the text to type either from a command-line argument or from a file.

## Features

- Simulates typing a string with random delays between keystrokes.
- Supports upper and lower case letters, numbers, and common punctuation.
- Reads input text from a command-line argument or from a file.
- Customizable start delay and inter-key delay with variance.

## Usage

### Command-Line Arguments

- `-t int`: Number of seconds to wait before starting (default: 5 seconds).
- `-i int`: Interval in milliseconds between each key press (default: 75 milliseconds).
- `-v int`: Variance of interval in milliseconds (default: 40 milliseconds).
- `-f string`: Path to a file with text you wish to be typed.
- `-s string`: String to type.

### Example

```sh
go run main.go -t 5 -i 75 -v 40 -s "Hello, world!"
```

This command will:

- Wait for 5 seconds.
- Type the string "Hello, world!" with an interval of 75 milliseconds between each key press, plus a variance of 40 milliseconds.
Alternatively, you can specify a file to read the text from:

```sh
go run main.go -t 5 -i 75 -v 40 -f example.txt
```

This command will:

- Wait for 5 seconds.
- Read the contents of example.txt and type it with the specified interval and variance.

## Installation
Ensure you have Go installed on your system. You can download it from golang.org. Then to install run:

```sh
go install github.com/tateexon/sendkeys@latest
```
