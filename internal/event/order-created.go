package event

import "time"

type OrderCreated struct {
	Name     string
	Payoload interface{}
}

func NewOrderCreated() *OrderCreated {
	return &OrderCreated{
		Name: "OrderCreated",
	}
}

func (o *OrderCreated) GetName() string {
	return o.Name
}

func (o *OrderCreated) GetPayload() interface{} {
	return o.Payoload
}

func (o *OrderCreated) SetPayload(payload interface{}) {
	o.Payoload = payload
}

func (o *OrderCreated) GetDateTime() time.Time {
	return time.Now()
}
