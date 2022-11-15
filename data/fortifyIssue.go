package data

type (
	//FortifyIssue is the object that holds the details of a fortify issue
	FortifyIssue struct {
		ProjectVersionID           int    `json:"projectVersionId"`
		LastScanID                 int    `json:"lastScanId"`
		ID                         int    `json:"id"`
		ProjectVersionName         string `json:"projectVersionName"`
		ProjectName                string `json:"projectName"`
		Revision                   int    `json:"revision"`
		FolderID                   int    `json:"folderId"`
		FolderGUID                 string `json:"folderGuid"`
		IssueInstanceID            string `json:"issueInstanceId"`
		IssueName                  string `json:"issueName"`
		PrimaryLocation            string `json:"primaryLocation"`
		LineNumber                 int    `json:"lineNumber"`
		FullFileName               string `json:"fullFileName"`
		Analyzer                   string `json:"analyzer"`
		Kingdom                    string `json:"kingdom"`
		Friority                   string `json:"friority"`
		Reviewed                   string `json:"reviewed"`
		BugURL                     string `json:"bugURL"`
		ExternalBugID              string `json:"externalBugId"`
		PrimaryTag                 string `json:"primaryTag"`
		HasAttachments             bool   `json:"hasAttachments"`
		HasCorrelatedIssues        bool   `json:"hasCorrelatedIssues"`
		ScanStatus                 string `json:"scanStatus"`
		FoundDate                  string `json:"foundDate"`
		RemovedDate                string `json:"removedDate"`
		EngineType                 string `json:"engineType"`
		DisplayEngineType          string `json:"displayEngineType"`
		EngineCategory             string `json:"engineCategory"`
		PrimaryRuleGUID            string `json:"primaryRuleGuid"`
		Impact                     int    `json:"impact"`
		Likelihood                 int    `json:"likelihood"`
		Severity                   int    `json:"severity"`
		Confidence                 int    `json:"confidence"`
		Audited                    bool   `json:"audited"`
		IssueStatus                string `json:"issueStatus"`
		PrimaryTagValueAutoApplied bool   `json:"primaryTagValueAutoApplied"`
		HasComments                bool   `json:"hasComments"`
		Removed                    bool   `json:"removed"`
		Suppressed                 bool   `json:"suppressed"`
		Hidden                     bool   `json:"hidden"`
		Href                       string `json:"_href"`
	}
)
