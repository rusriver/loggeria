package loggeria

import (
	"fmt"
	"time"
)

// Heterogeneous list
type IListBuffer interface {
	GetBytes() []byte
	AddFinalizer(func()) // added func will be called by Finalize(), just before closing the object
	Finalize()           // the finalizer; Finalize() must be called explicitly, to "close" the object. It must be idempotent.
	//
	List(list IListBuffer) IListBuffer // calls Finalize() on nested object automatically
	Map(m IMapBuffer) IListBuffer      // calls Finalize() on nested object automatically
	//
	StrFunc(f func() string) IListBuffer
	RawCBOR(v []byte) IMapBuffer
	RawJSON(v []byte) IMapBuffer
	//
	// 18 common types
	Bool(b bool) IListBuffer
	Dur(d time.Duration) IListBuffer
	Err(err error) IListBuffer
	Float32(f float32) IListBuffer
	Float64(f float64) IListBuffer
	Int(i int) IListBuffer
	Int16(i int16) IListBuffer
	Int32(i int32) IListBuffer
	Int64(i int64) IListBuffer
	Int8(i int8) IListBuffer
	Str(v string) IListBuffer
	Stringer(v fmt.Stringer) IMapBuffer
	Time(t time.Time) IListBuffer
	Uint(i uint) IListBuffer
	Uint16(i uint16) IListBuffer
	Uint32(i uint32) IListBuffer
	Uint64(i uint64) IListBuffer
	Uint8(i uint8) IListBuffer
}
