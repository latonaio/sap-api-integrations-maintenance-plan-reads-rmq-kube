package responses

type ToStrategyCycle struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			MaintenancePlan                string `json:"MaintenancePlan"`
			MaintenancePlanCycle           string `json:"MaintenancePlanCycle"`
			MaintenanceStrategy            string `json:"MaintenanceStrategy"`
			MaintPlanCycRcrrcIntervalQty   string `json:"MaintPlanCycRcrrcIntervalQty"`
			MaintPlanCycRcrrcIntervalUnit  string `json:"MaintPlanCycRcrrcIntervalUnit"`
			MaintPlanCycleDesc             string `json:"MaintPlanCycleDesc"`
			MaintPlanCycleStartOffsetValue string `json:"MaintPlanCycleStartOffsetValue"`
		} `json:"results"`
	} `json:"d"`
}
