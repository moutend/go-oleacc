package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTransformPattern struct {
	ole.IUnknown
}

type IUIAutomationTransformPatternVtbl struct {
	ole.IUnknownVtbl
	Move                 uintptr
	Resize               uintptr
	Rotate               uintptr
	Get_CurrentCanMove   uintptr
	Get_CurrentCanResize uintptr
	Get_CurrentCanRotate uintptr
	Get_CachedCanMove    uintptr
	Get_CachedCanResize  uintptr
	Get_CachedCanRotate  uintptr
}

func (v *IUIAutomationTransformPattern) VTable() *IUIAutomationTransformPatternVtbl {
	return (*IUIAutomationTransformPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTransformPattern) Move() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern) Resize() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern) Rotate() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern) Get_CurrentCanMove() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern) Get_CurrentCanResize() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern) Get_CurrentCanRotate() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern) Get_CachedCanMove() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern) Get_CachedCanResize() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern) Get_CachedCanRotate() error {
	return ole.NewError(ole.E_NOTIMPL)
}
