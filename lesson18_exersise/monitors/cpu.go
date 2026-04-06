package monitors

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type CPUMonitor struct {
}

func (m *CPUMonitor) Name() string {
	return "CPU Monitor"
}

func (m *CPUMonitor) Check(ctx context.Context) string {
	percent, err := cpu.PercentWithContext(ctx, 1*time.Second, false)
	if err != nil && len(percent) == 0 {
		panic(err)
	}

	value := fmt.Sprintf("%.2f%%", percent[0])
	return value
}
