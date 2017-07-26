package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElement4 struct {
	IUIAutomationElement3
}

type IUIAutomationElement4Vtbl struct {
	IUIAutomationElement3Vtbl
	Get_CurrentPositionInSet     uintptr
	Get_CurrentSizeOfSet         uintptr
	Get_CurrentLevel             uintptr
	Get_CurrentAnnotationTypes   uintptr
	Get_CurrentAnnotationObjects uintptr
	Get_CachedPositionInSet      uintptr
	Get_CachedSizeOfSet          uintptr
	Get_CachedLevel              uintptr
	Get_CachedAnnotationTypes    uintptr
	Get_CachedAnnotationObjects  uintptr
}

func (v *IUIAutomationElement4) VTable() *IUIAutomationElement4Vtbl {
	return (*IUIAutomationElement4Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationElement4) Get_CurrentPositionInSet() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement4) Get_CurrentSizeOfSet() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement4) Get_CurrentLevel() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement4) Get_CurrentAnnotationTypes() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement4) Get_CurrentAnnotationObjects() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement4) Get_CachedPositionInSet() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement4) Get_CachedSizeOfSet() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement4) Get_CachedLevel() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement4) Get_CachedAnnotationTypes() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement4) Get_CachedAnnotationObjects() error {
	return ole.NewError(ole.E_NOTIMPL)
}
