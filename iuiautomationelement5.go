package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElement5 struct {
	IUIAutomationElement4
}

type IUIAutomationElement5Vtbl struct {
	IUIAutomationElement4Vtbl
	Get_CurrentLandmarkType          uintptr
	Get_CurrentLocalizedLandmarkType uintptr
	Get_CachedLandmarkType           uintptr
	Get_CachedLocalizedLandmarkType  uintptr
}

func (v *IUIAutomationElement5) VTable() *IUIAutomationElement5Vtbl {
	return (*IUIAutomationElement5Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationElement5) Get_CurrentLandmarkType() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement5) Get_CurrentLocalizedLandmarkType() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement5) Get_CachedLandmarkType() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement5) Get_CachedLocalizedLandmarkType() error {
	return ole.NewError(ole.E_NOTIMPL)
}
