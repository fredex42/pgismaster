# pgismaster

## What is it?

pgismaster is a very simple little program designed to indicate to a process like keepalived whether
a given postgres server is a master or a standby.

This is done by checking the internal flag "pg_is_in_recovery();"; an active slave is permanently in a recovery
state and a master is not.

- If run on a slave, then a message is printed to indicate that it is a slave and an exit status of 1 is returned.
- If run on a master, then a message is printed to indicate that it is a master and an exit status of 1 is returned.

This allows it to be run in the "script" parameter of keepalived to allow only a master to assert a VIP.

## How do I run it?

Simply grab the compiled executable and deploy it to your server.  Then, call it in your script or config file.

## How do I compile it?

You'll need Go 1.11 or later available to compile (but not to run).
You can either head over to https://golang.org/doc/install#install or use docker (e.g. `docker run --rm golang:1.12-alpine3.9`)

Check out the sources: `git clone https://github.com/fredex42/pgismaster`

Compile it for Linux: `GOOS=linux go build`

Or for Windows: `GOOS=windows go build`

Or for Mac: `GOOS=darwin go build`

etc.

Dependencies are managed via go modules, so any required libraries should be automatically downloaded and compiled in the `build` command.

See the official Go documentation for cross-compilation options.