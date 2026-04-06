package processor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"darrenak.com/learn-go-basics/monitors"
)

func RunMonitor(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cpu := monitors.CPUMonitor{}
			cpuPercent := cpu.Check(ctx)
			fmt.Printf("%s: %s\n", cpu.Name(), cpuPercent)

			mem := monitors.MemMonitor{}
			memUsage := mem.Check(ctx)
			fmt.Printf("%s: %s\n", mem.Name(), memUsage)

		case <-ctx.Done():
			return
		}
	}
}
