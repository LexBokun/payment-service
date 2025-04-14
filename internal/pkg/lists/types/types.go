package types

// Type enumeratior
type Type string

func (t Type) String() string {
	return string(t)
}

var (
	Type_WITHDRAW Type = "WITHDRAW"
	Type_INVOICE  Type = "INVOICE"
	Type_EXCHANGE Type = "EXCHANGE"
)
