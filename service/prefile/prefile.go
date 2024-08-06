package prefile

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

// Info 生成Bot环境中的CPU，内存等信息
func Info() string {
	hostInfo, _ := host.Info()
	cpuInfo, _ := cpu.Percent(time.Second, false)
	memInfo, _ := mem.VirtualMemory()
	return fmt.Sprintf(
		"\nBot-486 运行环境:\n操作系统: %s\n系统架构: %s\nCPU利用率: %.2f%%\n内存利用率: %.2f%%",
		hostInfo.OS, hostInfo.KernelArch, cpuInfo[0], memInfo.UsedPercent)
}
