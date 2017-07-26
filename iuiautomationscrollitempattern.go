package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationScrollItemPattern struct {
	ole.IUnknown
}

type IUIAutomationScrollItemPatternVtbl struct {
	ole.IUnknownVtbl
}

func (v *IUIAutomationScrollItemPattern) VTable() *IUIAutomationScrollItemPatternVtbl {
	return (*IUIAutomationScrollItemPatternVtbl)(unsafe.Pointer(v.RawVTable))
}
