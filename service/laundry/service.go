package laundry

import "time"

func calculateCompletionDate(days int) string {
	stringDate := time.Now().AddDate(0, 0, days)
	return stringDate.Format(time.RFC3339)
}