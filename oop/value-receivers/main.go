package main

import "fmt"

type usage int

func (u usage) high() bool { return u >= 95 }
func (u usage) set(to int) { u = usage(to) } //nolint:staticcheck,ineffassign

func main() {
	var cpu usage = 99 // cpu is 99
	cpu.set(70)        // cpu is still 99
	fmt.Printf("cpu: %d%% high:%t\n", cpu, cpu.high())
}
