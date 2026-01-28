package meta

type CartStatus uint8

const (
	CartStatusActive CartStatus = 1

	CartStatusDone   CartStatus = 10
	CartStatusDelete CartStatus = 11
)
