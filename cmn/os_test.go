package cmn

import (
	"fmt"
	"testing"
)

func Test_os(t *testing.T) {
	Info(GetLocalIp())
}

func Test_os_mem(t *testing.T) {
	total, used, free, memPercent := MeasureMemory()
	Info("总内存", total, "，已使用", used, "，空余", free, "，使用占比", fmt.Sprintf("%.1f%%", memPercent))

	physicalCount, logicalCount, cpuPercent := MeasureCPU()
	Info("物理CPU核数", physicalCount, "，逻辑CPU核数", logicalCount, "，使用占比", fmt.Sprintf("%.1f%%", cpuPercent))

	total, used, free, diskPercent := MeasureDisk0()
	Info("磁盘总容量", total, "，已使用", used, "，空余", free, "，使用占比", fmt.Sprintf("%.1f%%", diskPercent))

}
