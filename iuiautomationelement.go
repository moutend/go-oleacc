package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElement struct {
	ole.IUnknown
}

type IUIAutomationElementVtbl struct {
	ole.IUnknownVtbl
	GetRuntimeId                    uintptr
	FindFirst                       uintptr
	FindAll                         uintptr
	FindFirstBuildCache             uintptr
	FindAllBuildCache               uintptr
	BuildUpdatedCache               uintptr
	GetCurrentPropertyValue         uintptr
	GetCurrentPropertyValueEx       uintptr
	GetCachedPropertyValue          uintptr
	GetCachedPropertyValueEx        uintptr
	GetCurrentPatternAs             uintptr
	GetCachedPatternAs              uintptr
	GetCurrentPattern               uintptr
	GetCachedPattern                uintptr
	GetCachedParent                 uintptr
	GetCachedChildren               uintptr
	Get_CurrentProcessId            uintptr
	Get_CurrentControlType          uintptr
	Get_CurrentLocalizedControlType uintptr
	Get_CurrentName                 uintptr
	Get_CurrentAcceleratorKey       uintptr
	Get_CurrentAccessKey            uintptr
	Get_CurrentHasKeyboardFocus     uintptr
	Get_CurrentIsKeyboardFocusable  uintptr
	Get_CurrentIsEnabled            uintptr
	Get_CurrentAutomationId         uintptr
	Get_CurrentClassName            uintptr
	Get_CurrentHelpText             uintptr
	Get_CurrentCulture              uintptr
	Get_CurrentIsControlElement     uintptr
	Get_CurrentIsContentElement     uintptr
	Get_CurrentIsPassword           uintptr
	Get_CurrentNativeWindowHandle   uintptr
	Get_CurrentItemType             uintptr
	Get_CurrentIsOffscreen          uintptr
	Get_CurrentOrientation          uintptr
	Get_CurrentFrameworkId          uintptr
	Get_CurrentIsRequiredForForm    uintptr
	Get_CurrentItemStatus           uintptr
	Get_CurrentBoundingRectangle    uintptr
	Get_CurrentLabeledBy            uintptr
	Get_CurrentAriaRole             uintptr
	Get_CurrentAriaProperties       uintptr
	Get_CurrentIsDataValidForForm   uintptr
	Get_CurrentControllerFor        uintptr
	Get_CurrentDescribedBy          uintptr
	Get_CurrentFlowsTo              uintptr
	Get_CurrentProviderDescription  uintptr
	Get_CachedProcessId             uintptr
	Get_CachedControlType           uintptr
	Get_CachedLocalizedControlType  uintptr
	Get_CachedName                  uintptr
	Get_CachedAcceleratorKey        uintptr
	Get_CachedAccessKey             uintptr
	Get_CachedHasKeyboardFocus      uintptr
	Get_CachedIsKeyboardFocusable   uintptr
	Get_CachedIsEnabled             uintptr
	Get_CachedAutomationId          uintptr
	Get_CachedClassName             uintptr
	Get_CachedHelpText              uintptr
	Get_CachedCulture               uintptr
	Get_CachedIsControlElement      uintptr
	Get_CachedIsContentElement      uintptr
	Get_CachedIsPassword            uintptr
	Get_CachedNativeWindowHandle    uintptr
	Get_CachedItemType              uintptr
	Get_CachedIsOffscreen           uintptr
	Get_CachedOrientation           uintptr
	Get_CachedFrameworkId           uintptr
	Get_CachedIsRequiredForForm     uintptr
	Get_CachedItemStatus            uintptr
	Get_CachedBoundingRectangle     uintptr
	Get_CachedLabeledBy             uintptr
	Get_CachedAriaRole              uintptr
	Get_CachedAriaProperties        uintptr
	Get_CachedIsDataValidForForm    uintptr
	Get_CachedControllerFor         uintptr
	Get_CachedDescribedBy           uintptr
	Get_CachedFlowsTo               uintptr
	Get_CachedProviderDescription   uintptr
	GetClickablePoint               uintptr
}

func (v *IUIAutomationElement) VTable() *IUIAutomationElementVtbl {
	return (*IUIAutomationElementVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationElement) GetRuntimeId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) FindFirst() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) FindAll() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) FindFirstBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) FindAllBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) BuildUpdatedCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) GetCurrentPropertyValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) GetCurrentPropertyValueEx() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) GetCachedPropertyValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) GetCachedPropertyValueEx() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) GetCurrentPatternAs() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) GetCachedPatternAs() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) GetCurrentPattern() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) GetCachedPattern() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) GetCachedParent() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) GetCachedChildren() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentProcessId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentControlType() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentLocalizedControlType() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentAcceleratorKey() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentAccessKey() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentHasKeyboardFocus() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentIsKeyboardFocusable() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentIsEnabled() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentAutomationId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentClassName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentHelpText() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentCulture() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentIsControlElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentIsContentElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentIsPassword() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentNativeWindowHandle() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentItemType() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentIsOffscreen() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentOrientation() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentFrameworkId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentIsRequiredForForm() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentItemStatus() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentBoundingRectangle() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentLabeledBy() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentAriaRole() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentAriaProperties() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentIsDataValidForForm() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentControllerFor() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentDescribedBy() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentFlowsTo() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CurrentProviderDescription() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedProcessId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedControlType() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedLocalizedControlType() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedAcceleratorKey() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedAccessKey() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedHasKeyboardFocus() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedIsKeyboardFocusable() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedIsEnabled() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedAutomationId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedClassName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedHelpText() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedCulture() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedIsControlElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedIsContentElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedIsPassword() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedNativeWindowHandle() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedItemType() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedIsOffscreen() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedOrientation() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedFrameworkId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedIsRequiredForForm() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedItemStatus() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedBoundingRectangle() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedLabeledBy() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedAriaRole() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedAriaProperties() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedIsDataValidForForm() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedControllerFor() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedDescribedBy() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedFlowsTo() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) Get_CachedProviderDescription() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement) GetClickablePoint() error {
	return ole.NewError(ole.E_NOTIMPL)
}
