# OS system statistics library for Go
This is a library to get system metrics like cpu load and memory usage.
The library is created for [mackerel-agent](https://github.com/mackerelio/mackerel-agent).

[![GoDev](https://pkg.go.dev/badge/github.com/mackerelio/go-osstat)](https://pkg.go.dev/github.com/mackerelio/go-osstat)
[![Build Status](https://github.com/mackerelio/go-osstat/workflows/Build/badge.svg)](https://github.com/mackerelio/go-osstat/workflows/Build/badge.svg)
[![Coverage Status](https://coveralls.io/repos/mackerelio/go-osstat/badge.svg)](https://coveralls.io/r/mackerelio/go-osstat)

## Example
```go
package main

import (
	"fmt"
	"os"

	"github.com/mackerelio/go-osstat/memory"
)

func main() {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	fmt.Printf("memory total: %d bytes\n", memory.Total)
	fmt.Printf("memory used: %d bytes\n", memory.Used)
	fmt.Printf("memory cached: %d bytes\n", memory.Cached)
	fmt.Printf("memory free: %d bytes\n", memory.Free)
}
```

## Supported OS

||loadavg|uptime|cpu|memory|network|disk i/o|
|:--:|:--:|:--:|:--:|:--:|:--:|:--:|
|Linux|yes|yes|yes|yes|yes|yes|
|Darwin|yes|yes|*1|yes|yes|no|
|FreeBSD|yes|yes|no|yes|yes|no|
|NetBSD|yes|yes|no|no|yes|no|
|OpenBSD|yes|yes|no|no|no|no|
|Windows|no|yes|no|yes|no|no|

*1: unavailable without cgo

## Note for counter values
This library returns the counter value for cpu, network and disk I/O statistics by design. To get the cpu usage in percent, network traffic in kB/s or disk IOPS, sleep for a while and calculate the difference.

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
)

func main() {
	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	total := float64(after.Total - before.Total)
	fmt.Printf("cpu user: %f %%\n", float64(after.User-before.User)/total*100)
	fmt.Printf("cpu system: %f %%\n", float64(after.System-before.System)/total*100)
	fmt.Printf("cpu idle: %f %%\n", float64(after.Idle-before.Idle)/total*100)
}
```

## LICENSE
```
Copyright 2017-2019 Hatena Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
