# xclient

A go client with web interface to test X-Road services

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/xclient run
```

Then open the following URL in your browser: http://127.0.0.1:5000/

This program will access files in the ui directory, so if you want to copy the program outside the source tree, please include the ui directory as well.

## Debugging

You can run code directly with go like this:
```console
$ go run main.go run
```

### Testing

``make test``