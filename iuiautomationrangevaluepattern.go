package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationRangeValuePattern struct {
	ole.IUnknown
}

type IUIAutomationRangeValuePatternVtbl struct {
	ole.IUnknownVtbl
	SetValue               uintptr
	Get_CurrentValue       uintptr
	Get_CurrentIsReadOnly  uintptr
	Get_CurrentMaximum     uintptr
	Get_CurrentMinimum     uintptr
	Get_CurrentLargeChange uintptr
	Get_CurrentSmallChange uintptr
	Get_CachedValue        uintptr
	Get_CachedIsReadOnly   uintptr
	Get_CachedMaximum      uintptr
	Get_CachedMinimum      uintptr
	Get_CachedLargeChange  uintptr
	Get_CachedSmallChange  uintptr
}

func (v *IUIAutomationRangeValuePattern) VTable() *IUIAutomationRangeValuePatternVtbl {
	return (*IUIAutomationRangeValuePatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationRangeValuePattern) SetValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationRangeValuePattern) Get_CurrentValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationRangeValuePattern) Get_CurrentIsReadOnly() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationRangeValuePattern) Get_CurrentMaximum() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationRangeValuePattern) Get_CurrentMinimum() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationRangeValuePattern) Get_CurrentLargeChange() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationRangeValuePattern) Get_CurrentSmallChange() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationRangeValuePattern) Get_CachedValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationRangeValuePattern) Get_CachedIsReadOnly() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationRangeValuePattern) Get_CachedMaximum() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationRangeValuePattern) Get_CachedMinimum() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationRangeValuePattern) Get_CachedLargeChange() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationRangeValuePattern) Get_CachedSmallChange() error {
	return ole.NewError(ole.E_NOTIMPL)
}
