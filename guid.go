package oleacc

import (
	"github.com/go-ole/go-ole"
)

var (
	IID_IAccessible = ole.NewGUID("{618736e0-3c3d-11cf-810c-00aa00389b71}")

	IID_IUIAutomationSpreadsheetPattern           = ole.NewGUID("{7517a7c8-faae-4de9-9f08-29b91e8595c1}")
	IID_IUIAutomationElement                      = ole.NewGUID("{d22108aa-8ac5-49a5-837b-37bbb3d7591e}")
	IID_IUIAutomation                             = ole.NewGUID("{30cbe57d-d9d0-452a-ab13-7ac5ac4825ee}")
	IID_IUIAutomation2                            = ole.NewGUID("{34723aff-0c9d-49d0-9896-7ab52df8cd8a}")
	IID_IUIAutomation3                            = ole.NewGUID("{73D768DA-9B51-4B89-936E-C209290973E7}")
	IID_IUIAutomation4                            = ole.NewGUID("{1189C02A-05F8-4319-8E21-E817E3DB2860]")
	IID_IUIAutomationCacheRequest                 = ole.NewGUID("{b32a92b5-bc25-4078-9c08-d7ee95c48e03}")
	IID_IUIAutomationCondition                    = ole.NewGUID("{352ffba8-0973-437c-a61f-f64cafd81df9}")
	IID_IUIAutomationTreeWalker                   = ole.NewGUID("{4042c624-389c-4afc-a630-9df854a541fc}")
	IID_IUIAutomationEventHandler                 = ole.NewGUID("{146c3c17-f12e-4e22-8c27-f894b9b79c69}")
	IID_IUIAutomationPropertyChangedEventHandler  = ole.NewGUID("{40cd37d4-c756-4b0c-8c6f-bddfeeb13b50}")
	IID_IUIAutomationStructureChangedEventHandler = ole.NewGUID("{e81d1b4e-11c5-42f8-9754-e7036c79f054}")
)
