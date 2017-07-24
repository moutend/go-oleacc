// +build !windows

package waa

import (
	"github.com/go-ole/go-ole"
)

func accessibleObjectFromEvent(hWindow uintptr, objectId, childId uint32, acc **IAccessible, varChild **ole.VARIANT) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accessibleObjectFromPoint(point Point, acc **IAccessible, varChild **ole.VARIANT) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
