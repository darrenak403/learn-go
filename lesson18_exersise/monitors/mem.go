package monitors

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

type MemMonitor struct {
}

func (m *MemMonitor) Name() string {
	return "Memory Monitor"
}

func (m *MemMonitor) Check(ctx context.Context) string {
	vmStat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		panic(err)
	}

	value := fmt.Sprintf("%.2f%%", vmStat.UsedPercent)
	return value
}
