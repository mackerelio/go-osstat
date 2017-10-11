// +build darwin

package memory

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

// Get memory statistics
func Get() (memory Memory, err error) {
	memory, err = collectMemoryStats(newMemoryGenerator())
	if err != nil {
		return
	}
	return
}

// Memory represents memory statistics for darwin
type Memory struct {
	Total    uint64
	Used     uint64
	Cached   uint64
	Free     uint64
	Active   uint64
	Inactive uint64
}

type memoryGenerator interface {
	Name() string
	Start() error
	Output() io.Reader
	Finish() error
}

type memoryGeneratorImpl struct {
	cmd *exec.Cmd
}

func newMemoryGenerator() *memoryGeneratorImpl {
	return &memoryGeneratorImpl{cmd: exec.Command("vm_stat")}
}

func (generator memoryGeneratorImpl) Start() error {
	return generator.cmd.Start()
}

func (generator memoryGeneratorImpl) Output() (io.Reader, error) {
	stdout, err := generator.cmd.StdoutPipe()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	return stdout
}

func (generator memoryGeneratorImpl) Finish() error {
	return generator.cmd.Wait()
}

const (
	freePages        = "Pages free"
	activePages      = "Pages active"
	inactivePages    = "Pages inactive"
	speculativePages = "Pages speculative"
	wiredDownPages   = "Pages wired down"
	purgeablePages   = "Pages purgeable"
	fileBackedPages  = "File-backed pages"
	compressedPages  = "Pages occupied by compressor"
)

// References:
//   - https://support.apple.com/en-us/HT201464#memory
//   - https://developer.apple.com/library/content/documentation/Performance/Conceptual/ManagingMemory/Articles/AboutMemory.html
//   - https://opensource.apple.com/source/system_cmds/system_cmds-790/vm_stat.tproj/
func collectMemoryStats(generator memoryGenerator) (Memory, error) {
	scanner := bufio.NewScanner(generator.Output())
	if err := generator.Start(); err != nil {
		return Memory{}, err
	}

	if !scanner.Scan() { // skip the first line
		return Memory{}, fmt.Errorf("failed to scan output of %s", generator.Name())
	}

	stats := make(map[string]uint64, 22)
	pageSize := uint64(4096)
	for scanner.Scan() {
		line := scanner.Text()
		i := strings.IndexRune(line, ':')
		if i < 0 {
			continue
		}
		val := strings.TrimRight(strings.TrimSpace(line[i+1:]), ".")
		if v, err := strconv.ParseUint(val, 10, 64); err == nil {
			stats[line[:i]] = v * pageSize
		}
	}

	wired := stats[wiredDownPages]
	compressed := stats[compressedPages]
	cached := stats[purgeablePages] + stats[fileBackedPages]
	active := stats[activePages]
	inactive := stats[inactivePages]
	used := wired + compressed + active + inactive + stats[speculativePages] - cached
	free := stats[freePages]

	if err := generator.Finish(); err != nil {
		return Memory{}, err
	}

	return Memory{
		Total:    used + cached + free,
		Used:     used,
		Cached:   cached,
		Free:     free,
		Active:   active,
		Inactive: inactive,
	}, nil
}
