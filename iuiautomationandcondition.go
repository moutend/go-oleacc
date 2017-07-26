package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationAndCondition struct {
	IUIAutomationCondition
}

type IUIAutomationAndConditionVtbl struct {
	IUIAutomationConditionVtbl
	Get_ChildCount           uintptr
	GetChildrenAsNativeArray uintptr
	GetChildren              uintptr
}

func (v *IUIAutomationAndCondition) VTable() *IUIAutomationAndConditionVtbl {
	return (*IUIAutomationAndConditionVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationAndCondition) Get_ChildCount() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationAndCondition) GetChildrenAsNativeArray() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationAndCondition) GetChildren() error {
	return ole.NewError(ole.E_NOTIMPL)
}
