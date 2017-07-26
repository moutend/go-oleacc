package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationAnnotationPattern struct {
	ole.IUnknown
}

type IUIAutomationAnnotationPatternVtbl struct {
	ole.IUnknownVtbl
	Get_CurrentAnnotationTypeId   uintptr
	Get_CurrentAnnotationTypeName uintptr
	Get_CurrentAuthor             uintptr
	Get_CurrentDateTime           uintptr
	Get_CurrentTarget             uintptr
	Get_CachedAnnotationTypeId    uintptr
	Get_CachedAnnotationTypeName  uintptr
	Get_CachedAuthor              uintptr
	Get_CachedDateTime            uintptr
	Get_CachedTarget              uintptr
}

func (v *IUIAutomationAnnotationPattern) VTable() *IUIAutomationAnnotationPatternVtbl {
	return (*IUIAutomationAnnotationPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationAnnotationPattern) Get_CurrentAnnotationTypeId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationAnnotationPattern) Get_CurrentAnnotationTypeName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationAnnotationPattern) Get_CurrentAuthor() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationAnnotationPattern) Get_CurrentDateTime() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationAnnotationPattern) Get_CurrentTarget() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationAnnotationPattern) Get_CachedAnnotationTypeId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationAnnotationPattern) Get_CachedAnnotationTypeName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationAnnotationPattern) Get_CachedAuthor() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationAnnotationPattern) Get_CachedDateTime() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationAnnotationPattern) Get_CachedTarget() error {
	return ole.NewError(ole.E_NOTIMPL)
}
