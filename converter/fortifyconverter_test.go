package converter

import (
	"os"
	"strconv"
	"testing"

	"github.com/bigbird023/fortify-json-parser-to-excel/data"
	"github.com/bigbird023/fortify-json-parser-to-excel/mock"
	"github.com/bigbird023/fortify-json-parser-to-excel/parser"
	"github.com/plandem/xlsx"
)

func compareNextColValue(t *testing.T, colNumber *int, row *xlsx.Row, expectedValue string) {
	*colNumber++
	cell := row.Cell(*colNumber)
	if cell.Value() != expectedValue {
		t.Logf("Expected value %s does not equal actual value %s", expectedValue, cell.Value())
		t.Fail()
	}
}

func TestHeaderToExcel(t *testing.T) {

	c := NewConverter("", "", nil)

	excelFile := xlsx.New()
	sheet := excelFile.AddSheet("fortifyIssues")

	c.headerToExcel(sheet)

	cols, rows := sheet.Dimension()
	if cols != len(c.headers) {
		t.Logf("%d cols are expected", len(c.headers))
		t.Fail()
	}

	if rows != 1 {
		t.Log("1 cols are expected")
		t.Fail()
	}

	row := sheet.Row(0)
	for cols := 0; cols < len(c.headers); cols++ {
		if row.Cell(cols).Value() != c.headers[cols] {
			t.Logf("Header Column %d should be %s", cols, c.headers[cols])
			t.Fail()
		}
	}

}

func TestIssueToExcel(t *testing.T) {

	c := NewConverter("", "", nil)

	excelFile := xlsx.New()
	sheet := excelFile.AddSheet("fortifyIssues")

	issue := newTestFortifyIssue()

	c.issueToExcel(issue, sheet)

	cols, rows := sheet.Dimension()
	if cols != len(c.headers) {
		t.Logf("%d cols are expected", len(c.headers))
		t.Fail()
	}

	if rows != 1 {
		t.Log("1 rows are expected")
		t.Fail()
	}
	row := sheet.Row(rows - 1)

	assertRowToIssue(t, row, issue)
}

func TestHeaderAndIssueToExcel(t *testing.T) {

	c := NewConverter("", "", nil)

	excelFile := xlsx.New()
	sheet := excelFile.AddSheet("fortifyIssues")

	issue := newTestFortifyIssue()

	c.headerToExcel(sheet)

	err := c.issueToExcel(issue, sheet)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	cols, rows := sheet.Dimension()
	if cols != len(c.headers) {
		t.Log("17 cols are expected")
		t.Fail()
	}

	if rows != 2 {
		t.Log("1 rows are expected")
		t.Fail()
	}
	row := sheet.Row(rows - 1)

	assertRowToIssue(t, row, issue)
}

func TestWriteToExcel(t *testing.T) {

	expected := "./local/output/testwritetoexcel.xlsx"
	fxp := parser.NewFortifyJSONParser()

	c := NewConverter("", expected, fxp)

	excelFile := mock.NewPackage()

	err := c.writeExcelToFile(excelFile)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if excelFile.SaveAsInterface != expected {
		t.Logf("excel saveas interface was not set correct expected %s, actual %s", expected, excelFile.SaveAsInterface)
		t.Fail()
	}

}

func TestWriteToExcelError(t *testing.T) {

	expected := "./local/output/testwritetoexcel.xlsx"
	fxp := parser.NewFortifyJSONParser()

	c := NewConverter("", expected, fxp)

	excelFile := mock.NewPackage()
	excelFile.ForceError = true

	err := c.writeExcelToFile(excelFile)
	if err == nil {
		t.Log("error not thrown")
		t.Fail()
	}

}

func TestConvert(t *testing.T) {

	expectedInput := "../local/examples/test.json"
	expectedOutput := "../local/output/test.xlsx"
	fxp := parser.NewFortifyJSONParser()

	c := NewConverter(expectedInput, expectedOutput, fxp)

	err := c.Convert()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

}

func TestConvertInputError(t *testing.T) {

	expectedInput := "./local/test/testmissing.json"
	expectedOutput := "./local/output/testwritetoexcel.xlsx"
	fxp := parser.NewFortifyJSONParser()

	c := NewConverter(expectedInput, expectedOutput, fxp)

	err := c.Convert()
	if err == nil {
		t.Log("Error expected, missing")
		t.Fail()
	}

}

func TestConvertOutputError(t *testing.T) {
	expectedOutput := "../local/output2/TestConvert.xlsx"

	//verify TestConvert.xlsx doesn't exist, delete if does
	_, err := os.Stat(expectedOutput)
	if err != nil {
		if !os.IsNotExist(err) {
			t.Log(err)
			t.Fail()
		}
	} else {
		err = os.Remove(expectedOutput)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
	}

	fxp := mock.NewFortifyJSONParser()

	c := NewConverter("", expectedOutput, fxp)

	err = c.Convert()
	if err == nil {
		t.Log("Should have generated error")
		t.Fail()
	}

}

