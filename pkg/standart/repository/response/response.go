package response

import "github.com/jinzhu/gorm"

type Response struct {
	err          error
	rowsAffected int64
}

func NewDefaultResponse(query *gorm.DB) *Response {
	return &Response{
		err:          query.Error,
		rowsAffected: query.RowsAffected,
	}
}

func (d *Response) RowsAffected() int64 {
	return d.rowsAffected
}

func (d *Response) Error() error {
	return d.err
}
