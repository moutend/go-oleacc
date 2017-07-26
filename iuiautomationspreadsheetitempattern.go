package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationSpreadsheetItemPattern struct {
	ole.IUnknown
}

type IUIAutomationSpreadsheetItemPatternVtbl struct {
	ole.IUnknownVtbl
	Get_CurrentFormula          uintptr
	GetCurrentAnnotationObjects uintptr
	GetCurrentAnnotationTypes   uintptr
	Get_CachedFormula           uintptr
	GetCachedAnnotationObjects  uintptr
	GetCachedAnnotationTypes    uintptr
}

func (v *IUIAutomationSpreadsheetItemPattern) VTable() *IUIAutomationSpreadsheetItemPatternVtbl {
	return (*IUIAutomationSpreadsheetItemPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationSpreadsheetItemPattern) Get_CurrentFormula() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationSpreadsheetItemPattern) GetCurrentAnnotationObjects() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationSpreadsheetItemPattern) GetCurrentAnnotationTypes() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationSpreadsheetItemPattern) Get_CachedFormula() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationSpreadsheetItemPattern) GetCachedAnnotationObjects() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationSpreadsheetItemPattern) GetCachedAnnotationTypes() error {
	return ole.NewError(ole.E_NOTIMPL)
}
