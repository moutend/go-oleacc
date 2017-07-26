package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomation3 struct {
	IUIAutomation2
}

type IUIAutomation3Vtbl struct {
	IUIAutomation2Vtbl
	AddTextEditTextChangedEventHandler    uintptr
	RemoveTextEditTextChangedEventHandler uintptr
}

func (v *IUIAutomation3) VTable() *IUIAutomation3Vtbl {
	return (*IUIAutomation3Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomation3) AddTextEditTextChangedEventHandler() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation3) RemoveTextEditTextChangedEventHandler() error {
	return ole.NewError(ole.E_NOTIMPL)
}
