package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationWindowPattern struct {
	ole.IUnknown
}

type IUIAutomationWindowPatternVtbl struct {
	ole.IUnknownVtbl
	WaitForInputIdle                  uintptr
	SetWindowVisualState              uintptr
	Get_CurrentCanMaximize            uintptr
	Get_CurrentCanMinimize            uintptr
	Get_CurrentIsModal                uintptr
	Get_CurrentIsTopmost              uintptr
	Get_CurrentWindowVisualState      uintptr
	Get_CurrentWindowInteractionState uintptr
	Get_CachedCanMaximize             uintptr
	Get_CachedCanMinimize             uintptr
	Get_CachedIsModal                 uintptr
	Get_CachedIsTopmost               uintptr
	Get_CachedWindowVisualState       uintptr
	Get_CachedWindowInteractionState  uintptr
}

func (v *IUIAutomationWindowPattern) VTable() *IUIAutomationWindowPatternVtbl {
	return (*IUIAutomationWindowPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationWindowPattern) WaitForInputIdle() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationWindowPattern) SetWindowVisualState() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationWindowPattern) Get_CurrentCanMaximize() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationWindowPattern) Get_CurrentCanMinimize() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationWindowPattern) Get_CurrentIsModal() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationWindowPattern) Get_CurrentIsTopmost() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationWindowPattern) Get_CurrentWindowVisualState() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationWindowPattern) Get_CurrentWindowInteractionState() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationWindowPattern) Get_CachedCanMaximize() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationWindowPattern) Get_CachedCanMinimize() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationWindowPattern) Get_CachedIsModal() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationWindowPattern) Get_CachedIsTopmost() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationWindowPattern) Get_CachedWindowVisualState() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationWindowPattern) Get_CachedWindowInteractionState() error {
	return ole.NewError(ole.E_NOTIMPL)
}
