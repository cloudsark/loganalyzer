package logalizer

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
)

// TablePrint ... print output in table
func TablePrint(field, field1 string, values map[string]int) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{field, field1})

	for _, kv := range SortMapByValue(values) {
		t.AppendRows([]table.Row{
			{kv.Key, kv.Value},
		})
	}
	t.SetStyle(table.StyleColoredBright)
	t.Render()
}

// TableBandwidth ... bandwidth table
func TabelBandwidth(totalSize int) {
	KB := totalSize / 1024
	MB := KB / 1024
	GB := MB / 1024
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Bandwidth"})
	t.AppendHeader(table.Row{"KB", "MB", "GB"})
	t.AppendRows([]table.Row{
		{KB, MB, GB},
	})

	t.SetStyle(table.StyleColoredBright)
	t.Render()

}
