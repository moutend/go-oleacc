package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationSelectionPattern struct {
	ole.IUnknown
}

type IUIAutomationSelectionPatternVtbl struct {
	ole.IUnknownVtbl
	GetCurrentSelection            uintptr
	Get_CurrentCanSelectMultiple   uintptr
	Get_CurrentIsSelectionRequired uintptr
	GetCachedSelection             uintptr
	Get_CachedCanSelectMultiple    uintptr
	Get_CachedIsSelectionRequired  uintptr
}

func (v *IUIAutomationSelectionPattern) VTable() *IUIAutomationSelectionPatternVtbl {
	return (*IUIAutomationSelectionPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationSelectionPattern) GetCurrentSelection() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationSelectionPattern) Get_CurrentCanSelectMultiple() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationSelectionPattern) Get_CurrentIsSelectionRequired() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationSelectionPattern) GetCachedSelection() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationSelectionPattern) Get_CachedCanSelectMultiple() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationSelectionPattern) Get_CachedIsSelectionRequired() error {
	return ole.NewError(ole.E_NOTIMPL)
}
