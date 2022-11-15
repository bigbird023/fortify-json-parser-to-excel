package data

type (
	FortifyIssue struct {
		ProjectVersionID           string `json:"projectVersionId"`
		LastScanID                 string `json:"lastScanId"`
		ID                         string `json:"id"`
		ProjectVersionName         string `json:"projectVersionName"`
		ProjectName                string `json:"projectName"`
		Revision                   string `json:"revision"`
		FolderID                   string `json:"folderId"`
		FolderGUID                 string `json:"folderGuid"`
		IssueInstanceID            string `json:"issueInstanceId"`
		IssueName                  string `json:"issueName"`
		PrimaryLocation            string `json:"primaryLocation"`
		LineNumber                 string `json:"lineNumber"`
		FullFileName               string `json:"fullFileName"`
		Analyzer                   string `json:"analyzer"`
		Kingdom                    string `json:"kingdom"`
		Friority                   string `json:"friority"`
		Reviewed                   string `json:"reviewed"`
		BugURL                     string `json:"bugURL"`
		ExternalBugID              string `json:"externalBugId"`
		PrimaryTag                 string `json:"primaryTag"`
		HasAttachments             string `json:"hasAttachments"`
		HasCorrelatedIssues        string `json:"hasCorrelatedIssues"`
		ScanStatus                 string `json:"scanStatus"`
		FoundDate                  string `json:"foundDate"`
		RemovedDate                string `json:"removedDate"`
		EngineType                 string `json:"engineType"`
		DisplayEngineType          string `json:"displayEngineType"`
		EngineCategory             string `json:"engineCategory"`
		PrimaryRuleGUID            string `json:"primaryRuleGuid"`
		Impact                     string `json:"impact"`
		Likelihood                 string `json:"likelihood"`
		Severity                   string `json:"severity"`
		Confidence                 string `json:"confidence"`
		Audited                    string `json:"audited"`
		IssueStatus                string `json:"issueStatus"`
		PrimaryTagValueAutoApplied string `json:"primaryTagValueAutoApplied"`
		HasComments                string `json:"hasComments"`
		Removed                    string `json:"removed"`
		Suppressed                 string `json:"suppressed"`
		Hidden                     string `json:"hidden"`
		Href                       string `json:"_href"`
	}
)
