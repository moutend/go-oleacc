package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationScrollPattern struct {
	ole.IUnknown
}

type IUIAutomationScrollPatternVtbl struct {
	ole.IUnknownVtbl
	Scroll                             uintptr
	SetScrollPercent                   uintptr
	Get_CurrentHorizontalScrollPercent uintptr
	Get_CurrentVerticalScrollPercent   uintptr
	Get_CurrentHorizontalViewSize      uintptr
	Get_CurrentVerticalViewSize        uintptr
	Get_CurrentHorizontallyScrollable  uintptr
	Get_CurrentVerticallyScrollable    uintptr
	Get_CachedHorizontalScrollPercent  uintptr
	Get_CachedVerticalScrollPercent    uintptr
	Get_CachedHorizontalViewSize       uintptr
	Get_CachedVerticalViewSize         uintptr
	Get_CachedHorizontallyScrollable   uintptr
	Get_CachedVerticallyScrollable     uintptr
}

func (v *IUIAutomationScrollPattern) VTable() *IUIAutomationScrollPatternVtbl {
	return (*IUIAutomationScrollPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationScrollPattern) Scroll() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationScrollPattern) SetScrollPercent() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationScrollPattern) Get_CurrentHorizontalScrollPercent() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationScrollPattern) Get_CurrentVerticalScrollPercent() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationScrollPattern) Get_CurrentHorizontalViewSize() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationScrollPattern) Get_CurrentVerticalViewSize() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationScrollPattern) Get_CurrentHorizontallyScrollable() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationScrollPattern) Get_CurrentVerticallyScrollable() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationScrollPattern) Get_CachedHorizontalScrollPercent() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationScrollPattern) Get_CachedVerticalScrollPercent() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationScrollPattern) Get_CachedHorizontalViewSize() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationScrollPattern) Get_CachedVerticalViewSize() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationScrollPattern) Get_CachedHorizontallyScrollable() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationScrollPattern) Get_CachedVerticallyScrollable() error {
	return ole.NewError(ole.E_NOTIMPL)
}
