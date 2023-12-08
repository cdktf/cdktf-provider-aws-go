// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtask

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTask",
		reflect.TypeOf((*SsmMaintenanceWindowTask)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "cutoffBehavior", GoGetter: "CutoffBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "cutoffBehaviorInput", GoGetter: "CutoffBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrency", GoGetter: "MaxConcurrency"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrencyInput", GoGetter: "MaxConcurrencyInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxErrors", GoGetter: "MaxErrors"},
			_jsii_.MemberProperty{JsiiProperty: "maxErrorsInput", GoGetter: "MaxErrorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "priorityInput", GoGetter: "PriorityInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putTargets", GoMethod: "PutTargets"},
			_jsii_.MemberMethod{JsiiMethod: "putTaskInvocationParameters", GoMethod: "PutTaskInvocationParameters"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCutoffBehavior", GoMethod: "ResetCutoffBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxConcurrency", GoMethod: "ResetMaxConcurrency"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxErrors", GoMethod: "ResetMaxErrors"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPriority", GoMethod: "ResetPriority"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceRoleArn", GoMethod: "ResetServiceRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargets", GoMethod: "ResetTargets"},
			_jsii_.MemberMethod{JsiiMethod: "resetTaskInvocationParameters", GoMethod: "ResetTaskInvocationParameters"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRoleArn", GoGetter: "ServiceRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRoleArnInput", GoGetter: "ServiceRoleArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "targets", GoGetter: "Targets"},
			_jsii_.MemberProperty{JsiiProperty: "targetsInput", GoGetter: "TargetsInput"},
			_jsii_.MemberProperty{JsiiProperty: "taskArn", GoGetter: "TaskArn"},
			_jsii_.MemberProperty{JsiiProperty: "taskArnInput", GoGetter: "TaskArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "taskInvocationParameters", GoGetter: "TaskInvocationParameters"},
			_jsii_.MemberProperty{JsiiProperty: "taskInvocationParametersInput", GoGetter: "TaskInvocationParametersInput"},
			_jsii_.MemberProperty{JsiiProperty: "taskType", GoGetter: "TaskType"},
			_jsii_.MemberProperty{JsiiProperty: "taskTypeInput", GoGetter: "TaskTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "windowId", GoGetter: "WindowId"},
			_jsii_.MemberProperty{JsiiProperty: "windowIdInput", GoGetter: "WindowIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "windowTaskId", GoGetter: "WindowTaskId"},
		},
		func() interface{} {
			j := jsiiProxy_SsmMaintenanceWindowTask{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskConfig",
		reflect.TypeOf((*SsmMaintenanceWindowTaskConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTargets",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTargets)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTargetsList",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTargetsList)(nil)).Elem(),
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
			j := jsiiProxy_SsmMaintenanceWindowTaskTargetsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTargetsOutputReference",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTargetsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
			_jsii_.MemberProperty{JsiiProperty: "valuesInput", GoGetter: "ValuesInput"},
		},
		func() interface{} {
			j := jsiiProxy_SsmMaintenanceWindowTaskTargetsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParameters",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "documentVersion", GoGetter: "DocumentVersion"},
			_jsii_.MemberProperty{JsiiProperty: "documentVersionInput", GoGetter: "DocumentVersionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "parameter", GoGetter: "Parameter"},
			_jsii_.MemberProperty{JsiiProperty: "parameterInput", GoGetter: "ParameterInput"},
			_jsii_.MemberMethod{JsiiMethod: "putParameter", GoMethod: "PutParameter"},
			_jsii_.MemberMethod{JsiiMethod: "resetDocumentVersion", GoMethod: "ResetDocumentVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameter", GoMethod: "ResetParameter"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameter",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameter)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameterList",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameterList)(nil)).Elem(),
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
			j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameterList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameterOutputReference",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameterOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameterOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clientContext", GoGetter: "ClientContext"},
			_jsii_.MemberProperty{JsiiProperty: "clientContextInput", GoGetter: "ClientContextInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "payload", GoGetter: "Payload"},
			_jsii_.MemberProperty{JsiiProperty: "payloadInput", GoGetter: "PayloadInput"},
			_jsii_.MemberProperty{JsiiProperty: "qualifier", GoGetter: "Qualifier"},
			_jsii_.MemberProperty{JsiiProperty: "qualifierInput", GoGetter: "QualifierInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientContext", GoMethod: "ResetClientContext"},
			_jsii_.MemberMethod{JsiiMethod: "resetPayload", GoMethod: "ResetPayload"},
			_jsii_.MemberMethod{JsiiMethod: "resetQualifier", GoMethod: "ResetQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "automationParameters", GoGetter: "AutomationParameters"},
			_jsii_.MemberProperty{JsiiProperty: "automationParametersInput", GoGetter: "AutomationParametersInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lambdaParameters", GoGetter: "LambdaParameters"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaParametersInput", GoGetter: "LambdaParametersInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAutomationParameters", GoMethod: "PutAutomationParameters"},
			_jsii_.MemberMethod{JsiiMethod: "putLambdaParameters", GoMethod: "PutLambdaParameters"},
			_jsii_.MemberMethod{JsiiMethod: "putRunCommandParameters", GoMethod: "PutRunCommandParameters"},
			_jsii_.MemberMethod{JsiiMethod: "putStepFunctionsParameters", GoMethod: "PutStepFunctionsParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutomationParameters", GoMethod: "ResetAutomationParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetLambdaParameters", GoMethod: "ResetLambdaParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetRunCommandParameters", GoMethod: "ResetRunCommandParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetStepFunctionsParameters", GoMethod: "ResetStepFunctionsParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "runCommandParameters", GoGetter: "RunCommandParameters"},
			_jsii_.MemberProperty{JsiiProperty: "runCommandParametersInput", GoGetter: "RunCommandParametersInput"},
			_jsii_.MemberProperty{JsiiProperty: "stepFunctionsParameters", GoGetter: "StepFunctionsParameters"},
			_jsii_.MemberProperty{JsiiProperty: "stepFunctionsParametersInput", GoGetter: "StepFunctionsParametersInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLogGroupName", GoGetter: "CloudwatchLogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLogGroupNameInput", GoGetter: "CloudwatchLogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchOutputEnabled", GoGetter: "CloudwatchOutputEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchOutputEnabledInput", GoGetter: "CloudwatchOutputEnabledInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLogGroupName", GoMethod: "ResetCloudwatchLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchOutputEnabled", GoMethod: "ResetCloudwatchOutputEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "notificationArn", GoGetter: "NotificationArn"},
			_jsii_.MemberProperty{JsiiProperty: "notificationArnInput", GoGetter: "NotificationArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "notificationEvents", GoGetter: "NotificationEvents"},
			_jsii_.MemberProperty{JsiiProperty: "notificationEventsInput", GoGetter: "NotificationEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "notificationType", GoGetter: "NotificationType"},
			_jsii_.MemberProperty{JsiiProperty: "notificationTypeInput", GoGetter: "NotificationTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotificationArn", GoMethod: "ResetNotificationArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotificationEvents", GoMethod: "ResetNotificationEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotificationType", GoMethod: "ResetNotificationType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchConfig", GoGetter: "CloudwatchConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchConfigInput", GoGetter: "CloudwatchConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "documentHash", GoGetter: "DocumentHash"},
			_jsii_.MemberProperty{JsiiProperty: "documentHashInput", GoGetter: "DocumentHashInput"},
			_jsii_.MemberProperty{JsiiProperty: "documentHashType", GoGetter: "DocumentHashType"},
			_jsii_.MemberProperty{JsiiProperty: "documentHashTypeInput", GoGetter: "DocumentHashTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "documentVersion", GoGetter: "DocumentVersion"},
			_jsii_.MemberProperty{JsiiProperty: "documentVersionInput", GoGetter: "DocumentVersionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "notificationConfig", GoGetter: "NotificationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "notificationConfigInput", GoGetter: "NotificationConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "outputS3Bucket", GoGetter: "OutputS3Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "outputS3BucketInput", GoGetter: "OutputS3BucketInput"},
			_jsii_.MemberProperty{JsiiProperty: "outputS3KeyPrefix", GoGetter: "OutputS3KeyPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "outputS3KeyPrefixInput", GoGetter: "OutputS3KeyPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "parameter", GoGetter: "Parameter"},
			_jsii_.MemberProperty{JsiiProperty: "parameterInput", GoGetter: "ParameterInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchConfig", GoMethod: "PutCloudwatchConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putNotificationConfig", GoMethod: "PutNotificationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putParameter", GoMethod: "PutParameter"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchConfig", GoMethod: "ResetCloudwatchConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetDocumentHash", GoMethod: "ResetDocumentHash"},
			_jsii_.MemberMethod{JsiiMethod: "resetDocumentHashType", GoMethod: "ResetDocumentHashType"},
			_jsii_.MemberMethod{JsiiMethod: "resetDocumentVersion", GoMethod: "ResetDocumentVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotificationConfig", GoMethod: "ResetNotificationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutputS3Bucket", GoMethod: "ResetOutputS3Bucket"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutputS3KeyPrefix", GoMethod: "ResetOutputS3KeyPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameter", GoMethod: "ResetParameter"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceRoleArn", GoMethod: "ResetServiceRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeoutSeconds", GoMethod: "ResetTimeoutSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRoleArn", GoGetter: "ServiceRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRoleArnInput", GoGetter: "ServiceRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutSeconds", GoGetter: "TimeoutSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutSecondsInput", GoGetter: "TimeoutSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameter",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameter)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterList",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterList)(nil)).Elem(),
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
			j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterOutputReference",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference",
		reflect.TypeOf((*SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "input", GoGetter: "Input"},
			_jsii_.MemberProperty{JsiiProperty: "inputInput", GoGetter: "InputInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInput", GoMethod: "ResetInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
