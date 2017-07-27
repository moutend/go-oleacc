// +build windows

package oleacc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

var (
	modOleacc, _ = syscall.LoadDLL("oleacc.dll")
	modUser32, _ = syscall.LoadDLL("user32.dll")

	procOpenInputDesktop, _ = modUser32.FindProc("OpenInputDesktop")
	procCloseDesktop, _     = modUser32.FindProc("CloseDesktop ")

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
	var p uint64

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, point)
	binary.Read(buf, binary.LittleEndian, &p)
	fmt.Printf("ptr: %x\n", p)
	hr, _, _ := procAccessibleObjectFromPoint.Call(
		uintptr(p),
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

func openInputDesktop(flags uint32, isInherit bool, accessMask uint32) HDESK {
	var isInheritValue uint32
	if isInherit {
		isInheritValue = 1
	}
	r, _, _ := procOpenInputDesktop.Call(
		uintptr(flags),
		uintptr(isInheritValue),
		uintptr(accessMask))
	return HDESK(r)
}

func closeDesktop(hDesk HDESK) bool {
	r, _, _ := procCloseDesktop.Call(uintptr(hDesk))
	if r == 0 {
		return false
	} else {
		return true
	}
}
