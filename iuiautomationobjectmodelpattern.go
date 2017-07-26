package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationObjectModelPattern struct {
	ole.IUnknown
}

type IUIAutomationObjectModelPatternVtbl struct {
	ole.IUnknownVtbl
	GetUnderlyingObjectModel uintptr
}

func (v *IUIAutomationObjectModelPattern) VTable() *IUIAutomationObjectModelPatternVtbl {
	return (*IUIAutomationObjectModelPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationObjectModelPattern) GetUnderlyingObjectModel() error {
	return ole.NewError(ole.E_NOTIMPL)
}
