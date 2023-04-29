package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/TcMits/bookingdotcom"
	"github.com/TcMits/bookingdotcom/dao"
	"github.com/TcMits/bookingdotcom/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StayAPIUseCaseFindStaysQuery struct {
	LimitOffsetConfig `json:",inline"`
	ProvinceCode      *int64 `json:"provinceCode,omitempty" query:"provinceCode"`
	DistrictCode      *int64 `json:"districtCode,omitempty" query:"districtCode"`
	WardCode          *int64 `json:"wardCode,omitempty" query:"wardCode"`
	Adults            *int64 `json:"adults,omitempty" query:"adults"`
	// Guests is a pointer to an array of 2 int64s, which is a bit weird.
	// first element is adults, second is children
	Guests []int64 `json:"guests,omitempty" query:"guests"`
	// CheckTimes is a pointer to an array of 2 types.DateTime, which is also weird.
	// first element is checkin, second is checkout
	CheckTimes []time.Time `json:"checkTimes,omitempty" query:"checkTimes"`
}

type StayAPIUseCaseFindStaysConfig struct {
	StayID                       *primitive.ObjectID `json:"stayId,omitempty" param:"id"`
	StayAPIUseCaseFindStaysQuery `json:",inline"`
}

func (f *StayAPIUseCaseFindStaysConfig) validate() error {
	err := f.LimitOffsetConfig.validate()
	if err != nil {
		return err
	}

	if f.StayID != nil && (f.Limit != nil || f.Offset != nil) {
		return errors.New("limit and offset are not allowed when stayId is specified")
	}

	if f.Guests != nil && len(f.Guests) != 2 {
		return errors.New("guests must be an array of 2 integers, adults and children")
	}

	if f.CheckTimes != nil && len(f.CheckTimes) != 2 {
		return errors.New("checkTimes must be an array of 2 time.Time, checkin and checkout")
	}

	return nil
}

func (f *StayAPIUseCaseFindStaysConfig) findOptions() []*options.FindOptions {
	opts := make([]*options.FindOptions, 0, 2)

	if f.Limit != nil {
		opts = append(opts, options.Find().SetLimit(*f.Limit))
	} else {
		opts = append(opts, options.Find().SetLimit(10))
	}

	if f.Offset != nil {
		opts = append(opts, options.Find().SetSkip(*f.Offset))
	} else {
		opts = append(opts, options.Find().SetSkip(0))
	}

	return opts
}

func (f *StayAPIUseCaseFindStaysConfig) countOptions() []*options.CountOptions {
	return nil
}

func (f *StayAPIUseCaseFindStaysConfig) filter() any {
	filter := make(bson.M, 6)

	if f.StayID != nil {
		filter["_id"] = *f.StayID
	}

	if f.ProvinceCode != nil {
		filter["provinceCode"] = *f.ProvinceCode
	}

	if f.DistrictCode != nil {
		filter["districtCode"] = *f.DistrictCode
	}

	if f.WardCode != nil {
		filter["wardCode"] = *f.WardCode
	}

	roomFilter := make(bson.M, 3)
	if f.Guests != nil {
		roomFilter["maxAdultGuests"] = bson.M{"$gte": f.Guests[0]}
		roomFilter["maxChildrenGuests"] = bson.M{"$gte": f.Guests[1]}
	}

	if f.CheckTimes != nil {
		roomFilter["reservedTimes"] = bson.M{
			"$not": bson.M{
				"$elemMatch": bson.M{
					"$or": bson.A{
						bson.M{
							"$and": bson.A{
								bson.M{"from": bson.M{"$gte": f.CheckTimes[0]}},
								bson.M{"from": bson.M{"$lte": f.CheckTimes[1]}},
							},
						},
						bson.M{
							"$and": bson.A{
								bson.M{"to": bson.M{"$gte": f.CheckTimes[0]}},
								bson.M{"to": bson.M{"$lte": f.CheckTimes[1]}},
							},
						},
					},
				},
			},
		}
	}

	if len(roomFilter) > 0 {
		filter["rooms"] = bson.M{"$elemMatch": roomFilter}
	}

	return filter
}

type StayAPIUseCaseReserveRoomBody struct {
	RoomCode           string `json:"roomCode"`
	model.ReservedTime `json:",inline"`
}

type StayAPIUseCaseReserveRoomConfig struct {
	StayID                        primitive.ObjectID `json:"stayId" param:"id"`
	StayAPIUseCaseReserveRoomBody `json:",inline"`
}

