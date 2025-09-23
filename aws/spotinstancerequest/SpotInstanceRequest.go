// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotinstancerequest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/spotinstancerequest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/spot_instance_request aws_spot_instance_request}.
type SpotInstanceRequest interface {
	cdktf.TerraformResource
	Ami() *string
	SetAmi(val *string)
	AmiInput() *string
	Arn() *string
	AssociatePublicIpAddress() interface{}
	SetAssociatePublicIpAddress(val interface{})
	AssociatePublicIpAddressInput() interface{}
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	CapacityReservationSpecification() SpotInstanceRequestCapacityReservationSpecificationOutputReference
	CapacityReservationSpecificationInput() *SpotInstanceRequestCapacityReservationSpecification
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CpuOptions() SpotInstanceRequestCpuOptionsOutputReference
	CpuOptionsInput() *SpotInstanceRequestCpuOptions
	CreditSpecification() SpotInstanceRequestCreditSpecificationOutputReference
	CreditSpecificationInput() *SpotInstanceRequestCreditSpecification
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableApiStop() interface{}
	SetDisableApiStop(val interface{})
	DisableApiStopInput() interface{}
	DisableApiTermination() interface{}
	SetDisableApiTermination(val interface{})
	DisableApiTerminationInput() interface{}
	EbsBlockDevice() SpotInstanceRequestEbsBlockDeviceList
	EbsBlockDeviceInput() interface{}
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	EnablePrimaryIpv6() interface{}
	SetEnablePrimaryIpv6(val interface{})
	EnablePrimaryIpv6Input() interface{}
	EnclaveOptions() SpotInstanceRequestEnclaveOptionsOutputReference
	EnclaveOptionsInput() *SpotInstanceRequestEnclaveOptions
	EphemeralBlockDevice() SpotInstanceRequestEphemeralBlockDeviceList
	EphemeralBlockDeviceInput() interface{}
	FetchPasswordData() interface{}
	SetFetchPasswordData(val interface{})
	FetchPasswordDataInput() interface{}
	ForceDestroy() interface{}
	SetForceDestroy(val interface{})
	ForceDestroyInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hibernation() interface{}
	SetHibernation(val interface{})
	HibernationInput() interface{}
	HostId() *string
	SetHostId(val *string)
	HostIdInput() *string
	HostResourceGroupArn() *string
	SetHostResourceGroupArn(val *string)
	HostResourceGroupArnInput() *string
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	IamInstanceProfileInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceInitiatedShutdownBehavior() *string
	SetInstanceInitiatedShutdownBehavior(val *string)
	InstanceInitiatedShutdownBehaviorInput() *string
	InstanceInterruptionBehavior() *string
	SetInstanceInterruptionBehavior(val *string)
	InstanceInterruptionBehaviorInput() *string
	InstanceState() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	Ipv6AddressCount() *float64
	SetIpv6AddressCount(val *float64)
	Ipv6AddressCountInput() *float64
	Ipv6Addresses() *[]*string
	SetIpv6Addresses(val *[]*string)
	Ipv6AddressesInput() *[]*string
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	LaunchGroup() *string
	SetLaunchGroup(val *string)
	LaunchGroupInput() *string
	LaunchTemplate() SpotInstanceRequestLaunchTemplateOutputReference
	LaunchTemplateInput() *SpotInstanceRequestLaunchTemplate
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenanceOptions() SpotInstanceRequestMaintenanceOptionsOutputReference
	MaintenanceOptionsInput() *SpotInstanceRequestMaintenanceOptions
	MetadataOptions() SpotInstanceRequestMetadataOptionsOutputReference
	MetadataOptionsInput() *SpotInstanceRequestMetadataOptions
	Monitoring() interface{}
	SetMonitoring(val interface{})
	MonitoringInput() interface{}
	NetworkInterface() SpotInstanceRequestNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	OutpostArn() *string
	PasswordData() *string
	PlacementGroup() *string
	SetPlacementGroup(val *string)
	PlacementGroupId() *string
	SetPlacementGroupId(val *string)
	PlacementGroupIdInput() *string
	PlacementGroupInput() *string
	PlacementPartitionNumber() *float64
	SetPlacementPartitionNumber(val *float64)
	PlacementPartitionNumberInput() *float64
	PrimaryNetworkInterface() SpotInstanceRequestPrimaryNetworkInterfaceList
	PrimaryNetworkInterfaceId() *string
	PrivateDns() *string
	PrivateDnsNameOptions() SpotInstanceRequestPrivateDnsNameOptionsOutputReference
	PrivateDnsNameOptionsInput() *SpotInstanceRequestPrivateDnsNameOptions
	PrivateIp() *string
	SetPrivateIp(val *string)
	PrivateIpInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicDns() *string
	PublicIp() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RootBlockDevice() SpotInstanceRequestRootBlockDeviceOutputReference
	RootBlockDeviceInput() *SpotInstanceRequestRootBlockDevice
	SecondaryPrivateIps() *[]*string
	SetSecondaryPrivateIps(val *[]*string)
	SecondaryPrivateIpsInput() *[]*string
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SourceDestCheck() interface{}
	SetSourceDestCheck(val interface{})
	SourceDestCheckInput() interface{}
	SpotBidStatus() *string
	SpotInstanceId() *string
	SpotPrice() *string
	SetSpotPrice(val *string)
	SpotPriceInput() *string
	SpotRequestState() *string
	SpotType() *string
	SetSpotType(val *string)
	SpotTypeInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	Tenancy() *string
	SetTenancy(val *string)
	TenancyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SpotInstanceRequestTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserData() *string
	SetUserData(val *string)
	UserDataBase64() *string
	SetUserDataBase64(val *string)
	UserDataBase64Input() *string
	UserDataInput() *string
	UserDataReplaceOnChange() interface{}
	SetUserDataReplaceOnChange(val interface{})
	UserDataReplaceOnChangeInput() interface{}
	ValidFrom() *string
	SetValidFrom(val *string)
	ValidFromInput() *string
	ValidUntil() *string
	SetValidUntil(val *string)
	ValidUntilInput() *string
	VolumeTags() *map[string]*string
	SetVolumeTags(val *map[string]*string)
	VolumeTagsInput() *map[string]*string
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	VpcSecurityGroupIdsInput() *[]*string
	WaitForFulfillment() interface{}
	SetWaitForFulfillment(val interface{})
	WaitForFulfillmentInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCapacityReservationSpecification(value *SpotInstanceRequestCapacityReservationSpecification)
	PutCpuOptions(value *SpotInstanceRequestCpuOptions)
	PutCreditSpecification(value *SpotInstanceRequestCreditSpecification)
	PutEbsBlockDevice(value interface{})
	PutEnclaveOptions(value *SpotInstanceRequestEnclaveOptions)
	PutEphemeralBlockDevice(value interface{})
	PutLaunchTemplate(value *SpotInstanceRequestLaunchTemplate)
	PutMaintenanceOptions(value *SpotInstanceRequestMaintenanceOptions)
	PutMetadataOptions(value *SpotInstanceRequestMetadataOptions)
	PutNetworkInterface(value interface{})
	PutPrivateDnsNameOptions(value *SpotInstanceRequestPrivateDnsNameOptions)
	PutRootBlockDevice(value *SpotInstanceRequestRootBlockDevice)
	PutTimeouts(value *SpotInstanceRequestTimeouts)
	ResetAmi()
	ResetAssociatePublicIpAddress()
	ResetAvailabilityZone()
	ResetCapacityReservationSpecification()
	ResetCpuOptions()
	ResetCreditSpecification()
	ResetDisableApiStop()
	ResetDisableApiTermination()
	ResetEbsBlockDevice()
	ResetEbsOptimized()
	ResetEnablePrimaryIpv6()
	ResetEnclaveOptions()
	ResetEphemeralBlockDevice()
	ResetFetchPasswordData()
	ResetForceDestroy()
	ResetHibernation()
	ResetHostId()
	ResetHostResourceGroupArn()
	ResetIamInstanceProfile()
	ResetId()
	ResetInstanceInitiatedShutdownBehavior()
	ResetInstanceInterruptionBehavior()
	ResetInstanceType()
	ResetIpv6AddressCount()
	ResetIpv6Addresses()
	ResetKeyName()
	ResetLaunchGroup()
	ResetLaunchTemplate()
	ResetMaintenanceOptions()
	ResetMetadataOptions()
	ResetMonitoring()
	ResetNetworkInterface()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacementGroup()
	ResetPlacementGroupId()
	ResetPlacementPartitionNumber()
	ResetPrivateDnsNameOptions()
	ResetPrivateIp()
	ResetRegion()
	ResetRootBlockDevice()
	ResetSecondaryPrivateIps()
	ResetSecurityGroups()
	ResetSourceDestCheck()
	ResetSpotPrice()
	ResetSpotType()
	ResetSubnetId()
	ResetTags()
	ResetTagsAll()
	ResetTenancy()
	ResetTimeouts()
	ResetUserData()
	ResetUserDataBase64()
	ResetUserDataReplaceOnChange()
	ResetValidFrom()
	ResetValidUntil()
	ResetVolumeTags()
	ResetVpcSecurityGroupIds()
	ResetWaitForFulfillment()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SpotInstanceRequest
type jsiiProxy_SpotInstanceRequest struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SpotInstanceRequest) Ami() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ami",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) AmiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) AssociatePublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) CapacityReservationSpecification() SpotInstanceRequestCapacityReservationSpecificationOutputReference {
	var returns SpotInstanceRequestCapacityReservationSpecificationOutputReference
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) CapacityReservationSpecificationInput() *SpotInstanceRequestCapacityReservationSpecification {
	var returns *SpotInstanceRequestCapacityReservationSpecification
	_jsii_.Get(
		j,
		"capacityReservationSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) CpuOptions() SpotInstanceRequestCpuOptionsOutputReference {
	var returns SpotInstanceRequestCpuOptionsOutputReference
	_jsii_.Get(
		j,
		"cpuOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) CpuOptionsInput() *SpotInstanceRequestCpuOptions {
	var returns *SpotInstanceRequestCpuOptions
	_jsii_.Get(
		j,
		"cpuOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) CreditSpecification() SpotInstanceRequestCreditSpecificationOutputReference {
	var returns SpotInstanceRequestCreditSpecificationOutputReference
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) CreditSpecificationInput() *SpotInstanceRequestCreditSpecification {
	var returns *SpotInstanceRequestCreditSpecification
	_jsii_.Get(
		j,
		"creditSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) DisableApiStop() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) DisableApiStopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) DisableApiTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) DisableApiTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) EbsBlockDevice() SpotInstanceRequestEbsBlockDeviceList {
	var returns SpotInstanceRequestEbsBlockDeviceList
	_jsii_.Get(
		j,
		"ebsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) EbsBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) EnablePrimaryIpv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrimaryIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) EnablePrimaryIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrimaryIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) EnclaveOptions() SpotInstanceRequestEnclaveOptionsOutputReference {
	var returns SpotInstanceRequestEnclaveOptionsOutputReference
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) EnclaveOptionsInput() *SpotInstanceRequestEnclaveOptions {
	var returns *SpotInstanceRequestEnclaveOptions
	_jsii_.Get(
		j,
		"enclaveOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) EphemeralBlockDevice() SpotInstanceRequestEphemeralBlockDeviceList {
	var returns SpotInstanceRequestEphemeralBlockDeviceList
	_jsii_.Get(
		j,
		"ephemeralBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) EphemeralBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) FetchPasswordData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchPasswordData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) FetchPasswordDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchPasswordDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Hibernation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) HibernationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) HostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) HostIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) HostResourceGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) HostResourceGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) InstanceInitiatedShutdownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) InstanceInitiatedShutdownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) InstanceInterruptionBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) InstanceInterruptionBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) InstanceState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Ipv6AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Ipv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Ipv6AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) LaunchGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) LaunchGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) LaunchTemplate() SpotInstanceRequestLaunchTemplateOutputReference {
	var returns SpotInstanceRequestLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) LaunchTemplateInput() *SpotInstanceRequestLaunchTemplate {
	var returns *SpotInstanceRequestLaunchTemplate
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) MaintenanceOptions() SpotInstanceRequestMaintenanceOptionsOutputReference {
	var returns SpotInstanceRequestMaintenanceOptionsOutputReference
	_jsii_.Get(
		j,
		"maintenanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) MaintenanceOptionsInput() *SpotInstanceRequestMaintenanceOptions {
	var returns *SpotInstanceRequestMaintenanceOptions
	_jsii_.Get(
		j,
		"maintenanceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) MetadataOptions() SpotInstanceRequestMetadataOptionsOutputReference {
	var returns SpotInstanceRequestMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) MetadataOptionsInput() *SpotInstanceRequestMetadataOptions {
	var returns *SpotInstanceRequestMetadataOptions
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Monitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) MonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) NetworkInterface() SpotInstanceRequestNetworkInterfaceList {
	var returns SpotInstanceRequestNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) OutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PasswordData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PlacementGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PlacementPartitionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"placementPartitionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PlacementPartitionNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"placementPartitionNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PrimaryNetworkInterface() SpotInstanceRequestPrimaryNetworkInterfaceList {
	var returns SpotInstanceRequestPrimaryNetworkInterfaceList
	_jsii_.Get(
		j,
		"primaryNetworkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PrimaryNetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryNetworkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PrivateDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PrivateDnsNameOptions() SpotInstanceRequestPrivateDnsNameOptionsOutputReference {
	var returns SpotInstanceRequestPrivateDnsNameOptionsOutputReference
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PrivateDnsNameOptionsInput() *SpotInstanceRequestPrivateDnsNameOptions {
	var returns *SpotInstanceRequestPrivateDnsNameOptions
	_jsii_.Get(
		j,
		"privateDnsNameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) RootBlockDevice() SpotInstanceRequestRootBlockDeviceOutputReference {
	var returns SpotInstanceRequestRootBlockDeviceOutputReference
	_jsii_.Get(
		j,
		"rootBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) RootBlockDeviceInput() *SpotInstanceRequestRootBlockDevice {
	var returns *SpotInstanceRequestRootBlockDevice
	_jsii_.Get(
		j,
		"rootBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SecondaryPrivateIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryPrivateIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SecondaryPrivateIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryPrivateIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SourceDestCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDestCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SourceDestCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDestCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SpotBidStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotBidStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SpotInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SpotPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SpotPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SpotRequestState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotRequestState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SpotType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SpotTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) TenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) Timeouts() SpotInstanceRequestTimeoutsOutputReference {
	var returns SpotInstanceRequestTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) UserDataBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) UserDataBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) UserDataReplaceOnChange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDataReplaceOnChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) UserDataReplaceOnChangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDataReplaceOnChangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) ValidFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) ValidFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) ValidUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) ValidUntilInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntilInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) VolumeTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) VolumeTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) WaitForFulfillment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForFulfillment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequest) WaitForFulfillmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForFulfillmentInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/spot_instance_request aws_spot_instance_request} Resource.
func NewSpotInstanceRequest(scope constructs.Construct, id *string, config *SpotInstanceRequestConfig) SpotInstanceRequest {
	_init_.Initialize()

	if err := validateNewSpotInstanceRequestParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpotInstanceRequest{}

	_jsii_.Create(
		"@cdktf/provider-aws.spotInstanceRequest.SpotInstanceRequest",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/spot_instance_request aws_spot_instance_request} Resource.
func NewSpotInstanceRequest_Override(s SpotInstanceRequest, scope constructs.Construct, id *string, config *SpotInstanceRequestConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.spotInstanceRequest.SpotInstanceRequest",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetAmi(val *string) {
	if err := j.validateSetAmiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ami",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetAssociatePublicIpAddress(val interface{}) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetDisableApiStop(val interface{}) {
	if err := j.validateSetDisableApiStopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiStop",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetDisableApiTermination(val interface{}) {
	if err := j.validateSetDisableApiTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiTermination",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetEnablePrimaryIpv6(val interface{}) {
	if err := j.validateSetEnablePrimaryIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrimaryIpv6",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetFetchPasswordData(val interface{}) {
	if err := j.validateSetFetchPasswordDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fetchPasswordData",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetHibernation(val interface{}) {
	if err := j.validateSetHibernationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hibernation",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetHostId(val *string) {
	if err := j.validateSetHostIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostId",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetHostResourceGroupArn(val *string) {
	if err := j.validateSetHostResourceGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostResourceGroupArn",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetInstanceInitiatedShutdownBehavior(val *string) {
	if err := j.validateSetInstanceInitiatedShutdownBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInitiatedShutdownBehavior",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetInstanceInterruptionBehavior(val *string) {
	if err := j.validateSetInstanceInterruptionBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInterruptionBehavior",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetIpv6AddressCount(val *float64) {
	if err := j.validateSetIpv6AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AddressCount",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetIpv6Addresses(val *[]*string) {
	if err := j.validateSetIpv6AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Addresses",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetLaunchGroup(val *string) {
	if err := j.validateSetLaunchGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchGroup",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetMonitoring(val interface{}) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetPlacementGroup(val *string) {
	if err := j.validateSetPlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetPlacementGroupId(val *string) {
	if err := j.validateSetPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroupId",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetPlacementPartitionNumber(val *float64) {
	if err := j.validateSetPlacementPartitionNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementPartitionNumber",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetPrivateIp(val *string) {
	if err := j.validateSetPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetSecondaryPrivateIps(val *[]*string) {
	if err := j.validateSetSecondaryPrivateIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryPrivateIps",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetSourceDestCheck(val interface{}) {
	if err := j.validateSetSourceDestCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDestCheck",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetSpotPrice(val *string) {
	if err := j.validateSetSpotPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotPrice",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetSpotType(val *string) {
	if err := j.validateSetSpotTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotType",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetTenancy(val *string) {
	if err := j.validateSetTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetUserDataBase64(val *string) {
	if err := j.validateSetUserDataBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDataBase64",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetUserDataReplaceOnChange(val interface{}) {
	if err := j.validateSetUserDataReplaceOnChangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDataReplaceOnChange",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetValidFrom(val *string) {
	if err := j.validateSetValidFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validFrom",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetValidUntil(val *string) {
	if err := j.validateSetValidUntilParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validUntil",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetVolumeTags(val *map[string]*string) {
	if err := j.validateSetVolumeTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeTags",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequest)SetWaitForFulfillment(val interface{}) {
	if err := j.validateSetWaitForFulfillmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForFulfillment",
		val,
	)
}

// Generates CDKTF code for importing a SpotInstanceRequest resource upon running "cdktf plan <stack-name>".
func SpotInstanceRequest_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSpotInstanceRequest_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.spotInstanceRequest.SpotInstanceRequest",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func SpotInstanceRequest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpotInstanceRequest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.spotInstanceRequest.SpotInstanceRequest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpotInstanceRequest_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpotInstanceRequest_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.spotInstanceRequest.SpotInstanceRequest",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpotInstanceRequest_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpotInstanceRequest_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.spotInstanceRequest.SpotInstanceRequest",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SpotInstanceRequest_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.spotInstanceRequest.SpotInstanceRequest",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SpotInstanceRequest) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpotInstanceRequest) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpotInstanceRequest) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpotInstanceRequest) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpotInstanceRequest) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpotInstanceRequest) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpotInstanceRequest) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpotInstanceRequest) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpotInstanceRequest) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpotInstanceRequest) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequest) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpotInstanceRequest) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) PutCapacityReservationSpecification(value *SpotInstanceRequestCapacityReservationSpecification) {
	if err := s.validatePutCapacityReservationSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCapacityReservationSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) PutCpuOptions(value *SpotInstanceRequestCpuOptions) {
	if err := s.validatePutCpuOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCpuOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) PutCreditSpecification(value *SpotInstanceRequestCreditSpecification) {
	if err := s.validatePutCreditSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCreditSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) PutEbsBlockDevice(value interface{}) {
	if err := s.validatePutEbsBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEbsBlockDevice",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) PutEnclaveOptions(value *SpotInstanceRequestEnclaveOptions) {
	if err := s.validatePutEnclaveOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEnclaveOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) PutEphemeralBlockDevice(value interface{}) {
	if err := s.validatePutEphemeralBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEphemeralBlockDevice",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) PutLaunchTemplate(value *SpotInstanceRequestLaunchTemplate) {
	if err := s.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) PutMaintenanceOptions(value *SpotInstanceRequestMaintenanceOptions) {
	if err := s.validatePutMaintenanceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMaintenanceOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) PutMetadataOptions(value *SpotInstanceRequestMetadataOptions) {
	if err := s.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) PutNetworkInterface(value interface{}) {
	if err := s.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) PutPrivateDnsNameOptions(value *SpotInstanceRequestPrivateDnsNameOptions) {
	if err := s.validatePutPrivateDnsNameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPrivateDnsNameOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) PutRootBlockDevice(value *SpotInstanceRequestRootBlockDevice) {
	if err := s.validatePutRootBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRootBlockDevice",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) PutTimeouts(value *SpotInstanceRequestTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetAmi() {
	_jsii_.InvokeVoid(
		s,
		"resetAmi",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		s,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		s,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetCapacityReservationSpecification() {
	_jsii_.InvokeVoid(
		s,
		"resetCapacityReservationSpecification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetCpuOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetCpuOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetCreditSpecification() {
	_jsii_.InvokeVoid(
		s,
		"resetCreditSpecification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetDisableApiStop() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableApiStop",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetDisableApiTermination() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableApiTermination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetEbsBlockDevice() {
	_jsii_.InvokeVoid(
		s,
		"resetEbsBlockDevice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		s,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetEnablePrimaryIpv6() {
	_jsii_.InvokeVoid(
		s,
		"resetEnablePrimaryIpv6",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetEnclaveOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetEnclaveOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetEphemeralBlockDevice() {
	_jsii_.InvokeVoid(
		s,
		"resetEphemeralBlockDevice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetFetchPasswordData() {
	_jsii_.InvokeVoid(
		s,
		"resetFetchPasswordData",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		s,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetHibernation() {
	_jsii_.InvokeVoid(
		s,
		"resetHibernation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetHostId() {
	_jsii_.InvokeVoid(
		s,
		"resetHostId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetHostResourceGroupArn() {
	_jsii_.InvokeVoid(
		s,
		"resetHostResourceGroupArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		s,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetInstanceInitiatedShutdownBehavior() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceInitiatedShutdownBehavior",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetInstanceInterruptionBehavior() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceInterruptionBehavior",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetInstanceType() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetIpv6AddressCount() {
	_jsii_.InvokeVoid(
		s,
		"resetIpv6AddressCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetIpv6Addresses() {
	_jsii_.InvokeVoid(
		s,
		"resetIpv6Addresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetKeyName() {
	_jsii_.InvokeVoid(
		s,
		"resetKeyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetLaunchGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetLaunchGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetMaintenanceOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetMaintenanceOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetMonitoring() {
	_jsii_.InvokeVoid(
		s,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetPlacementGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetPlacementGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetPlacementGroupId() {
	_jsii_.InvokeVoid(
		s,
		"resetPlacementGroupId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetPlacementPartitionNumber() {
	_jsii_.InvokeVoid(
		s,
		"resetPlacementPartitionNumber",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetPrivateDnsNameOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateDnsNameOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetRootBlockDevice() {
	_jsii_.InvokeVoid(
		s,
		"resetRootBlockDevice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetSecondaryPrivateIps() {
	_jsii_.InvokeVoid(
		s,
		"resetSecondaryPrivateIps",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetSourceDestCheck() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceDestCheck",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetSpotPrice() {
	_jsii_.InvokeVoid(
		s,
		"resetSpotPrice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetSpotType() {
	_jsii_.InvokeVoid(
		s,
		"resetSpotType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetSubnetId() {
	_jsii_.InvokeVoid(
		s,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetTenancy() {
	_jsii_.InvokeVoid(
		s,
		"resetTenancy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetUserData() {
	_jsii_.InvokeVoid(
		s,
		"resetUserData",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetUserDataBase64() {
	_jsii_.InvokeVoid(
		s,
		"resetUserDataBase64",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetUserDataReplaceOnChange() {
	_jsii_.InvokeVoid(
		s,
		"resetUserDataReplaceOnChange",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetValidFrom() {
	_jsii_.InvokeVoid(
		s,
		"resetValidFrom",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetValidUntil() {
	_jsii_.InvokeVoid(
		s,
		"resetValidUntil",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetVolumeTags() {
	_jsii_.InvokeVoid(
		s,
		"resetVolumeTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) ResetWaitForFulfillment() {
	_jsii_.InvokeVoid(
		s,
		"resetWaitForFulfillment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequest) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequest) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequest) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequest) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequest) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

