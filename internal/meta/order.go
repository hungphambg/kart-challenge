package meta

type OrderStatus uint8

const (
	OrderStatusPending OrderStatus = 1

	OrderStatusDone OrderStatus = 10
)
