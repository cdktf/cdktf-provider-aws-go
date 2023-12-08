// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomain

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomain",
		reflect.TypeOf((*ElasticsearchDomain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessPolicies", GoGetter: "AccessPolicies"},
			_jsii_.MemberProperty{JsiiProperty: "accessPoliciesInput", GoGetter: "AccessPoliciesInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "advancedOptions", GoGetter: "AdvancedOptions"},
			_jsii_.MemberProperty{JsiiProperty: "advancedOptionsInput", GoGetter: "AdvancedOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "advancedSecurityOptions", GoGetter: "AdvancedSecurityOptions"},
			_jsii_.MemberProperty{JsiiProperty: "advancedSecurityOptionsInput", GoGetter: "AdvancedSecurityOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "autoTuneOptions", GoGetter: "AutoTuneOptions"},
			_jsii_.MemberProperty{JsiiProperty: "autoTuneOptionsInput", GoGetter: "AutoTuneOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterConfig", GoGetter: "ClusterConfig"},
			_jsii_.MemberProperty{JsiiProperty: "clusterConfigInput", GoGetter: "ClusterConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "cognitoOptions", GoGetter: "CognitoOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cognitoOptionsInput", GoGetter: "CognitoOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "domainEndpointOptions", GoGetter: "DomainEndpointOptions"},
			_jsii_.MemberProperty{JsiiProperty: "domainEndpointOptionsInput", GoGetter: "DomainEndpointOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "domainId", GoGetter: "DomainId"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "domainNameInput", GoGetter: "DomainNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptions", GoGetter: "EbsOptions"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptionsInput", GoGetter: "EbsOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchVersion", GoGetter: "ElasticsearchVersion"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchVersionInput", GoGetter: "ElasticsearchVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryptAtRest", GoGetter: "EncryptAtRest"},
			_jsii_.MemberProperty{JsiiProperty: "encryptAtRestInput", GoGetter: "EncryptAtRestInput"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
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
			_jsii_.MemberProperty{JsiiProperty: "kibanaEndpoint", GoGetter: "KibanaEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logPublishingOptions", GoGetter: "LogPublishingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "logPublishingOptionsInput", GoGetter: "LogPublishingOptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodeToNodeEncryption", GoGetter: "NodeToNodeEncryption"},
			_jsii_.MemberProperty{JsiiProperty: "nodeToNodeEncryptionInput", GoGetter: "NodeToNodeEncryptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAdvancedSecurityOptions", GoMethod: "PutAdvancedSecurityOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putAutoTuneOptions", GoMethod: "PutAutoTuneOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putClusterConfig", GoMethod: "PutClusterConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putCognitoOptions", GoMethod: "PutCognitoOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putDomainEndpointOptions", GoMethod: "PutDomainEndpointOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEbsOptions", GoMethod: "PutEbsOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptAtRest", GoMethod: "PutEncryptAtRest"},
			_jsii_.MemberMethod{JsiiMethod: "putLogPublishingOptions", GoMethod: "PutLogPublishingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putNodeToNodeEncryption", GoMethod: "PutNodeToNodeEncryption"},
			_jsii_.MemberMethod{JsiiMethod: "putSnapshotOptions", GoMethod: "PutSnapshotOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "putVpcOptions", GoMethod: "PutVpcOptions"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessPolicies", GoMethod: "ResetAccessPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdvancedOptions", GoMethod: "ResetAdvancedOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdvancedSecurityOptions", GoMethod: "ResetAdvancedSecurityOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoTuneOptions", GoMethod: "ResetAutoTuneOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetClusterConfig", GoMethod: "ResetClusterConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetCognitoOptions", GoMethod: "ResetCognitoOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomainEndpointOptions", GoMethod: "ResetDomainEndpointOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetEbsOptions", GoMethod: "ResetEbsOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetElasticsearchVersion", GoMethod: "ResetElasticsearchVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptAtRest", GoMethod: "ResetEncryptAtRest"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogPublishingOptions", GoMethod: "ResetLogPublishingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetNodeToNodeEncryption", GoMethod: "ResetNodeToNodeEncryption"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSnapshotOptions", GoMethod: "ResetSnapshotOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcOptions", GoMethod: "ResetVpcOptions"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotOptions", GoGetter: "SnapshotOptions"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotOptionsInput", GoGetter: "SnapshotOptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOptions", GoGetter: "VpcOptions"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOptionsInput", GoGetter: "VpcOptionsInput"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomain{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainAdvancedSecurityOptions",
		reflect.TypeOf((*ElasticsearchDomainAdvancedSecurityOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions",
		reflect.TypeOf((*ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference",
		reflect.TypeOf((*ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "masterUserArn", GoGetter: "MasterUserArn"},
			_jsii_.MemberProperty{JsiiProperty: "masterUserArnInput", GoGetter: "MasterUserArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "masterUserName", GoGetter: "MasterUserName"},
			_jsii_.MemberProperty{JsiiProperty: "masterUserNameInput", GoGetter: "MasterUserNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "masterUserPassword", GoGetter: "MasterUserPassword"},
			_jsii_.MemberProperty{JsiiProperty: "masterUserPasswordInput", GoGetter: "MasterUserPasswordInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMasterUserArn", GoMethod: "ResetMasterUserArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetMasterUserName", GoMethod: "ResetMasterUserName"},
			_jsii_.MemberMethod{JsiiMethod: "resetMasterUserPassword", GoMethod: "ResetMasterUserPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainAdvancedSecurityOptionsOutputReference",
		reflect.TypeOf((*ElasticsearchDomainAdvancedSecurityOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalUserDatabaseEnabled", GoGetter: "InternalUserDatabaseEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "internalUserDatabaseEnabledInput", GoGetter: "InternalUserDatabaseEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "masterUserOptions", GoGetter: "MasterUserOptions"},
			_jsii_.MemberProperty{JsiiProperty: "masterUserOptionsInput", GoGetter: "MasterUserOptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMasterUserOptions", GoMethod: "PutMasterUserOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetInternalUserDatabaseEnabled", GoMethod: "ResetInternalUserDatabaseEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetMasterUserOptions", GoMethod: "ResetMasterUserOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainAutoTuneOptions",
		reflect.TypeOf((*ElasticsearchDomainAutoTuneOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainAutoTuneOptionsMaintenanceSchedule",
		reflect.TypeOf((*ElasticsearchDomainAutoTuneOptionsMaintenanceSchedule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration",
		reflect.TypeOf((*ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference",
		reflect.TypeOf((*ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleList",
		reflect.TypeOf((*ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleList)(nil)).Elem(),
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
			j := jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleOutputReference",
		reflect.TypeOf((*ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "cronExpressionForRecurrence", GoGetter: "CronExpressionForRecurrence"},
			_jsii_.MemberProperty{JsiiProperty: "cronExpressionForRecurrenceInput", GoGetter: "CronExpressionForRecurrenceInput"},
			_jsii_.MemberProperty{JsiiProperty: "duration", GoGetter: "Duration"},
			_jsii_.MemberProperty{JsiiProperty: "durationInput", GoGetter: "DurationInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putDuration", GoMethod: "PutDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "startAt", GoGetter: "StartAt"},
			_jsii_.MemberProperty{JsiiProperty: "startAtInput", GoGetter: "StartAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainAutoTuneOptionsOutputReference",
		reflect.TypeOf((*ElasticsearchDomainAutoTuneOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "desiredState", GoGetter: "DesiredState"},
			_jsii_.MemberProperty{JsiiProperty: "desiredStateInput", GoGetter: "DesiredStateInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "maintenanceSchedule", GoGetter: "MaintenanceSchedule"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceScheduleInput", GoGetter: "MaintenanceScheduleInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMaintenanceSchedule", GoMethod: "PutMaintenanceSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaintenanceSchedule", GoMethod: "ResetMaintenanceSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "resetRollbackOnDisable", GoMethod: "ResetRollbackOnDisable"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rollbackOnDisable", GoGetter: "RollbackOnDisable"},
			_jsii_.MemberProperty{JsiiProperty: "rollbackOnDisableInput", GoGetter: "RollbackOnDisableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainClusterConfig",
		reflect.TypeOf((*ElasticsearchDomainClusterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainClusterConfigColdStorageOptions",
		reflect.TypeOf((*ElasticsearchDomainClusterConfigColdStorageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainClusterConfigColdStorageOptionsOutputReference",
		reflect.TypeOf((*ElasticsearchDomainClusterConfigColdStorageOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainClusterConfigColdStorageOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainClusterConfigOutputReference",
		reflect.TypeOf((*ElasticsearchDomainClusterConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "coldStorageOptions", GoGetter: "ColdStorageOptions"},
			_jsii_.MemberProperty{JsiiProperty: "coldStorageOptionsInput", GoGetter: "ColdStorageOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dedicatedMasterCount", GoGetter: "DedicatedMasterCount"},
			_jsii_.MemberProperty{JsiiProperty: "dedicatedMasterCountInput", GoGetter: "DedicatedMasterCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "dedicatedMasterEnabled", GoGetter: "DedicatedMasterEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "dedicatedMasterEnabledInput", GoGetter: "DedicatedMasterEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "dedicatedMasterType", GoGetter: "DedicatedMasterType"},
			_jsii_.MemberProperty{JsiiProperty: "dedicatedMasterTypeInput", GoGetter: "DedicatedMasterTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "instanceCount", GoGetter: "InstanceCount"},
			_jsii_.MemberProperty{JsiiProperty: "instanceCountInput", GoGetter: "InstanceCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putColdStorageOptions", GoMethod: "PutColdStorageOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putZoneAwarenessConfig", GoMethod: "PutZoneAwarenessConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetColdStorageOptions", GoMethod: "ResetColdStorageOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDedicatedMasterCount", GoMethod: "ResetDedicatedMasterCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetDedicatedMasterEnabled", GoMethod: "ResetDedicatedMasterEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetDedicatedMasterType", GoMethod: "ResetDedicatedMasterType"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceCount", GoMethod: "ResetInstanceCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetWarmCount", GoMethod: "ResetWarmCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetWarmEnabled", GoMethod: "ResetWarmEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetWarmType", GoMethod: "ResetWarmType"},
			_jsii_.MemberMethod{JsiiMethod: "resetZoneAwarenessConfig", GoMethod: "ResetZoneAwarenessConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetZoneAwarenessEnabled", GoMethod: "ResetZoneAwarenessEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "warmCount", GoGetter: "WarmCount"},
			_jsii_.MemberProperty{JsiiProperty: "warmCountInput", GoGetter: "WarmCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "warmEnabled", GoGetter: "WarmEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "warmEnabledInput", GoGetter: "WarmEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "warmType", GoGetter: "WarmType"},
			_jsii_.MemberProperty{JsiiProperty: "warmTypeInput", GoGetter: "WarmTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "zoneAwarenessConfig", GoGetter: "ZoneAwarenessConfig"},
			_jsii_.MemberProperty{JsiiProperty: "zoneAwarenessConfigInput", GoGetter: "ZoneAwarenessConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "zoneAwarenessEnabled", GoGetter: "ZoneAwarenessEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "zoneAwarenessEnabledInput", GoGetter: "ZoneAwarenessEnabledInput"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainClusterConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainClusterConfigZoneAwarenessConfig",
		reflect.TypeOf((*ElasticsearchDomainClusterConfigZoneAwarenessConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference",
		reflect.TypeOf((*ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneCount", GoGetter: "AvailabilityZoneCount"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneCountInput", GoGetter: "AvailabilityZoneCountInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAvailabilityZoneCount", GoMethod: "ResetAvailabilityZoneCount"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainCognitoOptions",
		reflect.TypeOf((*ElasticsearchDomainCognitoOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainCognitoOptionsOutputReference",
		reflect.TypeOf((*ElasticsearchDomainCognitoOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityPoolId", GoGetter: "IdentityPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolIdInput", GoGetter: "IdentityPoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolId", GoGetter: "UserPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolIdInput", GoGetter: "UserPoolIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainConfig",
		reflect.TypeOf((*ElasticsearchDomainConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainDomainEndpointOptions",
		reflect.TypeOf((*ElasticsearchDomainDomainEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainDomainEndpointOptionsOutputReference",
		reflect.TypeOf((*ElasticsearchDomainDomainEndpointOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customEndpoint", GoGetter: "CustomEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "customEndpointCertificateArn", GoGetter: "CustomEndpointCertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "customEndpointCertificateArnInput", GoGetter: "CustomEndpointCertificateArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "customEndpointEnabled", GoGetter: "CustomEndpointEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "customEndpointEnabledInput", GoGetter: "CustomEndpointEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "customEndpointInput", GoGetter: "CustomEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "enforceHttps", GoGetter: "EnforceHttps"},
			_jsii_.MemberProperty{JsiiProperty: "enforceHttpsInput", GoGetter: "EnforceHttpsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCustomEndpoint", GoMethod: "ResetCustomEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomEndpointCertificateArn", GoMethod: "ResetCustomEndpointCertificateArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomEndpointEnabled", GoMethod: "ResetCustomEndpointEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnforceHttps", GoMethod: "ResetEnforceHttps"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsSecurityPolicy", GoMethod: "ResetTlsSecurityPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tlsSecurityPolicy", GoGetter: "TlsSecurityPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "tlsSecurityPolicyInput", GoGetter: "TlsSecurityPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainEbsOptions",
		reflect.TypeOf((*ElasticsearchDomainEbsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainEbsOptionsOutputReference",
		reflect.TypeOf((*ElasticsearchDomainEbsOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "ebsEnabled", GoGetter: "EbsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ebsEnabledInput", GoGetter: "EbsEnabledInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "iops", GoGetter: "Iops"},
			_jsii_.MemberProperty{JsiiProperty: "iopsInput", GoGetter: "IopsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIops", GoMethod: "ResetIops"},
			_jsii_.MemberMethod{JsiiMethod: "resetThroughput", GoMethod: "ResetThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "resetVolumeSize", GoMethod: "ResetVolumeSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetVolumeType", GoMethod: "ResetVolumeType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "throughput", GoGetter: "Throughput"},
			_jsii_.MemberProperty{JsiiProperty: "throughputInput", GoGetter: "ThroughputInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumeSize", GoGetter: "VolumeSize"},
			_jsii_.MemberProperty{JsiiProperty: "volumeSizeInput", GoGetter: "VolumeSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "volumeType", GoGetter: "VolumeType"},
			_jsii_.MemberProperty{JsiiProperty: "volumeTypeInput", GoGetter: "VolumeTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainEncryptAtRest",
		reflect.TypeOf((*ElasticsearchDomainEncryptAtRest)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainEncryptAtRestOutputReference",
		reflect.TypeOf((*ElasticsearchDomainEncryptAtRestOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainLogPublishingOptions",
		reflect.TypeOf((*ElasticsearchDomainLogPublishingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainLogPublishingOptionsList",
		reflect.TypeOf((*ElasticsearchDomainLogPublishingOptionsList)(nil)).Elem(),
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
			j := jsiiProxy_ElasticsearchDomainLogPublishingOptionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainLogPublishingOptionsOutputReference",
		reflect.TypeOf((*ElasticsearchDomainLogPublishingOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLogGroupArn", GoGetter: "CloudwatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLogGroupArnInput", GoGetter: "CloudwatchLogGroupArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberProperty{JsiiProperty: "logTypeInput", GoGetter: "LogTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainLogPublishingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainNodeToNodeEncryption",
		reflect.TypeOf((*ElasticsearchDomainNodeToNodeEncryption)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainNodeToNodeEncryptionOutputReference",
		reflect.TypeOf((*ElasticsearchDomainNodeToNodeEncryptionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
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
			j := jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainSnapshotOptions",
		reflect.TypeOf((*ElasticsearchDomainSnapshotOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainSnapshotOptionsOutputReference",
		reflect.TypeOf((*ElasticsearchDomainSnapshotOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "automatedSnapshotStartHour", GoGetter: "AutomatedSnapshotStartHour"},
			_jsii_.MemberProperty{JsiiProperty: "automatedSnapshotStartHourInput", GoGetter: "AutomatedSnapshotStartHourInput"},
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
			j := jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainTimeouts",
		reflect.TypeOf((*ElasticsearchDomainTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainTimeoutsOutputReference",
		reflect.TypeOf((*ElasticsearchDomainTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainVpcOptions",
		reflect.TypeOf((*ElasticsearchDomainVpcOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainVpcOptionsOutputReference",
		reflect.TypeOf((*ElasticsearchDomainVpcOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroupIds", GoMethod: "ResetSecurityGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnetIds", GoMethod: "ResetSubnetIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdsInput", GoGetter: "SecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdsInput", GoGetter: "SubnetIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
