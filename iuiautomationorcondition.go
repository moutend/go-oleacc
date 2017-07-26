package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationOrCondition struct {
	IUIAutomationCondition
}

type IUIAutomationOrConditionVtbl struct {
	IUIAutomationConditionVtbl
	Get_ChildCount           uintptr
	GetChildrenAsNativeArray uintptr
	GetChildren              uintptr
}

func (v *IUIAutomationOrCondition) VTable() *IUIAutomationOrConditionVtbl {
	return (*IUIAutomationOrConditionVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationOrCondition) Get_ChildCount() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationOrCondition) GetChildrenAsNativeArray() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationOrCondition) GetChildren() error {
	return ole.NewError(ole.E_NOTIMPL)
}
