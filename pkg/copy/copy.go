package copy

import "github.com/jinzhu/copier"

// MustCopy copies the src into dst and panics if an error occurs.
func MustCopy[Dst any, Src any](dst Dst, src Src) Dst {
	if err := copier.Copy(dst, src); err != nil {
		panic(err)
	}

	return dst
}
