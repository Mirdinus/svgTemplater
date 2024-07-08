package main

import "time"

type Event struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time
}

type NamedayResponse struct {
	Day     int `json:"day"`
	Month   int `json:"month"`
	Nameday struct {
		Name string `json:"cz"`
	} `json:"nameday"`
	Country string `json:"country"`
}

type TodoistResponse []struct {
	CreatorID    string    `json:"creator_id"`
	CreatedAt    time.Time `json:"created_at"`
	AssigneeID   string    `json:"assignee_id"`
	AssignerID   string    `json:"assigner_id"`
	CommentCount int       `json:"comment_count"`
	IsCompleted  bool      `json:"is_completed"`
	Content      string    `json:"content"`
	Description  string    `json:"description"`
	Due          struct {
		Date        string    `json:"date"`
		IsRecurring bool      `json:"is_recurring"`
		Datetime    time.Time `json:"datetime"`
		String      string    `json:"string"`
		Timezone    string    `json:"timezone"`
	} `json:"due"`
	Duration  any      `json:"duration"`
	ID        string   `json:"id"`
	Labels    []string `json:"labels"`
	Order     int      `json:"order"`
	Priority  int      `json:"priority"`
	ProjectID string   `json:"project_id"`
	SectionID string   `json:"section_id"`
	ParentID  string   `json:"parent_id"`
	URL       string   `json:"url"`
}

