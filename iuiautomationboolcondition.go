package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationBoolCondition struct {
	IUIAutomationCondition
}

type IUIAutomationBoolConditionVtbl struct {
	IUIAutomationConditionVtbl
	Get_BooleanValue uintptr
}

func (v *IUIAutomationBoolCondition) VTable() *IUIAutomationBoolConditionVtbl {
	return (*IUIAutomationBoolConditionVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationBoolCondition) Get_BooleanValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}
