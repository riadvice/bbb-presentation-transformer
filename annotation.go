package bigbluebutton

type Annotation struct {
	id, status, annotationType string
	annotationInfo             string // JSON map[string]Any
	position                   int
	color                      uint
}

type TextAnnotatin struct {
	Annotation
	x, y, calcedFontSize, thickness float64
	points                          []float64
}

type RectangleAnnotation struct {
	Annotation
}

const (
	TEXT_TYPE          = "text"
	PENCIL_TYPE        = "pencil"
	RECTANGLE_TYPE     = "rectangle"
	ELLIPSE_TYPE       = "ellipse"
	TRIANGLE_TYPE      = "triangle"
	LINE_TYPE          = "line"
	POLL_RESULT_TYPE   = "poll_result"
	DRAW_START_STATUS  = "DRAW_START"
	DRAW_UPDATE_STATUS = "DRAW_UPDATE"
	DRAW_END_STATUS    = "DRAW_END"
)
