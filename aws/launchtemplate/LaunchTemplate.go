// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package launchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/launchtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/launch_template aws_launch_template}.
type LaunchTemplate interface {
	cdktf.TerraformResource
	Arn() *string
	BlockDeviceMappings() LaunchTemplateBlockDeviceMappingsList
	BlockDeviceMappingsInput() interface{}
	CapacityReservationSpecification() LaunchTemplateCapacityReservationSpecificationOutputReference
	CapacityReservationSpecificationInput() *LaunchTemplateCapacityReservationSpecification
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
	CpuOptions() LaunchTemplateCpuOptionsOutputReference
	CpuOptionsInput() *LaunchTemplateCpuOptions
	CreditSpecification() LaunchTemplateCreditSpecificationOutputReference
	CreditSpecificationInput() *LaunchTemplateCreditSpecification
	DefaultVersion() *float64
	SetDefaultVersion(val *float64)
	DefaultVersionInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableApiStop() interface{}
	SetDisableApiStop(val interface{})
	DisableApiStopInput() interface{}
	DisableApiTermination() interface{}
	SetDisableApiTermination(val interface{})
	DisableApiTerminationInput() interface{}
	EbsOptimized() *string
	SetEbsOptimized(val *string)
	EbsOptimizedInput() *string
	ElasticGpuSpecifications() LaunchTemplateElasticGpuSpecificationsList
	ElasticGpuSpecificationsInput() interface{}
	ElasticInferenceAccelerator() LaunchTemplateElasticInferenceAcceleratorOutputReference
	ElasticInferenceAcceleratorInput() *LaunchTemplateElasticInferenceAccelerator
	EnclaveOptions() LaunchTemplateEnclaveOptionsOutputReference
	EnclaveOptionsInput() *LaunchTemplateEnclaveOptions
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HibernationOptions() LaunchTemplateHibernationOptionsOutputReference
	HibernationOptionsInput() *LaunchTemplateHibernationOptions
	IamInstanceProfile() LaunchTemplateIamInstanceProfileOutputReference
	IamInstanceProfileInput() *LaunchTemplateIamInstanceProfile
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	InstanceInitiatedShutdownBehavior() *string
	SetInstanceInitiatedShutdownBehavior(val *string)
	InstanceInitiatedShutdownBehaviorInput() *string
	InstanceMarketOptions() LaunchTemplateInstanceMarketOptionsOutputReference
	InstanceMarketOptionsInput() *LaunchTemplateInstanceMarketOptions
	InstanceRequirements() LaunchTemplateInstanceRequirementsOutputReference
	InstanceRequirementsInput() *LaunchTemplateInstanceRequirements
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	KernelId() *string
	SetKernelId(val *string)
	KernelIdInput() *string
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	LatestVersion() *float64
	LicenseSpecification() LaunchTemplateLicenseSpecificationList
	LicenseSpecificationInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenanceOptions() LaunchTemplateMaintenanceOptionsOutputReference
	MaintenanceOptionsInput() *LaunchTemplateMaintenanceOptions
	MetadataOptions() LaunchTemplateMetadataOptionsOutputReference
	MetadataOptionsInput() *LaunchTemplateMetadataOptions
	Monitoring() LaunchTemplateMonitoringOutputReference
	MonitoringInput() *LaunchTemplateMonitoring
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	NetworkInterfaces() LaunchTemplateNetworkInterfacesList
	NetworkInterfacesInput() interface{}
	// The tree node.
	Node() constructs.Node
	Placement() LaunchTemplatePlacementOutputReference
	PlacementInput() *LaunchTemplatePlacement
	PrivateDnsNameOptions() LaunchTemplatePrivateDnsNameOptionsOutputReference
	PrivateDnsNameOptionsInput() *LaunchTemplatePrivateDnsNameOptions
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	RamDiskId() *string
	SetRamDiskId(val *string)
	RamDiskIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	SecurityGroupNames() *[]*string
	SetSecurityGroupNames(val *[]*string)
	SecurityGroupNamesInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TagSpecifications() LaunchTemplateTagSpecificationsList
	TagSpecificationsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdateDefaultVersion() interface{}
	SetUpdateDefaultVersion(val interface{})
	UpdateDefaultVersionInput() interface{}
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	VpcSecurityGroupIdsInput() *[]*string
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBlockDeviceMappings(value interface{})
	PutCapacityReservationSpecification(value *LaunchTemplateCapacityReservationSpecification)
	PutCpuOptions(value *LaunchTemplateCpuOptions)
	PutCreditSpecification(value *LaunchTemplateCreditSpecification)
	PutElasticGpuSpecifications(value interface{})
	PutElasticInferenceAccelerator(value *LaunchTemplateElasticInferenceAccelerator)
	PutEnclaveOptions(value *LaunchTemplateEnclaveOptions)
	PutHibernationOptions(value *LaunchTemplateHibernationOptions)
	PutIamInstanceProfile(value *LaunchTemplateIamInstanceProfile)
	PutInstanceMarketOptions(value *LaunchTemplateInstanceMarketOptions)
	PutInstanceRequirements(value *LaunchTemplateInstanceRequirements)
	PutLicenseSpecification(value interface{})
	PutMaintenanceOptions(value *LaunchTemplateMaintenanceOptions)
	PutMetadataOptions(value *LaunchTemplateMetadataOptions)
	PutMonitoring(value *LaunchTemplateMonitoring)
	PutNetworkInterfaces(value interface{})
	PutPlacement(value *LaunchTemplatePlacement)
	PutPrivateDnsNameOptions(value *LaunchTemplatePrivateDnsNameOptions)
	PutTagSpecifications(value interface{})
	ResetBlockDeviceMappings()
	ResetCapacityReservationSpecification()
	ResetCpuOptions()
	ResetCreditSpecification()
	ResetDefaultVersion()
	ResetDescription()
	ResetDisableApiStop()
	ResetDisableApiTermination()
	ResetEbsOptimized()
	ResetElasticGpuSpecifications()
	ResetElasticInferenceAccelerator()
	ResetEnclaveOptions()
	ResetHibernationOptions()
	ResetIamInstanceProfile()
	ResetId()
	ResetImageId()
	ResetInstanceInitiatedShutdownBehavior()
	ResetInstanceMarketOptions()
	ResetInstanceRequirements()
	ResetInstanceType()
	ResetKernelId()
	ResetKeyName()
	ResetLicenseSpecification()
	ResetMaintenanceOptions()
	ResetMetadataOptions()
	ResetMonitoring()
	ResetName()
	ResetNamePrefix()
	ResetNetworkInterfaces()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacement()
	ResetPrivateDnsNameOptions()
	ResetRamDiskId()
	ResetSecurityGroupNames()
	ResetTags()
	ResetTagsAll()
	ResetTagSpecifications()
	ResetUpdateDefaultVersion()
	ResetUserData()
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

// The jsii proxy struct for LaunchTemplate
type jsiiProxy_LaunchTemplate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LaunchTemplate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) BlockDeviceMappings() LaunchTemplateBlockDeviceMappingsList {
	var returns LaunchTemplateBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) BlockDeviceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) CapacityReservationSpecification() LaunchTemplateCapacityReservationSpecificationOutputReference {
	var returns LaunchTemplateCapacityReservationSpecificationOutputReference
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) CapacityReservationSpecificationInput() *LaunchTemplateCapacityReservationSpecification {
	var returns *LaunchTemplateCapacityReservationSpecification
	_jsii_.Get(
		j,
		"capacityReservationSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) CpuOptions() LaunchTemplateCpuOptionsOutputReference {
	var returns LaunchTemplateCpuOptionsOutputReference
	_jsii_.Get(
		j,
		"cpuOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) CpuOptionsInput() *LaunchTemplateCpuOptions {
	var returns *LaunchTemplateCpuOptions
	_jsii_.Get(
		j,
		"cpuOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) CreditSpecification() LaunchTemplateCreditSpecificationOutputReference {
	var returns LaunchTemplateCreditSpecificationOutputReference
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) CreditSpecificationInput() *LaunchTemplateCreditSpecification {
	var returns *LaunchTemplateCreditSpecification
	_jsii_.Get(
		j,
		"creditSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) DefaultVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) DefaultVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) DisableApiStop() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) DisableApiStopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) DisableApiTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) DisableApiTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) EbsOptimized() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) EbsOptimizedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) ElasticGpuSpecifications() LaunchTemplateElasticGpuSpecificationsList {
	var returns LaunchTemplateElasticGpuSpecificationsList
	_jsii_.Get(
		j,
		"elasticGpuSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) ElasticGpuSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticGpuSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) ElasticInferenceAccelerator() LaunchTemplateElasticInferenceAcceleratorOutputReference {
	var returns LaunchTemplateElasticInferenceAcceleratorOutputReference
	_jsii_.Get(
		j,
		"elasticInferenceAccelerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) ElasticInferenceAcceleratorInput() *LaunchTemplateElasticInferenceAccelerator {
	var returns *LaunchTemplateElasticInferenceAccelerator
	_jsii_.Get(
		j,
		"elasticInferenceAcceleratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) EnclaveOptions() LaunchTemplateEnclaveOptionsOutputReference {
	var returns LaunchTemplateEnclaveOptionsOutputReference
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) EnclaveOptionsInput() *LaunchTemplateEnclaveOptions {
	var returns *LaunchTemplateEnclaveOptions
	_jsii_.Get(
		j,
		"enclaveOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) HibernationOptions() LaunchTemplateHibernationOptionsOutputReference {
	var returns LaunchTemplateHibernationOptionsOutputReference
	_jsii_.Get(
		j,
		"hibernationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) HibernationOptionsInput() *LaunchTemplateHibernationOptions {
	var returns *LaunchTemplateHibernationOptions
	_jsii_.Get(
		j,
		"hibernationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) IamInstanceProfile() LaunchTemplateIamInstanceProfileOutputReference {
	var returns LaunchTemplateIamInstanceProfileOutputReference
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) IamInstanceProfileInput() *LaunchTemplateIamInstanceProfile {
	var returns *LaunchTemplateIamInstanceProfile
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) InstanceInitiatedShutdownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) InstanceInitiatedShutdownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) InstanceMarketOptions() LaunchTemplateInstanceMarketOptionsOutputReference {
	var returns LaunchTemplateInstanceMarketOptionsOutputReference
	_jsii_.Get(
		j,
		"instanceMarketOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) InstanceMarketOptionsInput() *LaunchTemplateInstanceMarketOptions {
	var returns *LaunchTemplateInstanceMarketOptions
	_jsii_.Get(
		j,
		"instanceMarketOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) InstanceRequirements() LaunchTemplateInstanceRequirementsOutputReference {
	var returns LaunchTemplateInstanceRequirementsOutputReference
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) InstanceRequirementsInput() *LaunchTemplateInstanceRequirements {
	var returns *LaunchTemplateInstanceRequirements
	_jsii_.Get(
		j,
		"instanceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) KernelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) LatestVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) LicenseSpecification() LaunchTemplateLicenseSpecificationList {
	var returns LaunchTemplateLicenseSpecificationList
	_jsii_.Get(
		j,
		"licenseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) LicenseSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) MaintenanceOptions() LaunchTemplateMaintenanceOptionsOutputReference {
	var returns LaunchTemplateMaintenanceOptionsOutputReference
	_jsii_.Get(
		j,
		"maintenanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) MaintenanceOptionsInput() *LaunchTemplateMaintenanceOptions {
	var returns *LaunchTemplateMaintenanceOptions
	_jsii_.Get(
		j,
		"maintenanceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) MetadataOptions() LaunchTemplateMetadataOptionsOutputReference {
	var returns LaunchTemplateMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) MetadataOptionsInput() *LaunchTemplateMetadataOptions {
	var returns *LaunchTemplateMetadataOptions
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) Monitoring() LaunchTemplateMonitoringOutputReference {
	var returns LaunchTemplateMonitoringOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) MonitoringInput() *LaunchTemplateMonitoring {
	var returns *LaunchTemplateMonitoring
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) NetworkInterfaces() LaunchTemplateNetworkInterfacesList {
	var returns LaunchTemplateNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) NetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) Placement() LaunchTemplatePlacementOutputReference {
	var returns LaunchTemplatePlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) PlacementInput() *LaunchTemplatePlacement {
	var returns *LaunchTemplatePlacement
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) PrivateDnsNameOptions() LaunchTemplatePrivateDnsNameOptionsOutputReference {
	var returns LaunchTemplatePrivateDnsNameOptionsOutputReference
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) PrivateDnsNameOptionsInput() *LaunchTemplatePrivateDnsNameOptions {
	var returns *LaunchTemplatePrivateDnsNameOptions
	_jsii_.Get(
		j,
		"privateDnsNameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) RamDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) RamDiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) SecurityGroupNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) SecurityGroupNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) TagSpecifications() LaunchTemplateTagSpecificationsList {
	var returns LaunchTemplateTagSpecificationsList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) TagSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) UpdateDefaultVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateDefaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) UpdateDefaultVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateDefaultVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplate) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/launch_template aws_launch_template} Resource.
func NewLaunchTemplate(scope constructs.Construct, id *string, config *LaunchTemplateConfig) LaunchTemplate {
	_init_.Initialize()

	if err := validateNewLaunchTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LaunchTemplate{}

	_jsii_.Create(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/launch_template aws_launch_template} Resource.
func NewLaunchTemplate_Override(l LaunchTemplate, scope constructs.Construct, id *string, config *LaunchTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplate",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetDefaultVersion(val *float64) {
	if err := j.validateSetDefaultVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultVersion",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetDisableApiStop(val interface{}) {
	if err := j.validateSetDisableApiStopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiStop",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetDisableApiTermination(val interface{}) {
	if err := j.validateSetDisableApiTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiTermination",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetEbsOptimized(val *string) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetInstanceInitiatedShutdownBehavior(val *string) {
	if err := j.validateSetInstanceInitiatedShutdownBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInitiatedShutdownBehavior",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetKernelId(val *string) {
	if err := j.validateSetKernelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kernelId",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetRamDiskId(val *string) {
	if err := j.validateSetRamDiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ramDiskId",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetSecurityGroupNames(val *[]*string) {
	if err := j.validateSetSecurityGroupNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupNames",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetUpdateDefaultVersion(val interface{}) {
	if err := j.validateSetUpdateDefaultVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateDefaultVersion",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplate)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTF code for importing a LaunchTemplate resource upon running "cdktf plan <stack-name>".
func LaunchTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLaunchTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplate",
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
func LaunchTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLaunchTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LaunchTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLaunchTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LaunchTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLaunchTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LaunchTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LaunchTemplate) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LaunchTemplate) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LaunchTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplate) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplate) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LaunchTemplate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LaunchTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutBlockDeviceMappings(value interface{}) {
	if err := l.validatePutBlockDeviceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putBlockDeviceMappings",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutCapacityReservationSpecification(value *LaunchTemplateCapacityReservationSpecification) {
	if err := l.validatePutCapacityReservationSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCapacityReservationSpecification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutCpuOptions(value *LaunchTemplateCpuOptions) {
	if err := l.validatePutCpuOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCpuOptions",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutCreditSpecification(value *LaunchTemplateCreditSpecification) {
	if err := l.validatePutCreditSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCreditSpecification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutElasticGpuSpecifications(value interface{}) {
	if err := l.validatePutElasticGpuSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putElasticGpuSpecifications",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutElasticInferenceAccelerator(value *LaunchTemplateElasticInferenceAccelerator) {
	if err := l.validatePutElasticInferenceAcceleratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putElasticInferenceAccelerator",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutEnclaveOptions(value *LaunchTemplateEnclaveOptions) {
	if err := l.validatePutEnclaveOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putEnclaveOptions",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutHibernationOptions(value *LaunchTemplateHibernationOptions) {
	if err := l.validatePutHibernationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHibernationOptions",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutIamInstanceProfile(value *LaunchTemplateIamInstanceProfile) {
	if err := l.validatePutIamInstanceProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIamInstanceProfile",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutInstanceMarketOptions(value *LaunchTemplateInstanceMarketOptions) {
	if err := l.validatePutInstanceMarketOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putInstanceMarketOptions",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutInstanceRequirements(value *LaunchTemplateInstanceRequirements) {
	if err := l.validatePutInstanceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putInstanceRequirements",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutLicenseSpecification(value interface{}) {
	if err := l.validatePutLicenseSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLicenseSpecification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutMaintenanceOptions(value *LaunchTemplateMaintenanceOptions) {
	if err := l.validatePutMaintenanceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMaintenanceOptions",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutMetadataOptions(value *LaunchTemplateMetadataOptions) {
	if err := l.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutMonitoring(value *LaunchTemplateMonitoring) {
	if err := l.validatePutMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMonitoring",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutNetworkInterfaces(value interface{}) {
	if err := l.validatePutNetworkInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putNetworkInterfaces",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutPlacement(value *LaunchTemplatePlacement) {
	if err := l.validatePutPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPlacement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutPrivateDnsNameOptions(value *LaunchTemplatePrivateDnsNameOptions) {
	if err := l.validatePutPrivateDnsNameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPrivateDnsNameOptions",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) PutTagSpecifications(value interface{}) {
	if err := l.validatePutTagSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTagSpecifications",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetBlockDeviceMappings() {
	_jsii_.InvokeVoid(
		l,
		"resetBlockDeviceMappings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetCapacityReservationSpecification() {
	_jsii_.InvokeVoid(
		l,
		"resetCapacityReservationSpecification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetCpuOptions() {
	_jsii_.InvokeVoid(
		l,
		"resetCpuOptions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetCreditSpecification() {
	_jsii_.InvokeVoid(
		l,
		"resetCreditSpecification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetDefaultVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetDisableApiStop() {
	_jsii_.InvokeVoid(
		l,
		"resetDisableApiStop",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetDisableApiTermination() {
	_jsii_.InvokeVoid(
		l,
		"resetDisableApiTermination",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		l,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetElasticGpuSpecifications() {
	_jsii_.InvokeVoid(
		l,
		"resetElasticGpuSpecifications",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetElasticInferenceAccelerator() {
	_jsii_.InvokeVoid(
		l,
		"resetElasticInferenceAccelerator",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetEnclaveOptions() {
	_jsii_.InvokeVoid(
		l,
		"resetEnclaveOptions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetHibernationOptions() {
	_jsii_.InvokeVoid(
		l,
		"resetHibernationOptions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		l,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetImageId() {
	_jsii_.InvokeVoid(
		l,
		"resetImageId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetInstanceInitiatedShutdownBehavior() {
	_jsii_.InvokeVoid(
		l,
		"resetInstanceInitiatedShutdownBehavior",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetInstanceMarketOptions() {
	_jsii_.InvokeVoid(
		l,
		"resetInstanceMarketOptions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetInstanceRequirements() {
	_jsii_.InvokeVoid(
		l,
		"resetInstanceRequirements",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetInstanceType() {
	_jsii_.InvokeVoid(
		l,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetKernelId() {
	_jsii_.InvokeVoid(
		l,
		"resetKernelId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetKeyName() {
	_jsii_.InvokeVoid(
		l,
		"resetKeyName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetLicenseSpecification() {
	_jsii_.InvokeVoid(
		l,
		"resetLicenseSpecification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetMaintenanceOptions() {
	_jsii_.InvokeVoid(
		l,
		"resetMaintenanceOptions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		l,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetMonitoring() {
	_jsii_.InvokeVoid(
		l,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetName() {
	_jsii_.InvokeVoid(
		l,
		"resetName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		l,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetNetworkInterfaces() {
	_jsii_.InvokeVoid(
		l,
		"resetNetworkInterfaces",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetPlacement() {
	_jsii_.InvokeVoid(
		l,
		"resetPlacement",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetPrivateDnsNameOptions() {
	_jsii_.InvokeVoid(
		l,
		"resetPrivateDnsNameOptions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetRamDiskId() {
	_jsii_.InvokeVoid(
		l,
		"resetRamDiskId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetSecurityGroupNames() {
	_jsii_.InvokeVoid(
		l,
		"resetSecurityGroupNames",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetTagsAll() {
	_jsii_.InvokeVoid(
		l,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetTagSpecifications() {
	_jsii_.InvokeVoid(
		l,
		"resetTagSpecifications",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetUpdateDefaultVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetUpdateDefaultVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetUserData() {
	_jsii_.InvokeVoid(
		l,
		"resetUserData",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		l,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

