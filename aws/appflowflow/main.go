// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlow",
		reflect.TypeOf((*AppflowFlow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "destinationFlowConfig", GoGetter: "DestinationFlowConfig"},
			_jsii_.MemberProperty{JsiiProperty: "destinationFlowConfigInput", GoGetter: "DestinationFlowConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "flowStatus", GoGetter: "FlowStatus"},
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
			_jsii_.MemberProperty{JsiiProperty: "kmsArn", GoGetter: "KmsArn"},
			_jsii_.MemberProperty{JsiiProperty: "kmsArnInput", GoGetter: "KmsArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "metadataCatalogConfig", GoGetter: "MetadataCatalogConfig"},
			_jsii_.MemberProperty{JsiiProperty: "metadataCatalogConfigInput", GoGetter: "MetadataCatalogConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putDestinationFlowConfig", GoMethod: "PutDestinationFlowConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putMetadataCatalogConfig", GoMethod: "PutMetadataCatalogConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSourceFlowConfig", GoMethod: "PutSourceFlowConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putTask", GoMethod: "PutTask"},
			_jsii_.MemberMethod{JsiiMethod: "putTriggerConfig", GoMethod: "PutTriggerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsArn", GoMethod: "ResetKmsArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadataCatalogConfig", GoMethod: "ResetMetadataCatalogConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "sourceFlowConfig", GoGetter: "SourceFlowConfig"},
			_jsii_.MemberProperty{JsiiProperty: "sourceFlowConfigInput", GoGetter: "SourceFlowConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "task", GoGetter: "Task"},
			_jsii_.MemberProperty{JsiiProperty: "taskInput", GoGetter: "TaskInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "triggerConfig", GoGetter: "TriggerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "triggerConfigInput", GoGetter: "TriggerConfigInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlow{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowConfig",
		reflect.TypeOf((*AppflowFlowConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorProperties",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorProperties)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnector",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorErrorHandlingConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorErrorHandlingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorErrorHandlingConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorErrorHandlingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationError", GoGetter: "FailOnFirstDestinationError"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationErrorInput", GoGetter: "FailOnFirstDestinationErrorInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetBucketName", GoMethod: "ResetBucketName"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailOnFirstDestinationError", GoMethod: "ResetFailOnFirstDestinationError"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorErrorHandlingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customProperties", GoGetter: "CustomProperties"},
			_jsii_.MemberProperty{JsiiProperty: "customPropertiesInput", GoGetter: "CustomPropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "entityName", GoGetter: "EntityName"},
			_jsii_.MemberProperty{JsiiProperty: "entityNameInput", GoGetter: "EntityNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfig", GoGetter: "ErrorHandlingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfigInput", GoGetter: "ErrorHandlingConfigInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idFieldNames", GoGetter: "IdFieldNames"},
			_jsii_.MemberProperty{JsiiProperty: "idFieldNamesInput", GoGetter: "IdFieldNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putErrorHandlingConfig", GoMethod: "PutErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomProperties", GoMethod: "ResetCustomProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorHandlingConfig", GoMethod: "ResetErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdFieldNames", GoMethod: "ResetIdFieldNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetWriteOperationType", GoMethod: "ResetWriteOperationType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "writeOperationType", GoGetter: "WriteOperationType"},
			_jsii_.MemberProperty{JsiiProperty: "writeOperationTypeInput", GoGetter: "WriteOperationTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfilesOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfilesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "domainNameInput", GoGetter: "DomainNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "objectTypeName", GoGetter: "ObjectTypeName"},
			_jsii_.MemberProperty{JsiiProperty: "objectTypeNameInput", GoGetter: "ObjectTypeNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetObjectTypeName", GoMethod: "ResetObjectTypeName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfilesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridge",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridge)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridgeErrorHandlingConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridgeErrorHandlingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridgeErrorHandlingConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridgeErrorHandlingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationError", GoGetter: "FailOnFirstDestinationError"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationErrorInput", GoGetter: "FailOnFirstDestinationErrorInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetBucketName", GoMethod: "ResetBucketName"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailOnFirstDestinationError", GoMethod: "ResetFailOnFirstDestinationError"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridgeErrorHandlingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridgeOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridgeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfig", GoGetter: "ErrorHandlingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfigInput", GoGetter: "ErrorHandlingConfigInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "putErrorHandlingConfig", GoMethod: "PutErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorHandlingConfig", GoMethod: "ResetErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridgeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycode",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycode)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycodeErrorHandlingConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycodeErrorHandlingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycodeErrorHandlingConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycodeErrorHandlingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationError", GoGetter: "FailOnFirstDestinationError"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationErrorInput", GoGetter: "FailOnFirstDestinationErrorInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetBucketName", GoMethod: "ResetBucketName"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailOnFirstDestinationError", GoMethod: "ResetFailOnFirstDestinationError"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycodeErrorHandlingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycodeOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycodeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfig", GoGetter: "ErrorHandlingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfigInput", GoGetter: "ErrorHandlingConfigInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "putErrorHandlingConfig", GoMethod: "PutErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorHandlingConfig", GoMethod: "ResetErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycodeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetrics",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetrics)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetricsOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetricsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetricsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketo",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationError", GoGetter: "FailOnFirstDestinationError"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationErrorInput", GoGetter: "FailOnFirstDestinationErrorInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetBucketName", GoMethod: "ResetBucketName"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailOnFirstDestinationError", GoMethod: "ResetFailOnFirstDestinationError"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketoOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketoOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfig", GoGetter: "ErrorHandlingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfigInput", GoGetter: "ErrorHandlingConfigInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "putErrorHandlingConfig", GoMethod: "PutErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorHandlingConfig", GoMethod: "ResetErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customConnector", GoGetter: "CustomConnector"},
			_jsii_.MemberProperty{JsiiProperty: "customConnectorInput", GoGetter: "CustomConnectorInput"},
			_jsii_.MemberProperty{JsiiProperty: "customerProfiles", GoGetter: "CustomerProfiles"},
			_jsii_.MemberProperty{JsiiProperty: "customerProfilesInput", GoGetter: "CustomerProfilesInput"},
			_jsii_.MemberProperty{JsiiProperty: "eventBridge", GoGetter: "EventBridge"},
			_jsii_.MemberProperty{JsiiProperty: "eventBridgeInput", GoGetter: "EventBridgeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "honeycode", GoGetter: "Honeycode"},
			_jsii_.MemberProperty{JsiiProperty: "honeycodeInput", GoGetter: "HoneycodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lookoutMetrics", GoGetter: "LookoutMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "lookoutMetricsInput", GoGetter: "LookoutMetricsInput"},
			_jsii_.MemberProperty{JsiiProperty: "marketo", GoGetter: "Marketo"},
			_jsii_.MemberProperty{JsiiProperty: "marketoInput", GoGetter: "MarketoInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomConnector", GoMethod: "PutCustomConnector"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomerProfiles", GoMethod: "PutCustomerProfiles"},
			_jsii_.MemberMethod{JsiiMethod: "putEventBridge", GoMethod: "PutEventBridge"},
			_jsii_.MemberMethod{JsiiMethod: "putHoneycode", GoMethod: "PutHoneycode"},
			_jsii_.MemberMethod{JsiiMethod: "putLookoutMetrics", GoMethod: "PutLookoutMetrics"},
			_jsii_.MemberMethod{JsiiMethod: "putMarketo", GoMethod: "PutMarketo"},
			_jsii_.MemberMethod{JsiiMethod: "putRedshift", GoMethod: "PutRedshift"},
			_jsii_.MemberMethod{JsiiMethod: "putS3", GoMethod: "PutS3"},
			_jsii_.MemberMethod{JsiiMethod: "putSalesforce", GoMethod: "PutSalesforce"},
			_jsii_.MemberMethod{JsiiMethod: "putSapoData", GoMethod: "PutSapoData"},
			_jsii_.MemberMethod{JsiiMethod: "putSnowflake", GoMethod: "PutSnowflake"},
			_jsii_.MemberMethod{JsiiMethod: "putUpsolver", GoMethod: "PutUpsolver"},
			_jsii_.MemberMethod{JsiiMethod: "putZendesk", GoMethod: "PutZendesk"},
			_jsii_.MemberProperty{JsiiProperty: "redshift", GoGetter: "Redshift"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftInput", GoGetter: "RedshiftInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomConnector", GoMethod: "ResetCustomConnector"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomerProfiles", GoMethod: "ResetCustomerProfiles"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventBridge", GoMethod: "ResetEventBridge"},
			_jsii_.MemberMethod{JsiiMethod: "resetHoneycode", GoMethod: "ResetHoneycode"},
			_jsii_.MemberMethod{JsiiMethod: "resetLookoutMetrics", GoMethod: "ResetLookoutMetrics"},
			_jsii_.MemberMethod{JsiiMethod: "resetMarketo", GoMethod: "ResetMarketo"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedshift", GoMethod: "ResetRedshift"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3", GoMethod: "ResetS3"},
			_jsii_.MemberMethod{JsiiMethod: "resetSalesforce", GoMethod: "ResetSalesforce"},
			_jsii_.MemberMethod{JsiiMethod: "resetSapoData", GoMethod: "ResetSapoData"},
			_jsii_.MemberMethod{JsiiMethod: "resetSnowflake", GoMethod: "ResetSnowflake"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpsolver", GoMethod: "ResetUpsolver"},
			_jsii_.MemberMethod{JsiiMethod: "resetZendesk", GoMethod: "ResetZendesk"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3", GoGetter: "S3"},
			_jsii_.MemberProperty{JsiiProperty: "s3Input", GoGetter: "S3Input"},
			_jsii_.MemberProperty{JsiiProperty: "salesforce", GoGetter: "Salesforce"},
			_jsii_.MemberProperty{JsiiProperty: "salesforceInput", GoGetter: "SalesforceInput"},
			_jsii_.MemberProperty{JsiiProperty: "sapoData", GoGetter: "SapoData"},
			_jsii_.MemberProperty{JsiiProperty: "sapoDataInput", GoGetter: "SapoDataInput"},
			_jsii_.MemberProperty{JsiiProperty: "snowflake", GoGetter: "Snowflake"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeInput", GoGetter: "SnowflakeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "upsolver", GoGetter: "Upsolver"},
			_jsii_.MemberProperty{JsiiProperty: "upsolverInput", GoGetter: "UpsolverInput"},
			_jsii_.MemberProperty{JsiiProperty: "zendesk", GoGetter: "Zendesk"},
			_jsii_.MemberProperty{JsiiProperty: "zendeskInput", GoGetter: "ZendeskInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshift",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshift)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshiftErrorHandlingConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshiftErrorHandlingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshiftErrorHandlingConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshiftErrorHandlingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationError", GoGetter: "FailOnFirstDestinationError"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationErrorInput", GoGetter: "FailOnFirstDestinationErrorInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetBucketName", GoMethod: "ResetBucketName"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailOnFirstDestinationError", GoMethod: "ResetFailOnFirstDestinationError"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshiftErrorHandlingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshiftOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshiftOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfig", GoGetter: "ErrorHandlingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfigInput", GoGetter: "ErrorHandlingConfigInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "intermediateBucketName", GoGetter: "IntermediateBucketName"},
			_jsii_.MemberProperty{JsiiProperty: "intermediateBucketNameInput", GoGetter: "IntermediateBucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "putErrorHandlingConfig", GoMethod: "PutErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorHandlingConfig", GoMethod: "ResetErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshiftOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3OutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3OutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putS3OutputFormatConfig", GoMethod: "PutS3OutputFormatConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3OutputFormatConfig", GoMethod: "ResetS3OutputFormatConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputFormatConfig", GoGetter: "S3OutputFormatConfig"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputFormatConfigInput", GoGetter: "S3OutputFormatConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aggregationType", GoGetter: "AggregationType"},
			_jsii_.MemberProperty{JsiiProperty: "aggregationTypeInput", GoGetter: "AggregationTypeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAggregationType", GoMethod: "ResetAggregationType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetFileSize", GoMethod: "ResetTargetFileSize"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "targetFileSize", GoGetter: "TargetFileSize"},
			_jsii_.MemberProperty{JsiiProperty: "targetFileSizeInput", GoGetter: "TargetFileSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aggregationConfig", GoGetter: "AggregationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "aggregationConfigInput", GoGetter: "AggregationConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileType", GoGetter: "FileType"},
			_jsii_.MemberProperty{JsiiProperty: "fileTypeInput", GoGetter: "FileTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "prefixConfig", GoGetter: "PrefixConfig"},
			_jsii_.MemberProperty{JsiiProperty: "prefixConfigInput", GoGetter: "PrefixConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "preserveSourceDataTyping", GoGetter: "PreserveSourceDataTyping"},
			_jsii_.MemberProperty{JsiiProperty: "preserveSourceDataTypingInput", GoGetter: "PreserveSourceDataTypingInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAggregationConfig", GoMethod: "PutAggregationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putPrefixConfig", GoMethod: "PutPrefixConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAggregationConfig", GoMethod: "ResetAggregationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileType", GoMethod: "ResetFileType"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixConfig", GoMethod: "ResetPrefixConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreserveSourceDataTyping", GoMethod: "ResetPreserveSourceDataTyping"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "prefixFormat", GoGetter: "PrefixFormat"},
			_jsii_.MemberProperty{JsiiProperty: "prefixFormatInput", GoGetter: "PrefixFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "prefixHierarchy", GoGetter: "PrefixHierarchy"},
			_jsii_.MemberProperty{JsiiProperty: "prefixHierarchyInput", GoGetter: "PrefixHierarchyInput"},
			_jsii_.MemberProperty{JsiiProperty: "prefixType", GoGetter: "PrefixType"},
			_jsii_.MemberProperty{JsiiProperty: "prefixTypeInput", GoGetter: "PrefixTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixFormat", GoMethod: "ResetPrefixFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixHierarchy", GoMethod: "ResetPrefixHierarchy"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixType", GoMethod: "ResetPrefixType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationError", GoGetter: "FailOnFirstDestinationError"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationErrorInput", GoGetter: "FailOnFirstDestinationErrorInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetBucketName", GoMethod: "ResetBucketName"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailOnFirstDestinationError", GoMethod: "ResetFailOnFirstDestinationError"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfig", GoGetter: "ErrorHandlingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfigInput", GoGetter: "ErrorHandlingConfigInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idFieldNames", GoGetter: "IdFieldNames"},
			_jsii_.MemberProperty{JsiiProperty: "idFieldNamesInput", GoGetter: "IdFieldNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "putErrorHandlingConfig", GoMethod: "PutErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorHandlingConfig", GoMethod: "ResetErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdFieldNames", GoMethod: "ResetIdFieldNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetWriteOperationType", GoMethod: "ResetWriteOperationType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "writeOperationType", GoGetter: "WriteOperationType"},
			_jsii_.MemberProperty{JsiiProperty: "writeOperationTypeInput", GoGetter: "WriteOperationTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoData",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationError", GoGetter: "FailOnFirstDestinationError"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationErrorInput", GoGetter: "FailOnFirstDestinationErrorInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetBucketName", GoMethod: "ResetBucketName"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailOnFirstDestinationError", GoMethod: "ResetFailOnFirstDestinationError"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfig", GoGetter: "ErrorHandlingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfigInput", GoGetter: "ErrorHandlingConfigInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idFieldNames", GoGetter: "IdFieldNames"},
			_jsii_.MemberProperty{JsiiProperty: "idFieldNamesInput", GoGetter: "IdFieldNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "objectPath", GoGetter: "ObjectPath"},
			_jsii_.MemberProperty{JsiiProperty: "objectPathInput", GoGetter: "ObjectPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "putErrorHandlingConfig", GoMethod: "PutErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSuccessResponseHandlingConfig", GoMethod: "PutSuccessResponseHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorHandlingConfig", GoMethod: "ResetErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdFieldNames", GoMethod: "ResetIdFieldNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuccessResponseHandlingConfig", GoMethod: "ResetSuccessResponseHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetWriteOperationType", GoMethod: "ResetWriteOperationType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "successResponseHandlingConfig", GoGetter: "SuccessResponseHandlingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "successResponseHandlingConfigInput", GoGetter: "SuccessResponseHandlingConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "writeOperationType", GoGetter: "WriteOperationType"},
			_jsii_.MemberProperty{JsiiProperty: "writeOperationTypeInput", GoGetter: "WriteOperationTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataSuccessResponseHandlingConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataSuccessResponseHandlingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataSuccessResponseHandlingConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataSuccessResponseHandlingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetBucketName", GoMethod: "ResetBucketName"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataSuccessResponseHandlingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflake",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflake)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeErrorHandlingConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeErrorHandlingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationError", GoGetter: "FailOnFirstDestinationError"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationErrorInput", GoGetter: "FailOnFirstDestinationErrorInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetBucketName", GoMethod: "ResetBucketName"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailOnFirstDestinationError", GoMethod: "ResetFailOnFirstDestinationError"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfig", GoGetter: "ErrorHandlingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfigInput", GoGetter: "ErrorHandlingConfigInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "intermediateBucketName", GoGetter: "IntermediateBucketName"},
			_jsii_.MemberProperty{JsiiProperty: "intermediateBucketNameInput", GoGetter: "IntermediateBucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "putErrorHandlingConfig", GoMethod: "PutErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorHandlingConfig", GoMethod: "ResetErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolver",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolver)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putS3OutputFormatConfig", GoMethod: "PutS3OutputFormatConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputFormatConfig", GoGetter: "S3OutputFormatConfig"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputFormatConfigInput", GoGetter: "S3OutputFormatConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aggregationType", GoGetter: "AggregationType"},
			_jsii_.MemberProperty{JsiiProperty: "aggregationTypeInput", GoGetter: "AggregationTypeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAggregationType", GoMethod: "ResetAggregationType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aggregationConfig", GoGetter: "AggregationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "aggregationConfigInput", GoGetter: "AggregationConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileType", GoGetter: "FileType"},
			_jsii_.MemberProperty{JsiiProperty: "fileTypeInput", GoGetter: "FileTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "prefixConfig", GoGetter: "PrefixConfig"},
			_jsii_.MemberProperty{JsiiProperty: "prefixConfigInput", GoGetter: "PrefixConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAggregationConfig", GoMethod: "PutAggregationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putPrefixConfig", GoMethod: "PutPrefixConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAggregationConfig", GoMethod: "ResetAggregationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileType", GoMethod: "ResetFileType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "prefixFormat", GoGetter: "PrefixFormat"},
			_jsii_.MemberProperty{JsiiProperty: "prefixFormatInput", GoGetter: "PrefixFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "prefixHierarchy", GoGetter: "PrefixHierarchy"},
			_jsii_.MemberProperty{JsiiProperty: "prefixHierarchyInput", GoGetter: "PrefixHierarchyInput"},
			_jsii_.MemberProperty{JsiiProperty: "prefixType", GoGetter: "PrefixType"},
			_jsii_.MemberProperty{JsiiProperty: "prefixTypeInput", GoGetter: "PrefixTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixFormat", GoMethod: "ResetPrefixFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixHierarchy", GoMethod: "ResetPrefixHierarchy"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendesk",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendesk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfig",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationError", GoGetter: "FailOnFirstDestinationError"},
			_jsii_.MemberProperty{JsiiProperty: "failOnFirstDestinationErrorInput", GoGetter: "FailOnFirstDestinationErrorInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetBucketName", GoMethod: "ResetBucketName"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailOnFirstDestinationError", GoMethod: "ResetFailOnFirstDestinationError"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendeskOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendeskOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfig", GoGetter: "ErrorHandlingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "errorHandlingConfigInput", GoGetter: "ErrorHandlingConfigInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idFieldNames", GoGetter: "IdFieldNames"},
			_jsii_.MemberProperty{JsiiProperty: "idFieldNamesInput", GoGetter: "IdFieldNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "putErrorHandlingConfig", GoMethod: "PutErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorHandlingConfig", GoMethod: "ResetErrorHandlingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdFieldNames", GoMethod: "ResetIdFieldNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetWriteOperationType", GoMethod: "ResetWriteOperationType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "writeOperationType", GoGetter: "WriteOperationType"},
			_jsii_.MemberProperty{JsiiProperty: "writeOperationTypeInput", GoGetter: "WriteOperationTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendeskOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigList",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_AppflowFlowDestinationFlowConfigList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigOutputReference",
		reflect.TypeOf((*AppflowFlowDestinationFlowConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersionInput", GoGetter: "ApiVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "connectorProfileName", GoGetter: "ConnectorProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "connectorProfileNameInput", GoGetter: "ConnectorProfileNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
			_jsii_.MemberProperty{JsiiProperty: "connectorTypeInput", GoGetter: "ConnectorTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinationConnectorProperties", GoGetter: "DestinationConnectorProperties"},
			_jsii_.MemberProperty{JsiiProperty: "destinationConnectorPropertiesInput", GoGetter: "DestinationConnectorPropertiesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putDestinationConnectorProperties", GoMethod: "PutDestinationConnectorProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiVersion", GoMethod: "ResetApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectorProfileName", GoMethod: "ResetConnectorProfileName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowDestinationFlowConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowMetadataCatalogConfig",
		reflect.TypeOf((*AppflowFlowMetadataCatalogConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowMetadataCatalogConfigGlueDataCatalog",
		reflect.TypeOf((*AppflowFlowMetadataCatalogConfigGlueDataCatalog)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowMetadataCatalogConfigGlueDataCatalogOutputReference",
		reflect.TypeOf((*AppflowFlowMetadataCatalogConfigGlueDataCatalogOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "databaseNameInput", GoGetter: "DatabaseNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "tablePrefix", GoGetter: "TablePrefix"},
			_jsii_.MemberProperty{JsiiProperty: "tablePrefixInput", GoGetter: "TablePrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowMetadataCatalogConfigGlueDataCatalogOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowMetadataCatalogConfigOutputReference",
		reflect.TypeOf((*AppflowFlowMetadataCatalogConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "glueDataCatalog", GoGetter: "GlueDataCatalog"},
			_jsii_.MemberProperty{JsiiProperty: "glueDataCatalogInput", GoGetter: "GlueDataCatalogInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putGlueDataCatalog", GoMethod: "PutGlueDataCatalog"},
			_jsii_.MemberMethod{JsiiMethod: "resetGlueDataCatalog", GoMethod: "ResetGlueDataCatalog"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowMetadataCatalogConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfig",
		reflect.TypeOf((*AppflowFlowSourceFlowConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigIncrementalPullConfig",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigIncrementalPullConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigIncrementalPullConfigOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigIncrementalPullConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "datetimeTypeFieldName", GoGetter: "DatetimeTypeFieldName"},
			_jsii_.MemberProperty{JsiiProperty: "datetimeTypeFieldNameInput", GoGetter: "DatetimeTypeFieldNameInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDatetimeTypeFieldName", GoMethod: "ResetDatetimeTypeFieldName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigIncrementalPullConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersionInput", GoGetter: "ApiVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "connectorProfileName", GoGetter: "ConnectorProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "connectorProfileNameInput", GoGetter: "ConnectorProfileNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
			_jsii_.MemberProperty{JsiiProperty: "connectorTypeInput", GoGetter: "ConnectorTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "incrementalPullConfig", GoGetter: "IncrementalPullConfig"},
			_jsii_.MemberProperty{JsiiProperty: "incrementalPullConfigInput", GoGetter: "IncrementalPullConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIncrementalPullConfig", GoMethod: "PutIncrementalPullConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSourceConnectorProperties", GoMethod: "PutSourceConnectorProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiVersion", GoMethod: "ResetApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectorProfileName", GoMethod: "ResetConnectorProfileName"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncrementalPullConfig", GoMethod: "ResetIncrementalPullConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourceConnectorProperties", GoGetter: "SourceConnectorProperties"},
			_jsii_.MemberProperty{JsiiProperty: "sourceConnectorPropertiesInput", GoGetter: "SourceConnectorPropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorProperties",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorProperties)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitude",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitude)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitudeOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitudeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitudeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnector",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnectorOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnectorOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customProperties", GoGetter: "CustomProperties"},
			_jsii_.MemberProperty{JsiiProperty: "customPropertiesInput", GoGetter: "CustomPropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "entityName", GoGetter: "EntityName"},
			_jsii_.MemberProperty{JsiiProperty: "entityNameInput", GoGetter: "EntityNameInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCustomProperties", GoMethod: "ResetCustomProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnectorOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadog",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadog)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadogOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadogOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadogOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatrace",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatrace)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatraceOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatraceOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatraceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalytics",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalytics)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalyticsOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalyticsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalyticsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexus",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexus)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexusOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexusOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexusOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketo",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketoOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketoOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amplitude", GoGetter: "Amplitude"},
			_jsii_.MemberProperty{JsiiProperty: "amplitudeInput", GoGetter: "AmplitudeInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customConnector", GoGetter: "CustomConnector"},
			_jsii_.MemberProperty{JsiiProperty: "customConnectorInput", GoGetter: "CustomConnectorInput"},
			_jsii_.MemberProperty{JsiiProperty: "datadog", GoGetter: "Datadog"},
			_jsii_.MemberProperty{JsiiProperty: "datadogInput", GoGetter: "DatadogInput"},
			_jsii_.MemberProperty{JsiiProperty: "dynatrace", GoGetter: "Dynatrace"},
			_jsii_.MemberProperty{JsiiProperty: "dynatraceInput", GoGetter: "DynatraceInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "googleAnalytics", GoGetter: "GoogleAnalytics"},
			_jsii_.MemberProperty{JsiiProperty: "googleAnalyticsInput", GoGetter: "GoogleAnalyticsInput"},
			_jsii_.MemberProperty{JsiiProperty: "inforNexus", GoGetter: "InforNexus"},
			_jsii_.MemberProperty{JsiiProperty: "inforNexusInput", GoGetter: "InforNexusInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "marketo", GoGetter: "Marketo"},
			_jsii_.MemberProperty{JsiiProperty: "marketoInput", GoGetter: "MarketoInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAmplitude", GoMethod: "PutAmplitude"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomConnector", GoMethod: "PutCustomConnector"},
			_jsii_.MemberMethod{JsiiMethod: "putDatadog", GoMethod: "PutDatadog"},
			_jsii_.MemberMethod{JsiiMethod: "putDynatrace", GoMethod: "PutDynatrace"},
			_jsii_.MemberMethod{JsiiMethod: "putGoogleAnalytics", GoMethod: "PutGoogleAnalytics"},
			_jsii_.MemberMethod{JsiiMethod: "putInforNexus", GoMethod: "PutInforNexus"},
			_jsii_.MemberMethod{JsiiMethod: "putMarketo", GoMethod: "PutMarketo"},
			_jsii_.MemberMethod{JsiiMethod: "putS3", GoMethod: "PutS3"},
			_jsii_.MemberMethod{JsiiMethod: "putSalesforce", GoMethod: "PutSalesforce"},
			_jsii_.MemberMethod{JsiiMethod: "putSapoData", GoMethod: "PutSapoData"},
			_jsii_.MemberMethod{JsiiMethod: "putServiceNow", GoMethod: "PutServiceNow"},
			_jsii_.MemberMethod{JsiiMethod: "putSingular", GoMethod: "PutSingular"},
			_jsii_.MemberMethod{JsiiMethod: "putSlack", GoMethod: "PutSlack"},
			_jsii_.MemberMethod{JsiiMethod: "putTrendmicro", GoMethod: "PutTrendmicro"},
			_jsii_.MemberMethod{JsiiMethod: "putVeeva", GoMethod: "PutVeeva"},
			_jsii_.MemberMethod{JsiiMethod: "putZendesk", GoMethod: "PutZendesk"},
			_jsii_.MemberMethod{JsiiMethod: "resetAmplitude", GoMethod: "ResetAmplitude"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomConnector", GoMethod: "ResetCustomConnector"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatadog", GoMethod: "ResetDatadog"},
			_jsii_.MemberMethod{JsiiMethod: "resetDynatrace", GoMethod: "ResetDynatrace"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleAnalytics", GoMethod: "ResetGoogleAnalytics"},
			_jsii_.MemberMethod{JsiiMethod: "resetInforNexus", GoMethod: "ResetInforNexus"},
			_jsii_.MemberMethod{JsiiMethod: "resetMarketo", GoMethod: "ResetMarketo"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3", GoMethod: "ResetS3"},
			_jsii_.MemberMethod{JsiiMethod: "resetSalesforce", GoMethod: "ResetSalesforce"},
			_jsii_.MemberMethod{JsiiMethod: "resetSapoData", GoMethod: "ResetSapoData"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceNow", GoMethod: "ResetServiceNow"},
			_jsii_.MemberMethod{JsiiMethod: "resetSingular", GoMethod: "ResetSingular"},
			_jsii_.MemberMethod{JsiiMethod: "resetSlack", GoMethod: "ResetSlack"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrendmicro", GoMethod: "ResetTrendmicro"},
			_jsii_.MemberMethod{JsiiMethod: "resetVeeva", GoMethod: "ResetVeeva"},
			_jsii_.MemberMethod{JsiiMethod: "resetZendesk", GoMethod: "ResetZendesk"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3", GoGetter: "S3"},
			_jsii_.MemberProperty{JsiiProperty: "s3Input", GoGetter: "S3Input"},
			_jsii_.MemberProperty{JsiiProperty: "salesforce", GoGetter: "Salesforce"},
			_jsii_.MemberProperty{JsiiProperty: "salesforceInput", GoGetter: "SalesforceInput"},
			_jsii_.MemberProperty{JsiiProperty: "sapoData", GoGetter: "SapoData"},
			_jsii_.MemberProperty{JsiiProperty: "sapoDataInput", GoGetter: "SapoDataInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNow", GoGetter: "ServiceNow"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNowInput", GoGetter: "ServiceNowInput"},
			_jsii_.MemberProperty{JsiiProperty: "singular", GoGetter: "Singular"},
			_jsii_.MemberProperty{JsiiProperty: "singularInput", GoGetter: "SingularInput"},
			_jsii_.MemberProperty{JsiiProperty: "slack", GoGetter: "Slack"},
			_jsii_.MemberProperty{JsiiProperty: "slackInput", GoGetter: "SlackInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trendmicro", GoGetter: "Trendmicro"},
			_jsii_.MemberProperty{JsiiProperty: "trendmicroInput", GoGetter: "TrendmicroInput"},
			_jsii_.MemberProperty{JsiiProperty: "veeva", GoGetter: "Veeva"},
			_jsii_.MemberProperty{JsiiProperty: "veevaInput", GoGetter: "VeevaInput"},
			_jsii_.MemberProperty{JsiiProperty: "zendesk", GoGetter: "Zendesk"},
			_jsii_.MemberProperty{JsiiProperty: "zendeskInput", GoGetter: "ZendeskInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3OutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3OutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putS3InputFormatConfig", GoMethod: "PutS3InputFormatConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3InputFormatConfig", GoMethod: "ResetS3InputFormatConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3InputFormatConfig", GoGetter: "S3InputFormatConfig"},
			_jsii_.MemberProperty{JsiiProperty: "s3InputFormatConfigInput", GoGetter: "S3InputFormatConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3S3InputFormatConfig",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3S3InputFormatConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3S3InputFormatConfigOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3S3InputFormatConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetS3InputFileType", GoMethod: "ResetS3InputFileType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3InputFileType", GoGetter: "S3InputFileType"},
			_jsii_.MemberProperty{JsiiProperty: "s3InputFileTypeInput", GoGetter: "S3InputFileTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3S3InputFormatConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforce",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforce)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforceOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableDynamicFieldUpdate", GoGetter: "EnableDynamicFieldUpdate"},
			_jsii_.MemberProperty{JsiiProperty: "enableDynamicFieldUpdateInput", GoGetter: "EnableDynamicFieldUpdateInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includeDeletedRecords", GoGetter: "IncludeDeletedRecords"},
			_jsii_.MemberProperty{JsiiProperty: "includeDeletedRecordsInput", GoGetter: "IncludeDeletedRecordsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableDynamicFieldUpdate", GoMethod: "ResetEnableDynamicFieldUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeDeletedRecords", GoMethod: "ResetIncludeDeletedRecords"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "objectPath", GoGetter: "ObjectPath"},
			_jsii_.MemberProperty{JsiiProperty: "objectPathInput", GoGetter: "ObjectPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNow",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNow)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNowOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNowOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNowOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingular",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingular)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingularOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingularOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingularOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlack",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlack)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlackOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlackOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlackOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicro",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicro)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicroOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicroOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicroOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "documentType", GoGetter: "DocumentType"},
			_jsii_.MemberProperty{JsiiProperty: "documentTypeInput", GoGetter: "DocumentTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includeAllVersions", GoGetter: "IncludeAllVersions"},
			_jsii_.MemberProperty{JsiiProperty: "includeAllVersionsInput", GoGetter: "IncludeAllVersionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeRenditions", GoGetter: "IncludeRenditions"},
			_jsii_.MemberProperty{JsiiProperty: "includeRenditionsInput", GoGetter: "IncludeRenditionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeSourceFiles", GoGetter: "IncludeSourceFiles"},
			_jsii_.MemberProperty{JsiiProperty: "includeSourceFilesInput", GoGetter: "IncludeSourceFilesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDocumentType", GoMethod: "ResetDocumentType"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeAllVersions", GoMethod: "ResetIncludeAllVersions"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeRenditions", GoMethod: "ResetIncludeRenditions"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeSourceFiles", GoMethod: "ResetIncludeSourceFiles"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendesk",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendesk)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendeskOutputReference",
		reflect.TypeOf((*AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendeskOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberProperty{JsiiProperty: "objectInput", GoGetter: "ObjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendeskOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTask",
		reflect.TypeOf((*AppflowFlowTask)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTaskConnectorOperator",
		reflect.TypeOf((*AppflowFlowTaskConnectorOperator)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTaskConnectorOperatorList",
		reflect.TypeOf((*AppflowFlowTaskConnectorOperatorList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_AppflowFlowTaskConnectorOperatorList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTaskConnectorOperatorOutputReference",
		reflect.TypeOf((*AppflowFlowTaskConnectorOperatorOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amplitude", GoGetter: "Amplitude"},
			_jsii_.MemberProperty{JsiiProperty: "amplitudeInput", GoGetter: "AmplitudeInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customConnector", GoGetter: "CustomConnector"},
			_jsii_.MemberProperty{JsiiProperty: "customConnectorInput", GoGetter: "CustomConnectorInput"},
			_jsii_.MemberProperty{JsiiProperty: "datadog", GoGetter: "Datadog"},
			_jsii_.MemberProperty{JsiiProperty: "datadogInput", GoGetter: "DatadogInput"},
			_jsii_.MemberProperty{JsiiProperty: "dynatrace", GoGetter: "Dynatrace"},
			_jsii_.MemberProperty{JsiiProperty: "dynatraceInput", GoGetter: "DynatraceInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "googleAnalytics", GoGetter: "GoogleAnalytics"},
			_jsii_.MemberProperty{JsiiProperty: "googleAnalyticsInput", GoGetter: "GoogleAnalyticsInput"},
			_jsii_.MemberProperty{JsiiProperty: "inforNexus", GoGetter: "InforNexus"},
			_jsii_.MemberProperty{JsiiProperty: "inforNexusInput", GoGetter: "InforNexusInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "marketo", GoGetter: "Marketo"},
			_jsii_.MemberProperty{JsiiProperty: "marketoInput", GoGetter: "MarketoInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAmplitude", GoMethod: "ResetAmplitude"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomConnector", GoMethod: "ResetCustomConnector"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatadog", GoMethod: "ResetDatadog"},
			_jsii_.MemberMethod{JsiiMethod: "resetDynatrace", GoMethod: "ResetDynatrace"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleAnalytics", GoMethod: "ResetGoogleAnalytics"},
			_jsii_.MemberMethod{JsiiMethod: "resetInforNexus", GoMethod: "ResetInforNexus"},
			_jsii_.MemberMethod{JsiiMethod: "resetMarketo", GoMethod: "ResetMarketo"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3", GoMethod: "ResetS3"},
			_jsii_.MemberMethod{JsiiMethod: "resetSalesforce", GoMethod: "ResetSalesforce"},
			_jsii_.MemberMethod{JsiiMethod: "resetSapoData", GoMethod: "ResetSapoData"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceNow", GoMethod: "ResetServiceNow"},
			_jsii_.MemberMethod{JsiiMethod: "resetSingular", GoMethod: "ResetSingular"},
			_jsii_.MemberMethod{JsiiMethod: "resetSlack", GoMethod: "ResetSlack"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrendmicro", GoMethod: "ResetTrendmicro"},
			_jsii_.MemberMethod{JsiiMethod: "resetVeeva", GoMethod: "ResetVeeva"},
			_jsii_.MemberMethod{JsiiMethod: "resetZendesk", GoMethod: "ResetZendesk"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3", GoGetter: "S3"},
			_jsii_.MemberProperty{JsiiProperty: "s3Input", GoGetter: "S3Input"},
			_jsii_.MemberProperty{JsiiProperty: "salesforce", GoGetter: "Salesforce"},
			_jsii_.MemberProperty{JsiiProperty: "salesforceInput", GoGetter: "SalesforceInput"},
			_jsii_.MemberProperty{JsiiProperty: "sapoData", GoGetter: "SapoData"},
			_jsii_.MemberProperty{JsiiProperty: "sapoDataInput", GoGetter: "SapoDataInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNow", GoGetter: "ServiceNow"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNowInput", GoGetter: "ServiceNowInput"},
			_jsii_.MemberProperty{JsiiProperty: "singular", GoGetter: "Singular"},
			_jsii_.MemberProperty{JsiiProperty: "singularInput", GoGetter: "SingularInput"},
			_jsii_.MemberProperty{JsiiProperty: "slack", GoGetter: "Slack"},
			_jsii_.MemberProperty{JsiiProperty: "slackInput", GoGetter: "SlackInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trendmicro", GoGetter: "Trendmicro"},
			_jsii_.MemberProperty{JsiiProperty: "trendmicroInput", GoGetter: "TrendmicroInput"},
			_jsii_.MemberProperty{JsiiProperty: "veeva", GoGetter: "Veeva"},
			_jsii_.MemberProperty{JsiiProperty: "veevaInput", GoGetter: "VeevaInput"},
			_jsii_.MemberProperty{JsiiProperty: "zendesk", GoGetter: "Zendesk"},
			_jsii_.MemberProperty{JsiiProperty: "zendeskInput", GoGetter: "ZendeskInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTaskList",
		reflect.TypeOf((*AppflowFlowTaskList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_AppflowFlowTaskList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTaskOutputReference",
		reflect.TypeOf((*AppflowFlowTaskOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "connectorOperator", GoGetter: "ConnectorOperator"},
			_jsii_.MemberProperty{JsiiProperty: "connectorOperatorInput", GoGetter: "ConnectorOperatorInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinationField", GoGetter: "DestinationField"},
			_jsii_.MemberProperty{JsiiProperty: "destinationFieldInput", GoGetter: "DestinationFieldInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putConnectorOperator", GoMethod: "PutConnectorOperator"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectorOperator", GoMethod: "ResetConnectorOperator"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestinationField", GoMethod: "ResetDestinationField"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceFields", GoMethod: "ResetSourceFields"},
			_jsii_.MemberMethod{JsiiMethod: "resetTaskProperties", GoMethod: "ResetTaskProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourceFields", GoGetter: "SourceFields"},
			_jsii_.MemberProperty{JsiiProperty: "sourceFieldsInput", GoGetter: "SourceFieldsInput"},
			_jsii_.MemberProperty{JsiiProperty: "taskProperties", GoGetter: "TaskProperties"},
			_jsii_.MemberProperty{JsiiProperty: "taskPropertiesInput", GoGetter: "TaskPropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "taskType", GoGetter: "TaskType"},
			_jsii_.MemberProperty{JsiiProperty: "taskTypeInput", GoGetter: "TaskTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowTaskOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTriggerConfig",
		reflect.TypeOf((*AppflowFlowTriggerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTriggerConfigOutputReference",
		reflect.TypeOf((*AppflowFlowTriggerConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "putTriggerProperties", GoMethod: "PutTriggerProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetTriggerProperties", GoMethod: "ResetTriggerProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggerProperties", GoGetter: "TriggerProperties"},
			_jsii_.MemberProperty{JsiiProperty: "triggerPropertiesInput", GoGetter: "TriggerPropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "triggerType", GoGetter: "TriggerType"},
			_jsii_.MemberProperty{JsiiProperty: "triggerTypeInput", GoGetter: "TriggerTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowTriggerConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTriggerConfigTriggerProperties",
		reflect.TypeOf((*AppflowFlowTriggerConfigTriggerProperties)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTriggerConfigTriggerPropertiesOutputReference",
		reflect.TypeOf((*AppflowFlowTriggerConfigTriggerPropertiesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "putScheduled", GoMethod: "PutScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduled", GoMethod: "ResetScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scheduled", GoGetter: "Scheduled"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledInput", GoGetter: "ScheduledInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTriggerConfigTriggerPropertiesScheduled",
		reflect.TypeOf((*AppflowFlowTriggerConfigTriggerPropertiesScheduled)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference",
		reflect.TypeOf((*AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataPullMode", GoGetter: "DataPullMode"},
			_jsii_.MemberProperty{JsiiProperty: "dataPullModeInput", GoGetter: "DataPullModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "firstExecutionFrom", GoGetter: "FirstExecutionFrom"},
			_jsii_.MemberProperty{JsiiProperty: "firstExecutionFromInput", GoGetter: "FirstExecutionFromInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDataPullMode", GoMethod: "ResetDataPullMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetFirstExecutionFrom", GoMethod: "ResetFirstExecutionFrom"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduleEndTime", GoMethod: "ResetScheduleEndTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduleOffset", GoMethod: "ResetScheduleOffset"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduleStartTime", GoMethod: "ResetScheduleStartTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimezone", GoMethod: "ResetTimezone"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleEndTime", GoGetter: "ScheduleEndTime"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleEndTimeInput", GoGetter: "ScheduleEndTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleExpression", GoGetter: "ScheduleExpression"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleExpressionInput", GoGetter: "ScheduleExpressionInput"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleOffset", GoGetter: "ScheduleOffset"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleOffsetInput", GoGetter: "ScheduleOffsetInput"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleStartTime", GoGetter: "ScheduleStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleStartTimeInput", GoGetter: "ScheduleStartTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timezone", GoGetter: "Timezone"},
			_jsii_.MemberProperty{JsiiProperty: "timezoneInput", GoGetter: "TimezoneInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
