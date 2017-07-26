package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomation4 struct {
	IUIAutomation3
}

type IUIAutomation4Vtbl struct {
	IUIAutomation3Vtbl
	AddChangesEventHandler    uintptr
	RemoveChangesEventHandler uintptr
}

func (v *IUIAutomation4) VTable() *IUIAutomation4Vtbl {
	return (*IUIAutomation4Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomation4) AddChangesEventHandler() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation4) RemoveChangesEventHandler() error {
	return ole.NewError(ole.E_NOTIMPL)
}
