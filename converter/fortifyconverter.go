package converter

import (
	"fmt"
	"strconv"

	"github.com/bigbird023/fortify-json-parser-to-excel/data"
	"github.com/bigbird023/fortify-json-parser-to-excel/parser"
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx"
)

type (
	//FortifyConverter is the object for controlling the conversion process from JSON to Excel
	FortifyConverter struct {
		headers     []string
		inputFile   string
		outputFile  string
		fortifyjson parser.FortifyJSONParserInterface
	}
)

//NewConverter is the ability to create the default object configuration
func NewConverter(input string, output string, fxp parser.FortifyJSONParserInterface) *FortifyConverter {
	conv := &FortifyConverter{
		headers:     []string{"ProjectVersionID", "LastScanID", "ID", "ProjectVersionName", "ProjectName", "Revision", "FolderID", "FolderGUID", "IssueInstanceID", "IssueName", "PrimaryLocation", "LineNumber", "FullFileName", "Analyzer", "Kingdom", "Friority", "Reviewed", "BugURL", "ExternalBugID", "PrimaryTag", "HasAttachments", "HasCorrelatedIssues", "ScanStatus", "FoundDate", "RemovedDate", "EngineType", "DisplayEngineType", "EngineCategory", "PrimaryRuleGUID", "Impact", "Likelihood", "Severity", "Confidence", "Audited", "IssueStatus", "PrimaryTagValueAutoApplied", "HasComments", "Removed", "Suppressed", "Hidden", "Href"},
		inputFile:   input,
		outputFile:  output,
		fortifyjson: fxp,
	}
	return conv
}

//Convert is the main execution of the conversion
func (c *FortifyConverter) Convert() error {
	excelFile := xlsx.New()
	sheet := excelFile.AddSheet("fortifyIssues")

	fortifyJSON, err := c.fortifyjson.JSONParse(c.inputFile)
	if err != nil {
		return err
	}

	c.headerToExcel(sheet)
	if err != nil {
		return err
	}

	if fortifyJSON != nil {
		for rsloop := 0; rsloop < len(*fortifyJSON); rsloop++ {
			err := c.issueToExcel(&(*fortifyJSON)[rsloop], sheet)
			if err != nil {
				err = fmt.Errorf("error converting issue to excel")
				return err
			}
		}
	}

	err = c.writeExcelToFile(excelFile)
	if err != nil {
		return err
	}

	return nil
}

func (c *FortifyConverter) headerToExcel(sheet xlsx.Sheet) {

	row := sheet.Row(0)

	for p, v := range c.headers {
		cell := row.Cell(p)
		cell.SetValue(v)
	}
}

func (c *FortifyConverter) issueToExcel(issue *data.FortifyIssue, sheet xlsx.Sheet) error {

	_, totalRows := sheet.Dimension()
	row := sheet.Row(totalRows - 1)

	if row.Cell(0).Value() != "" {
		//if headers are set, move down
		row = sheet.Row(totalRows)
	}

	col := -1
	c.setNextCell(&col, row, strconv.Itoa(issue.ProjectVersionID))
	c.setNextCell(&col, row, strconv.Itoa(issue.LastScanID))
	c.setNextCell(&col, row, strconv.Itoa(issue.ID))
	c.setNextCell(&col, row, issue.ProjectVersionName)
	c.setNextCell(&col, row, issue.ProjectName)
	c.setNextCell(&col, row, strconv.Itoa(issue.Revision))
	c.setNextCell(&col, row, strconv.Itoa(issue.FolderID))
	c.setNextCell(&col, row, issue.FolderGUID)
	c.setNextCell(&col, row, issue.IssueInstanceID)
	c.setNextCell(&col, row, issue.IssueName)
	c.setNextCell(&col, row, issue.PrimaryLocation)
	c.setNextCell(&col, row, strconv.Itoa(issue.LineNumber))
	c.setNextCell(&col, row, issue.FullFileName)
	c.setNextCell(&col, row, issue.Analyzer)
	c.setNextCell(&col, row, issue.Kingdom)
	c.setNextCell(&col, row, issue.Friority)
	c.setNextCell(&col, row, issue.Reviewed)
	c.setNextCell(&col, row, issue.BugURL)
	c.setNextCell(&col, row, issue.ExternalBugID)
	c.setNextCell(&col, row, issue.PrimaryTag)
	c.setNextCell(&col, row, strconv.FormatBool(issue.HasAttachments))
	c.setNextCell(&col, row, strconv.FormatBool(issue.HasCorrelatedIssues))
	c.setNextCell(&col, row, issue.ScanStatus)
	c.setNextCell(&col, row, issue.FoundDate)
	c.setNextCell(&col, row, issue.RemovedDate)
	c.setNextCell(&col, row, issue.EngineType)
	c.setNextCell(&col, row, issue.DisplayEngineType)
	c.setNextCell(&col, row, issue.EngineCategory)
	c.setNextCell(&col, row, issue.PrimaryRuleGUID)
	c.setNextCell(&col, row, strconv.Itoa(issue.Impact))
	c.setNextCell(&col, row, strconv.Itoa(issue.Likelihood))
	c.setNextCell(&col, row, strconv.Itoa(issue.Severity))
	c.setNextCell(&col, row, strconv.Itoa(issue.Confidence))
	c.setNextCell(&col, row, strconv.FormatBool(issue.Audited))
	c.setNextCell(&col, row, issue.IssueStatus)
	c.setNextCell(&col, row, strconv.FormatBool(issue.PrimaryTagValueAutoApplied))
	c.setNextCell(&col, row, strconv.FormatBool(issue.HasComments))
	c.setNextCell(&col, row, strconv.FormatBool(issue.Removed))
	c.setNextCell(&col, row, strconv.FormatBool(issue.Suppressed))
	c.setNextCell(&col, row, strconv.FormatBool(issue.Hidden))
	c.setNextCell(&col, row, issue.Href)

	return nil
}

func (c *FortifyConverter) setNextCell(colNumber *int, row *xlsx.Row, value string) {
	*colNumber++
	cell := row.Cell(*colNumber)
	cell.SetValue(value)
}

func (c *FortifyConverter) writeExcelToFile(excelFile ooxml.Package) error {

	// Save the XLSX file under different name
	err := excelFile.SaveAs(c.outputFile)
	if err != nil {
		return err
	}

	return nil
}
