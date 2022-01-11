package sap_api_output_formatter

type MaintenancePlan struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	APISchema          string `json:"api_schema"`
	MaintenancePlan    string `json:"maintenance_plan"`
	Deleted            bool   `json:"deleted"`
}

type Header struct {
	MaintenancePlan                string      `json:"MaintenancePlan"`
	MaintenancePlanDesc            string      `json:"MaintenancePlanDesc"`
	CreationDate                   string      `json:"CreationDate"`
	LastChangeDate                 string      `json:"LastChangeDate"`
	MaintenanceStrategy            string      `json:"MaintenanceStrategy"`
	SchedulingDuration             string      `json:"SchedulingDuration"`
	SchedulingDurationUnit         string      `json:"SchedulingDurationUnit"`
	NumberOfMaintenanceItems       string      `json:"NumberOfMaintenanceItems"`
	CycleModificationRatio         string      `json:"CycleModificationRatio"`
	MaintPlanSchedgIndicator       string      `json:"MaintPlanSchedgIndicator"`
	MaintenancePlanInternalID      string      `json:"MaintenancePlanInternalID"`
	MaintenanceCall                int         `json:"MaintenanceCall"`
	MaintenancePlanCategory        string      `json:"MaintenancePlanCategory"`
	MaintPlanFreeDefinedAttrib     string      `json:"MaintPlanFreeDefinedAttrib"`
	BasicStartDate                 string      `json:"BasicStartDate"`
	SchedulingStartDate            string      `json:"SchedulingStartDate"`
	SchedulingStartTime            string      `json:"SchedulingStartTime"`
	MaintPlanStartCntrReadingValue string      `json:"MaintPlanStartCntrReadingValue"`
	MaintPlnStrtBufDurationInDays  string      `json:"MaintPlnStrtBufDurationInDays"`
	MaintPlanStartBufferUnit       string      `json:"MaintPlanStartBufferUnit"`
	FactoryCalendar                string      `json:"FactoryCalendar"`
	LateCompletionShiftInPercent   string      `json:"LateCompletionShiftInPercent"`
	LateCompletionTolerancePercent string      `json:"LateCompletionTolerancePercent"`
	EarlyCompletionShiftInPercent  string      `json:"EarlyCompletionShiftInPercent"`
	EarlyCompletionTolerancePct    string      `json:"EarlyCompletionTolerancePct"`
	MaintPlanLogicalOperatorCode   string      `json:"MaintPlanLogicalOperatorCode"`
	SchedulingEndDate              string      `json:"SchedulingEndDate"`
	MaintPlanEndCntrReadingValue   string      `json:"MaintPlanEndCntrReadingValue"`
	LastChangeDateTime             string      `json:"LastChangeDateTime"`
	MultipleCounterPlanShiftFactor string      `json:"MultipleCounterPlanShiftFactor"`
	MaintenanceLeadFloatInDays     string      `json:"MaintenanceLeadFloatInDays"`
	MaintenancePlanCallObject      string      `json:"MaintenancePlanCallObject"`
	MaintenancePlanSystemStatus    string      `json:"MaintenancePlanSystemStatus"`
	ToStrategyCycle                string      `json:"to_StrategyCycle"`
	ToItem                         string      `json:"to_Item"`
}

type ToStrategyCycle struct {
	MaintenancePlan                string `json:"MaintenancePlan"`
	MaintenancePlanCycle           string `json:"MaintenancePlanCycle"`
	MaintenanceStrategy            string `json:"MaintenanceStrategy"`
	MaintPlanCycRcrrcIntervalQty   string `json:"MaintPlanCycRcrrcIntervalQty"`
	MaintPlanCycRcrrcIntervalUnit  string `json:"MaintPlanCycRcrrcIntervalUnit"`
	MaintPlanCycleDesc             string `json:"MaintPlanCycleDesc"`
	MaintPlanCycleStartOffsetValue string `json:"MaintPlanCycleStartOffsetValue"`
}

