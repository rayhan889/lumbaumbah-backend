package events

type LaundryRequestStatusUpdated struct {
	RequestID string
	UserID    string
	AdminID   string
	Status    string
}

func (e LaundryRequestStatusUpdated) Eventname() string {
	return "LaundryRequestStatusUpdated"
}
