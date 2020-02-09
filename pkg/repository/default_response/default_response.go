package defaultresponse

import "github.com/jinzhu/gorm"

type DefaultResponse struct {
	err          error
	rowsAffected int64
}

func NewDefaultResponse(query *gorm.DB) *DefaultResponse {
	return &DefaultResponse{
		err:          query.Error,
		rowsAffected: query.RowsAffected,
	}
}

func (d *DefaultResponse) RowsAffected() int64 {
	return d.rowsAffected
}

func (d *DefaultResponse) Error() error {
	return d.err
}
