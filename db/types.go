package db

type SensorData struct {
	Id int64

	ScreenData struct {
		AvailWidth  int `json:"availWidth"`
		AvailHeight int `json:"availHeight"`
		Width       int `json:"width"`
		Height      int `json:"height"`
		ColorDepth  int `json:"colorDepth"`
		PixelDepth  int `json:"pixelDepth"`
		InnerWidth  int `json:"innerWidth"`
		InnerHeight int `json:"innerHeight"`
		OuterWidth  int `json:"outerWidth"`
	} `json:"screenData"`
	Navigator struct {
		UserAgent                     string   `json:"userAgent"`
		ProductSub                    string   `json:"productSub"`
		Language                      string   `json:"language"`
		Product                       string   `json:"product"`
		OnLine                        bool     `json:"onLine"`
		Vibrate                       bool     `json:"vibrate"`
		GetBattery                    bool     `json:"getBattery"`
		Credentials                   bool     `json:"credentials"`
		AppMinorVersion               bool     `json:"appMinorVersion"`
		Bluetooth                     bool     `json:"bluetooth"`
		Storage                       bool     `json:"storage"`
		GetGamepads                   bool     `json:"getGamepads"`
		GetStorageUpdates             bool     `json:"getStorageUpdates"`
		HardwareConcurrency           bool     `json:"hardwareConcurrency"`
		MediaDevices                  bool     `json:"mediaDevices"`
		MozAlarms                     bool     `json:"mozAlarms"`
		MozConnection                 bool     `json:"mozConnection"`
		MozIsLocallyAvailable         bool     `json:"mozIsLocallyAvailable"`
		MozPhoneNumberService         bool     `json:"mozPhoneNumberService"`
		MsManipulationViewsEnabled    bool     `json:"msManipulationViewsEnabled"`
		Permissions                   bool     `json:"permissions"`
		RegisterProtocolHandler       bool     `json:"registerProtocolHandler"`
		RequestMediaKeySystemAccess   bool     `json:"requestMediaKeySystemAccess"`
		RequestWakeLock               bool     `json:"requestWakeLock"`
		SendBeacon                    bool     `json:"sendBeacon"`
		ServiceWorker                 bool     `json:"serviceWorker"`
		StoreWebWideTrackingException bool     `json:"storeWebWideTrackingException"`
		WebkitGetGamepads             bool     `json:"webkitGetGamepads"`
		WebkitTemporaryStorage        bool     `json:"webkitTemporaryStorage"`
		CookieEnabled                 bool     `json:"cookieEnabled"`
		JavaEnabled                   bool     `json:"javaEnabled"`
		DoNotTrack                    int      `json:"doNotTrack"`
		Plugins                       []string `json:"plugins"`
	} `json:"navigator"`
	AddEventListener       bool `json:"addEventListener"`
	XMLHTTPRequest         bool `json:"XMLHttpRequest"`
	XDomainRequest         bool `json:"XDomainRequest"`
	DeviceOrientationEvent bool `json:"DeviceOrientationEvent"`
	DeviceMotionEvent      bool `json:"DeviceMotionEvent"`
	TouchEvent             bool `json:"TouchEvent"`
	Chrome                 struct {
		App struct {
			IsInstalled  bool `json:"isInstalled"`
			InstallState struct {
				DISABLED     string `json:"DISABLED"`
				INSTALLED    string `json:"INSTALLED"`
				NOTINSTALLED string `json:"NOT_INSTALLED"`
			} `json:"InstallState"`
			RunningState struct {
				CANNOTRUN  string `json:"CANNOT_RUN"`
				READYTORUN string `json:"READY_TO_RUN"`
				RUNNING    string `json:"RUNNING"`
			} `json:"RunningState"`
		} `json:"app"`
	} `json:"chrome"`
	PrototypeBind bool `json:"prototype_bind"`
	PointerEvent  bool `json:"PointerEvent"`
	CCON          bool `json:"CC_ON"`
	Document      struct {
		DocumentMode string `json:"documentMode"`
		Webdriver    bool   `json:"webdriver"`
		Driver       bool   `json:"driver"`
		Selenium     bool   `json:"selenium"`
		Hidden       bool   `json:"hidden"`
		WebkitHidden bool   `json:"webkitHidden"`
	} `json:"document"`
	InstallTrigger   bool `json:"InstallTrigger"`
	HTMLElement      bool `json:"HTMLElement"`
	WebRTC           bool `json:"webRTC"`
	MozInnerScreenY  int  `json:"mozInnerScreenY"`
	PrototypeForEach bool `json:"prototype_forEach"`
	FileReader       bool `json:"FileReader"`
	Imul             bool `json:"imul"`
	ParseInt         bool `json:"parseInt"`
	Hypot            bool `json:"hypot"`
	Value1           bool `json:"value1"`
	XPathResult      bool `json:"XPathResult"`
	SessionStorage   bool `json:"sessionStorage"`
	LocalStorage     bool `json:"localStorage"`
	IndexedDB        bool `json:"indexedDB"`
	Canvas           struct {
		Value1 string `json:"value1"`
		Value2 string `json:"value2"`
	} `json:"canvas"`
	RCFP      string `json:"rCFP"`
	FontsOptm string `json:"fonts_optm"`
	Fonts     string `json:"fonts"`
	FpValstr  string `json:"fpValstr"`
	SSH       string `json:"ssh"`
	Mr        string `json:"mr"`
	Brave     string `json:"brave"`
	Wv        string `json:"wv"`
	Wr        string `json:"wr"`
	Weh       string `json:"weh"`
	Wl        int    `json:"wl"`
	Fmh       string `json:"fmh"`
	Fmz       int    `json:"fmz"`
	Np        string `json:"np"`

	PX400 string `json:"px400"`
}
