// +build windows

package oleacc

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

var (
	modOleacc, _ = syscall.LoadDLL("oleacc.dll")

	procAccessibleObjectFromEvent, _  = modOleacc.FindProc("AccessibleObjectFromEvent")
	procAccessibleObjectFromPoint, _  = modOleacc.FindProc("AccessibleObjectFromPoint")
	procWindowFromAccessibleObject, _ = modOleacc.FindProc("WindowFromAccessibleObject")
)

func accessibleObjectFromEvent(hWindow uintptr, objectId, childId uint32, acc **IAccessible, varChild **ole.VARIANT) (err error) {
	hr, _, _ := procAccessibleObjectFromEvent.Call(
		hWindow,
		uintptr(objectId),
		uintptr(childId),
		uintptr(unsafe.Pointer(acc)),
		uintptr(unsafe.Pointer(varChild)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func accessibleObjectFromPoint(point Point, acc **IAccessible, varChild **ole.VARIANT) (err error) {
	hr, _, _ := procAccessibleObjectFromPoint.Call(
		uintptr(unsafe.Pointer(&point)),
		uintptr(unsafe.Pointer(acc)),
		uintptr(unsafe.Pointer(varChild)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func windowFromAccessibleObject(acc *IAccessible, hWindow uintptr) (err error) {
	hr, _, _ := procWindowFromAccessibleObject.Call(
		uintptr(unsafe.Pointer(acc)),
		hWindow)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
