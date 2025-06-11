package notification

import (
	"log"
	"time"

	"github.com/rayhan889/lumbaumbah-backend/domain"
	"github.com/rayhan889/lumbaumbah-backend/domain/events"
	"github.com/rayhan889/lumbaumbah-backend/types"
	"github.com/rayhan889/lumbaumbah-backend/utils"
)

func RegisterHandlers(notifStore types.NotificationStore, adminStore types.AdminStore) {
	domain.RegisterHandler("LaundryRequestStatusUpdated", func(e domain.Event) {
		evt := e.(events.LaundryRequestStatusUpdated)

		var adminId *string
		dt, _ := adminStore.GetAdminByID(evt.AdminID)
		if dt.ID != "" {
			adminId = &dt.ID
		}

		notif := types.Notification{
			ID:               utils.GenerateUUID(),
			UserID:           &evt.UserID,
			AdminID:          adminId,
			LaundryRequestID: &evt.RequestID,
			Message:          "Your laundry request status has been updated to " + evt.Status,
			IsRead:           false,
			CreatedAt:        time.Now().Format(time.RFC3339),
		}

		err := notifStore.CreateNotification(notif)
		if err != nil {
			log.Printf("Failed to create notification: %v", err)
			return
		}
	})
}
