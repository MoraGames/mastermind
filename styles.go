package main

type IndicatorStyle string

func NewArrowStyle(style string) *IndicatorStyle {
	newIndicatorStyle := IndicatorStyle(style)
	return &newIndicatorStyle
}

func (is *IndicatorStyle) Multiply(num int) {
	for i := 1; i < num; i++ {
		*is += *is
	}
}

func (is IndicatorStyle) Length() int {
	return len(is)
}

type SelectionStyle [2]IndicatorStyle

func NewSelectionStyle(leftIndicatorStyle, rightIndicatorStyle IndicatorStyle) *SelectionStyle {
	return &SelectionStyle{leftIndicatorStyle, rightIndicatorStyle}
}

func (ss SelectionStyle) Length() int {
	return len(ss[0]) + len (ss[1]) + 2
}