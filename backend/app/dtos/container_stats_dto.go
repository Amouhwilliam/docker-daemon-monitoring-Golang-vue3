package dtos

type Stats struct {
	TotalCpus       int     `json:"totalCpus"`
	CpuUsage        float64 `json:"cpuUsage"`
	UsedMemory      uint64  `json:"usedMemory"`
	UsedMemoryMb    float64 `json:"usedMemoryMb"`
	TotalMemory     uint64  `json:"totalMemory"`
	TotalMemoryGb   float64 `json:"totalMemoryGb"`
	MemoryUsage     float64 `json:"memoryUsage"`
	NetworkInput    uint64  `json:"networkInput"`    // Bytes received
	NetworkOutput   uint64  `json:"networkOutput"`   // Bytes transmitted
	NetworkInputMb  float64 `json:"networkInputMb"`  // Bytes received
	NetworkOutputMb float64 `json:"networkOutputMb"` // Bytes transmitted
}
