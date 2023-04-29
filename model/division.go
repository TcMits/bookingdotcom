package model

type Division struct {
	ProvinceName string `json:"provinceName" bson:"provinceName" fake:"Thành phố Hồ Chí Minh"`
	ProvinceCode int64  `json:"provinceCode" bson:"provinceCode" fake:"79"`
	DistrictName string `json:"districtName" bson:"districtName" fake:"Quận 3"`
	DistrictCode int64  `json:"districtCode" bson:"districtCode" fake:"770"`
	WardName     string `json:"wardName" bson:"wardName" fake:"Phường 14"`
	WardCode     int64  `json:"wardCode" bson:"wardCode" fake:"27127"`
}

type Address struct {
	Division `json:",inline" bson:"inline"`
  Address  string   `json:"address" bson:"address" fake:"{street}"`
}
