package models

type InstanceInfo struct {
	InstanceID          string  `json:"instance_id"`
	MachineID           string  `json:"machine_id"`
	MachineName         string  `json:"machine_name"`
	RegionID            string  `json:"region_id"`
	RegionName          string  `json:"region_name"`
	DriverVersion       string  `json:"driver_version"`
	Status              string  `json:"status"`
	SSHProxyHost        string  `json:"ssh_proxy_host"`
	SSHCommand          string  `json:"ssh_command"`
	SSHRootPassword     string  `json:"ssh_root_password"`
	JupyterDomain       string  `json:"jupyter_domain"`
	JupyterToken        string  `json:"jupyter_token"`
	JupyterPort         int     `json:"jupyter_port"`
	TensorboardDomain   string  `json:"tensorboard_domain"`
	TensorboardPort     int     `json:"tensorboard_port"`
	CustomServiceStatus string  `json:"custom_service_status"`
	CustomUrl           string  `json:"custom_url"`
	ImageID             string  `json:"image_id"`
	SkuID               string  `json:"sku_id"`
	ReqGpuAmount        int     `json:"req_gpu_amount"`
	InstanceGPUNum      int     `json:"instance_gpu_num"`
	GPUIdleNum          int     `json:"gpu_idle_num"`
	StartedAt           string  `json:"started_at"`
	StoppedAt           string  `json:"stopped_at"`
	ScheduleStopTime    string  `json:"schedule_stop_time"`
	ScheduleReleaseTime string  `json:"schedule_release_time"`
	DiskUsedRate        float64 `json:"disk_used_rate"`
	TmpDiskUsedRate     float64 `json:"tmp_disk_used_rate"`
	AutoRenew           string  `json:"auto_renew"`
}

type ListInstancesResponse struct {
	Pagination Pagination      `json:"pagination"`
	Result     []*InstanceInfo `json:"result"`
}

type CreateInstanceRequest struct {
	MachineID      string `json:"machine_id"`       // 机器ID
	ReqGPUAmount   int32  `json:"req_gpu_amount"`   // 请求GPU数量
	ImageType      string `json:"image_type"`       // 镜像类型 官方镜像 用户自定义镜像
	Image          string `json:"image"`            // 镜像ID (官方镜像)
	PrivateImage   string `json:"private_image"`    // 用户自定义镜像ID
	InstanceName   string `json:"instance_name"`    // 实例名称
	ExpandDataDisk int64  `json:"expand_data_disk"` // 扩容数据盘大小
	SkuID          string `json:"sku_id"`           // 计费类型 按量计费 包年包月
	Duration       int64  `json:"duration"`         // 时长
	AutoRenew      string `json:"auto_renew"`       // auto 自动续费  manual 手动续费
}
