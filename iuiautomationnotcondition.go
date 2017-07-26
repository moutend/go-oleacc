package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationNotCondition struct {
	IUIAutomationCondition
}

type IUIAutomationNotConditionVtbl struct {
	IUIAutomationConditionVtbl
	GetChild uintptr
}

func (v *IUIAutomationNotCondition) VTable() *IUIAutomationNotConditionVtbl {
	return (*IUIAutomationNotConditionVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationNotCondition) GetChild() error {
	return ole.NewError(ole.E_NOTIMPL)
}
