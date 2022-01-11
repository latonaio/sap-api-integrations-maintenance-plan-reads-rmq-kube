package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
			ToCallObjects struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_MaintPlanCallObjects"`
		} `json:"results"`
	} `json:"d"`
}
