package order

type Order struct {
	OrderId   int  `json:"OrderId"`
	UserId    int  `json:"UserId"`
	ServiceId int  `json:"ServiceId"`
	Price     int  `json:"Price"`
	Reserved  bool `json:"Reserved"`
}
