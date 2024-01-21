package posts

type TwoIPCEvent struct {
	Addition    any    `json:"Addition"`
	AlarmID     string `json:"AlarmId"`
	BoardID     string `json:"BoardId"`
	BoardIP     string `json:"BoardIp"`
	DataID      int    `json:"DataId"`
	GBDeviceID  string `json:"GBDeviceId"`
	GBTaskChnID string `json:"GBTaskChnId"`
	GPS         struct {
		AngleCourse float64 `json:"angleCourse"`
		Available   bool    `json:"available"`
		KSpeed      float64 `json:"kSpeed"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		NSpeed      float64 `json:"nSpeed"`
	} `json:"GPS"`
	ID               int    `json:"Id"`
	ImageData        string `json:"ImageData"`
	ImageDataLabeled string `json:"ImageDataLabeled"`
	LocalLabeledPath string `json:"LocalLabeledPath"`
	LocalRawPath     string `json:"LocalRawPath"`
	Media            struct {
		GBTransport   bool   `json:"GBTransport"`
		MediaDesc     string `json:"MediaDesc"`
		MediaHeight   int    `json:"MediaHeight"`
		MediaName     string `json:"MediaName"`
		MediaURL      string `json:"MediaUrl"`
		MediaWidth    int    `json:"MediaWidth"`
		Params        any    `json:"Params"`
		RtspTransport bool   `json:"RtspTransport"`
	} `json:"Media"`
	Result struct {
		Calling     bool      `json:"Calling"`
		Cropped     bool      `json:"Cropped"`
		Description string    `json:"Description"`
		Properties  []any     `json:"Properties"`
		RegType     string    `json:"RegType"`
		RelativeBox []float64 `json:"RelativeBox"`
		Type        string    `json:"Type"`
	} `json:"Result"`
	Summary     string `json:"Summary"`
	TaskDesc    string `json:"TaskDesc"`
	TaskSession string `json:"TaskSession"`
	Time        string `json:"Time"`
	TimeStamp   int64  `json:"TimeStamp"`
	Type        int    `json:"Type"`
	UniqueID    string `json:"UniqueId"`
	VideoFile   string `json:"VideoFile"`
}
