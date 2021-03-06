# 📝 logit

[![License](_icon/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Go Doc](_icon/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/logit?tab=doc)

**logit** is an easy-to-use, also level-based and config file first logger for [GoLang](https://golang.org) applications.

[阅读中文版的 Read me](./README.md).

### 🥇 Features

* Modularization design, easy to extend your logger with wrapper and handler
* Level-based logging, and there are four levels to use
* Config file supports, you can use a config file to change you logger flexibility even it has been a binary
* Log Function supports, it is a better way to output a very long log
* Enable or disable Logger, you can disable or switch to a higher level in your production environment
* Log file supports, and you can customer the name of your log file
* Duration rolling supports, which means it will roll to a new log file by duration automatically, such as one day one log file
* File size rolling supports, which means it will roll to a new log file by file size automatically, such as one 64 MB one log file
* Log handler supports, you can extend logger with your own log handler easily
* High-performance supports, by avoiding to call runtime.Caller
* Time format supports, you can format time in your way
* Log as Json string supports, by using provided JsonLoggerHandler

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

> v0.1.x is an interim version and will not be updated but fixed from now on. Next version v0.2.x is a big update which will bring new features and user experience, also it will be updated and maintained for a long time!

### 🚀 Installation

The only requirement is the [Golang Programming Language](https://golang.org).

> Go modules

```bash
$ go get -u github.com/FishGoddess/logit
```

Or edit your project's go.mod file and execute _**go build**_.

```bash
module your_project_name

go 1.14

require (
    github.com/FishGoddess/logit v0.2.3
)
```

> Go path

```bash
$ go get -u github.com/FishGoddess/logit
```

logit has no more external dependencies.

```go
package main

import (
    "math/rand"
    "strconv"
    "time"
    
    "github.com/FishGoddess/logit"
)

func main() {
    
	// Log messages with four levels.
	logit.Debug("I am a debug message!")
	logit.Info("I am an info message!")
	logit.Warn("I am a warn message!")
	logit.Error("I am an error message!")

	// Notice that logit has blocked some methods for more refreshing method list.
	// If you want to use some higher level methods, you should call logit.Me() to
	// get the fully functional logger, then call what you want to call.
	// For example, if you want to output log with file info, try this:
	logit.Me().EnableFileInfo()
	logit.Info("Show file info!")

	// If you have a long log and it is made of many variables, try this:
	// The msg is the return value of msgGenerator.
	logit.DebugFunc(func() string {
		// Use time as the source of random number generator.
		r := rand.New(rand.NewSource(time.Now().Unix()))
		return "debug rand int: " + strconv.Itoa(r.Intn(100))
	})

	// If a config file "logit.conf" in "./", then logit will load it automatically.
	// This is more convenience to use config file and logger.
}
```

### 📖 Examples

* [basic](./_examples/basic.go)
* [logger](./_examples/logger.go)
* [level_and_disable](./_examples/level_and_disable.go)
* [config_file](./_examples/config_file.go)
* [handler](./_examples/handler.go)
* [writer](./_examples/writer.go)
* [log_to_file](./_examples/log_to_file.go)

_Check more examples in [_examples](./_examples)._

_Learn more about config file in [_examples/config](./_examples/config)._

### 🔥 Benchmarks

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=10s
```

> Benchmark file：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

| test case | times ran (large is better) |  ns/op (small is better) | features | extension |
| -----------|--------|-------------|-------------|-------------|
| **logit** | **6429907** | **1855 ns/op** | powerful | high |
| golog | 3361483 | 3589 ns/op | easy | common |
| zap | 2971119 | 4066 ns/op | complex | normal |
| logrus | 1553419 | 7869 ns/op | normal | normal |

> Environment：I7-6700HQ CPU @ 2.6 GHZ, 16 GB RAM

**Notice:**

**1. Fetching file info will call runtime.Caller, which is expensive.**
**However, we think file info is useful in check errors,**
**so we keep this feature, and provide a switch to turn off it for high-performance.**

**2. v0.0.7 and lower versions use some functions of fmt, and these functions is expensive**
**because of reflect (for judging the parameter v interface{}). Actually, these judgements**
**are redundant in a logger. The more effective output is used in v0.0.8 and higher versions.**

**3. After checking the benchmarks of v0.0.8 version, we found that time format takes a lots of time**
**because of time.Time.AppendFormat. In v0.0.11 and higher versions, we use time cache mechanism to**
**reduce the times of time format. However, is it worth to replace time format operation with concurrent competition?**
**The answer is no, so we cancel this mechanism in v0.1.1-alpha and higher versions.**

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects using logit

| Project | Author | Description |
| -----------|--------|-------------|
|  |  |  |

