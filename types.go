package oleacc

import (
	"github.com/go-ole/go-ole"
)

type Point struct {
	X int32
	Y int32
}

type UiaChangeInfo struct {
	uiaId     int32
	payload   ole.VARIANT
	extraInfo ole.VARIANT
}
