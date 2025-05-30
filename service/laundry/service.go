package laundry

import (
	"time"

	"github.com/rayhan889/lumbaumbah-backend/types"
)

var statusLists = []string{
	string(types.StatusPending),
	string(types.StatusCanceled),
	string(types.StatusCompleted),
	string(types.StatusProcessed),
}

var allowedStatusTransitions = map[string]map[string]bool {
	string(types.StatusPending): {
		string(types.StatusCanceled): true,
		string(types.StatusCompleted): false,
		string(types.StatusProcessed): true,
	},
	string(types.StatusProcessed): {
		string(types.StatusCanceled): true,
		string(types.StatusCompleted): true,
		string(types.StatusPending): false,
	},
	string(types.StatusCompleted): {
		string(types.StatusCanceled): false,
		string(types.StatusProcessed): false,
		string(types.StatusPending): false,
	},
	string(types.StatusCanceled): {
		string(types.StatusProcessed): false,
		string(types.StatusCompleted): false,
		string(types.StatusPending): false,
	},
}

func calculateCompletionDate(days int) string {
	stringDate := time.Now().AddDate(0, 0, days)
	return stringDate.Format(time.RFC3339)
}

func checkUpdateAbility(curr_status string, status string) bool {
	for _, s := range statusLists {
		if curr_status == s && status == s {
			return false
		}
	}
	if next, ok := allowedStatusTransitions[curr_status]; ok {
		if allowed, exists := next[status]; exists {
			return allowed
		}
	}

	return false
}

func checkUpdaterAccess(role string, status string) bool {
	permission := false
	if role == "admin" && status == string(types.StatusCompleted) {
		permission = true
	}
	if role == "admin" && status == string(types.StatusCanceled) {
		permission = true
	}
	if role == "admin" && status == string(types.StatusProcessed) {
		permission = true
	}
	if role == "user" && status == string(types.StatusCanceled) {
		permission = true
	}
	if role == "user" && status == string(types.StatusCompleted) {
		permission = false
	}
	if role == "user" && status == string(types.StatusProcessed) {
		permission = false
	}

	return permission
}