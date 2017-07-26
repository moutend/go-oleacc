package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationDragPattern struct {
	ole.IUnknown
}

type IUIAutomationDragPatternVtbl struct {
	ole.IUnknownVtbl
	Get_CurrentIsGrabbed   uintptr
	Get_CachedIsGrabbed    uintptr
	Get_CurrentDropEffect  uintptr
	Get_CachedDropEffect   uintptr
	Get_CurrentDropEffects uintptr
	Get_CachedDropEffects  uintptr
	GetCurrentGrabbedItems uintptr
	GetCachedGrabbedItems  uintptr
}

func (v *IUIAutomationDragPattern) VTable() *IUIAutomationDragPatternVtbl {
	return (*IUIAutomationDragPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationDragPattern) Get_CurrentIsGrabbed() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationDragPattern) Get_CachedIsGrabbed() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationDragPattern) Get_CurrentDropEffect() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationDragPattern) Get_CachedDropEffect() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationDragPattern) Get_CurrentDropEffects() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationDragPattern) Get_CachedDropEffects() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationDragPattern) GetCurrentGrabbedItems() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationDragPattern) GetCachedGrabbedItems() error {
	return ole.NewError(ole.E_NOTIMPL)
}
