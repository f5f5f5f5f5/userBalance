package balance

type User struct {
	UserId  int `json:"UserId"`
	Balance int `json:"Balance"`
}

type Service struct {
	ServiceId int `json:"ServiceId"`
	Price     int `json:"Price"`
}

type Order struct {
	OrderId   int  `json:"OrderId"`
	UserId    int  `json:"UserId"`
	ServiceId int  `json:"ServiceId"`
	Price     int  `json:"Price"`
	Reserved  bool `json:"Reserved"`
}