func (f *StayAPIUseCaseReserveRoomConfig) validate() error {
	if f.StayID.IsZero() {
		return errors.New("id is required")
	}

	now := time.Now()
	if f.RoomCode == "" {
		return errors.New("roomCode is required")
	}

	fromTime := f.From.Time()
	if fromTime.IsZero() || fromTime.Before(now) {
		return errors.New("from is invalid, from must be in the future")
	}

	toTime := f.To.Time()
	if toTime.IsZero() || toTime.Before(now) || toTime.Before(fromTime) {
		return errors.New("to to invalid, to must be in the future and after from")
	}

	receiveTime := f.ReceiveTime.Time()
	if receiveTime.IsZero() || receiveTime.Before(now) || receiveTime.After(fromTime) {
		return errors.New("receiveTime is invalid, receiveTime must be before from")
	}

	if f.Name == "" {
		return errors.New("name is required")
	}

	if f.Phone == "" {
		return errors.New("phone is required")
	}

	if f.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

func (f *StayAPIUseCaseReserveRoomConfig) findOptions() []*options.FindOptions {
	return nil
}

func (f *StayAPIUseCaseReserveRoomConfig) countOptions() []*options.CountOptions {
	return nil
}

func (f *StayAPIUseCaseReserveRoomConfig) filter() any {
	fmt.Println(f.RoomCode)
	return bson.M{
		"_id": f.StayID,
		"rooms": bson.M{
			"$elemMatch": bson.M{
				"code": f.RoomCode,
				"reservedTimes": bson.M{
					"$not": bson.M{
						"$elemMatch": bson.M{
							"$or": bson.A{
								bson.M{
									"$and": bson.A{
										bson.M{"from": bson.M{"$gte": f.From.Time()}},
										bson.M{"from": bson.M{"$lte": f.To.Time()}},
									},
								},
								bson.M{
									"$and": bson.A{
										bson.M{"to": bson.M{"$gte": f.From.Time()}},
										bson.M{"to": bson.M{"$lte": f.To.Time()}},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

type StayAPIUseCaseFindStaysResult struct {
	Count int64         `json:"count"`
	Items []*model.Stay `json:"items"`
}

type StayAPIUseCaseReserveRoomResult struct {
	*model.Stay `json:",inline"`
}

type StayAPIUseCase struct {
	bookingdotcom bookingdotcom.BookingDotCom
}

func NewStayAPIUseCase(bookingdotcom bookingdotcom.BookingDotCom) *StayAPIUseCase {
	return &StayAPIUseCase{
		bookingdotcom: bookingdotcom,
	}
}

func (u *StayAPIUseCase) FindStays(ctx context.Context, config *StayAPIUseCaseFindStaysConfig) (*StayAPIUseCaseFindStaysResult, error) {
	var err error
	if config == nil {
		return nil, errors.New("config is required")
	}

	if err = config.validate(); err != nil {
		return nil, err
	}

	result := &StayAPIUseCaseFindStaysResult{}
	f := config.filter()

	d := dao.NewDao(u.bookingdotcom)
	result.Count, err = d.Count(ctx, &dao.DaoCountConfig{
		CollectionName: model.CollectionNameStays,
		Filter:         f,
		Options:        config.countOptions(),
	})
	if err != nil {
		return nil, err
	}

	if config.Limit != nil {
		result.Items = make([]*model.Stay, 0, *config.Limit)
	} else {
		result.Items = make([]*model.Stay, 0, 10)
	}

	if err = d.Find(ctx, &dao.DaoFindConfig{
		CollectionName: model.CollectionNameStays,
		Destination:    &result.Items,
		Filter:         f,
		Options:        config.findOptions(),
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (u *StayAPIUseCase) ReserveRoom(ctx context.Context, config *StayAPIUseCaseReserveRoomConfig) (*StayAPIUseCaseReserveRoomResult, error) {
	var err error
	if config == nil {
		return nil, errors.New("config is required")
	}

	if err = config.validate(); err != nil {
		return nil, err
	}

	result := &StayAPIUseCaseReserveRoomResult{}
	items := make([]*model.Stay, 0, 1)

	d := dao.NewDao(u.bookingdotcom)
	if err := d.Find(ctx, &dao.DaoFindConfig{
		CollectionName: model.CollectionNameStays,
		Destination:    &items,
		Filter:         config.filter(),
		Options:        config.findOptions(),
	}); err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, errors.New("can't reserve this room")
	}

	stay := items[0]
	result.Stay = stay
	for i := range stay.Rooms {
		if stay.Rooms[i].Code != config.RoomCode {
			continue
		}

		stay.Rooms[i].ReservedTimes = append(stay.Rooms[i].ReservedTimes, config.ReservedTime)
		break
	}

	if err := d.Mutate(ctx, &dao.DaoMutateConfig{
		Model: stay,
	}); err != nil {
		return nil, err
	}

	return result, nil
}
