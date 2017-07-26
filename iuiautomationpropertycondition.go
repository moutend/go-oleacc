package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationPropertyCondition struct {
	IUIAutomationCondition
}

type IUIAutomationPropertyConditionVtbl struct {
	IUIAutomationConditionVtbl
	Get_PropertyId             uintptr
	Get_PropertyValue          uintptr
	Get_PropertyConditionFlags uintptr
}

func (v *IUIAutomationPropertyCondition) VTable() *IUIAutomationPropertyConditionVtbl {
	return (*IUIAutomationPropertyConditionVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationPropertyCondition) Get_PropertyId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationPropertyCondition) Get_PropertyValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationPropertyCondition) Get_PropertyConditionFlags() error {
	return ole.NewError(ole.E_NOTIMPL)
}
