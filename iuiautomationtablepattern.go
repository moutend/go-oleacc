package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTablePattern struct {
	ole.IUnknown
}

type IUIAutomationTablePatternVtbl struct {
	ole.IUnknownVtbl
	GetCurrentRowHeaders        uintptr
	GetCurrentColumnHeaders     uintptr
	Get_CurrentRowOrColumnMajor uintptr
	GetCachedRowHeaders         uintptr
	GetCachedColumnHeaders      uintptr
	Get_CachedRowOrColumnMajor  uintptr
}

func (v *IUIAutomationTablePattern) VTable() *IUIAutomationTablePatternVtbl {
	return (*IUIAutomationTablePatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTablePattern) GetCurrentRowHeaders() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTablePattern) GetCurrentColumnHeaders() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTablePattern) Get_CurrentRowOrColumnMajor() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTablePattern) GetCachedRowHeaders() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTablePattern) GetCachedColumnHeaders() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTablePattern) Get_CachedRowOrColumnMajor() error {
	return ole.NewError(ole.E_NOTIMPL)
}
