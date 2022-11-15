package converter

import (
	"os"
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
		ProjectVersionID:           "ProjectVersionID",
		LastScanID:                 "LastScanID",
		ID:                         "ID",
		ProjectVersionName:         "ProjectVersionName",
		ProjectName:                "ProjectName",
		Revision:                   "Revision",
		FolderID:                   "FolderID",
		FolderGUID:                 "FolderGUID",
		IssueInstanceID:            "IssueInstanceID",
		IssueName:                  "IssueName",
		PrimaryLocation:            "PrimaryLocation",
		LineNumber:                 "LineNumber",
		FullFileName:               "FullFileName",
		Analyzer:                   "Analyzer",
		Kingdom:                    "Kingdom",
		Friority:                   "Friority",
		Reviewed:                   "Reviewed",
		BugURL:                     "BugURL",
		ExternalBugID:              "ExternalBugID",
		PrimaryTag:                 "PrimaryTag",
		HasAttachments:             "HasAttachments",
		HasCorrelatedIssues:        "HasCorrelatedIssues",
		ScanStatus:                 "ScanStatus",
		FoundDate:                  "FoundDate",
		RemovedDate:                "RemovedDate",
		EngineType:                 "EngineType",
		DisplayEngineType:          "DisplayEngineType",
		EngineCategory:             "EngineCategory",
		PrimaryRuleGUID:            "PrimaryRuleGUID",
		Impact:                     "Impact",
		Likelihood:                 "Likelihood",
		Severity:                   "Severity",
		Confidence:                 "Confidence",
		Audited:                    "Audited",
		IssueStatus:                "IssueStatus",
		PrimaryTagValueAutoApplied: "PrimaryTagValueAutoApplied",
		HasComments:                "HasComments",
		Removed:                    "Removed",
		Suppressed:                 "Suppressed",
		Hidden:                     "Hidden",
		Href:                       "Href",
	}
}

func assertRowToIssue(t *testing.T, row *xlsx.Row, issue *data.FortifyIssue) {
	col := -1
	compareNextColValue(t, &col, row, issue.ProjectVersionID)
	compareNextColValue(t, &col, row, issue.LastScanID)
	compareNextColValue(t, &col, row, issue.ID)
	compareNextColValue(t, &col, row, issue.ProjectVersionName)
	compareNextColValue(t, &col, row, issue.ProjectName)
	compareNextColValue(t, &col, row, issue.Revision)
	compareNextColValue(t, &col, row, issue.FolderID)
	compareNextColValue(t, &col, row, issue.FolderGUID)
	compareNextColValue(t, &col, row, issue.IssueInstanceID)
	compareNextColValue(t, &col, row, issue.IssueName)
	compareNextColValue(t, &col, row, issue.PrimaryLocation)
	compareNextColValue(t, &col, row, issue.LineNumber)
	compareNextColValue(t, &col, row, issue.FullFileName)
	compareNextColValue(t, &col, row, issue.Analyzer)
	compareNextColValue(t, &col, row, issue.Kingdom)
	compareNextColValue(t, &col, row, issue.Friority)
	compareNextColValue(t, &col, row, issue.Reviewed)
	compareNextColValue(t, &col, row, issue.BugURL)
	compareNextColValue(t, &col, row, issue.ExternalBugID)
	compareNextColValue(t, &col, row, issue.PrimaryTag)
	compareNextColValue(t, &col, row, issue.HasAttachments)
	compareNextColValue(t, &col, row, issue.HasCorrelatedIssues)
	compareNextColValue(t, &col, row, issue.ScanStatus)
	compareNextColValue(t, &col, row, issue.FoundDate)
	compareNextColValue(t, &col, row, issue.RemovedDate)
	compareNextColValue(t, &col, row, issue.EngineType)
	compareNextColValue(t, &col, row, issue.DisplayEngineType)
	compareNextColValue(t, &col, row, issue.EngineCategory)
	compareNextColValue(t, &col, row, issue.PrimaryRuleGUID)
	compareNextColValue(t, &col, row, issue.Impact)
	compareNextColValue(t, &col, row, issue.Likelihood)
	compareNextColValue(t, &col, row, issue.Severity)
	compareNextColValue(t, &col, row, issue.Confidence)
	compareNextColValue(t, &col, row, issue.Audited)
	compareNextColValue(t, &col, row, issue.IssueStatus)
	compareNextColValue(t, &col, row, issue.PrimaryTagValueAutoApplied)
	compareNextColValue(t, &col, row, issue.HasComments)
	compareNextColValue(t, &col, row, issue.Removed)
	compareNextColValue(t, &col, row, issue.Suppressed)
	compareNextColValue(t, &col, row, issue.Hidden)
	compareNextColValue(t, &col, row, issue.Href)
}
