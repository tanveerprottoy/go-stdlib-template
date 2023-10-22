package typesext

// this package contains custom/derived/union types

// Status represents the status of a resource
// it's like an enum
type Status string

// ContextKey represents the key of a context
type ContextKey string

// LogLevel represents the level of a log
type LogLevel string

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Union interface {
	Numeric | string | bool
}
