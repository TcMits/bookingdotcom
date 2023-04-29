package usecase

import (
	"errors"
	"strings"

	"github.com/TcMits/vnprovince"
)

var skip = errors.New("skip")

type DivisionAPIUseCaseFindConfig struct {
	LimitOffsetConfig `json:",inline"`
	Search            *string `json:"search,omitempty" query:"search"`
}

type DivisionAPIUseCaseFindResult struct {
	Count int64                  `json:"count"`
	Items []*vnprovince.Division `json:"items"`
}

type DivisionAPIUseCase struct{}

func NewDivisionAPIUseCase() *DivisionAPIUseCase {
	return &DivisionAPIUseCase{}
}

func (*DivisionAPIUseCase) Find(config *DivisionAPIUseCaseFindConfig) (*DivisionAPIUseCaseFindResult, error) {
	if config == nil {
		return nil, errors.New("config is required")
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	result := &DivisionAPIUseCaseFindResult{
    Count: vnprovince.DivisionsLength,
  }

	if config.Limit != nil {
		result.Items = make([]*vnprovince.Division, 0, *config.Limit)
	} else {
		result.Items = make([]*vnprovince.Division, 0, 10)
	}

	index := 0
	err := vnprovince.EachDivision(func(d vnprovince.Division) error {
		defer func() {
			index += 1
		}()

		limit := int64(10)
		if config.Limit != nil {
			limit = *config.Limit
		}

		offset := int64(0)
		if config.Offset != nil {
			offset = *config.Offset
		}

		if int64(len(result.Items)) >= limit {
			return skip
		}

		if int64(index) < offset {
			return nil
		}

		if config.Search != nil && *config.Search != "" {
			if searchText := strings.ToLower(
				d.ProvinceName + " " + d.DistrictName + " " + d.WardName,
			); !strings.Contains(searchText, strings.ToLower(*config.Search)) {
				return nil
			}
		}

		result.Items = append(result.Items, &d)
		return nil
	})
	if err != nil && err != skip {
		return nil, err
	}

	return result, nil
}
