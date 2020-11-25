package px

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/GuapAIO/akamai-api/db"
	"github.com/google/uuid"
	"github.com/obito/cclient"
	utls "github.com/refraction-networking/utls"
)

// CollectorDo is used to transform body to json
type CollectorDo struct {
	Do []string `json:"do"`
}

type Action struct {
	Value          string
	Value2         string
	CookieName     string
	ExpirationTime string
}

func GenCookie(targetWebsite, productPage string, sensorInfo db.SensorData) (string, error) {
	id := uuid.New().String()

	payload := createFirstPayload(id, productPage)

	pxDo, err := sendCollector(targetWebsite, payload, "PXu6b0qd2S", "v6.0.2", id, "179", false, "", "", "")
	if err != nil {
		return "", err
	}

	// Add all the action into a map
	toDoFirst := saveActions(pxDo)

	sts := toDoFirst["sts"].Value
	cls1 := toDoFirst["cls"].Value
	cls2 := toDoFirst["cls"].Value2
	wcs := toDoFirst["wcs"].Value
	drc := toDoFirst["drc"].Value

	vid := toDoFirst["vid"].Value
	sid := toDoFirst["sid"].Value
	cs := toDoFirst["cs"].Value

	secondPayload := createSecondPayload(id, sensorInfo, sts, cls1, cls2, wcs, drc, vid, sid, productPage)

	pxDoSecond, err := sendCollector(targetWebsite, secondPayload, "PXu6b0qd2S", "v6.0.2", id, "179", true, vid, sid, cs)
	if err != nil {
		return "", err
	}

	/*
		log.Print(pxDoSecond)

		toDoSecond := saveActions(pxDoSecond)

		log.Print(toDoSecond)

		px3 := toDoSecond["bake"].Value
		pxde := toDoSecond["en"].Value

		log.Printf("PX3: %v - PXDE: %v", px3, pxde)
	*/
	out, err := json.Marshal(pxDoSecond)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func saveActions(toDo *CollectorDo) map[string]Action {
	actions := make(map[string]Action)

	for _, action := range toDo.Do {
		parsedAction := strings.Split(action, "|")

		if len(parsedAction) == 3 {
			actions[parsedAction[0]] = Action{
				Value:  parsedAction[1],
				Value2: parsedAction[2],
			}
		} else if len(parsedAction) <= 4 {
			actions[parsedAction[0]] = Action{
				Value: parsedAction[1],
			}
		} else {
			actions[parsedAction[0]] = Action{
				Value:          parsedAction[3],
				CookieName:     parsedAction[1],
				ExpirationTime: parsedAction[2],
			}
		}
	}

	return actions
}

func createFirstPayload(id, productLink string) string {
	currentTime := time.Now()

	tUnixMicro := int64(time.Nanosecond) * currentTime.UnixNano() / int64(time.Microsecond)

	payload := fmt.Sprintf(`[
		{
			"t": "PX2",
			"d": {
				"PX96": "%v",
				"PX63": "Win32",
				"PX191": 0,
				"PX850": 0,
				"PX851": 872,
				"PX1008": 3600,
				"PX1055": %v,
				"PX1056": %v,
				"PX1038": "%v",
				"PX371": true
			}
		}
	]`, productLink, tUnixMicro, tUnixMicro+3, id)

	return payload
}

// send the second payload with much more info
// NOTES: here are the value that you get from the actions of the first payload
// sts: PX982 ; cls1: PX983 ; cls2: PX986 ; drc: PX985 ; wcs: PX943
func createSecondPayload(id string, sensorInfo db.SensorData, sts, cls1, cls2, wcs, drc, vid, sid, productPage string) string {
	currentTime := time.Now()
	tUnixMicro := int64(time.Nanosecond) * currentTime.UnixNano() / int64(time.Microsecond)
	px357 := CreatePX(vid, sensorInfo.Navigator.UserAgent)
	px358 := CreatePX(sid, sensorInfo.Navigator.UserAgent)
	px359 := CreatePX(id, sensorInfo.Navigator.UserAgent)

	var cleanPlugins []string

	for _, plugin := range sensorInfo.Navigator.Plugins {
		cleanPlugins = append(cleanPlugins, fmt.Sprintf(`"%s"`, plugin))
	}

	formatedDate := currentTime.Format("Mon Jan 02 2006 15:04:05") + " GMT+0000 (Coordinated Universal Time)"

	payload := fmt.Sprintf(`[
		{
			"t": "PX3",
			"d": {
				"PX234": false,
				"PX235": false,
				"PX151": false,
				"PX239": false,
				"PX240": false,
				"PX152": false,
				"PX153": false,
				"PX314": false,
				"PX192": false,
				"PX196": false,
				"PX207": false,
				"PX251": false,
				"PX982": %v,
				"PX983": "%v",
				"PX986": "%v",
				"PX985": %v,
				"PX1033": "e0eaf10e",
				"PX1019": "3f2b5a30",
				"PX1020": "7766a52d",
				"PX1021": "714b1317",
				"PX1022": "6a90378d",
				"PX1035": false,
				"PX1025": false,
				"PX359": "%v",
				"PX943": "%v",
				"PX357": "%v",
				"PX358": "%v",
				"PX229": 24,
				"PX230": 24,
				"PX91": %v,
				"PX92": %v,
				"PX269": %v,
				"PX270": %v,
				"PX93": "%v",
				"PX185": %v,
				"PX186": %v,
				"PX187": 0,
				"PX188": 0,
				"PX95": true,
				"PX400": %v,
				"PX404": "144|54|54|180|68",
				"PX90": [
					"loadTimes",
					"csi",
					"app",
					"runtime"
				],
				"PX190": "",
				"PX552": "undefined",
				"PX399": "undefined",
				"PX549": 0,
				"PX411": 0,
				"PX402": 1,
				"PX548": 1,
				"PX405": true,
				"PX547": true,
				"PX134": true,
				"PX89": true,
				"PX170": 3,
				"PX85": [
					%s
				],
				"PX59": "%v",
				"PX61": "%v",
				"PX313": [
					"%v",
					"fr",
					"en-US",
					"en"
				],
				"PX63": "Win32",
				"PX86": true,
				"PX154": -60,
				"PX133": true,
				"PX88": true,
				"PX169": 4,
				"PX62": "Gecko",
				"PX69": "%v",
				"PX64": "5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36",
				"PX65": "Netscape",
				"PX66": "Mozilla",
				"PX60": true,
				"PX87": true,
				"PX821": 4294705152,
				"PX822": 28398146,
				"PX823": 19947142,
				"PX147": false,
				"PX155": "%v",
				"PX236": false,
				"PX194": false,
				"PX195": true,
				"PX237": 0,
				"PX238": "missing",
				"PX208": "visible",
				"PX218": 0,
				"PX231": 1560,
				"PX232": 3840,
				"PX254": false,
				"PX295": false,
				"PX268": false,
				"PX166": true,
				"PX138": true,
				"PX143": true,
				"PX714": "64556c77",
				"PX715": "",
				"PX724": "10207b2f",
				"PX725": "10207b2f",
				"PX729": "",
				"PX443": true,
				"PX466": true,
				"PX467": true,
				"PX468": true,
				"PX191": 0,
				"PX94": 2,
				"PX120": [],
				"PX141": false,
				"PX96": "%v",
				"PX55": "",
				"PX1065": 1,
				"PX850": 1,
				"PX851": 1518,
				"PX1054": %v,
				"PX1008": 3600,
				"PX1055": %v,
				"PX1056": %v,
				"PX1038": "%v",
				"PX371": true
			}
		}
	]`, sts, cls1, cls2, drc, px359, wcs, px357, px358, sensorInfo.ScreenData.Width, sensorInfo.ScreenData.Height, sensorInfo.ScreenData.AvailWidth, sensorInfo.ScreenData.AvailHeight, fmt.Sprintf("%vX%v", sensorInfo.ScreenData.Width, sensorInfo.ScreenData.Height), sensorInfo.ScreenData.InnerHeight, sensorInfo.ScreenData.InnerWidth, sensorInfo.PX400, strings.Join(cleanPlugins, ",\n"), sensorInfo.Navigator.UserAgent, sensorInfo.Navigator.Language, sensorInfo.Navigator.Language, sensorInfo.Navigator.ProductSub, formatedDate, productPage, tUnixMicro, tUnixMicro+420, tUnixMicro, id)

	return payload
}

// send third payload, with mouse events
func createThirdPayload(id string) string {

	return ""
}

func sendCollector(website, payload, appID, tag, uuid, ft string, secondPayload bool, vid, sid, cs string) (*CollectorDo, error) {
	obfuscatedPayload := ObfuscateString(payload, 50)

	encodedPayload := b64.StdEncoding.EncodeToString([]byte(obfuscatedPayload))

	pc := CreatePC(payload, fmt.Sprintf("%v:%v:%v", uuid, tag, ft))

	var body *strings.Reader

	if secondPayload == false {
		body = strings.NewReader("payload=" + encodedPayload + "&appId=" + appID + "&tag=" + tag + "&uuid=" + uuid + "&ft=" + ft + "&seq=0&en=NTA&pc=" + pc + "&rsc=1")
	} else {
		body = strings.NewReader("payload=" + encodedPayload + "&appId=" + appID + "&tag=" + tag + "&uuid=" + uuid + "&ft=" + ft + "&seq=0&en=NTA&cs=" + cs + "&pc=" + pc + "&sid=" + sid + "&vid=" + vid + "&rsc=2")
	}

	client, err := cclient.NewClient(utls.HelloChrome_Auto, false)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://collector-pxu6b0qd2s.px-cloud.net/api/v2/collector", body)
	if err != nil {
		return nil, err
	}

	req.Host = "collector-pxu6b0qd2s.px-cloud.net"
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"86\", \"\"Not\\A;Brand\";v=\"99\", \"Google Chrome\";v=\"86\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", website)
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", website)
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pxDo CollectorDo
	err = json.Unmarshal(bodyResp, &pxDo)
	if err != nil {
		return nil, err
	}

	return &pxDo, nil
}
