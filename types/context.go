package types

type AppContext interface {
	Value(key any) any
}
