package instance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/instance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/instance aws_instance}.
type Instance interface {
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
	CapacityReservationSpecification() InstanceCapacityReservationSpecificationOutputReference
	CapacityReservationSpecificationInput() *InstanceCapacityReservationSpecification
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CpuCoreCount() *float64
	SetCpuCoreCount(val *float64)
	CpuCoreCountInput() *float64
	CpuThreadsPerCore() *float64
	SetCpuThreadsPerCore(val *float64)
	CpuThreadsPerCoreInput() *float64
	CreditSpecification() InstanceCreditSpecificationOutputReference
	CreditSpecificationInput() *InstanceCreditSpecification
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
	EbsBlockDevice() InstanceEbsBlockDeviceList
	EbsBlockDeviceInput() interface{}
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	EnclaveOptions() InstanceEnclaveOptionsOutputReference
	EnclaveOptionsInput() *InstanceEnclaveOptions
	EphemeralBlockDevice() InstanceEphemeralBlockDeviceList
	EphemeralBlockDeviceInput() interface{}
	FetchPasswordData() interface{}
	SetFetchPasswordData(val interface{})
	FetchPasswordDataInput() interface{}
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
	LaunchTemplate() InstanceLaunchTemplateOutputReference
	LaunchTemplateInput() *InstanceLaunchTemplate
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenanceOptions() InstanceMaintenanceOptionsOutputReference
	MaintenanceOptionsInput() *InstanceMaintenanceOptions
	MetadataOptions() InstanceMetadataOptionsOutputReference
	MetadataOptionsInput() *InstanceMetadataOptions
	Monitoring() interface{}
	SetMonitoring(val interface{})
	MonitoringInput() interface{}
	NetworkInterface() InstanceNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	OutpostArn() *string
	PasswordData() *string
	PlacementGroup() *string
	SetPlacementGroup(val *string)
	PlacementGroupInput() *string
	PlacementPartitionNumber() *float64
	SetPlacementPartitionNumber(val *float64)
	PlacementPartitionNumberInput() *float64
	PrimaryNetworkInterfaceId() *string
	PrivateDns() *string
	PrivateDnsNameOptions() InstancePrivateDnsNameOptionsOutputReference
	PrivateDnsNameOptionsInput() *InstancePrivateDnsNameOptions
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
	RootBlockDevice() InstanceRootBlockDeviceOutputReference
	RootBlockDeviceInput() *InstanceRootBlockDevice
	SecondaryPrivateIps() *[]*string
	SetSecondaryPrivateIps(val *[]*string)
	SecondaryPrivateIpsInput() *[]*string
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SourceDestCheck() interface{}
	SetSourceDestCheck(val interface{})
	SourceDestCheckInput() interface{}
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
	Timeouts() InstanceTimeoutsOutputReference
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
	VolumeTags() *map[string]*string
	SetVolumeTags(val *map[string]*string)
	VolumeTagsInput() *map[string]*string
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	VpcSecurityGroupIdsInput() *[]*string
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
	PutCapacityReservationSpecification(value *InstanceCapacityReservationSpecification)
	PutCreditSpecification(value *InstanceCreditSpecification)
	PutEbsBlockDevice(value interface{})
	PutEnclaveOptions(value *InstanceEnclaveOptions)
	PutEphemeralBlockDevice(value interface{})
	PutLaunchTemplate(value *InstanceLaunchTemplate)
	PutMaintenanceOptions(value *InstanceMaintenanceOptions)
	PutMetadataOptions(value *InstanceMetadataOptions)
	PutNetworkInterface(value interface{})
	PutPrivateDnsNameOptions(value *InstancePrivateDnsNameOptions)
	PutRootBlockDevice(value *InstanceRootBlockDevice)
	PutTimeouts(value *InstanceTimeouts)
	ResetAmi()
	ResetAssociatePublicIpAddress()
	ResetAvailabilityZone()
	ResetCapacityReservationSpecification()
	ResetCpuCoreCount()
	ResetCpuThreadsPerCore()
	ResetCreditSpecification()
	ResetDisableApiStop()
	ResetDisableApiTermination()
	ResetEbsBlockDevice()
	ResetEbsOptimized()
	ResetEnclaveOptions()
	ResetEphemeralBlockDevice()
	ResetFetchPasswordData()
	ResetHibernation()
	ResetHostId()
	ResetHostResourceGroupArn()
	ResetIamInstanceProfile()
	ResetId()
	ResetInstanceInitiatedShutdownBehavior()
	ResetInstanceType()
	ResetIpv6AddressCount()
	ResetIpv6Addresses()
	ResetKeyName()
	ResetLaunchTemplate()
	ResetMaintenanceOptions()
	ResetMetadataOptions()
	ResetMonitoring()
	ResetNetworkInterface()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacementGroup()
	ResetPlacementPartitionNumber()
	ResetPrivateDnsNameOptions()
	ResetPrivateIp()
	ResetRootBlockDevice()
	ResetSecondaryPrivateIps()
	ResetSecurityGroups()
	ResetSourceDestCheck()
	ResetSubnetId()
	ResetTags()
	ResetTagsAll()
	ResetTenancy()
	ResetTimeouts()
	ResetUserData()
	ResetUserDataBase64()
	ResetUserDataReplaceOnChange()
	ResetVolumeTags()
	ResetVpcSecurityGroupIds()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Instance
type jsiiProxy_Instance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Instance) Ami() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ami",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) AmiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) AssociatePublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) CapacityReservationSpecification() InstanceCapacityReservationSpecificationOutputReference {
	var returns InstanceCapacityReservationSpecificationOutputReference
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) CapacityReservationSpecificationInput() *InstanceCapacityReservationSpecification {
	var returns *InstanceCapacityReservationSpecification
	_jsii_.Get(
		j,
		"capacityReservationSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) CpuCoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) CpuThreadsPerCore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuThreadsPerCore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) CpuThreadsPerCoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuThreadsPerCoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) CreditSpecification() InstanceCreditSpecificationOutputReference {
	var returns InstanceCreditSpecificationOutputReference
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) CreditSpecificationInput() *InstanceCreditSpecification {
	var returns *InstanceCreditSpecification
	_jsii_.Get(
		j,
		"creditSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) DisableApiStop() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) DisableApiStopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) DisableApiTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) DisableApiTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) EbsBlockDevice() InstanceEbsBlockDeviceList {
	var returns InstanceEbsBlockDeviceList
	_jsii_.Get(
		j,
		"ebsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) EbsBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) EnclaveOptions() InstanceEnclaveOptionsOutputReference {
	var returns InstanceEnclaveOptionsOutputReference
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) EnclaveOptionsInput() *InstanceEnclaveOptions {
	var returns *InstanceEnclaveOptions
	_jsii_.Get(
		j,
		"enclaveOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) EphemeralBlockDevice() InstanceEphemeralBlockDeviceList {
	var returns InstanceEphemeralBlockDeviceList
	_jsii_.Get(
		j,
		"ephemeralBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) EphemeralBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) FetchPasswordData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchPasswordData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) FetchPasswordDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchPasswordDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Hibernation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) HibernationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) HostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) HostIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) HostResourceGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) HostResourceGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) InstanceInitiatedShutdownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) InstanceInitiatedShutdownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) InstanceState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Ipv6AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Ipv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Ipv6AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) LaunchTemplate() InstanceLaunchTemplateOutputReference {
	var returns InstanceLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) LaunchTemplateInput() *InstanceLaunchTemplate {
	var returns *InstanceLaunchTemplate
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) MaintenanceOptions() InstanceMaintenanceOptionsOutputReference {
	var returns InstanceMaintenanceOptionsOutputReference
	_jsii_.Get(
		j,
		"maintenanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) MaintenanceOptionsInput() *InstanceMaintenanceOptions {
	var returns *InstanceMaintenanceOptions
	_jsii_.Get(
		j,
		"maintenanceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) MetadataOptions() InstanceMetadataOptionsOutputReference {
	var returns InstanceMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) MetadataOptionsInput() *InstanceMetadataOptions {
	var returns *InstanceMetadataOptions
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Monitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) MonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) NetworkInterface() InstanceNetworkInterfaceList {
	var returns InstanceNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) OutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) PasswordData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) PlacementGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) PlacementPartitionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"placementPartitionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) PlacementPartitionNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"placementPartitionNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) PrimaryNetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryNetworkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) PrivateDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) PrivateDnsNameOptions() InstancePrivateDnsNameOptionsOutputReference {
	var returns InstancePrivateDnsNameOptionsOutputReference
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) PrivateDnsNameOptionsInput() *InstancePrivateDnsNameOptions {
	var returns *InstancePrivateDnsNameOptions
	_jsii_.Get(
		j,
		"privateDnsNameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) PublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) RootBlockDevice() InstanceRootBlockDeviceOutputReference {
	var returns InstanceRootBlockDeviceOutputReference
	_jsii_.Get(
		j,
		"rootBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) RootBlockDeviceInput() *InstanceRootBlockDevice {
	var returns *InstanceRootBlockDevice
	_jsii_.Get(
		j,
		"rootBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) SecondaryPrivateIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryPrivateIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) SecondaryPrivateIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryPrivateIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) SourceDestCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDestCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) SourceDestCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDestCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) TenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) Timeouts() InstanceTimeoutsOutputReference {
	var returns InstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) UserDataBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) UserDataBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) UserDataReplaceOnChange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDataReplaceOnChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) UserDataReplaceOnChangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDataReplaceOnChangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) VolumeTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) VolumeTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Instance) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/instance aws_instance} Resource.
