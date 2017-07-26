package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationGridPattern struct {
	ole.IUnknown
}

type IUIAutomationGridPatternVtbl struct {
	ole.IUnknownVtbl
	GetItem                uintptr
	Get_CurrentRowCount    uintptr
	Get_CurrentColumnCount uintptr
	Get_CachedRowCount     uintptr
	Get_CachedColumnCount  uintptr
}

func (v *IUIAutomationGridPattern) VTable() *IUIAutomationGridPatternVtbl {
	return (*IUIAutomationGridPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationGridPattern) GetItem() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationGridPattern) Get_CurrentRowCount() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationGridPattern) Get_CurrentColumnCount() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationGridPattern) Get_CachedRowCount() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationGridPattern) Get_CachedColumnCount() error {
	return ole.NewError(ole.E_NOTIMPL)
}
