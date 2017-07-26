package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationDropTargetPattern struct {
	ole.IUnknown
}

type IUIAutomationDropTargetPatternVtbl struct {
	ole.IUnknownVtbl
	Get_CurrentDropTargetEffect  uintptr
	Get_CachedDropTargetEffect   uintptr
	Get_CurrentDropTargetEffects uintptr
	Get_CachedDropTargetEffects  uintptr
}

func (v *IUIAutomationDropTargetPattern) VTable() *IUIAutomationDropTargetPatternVtbl {
	return (*IUIAutomationDropTargetPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationDropTargetPattern) Get_CurrentDropTargetEffect() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationDropTargetPattern) Get_CachedDropTargetEffect() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationDropTargetPattern) Get_CurrentDropTargetEffects() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationDropTargetPattern) Get_CachedDropTargetEffects() error {
	return ole.NewError(ole.E_NOTIMPL)
}
