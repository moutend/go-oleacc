package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElement3 struct {
	IUIAutomationElement2
}

type IUIAutomationElement3Vtbl struct {
	IUIAutomationElement2Vtbl
	Get_CurrentIsPeripheral uintptr
	Get_CachedIsPeripheral  uintptr
}

func (v *IUIAutomationElement3) VTable() *IUIAutomationElement3Vtbl {
	return (*IUIAutomationElement3Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationElement3) Get_CurrentIsPeripheral() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement3) Get_CachedIsPeripheral() error {
	return ole.NewError(ole.E_NOTIMPL)
}