type ToItem struct {
	MaintenanceItem              string      `json:"MaintenanceItem"`
	MaintenanceItemDescription   string      `json:"MaintenanceItemDescription"`
	MaintenanceStrategy          string      `json:"MaintenanceStrategy"`
	MaintenancePlanCategory      string      `json:"MaintenancePlanCategory"`
	MaintenancePlanCallObject    string      `json:"MaintenancePlanCallObject"`
	MaintenancePlan              string      `json:"MaintenancePlan"`
	MaintenancePlanItemPosition  string      `json:"MaintenancePlanItemPosition"`
	MaintenanceItemObjectList    string      `json:"MaintenanceItemObjectList"`
	FunctionalLocationLabelName  string      `json:"FunctionalLocationLabelName"`
	Equipment                    string      `json:"Equipment"`
	Assembly                     string      `json:"Assembly"`
	AdditionalDeviceData         string      `json:"AdditionalDeviceData"`
	TaskListType                 string      `json:"TaskListType"`
	TaskListGroup                string      `json:"TaskListGroup"`
	TaskListGroupCounter         string      `json:"TaskListGroupCounter"`
	OperationSystemCondition     string      `json:"OperationSystemCondition"`
	NumberOfTaskListExecutions   string      `json:"NumberOfTaskListExecutions"`
	MaintNotifTskIsAutomlyDtmnd  string      `json:"MaintNotifTskIsAutomlyDtmnd"`
	MaintenancePlanningPlant     string      `json:"MaintenancePlanningPlant"`
	MaintenancePlannerGroup      string      `json:"MaintenancePlannerGroup"`
	MaintenanceOrderType         string      `json:"MaintenanceOrderType"`
	NotificationType             string      `json:"NotificationType"`
	MaintenanceActivityType      string      `json:"MaintenanceActivityType"`
	MainWorkCenter               string      `json:"MainWorkCenter"`
	MainWorkCenterPlant          string      `json:"MainWorkCenterPlant"`
	MaintPriority                string      `json:"MaintPriority"`
	MaintPriorityType            string      `json:"MaintPriorityType"`
	BusinessArea                 string      `json:"BusinessArea"`
	ImmediateReleaseIsBlocked    bool        `json:"ImmediateReleaseIsBlocked"`
	Material                     string      `json:"Material"`
	SerialNumber                 string      `json:"SerialNumber"`
	ServiceDocumentType          string      `json:"ServiceDocumentType"`
	ServiceContract              string      `json:"ServiceContract"`
	ServiceContractItem          string      `json:"ServiceContractItem"`
	ServiceOrderTemplate         string      `json:"ServiceOrderTemplate"`
	ServiceDocumentPriority      string      `json:"ServiceDocumentPriority"`
	Product                      string      `json:"Product"`
	MaintenancePlant             string      `json:"MaintenancePlant"`
	AssetLocation                string      `json:"AssetLocation"`
	AssetRoom                    string      `json:"AssetRoom"`
	PlantSection                 string      `json:"PlantSection"`
	WorkCenter                   string      `json:"WorkCenter"`
	ABCIndicator                 string      `json:"ABCIndicator"`
	MaintObjectFreeDefinedAttrib string      `json:"MaintObjectFreeDefinedAttrib"`
	CompanyCode                  string      `json:"CompanyCode"`
	MasterFixedAsset             string      `json:"MasterFixedAsset"`
	FixedAsset                   string      `json:"FixedAsset"`
	LocAcctAssgmtBusinessArea    string      `json:"LocAcctAssgmtBusinessArea"`
	CostCenter                   string      `json:"CostCenter"`
	ControllingArea              string      `json:"ControllingArea"`
	WBSElement                   string      `json:"WBSElement"`
	SettlementOrder              string      `json:"SettlementOrder"`
	CycleSetSequence             string      `json:"CycleSetSequence"`
	StandingOrderNumber          string      `json:"StandingOrderNumber"`
	CreationDate                 string      `json:"CreationDate"`
	LastChangeDate               string      `json:"LastChangeDate"`
	LastChangeDateTime           string      `json:"LastChangeDateTime"`
	ToCallObjects                string      `json:"to_MaintPlanCallObjects"`
}

type ToCallObjects struct {
	MaintenancePlan              string      `json:"MaintenancePlan"`
	MaintenancePlanCallNumber    int         `json:"MaintenancePlanCallNumber"`
	MaintenanceItem              string      `json:"MaintenanceItem"`
	MaintenanceOrder             string      `json:"MaintenanceOrder"`
	MaintenanceNotification      string      `json:"MaintenanceNotification"`
	ServiceOrder                 string      `json:"ServiceOrder"`
	MaintCallHorizonIsNotReached bool        `json:"MaintCallHorizonIsNotReached"`
	SchedulingStatus             string      `json:"SchedulingStatus"`
	PlannedStartDate             string      `json:"PlannedStartDate"`
}
