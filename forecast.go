package main
import (
	// "fmt"
	"encoding/json"
	"net/http"
	"github.com/Doc0160/Marcy/slack"
	"strconv"
	"time"
)
func do_forecastHourly(ct *CT, s Slack.OMNI) {
	e, _ := explode_cmd(s.Text)
	if len(e) > 1 {
		Typing(ct.Websocket, s)
		x, y, err := forecast(e[1])
		Typing(ct.Websocket, s)
		if err == nil {
			f := []map[string]interface{}{
				map[string]interface{}{
					"title": y.Results[0].FormattedAddress,
					"text":  x.Hourly.Summary,
					// "color" : "#0000ff",
				},
			}
			for _, v := range x.Hourly.Data {
				d := v.Time - time.Now().Unix()
				h := int(d) / 60 / 60
				if d > 0 && h < 24 {
					f = append(f, map[string]interface{}{
						"title": strconv.Itoa(time.Unix(v.Time, 0).Day()) + "/" + strconv.Itoa(int(time.Unix(v.Time, 0).Month())) + "/" + strconv.Itoa(time.Unix(v.Time, 0).Year()) + ", " + strconv.Itoa(time.Unix(v.Time, 0).Hour()) + "h",
						"text":  v.Summary,
						"fields": []map[string]interface{}{
							map[string]interface{}{
								"title": "Température",
								"value": strconv.FormatFloat(v.Temperature, 'f', -1, 64) + "°C",
								"short": true,
							},
							map[string]interface{}{
								"title": "Humidité",
								"value": strconv.FormatFloat(v.Humidity, 'f', -1, 64),
								"short": true,
							},
						},
					})
				}
			}
			Typing(ct.Websocket, s)
			_, err := ct.Slack.API_CALL("chat.postMessage", map[string]interface{}{
				"as_user":     "true",
				"channel":     s.Channel,
				"attachments": f,
			})
			if err != nil {
				Message(ct.Websocket, s, "Y'a une couille dans le paté !\n"+err.Error())
			}
		} else {
			Message(ct.Websocket, s, "Y'a une couille dans le paté !\n"+err.Error())
		}
	} else {
		Message(ct.Websocket, s, "Veuillez présciser la ville à consulter.")
	}
}

func forecast(where string) (Forecast, Coords, error) {
	client := &http.Client{}
	var a Coords
	var b Forecast
	req, err := http.NewRequest("GET", "https://maps.googleapis.com/maps/api/geocode/json?sensor=false&address="+where, nil)
	req.Header.Set("User-Agent", "I do things ? ... I'm a stupid slack bot ! (tristan.magniez@viacesi.fr)")
	r, err := client.Do(req)
	if err != nil {
		return b, a, err
	} else {
		err := json.NewDecoder(r.Body).Decode(&a)
		r.Body.Close()
		if err != nil {
			return b, a, err
		} else {
			req, err := http.NewRequest("GET", "https://api.forecast.io/forecast/16efd5f80379199411b16e47ed7c9f93/"+strconv.FormatFloat(a.Results[0].Geometry.Location.Lat, 'f', -1, 32)+","+strconv.FormatFloat(a.Results[0].Geometry.Location.Lng, 'f', -1, 32)+"/?units=si&lang=fr", nil)
			req.Header.Set("User-Agent", "I do things ? ... I'm a stupid slack bot ! (tristan.magniez@viacesi.fr)")
			r, err := client.Do(req)
			if err != nil {
				return b, a, err
			} else {
				err := json.NewDecoder(r.Body).Decode(&b)
				r.Body.Close()
				if err != nil {
					return b, a, err
				} else {
					return b, a, nil
				}
			}
		}
	}
}
type Coords struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Bounds struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"bounds"`
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceID string   `json:"place_id"`
		Types   []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}
type Forecast struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Offset    float64 `json:"offset"`
	Currently struct {
		Time                int64   `json:"time"`
		Summary             string  `json:"summary"`
		Icon                string  `json:"icon"`
		PrecipIntensity     float64 `json:"precipIntensity"`
		PrecipProbability   float64 `json:"precipProbability"`
		Temperature         float64 `json:"temperature"`
		ApparentTemperature float64 `json:"apparentTemperature"`
		DewPoint            float64 `json:"dewPoint"`
		Humidity            float64 `json:"humidity"`
		WindSpeed           float64 `json:"windSpeed"`
		WindBearing         float64 `json:"windBearing"`
		CloudCover          float64 `json:"cloudCover"`
		Pressure            float64 `json:"pressure"`
		Ozone               float64 `json:"ozone"`
	} `json:"currently"`
	Hourly struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                int64   `json:"time"`
			Summary             string  `json:"summary"`
			Icon                string  `json:"icon"`
			PrecipIntensity     float64 `json:"precipIntensity"`
			PrecipProbability   float64 `json:"precipProbability"`
			Temperature         float64 `json:"temperature"`
			ApparentTemperature float64 `json:"apparentTemperature"`
			DewPoint            float64 `json:"dewPoint"`
			Humidity            float64 `json:"humidity"`
			WindSpeed           float64 `json:"windSpeed"`
			WindBearing         float64 `json:"windBearing"`
			CloudCover          float64 `json:"cloudCover"`
			Pressure            float64 `json:"pressure"`
			Ozone               float64 `json:"ozone"`
		} `json:"data"`
	} `json:"hourly"`
	Daily struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                       int64   `json:"time"`
			Summary                    string  `json:"summary"`
			Icon                       string  `json:"icon"`
			SunriseTime                float64 `json:"sunriseTime"`
			SunsetTime                 float64 `json:"sunsetTime"`
			MoonPhase                  float64 `json:"moonPhase"`
			PrecipIntensity            float64 `json:"precipIntensity"`
			PrecipIntensityMax         float64 `json:"precipIntensityMax"`
			PrecipProbability          float64 `json:"precipProbability"`
			TemperatureMin             float64 `json:"temperatureMin"`
			TemperatureMinTime         float64 `json:"temperatureMinTime"`
			TemperatureMax             float64 `json:"temperatureMax"`
			TemperatureMaxTime         float64 `json:"temperatureMaxTime"`
			ApparentTemperatureMin     float64 `json:"apparentTemperatureMin"`
			ApparentTemperatureMinTime float64 `json:"apparentTemperatureMinTime"`
			ApparentTemperatureMax     float64 `json:"apparentTemperatureMax"`
			ApparentTemperatureMaxTime float64 `json:"apparentTemperatureMaxTime"`
			DewPoint                   float64 `json:"dewPoint"`
			Humidity                   float64 `json:"humidity"`
			WindSpeed                  float64 `json:"windSpeed"`
			WindBearing                float64 `json:"windBearing"`
			CloudCover                 float64 `json:"cloudCover"`
			Pressure                   float64 `json:"pressure"`
			Ozone                      float64 `json:"ozone"`
			PrecipIntensityMaxTime     float64 `json:"precipIntensityMaxTime,omitempty"`
			PrecipType                 string  `json:"precipType,omitempty"`
		} `json:"data"`
	} `json:"daily"`
	Flags struct {
		Sources       []string `json:"sources"`
		MetnoLicense  string   `json:"metno-license"`
		IsdStations   []string `json:"isd-stations"`
		MadisStations []string `json:"madis-stations"`
		Units         string   `json:"units"`
	} `json:"flags"`
}
