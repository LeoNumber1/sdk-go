package models

import "time"

type Region struct {
	RegionID   string `json:"region_id"`
	Name       string `json:"name"`
	Datacenter string `json:"datacenter"`
	Visible    string `json:"visible"`
	UsedFor    string `json:"used_for"`
}

type Sku struct {
	SkuID string `json:"skuID"`
	Name  string `json:"name"`
}

type MarketResponse struct {
	SkuList []Sku    `json:"sku_list"` // 计费方式
	Region  []Region `json:"region"`   // 地域
}

type ListMachinesResponse struct {
	Pagination Pagination    `json:"pagination"`
	Result     []*GpuMachine `json:"result"`
}

type GpuMachine struct {
	MachineId         string                 `json:"machine_id"`
	MachineName       string                 `json:"machine_name"`
	MachineSupply     string                 `json:"machine_supply"`
	GpuName           string                 `json:"gpu_name"`
	GpuNum            int32                  `json:"gpu_num"`
	GpuIdle           int32                  `json:"gpu_idle"`
	GpuUsed           int32                  `json:"gpu_used"`
	GpuMemory         int64                  `json:"gpu_memory" ` // GPU显存大小 单位 GB
	RentStartTime     time.Time              `json:"rent_start_time"`
	RentEndTime       time.Time              `json:"rent_end_time"`
	Status            int32                  `json:"status"`
	RentHours         int32                  `json:"rent_hours"`
	RentTimes         int32                  `json:"rent_times"`
	CpuPerGpu         int32                  `json:"cpu_per_gpu"`
	MemoryPerGpu      int64                  `json:"memory_per_gpu"`
	MaxDiskExpand     int64                  `json:"max_disk_expand"`
	Level             string                 `json:"level"`
	Score             int32                  `json:"score"`
	Nas               int32                  `json:"nas"`
	Tag               string                 `json:"tag"`
	IsOnline          int32                  `json:"is_online"`
	RegionName        string                 `json:"region_name"`
	RegionId          string                 `json:"region_id"`
	GpuLowPrice       string                 `json:"gpu_low_price"`
	GpuHighPrice      string                 `json:"gpu_high_price"`
	BaseInfo          *MarketMachineBaseInfo `json:"base_info"`
	Sku               []SkuInfo              `json:"sku"`
	SupplierID        string                 `json:"supplier_id"`
	MainTenanceStatus int32                  `json:"maintenance_status"`
}

type SkuInfo struct {
	SkuId         string `json:"sku_id"`
	Price         string `json:"price"`
	OriginalPrice string `json:"original_price"`
	Discount      string `json:"discount"`
	Status        int32  `json:"status"`
	Enable        int32  `json:"enable"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
}

type MarketMachineBaseInfo struct {
	Id                int64     `json:"id"` // id
	MachineId         string    `json:"machine_id"`
	MachineType       string    `json:"machine_type"`
	PhysicalMachineID string    `json:"physical_machine_id"`
	CPUCoreNum        int       `json:"cpu_core_num"`
	CPUBaseFrequency  string    `json:"cpu_base_frequency"`
	CPUName           string    `json:"cpu_name"`
	PhysicalCPUNum    int       `json:"physical_cpu_num"`
	DiskSize          int64     `json:"disk_size"`
	DiskType          string    `json:"disk_type"`
	GPUToolkitName    string    `json:"gpu_toolkit_name"`
	GPUToolkitVersion string    `json:"gpu_toolkit_version"`
	GPUDriverVersion  string    `json:"gpu_driver_version"`
	GPUName           string    `json:"gpu_name"`
	GPUNum            int       `json:"gpu_num"`
	MemoryTotal       int64     `json:"memory_total"`
	NetDownloadSpeed  int64     `json:"net_download_speed"`
	NetUploadSpeed    int64     `json:"net_upload_speed"`
	OSName            string    `json:"os_name"`
	Hostname          string    `json:"hostname"`
	RegionID          string    `json:"region_id"`
	VMType            string    `json:"vm_type"`
	CreateTime        time.Time `json:"create_time" `
	UpdateTime        time.Time `json:"update_time"`
}

type ImageItem struct {
	Label    string      `json:"label"`
	Value    string      `json:"value"`
	Children []ImageItem `json:"children"`
	Uuid     *string     `json:"uuid"`
}
