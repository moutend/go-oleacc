package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationSelectionItemPattern struct {
	ole.IUnknown
}

type IUIAutomationSelectionItemPatternVtbl struct {
	ole.IUnknownVtbl
	Get_CurrentIsSelected         uintptr
	Get_CurrentSelectionContainer uintptr
	Get_CachedIsSelected          uintptr
	Get_CachedSelectionContainer  uintptr
}

func (v *IUIAutomationSelectionItemPattern) VTable() *IUIAutomationSelectionItemPatternVtbl {
	return (*IUIAutomationSelectionItemPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationSelectionItemPattern) Get_CurrentIsSelected() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationSelectionItemPattern) Get_CurrentSelectionContainer() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationSelectionItemPattern) Get_CachedIsSelected() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationSelectionItemPattern) Get_CachedSelectionContainer() error {
	return ole.NewError(ole.E_NOTIMPL)
}
