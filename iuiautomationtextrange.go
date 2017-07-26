package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTextRange struct {
	ole.IUnknown
}

type IUIAutomationTextRangeVtbl struct {
	ole.IUnknownVtbl
	Clone                 uintptr
	Compare               uintptr
	CompareEndpoints      uintptr
	ExpandToEnclosingUnit uintptr
	FindAttribute         uintptr
	FindText              uintptr
	GetAttributeValue     uintptr
	GetBoundingRectangles uintptr
	GetEnclosingElement   uintptr
	GetText               uintptr
	Move                  uintptr
	MoveEndpointByUnit    uintptr
	MoveEndpointByRange   uintptr
	ScrollIntoView        uintptr
	GetChildren           uintptr
}

func (v *IUIAutomationTextRange) VTable() *IUIAutomationTextRangeVtbl {
	return (*IUIAutomationTextRangeVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTextRange) Clone() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) Compare() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) CompareEndpoints() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) ExpandToEnclosingUnit() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) FindAttribute() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) FindText() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) GetAttributeValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) GetBoundingRectangles() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) GetEnclosingElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) GetText() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) Move() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) MoveEndpointByUnit() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) MoveEndpointByRange() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) ScrollIntoView() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange) GetChildren() error {
	return ole.NewError(ole.E_NOTIMPL)
}
