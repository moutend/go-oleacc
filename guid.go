package oleacc

import (
	"github.com/go-ole/go-ole"
)

var (
	IID_IAccessible = ole.NewGUID("{618736e0-3c3d-11cf-810c-00aa00389b71}")

	IID_IUIAutomationSpreadsheetPattern = ole.NewGUID("{7517a7c8-faae-4de9-9f08-29b91e8595c1}")
	IID_IUIAutomationElement            = ole.NewGUID("{d22108aa-8ac5-49a5-837b-37bbb3d7591e}")
	IID_IUIAutomation                   = ole.NewGUID("{30cbe57d-d9d0-452a-ab13-7ac5ac4825ee}")
	IID_IUIAutomation2                  = ole.NewGUID("{34723aff-0c9d-49d0-9896-7ab52df8cd8a}")
)
