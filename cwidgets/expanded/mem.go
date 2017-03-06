package expanded

import (
	"github.com/bcicen/ctop/cwidgets"
	ui "github.com/gizak/termui"
)

type ExpandedMem struct {
	*ui.BarChart
	hist *IntHist
}

func NewExpandedMem() *ExpandedMem {
	mem := &ExpandedMem{
		ui.NewBarChart(),
		NewIntHist(10),
	}
	mem.BorderLabel = "MEM"
	mem.Height = 10
	mem.Width = colWidth[0]
	mem.BarWidth = 5
	mem.BarGap = 1
	mem.X = 0
	mem.Y = 14
	mem.TextColor = ui.ColorDefault
	mem.Data = mem.hist.Data
	mem.BarColor = ui.ColorGreen
	mem.DataLabels = mem.hist.Labels
	mem.NumFmt = cwidgets.ByteFormatInt
	return mem
}

func (w *ExpandedMem) Update(val int, limit int) {
	// implement our own scaling for mem graph
	if val*4 < limit {
		w.SetMax(val * 4)
	} else {
		w.SetMax(limit)
	}
	w.hist.Append(val)
}