func NewInstance(scope constructs.Construct, id *string, config *InstanceConfig) Instance {
	_init_.Initialize()

	if err := validateNewInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Instance{}

	_jsii_.Create(
		"@cdktf/provider-aws.instance.Instance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/instance aws_instance} Resource.
func NewInstance_Override(i Instance, scope constructs.Construct, id *string, config *InstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.instance.Instance",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_Instance)SetAmi(val *string) {
	if err := j.validateSetAmiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ami",
		val,
	)
}

func (j *jsiiProxy_Instance)SetAssociatePublicIpAddress(val interface{}) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_Instance)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_Instance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Instance)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Instance)SetCpuCoreCount(val *float64) {
	if err := j.validateSetCpuCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCoreCount",
		val,
	)
}

func (j *jsiiProxy_Instance)SetCpuThreadsPerCore(val *float64) {
	if err := j.validateSetCpuThreadsPerCoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuThreadsPerCore",
		val,
	)
}

func (j *jsiiProxy_Instance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Instance)SetDisableApiStop(val interface{}) {
	if err := j.validateSetDisableApiStopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiStop",
		val,
	)
}

func (j *jsiiProxy_Instance)SetDisableApiTermination(val interface{}) {
	if err := j.validateSetDisableApiTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiTermination",
		val,
	)
}

