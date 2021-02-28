# go-web-crawler-starter

This project is a working example of a web crawler with multiple controllers.

As this project is an example, I have chosen to retrieve some statistics from the following PUBG (video game) statistics
web pages;

- https://pubg.op.gg/
- https://pubglookup.com/

Given a PUBG username, controllers retrieves some data from the pages.

## Controllers

For each page listed above, there is a similarly named controller implemented in the source code. These controllers
reside in;
 
`src/controller/<DOMAIN_NAME>/controller.go` (example <DOMAIN_NAME>="pubg.op.gg")

Each controller implements the following `interface`;

```
type Controller interface {
	GetName() string
	Run(controllerConfiguration *configuration.ControllerConfiguration)
}
```

Controllers gets registered at `src/cmd/main.go:registerControllers()`.

## Build

Build `src/cmd/main.go`

## Usage

Use `--help` flag to print flags

```
  -alsologtostderr
        log to standard error as well as files
  -controllers string
        comma separated controllers list
  -log_backtrace_at value
        when logging hits line file:N, emit a stack trace
  -log_dir string
        If non-empty, write log files in this directory
  -logtostderr
        log to standard error instead of files
  -pubg-username string
        PUBG username
  -stderrthreshold value
        logs at or above this threshold go to stderr
  -v value
        log level for V logs
  -vmodule value
        comma-separated list of pattern=N settings for file-filtered logging
```

Required to run;

`-controllers="pubg.op.gg,pubglookup.com" -pubg-username="<PUBG-USERNAME>"`

Add `-logtostderr=true` to see logs on standard output.

