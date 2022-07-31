//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) AverageBandwidthUsage() int {
	total := 0
	for i := 0; i < len(b.amount); i++ {
		total += int(b.amount[i])
	}

	return total / len(b.amount)
}

type CpuTemp struct {
	temp []Celcius
}

func (c *CpuTemp) AverageCpuTemp() float32 {
	total := 0
	for i := 0; i < len(c.temp); i++ {
		total += int(c.temp[i])
	}

	return float32(total / len(c.temp))
}

type MemoryUsage struct {
	amount []Bytes
}

func (m *MemoryUsage) AverageMemoryUsage() int {
	total := 0
	for i := 0; i < len(m.amount); i++ {
		total += int(m.amount[i])
	}

	return total / len(m.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func (d *Dashboard) promoted() {
	fmt.Printf("Average bandwidth usage: %v\n", d.AverageBandwidthUsage())
	fmt.Printf("Average cputemp: %v\n", d.AverageCpuTemp())
	fmt.Printf("Average memory usage: %v\n", d.AverageMemoryUsage())
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboard := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}
	dashboard.promoted()
}
