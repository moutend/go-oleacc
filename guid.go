package oleacc

import (
	"github.com/go-ole/go-ole"
)

var (
	IID_IAccessible = ole.NewGUID("{618736e0-3c3d-11cf-810c-00aa00389b71}")

	IID_IUIAutomationSpreadsheetPattern = ole.NewGUID("{7517a7c8-faae-4de9-9f08-29b91e8595c1}")
	IID_IUIAutomationElement            = ole.NewGUID("{d22108aa-8ac5-49a5-837b-37bbb3d7591e}")
)
