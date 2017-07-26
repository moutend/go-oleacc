package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationGridItemPattern struct {
	ole.IUnknown
}

type IUIAutomationGridItemPatternVtbl struct {
	ole.IUnknownVtbl
	Get_CurrentContainingGrid uintptr
	Get_CurrentRow            uintptr
	Get_CurrentColumn         uintptr
	Get_CurrentRowSpan        uintptr
	Get_CurrentColumnSpan     uintptr
	Get_CachedContainingGrid  uintptr
	Get_CachedRow             uintptr
	Get_CachedColumn          uintptr
	Get_CachedRowSpan         uintptr
	Get_CachedColumnSpan      uintptr
}

func (v *IUIAutomationGridItemPattern) VTable() *IUIAutomationGridItemPatternVtbl {
	return (*IUIAutomationGridItemPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationGridItemPattern) Get_CurrentContainingGrid() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationGridItemPattern) Get_CurrentRow() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationGridItemPattern) Get_CurrentColumn() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationGridItemPattern) Get_CurrentRowSpan() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationGridItemPattern) Get_CurrentColumnSpan() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationGridItemPattern) Get_CachedContainingGrid() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationGridItemPattern) Get_CachedRow() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationGridItemPattern) Get_CachedColumn() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationGridItemPattern) Get_CachedRowSpan() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationGridItemPattern) Get_CachedColumnSpan() error {
	return ole.NewError(ole.E_NOTIMPL)
}
