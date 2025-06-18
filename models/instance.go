package models

type InstanceInfo struct {
	InstanceID          string `json:"instance_id"`
	MachineID           string `json:"machine_id"`
	MachineAlias        string `json:"machine_alias"`
	RegionID            string `json:"region_id"`
	RegionName          string `json:"region_name"`
	RentDeadline        string `json:"rent_deadline"`
	HasAutoPanel        bool   `json:"has_auto_panel"`
	DiskExpandAvailable int64  `json:"disk_expand_available"`
	DriverVersion       string `json:"driver_version"`
	HighestCudaVersion  string `json:"highest_cuda_version"`
	Status              string `json:"status"`
	SubStatus           string `json:"sub_status"`
	StatusAt            string `json:"status_at"`
	OOMKilled           bool   `json:"oom_killed"`

	SSHCommand           string `json:"ssh_command"`
	RootPassword         string `json:"root_password"`
	ProxyHost            string `json:"proxy_host"`
	ProxyHostIP          string `json:"proxy_host_ip"`
	SnapshotGPUAliasName string `json:"snapshot_gpu_alias_name"`

	JupyterToken      string `json:"jupyter_token"`
	JupyterPort       int64  `json:"jupyter_port"`
	TensorboardPort   int64  `json:"tensorboard_port"`
	JupyterDomain     string `json:"jupyter_domain"`
	TensorboardDomain string `json:"tensorboard_domain"`

	CustomServiceStatus string `json:"custom_service"`
	CustomUrl           string `json:"custom_url"`

	ScheduleStopTime string `json:"schedule_stop_time"`
	ScheduleRelease  string `json:"schedule_release"`

	RegionCustomerPortVisible string         `json:"region_customer_port_visible"`
	SSHPort                   int64          `json:"ssh_port"`
	Image                     string         `json:"image"`
	PrivateImageUUID          string         `json:"private_image_uuid"`
	ReproductionUUID          string         `json:"reproduction_uuid"`
	ReproductionID            int64          `json:"reproduction_id"`
	CPULimit                  int64          `json:"cpu_limit"`
	MemLimitInByte            int64          `json:"mem_limit_in_byte"`
	StartMode                 string         `json:"start_mode"`
	ShmSize                   int64          `json:"shm_size"`
	ChargeType                string         `json:"charge_type"`
	ReqGPUAmount              int            `json:"req_gpu_amount"`
	PaygPrice                 int64          `json:"payg_price"`
	ExpiredAt                 TimeValid      `json:"expired_at_valid"`
	StartedAt                 TimeValid      `json:"started_at_valid"`
	StoppedAt                 TimeValid      `json:"stopped_at_valid"`
	RootFSUsedRate            float64        `json:"root_fs_used_rate"`
	TmpFSUsedRate             float64        `json:"tmp_fs_used_rate"`
	DiskHealthStatus          string         `json:"disk_health_status"`
	CPUUsagePercent           float64        `json:"cpu_usage_percent"`
	MemUsagePercent           float64        `json:"mem_usage_percent"`
	MemUsage                  int64          `json:"mem_usage"`
	MemLimit                  int64          `json:"mem_limit"`
	PullImageProgress         int64          `json:"pull_image_progress"`
	UsageInfo                 UsageInfo      `json:"usage_info"`
	ExpandDataDiskSize        int64          `json:"expand_data_disk_size"`
	RootFSUsedSize            int64          `json:"root_fs_used_size"`
	RootFSTotalSize           int64          `json:"root_fs_total_size"`
	StorageFSUsage            string         `json:"storage_fs_usage"`
	Phone                     string         `json:"phone"`
	Name                      string         `json:"name"`
	Description               string         `json:"description"`
	TimedShutdownAt           TimeValid      `json:"timed_shutdown_at_valid"`
	GPUAllNum                 int32          `json:"gpu_all_num"`
	InstanceGPUNum            int32          `json:"instance_gpu_num"`
	GPUIdleNum                int32          `json:"gpu_idle_num"`
	RentGpuNum                int32          `json:"rent_gpu_num"`
	GPUReleased               int32          `json:"gpu_released"`
	SupplierID                string         `json:"supplier_id"`
	UserID                    string         `json:"user_id"`
	GpuType                   string         `json:"gpu_type"`
	MachineFollowed           int64          `json:"machine_followed"`
	AutoTransToPayg           int64          `json:"auto_trans_to_payg"`
	AdditionalInfo            AdditionalInfo `json:"additional_info"`

	Maintenance          int32  `json:"maintenance"`
	MaintenanceMsg       string `json:"maintenance_msg"`
	MaintenanceStartTime string `json:"maintenance_start_time"`
	MaintenanceEndTime   string `json:"maintenance_end_time"`

	AutoRenew string `json:"auto_renew"`
	CreatedAt string `json:"created_at"`
}

type TimeValid struct {
	Time  string `json:"Time"`
	Valid bool   `json:"Valid"`
}

type UsageInfo struct {
	ContainerID             string  `json:"container_id"`
	ValidAt                 string  `json:"valid_at"`
	CPUUsagePercent         float64 `json:"cpu_usage_percent"`
	MemUsagePercent         float64 `json:"mem_usage_percent"`
	MemUsage                int64   `json:"mem_usage"`
	MemLimit                int64   `json:"mem_limit"`
	RootFSUsedSize          int64   `json:"root_fs_used_size"`
	RootFSTotalSize         int64   `json:"root_fs_total_size"`
	DataDiskTotalSize       int64   `json:"data_disk_total_size"`
	DataDiskUsedSize        int64   `json:"data_disk_used_size"`
	StorageFSUsage          string  `json:"storage_fs_usage"`
	PullImageProgress       int64   `json:"pull_image_progress"`
	DownloadImageProgress   int64   `json:"download_image_progress"`
	DownloadOSSFileProgress int64   `json:"download_oss_file_progress"`
	IsNew                   bool    `json:"is_new"`
}

type AdditionalInfo struct {
	CopyDataDiskAfterCloneTaskID string `json:"copy_data_disk_after_clone_task_id"`
	CopyDataDiskAfterCloneStatus string `json:"copy_data_disk_after_clone_status"`
}

type ListInstancesResponse struct {
	Pagination struct {
		Page  int `json:"page"`
		Size  int `json:"size"`
		Total int `json:"total"`
	} `json:"pagination"`
	Result []*InstanceInfo `json:"result"`
}
