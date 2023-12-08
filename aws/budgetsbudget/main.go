// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudget",
		reflect.TypeOf((*BudgetsBudget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "autoAdjustData", GoGetter: "AutoAdjustData"},
			_jsii_.MemberProperty{JsiiProperty: "autoAdjustDataInput", GoGetter: "AutoAdjustDataInput"},
			_jsii_.MemberProperty{JsiiProperty: "budgetType", GoGetter: "BudgetType"},
			_jsii_.MemberProperty{JsiiProperty: "budgetTypeInput", GoGetter: "BudgetTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "costFilter", GoGetter: "CostFilter"},
			_jsii_.MemberProperty{JsiiProperty: "costFilterInput", GoGetter: "CostFilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "costTypes", GoGetter: "CostTypes"},
			_jsii_.MemberProperty{JsiiProperty: "costTypesInput", GoGetter: "CostTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "limitAmount", GoGetter: "LimitAmount"},
			_jsii_.MemberProperty{JsiiProperty: "limitAmountInput", GoGetter: "LimitAmountInput"},
			_jsii_.MemberProperty{JsiiProperty: "limitUnit", GoGetter: "LimitUnit"},
			_jsii_.MemberProperty{JsiiProperty: "limitUnitInput", GoGetter: "LimitUnitInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namePrefix", GoGetter: "NamePrefix"},
			_jsii_.MemberProperty{JsiiProperty: "namePrefixInput", GoGetter: "NamePrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notification", GoGetter: "Notification"},
			_jsii_.MemberProperty{JsiiProperty: "notificationInput", GoGetter: "NotificationInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "plannedLimit", GoGetter: "PlannedLimit"},
			_jsii_.MemberProperty{JsiiProperty: "plannedLimitInput", GoGetter: "PlannedLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAutoAdjustData", GoMethod: "PutAutoAdjustData"},
			_jsii_.MemberMethod{JsiiMethod: "putCostFilter", GoMethod: "PutCostFilter"},
			_jsii_.MemberMethod{JsiiMethod: "putCostTypes", GoMethod: "PutCostTypes"},
			_jsii_.MemberMethod{JsiiMethod: "putNotification", GoMethod: "PutNotification"},
			_jsii_.MemberMethod{JsiiMethod: "putPlannedLimit", GoMethod: "PutPlannedLimit"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountId", GoMethod: "ResetAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoAdjustData", GoMethod: "ResetAutoAdjustData"},
			_jsii_.MemberMethod{JsiiMethod: "resetCostFilter", GoMethod: "ResetCostFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetCostTypes", GoMethod: "ResetCostTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLimitAmount", GoMethod: "ResetLimitAmount"},
			_jsii_.MemberMethod{JsiiMethod: "resetLimitUnit", GoMethod: "ResetLimitUnit"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamePrefix", GoMethod: "ResetNamePrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotification", GoMethod: "ResetNotification"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlannedLimit", GoMethod: "ResetPlannedLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimePeriodEnd", GoMethod: "ResetTimePeriodEnd"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimePeriodStart", GoMethod: "ResetTimePeriodStart"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timePeriodEnd", GoGetter: "TimePeriodEnd"},
			_jsii_.MemberProperty{JsiiProperty: "timePeriodEndInput", GoGetter: "TimePeriodEndInput"},
			_jsii_.MemberProperty{JsiiProperty: "timePeriodStart", GoGetter: "TimePeriodStart"},
			_jsii_.MemberProperty{JsiiProperty: "timePeriodStartInput", GoGetter: "TimePeriodStartInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeUnit", GoGetter: "TimeUnit"},
			_jsii_.MemberProperty{JsiiProperty: "timeUnitInput", GoGetter: "TimeUnitInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudget{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetAutoAdjustData",
		reflect.TypeOf((*BudgetsBudgetAutoAdjustData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetAutoAdjustDataHistoricalOptions",
		reflect.TypeOf((*BudgetsBudgetAutoAdjustDataHistoricalOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetAutoAdjustDataHistoricalOptionsOutputReference",
		reflect.TypeOf((*BudgetsBudgetAutoAdjustDataHistoricalOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "budgetAdjustmentPeriod", GoGetter: "BudgetAdjustmentPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "budgetAdjustmentPeriodInput", GoGetter: "BudgetAdjustmentPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lookbackAvailablePeriods", GoGetter: "LookbackAvailablePeriods"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetAutoAdjustDataHistoricalOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetAutoAdjustDataOutputReference",
		reflect.TypeOf((*BudgetsBudgetAutoAdjustDataOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoAdjustType", GoGetter: "AutoAdjustType"},
			_jsii_.MemberProperty{JsiiProperty: "autoAdjustTypeInput", GoGetter: "AutoAdjustTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "historicalOptions", GoGetter: "HistoricalOptions"},
			_jsii_.MemberProperty{JsiiProperty: "historicalOptionsInput", GoGetter: "HistoricalOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastAutoAdjustTime", GoGetter: "LastAutoAdjustTime"},
			_jsii_.MemberMethod{JsiiMethod: "putHistoricalOptions", GoMethod: "PutHistoricalOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetHistoricalOptions", GoMethod: "ResetHistoricalOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetConfig",
		reflect.TypeOf((*BudgetsBudgetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetCostFilter",
		reflect.TypeOf((*BudgetsBudgetCostFilter)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetCostFilterList",
		reflect.TypeOf((*BudgetsBudgetCostFilterList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetCostFilterList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetCostFilterOutputReference",
		reflect.TypeOf((*BudgetsBudgetCostFilterOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
			_jsii_.MemberProperty{JsiiProperty: "valuesInput", GoGetter: "ValuesInput"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetCostFilterOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetCostTypes",
		reflect.TypeOf((*BudgetsBudgetCostTypes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetCostTypesOutputReference",
		reflect.TypeOf((*BudgetsBudgetCostTypesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "includeCredit", GoGetter: "IncludeCredit"},
			_jsii_.MemberProperty{JsiiProperty: "includeCreditInput", GoGetter: "IncludeCreditInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeDiscount", GoGetter: "IncludeDiscount"},
			_jsii_.MemberProperty{JsiiProperty: "includeDiscountInput", GoGetter: "IncludeDiscountInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeOtherSubscription", GoGetter: "IncludeOtherSubscription"},
			_jsii_.MemberProperty{JsiiProperty: "includeOtherSubscriptionInput", GoGetter: "IncludeOtherSubscriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeRecurring", GoGetter: "IncludeRecurring"},
			_jsii_.MemberProperty{JsiiProperty: "includeRecurringInput", GoGetter: "IncludeRecurringInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeRefund", GoGetter: "IncludeRefund"},
			_jsii_.MemberProperty{JsiiProperty: "includeRefundInput", GoGetter: "IncludeRefundInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeSubscription", GoGetter: "IncludeSubscription"},
			_jsii_.MemberProperty{JsiiProperty: "includeSubscriptionInput", GoGetter: "IncludeSubscriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeSupport", GoGetter: "IncludeSupport"},
			_jsii_.MemberProperty{JsiiProperty: "includeSupportInput", GoGetter: "IncludeSupportInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeTax", GoGetter: "IncludeTax"},
			_jsii_.MemberProperty{JsiiProperty: "includeTaxInput", GoGetter: "IncludeTaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeUpfront", GoGetter: "IncludeUpfront"},
			_jsii_.MemberProperty{JsiiProperty: "includeUpfrontInput", GoGetter: "IncludeUpfrontInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeCredit", GoMethod: "ResetIncludeCredit"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeDiscount", GoMethod: "ResetIncludeDiscount"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeOtherSubscription", GoMethod: "ResetIncludeOtherSubscription"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeRecurring", GoMethod: "ResetIncludeRecurring"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeRefund", GoMethod: "ResetIncludeRefund"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeSubscription", GoMethod: "ResetIncludeSubscription"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeSupport", GoMethod: "ResetIncludeSupport"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeTax", GoMethod: "ResetIncludeTax"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeUpfront", GoMethod: "ResetIncludeUpfront"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseAmortized", GoMethod: "ResetUseAmortized"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseBlended", GoMethod: "ResetUseBlended"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "useAmortized", GoGetter: "UseAmortized"},
			_jsii_.MemberProperty{JsiiProperty: "useAmortizedInput", GoGetter: "UseAmortizedInput"},
			_jsii_.MemberProperty{JsiiProperty: "useBlended", GoGetter: "UseBlended"},
			_jsii_.MemberProperty{JsiiProperty: "useBlendedInput", GoGetter: "UseBlendedInput"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetCostTypesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetNotification",
		reflect.TypeOf((*BudgetsBudgetNotification)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetNotificationList",
		reflect.TypeOf((*BudgetsBudgetNotificationList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetNotificationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetNotificationOutputReference",
		reflect.TypeOf((*BudgetsBudgetNotificationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "comparisonOperator", GoGetter: "ComparisonOperator"},
			_jsii_.MemberProperty{JsiiProperty: "comparisonOperatorInput", GoGetter: "ComparisonOperatorInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "notificationType", GoGetter: "NotificationType"},
			_jsii_.MemberProperty{JsiiProperty: "notificationTypeInput", GoGetter: "NotificationTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubscriberEmailAddresses", GoMethod: "ResetSubscriberEmailAddresses"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubscriberSnsTopicArns", GoMethod: "ResetSubscriberSnsTopicArns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "subscriberEmailAddresses", GoGetter: "SubscriberEmailAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "subscriberEmailAddressesInput", GoGetter: "SubscriberEmailAddressesInput"},
			_jsii_.MemberProperty{JsiiProperty: "subscriberSnsTopicArns", GoGetter: "SubscriberSnsTopicArns"},
			_jsii_.MemberProperty{JsiiProperty: "subscriberSnsTopicArnsInput", GoGetter: "SubscriberSnsTopicArnsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdInput", GoGetter: "ThresholdInput"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdType", GoGetter: "ThresholdType"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdTypeInput", GoGetter: "ThresholdTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetNotificationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetPlannedLimit",
		reflect.TypeOf((*BudgetsBudgetPlannedLimit)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetPlannedLimitList",
		reflect.TypeOf((*BudgetsBudgetPlannedLimitList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetPlannedLimitList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetPlannedLimitOutputReference",
		reflect.TypeOf((*BudgetsBudgetPlannedLimitOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amount", GoGetter: "Amount"},
			_jsii_.MemberProperty{JsiiProperty: "amountInput", GoGetter: "AmountInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "startTime", GoGetter: "StartTime"},
			_jsii_.MemberProperty{JsiiProperty: "startTimeInput", GoGetter: "StartTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetPlannedLimitOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