func (j *jsiiProxy_Instance)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_Instance)SetFetchPasswordData(val interface{}) {
	if err := j.validateSetFetchPasswordDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fetchPasswordData",
		val,
	)
}

func (j *jsiiProxy_Instance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Instance)SetHibernation(val interface{}) {
	if err := j.validateSetHibernationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hibernation",
		val,
	)
}

func (j *jsiiProxy_Instance)SetHostId(val *string) {
	if err := j.validateSetHostIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostId",
		val,
	)
}

func (j *jsiiProxy_Instance)SetHostResourceGroupArn(val *string) {
	if err := j.validateSetHostResourceGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostResourceGroupArn",
		val,
	)
}

func (j *jsiiProxy_Instance)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_Instance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Instance)SetInstanceInitiatedShutdownBehavior(val *string) {
	if err := j.validateSetInstanceInitiatedShutdownBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInitiatedShutdownBehavior",
		val,
	)
}

func (j *jsiiProxy_Instance)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_Instance)SetIpv6AddressCount(val *float64) {
	if err := j.validateSetIpv6AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AddressCount",
		val,
	)
}

func (j *jsiiProxy_Instance)SetIpv6Addresses(val *[]*string) {
	if err := j.validateSetIpv6AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Addresses",
		val,
	)
}

