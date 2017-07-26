package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTogglePattern struct {
	ole.IUnknown
}

type IUIAutomationTogglePatternVtbl struct {
	ole.IUnknownVtbl
	Get_CurrentToggleState uintptr
	Get_CachedToggleState  uintptr
}

func (v *IUIAutomationTogglePattern) VTable() *IUIAutomationTogglePatternVtbl {
	return (*IUIAutomationTogglePatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTogglePattern) Get_CurrentToggleState() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTogglePattern) Get_CachedToggleState() error {
	return ole.NewError(ole.E_NOTIMPL)
}
