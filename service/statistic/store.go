package statistic

import (
	"github.com/rayhan889/lumbaumbah-backend/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUsersList() ([]types.User, error) {
	var users []types.User
	result := s.db.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}

	return users, nil
}

func (s *Store) GetUserStatistics(uId string) (types.UserStatistics, error) {
	var stats types.UserStatistics
	var requests []types.LaundryRequestResponse
	var tempRows []types.LaundryRequestResponse

	err := s.db.Model(&types.LaundryRequest{}).Where("user_id = ?", uId).Count(&stats.LaundryRequestsCount).Error
	if err != nil {
		return stats, err
	}

	err = s.db.Table("laundry_requests lr").
		Select("SUM(lt.price * lr.weight) as total_price").
		Joins("LEFT JOIN laundry_types lt ON lr.laundry_type_id = lt.id").
		Scan(&stats.TotalPrice).Error
	if err != nil {
		return stats, err
	}

	err = s.db.Model(&types.LaundryRequest{}).Where("user_id = ? AND status = ?", uId, types.StatusCompleted).Count(&stats.TotalCompleted).Error
	if err != nil {
		return stats, err
	}

	err = s.db.Model(&types.LaundryRequest{}).Select("SUM(weight) as total_weight").Where("user_id = ?", uId).Scan(&stats.TotalWeight).Error
	if err != nil {
		return stats, err
	}

	result := s.db.
		Table("laundry_requests lr").
		Select("lr.id, u.username, lr.weight, lt.name AS laundry_type, lr.notes, lr.status AS current_status, lr.completion_date, lt.price * lr.weight AS total_price").
		Joins("LEFT JOIN laundry_types lt ON lr.laundry_type_id = lt.id").
		Joins("LEFT JOIN users u ON lr.user_id = u.id").
		Where("lr.user_id = ?", uId).
		Limit(5).
		Scan(&tempRows)
	if result.Error != nil {
		return stats, result.Error
	}

	for _, row := range tempRows {
		var statusHistories []types.StatusHistoryResponse

		result := s.db.
			Table("status_histories sh").
			Where("sh.laundry_request_id = ?", row.ID).
			Scan(&statusHistories)
		if result.Error != nil {
			return stats, result.Error
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
	stats.LatestRequests = requests

	return stats, nil
}