func (j *jsiiProxy_Instance)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_Instance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Instance)SetMonitoring(val interface{}) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_Instance)SetPlacementGroup(val *string) {
	if err := j.validateSetPlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_Instance)SetPlacementPartitionNumber(val *float64) {
	if err := j.validateSetPlacementPartitionNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementPartitionNumber",
		val,
	)
}

func (j *jsiiProxy_Instance)SetPrivateIp(val *string) {
	if err := j.validateSetPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_Instance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Instance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Instance)SetSecondaryPrivateIps(val *[]*string) {
	if err := j.validateSetSecondaryPrivateIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryPrivateIps",
		val,
	)
}

func (j *jsiiProxy_Instance)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_Instance)SetSourceDestCheck(val interface{}) {
	if err := j.validateSetSourceDestCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDestCheck",
		val,
	)
}

func (j *jsiiProxy_Instance)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_Instance)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Instance)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_Instance)SetTenancy(val *string) {
	if err := j.validateSetTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_Instance)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_Instance)SetUserDataBase64(val *string) {
	if err := j.validateSetUserDataBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDataBase64",
		val,
	)
}

func (j *jsiiProxy_Instance)SetUserDataReplaceOnChange(val interface{}) {
	if err := j.validateSetUserDataReplaceOnChangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDataReplaceOnChange",
		val,
	)
}

func (j *jsiiProxy_Instance)SetVolumeTags(val *map[string]*string) {
	if err := j.validateSetVolumeTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeTags",
		val,
	)
}

