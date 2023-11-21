package tools

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

const (
	keyProcessId       = "ProcessId"
	keyParentProcessId = "ParentProcessId"
	keyName            = "Name"
)

func getQueryHeaders() string {
	return fmt.Sprintf("%s,%s,%s", keyProcessId, keyParentProcessId, keyName)
}

type process struct {
	name    string
	pid     string
	ppid    string
	printed bool

	children []*process
}

func (p *process) toString() string {
	return fmt.Sprintf("%s %s %s", p.name, p.pid, p.ppid)
}

func printProcess(p *process, indentation int) {
	if p.printed {
		return
	}

	indent := " "
	if indentation > 0 {
		for i := 0; i < indentation; i++ {
			indent += "\t"
		}
	}
	fmt.Printf("%s %s\n", indent, p.toString())
	p.printed = true

	for _, pc := range p.children {
		printProcess(pc, indentation+1)
	}
}

func getIndexFromHeader(lines []string) (nameindex, pidindex, ppidindex int, err error) {
	for _, l := range lines {
		if !strings.Contains(l, keyParentProcessId) {
			continue
		}
		filelds := strings.Split(strings.Trim(l, "\r\n "), ",")
		for i, v := range filelds {
			if v == keyProcessId {
				pidindex = i
			} else if v == keyParentProcessId {
				ppidindex = i
			} else if v == keyName {
				nameindex = i
			}
		}
		return
	}

	err = fmt.Errorf("not found headers")
	return
}

func Pstree() error {
	if runtime.GOOS != "windows" {
		return fmt.Errorf("only implement on windows now, you can install pstree on others")
	}

	out, err := exec.Command("wmic", "process", "get", getQueryHeaders(), "/format:csv").Output()
	if err != nil {
		fmt.Println("执行命令出错:", err)
		return nil
	}

	lines := strings.Split(string(out), "\n")
	// get right result index for each key
	nindex, pindex, ppindex, err := getIndexFromHeader(lines)
	if err != nil {
		return err
	}

	processarr := make([]*process, 0, 100)
	for _, l := range lines {
		if strings.Contains(l, keyParentProcessId) {
			continue
		}
		filelds := strings.Split(strings.Trim(l, "\r\n "), ",")
		if len(filelds) == 4 {
			processarr = append(processarr, &process{
				name: filelds[nindex],
				ppid: filelds[ppindex],
				pid:  filelds[pindex],
			})
		}
	}

	for _, p := range processarr {
		for _, p1 := range processarr {
			if p1.ppid == p.pid && p.pid != p.ppid {
				if p.children == nil {
					p.children = make([]*process, 0, 10)
				}

				p.children = append(p.children, p1)
			}
		}
	}

	for _, p := range processarr {
		printProcess(p, 0)
	}

	return nil
}
