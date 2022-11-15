package mock

import (
	"encoding/json"
	"fmt"

	"github.com/bigbird023/fortify-json-parser-to-excel/parser"
)

//FortifyJSONParser is the main mocked fortify json parser
type FortifyJSONParser struct {
	EmptyReportDefinition bool
	ForceError            bool
}

//NewFortifyJSONParser creates a new MockFortifyJsonParser
func NewFortifyJSONParser() *FortifyJSONParser {
	return &FortifyJSONParser{
		EmptyReportDefinition: false,
		ForceError:            false,
	}
}

//JSONParse will parse json into a golang struct
func (m *FortifyJSONParser) JSONParse(inputFile string) (*parser.FortifyJSON, error) {
	if m.ForceError {
		return nil, fmt.Errorf("forced error")
	}
	var rd *parser.FortifyJSON
	if m.EmptyReportDefinition {
		rd = &parser.FortifyJSON{}
	} else {
		rd = newFortifyJSON()
	}

	return rd, nil
}

func newFortifyJSON() *parser.FortifyJSON {

	jsonDoc := `[
		{
		  "projectVersionId": 3112438,
		  "lastScanId": 3112526,
		  "id": 3144096,
		  "projectVersionName": null,
		  "projectName": null,
		  "revision": 0,
		  "folderId": 3112503,
		  "folderGuid": "bb824e8d-b401-40be-13bd-5d156696a685",
		  "issueInstanceId": "E01A09AFEFC5B1C8334A8B767D9AD59A",
		  "issueName": "Poor Error Handling: Overly Broad Throws",
		  "primaryLocation": "ACABulkAckStatusClient.java",
		  "lineNumber": 35,
		  "fullFileName": "evoBrixX_JAVA_CORR/cld-crsqueue-ejb/src/main/java/com/cnsi/webservices/member/irsclient/BulkSbmtAckStatus/ACABulkAckStatusClient.java",
		  "analyzer": "Structural",
		  "kingdom": "Errors",
		  "friority": "Low",
		  "reviewed": null,
		  "bugURL": null,
		  "externalBugId": null,
		  "primaryTag": null,
		  "hasAttachments": false,
		  "hasCorrelatedIssues": false,
		  "scanStatus": "NEW",
		  "foundDate": "2021-04-01T11:59:41.000+0000",
		  "removedDate": null,
		  "engineType": "SCA",
		  "displayEngineType": "SCA",
		  "engineCategory": "STATIC",
		  "primaryRuleGuid": "572EA1F6-FC86-443E-B1A9-A227D5AD17CC",
		  "impact": 1,
		  "likelihood": 1,
		  "severity": 2,
		  "confidence": 5,
		  "audited": false,
		  "issueStatus": "Unreviewed",
		  "primaryTagValueAutoApplied": false,
		  "hasComments": false,
		  "removed": false,
		  "suppressed": false,
		  "hidden": false,
		  "_href": "https://ctiappvm18.cns-inc.com:10181/ssc/api/v1/projectVersions/3112438/issues/3144096"
		},
		{
		  "projectVersionId": 3112438,
		  "lastScanId": 3112526,
		  "id": 3137078,
		  "projectVersionName": null,
		  "projectName": null,
		  "revision": 0,
		  "folderId": 3112503,
		  "folderGuid": "bb824e8d-b401-40be-13bd-5d156696a685",
		  "issueInstanceId": "6545C65629DC8590797A1FCFA6F6F26A",
		  "issueName": "Poor Error Handling: Overly Broad Throws",
		  "primaryLocation": "ACABulkRequestClient.java",
		  "lineNumber": 82,
		  "fullFileName": "evoBrixX_JAVA_CORR/cld-crsqueue-ejb/src/main/java/com/cnsi/webservices/member/irsclient/BulkSbmtRqst/ACABulkRequestClient.java",
		  "analyzer": "Structural",
		  "kingdom": "Errors",
		  "friority": "Low",
		  "reviewed": null,
		  "bugURL": null,
		  "externalBugId": null,
		  "primaryTag": null,
		  "hasAttachments": false,
		  "hasCorrelatedIssues": false,
		  "scanStatus": "NEW",
		  "foundDate": "2021-04-01T11:59:41.000+0000",
		  "removedDate": null,
		  "engineType": "SCA",
		  "displayEngineType": "SCA",
		  "engineCategory": "STATIC",
		  "primaryRuleGuid": "572EA1F6-FC86-443E-B1A9-A227D5AD17CC",
		  "impact": 1,
		  "likelihood": 1,
		  "severity": 2,
		  "confidence": 5,
		  "audited": false,
		  "issueStatus": "Unreviewed",
		  "primaryTagValueAutoApplied": false,
		  "hasComments": false,
		  "removed": false,
		  "suppressed": false,
		  "hidden": false,
		  "_href": "https://ctiappvm18.cns-inc.com:10181/ssc/api/v1/projectVersions/3112438/issues/3137078"
		},
		{
		  "projectVersionId": 3112438,
		  "lastScanId": 3112526,
		  "id": 3140697,
		  "projectVersionName": null,
		  "projectName": null,
		  "revision": 0,
		  "folderId": 3112503,
		  "folderGuid": "bb824e8d-b401-40be-13bd-5d156696a685",
		  "issueInstanceId": "A1FBC5F3376291D1A965BF6C70DFFE68",
		  "issueName": "Dead Code: Expression is Always true",
		  "primaryLocation": "AccidentDateCheck.java",
		  "lineNumber": 76,
		  "fullFileName": "evoBrixX_JAVA_EDI/cld-ediengine-service/src/main/java/com/cnsi/evobrix/edi/edits/impl/inbnd/healthcareclaim/AccidentDateCheck.java",
		  "analyzer": "Structural",
		  "kingdom": "Code Quality",
		  "friority": "Low",
		  "reviewed": null,
		  "bugURL": null,
		  "externalBugId": null,
		  "primaryTag": null,
		  "hasAttachments": false,
		  "hasCorrelatedIssues": false,
		  "scanStatus": "NEW",
		  "foundDate": "2021-04-01T11:59:41.000+0000",
		  "removedDate": null,
		  "engineType": "SCA",
		  "displayEngineType": "SCA",
		  "engineCategory": "STATIC",
		  "primaryRuleGuid": "0C82D0B5-1B23-4D56-B38E-F6263A454766",
		  "impact": 1,
		  "likelihood": 0.6,
		  "severity": 2,
		  "confidence": 5,
		  "audited": false,
		  "issueStatus": "Unreviewed",
		  "primaryTagValueAutoApplied": false,
		  "hasComments": false,
		  "removed": false,
		  "suppressed": false,
		  "hidden": false,
		  "_href": "https://ctiappvm18.cns-inc.com:10181/ssc/api/v1/projectVersions/3112438/issues/3140697"
		}
	]
	`

	byteValue := []byte(jsonDoc)

	var fortifyJSON *parser.FortifyJSON

	json.Unmarshal(byteValue, &fortifyJSON)

	return fortifyJSON
}
