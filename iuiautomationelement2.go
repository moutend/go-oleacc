package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElement2 struct {
	IUIAutomationElement
}

type IUIAutomationElement2Vtbl struct {
	IUIAutomationElementVtbl
	Get_CurrentOptimizeForVisualContent uintptr
	Get_CachedOptimizeForVisualContent  uintptr
	Get_CurrentLiveSetting              uintptr
	Get_CachedLiveSetting               uintptr
	Get_CurrentFlowsFrom                uintptr
	Get_CachedFlowsFrom                 uintptr
}

func (v *IUIAutomationElement2) VTable() *IUIAutomationElement2Vtbl {
	return (*IUIAutomationElement2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationElement2) Get_CurrentOptimizeForVisualContent() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement2) Get_CachedOptimizeForVisualContent() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement2) Get_CurrentLiveSetting() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement2) Get_CachedLiveSetting() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement2) Get_CurrentFlowsFrom() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement2) Get_CachedFlowsFrom() error {
	return ole.NewError(ole.E_NOTIMPL)
}
