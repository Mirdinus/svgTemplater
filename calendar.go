package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/nav-inc/datetime"
)

// fetchICS fetches the ICS calendar data from the provided URL.
func fetchICS(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("unable to fetch ICS: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read ICS data: %v", err)
	}

	return string(data), nil
}

// parseICS parses the ICS calendar data and prints the events.
func parseICS(data string) []Event {
	calendar, err := ics.ParseCalendar(strings.NewReader(data))
	if err != nil {
		log.Fatalf("unable to parse ICS data: %v", err)
	}

	events := make([]Event, 0)

	for _, event := range calendar.Events() {
		name := event.GetProperty(ics.ComponentPropertySummary).Value
		start := event.GetProperty(ics.ComponentPropertyDtStart).Value
		end := event.GetProperty(ics.ComponentPropertyDtEnd).Value

		startTime, err := datetime.Parse(start, time.UTC)
		if err != nil {
			log.Fatalf("unable to parse start time: %v", err)
		}
		endTime, err := datetime.Parse(end, time.UTC)
		if err != nil {
			log.Fatalf("unable to parse end time: %v", err)
		}

		events = append(events, Event{
			Name:      name,
			StartTime: startTime,
			EndTime:   endTime,
		})
	}

	return events
}

func formatEvent(event Event) string {
	// duration over 23 hours
	loc, _ := time.LoadLocation(Config.Calendar.LocalTimezone)

	event.StartTime = event.StartTime.In(loc)
	event.EndTime = event.EndTime.In(loc)

	if event.EndTime.Sub(event.StartTime).Hours() > 23 {
		return fmt.Sprintf("All day      %s", event.Name)
	}
	return fmt.Sprintf("%02d:%02d-%02d:%02d %s", event.StartTime.Hour(), event.StartTime.Minute(), event.EndTime.Hour(), event.EndTime.Minute(), event.Name)
}

func getNameday(day, month int) string {
	response, err := http.Get(fmt.Sprintf("%s?day=%d&month=%d&country=cz", Config.Calendar.NamedayUrl, day, month))
	if err != nil {
		log.Fatalf("Error fetching nameday: %v", err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading nameday: %v", err)
	}

	nameday := NamedayResponse{}
	err = json.Unmarshal(data, &nameday)
	if err != nil {
		log.Fatalf("Error unmarshalling nameday: %v", err)
	}

	if nameday.Nameday.Name != "" {
		return nameday.Nameday.Name
	}
	return ""
}

func getDayOfWeek(t time.Time) string {
	switch t.Weekday() {
	case time.Monday:
		return "Pondělí"
	case time.Tuesday:
		return "Úterý"
	case time.Wednesday:
		return "Středa"
	case time.Thursday:
		return "Čtvrtek"
	case time.Friday:
		return "Pátek"
	case time.Saturday:
		return "Sobota"
	case time.Sunday:
		return "Neděle"
	}
	return ""
}

func processDay(template, prefix string, day time.Time) string {
	dayEvents := make([]string, 0)
	for _, event := range events {
		if event.StartTime.Day() == day.Day() && event.StartTime.Month() == day.Month() && event.StartTime.Year() == day.Year() && event.EndTime.Sub(event.StartTime).Hours() <= 24 {
			dayEvents = append(dayEvents, formatEvent(event))
		}
	}

	for i := 0; i < len(dayEvents); i++ {
		template = strings.ReplaceAll(template, fmt.Sprintf("{%s_event_%d}", prefix, i+1), dayEvents[i])
	}
	if len(dayEvents) > 2 {
		template = strings.ReplaceAll(template, fmt.Sprintf("{%s_more_events}", prefix), "...")
	} else {
		template = strings.ReplaceAll(template, fmt.Sprintf("{%s_more_events}", prefix), "")
	}

	for i := 0; i < 6; i++ {
		template = strings.ReplaceAll(template, fmt.Sprintf("{%s_event_%d}", prefix, i+1), "")
	}

	nameDay := getNameday(day.Day(), int(day.Month()))
	// cap nameday to 30 characters
	if len(nameDay) > 30 {
		nameDay = nameDay[:30]
	}
	template = strings.ReplaceAll(template, fmt.Sprintf("{%s_nameday}", prefix), nameDay)

	template = strings.ReplaceAll(template, fmt.Sprintf("{%s_date}", prefix), fmt.Sprintf("%d.%d", day.Day(), int(day.Month())))
	template = strings.ReplaceAll(template, fmt.Sprintf("{%s_day}", prefix), getDayOfWeek(day))

	return template
}

func ProcessCalendar(formated string) string {
	// Replace with your ICS link
	icsURL := Config.Calendar.ICSUrl

	data, err := fetchICS(icsURL)
	if err != nil {
		log.Fatalf("Error fetching ICS: %v", err)
	}

	parseICS(data)

	events = parseICS(data)

	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	day_after_tomorrow := today.AddDate(0, 0, 2)

	formated = processDay(formated, "today", today)
	formated = processDay(formated, "tomorrow", tomorrow)
	formated = processDay(formated, "day_after_tomorrow", day_after_tomorrow)

	return formated
}
