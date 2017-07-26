package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTextPattern2 struct {
	IUIAutomationTextPattern
}

type IUIAutomationTextPattern2Vtbl struct {
	IUIAutomationTextPatternVtbl
	RangeFromAnnotation uintptr
	GetCaretRange       uintptr
}

func (v *IUIAutomationTextPattern2) VTable() *IUIAutomationTextPattern2Vtbl {
	return (*IUIAutomationTextPattern2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTextPattern2) RangeFromAnnotation() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextPattern2) GetCaretRange() error {
	return ole.NewError(ole.E_NOTIMPL)
}