type ForecastWeatherData struct {
	Location struct {
		Name           string  `json:"name,omitempty"`
		Region         string  `json:"region,omitempty"`
		Country        string  `json:"country,omitempty"`
		Lat            float64 `json:"lat,omitempty"`
		Lon            float64 `json:"lon,omitempty"`
		TzID           string  `json:"tz_id,omitempty"`
		LocaltimeEpoch int     `json:"localtime_epoch,omitempty"`
		Localtime      string  `json:"localtime,omitempty"`
	} `json:"location,omitempty"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch,omitempty"`
		LastUpdated      string  `json:"last_updated,omitempty"`
		TempC            float64 `json:"temp_c,omitempty"`
		TempF            float64 `json:"temp_f,omitempty"`
		IsDay            int     `json:"is_day,omitempty"`
		Condition        struct {
			Text string `json:"text,omitempty"`
			Icon string `json:"icon,omitempty"`
			Code int    `json:"code,omitempty"`
		} `json:"condition,omitempty"`
		WindMph    float64 `json:"wind_mph,omitempty"`
		WindKph    float64 `json:"wind_kph,omitempty"`
		WindDegree int     `json:"wind_degree,omitempty"`
		WindDir    string  `json:"wind_dir,omitempty"`
		PressureMb float64 `json:"pressure_mb,omitempty"`
		PressureIn float64 `json:"pressure_in,omitempty"`
		PrecipMm   float64 `json:"precip_mm,omitempty"`
		PrecipIn   float64 `json:"precip_in,omitempty"`
		Humidity   int     `json:"humidity,omitempty"`
		Cloud      int     `json:"cloud,omitempty"`
		FeelslikeC float64 `json:"feelslike_c,omitempty"`
		FeelslikeF float64 `json:"feelslike_f,omitempty"`
		WindchillC float64 `json:"windchill_c,omitempty"`
		WindchillF float64 `json:"windchill_f,omitempty"`
		HeatindexC float64 `json:"heatindex_c,omitempty"`
		HeatindexF float64 `json:"heatindex_f,omitempty"`
		DewpointC  float64 `json:"dewpoint_c,omitempty"`
		DewpointF  float64 `json:"dewpoint_f,omitempty"`
		VisKm      float64 `json:"vis_km,omitempty"`
		VisMiles   float64 `json:"vis_miles,omitempty"`
		Uv         float64 `json:"uv,omitempty"`
		GustMph    float64 `json:"gust_mph,omitempty"`
		GustKph    float64 `json:"gust_kph,omitempty"`
	} `json:"current,omitempty"`
	Forecast struct {
		Forecastday []struct {
			Date      string `json:"date,omitempty"`
			DateEpoch int    `json:"date_epoch,omitempty"`
			Day       struct {
				MaxtempC          float64 `json:"maxtemp_c,omitempty"`
				MaxtempF          float64 `json:"maxtemp_f,omitempty"`
				MintempC          float64 `json:"mintemp_c,omitempty"`
				MintempF          float64 `json:"mintemp_f,omitempty"`
				AvgtempC          float64 `json:"avgtemp_c,omitempty"`
				AvgtempF          float64 `json:"avgtemp_f,omitempty"`
				MaxwindMph        float64 `json:"maxwind_mph,omitempty"`
				MaxwindKph        float64 `json:"maxwind_kph,omitempty"`
				TotalprecipMm     float64 `json:"totalprecip_mm,omitempty"`
				TotalprecipIn     float64 `json:"totalprecip_in,omitempty"`
				TotalsnowCm       float64 `json:"totalsnow_cm,omitempty"`
				AvgvisKm          float64 `json:"avgvis_km,omitempty"`
				AvgvisMiles       float64 `json:"avgvis_miles,omitempty"`
				Avghumidity       int     `json:"avghumidity,omitempty"`
				DailyWillItRain   int     `json:"daily_will_it_rain,omitempty"`
				DailyChanceOfRain int     `json:"daily_chance_of_rain,omitempty"`
				DailyWillItSnow   int     `json:"daily_will_it_snow,omitempty"`
				DailyChanceOfSnow int     `json:"daily_chance_of_snow,omitempty"`
				Condition         struct {
					Text string `json:"text,omitempty"`
					Icon string `json:"icon,omitempty"`
					Code int    `json:"code,omitempty"`
				} `json:"condition,omitempty"`
				Uv float64 `json:"uv,omitempty"`
			} `json:"day,omitempty"`
			Astro struct {
				Sunrise          string `json:"sunrise,omitempty"`
				Sunset           string `json:"sunset,omitempty"`
				Moonrise         string `json:"moonrise,omitempty"`
				Moonset          string `json:"moonset,omitempty"`
				MoonPhase        string `json:"moon_phase,omitempty"`
				MoonIllumination int    `json:"moon_illumination,omitempty"`
				IsMoonUp         int    `json:"is_moon_up,omitempty"`
				IsSunUp          int    `json:"is_sun_up,omitempty"`
			} `json:"astro,omitempty"`
			Hour []struct {
				TimeEpoch int     `json:"time_epoch,omitempty"`
				Time      string  `json:"time,omitempty"`
				TempC     float64 `json:"temp_c,omitempty"`
				TempF     float64 `json:"temp_f,omitempty"`
				IsDay     int     `json:"is_day,omitempty"`
				Condition struct {
					Text string `json:"text,omitempty"`
					Icon string `json:"icon,omitempty"`
					Code int    `json:"code,omitempty"`
				} `json:"condition,omitempty"`
				WindMph      float64 `json:"wind_mph,omitempty"`
				WindKph      float64 `json:"wind_kph,omitempty"`
				WindDegree   int     `json:"wind_degree,omitempty"`
				WindDir      string  `json:"wind_dir,omitempty"`
				PressureMb   float64 `json:"pressure_mb,omitempty"`
				PressureIn   float64 `json:"pressure_in,omitempty"`
				PrecipMm     float64 `json:"precip_mm,omitempty"`
				PrecipIn     float64 `json:"precip_in,omitempty"`
				SnowCm       float64 `json:"snow_cm,omitempty"`
				Humidity     int     `json:"humidity,omitempty"`
				Cloud        int     `json:"cloud,omitempty"`
				FeelslikeC   float64 `json:"feelslike_c,omitempty"`
				FeelslikeF   float64 `json:"feelslike_f,omitempty"`
				WindchillC   float64 `json:"windchill_c,omitempty"`
				WindchillF   float64 `json:"windchill_f,omitempty"`
				HeatindexC   float64 `json:"heatindex_c,omitempty"`
				HeatindexF   float64 `json:"heatindex_f,omitempty"`
				DewpointC    float64 `json:"dewpoint_c,omitempty"`
				DewpointF    float64 `json:"dewpoint_f,omitempty"`
				WillItRain   int     `json:"will_it_rain,omitempty"`
				ChanceOfRain int     `json:"chance_of_rain,omitempty"`
				WillItSnow   int     `json:"will_it_snow,omitempty"`
				ChanceOfSnow int     `json:"chance_of_snow,omitempty"`
				VisKm        float64 `json:"vis_km,omitempty"`
				VisMiles     float64 `json:"vis_miles,omitempty"`
				GustMph      float64 `json:"gust_mph,omitempty"`
				GustKph      float64 `json:"gust_kph,omitempty"`
				Uv           float64 `json:"uv,omitempty"`
			} `json:"hour,omitempty"`
		} `json:"forecastday,omitempty"`
	} `json:"forecast,omitempty"`
}
