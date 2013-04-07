locald
======

### Overview

Ever need to run an HTTP server to serve static content out of a directory for testing?
Well, me too!
That's what locald does.
In fact, that's all it does.
Even better is that it doesn't have any dependencies outside of Go.
Once you build the server you end up with a nice single, statically compiled binary.


### Getting Started

To get started, make sure you install [Go](http://golang.org/) and then build the server:

```sh
$ go build
```

Now you have a `locald` executable in your current directory! Woohoo!
Run it like so:

```sh
$ ./locald
locald v0.1.0
Serving /your/current/path on http://localhost:9000
```

The server is dead simple so there's not a lot of options.
Actually... there's just one.
You can specify the port using the `-p PORT` or `--port PORT` command line arguments.
By default it'll use port 9000.


### Contributions

Unless you can think of a way to make this server any simpler then you probably shouldn't bother submitting a pull request.
It's meant to be simple.
If you need more features then fork the code or use ~~Apache~~ Nginx.
