package model

type Visits struct {
	Short     string `json:"short"`
	Url       string `json:"url"`
	Ip        string `json:"ip"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Time      int64  `json:"time"`
	UA        UA     `json:"ua"`
}

type UA struct {
	UA      string  `json:"ua"`
	Browser Browser `json:"browser"`
	Engine  Engine  `json:"engine"`
	OS      OS      `json:"os"`
	Device  Device  `json:"device"`
	CPU     CPU     `json:"CPU"`
}

type Browser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Major   string `json:"major"`
}

type Engine struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type OS struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Device struct {
	Model  string `json:"model"`
	Type   string `json:"type"`
	Vendor string `json:"vendor"`
}

type CPU struct {
	Architecture string `json:"architecture"`
}

type VisitsLog struct {
	UUID            string `json:"uuid" gorm:"primarykey"`
	Short           string `json:"short"`
	Url             string `json:"url"`
	Ip              string `json:"ip"`
	Region          string `json:"region"`
	Country         string `json:"country"`
	City            string `json:"city"`
	Longitude       string `json:"longitude"`
	Latitude        string `json:"latitude"`
	UA              string `json:"ua"`
	BrowserName     string `json:"browser_name"`
	BrowserVersion  string `json:"browser_version"`
	BrowserMajor    string `json:"browser_major"`
	EngineName      string `json:"engine_name"`
	EngineVersion   string `json:"engine_version"`
	OSName          string `json:"os_name"`
	OSVersion       string `json:"os_version"`
	DeviceModel     string `json:"device_model"`
	DeviceType      string `json:"device_type"`
	DeviceVendor    string `json:"device_vendor"`
	CPUArchitecture string `json:"cpu_architecture"`
	Time            int64  `json:"time"`
}

func (v *VisitsLog) TableName() string {
	return "visits"
}
