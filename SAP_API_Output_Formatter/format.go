package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-maintenance-plan-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			MaintenancePlan:                data.MaintenancePlan,
			MaintenancePlanDesc:            data.MaintenancePlanDesc,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			MaintenanceStrategy:            data.MaintenanceStrategy,
			SchedulingDuration:             data.SchedulingDuration,
			SchedulingDurationUnit:         data.SchedulingDurationUnit,
			NumberOfMaintenanceItems:       data.NumberOfMaintenanceItems,
			CycleModificationRatio:         data.CycleModificationRatio,
			MaintPlanSchedgIndicator:       data.MaintPlanSchedgIndicator,
			MaintenancePlanInternalID:      data.MaintenancePlanInternalID,
			MaintenanceCall:                data.MaintenanceCall,
			MaintenancePlanCategory:        data.MaintenancePlanCategory,
			MaintPlanFreeDefinedAttrib:     data.MaintPlanFreeDefinedAttrib,
			BasicStartDate:                 data.BasicStartDate,
			SchedulingStartDate:            data.SchedulingStartDate,
			SchedulingStartTime:            data.SchedulingStartTime,
			MaintPlanStartCntrReadingValue: data.MaintPlanStartCntrReadingValue,
			MaintPlnStrtBufDurationInDays:  data.MaintPlnStrtBufDurationInDays,
			MaintPlanStartBufferUnit:       data.MaintPlanStartBufferUnit,
			FactoryCalendar:                data.FactoryCalendar,
			LateCompletionShiftInPercent:   data.LateCompletionShiftInPercent,
			LateCompletionTolerancePercent: data.LateCompletionTolerancePercent,
			EarlyCompletionShiftInPercent:  data.EarlyCompletionShiftInPercent,
			EarlyCompletionTolerancePct:    data.EarlyCompletionTolerancePct,
			MaintPlanLogicalOperatorCode:   data.MaintPlanLogicalOperatorCode,
			SchedulingEndDate:              data.SchedulingEndDate,
			MaintPlanEndCntrReadingValue:   data.MaintPlanEndCntrReadingValue,
			LastChangeDateTime:             data.LastChangeDateTime,
			MultipleCounterPlanShiftFactor: data.MultipleCounterPlanShiftFactor,
			MaintenanceLeadFloatInDays:     data.MaintenanceLeadFloatInDays,
			MaintenancePlanCallObject:      data.MaintenancePlanCallObject,
			MaintenancePlanSystemStatus:    data.MaintenancePlanSystemStatus,
			ToStrategyCycle:                data.ToStrategyCycle.Deferred.URI,
			ToItem:                         data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToToStrategyCycle(raw []byte, l *logger.Logger) ([]ToStrategyCycle, error) {
	pm := &responses.ToStrategyCycle{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToStrategyCycle. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toStrategyCycle := make([]ToStrategyCycle, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toStrategyCycle = append(toStrategyCycle, ToStrategyCycle{
			MaintenancePlan:                data.MaintenancePlan,
			MaintenancePlanCycle:           data.MaintenancePlanCycle,
			MaintenanceStrategy:            data.MaintenanceStrategy,
			MaintPlanCycRcrrcIntervalQty:   data.MaintPlanCycRcrrcIntervalQty,
			MaintPlanCycRcrrcIntervalUnit:  data.MaintPlanCycRcrrcIntervalUnit,
			MaintPlanCycleDesc:             data.MaintPlanCycleDesc,
			MaintPlanCycleStartOffsetValue: data.MaintPlanCycleStartOffsetValue,
		})
	}

	return toStrategyCycle, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
			MaintenanceItem:              data.MaintenanceItem,
			MaintenanceItemDescription:   data.MaintenanceItemDescription,
			MaintenanceStrategy:          data.MaintenanceStrategy,
			MaintenancePlanCategory:      data.MaintenancePlanCategory,
			MaintenancePlanCallObject:    data.MaintenancePlanCallObject,
			MaintenancePlan:              data.MaintenancePlan,
			MaintenancePlanItemPosition:  data.MaintenancePlanItemPosition,
			MaintenanceItemObjectList:    data.MaintenanceItemObjectList,
			FunctionalLocationLabelName:  data.FunctionalLocationLabelName,
			Equipment:                    data.Equipment,
			Assembly:                     data.Assembly,
			AdditionalDeviceData:         data.AdditionalDeviceData,
			TaskListType:                 data.TaskListType,
			TaskListGroup:                data.TaskListGroup,
			TaskListGroupCounter:         data.TaskListGroupCounter,
			OperationSystemCondition:     data.OperationSystemCondition,
			NumberOfTaskListExecutions:   data.NumberOfTaskListExecutions,
			MaintNotifTskIsAutomlyDtmnd:  data.MaintNotifTskIsAutomlyDtmnd,
			MaintenancePlanningPlant:     data.MaintenancePlanningPlant,
			MaintenancePlannerGroup:      data.MaintenancePlannerGroup,
			MaintenanceOrderType:         data.MaintenanceOrderType,
			NotificationType:             data.NotificationType,
			MaintenanceActivityType:      data.MaintenanceActivityType,
			MainWorkCenter:               data.MainWorkCenter,
			MainWorkCenterPlant:          data.MainWorkCenterPlant,
			MaintPriority:                data.MaintPriority,
			MaintPriorityType:            data.MaintPriorityType,
			BusinessArea:                 data.BusinessArea,
			ImmediateReleaseIsBlocked:    data.ImmediateReleaseIsBlocked,
			Material:                     data.Material,
			SerialNumber:                 data.SerialNumber,
			ServiceDocumentType:          data.ServiceDocumentType,
			ServiceContract:              data.ServiceContract,
			ServiceContractItem:          data.ServiceContractItem,
			ServiceOrderTemplate:         data.ServiceOrderTemplate,
			ServiceDocumentPriority:      data.ServiceDocumentPriority,
			Product:                      data.Product,
			MaintenancePlant:             data.MaintenancePlant,
			AssetLocation:                data.AssetLocation,
			AssetRoom:                    data.AssetRoom,
			PlantSection:                 data.PlantSection,
			WorkCenter:                   data.WorkCenter,
			ABCIndicator:                 data.ABCIndicator,
			MaintObjectFreeDefinedAttrib: data.MaintObjectFreeDefinedAttrib,
			CompanyCode:                  data.CompanyCode,
			MasterFixedAsset:             data.MasterFixedAsset,
			FixedAsset:                   data.FixedAsset,
			LocAcctAssgmtBusinessArea:    data.LocAcctAssgmtBusinessArea,
			CostCenter:                   data.CostCenter,
			ControllingArea:              data.ControllingArea,
			WBSElement:                   data.WBSElement,
			SettlementOrder:              data.SettlementOrder,
			CycleSetSequence:             data.CycleSetSequence,
			StandingOrderNumber:          data.StandingOrderNumber,
			CreationDate:                 data.CreationDate,
			LastChangeDate:               data.LastChangeDate,
			LastChangeDateTime:           data.LastChangeDateTime,
			ToCallObjects:                data.ToCallObjects.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToCallObjects(raw []byte, l *logger.Logger) ([]ToCallObjects, error) {
	pm := &responses.ToCallObjects{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToCallObjects. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toCallObjects := make([]ToCallObjects, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toCallObjects = append(toCallObjects, ToCallObjects{
			MaintenancePlan:              data.MaintenancePlan,
			MaintenancePlanCallNumber:    data.MaintenancePlanCallNumber,
			MaintenanceItem:              data.MaintenanceItem,
			MaintenanceOrder:             data.MaintenanceOrder,
			MaintenanceNotification:      data.MaintenanceNotification,
			ServiceOrder:                 data.ServiceOrder,
			MaintCallHorizonIsNotReached: data.MaintCallHorizonIsNotReached,
			SchedulingStatus:             data.SchedulingStatus,
			PlannedStartDate:             data.PlannedStartDate,
		})
	}

	return toCallObjects, nil
}
