package dataframe

import "reflect"

type Series struct {
	Name   string
	rawbuf []byte
	_type  reflect.Type
}

// Load stores the i-th element of the series into `addr`. `addr` must be a pointer.
func (s *Series) Load(addr interface{}, i int) {
}

// Store stores `value` to the i-th location of the series.
func (s *Series) Store(value interface{}, i int) {
}

// Append appends `value` to the series.
func (s *Series) Append(value interface{}) {
}
