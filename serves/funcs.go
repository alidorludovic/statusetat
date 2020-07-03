package serves

import (
	"encoding/json"
	"html/template"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/ArthurHlt/statusetat/markdown"
	"github.com/ArthurHlt/statusetat/models"
	"github.com/dustin/go-humanize"
)

func iconState(state models.ComponentState) string {
	switch state {
	case models.DegradedPerformance:
		return "error"
	case models.PartialOutage:
		return "remove_circle"
	case models.UnderMaintenance:
		return "watch_later"
	case models.MajorOutage:
		return "cancel"
	}
	return "check_circle"
}

func colorState(state models.ComponentState) string {
	switch state {
	case models.DegradedPerformance:
		return "purple"
	case models.PartialOutage:
		return "deep-orange"
	case models.UnderMaintenance:
		return "grey"
	case models.MajorOutage:
		return "red"
	}
	return "green"
}

func colorIncidentState(state models.IncidentState) string {
	switch state {
	case models.Unresolved:
		return "deep-orange"
	case models.Monitoring:
		return "blue"
	}
	return "green"
}

func timeFormat(t time.Time) string {

	return t.Format("Jan 02, 15:04 MST")
}

func timeFmtCustom(layout string, t time.Time) string {
	return t.Format(layout)
}

func timeStdFormat(t time.Time) string {
	return t.Format(time.RFC3339)
}

func stateFromIncidents(incidents []models.Incident) models.ComponentState {
	state := models.Operational

	for _, incident := range incidents {
		if incident.ComponentState > state {
			state = incident.ComponentState
		}
	}
	return state
}

func safeHTML(content string) template.HTML {
	return template.HTML(content)
}

func jsonify(content interface{}) template.JS {
	b, _ := json.Marshal(content)
	return template.JS(b)
}

func listMap(strs []string) template.JS {
	tags := make(map[string]interface{})
	for _, c := range strs {
		tags[c] = nil
	}
	return jsonify(tags)
}

func ref(d interface{}) interface{} {
	if reflect.TypeOf(d).Kind() != reflect.Ptr {
		return d
	}
	value := reflect.ValueOf(d)
	if !value.IsZero() {
		return value.Elem().Interface()
	}
	return reflect.New(value.Type().Elem()).Elem().Interface()
}

func tagify(strs []string) template.JS {
	tags := make([]map[string]interface{}, 0)
	for _, c := range strs {
		tags = append(tags, map[string]interface{}{
			"tag": c,
		})
	}
	return jsonify(tags)
}

func humanTime(t time.Time) string {
	return humanize.Time(t)
}

func (a Serve) timeNow() time.Time {
	return time.Now().In(a.loc)
}

func isAfterNow(t time.Time) bool {
	return time.Now().After(t)
}

func (a Serve) baseUrl() *url.URL {
	u, _ := url.Parse(a.baseInfo.BaseURL)
	return u
}

func markdownNoParaph(content string) template.HTML {
	b := markdown.Convert([]byte(content))
	content = strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(string(b)), "<p>"), "</p>")
	return template.HTML(content)
}
