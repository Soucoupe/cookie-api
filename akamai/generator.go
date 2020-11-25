package akamai

import (
	"fmt"
	"math"
	"strconv"

	"github.com/GuapAIO/akamai-api/db"
)

type Fpcf struct {
	FpValStr string  `json:"fpValstr"`
	RVal     string  `json:"rVal"`
	RCfp     string  `json:"rCFP"` // need to collect this from users
	Td       float64 `json:"td"`
}

type Bmak struct {
	UserAgent string `json:"useragent"`
	URL       string `json:"url"` // URL to generate abck for

	Z1   int64 `json:"z1"`
	Abck string

	StartTimeStamp int64   `json:"start_ts"`
	Xagg           string  `json:"xagg"`
	Psub           string  `json:"psub"`
	Lang           string  `json:"lang"`
	Prod           string  `json:"prod"`
	Plen           string  `json:"plen"`
	Pen            string  `json:"pen"`
	Wen            string  `json:"wen"`
	Den            string  `json:"den"`
	DeviceEvents   string  `json:"device_events"`
	Fas            string  `json:"fas"`
	Sed            string  `json:"sed"`
	Kact           string  `json:"kact"`
	Mact           string  `json:"mact"`
	Tact           string  `json:"tact"`
	Doact          string  `json:"doact"`
	Dmact          string  `json:"dmact"`
	Pact           string  `json:"pact"`
	Vcact          string  `json:"vcact"`
	KeVel          float64 `json:"ke_vel"`
	MeVel          float64 `json:"me_vel"`
	TeVel          float64 `json:"te_vel"`
	DoeVel         float64 `json:"doe_vel"`
	DmeVel         float64 `json:"dme_vel"`
	PeVel          float64 `json:"pe_vel"`
	InitTime       float64 `json:"init_time"`
	KeCnt          float64 `json:"ke_cnt"`
	MeCnt          float64 `json:"me_cnt"`
	PeCnt          float64 `json:"pe_cnt"`
	TeCnt          float64 `json:"te_cnt"`
	Ta             float64 `json:"ta"`
	NCk            float64 `json:"n_ck"`
	AjType         string  `json:"aj_type"`
	AjIndx         string  `json:"aj_indx"`
	Mr             string  `json:"mr"`
	NavPerm        string  `json:"nav_perm"`
	APIPublicKey   string  `json:"api_public_key"`
	Cs             string  `json:"cs"`
	Tst            int64   `json:"tst"`
	WebGLRender    string  `json:"webglRender"`
	Ts             float64 `json:"ts"`

	Fpcf Fpcf `json:"fpcf"`
}

func GenCookie(url string, abckCookie string, sensorInfo db.SensorData) (string, error) {
	startTimestamp := GetCfDate()

	abck := &Bmak{
		UserAgent:      sensorInfo.Navigator.UserAgent,
		URL:            url,
		StartTimeStamp: startTimestamp,
		Z1:             Z1(startTimestamp),
		Abck:           abckCookie,
		Xagg:           "12147",
		Psub:           sensorInfo.Navigator.ProductSub,
		Lang:           sensorInfo.Navigator.Language,
		Prod:           "Gecko",
		Plen:           "3",
		Pen:            "0",
		Wen:            "0",
		Den:            "0",
		DeviceEvents:   "do_en,dm_en,t_en",
		Fas:            "30261693",
		Sed:            "0,0,0,0,1,0,0",
		Kact:           "",
		Mact:           "",
		Tact:           "",
		Doact:          "",
		Dmact:          "",
		Pact:           "",
		Vcact:          "",
		KeVel:          0,
		MeVel:          0,
		TeVel:          0,
		DoeVel:         0,
		DmeVel:         0,
		PeVel:          0,
		InitTime:       0,
		KeCnt:          0,
		MeCnt:          0,
		PeCnt:          0,
		TeCnt:          0,
		Ta:             0,
		NCk:            0,
		AjType:         "0",
		AjIndx:         "0",
		Mr:             sensorInfo.Mr,
		NavPerm:        "8",
		APIPublicKey:   "afSbep8yjnZUjq3aL010jO15Sawj2VZfdYK8uY90uxq",
		Cs:             "0a46G5m17Vrp4o4c",
		Tst:            -1,
		WebGLRender:    "",
		Ts:             float64(GetCfDate()) - math.Floor(randomNumberFloat()*(11-5+1)+5),

		Fpcf: Fpcf{
			FpValStr: sensorInfo.FpValstr,
			RVal:     "-1",
			RCfp:     sensorInfo.RCFP,
			Td:       -999999,
		},
	}

	akamaiVersion := "1.66"

	d3 := D3()

	gd, err := Gd("chrome", BrowserData{
		UserAgent: abck.UserAgent,
		Lang:      abck.Lang,
		ScreenData: ScreenData{
			AvailHeight: sensorInfo.ScreenData.AvailHeight,
			AvailWidth:  sensorInfo.ScreenData.AvailWidth,
			Height:      sensorInfo.ScreenData.Height,
			Width:       sensorInfo.ScreenData.Width,
			InnerHeight: sensorInfo.ScreenData.InnerHeight,
			InnerWidth:  sensorInfo.ScreenData.InnerWidth,
		},
	}, int64(abck.StartTimeStamp), d3)
	if err != nil {
		return "", err
	}

	abck = DeviceAct(abck)

	v := GetH(abck)

	informInfo := "0,0,0,0,-1,113,0;0,0,0,1,1125,-1,0;0,-1,0,0,993,-1,0;"
	ajCount := abck.AjType + "," + abck.AjIndx

	sensorData := akamaiVersion + "-1,2,-94,-100," + gd + "-1,2,-94,-101," + abck.DeviceEvents + "-1,2,-94,-105," + informInfo + "-1,2,-94,-102," + informInfo + "-1,2,-94,-108," + abck.Kact + "-1,2,-94,-110," + abck.Mact + "-1,2,-94,-114," + abck.Tact + "-1,2,-94,-111," + abck.Doact + "-1,2,-94,-109," + abck.Dmact + "-1,2,-94,-114," + abck.Pact + "-1,2,-94,-103," + abck.Vcact + "-1,2,-94,-112," + abck.URL + "-1,2,-94,-115," + v + "-1,2,-94,-106," + ajCount + "-1,2,-94,-119," + abck.Mr + "-1,2,-94,-122," + abck.Sed + "-1,2,-94,-123," + "-1,2,-94,-124," + "-1,2,-94,-126," + "-1,2,-94,-127," + abck.NavPerm

	sdHash := 24 ^ Ab(sensorData)

	sensorData = sensorData + "-1,2,-94,-70," + abck.Fpcf.FpValStr + "-1,2,-94,-80," + "" + strconv.Itoa(Ab(abck.Fpcf.FpValStr)) + "-1,2,-94,-116," + strconv.Itoa(O9(d3)) + "-1,2,-94,-118," + strconv.Itoa(sdHash) + "-1,2,-94,-129," + abck.WebGLRender + "-1,2,-94,-121,"

	t := Od(abck.Cs, abck.APIPublicKey)[0:16]
	f := math.Floor(float64(GetCfDate() / 36e5))
	d := GetCfDate()

	b := t + Od(fmt.Sprintf("%.f", f), t) + sensorData

	sensorData = b + ";" + strconv.Itoa(int(GetCfDate()-abck.StartTimeStamp)) + ";" + strconv.Itoa(int(abck.Tst)) + ";" + strconv.Itoa(int(GetCfDate()-d))

	return sensorData, nil
}
