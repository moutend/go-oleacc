package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationMultipleViewPattern struct {
	ole.IUnknown
}

type IUIAutomationMultipleViewPatternVtbl struct {
	ole.IUnknownVtbl
	GetViewName              uintptr
	SetCurrentView           uintptr
	Get_CurrentCurrentView   uintptr
	GetCurrentSupportedViews uintptr
	Get_CachedCurrentView    uintptr
	GetCachedSupportedViews  uintptr
}

func (v *IUIAutomationMultipleViewPattern) VTable() *IUIAutomationMultipleViewPatternVtbl {
	return (*IUIAutomationMultipleViewPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationMultipleViewPattern) GetViewName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationMultipleViewPattern) SetCurrentView() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationMultipleViewPattern) Get_CurrentCurrentView() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationMultipleViewPattern) GetCurrentSupportedViews() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationMultipleViewPattern) Get_CachedCurrentView() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationMultipleViewPattern) GetCachedSupportedViews() error {
	return ole.NewError(ole.E_NOTIMPL)
}
