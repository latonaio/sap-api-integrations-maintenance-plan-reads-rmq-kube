package responses

type ToCallObjects struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			MaintenancePlan              string      `json:"MaintenancePlan"`
			MaintenancePlanCallNumber    int         `json:"MaintenancePlanCallNumber"`
			MaintenanceItem              string      `json:"MaintenanceItem"`
			MaintenanceOrder             string      `json:"MaintenanceOrder"`
			MaintenanceNotification      string      `json:"MaintenanceNotification"`
			ServiceOrder                 string      `json:"ServiceOrder"`
			MaintCallHorizonIsNotReached bool        `json:"MaintCallHorizonIsNotReached"`
			SchedulingStatus             string      `json:"SchedulingStatus"`
			PlannedStartDate             string      `json:"PlannedStartDate"`
		} `json:"results"`
	} `json:"d"`
}
