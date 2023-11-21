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
	keyWorkingSetSize  = "WorkingSetSize"
	keyCreationDate    = "CreationDate"

	keyNum = 5 + 1 // query keys and default node
)

func getQueryHeaders() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s", keyProcessId, keyParentProcessId, keyName, keyWorkingSetSize, keyCreationDate)
}

type process struct {
	name           string
	pid            string
	ppid           string
	workingSetSize string
	creationDate   string

	printed bool

	children []*process
}

func (p *process) toString() string {
	return fmt.Sprintf("%s %s %s %s %s ", p.name, p.pid, p.ppid, p.workingSetSize, p.creationDate)
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

func getIndexFromHeader(lines []string) (nameindex, pidindex, ppidindex, wsindex, cdindex int, err error) {
	for _, l := range lines {
		if !strings.Contains(l, keyParentProcessId) {
			continue
		}
		filelds := strings.Split(strings.Trim(l, "\r\n "), ",")
		for i, v := range filelds {
			switch v {
			case keyProcessId:
				pidindex = i
				break
			case keyParentProcessId:
				ppidindex = i
				break
			case keyName:
				nameindex = i
				break
			case keyWorkingSetSize:
				wsindex = i
				break
			case keyCreationDate:
				cdindex = i
				break
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
	// fmt.Printf("%s\r\n", string(out))

	lines := strings.Split(string(out), "\n")
	// get right result index for each key
	nindex, pindex, ppindex, wsindex, cdindex, err := getIndexFromHeader(lines)
	if err != nil {
		return err
	}

	processarr := make([]*process, 0, 100)
	for _, l := range lines {
		if strings.Contains(l, keyParentProcessId) {
			continue
		}
		filelds := strings.Split(strings.Trim(l, "\r\n "), ",")
		if len(filelds) >= keyNum {
			processarr = append(processarr, &process{
				name:           filelds[nindex],
				ppid:           filelds[ppindex],
				pid:            filelds[pindex],
				workingSetSize: filelds[wsindex],
				creationDate:   filelds[cdindex],
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
