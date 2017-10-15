# OS system statistics library for Go
This is a library to get system metrics like cpu load and memory usage.
The library is created for [mackerel-agent](https://github.com/mackerelio/mackerel-agent).

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
		fmt.Fprintf(os.Stdout, "%s\n", err)
		return
	}
	fmt.Printf("memory total: %d bytes\n", memory.Total)
	fmt.Printf("memory used: %d bytes\n", memory.Used)
	fmt.Printf("memory cached: %d bytes\n", memory.Cached)
	fmt.Printf("memory free: %d bytes\n", memory.Free)
}
```

## Supported OS

||loadavg|cpu|memory|network|disk|
|:--:|:--:|:--:|:--:|:--:|:--:|
|Linux|o|o|o|o|x|
|Darwin|o|△`*`|o|o|x|
|Windows|x|x|o|x|x|

`*`: unavailable without cgo

## Note for counter values
This library returns the counter value for cpu and network statistics by design. To get the cpu usage in percent or network traffic in kB/s, sleep for a while and calculate the difference.

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
		fmt.Fprintf(os.Stdout, "%s\n", err)
		return
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		return
	}
	total := float64(after.Total - before.Total)
	fmt.Printf("cpu user: %f %%\n", float64(after.User-before.User)/total*100)
	fmt.Printf("cpu system: %f %%\n", float64(after.System-before.System)/total*100)
	fmt.Printf("cpu idle: %f %%\n", float64(after.Idle-before.Idle)/total*100)
}
```

## LICENSE
Copyright 2017 Hatena Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
