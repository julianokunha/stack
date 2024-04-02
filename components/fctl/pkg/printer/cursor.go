package printer

import (
	"fmt"
	"io"

	"github.com/formancehq/fctl/membershipclient"
	"github.com/pterm/pterm"
)

func Cursor(writer io.Writer, cursor *membershipclient.Cursor) error {
	tableData := pterm.TableData{}
	tableData = append(tableData, []string{pterm.LightCyan("HasMore"), fmt.Sprintf("%v", cursor.HasMore)})
	tableData = append(tableData, []string{pterm.LightCyan("PageSize"), fmt.Sprintf("%d", cursor.PageSize)})
	tableData = append(tableData, []string{pterm.LightCyan("Next"), func() string {
		if cursor.Next == nil {
			return ""
		}
		return *cursor.Next
	}()})
	tableData = append(tableData, []string{pterm.LightCyan("Previous"), func() string {
		if cursor.Previous == nil {
			return ""
		}
		return *cursor.Previous
	}()})

	if err := pterm.DefaultTable.
		WithWriter(writer).
		WithData(tableData).
		Render(); err != nil {
		return err
	}

	return nil
}
