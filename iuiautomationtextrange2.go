package oleacc

import (
	"unsafe"
)

type IUIAutomationTextRange2 struct {
	IUIAutomationTextRange
}

type IUIAutomationTextRange2Vtbl struct {
	IUIAutomationTextRangeVtbl
}

func (v *IUIAutomationTextRange2) VTable() *IUIAutomationTextRange2Vtbl {
	return (*IUIAutomationTextRange2Vtbl)(unsafe.Pointer(v.RawVTable))
}
