// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagegatewaygateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/storagegatewaygateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/storagegateway_gateway aws_storagegateway_gateway}.
type StoragegatewayGateway interface {
	cdktf.TerraformResource
	ActivationKey() *string
	SetActivationKey(val *string)
	ActivationKeyInput() *string
	Arn() *string
	AverageDownloadRateLimitInBitsPerSec() *float64
	SetAverageDownloadRateLimitInBitsPerSec(val *float64)
	AverageDownloadRateLimitInBitsPerSecInput() *float64
	AverageUploadRateLimitInBitsPerSec() *float64
	SetAverageUploadRateLimitInBitsPerSec(val *float64)
	AverageUploadRateLimitInBitsPerSecInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudwatchLogGroupArn() *string
	SetCloudwatchLogGroupArn(val *string)
	CloudwatchLogGroupArnInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Ec2InstanceId() *string
	EndpointType() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayId() *string
	GatewayIpAddress() *string
	SetGatewayIpAddress(val *string)
	GatewayIpAddressInput() *string
	GatewayName() *string
	SetGatewayName(val *string)
	GatewayNameInput() *string
	GatewayNetworkInterface() StoragegatewayGatewayGatewayNetworkInterfaceList
	GatewayTimezone() *string
	SetGatewayTimezone(val *string)
	GatewayTimezoneInput() *string
	GatewayType() *string
	SetGatewayType(val *string)
	GatewayTypeInput() *string
	GatewayVpcEndpoint() *string
	SetGatewayVpcEndpoint(val *string)
	GatewayVpcEndpointInput() *string
	HostEnvironment() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenanceStartTime() StoragegatewayGatewayMaintenanceStartTimeOutputReference
	MaintenanceStartTimeInput() *StoragegatewayGatewayMaintenanceStartTime
	MediumChangerType() *string
	SetMediumChangerType(val *string)
	MediumChangerTypeInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SmbActiveDirectorySettings() StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference
	SmbActiveDirectorySettingsInput() *StoragegatewayGatewaySmbActiveDirectorySettings
	SmbFileShareVisibility() interface{}
	SetSmbFileShareVisibility(val interface{})
	SmbFileShareVisibilityInput() interface{}
	SmbGuestPassword() *string
	SetSmbGuestPassword(val *string)
	SmbGuestPasswordInput() *string
	SmbSecurityStrategy() *string
	SetSmbSecurityStrategy(val *string)
	SmbSecurityStrategyInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TapeDriveType() *string
	SetTapeDriveType(val *string)
	TapeDriveTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() StoragegatewayGatewayTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutMaintenanceStartTime(value *StoragegatewayGatewayMaintenanceStartTime)
	PutSmbActiveDirectorySettings(value *StoragegatewayGatewaySmbActiveDirectorySettings)
	PutTimeouts(value *StoragegatewayGatewayTimeouts)
	ResetActivationKey()
	ResetAverageDownloadRateLimitInBitsPerSec()
	ResetAverageUploadRateLimitInBitsPerSec()
	ResetCloudwatchLogGroupArn()
	ResetGatewayIpAddress()
	ResetGatewayType()
	ResetGatewayVpcEndpoint()
	ResetId()
	ResetMaintenanceStartTime()
	ResetMediumChangerType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSmbActiveDirectorySettings()
	ResetSmbFileShareVisibility()
	ResetSmbGuestPassword()
	ResetSmbSecurityStrategy()
	ResetTags()
	ResetTagsAll()
	ResetTapeDriveType()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for StoragegatewayGateway
type jsiiProxy_StoragegatewayGateway struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewayGateway) ActivationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) ActivationKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) AverageDownloadRateLimitInBitsPerSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"averageDownloadRateLimitInBitsPerSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) AverageDownloadRateLimitInBitsPerSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"averageDownloadRateLimitInBitsPerSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) AverageUploadRateLimitInBitsPerSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"averageUploadRateLimitInBitsPerSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) AverageUploadRateLimitInBitsPerSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"averageUploadRateLimitInBitsPerSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) CloudwatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) CloudwatchLogGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Ec2InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayNetworkInterface() StoragegatewayGatewayGatewayNetworkInterfaceList {
	var returns StoragegatewayGatewayGatewayNetworkInterfaceList
	_jsii_.Get(
		j,
		"gatewayNetworkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayVpcEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayVpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayVpcEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayVpcEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) HostEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) MaintenanceStartTime() StoragegatewayGatewayMaintenanceStartTimeOutputReference {
	var returns StoragegatewayGatewayMaintenanceStartTimeOutputReference
	_jsii_.Get(
		j,
		"maintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) MaintenanceStartTimeInput() *StoragegatewayGatewayMaintenanceStartTime {
	var returns *StoragegatewayGatewayMaintenanceStartTime
	_jsii_.Get(
		j,
		"maintenanceStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) MediumChangerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediumChangerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) MediumChangerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediumChangerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbActiveDirectorySettings() StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference {
	var returns StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference
	_jsii_.Get(
		j,
		"smbActiveDirectorySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbActiveDirectorySettingsInput() *StoragegatewayGatewaySmbActiveDirectorySettings {
	var returns *StoragegatewayGatewaySmbActiveDirectorySettings
	_jsii_.Get(
		j,
		"smbActiveDirectorySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbFileShareVisibility() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbFileShareVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbFileShareVisibilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbFileShareVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbGuestPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbGuestPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbGuestPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbGuestPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbSecurityStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbSecurityStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbSecurityStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbSecurityStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TapeDriveType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tapeDriveType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TapeDriveTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tapeDriveTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Timeouts() StoragegatewayGatewayTimeoutsOutputReference {
	var returns StoragegatewayGatewayTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/storagegateway_gateway aws_storagegateway_gateway} Resource.
func NewStoragegatewayGateway(scope constructs.Construct, id *string, config *StoragegatewayGatewayConfig) StoragegatewayGateway {
	_init_.Initialize()

	if err := validateNewStoragegatewayGatewayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StoragegatewayGateway{}

	_jsii_.Create(
		"@cdktf/provider-aws.storagegatewayGateway.StoragegatewayGateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/storagegateway_gateway aws_storagegateway_gateway} Resource.
func NewStoragegatewayGateway_Override(s StoragegatewayGateway, scope constructs.Construct, id *string, config *StoragegatewayGatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.storagegatewayGateway.StoragegatewayGateway",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetActivationKey(val *string) {
	if err := j.validateSetActivationKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activationKey",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetAverageDownloadRateLimitInBitsPerSec(val *float64) {
	if err := j.validateSetAverageDownloadRateLimitInBitsPerSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"averageDownloadRateLimitInBitsPerSec",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetAverageUploadRateLimitInBitsPerSec(val *float64) {
	if err := j.validateSetAverageUploadRateLimitInBitsPerSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"averageUploadRateLimitInBitsPerSec",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetCloudwatchLogGroupArn(val *string) {
	if err := j.validateSetCloudwatchLogGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchLogGroupArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetGatewayIpAddress(val *string) {
	if err := j.validateSetGatewayIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayIpAddress",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetGatewayName(val *string) {
	if err := j.validateSetGatewayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayName",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetGatewayTimezone(val *string) {
	if err := j.validateSetGatewayTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayTimezone",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetGatewayType(val *string) {
	if err := j.validateSetGatewayTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayType",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetGatewayVpcEndpoint(val *string) {
	if err := j.validateSetGatewayVpcEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayVpcEndpoint",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetMediumChangerType(val *string) {
	if err := j.validateSetMediumChangerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediumChangerType",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetSmbFileShareVisibility(val interface{}) {
	if err := j.validateSetSmbFileShareVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbFileShareVisibility",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetSmbGuestPassword(val *string) {
	if err := j.validateSetSmbGuestPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbGuestPassword",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetSmbSecurityStrategy(val *string) {
	if err := j.validateSetSmbSecurityStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbSecurityStrategy",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway)SetTapeDriveType(val *string) {
	if err := j.validateSetTapeDriveTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tapeDriveType",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func StoragegatewayGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStoragegatewayGateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewayGateway.StoragegatewayGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StoragegatewayGateway_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStoragegatewayGateway_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewayGateway.StoragegatewayGateway",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StoragegatewayGateway_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStoragegatewayGateway_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewayGateway.StoragegatewayGateway",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewayGateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.storagegatewayGateway.StoragegatewayGateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StoragegatewayGateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StoragegatewayGateway) PutMaintenanceStartTime(value *StoragegatewayGatewayMaintenanceStartTime) {
	if err := s.validatePutMaintenanceStartTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMaintenanceStartTime",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewayGateway) PutSmbActiveDirectorySettings(value *StoragegatewayGatewaySmbActiveDirectorySettings) {
	if err := s.validatePutSmbActiveDirectorySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSmbActiveDirectorySettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewayGateway) PutTimeouts(value *StoragegatewayGatewayTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetActivationKey() {
	_jsii_.InvokeVoid(
		s,
		"resetActivationKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetAverageDownloadRateLimitInBitsPerSec() {
	_jsii_.InvokeVoid(
		s,
		"resetAverageDownloadRateLimitInBitsPerSec",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetAverageUploadRateLimitInBitsPerSec() {
	_jsii_.InvokeVoid(
		s,
		"resetAverageUploadRateLimitInBitsPerSec",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetCloudwatchLogGroupArn() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudwatchLogGroupArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetGatewayIpAddress() {
	_jsii_.InvokeVoid(
		s,
		"resetGatewayIpAddress",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetGatewayType() {
	_jsii_.InvokeVoid(
		s,
		"resetGatewayType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetGatewayVpcEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetGatewayVpcEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetMaintenanceStartTime() {
	_jsii_.InvokeVoid(
		s,
		"resetMaintenanceStartTime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetMediumChangerType() {
	_jsii_.InvokeVoid(
		s,
		"resetMediumChangerType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetSmbActiveDirectorySettings() {
	_jsii_.InvokeVoid(
		s,
		"resetSmbActiveDirectorySettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetSmbFileShareVisibility() {
	_jsii_.InvokeVoid(
		s,
		"resetSmbFileShareVisibility",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetSmbGuestPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetSmbGuestPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetSmbSecurityStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetSmbSecurityStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetTapeDriveType() {
	_jsii_.InvokeVoid(
		s,
		"resetTapeDriveType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

