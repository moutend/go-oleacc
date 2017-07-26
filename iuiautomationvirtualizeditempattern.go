package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationVirtualizedItemPattern struct {
	ole.IUnknown
}

type IUIAutomationVirtualizedItemPatternVtbl struct {
	ole.IUnknownVtbl
}

func (v *IUIAutomationVirtualizedItemPattern) VTable() *IUIAutomationVirtualizedItemPatternVtbl {
	return (*IUIAutomationVirtualizedItemPatternVtbl)(unsafe.Pointer(v.RawVTable))
}
