// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package walk

func (wb *WindowBase) CalculateTextSizeImpl(text string) Size {
	return wb.calculateTextSizeImpl(text)
}

func (wb *WindowBase) CalculateTextSizeImplForWidth(text string, width int) Size {
	return wb.calculateTextSizeImplForWidth(text, width)
}
