package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTextRangeArray struct {
	ole.IUnknown
}

type IUIAutomationTextRangeArrayVtbl struct {
	ole.IUnknownVtbl
	Get_Length uintptr
	GetElement uintptr
}

func (v *IUIAutomationTextRangeArray) VTable() *IUIAutomationTextRangeArrayVtbl {
	return (*IUIAutomationTextRangeArrayVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTextRangeArray) Get_Length() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRangeArray) GetElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}
