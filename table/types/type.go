package types

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal"
)

// Type describes YDB data types.
type Type interface {
	internal.T
}

func List(T Type) Type {
	return internal.ListType{T: T}
}

func Tuple(elems ...Type) Type {
	es := make([]internal.T, len(elems))
	for i, el := range elems {
		es[i] = el
	}
	return internal.TupleType{
		Elems: es,
	}
}

type tStructType internal.StructType

type StructOption func(*tStructType)

func StructField(name string, typ Type) StructOption {
	return func(s *tStructType) {
		s.Fields = append(s.Fields, internal.StructField{
			Name: name,
			Type: typ,
		})
	}
}

func Struct(opts ...StructOption) Type {
	var s tStructType
	for _, opt := range opts {
		opt(&s)
	}
	return internal.StructType(s)
}

func Variant(x Type) Type {
	switch v := x.(type) {
	case internal.TupleType:
		return internal.VariantType{
			T: v,
		}
	case internal.StructType:
		return internal.VariantType{
			S: v,
		}
	default:
		panic(fmt.Sprintf("unsupported types for variant: %s", v))
	}
}

func Void() Type {
	return internal.VoidType{}
}

func Optional(T Type) Type {
	return internal.OptionalType{T: T}
}

// TODO(kamardin): rename types to consistent format like values: BoolType,
// IntType and so on. Do not forget about code generation.

// Primitive types known by YDB.
const (
	TypeUnknown      = internal.TypeUnknown
	TypeBool         = internal.TypeBool
	TypeInt8         = internal.TypeInt8
	TypeUint8        = internal.TypeUint8
	TypeInt16        = internal.TypeInt16
	TypeUint16       = internal.TypeUint16
	TypeInt32        = internal.TypeInt32
	TypeUint32       = internal.TypeUint32
	TypeInt64        = internal.TypeInt64
	TypeUint64       = internal.TypeUint64
	TypeFloat        = internal.TypeFloat
	TypeDouble       = internal.TypeDouble
	TypeDate         = internal.TypeDate
	TypeDatetime     = internal.TypeDatetime
	TypeTimestamp    = internal.TypeTimestamp
	TypeInterval     = internal.TypeInterval
	TypeTzDate       = internal.TypeTzDate
	TypeTzDatetime   = internal.TypeTzDatetime
	TypeTzTimestamp  = internal.TypeTzTimestamp
	TypeString       = internal.TypeString
	TypeUTF8         = internal.TypeUTF8
	TypeYSON         = internal.TypeYSON
	TypeJSON         = internal.TypeJSON
	TypeUUID         = internal.TypeUUID
	TypeJSONDocument = internal.TypeJSONDocument
	TypeDyNumber     = internal.TypeDyNumber
)

func WriteTypeStringTo(buf *bytes.Buffer, t Type) {
	internal.WriteTypeStringTo(buf, t)
}

// RawValue scanning non-primitive yql types or for own implementation scanner native API
type RawValue interface {
	Path() string
	WritePathTo(w io.Writer) (n int64, err error)
	Type() Type
	Bool() (v bool)
	Int8() (v int8)
	Uint8() (v uint8)
	Int16() (v int16)
	Uint16() (v uint16)
	Int32() (v int32)
	Uint32() (v uint32)
	Int64() (v int64)
	Uint64() (v uint64)
	Float() (v float32)
	Double() (v float64)
	Date() (v time.Time)
	Datetime() (v time.Time)
	Timestamp() (v time.Time)
	Interval() (v time.Duration)
	TzDate() (v time.Time)
	TzDatetime() (v time.Time)
	TzTimestamp() (v time.Time)
	String() (v string)
	UTF8() (v string)
	YSON() (v []byte)
	JSON() (v []byte)
	UUID() (v [16]byte)
	JSONDocument() (v []byte)
	DyNumber() (v string)
	Value() Value

	// Any returns any primitive or optional value.
	// Currently, it may return one of these types:
	//
	//   bool
	//   int8
	//   uint8
	//   int16
	//   uint16
	//   int32
	//   uint32
	//   int64
	//   uint64
	//   float32
	//   float64
	//   []byte
	//   string
	//   [16]byte
	//
	Any() interface{}

	// Unwrap unwraps current item under scan interpreting it as Optional<T> types.
	Unwrap()
	AssertType(t Type) bool
	IsNull() bool
	IsOptional() bool

	// ListIn interprets current item under scan as a ydb's list.
	// It returns the size of the nested items.
	// If current item under scan is not a list types, it returns -1.
	ListIn() (size int)

	// ListItem selects current item i-th element as an item to scan.
	// ListIn() must be called before.
	ListItem(i int)

	// ListOut leaves list entered before by ListIn() call.
	ListOut()

	// TupleIn interprets current item under scan as a ydb's tuple.
	// It returns the size of the nested items.
	TupleIn() (size int)

	// TupleItem selects current item i-th element as an item to scan.
	// Note that TupleIn() must be called before.
	// It panics if it is out of bounds.
	TupleItem(i int)

	// TupleOut leaves tuple entered before by TupleIn() call.
	TupleOut()

	// StructIn interprets current item under scan as a ydb's struct.
	// It returns the size of the nested items – the struct fields values.
	// If there is no current item under scan it returns -1.
	StructIn() (size int)

	// StructField selects current item i-th field value as an item to scan.
	// Note that StructIn() must be called before.
	// It panics if i is out of bounds.
	StructField(i int) (name string)

	// StructOut leaves struct entered before by StructIn() call.
	StructOut()

	// DictIn interprets current item under scan as a ydb's dict.
	// It returns the size of the nested items pairs.
	// If there is no current item under scan it returns -1.
	DictIn() (size int)

	// DictKey selects current item i-th pair key as an item to scan.
	// Note that DictIn() must be called before.
	// It panics if i is out of bounds.
	DictKey(i int)

	// DictPayload selects current item i-th pair value as an item to scan.
	// Note that DictIn() must be called before.
	// It panics if i is out of bounds.
	DictPayload(i int)

	// DictOut leaves dict entered before by DictIn() call.
	DictOut()

	// Variant unwraps current item under scan interpreting it as Variant<T> types.
	// It returns non-empty name of a field that is filled for struct-based
	// variant.
	// It always returns an index of filled field of a T.
	Variant() (name string, index uint32)

	// Decimal returns decimal value represented by big-endian 128 bit signed integer.
	Decimal(t Type) (v [16]byte)

	// UnwrapDecimal returns decimal value represented by big-endian 128 bit signed
	// integer and its types information.
	UnwrapDecimal() Decimal
	IsDecimal() bool
	Err() error
}

// Scanner scanning non-primitive yql types
type Scanner interface {
	UnmarshalYDB(res RawValue) error
}

type Decimal struct {
	Bytes     [16]byte
	Precision uint32
	Scale     uint32
}