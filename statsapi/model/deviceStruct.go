package model

const (
	C500       = "C500"
	C404       = "C404"
	FAIL       = "0099"
	SUCCESS    = "0000"
	RETRY      = "1000"
	VALIDATION = "2000"
)
const (
	SUCCMESSAGE   = "정상"
	FAILMESSAGE   = "비정상"
	FILEERMESSAGE = "파일 에러"
	TYPE          = "STATS"
)

const PerRowCount = 10

// update, insert, delete response format
type RespIUDFormat struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        interface{}  `json:"data"`
}

// 장비 리스트
type RespGetDeviceListAll struct {
	Code        string             `json:"code"`
	Message     MessageValue       `json:"message"`
	ServiceName string             `json:"serviceName"`
	Data        []GetDeviceListAll `json:"data"`
}

type GetDeviceListAll struct {
	Uuid       string  `json:"uuid"`
	DeviceName string  `json:"deviceName"`
	PublicIp   string  `json:"publicIp"`
	Cpu        float64 `json:"cpu"`
	Memory     float64 `json:"memory"`
	Storage    float64 `json:"storage"`
	Network    float64 `json:"network"`
	Time       string  `json:"time"`
}

// 장비 정보
type RespGetDeviceInfoOne struct {
	Code        string           `json:"code"`
	Message     MessageValue     `json:"message"`
	ServiceName string           `json:"serviceName"`
	Data        GetDeviceInfoOne `json:"data"`
}

type GetDeviceInfoOne struct {
	Uuid       string  `json:"uuid"`
	DeviceName string  `json:"deviceName"`
	PublicIp   string  `json:"publicIp"`
	Cpu        float64 `json:"cpu"`
	Memory     float64 `json:"memory"`
	Storage    float64 `json:"storage"`
	Network    float64 `json:"network"`
	Time       string  `json:"time"`
}

// 장비 그래프 평균 최대 최저
type RespGetDeviceUsageOne struct {
	Code        string            `json:"code"`
	Message     MessageValue      `json:"message"`
	ServiceName string            `json:"serviceName"`
	Data        GetDeviceUsageOne `json:"data"`
}

type GetDeviceUsageOne struct {
	Uuid           string  `json:"uuid"`
	CpuAverage     float64 `json:"cpuAverage"`
	CpuMax         float64 `json:"cpuMax"`
	CpuMin         float64 `json:"cpuMin"`
	MemoryAverage  float64 `json:"memoryAverage"`
	MemoryMax      float64 `json:"memoryMax"`
	MemoryMin      float64 `json:"memoryMin"`
	StorageAverage float64 `json:"storageAverage"`
	StorageMax     float64 `json:"storageMax"`
	StorageMin     float64 `json:"storageMin"`
	NetworkAverage float64 `json:"networkAverage"`
	NetworkMax     float64 `json:"networkMax"`
	NetworkMin     float64 `json:"networkMin"`
}

// 장비 그래프 일 쿼리
type RespGetDeviceDayQueryAll struct {
	Code        string                 `json:"code"`
	Message     MessageValue           `json:"message"`
	ServiceName string                 `json:"serviceName"`
	Data        []GetDeviceDayQueryAll `json:"data"`
}

type GetDeviceDayQueryAll struct {
	Uuid           string  `json:"uuid"`
	Kst            string  `json:"kst"`
	CpuAverage     float64 `json:"cpuAverage"`
	MemoryAverage  float64 `json:"memoryAverage"`
	StorageAverage float64 `json:"storageAverage"`
	NetworkAverage float64 `json:"networkAverage"`
}

// 장비 그래프 월 쿼리
type RespGetDeviceMonthQueryAll struct {
	Code        string                   `json:"code"`
	Message     MessageValue             `json:"message"`
	ServiceName string                   `json:"serviceName"`
	Data        []GetDeviceMonthQueryAll `json:"data"`
}

type GetDeviceMonthQueryAll struct {
	Uuid           string  `json:"uuid"`
	Kst            string  `json:"kst"`
	CpuAverage     float64 `json:"cpuAverage"`
	MemoryAverage  float64 `json:"memoryAverage"`
	StorageAverage float64 `json:"storageAverage"`
	NetworkAverage float64 `json:"networkAverage"`
}
