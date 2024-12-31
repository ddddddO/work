package contract

type Protocol interface {
	GetStr() string
	Name(x string) string
	Port() uint64
}

var p Protocol

func Register(proto Protocol) {
	p = proto
}
