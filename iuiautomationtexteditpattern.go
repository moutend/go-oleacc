package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTextEditPattern struct {
	IUIAutomationTextPattern
}

type IUIAutomationTextEditPatternVtbl struct {
	IUIAutomationTextPatternVtbl
	GetActiveComposition uintptr
	GetConversionTarget  uintptr
}

func (v *IUIAutomationTextEditPattern) VTable() *IUIAutomationTextEditPatternVtbl {
	return (*IUIAutomationTextEditPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTextEditPattern) GetActiveComposition() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextEditPattern) GetConversionTarget() error {
	return ole.NewError(ole.E_NOTIMPL)
}
