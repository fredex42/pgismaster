# pgismaster

## What is it?

pgismaster is a very simple little program designed to indicate to a process like keepalived whether
a given postgres server is a master or a standby.

This is done by checking the internal flag "pg_is_in_recovery();"; an active slave is permanently in a recovery
state and a master is not.

- If run on a slave, then a message is printed to indicate that it is a slave and an exit status of 1 is returned.
- If run on a master, then a message is printed to indicate that it is a master and an exit status of 0 is returned.

This allows it to be run in the "script" parameter of keepalived to allow only a master to assert a VIP.

## How do I run it?

Simply grab the compiled executable and copy it to your server.  Then, call it in your script or config file.

For example, in `keepalived.conf`:

```
global_defs {
	enable_script_security
	script_user postgres    ###server is configured to allow local access here
}

vrrp_script chk_postgres {               # Requires keepalived-1.1.13
        script "/usr/local/bin/pgismaster"
        interval 10                      # check every 10 seconds
        fall 2       # require 2 failures for KO
        rise 2       # require 2 successes for OK
}

vrrp_instance VI_1 {
        interface ens192
        virtual_router_id 42
        priority 150
        state MASTER
        virtual_ipaddress {
            1.2.3.4
        }
        track_script {
            chk_postgres
        }
}
```

Or just run from the commandline:
```
$ /usr/local/bin/pgismaster
2019/10/14 16:45:31 pgismaster, Andy Gallagher 2019. See https://github.com/fredex42/pgismaster for details.
2019/10/14 16:45:31 Server is a standby
$ echo $?
1
```

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
