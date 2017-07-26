package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationValuePattern struct {
	ole.IUnknown
}

type IUIAutomationValuePatternVtbl struct {
	ole.IUnknownVtbl
	SetValue              uintptr
	Get_CurrentValue      uintptr
	Get_CurrentIsReadOnly uintptr
	Get_CachedValue       uintptr
	Get_CachedIsReadOnly  uintptr
}

func (v *IUIAutomationValuePattern) VTable() *IUIAutomationValuePatternVtbl {
	return (*IUIAutomationValuePatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationValuePattern) SetValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationValuePattern) Get_CurrentValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationValuePattern) Get_CurrentIsReadOnly() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationValuePattern) Get_CachedValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationValuePattern) Get_CachedIsReadOnly() error {
	return ole.NewError(ole.E_NOTIMPL)
}
