package servicenow

import (
	"encoding/json"
	"net/url"
)

const TableChangeRequests = "change_request"

func (c Client) GetChangeRequests(query url.Values) ([]ChangeRequest, error) {
	var res struct {
		Records []ChangeRequest
	}
	err := c.GetRecordsFor(TableChangeRequests, query, &res)
	return res.Records, err
}

type ChangeRequest struct {
	Status                   string      `json:"__status"`
	Active                   bool        `json:"active,string"`
	ActivityDue              string      `json:"activity_due"`
	AdditionalAssigneeList   string      `json:"additional_assignee_list"`
	Approval                 string      `json:"approval"`
	ApprovalHistory          string      `json:"approval_history"`
	ApprovalSet              string      `json:"approval_set"`
	AssignedTo               string      `json:"assigned_to"`
	AssignmentGroup          string      `json:"assignment_group"`
	BackoutPlan              string      `json:"backout_plan"`
	BusinessDuration         string      `json:"business_duration"`
	BusinessService          string      `json:"business_service"`
	CabDate                  SNTime      `json:"cab_date"`
	CabDelegate              string      `json:"cab_delegate"`
	CabRecommendation        string      `json:"cab_recommendation"`
	CabRequired              bool        `json:"cab_required,string"`
	CalendarDuration         string      `json:"calendar_duration"`
	Category                 string      `json:"category"`
	ChangePlan               string      `json:"change_plan"`
	CloseCode                string      `json:"close_code"`
	CloseNotes               string      `json:"close_notes"`
	ClosedAt                 SNTime      `json:"closed_at"`
	ClosedBy                 string      `json:"closed_by"`
	CmdbCi                   string      `json:"cmdb_ci"`
	Comments                 string      `json:"comments"`
	CommentsAndWorkNotes     string      `json:"comments_and_work_notes"`
	Company                  string      `json:"company"`
	ConflictLastRun          string      `json:"conflict_last_run"`
	ConflictStatus           string      `json:"conflict_status"`
	ContactType              string      `json:"contact_type"`
	CorrelationDisplay       string      `json:"correlation_display"`
	CorrelationID            string      `json:"correlation_id"`
	DeliveryPlan             string      `json:"delivery_plan"`
	DeliveryTask             string      `json:"delivery_task"`
	Description              string      `json:"description"`
	DueDate                  SNTime      `json:"due_date"`
	EndDate                  SNTime      `json:"end_date"`
	Escalation               json.Number `json:"escalation"`
	ExpectedStart            SNTime      `json:"expected_start"`
	FollowUp                 string      `json:"follow_up"`
	GroupList                string      `json:"group_list"`
	Impact                   json.Number `json:"impact"`
	ImplementationPlan       string      `json:"implementation_plan"`
	Justification            string      `json:"justification"`
	Knowledge                bool        `json:"knowledge,string"`
	Location                 string      `json:"location"`
	MadeSLA                  bool        `json:"made_sla,string"`
	Number                   string      `json:"number"`
	OnHold                   bool        `json:"on_hold,string"`
	OnHoldReason             string      `json:"on_hold_reason"`
	OpenedAt                 SNTime      `json:"opened_at"`
	OpenedBy                 string      `json:"opened_by"`
	Order                    string      `json:"order"`
	OutsideMaintSchedule     bool        `json:"outside_maintenance_schedule,string"`
	Parent                   string      `json:"parent"`
	Phase                    string      `json:"phase"`
	PhaseState               string      `json:"phase_state"`
	Priority                 json.Number `json:"priority"`
	ProductionSystem         bool        `json:"production_system,string"`
	Reason                   string      `json:"reason"`
	ReassignmentCount        json.Number `json:"reassignment_count"`
	RequestedBy              string      `json:"requested_by"`
	RequestedByDate          SNTime      `json:"requested_by_date"`
	ReviewComments           string      `json:"review_comments"`
	ReviewDate               SNTime      `json:"review_date"`
	ReviewStatus             string      `json:"review_status"`
	Risk                     json.Number `json:"risk"`
	RiskImpactAnalysis       string      `json:"risk_impact_analysis"`
	Scope                    json.Number `json:"scope"`
	ShortDescription         string      `json:"short_description"`
	SLADue                   string      `json:"sla_due"`
	StartDate                SNTime      `json:"start_date"`
	State                    string      `json:"state"`
	StdChangeProducerVersion string      `json:"std_change_producer_version"`
	SysClassName             string      `json:"sys_class_name"`
	SysCreatedBy             string      `json:"sys_created_by"`
	SysCreatedOn             SNTime      `json:"sys_created_on"`
	SysDomain                string      `json:"sys_domain"`
	SysDomainPath            string      `json:"sys_domain_path"`
	SysID                    string      `json:"sys_id"`
	SysModCount              json.Number `json:"sys_mod_count"`
	SysTags                  string      `json:"sys_tags"`
	SysUpdatedBy             string      `json:"sys_updated_by"`
	SysUpdatedOn             SNTime      `json:"sys_updated_on"`
	TestPlan                 string      `json:"test_plan"`
	TimeWorked               string      `json:"time_worked"`
	Type                     string      `json:"type"`
	UponApproval             string      `json:"upon_approval"`
	UponReject               string      `json:"upon_reject"`
	Urgency                  string      `json:"urgency"`
	UserInput                string      `json:"user_input"`
	WatchList                string      `json:"watch_list"`
	WorkEnd                  string      `json:"work_end"`
	WorkNotes                string      `json:"work_notes"`
	WorkNotesList            string      `json:"work_notes_list"`
	WorkStart                string      `json:"work_start"`
}
