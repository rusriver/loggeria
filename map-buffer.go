package loggeria

import (
	"fmt"
	"time"
)

// Can be used on its own, and a "chaining API" DSL, to build data messages directly.
type IMapBuffer interface {
	GetBytes() []byte
	AddFinalizer(func()) // added func will be called by Finalize(), just before closing the object
	Finalize()           // the finalizer
	//
	List(k string, list IListBuffer) IMapBuffer // calls Finalize() on nested object automatically
	Map(m IMapBuffer) IMapBuffer                // calls Finalize() on nested object automatically
	//
	// special and shortcuts
	Caller(skip ...int) IMapBuffer
	ErrWithStacktrace(k string, err error) IMapBuffer
	Fields(fields interface{}) IMapBuffer
	Msgid(id string) IMapBuffer
	Msg(msg string) IMapBuffer
	MsgFunc(f func() string) IMapBuffer
	Msgf(format string, v ...interface{}) IMapBuffer
	RawCBOR(k string, v []byte) IMapBuffer
	RawJSON(k string, v []byte) IMapBuffer
	StrFunc(k string, f func() string) IMapBuffer
	TimeDiff(k string, t1 time.Time, t2 time.Time) IMapBuffer
	Type(k string, v interface{}) IMapBuffer
	//
	// 18 common types
	Bool(b bool) IMapBuffer
	Dur(d time.Duration) IMapBuffer
	Err(k string, err error) IMapBuffer
	Float32(f float32) IMapBuffer
	Float64(f float64) IMapBuffer
	Int(i int) IMapBuffer
	Int16(i int16) IMapBuffer
	Int32(i int32) IMapBuffer
	Int64(i int64) IMapBuffer
	Int8(i int8) IMapBuffer
	Str(k string, v string) IMapBuffer
	Stringer(k string, v fmt.Stringer) IMapBuffer
	Time(t time.Time) IMapBuffer
	Uint(i uint) IMapBuffer
	Uint16(i uint16) IMapBuffer
	Uint32(i uint32) IMapBuffer
	Uint64(i uint64) IMapBuffer
	Uint8(i uint8) IMapBuffer
	//
	// 18 homogeneous lists
	Bools(k string, b []bool) IMapBuffer
	Durs(k string, d []time.Duration) IMapBuffer
	Errs(k string, errs []error) IMapBuffer
	Floats32(k string, f []float32) IMapBuffer
	Floats64(k string, f []float64) IMapBuffer
	Ints(k string, i []int) IMapBuffer
	Ints16(k string, i []int16) IMapBuffer
	Ints32(k string, i []int32) IMapBuffer
	Ints64(k string, i []int64) IMapBuffer
	Ints8(k string, i []int8) IMapBuffer
	Stringers(k string, vals []fmt.Stringer) IMapBuffer
	Strs(k string, vals []string) IMapBuffer
	Times(k string, t []time.Time) IMapBuffer
	Uints(k string, i []uint) IMapBuffer
	Uints16(k string, i []uint16) IMapBuffer
	Uints32(k string, i []uint32) IMapBuffer
	Uints64(k string, i []uint64) IMapBuffer
	Uints8(k string, i []uint8) IMapBuffer
}