func TestConvertIssueToExcelError(t *testing.T) {
	expectedOutput := "../local/output2/TestConvert.xlsx"

	//verify TestConvert.xlsx doesn't exist, delete if does
	_, err := os.Stat(expectedOutput)
	if err != nil {
		if !os.IsNotExist(err) {
			t.Log(err)
			t.Fail()
		}
	} else {
		err = os.Remove(expectedOutput)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
	}

	fxp := mock.NewFortifyJSONParser()
	fxp.EmptyReportDefinition = true

	c := NewConverter("", expectedOutput, fxp)

	err = c.Convert()
	if err == nil {
		t.Log("Should have generated error")
		t.Fail()
	}
}

func newTestFortifyIssue() *data.FortifyIssue {
	return &data.FortifyIssue{
		ProjectVersionID:           1,
		LastScanID:                 2,
		ID:                         3,
		ProjectVersionName:         "ProjectVersionName",
		ProjectName:                "ProjectName",
		Revision:                   4,
		FolderID:                   5,
		FolderGUID:                 "FolderGUID",
		IssueInstanceID:            "IssueInstanceID",
		IssueName:                  "IssueName",
		PrimaryLocation:            "PrimaryLocation",
		LineNumber:                 6,
		FullFileName:               "FullFileName",
		Analyzer:                   "Analyzer",
		Kingdom:                    "Kingdom",
		Friority:                   "Friority",
		Reviewed:                   "Reviewed",
		BugURL:                     "BugURL",
		ExternalBugID:              "ExternalBugID",
		PrimaryTag:                 "PrimaryTag",
		HasAttachments:             true,
		HasCorrelatedIssues:        true,
		ScanStatus:                 "ScanStatus",
		FoundDate:                  "FoundDate",
		RemovedDate:                "RemovedDate",
		EngineType:                 "EngineType",
		DisplayEngineType:          "DisplayEngineType",
		EngineCategory:             "EngineCategory",
		PrimaryRuleGUID:            "PrimaryRuleGUID",
		Impact:                     7,
		Likelihood:                 8,
		Severity:                   9,
		Confidence:                 10,
		Audited:                    true,
		IssueStatus:                "IssueStatus",
		PrimaryTagValueAutoApplied: true,
		HasComments:                true,
		Removed:                    true,
		Suppressed:                 true,
		Hidden:                     true,
		Href:                       "Href",
	}
}

func assertRowToIssue(t *testing.T, row *xlsx.Row, issue *data.FortifyIssue) {
	col := -1
	compareNextColValue(t, &col, row, strconv.Itoa(issue.ProjectVersionID))
	compareNextColValue(t, &col, row, strconv.Itoa(issue.LastScanID))
	compareNextColValue(t, &col, row, strconv.Itoa(issue.ID))
	compareNextColValue(t, &col, row, issue.ProjectVersionName)
	compareNextColValue(t, &col, row, issue.ProjectName)
	compareNextColValue(t, &col, row, strconv.Itoa(issue.Revision))
	compareNextColValue(t, &col, row, strconv.Itoa(issue.FolderID))
	compareNextColValue(t, &col, row, issue.FolderGUID)
	compareNextColValue(t, &col, row, issue.IssueInstanceID)
	compareNextColValue(t, &col, row, issue.IssueName)
	compareNextColValue(t, &col, row, issue.PrimaryLocation)
	compareNextColValue(t, &col, row, strconv.Itoa(issue.LineNumber))
	compareNextColValue(t, &col, row, issue.FullFileName)
	compareNextColValue(t, &col, row, issue.Analyzer)
	compareNextColValue(t, &col, row, issue.Kingdom)
	compareNextColValue(t, &col, row, issue.Friority)
	compareNextColValue(t, &col, row, issue.Reviewed)
	compareNextColValue(t, &col, row, issue.BugURL)
	compareNextColValue(t, &col, row, issue.ExternalBugID)
	compareNextColValue(t, &col, row, issue.PrimaryTag)
	compareNextColValue(t, &col, row, strconv.FormatBool(issue.HasAttachments))
	compareNextColValue(t, &col, row, strconv.FormatBool(issue.HasCorrelatedIssues))
	compareNextColValue(t, &col, row, issue.ScanStatus)
	compareNextColValue(t, &col, row, issue.FoundDate)
	compareNextColValue(t, &col, row, issue.RemovedDate)
	compareNextColValue(t, &col, row, issue.EngineType)
	compareNextColValue(t, &col, row, issue.DisplayEngineType)
	compareNextColValue(t, &col, row, issue.EngineCategory)
	compareNextColValue(t, &col, row, issue.PrimaryRuleGUID)
	compareNextColValue(t, &col, row, strconv.Itoa(issue.Impact))
	compareNextColValue(t, &col, row, strconv.Itoa(issue.Likelihood))
	compareNextColValue(t, &col, row, strconv.Itoa(issue.Severity))
	compareNextColValue(t, &col, row, strconv.Itoa(issue.Confidence))
	compareNextColValue(t, &col, row, strconv.FormatBool(issue.Audited))
	compareNextColValue(t, &col, row, issue.IssueStatus)
	compareNextColValue(t, &col, row, strconv.FormatBool(issue.PrimaryTagValueAutoApplied))
	compareNextColValue(t, &col, row, strconv.FormatBool(issue.HasComments))
	compareNextColValue(t, &col, row, strconv.FormatBool(issue.Removed))
	compareNextColValue(t, &col, row, strconv.FormatBool(issue.Suppressed))
	compareNextColValue(t, &col, row, strconv.FormatBool(issue.Hidden))
	compareNextColValue(t, &col, row, issue.Href)
}
