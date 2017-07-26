package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationStylesPattern struct {
	ole.IUnknown
}

type IUIAutomationStylesPatternVtbl struct {
	ole.IUnknownVtbl
	Get_CurrentStyleId                  uintptr
	Get_CurrentStyleName                uintptr
	Get_CurrentFillColor                uintptr
	Get_CurrentFillPatternStyle         uintptr
	Get_CurrentShape                    uintptr
	Get_CurrentFillPatternColor         uintptr
	Get_CurrentExtendedProperties       uintptr
	GetCurrentExtendedPropertiesAsArray uintptr
	Get_CachedStyleId                   uintptr
	Get_CachedStyleName                 uintptr
	Get_CachedFillColor                 uintptr
	Get_CachedFillPatternStyle          uintptr
	Get_CachedShape                     uintptr
	Get_CachedFillPatternColor          uintptr
	Get_CachedExtendedProperties        uintptr
	GetCachedExtendedPropertiesAsArray  uintptr
}

func (v *IUIAutomationStylesPattern) VTable() *IUIAutomationStylesPatternVtbl {
	return (*IUIAutomationStylesPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationStylesPattern) Get_CurrentStyleId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) Get_CurrentStyleName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) Get_CurrentFillColor() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) Get_CurrentFillPatternStyle() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) Get_CurrentShape() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) Get_CurrentFillPatternColor() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) Get_CurrentExtendedProperties() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) GetCurrentExtendedPropertiesAsArray() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) Get_CachedStyleId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) Get_CachedStyleName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) Get_CachedFillColor() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) Get_CachedFillPatternStyle() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) Get_CachedShape() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) Get_CachedFillPatternColor() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) Get_CachedExtendedProperties() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationStylesPattern) GetCachedExtendedPropertiesAsArray() error {
	return ole.NewError(ole.E_NOTIMPL)
}
