package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTextPattern struct {
	ole.IUnknown
}

type IUIAutomationTextPatternVtbl struct {
	ole.IUnknownVtbl
	RangeFromPoint             uintptr
	RangeFromChild             uintptr
	GetSelection               uintptr
	GetVisibleRanges           uintptr
	Get_DocumentRange          uintptr
	Get_SupportedTextSelection uintptr
}

func (v *IUIAutomationTextPattern) VTable() *IUIAutomationTextPatternVtbl {
	return (*IUIAutomationTextPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTextPattern) RangeFromPoint() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextPattern) RangeFromChild() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextPattern) GetSelection() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextPattern) GetVisibleRanges() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextPattern) Get_DocumentRange() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextPattern) Get_SupportedTextSelection() error {
	return ole.NewError(ole.E_NOTIMPL)
}
