package usecase

import "errors"

type LimitOffsetConfig struct {
	Limit        *int64              `json:"limit,omitempty" query:"limit"`
	Offset       *int64              `json:"offset,omitempty" query:"offset"`
}

func (f *LimitOffsetConfig) validate() error {
  if f.Limit != nil && (*f.Limit > 500 || *f.Limit < 1) {
    return errors.New("limit must be between 1 and 500")
  }

  if f.Offset != nil && *f.Offset < 0 {
    return errors.New("offset must be greater than or equal to 0")
  }

  return nil
}
