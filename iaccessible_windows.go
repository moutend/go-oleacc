// +build windows

package waa

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func accNavigate(acc *IAccessible, navdir uint32, varStart *ole.VARIANT, varEndUpAt **ole.VARIANT) (err error) {
	hr, _, _ := syscall.Syscall6(
		acc.VTable().AccNavigate,
		4,
		uintptr(unsafe.Pointer(acc)),
		uintptr(navdir),
		uintptr(unsafe.Pointer(varStart)),
		uintptr(unsafe.Pointer(varEndUpAt)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
