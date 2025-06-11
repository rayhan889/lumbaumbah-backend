package laundry

import (
	"fmt"
	"time"

	"github.com/rayhan889/lumbaumbah-backend/types"
	"github.com/rayhan889/lumbaumbah-backend/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateLaundryType(laundryType types.LaundryType) error {
	result := s.db.Create(&laundryType)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) GetLaundryRequestByID(id string) (types.LaundryRequest, error) {
	var laundryRequest types.LaundryRequest
	result := s.db.Where("id = ?", id).First(&laundryRequest)
	if result.Error != nil {
		return laundryRequest, result.Error
	}

	return laundryRequest, nil
}

func (s *Store) GetLaundryRequests() ([]types.LaundryRequestResponse, error) {
	var requests []types.LaundryRequestResponse

	var tempRows []types.LaundryRequestResponse
	result := s.db.
		Table("laundry_requests lr").
		Select("lr.id, u.username, lr.weight, lt.name AS laundry_type, lr.notes, lr.status AS current_status, lr.completion_date, lt.price * lr.weight AS total_price").
		Joins("LEFT JOIN laundry_types lt ON lr.laundry_type_id = lt.id").
		Joins("LEFT JOIN users u ON lr.user_id = u.id").
		Scan(&tempRows)
	if result.Error != nil {
		return requests, result.Error
	}

	fmt.Printf("Query result: %+v\n", tempRows)

	for _, row := range tempRows {
		var statusHistories []types.StatusHistoryResponse

		result := s.db.
			Table("status_histories sh").
			Where("sh.laundry_request_id = ?", row.ID).
			Scan(&statusHistories)
		if result.Error != nil {
			return requests, result.Error
		}

		requests = append(requests, types.LaundryRequestResponse{
			ID:              row.ID,
			Username:        row.Username,
			Weight:          row.Weight,
			LaundryType:     row.LaundryType,
			Notes:           row.Notes,
			CurrentStatus:   row.CurrentStatus,
			CompletionDate:  row.CompletionDate,
			TotalPrice:      row.TotalPrice,
			StatusHistories: statusHistories,
		})
	}

	return requests, nil
}

func (s *Store) GetLaundryRequestsByUserID(id string) ([]types.LaundryRequestResponse, error) {
	var requests []types.LaundryRequestResponse

	var tempRows []types.LaundryRequestResponse
	result := s.db.
		Table("laundry_requests lr").
		Select("lr.id, lr.weight, lt.name AS laundry_type, lr.notes, lr.status AS current_status, lr.completion_date, lt.price * lr.weight AS total_price").
		Joins("LEFT JOIN laundry_types lt ON lr.laundry_type_id = lt.id").
		Where("lr.user_id = ?", id).
		Scan(&tempRows)
	if result.Error != nil {
		return requests, result.Error
	}

	fmt.Printf("Query result: %+v\n", tempRows)

	for _, row := range tempRows {
		var statusHistories []types.StatusHistoryResponse

		result := s.db.
			Table("status_histories sh").
			Where("sh.laundry_request_id = ?", row.ID).
			Scan(&statusHistories)
		if result.Error != nil {
			return requests, result.Error
		}

		requests = append(requests, types.LaundryRequestResponse{
			ID:              row.ID,
			Weight:          row.Weight,
			LaundryType:     row.LaundryType,
			Notes:           row.Notes,
			CurrentStatus:   row.CurrentStatus,
			CompletionDate:  row.CompletionDate,
			TotalPrice:      row.TotalPrice,
			StatusHistories: statusHistories,
		})
	}

	return requests, nil
}

func (s *Store) GetLaundryTypes() ([]types.LaundryType, error) {
	var types []types.LaundryType
	result := s.db.Find(&types)
	if result.Error != nil {
		return types, result.Error
	}

	return types, nil
}

func (s *Store) CreateLaundryRequest(laundryRequest types.LaundryRequest) error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&laundryRequest).Error; err != nil {
		tx.Rollback()
		return err
	}
	requestId := laundryRequest.ID

	statusHistory := types.StatusHistory{
		ID:               utils.GenerateUUID(),
		LaundryRequestID: requestId,
		Status:           string(types.StatusPending),
		UpdatedAt:        time.Now().Format(time.RFC3339),
		UpdatedBy:        laundryRequest.UserID,
	}
	if err := tx.Create(&statusHistory).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *Store) GetLaundryTypeByID(id string) (types.LaundryType, error) {
	var laundryType types.LaundryType
	result := s.db.Where("id = ?", id).First(&laundryType)
	if result.Error != nil {
		return laundryType, result.Error
	}

	return laundryType, nil
}

func (s *Store) UpdateLaundryRequestStatus(status string, rId string, uId string) error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var returnedLaundryType struct {
		LaundryTypeID string  `gorm:"column:laundry_type_id"`
		Weight        float64 `gorm:"column:weight"`
		UserID        string  `gorm:"column:user_id"`
	}

	if err := tx.Model(&types.LaundryRequest{}).
		Clauses(clause.Returning{Columns: []clause.Column{{Name: "laundry_type_id"}, {Name: "weight"}, {Name: "user_id"}}}).
		Where("id = ?", rId).
		Update("status", status).
		Scan(&returnedLaundryType).Error; err != nil {
		tx.Rollback()
		return err
	}

	sh := types.StatusHistory{
		ID:               utils.GenerateUUID(),
		LaundryRequestID: rId,
		Status:           status,
		UpdatedAt:        time.Now().Format(time.RFC3339),
		UpdatedBy:        uId,
	}

	if status == string(types.StatusCompleted) && returnedLaundryType.LaundryTypeID != "" {
		var price float64
		ltId := returnedLaundryType.LaundryTypeID
		err := tx.Model(&types.LaundryType{}).Select("price").Where("id = ?", ltId).Pluck("price", &price).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		totalPrice := price * returnedLaundryType.Weight
		inv := types.Invoice{
			ID:               utils.GenerateUUID(),
			UserID:           returnedLaundryType.UserID,
			AdminID:          nil,
			Amount:           totalPrice,
			PaymentMethod:    string(types.PaymentMethodBankTransfer),
			Status:           string(types.StatusPending),
			IssuedAt:         time.Now().Format(time.RFC3339),
			LaundryRequestID: rId,
		}

		if err := tx.Create(&inv).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Create(&sh).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
