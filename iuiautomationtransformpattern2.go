package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTransformPattern2 struct {
	IUIAutomationTransformPattern
}

type IUIAutomationTransformPattern2Vtbl struct {
	IUIAutomationTransformPatternVtbl
	Zoom                   uintptr
	ZoomByUnit             uintptr
	Get_CurrentCanZoom     uintptr
	Get_CachedCanZoom      uintptr
	Get_CurrentZoomLevel   uintptr
	Get_CachedZoomLevel    uintptr
	Get_CurrentZoomMinimum uintptr
	Get_CachedZoomMinimum  uintptr
	Get_CurrentZoomMaximum uintptr
	Get_CachedZoomMaximum  uintptr
}

func (v *IUIAutomationTransformPattern2) VTable() *IUIAutomationTransformPattern2Vtbl {
	return (*IUIAutomationTransformPattern2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTransformPattern2) Zoom() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern2) ZoomByUnit() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern2) Get_CurrentCanZoom() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern2) Get_CachedCanZoom() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern2) Get_CurrentZoomLevel() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern2) Get_CachedZoomLevel() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern2) Get_CurrentZoomMinimum() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern2) Get_CachedZoomMinimum() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern2) Get_CurrentZoomMaximum() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTransformPattern2) Get_CachedZoomMaximum() error {
	return ole.NewError(ole.E_NOTIMPL)
}
