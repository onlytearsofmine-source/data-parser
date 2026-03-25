package helpers

import (
	"fmt"
	"strings"
	"time"

	"github.com/yourproject/data-parser/models"
)

const (
	defaultDateFormat = "2006-01-02"
)

type duration struct {
	time.Duration
}

func (d duration) String() string {
	if d.Duration == 0 {
		return "0s"
	}
	dur := time.Duration(d.Duration)
	hours := int(dur.Hours())
	if hours > 0 {
		return fmt.Sprintf("%dd %dh", hours/24, hours%24)
	}
	days := int(dur.Hours() / 24)
	if days > 0 {
		return fmt.Sprintf("%dd", days)
	}
	minutes := int(dur.Minutes())
	if minutes > 0 {
		return fmt.Sprintf("%dm", minutes)
	}
	return fmt.Sprintf("%ds", int(dur.Seconds()))
}

func extractDuration(str string) (duration, error) {
	dur := strings.Split(str, " ")
	if len(dur) != 2 {
		return duration{}, fmt.Errorf("invalid duration format: %s", str)
	}
	var (
		num  float64
		unit string
	)
	var err error
	num, err = strconv.ParseFloat(dur[0], 64)
	if err != nil {
		return duration{}, err
	}
	unit = dur[1]
	switch unit {
	case "d", "days":
		return duration{time.Duration(num * 24 * time.Hour)}, nil
	case "h", "hours":
		return duration{time.Duration(num * time.Hour)}, nil
	case "m", "minutes":
		return duration{time.Duration(num * time.Minute)}, nil
	case "s", "seconds":
		return duration{time.Duration(num * time.Second)}, nil
	default:
		return duration{}, fmt.Errorf("unsupported duration unit: %s", unit)
	}
}

func stripPrefix(str, prefix string) string {
	if !strings.HasPrefix(str, prefix) {
		return str
	}
	return str[len(prefix):]
}

func validateDate(date string) (bool, error) {
	t, err := time.Parse(defaultDateFormat, date)
	if err != nil {
		return false, err
	}
	return true, nil
}

func stripWhitespace(s string) string {
	return strings.TrimSpace(s)
}

func convertToModel(value interface{}) (models.Data, error) {
	switch v := value.(type) {
	case string:
		return models.Data{Value: v}, nil
	case float64:
		return models.Data{Value: v}, nil
	case time.Time:
		return models.Data{Value: v.Format(defaultDateFormat)}, nil
	default:
		return models.Data{}, fmt.Errorf("unsupported type: %T", v)
	}
}