func (j *jsiiProxy_Instance)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
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
func Instance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.instance.Instance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Instance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.instance.Instance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_Instance) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_Instance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Instance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Instance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Instance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Instance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Instance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Instance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Instance) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Instance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Instance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Instance) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_Instance) PutCapacityReservationSpecification(value *InstanceCapacityReservationSpecification) {
	if err := i.validatePutCapacityReservationSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCapacityReservationSpecification",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Instance) PutCreditSpecification(value *InstanceCreditSpecification) {
	if err := i.validatePutCreditSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCreditSpecification",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Instance) PutEbsBlockDevice(value interface{}) {
	if err := i.validatePutEbsBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEbsBlockDevice",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Instance) PutEnclaveOptions(value *InstanceEnclaveOptions) {
	if err := i.validatePutEnclaveOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEnclaveOptions",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Instance) PutEphemeralBlockDevice(value interface{}) {
	if err := i.validatePutEphemeralBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEphemeralBlockDevice",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Instance) PutLaunchTemplate(value *InstanceLaunchTemplate) {
	if err := i.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Instance) PutMaintenanceOptions(value *InstanceMaintenanceOptions) {
	if err := i.validatePutMaintenanceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMaintenanceOptions",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Instance) PutMetadataOptions(value *InstanceMetadataOptions) {
	if err := i.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Instance) PutNetworkInterface(value interface{}) {
	if err := i.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Instance) PutPrivateDnsNameOptions(value *InstancePrivateDnsNameOptions) {
	if err := i.validatePutPrivateDnsNameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPrivateDnsNameOptions",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Instance) PutRootBlockDevice(value *InstanceRootBlockDevice) {
	if err := i.validatePutRootBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRootBlockDevice",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Instance) PutTimeouts(value *InstanceTimeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Instance) ResetAmi() {
	_jsii_.InvokeVoid(
		i,
		"resetAmi",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		i,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		i,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetCapacityReservationSpecification() {
	_jsii_.InvokeVoid(
		i,
		"resetCapacityReservationSpecification",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetCpuCoreCount() {
	_jsii_.InvokeVoid(
		i,
		"resetCpuCoreCount",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetCpuThreadsPerCore() {
	_jsii_.InvokeVoid(
		i,
		"resetCpuThreadsPerCore",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetCreditSpecification() {
	_jsii_.InvokeVoid(
		i,
		"resetCreditSpecification",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetDisableApiStop() {
	_jsii_.InvokeVoid(
		i,
		"resetDisableApiStop",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetDisableApiTermination() {
	_jsii_.InvokeVoid(
		i,
		"resetDisableApiTermination",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetEbsBlockDevice() {
	_jsii_.InvokeVoid(
		i,
		"resetEbsBlockDevice",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		i,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetEnclaveOptions() {
	_jsii_.InvokeVoid(
		i,
		"resetEnclaveOptions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetEphemeralBlockDevice() {
	_jsii_.InvokeVoid(
		i,
		"resetEphemeralBlockDevice",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetFetchPasswordData() {
	_jsii_.InvokeVoid(
		i,
		"resetFetchPasswordData",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetHibernation() {
	_jsii_.InvokeVoid(
		i,
		"resetHibernation",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetHostId() {
	_jsii_.InvokeVoid(
		i,
		"resetHostId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetHostResourceGroupArn() {
	_jsii_.InvokeVoid(
		i,
		"resetHostResourceGroupArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		i,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetInstanceInitiatedShutdownBehavior() {
	_jsii_.InvokeVoid(
		i,
		"resetInstanceInitiatedShutdownBehavior",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetInstanceType() {
	_jsii_.InvokeVoid(
		i,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetIpv6AddressCount() {
	_jsii_.InvokeVoid(
		i,
		"resetIpv6AddressCount",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetIpv6Addresses() {
	_jsii_.InvokeVoid(
		i,
		"resetIpv6Addresses",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetKeyName() {
	_jsii_.InvokeVoid(
		i,
		"resetKeyName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		i,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetMaintenanceOptions() {
	_jsii_.InvokeVoid(
		i,
		"resetMaintenanceOptions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		i,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetMonitoring() {
	_jsii_.InvokeVoid(
		i,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		i,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetPlacementGroup() {
	_jsii_.InvokeVoid(
		i,
		"resetPlacementGroup",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetPlacementPartitionNumber() {
	_jsii_.InvokeVoid(
		i,
		"resetPlacementPartitionNumber",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetPrivateDnsNameOptions() {
	_jsii_.InvokeVoid(
		i,
		"resetPrivateDnsNameOptions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		i,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetRootBlockDevice() {
	_jsii_.InvokeVoid(
		i,
		"resetRootBlockDevice",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetSecondaryPrivateIps() {
	_jsii_.InvokeVoid(
		i,
		"resetSecondaryPrivateIps",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		i,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetSourceDestCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetSourceDestCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetSubnetId() {
	_jsii_.InvokeVoid(
		i,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetTagsAll() {
	_jsii_.InvokeVoid(
		i,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetTenancy() {
	_jsii_.InvokeVoid(
		i,
		"resetTenancy",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetUserData() {
	_jsii_.InvokeVoid(
		i,
		"resetUserData",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetUserDataBase64() {
	_jsii_.InvokeVoid(
		i,
		"resetUserDataBase64",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetUserDataReplaceOnChange() {
	_jsii_.InvokeVoid(
		i,
		"resetUserDataReplaceOnChange",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetVolumeTags() {
	_jsii_.InvokeVoid(
		i,
		"resetVolumeTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		i,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Instance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Instance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Instance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Instance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

