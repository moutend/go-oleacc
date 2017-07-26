package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationInvokePattern struct {
	ole.IUnknown
}

type IUIAutomationInvokePatternVtbl struct {
	ole.IUnknownVtbl
	Invoke uintptr
}

func (v *IUIAutomationInvokePattern) VTable() *IUIAutomationInvokePatternVtbl {
	return (*IUIAutomationInvokePatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationInvokePattern) Invoke() error {
	return ole.NewError(ole.E_NOTIMPL)
}
