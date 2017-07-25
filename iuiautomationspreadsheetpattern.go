package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationSpreadsheetPattern struct {
	ole.IUnknown
}

type IUIAutomationSpreadsheetPatternVtbl struct {
	ole.IUnknownVtbl
	GetItemByName uintptr
}

func (v *IUIAutomationSpreadsheetPattern) VTable() *IUIAutomationSpreadsheetPatternVtbl {
	return (*IUIAutomationSpreadsheetPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationSpreadsheetPattern) GetItemByName() error {
	return ole.NewError(ole.E_NOTIMPL)
}
