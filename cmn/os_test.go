package cmn

import (
	"testing"
)

func Test_measures(t *testing.T) {
	total, used, free, memPercent := MeasureMemory()
	Info("总内存", total, "，已使用", used, "，空余", free, "，使用占比", Float64ToString(Round(memPercent, 1))+"%")

	total, used, free, swapPercent := MeasureSwap()
	Info("总虚拟内存", total, "，已使用", used, "，空余", free, "，使用占比", Float64ToString(Round(swapPercent, 1))+"%")

	physicalCount, logicalCount, cpuPercent := MeasureCPU()
	Info("物理CPU核数", physicalCount, "，逻辑CPU核数", logicalCount, "，使用占比", Float64ToString(Round(cpuPercent, 1))+"%")

	total, used, free, diskPercent := MeasureDisk()
	Info("磁盘总容量", total, "，已使用", used, "，空余", free, "，使用占比", Float64ToString(Round(diskPercent, 1))+"%")

	Info(MeasureHost())
	Info(MeasureSummary())
}

func Test_measures2(t *testing.T) {
	Info(MeasureDiskFreeSpace("d:/Csa\\s/a/f"))
}
