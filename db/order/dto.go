package order

type CreateOrderDto struct {
	UserId    int  `json:"UserId"`
	ServiceId int  `json:"ServiceId"`
	Price     int  `json:"Price"`
	Reserved  bool `json:"Reserved"`
}
