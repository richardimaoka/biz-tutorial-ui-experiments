package state

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type SourceCodeTooltipTiming string

const (
	SOURCE_TOOLTIP_START SourceCodeTooltipTiming = "START"
	SOURCE_TOOLTIP_END   SourceCodeTooltipTiming = "END"
)

func (t SourceCodeTooltipTiming) toGraphQL() model.SourceCodeTooltipTiming {
	switch t {
	case SOURCE_TOOLTIP_START:
		return model.SourceCodeTooltipTimingStart
	case SOURCE_TOOLTIP_END:
		return model.SourceCodeTooltipTimingEnd
	default:
		panic(fmt.Sprintf("SourceCodeTooltipTiming has an invalid value = '%s'", t))
	}
}

type SourceCodeTooltip struct {
	markdownBody string
	timing       SourceCodeTooltipTiming
	lineNumber   int
}

func (t *SourceCodeTooltip) ToGraphQL() *model.SourceCodeTooltip {
	timing := t.timing.toGraphQL()

	return &model.SourceCodeTooltip{
		MarkdownBody: t.markdownBody,
		LineNumber:   t.lineNumber,
		Timing:       &timing,
	}
}
