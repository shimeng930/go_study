package common_util

var PointerConverter pointerConverter

type pointerConverter struct {

}


func (r *pointerConverter) Uint32(v uint32) *uint32 {
	return &v
}

func (r *pointerConverter) Int64(v int64) *int64 {
	return &v
}

func (r *pointerConverter) Uint64(v uint64) *uint64 {
	return &v
}