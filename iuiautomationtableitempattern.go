package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTableItemPattern struct {
	ole.IUnknown
}

type IUIAutomationTableItemPatternVtbl struct {
	ole.IUnknownVtbl
	GetCurrentRowHeaderItems    uintptr
	GetCurrentColumnHeaderItems uintptr
	GetCachedRowHeaderItems     uintptr
	GetCachedColumnHeaderItems  uintptr
}

func (v *IUIAutomationTableItemPattern) VTable() *IUIAutomationTableItemPatternVtbl {
	return (*IUIAutomationTableItemPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTableItemPattern) GetCurrentRowHeaderItems() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTableItemPattern) GetCurrentColumnHeaderItems() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTableItemPattern) GetCachedRowHeaderItems() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTableItemPattern) GetCachedColumnHeaderItems() error {
	return ole.NewError(ole.E_NOTIMPL)
}
