//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	. "github.com/ying32/govcl/vcl/api"
)

// SetShortCutFromString
//
// 设置快捷键字符
//
// Set shortcut key characters.
func (m *TMenuItem) SetShortCutFromString(s string) {
	MenuItem_SetShortCut(m._instance(), DTextToShortCut(s))
}

// ShortCutFromString
//
// 获取快捷键字符
//
// Get shortcut key characters.
func (m *TMenuItem) ShortCutFromString() string {
	return DShortCutToText(MenuItem_GetShortCut(m._instance()))
}
