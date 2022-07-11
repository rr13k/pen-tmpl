package machineservers

import (
	"testing"
	"time"
)

// 测试申请机器
func Test_Apple(t *testing.T) {
	jdosCloudSdk := newJdosCloudSdk(&MachineConfig{
		Cpu:  8,
		Mem:  32,
		Disk: 200,
	}, nil)
	jdosCloudSdk.getAppGroupConfig()
}

func Test_Apple_JDos(t *testing.T) {
	Apple(&MachineConfig{
		Cpu:  8,
		Mem:  32,
		Disk: 50,
	}, "test")

	time.Sleep(100000 * time.Second) // 注意删除
}

func Test_Apple_JdCloud(t *testing.T) {
	Apple(&MachineConfig{
		Cpu:  12,
		Mem:  32,
		Disk: 50,
	}, "test")

	time.Sleep(100000 * time.Second) // 注意删除
}
