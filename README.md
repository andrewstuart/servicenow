# servicenow
--
    import "astuart.co/servicenow"

Package servicenow currently provides primarily low-level abstractions for
communicating with the ServiceNow API, and data types useful for unmarshaling
the nonstandard JSON types that ServiceNow returns (e.g. non-ISO datetime
format). As needs arise, higher-level abstractions may be added that provide
easy access to standard structured objects in the ServiceNow API, for example a
method to easily access structs representing change requests.

## Usage

```go
const TableCMDB = "cmdb_ci"
```

```go
const TableChangeRequests = "change_request"
```

#### type CMDBItem

```go
type CMDBItem struct {
	Status              string `json:"__status"`
	Asset               string `json:"asset"`
	AssetTag            string `json:"asset_tag"`
	Assigned            string `json:"assigned"`
	AssignedTo          string `json:"assigned_to"`
	AssignmentGroup     string `json:"assignment_group"`
	Attributes          string `json:"attributes"`
	CanPrint            string `json:"can_print"`
	Category            string `json:"category"`
	ChangeControl       string `json:"change_control"`
	CheckedIn           string `json:"checked_in"`
	CheckedOut          string `json:"checked_out"`
	Comments            string `json:"comments"`
	Company             string `json:"company"`
	CorrelationID       string `json:"correlation_id"`
	Cost                string `json:"cost"`
	CostCc              string `json:"cost_cc"`
	CostCenter          string `json:"cost_center"`
	DeliveryDate        string `json:"delivery_date"`
	Department          string `json:"department"`
	DiscoverySource     string `json:"discovery_source"`
	DNSDomain           string `json:"dns_domain"`
	Due                 string `json:"due"`
	DueIn               string `json:"due_in"`
	FaultCount          string `json:"fault_count"`
	FirstDiscovered     string `json:"first_discovered"`
	Fqdn                string `json:"fqdn"`
	GlAccount           string `json:"gl_account"`
	InstallDate         string `json:"install_date"`
	InstallStatus       string `json:"install_status"`
	InvoiceNumber       string `json:"invoice_number"`
	IPAddress           string `json:"ip_address"`
	Justification       string `json:"justification"`
	LastDiscovered      string `json:"last_discovered"`
	LeaseID             string `json:"lease_id"`
	Location            string `json:"location"`
	MacAddress          string `json:"mac_address"`
	MaintenanceSchedule string `json:"maintenance_schedule"`
	ManagedBy           string `json:"managed_by"`
	Manufacturer        string `json:"manufacturer"`
	ModelID             string `json:"model_id"`
	ModelNumber         string `json:"model_number"`
	Monitor             string `json:"monitor"`
	Name                string `json:"name"`
	OperationalStatus   string `json:"operational_status"`
	OrderDate           string `json:"order_date"`
	OwnedBy             string `json:"owned_by"`
	PoNumber            string `json:"po_number"`
	PurchaseDate        string `json:"purchase_date"`
	Schedule            string `json:"schedule"`
	SerialNumber        string `json:"serial_number"`
	ShortDescription    string `json:"short_description"`
	SkipSync            string `json:"skip_sync"`
	StartDate           string `json:"start_date"`
	Subcategory         string `json:"subcategory"`
	SupportGroup        string `json:"support_group"`
	SupportedBy         string `json:"supported_by"`
	SysClassName        string `json:"sys_class_name"`
	SysClassPath        string `json:"sys_class_path"`
	SysCreatedBy        string `json:"sys_created_by"`
	SysCreatedOn        string `json:"sys_created_on"`
	SysDomain           string `json:"sys_domain"`
	SysDomainPath       string `json:"sys_domain_path"`
	SysID               string `json:"sys_id"`
	SysModCount         string `json:"sys_mod_count"`
	SysTags             string `json:"sys_tags"`
	SysUpdatedBy        string `json:"sys_updated_by"`
	SysUpdatedOn        string `json:"sys_updated_on"`
	Unverified          string `json:"unverified"`
	Vendor              string `json:"vendor"`
	WarrantyExpiration  string `json:"warranty_expiration"`
}
```


#### type ChangeRequest

```go
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
```


#### type Client

```go
type Client struct {
	Username, Password, Instance string
}
```

Client helps users interact with a ServiceNow instance.

#### func (Client) GetCMDBItems

```go
func (c Client) GetCMDBItems(query url.Values) ([]CMDBItem, error)
```

#### func (Client) GetChangeRequests

```go
func (c Client) GetChangeRequests(query url.Values) ([]ChangeRequest, error)
```

#### func (Client) GetFor

```go
func (c Client) GetFor(table string, id string, opts url.Values, out interface{}) error
```
GetFor performs a servicenow get to the specified table, with options, and
unmarhals JSON into the output parameter.

#### func (Client) GetRecords

```go
func (c Client) GetRecords(table string, opts url.Values) ([]map[string]interface{}, error)
```
GetRecords performs a servicenow getRecords to the specified table, with
options, and returns a map for each item

#### func (Client) GetRecordsFor

```go
func (c Client) GetRecordsFor(table string, opts url.Values, out interface{}) error
```
GetRecordsFor performs a servicenow getRecords to the specified table, with
options, and unmarhals JSON into the output parameter.

#### func (Client) Insert

```go
func (c Client) Insert(table string, obj, out interface{}) error
```

#### func (Client) PerformFor

```go
func (c Client) PerformFor(table, action, id string, opts url.Values, body interface{}, out interface{}) error
```
PerformFor creates and executes an authenticated HTTP request to an instance,
for the given table, action and optional id, with the passed options, and
unmarhals the JSON into the passed output interface pointer, returning an error.

#### type Err

```go
type Err struct {
	Err    string `json:"error"`
	Reason string
}
```


#### func (Err) Error

```go
func (e Err) Error() string
```

#### type SNTime

```go
type SNTime struct {
	time.Time
}
```


#### func (*SNTime) MarshalJSON

```go
func (s *SNTime) MarshalJSON() ([]byte, error)
```

#### func (*SNTime) UnmarshalJSON

```go
func (s *SNTime) UnmarshalJSON(bs []byte) error
```
