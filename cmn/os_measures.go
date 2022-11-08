package cmn

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

// 检测内存
func MeasureMemory() (total uint64, used uint64, free uint64, usePercent float64) {
	v, _ := mem.VirtualMemory()
	total = v.Total
	used = v.Total - v.Free
	free = v.Free
	usePercent = v.UsedPercent
	return
}

// 检测CPU
func MeasureCPU() (physicalCount int, logicalCount int, usePercent float64) {
	physicalCount, _ = cpu.Counts(false)       // 物理核数
	logicalCount, _ = cpu.Counts(true)         // 逻辑核数
	pers, _ := cpu.Percent(time.Second, false) // cpu总体使用率
	usePercent = pers[0]
	return
}

// 检测磁盘(当前盘)
func MeasureDisk0() (total uint64, used uint64, free uint64, usePercent float64) {
	parts, _ := disk.Partitions(false)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	total = diskInfo.Total
	used = diskInfo.Used
	free = diskInfo.Free
	usePercent = diskInfo.UsedPercent
	return
}
