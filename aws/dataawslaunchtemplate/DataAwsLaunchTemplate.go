// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslaunchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dataawslaunchtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/launch_template aws_launch_template}.
type DataAwsLaunchTemplate interface {
	cdktf.TerraformDataSource
	Arn() *string
	BlockDeviceMappings() DataAwsLaunchTemplateBlockDeviceMappingsList
	CapacityReservationSpecification() DataAwsLaunchTemplateCapacityReservationSpecificationList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuOptions() DataAwsLaunchTemplateCpuOptionsList
	CreditSpecification() DataAwsLaunchTemplateCreditSpecificationList
	DefaultVersion() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	DisableApiStop() cdktf.IResolvable
	DisableApiTermination() cdktf.IResolvable
	EbsOptimized() *string
	ElasticGpuSpecifications() DataAwsLaunchTemplateElasticGpuSpecificationsList
	ElasticInferenceAccelerator() DataAwsLaunchTemplateElasticInferenceAcceleratorList
	EnclaveOptions() DataAwsLaunchTemplateEnclaveOptionsList
	Filter() DataAwsLaunchTemplateFilterList
	FilterInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HibernationOptions() DataAwsLaunchTemplateHibernationOptionsList
	IamInstanceProfile() DataAwsLaunchTemplateIamInstanceProfileList
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageId() *string
	InstanceInitiatedShutdownBehavior() *string
	InstanceMarketOptions() DataAwsLaunchTemplateInstanceMarketOptionsList
	InstanceRequirements() DataAwsLaunchTemplateInstanceRequirementsList
	InstanceType() *string
	KernelId() *string
	KeyName() *string
	LatestVersion() *float64
	LicenseSpecification() DataAwsLaunchTemplateLicenseSpecificationList
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenanceOptions() DataAwsLaunchTemplateMaintenanceOptionsList
	MetadataOptions() DataAwsLaunchTemplateMetadataOptionsList
	Monitoring() DataAwsLaunchTemplateMonitoringList
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterfaces() DataAwsLaunchTemplateNetworkInterfacesList
	// The tree node.
	Node() constructs.Node
	Placement() DataAwsLaunchTemplatePlacementList
	PrivateDnsNameOptions() DataAwsLaunchTemplatePrivateDnsNameOptionsList
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	RamDiskId() *string
	// Experimental.
	RawOverrides() interface{}
	SecurityGroupNames() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TagSpecifications() DataAwsLaunchTemplateTagSpecificationsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAwsLaunchTemplateTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserData() *string
	VpcSecurityGroupIds() *[]*string
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
	PutFilter(value interface{})
	PutTimeouts(value *DataAwsLaunchTemplateTimeouts)
	ResetFilter()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAwsLaunchTemplate
type jsiiProxy_DataAwsLaunchTemplate struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) BlockDeviceMappings() DataAwsLaunchTemplateBlockDeviceMappingsList {
	var returns DataAwsLaunchTemplateBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) CapacityReservationSpecification() DataAwsLaunchTemplateCapacityReservationSpecificationList {
	var returns DataAwsLaunchTemplateCapacityReservationSpecificationList
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) CpuOptions() DataAwsLaunchTemplateCpuOptionsList {
	var returns DataAwsLaunchTemplateCpuOptionsList
	_jsii_.Get(
		j,
		"cpuOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) CreditSpecification() DataAwsLaunchTemplateCreditSpecificationList {
	var returns DataAwsLaunchTemplateCreditSpecificationList
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) DefaultVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) DisableApiStop() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableApiStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) DisableApiTermination() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableApiTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) EbsOptimized() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) ElasticGpuSpecifications() DataAwsLaunchTemplateElasticGpuSpecificationsList {
	var returns DataAwsLaunchTemplateElasticGpuSpecificationsList
	_jsii_.Get(
		j,
		"elasticGpuSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) ElasticInferenceAccelerator() DataAwsLaunchTemplateElasticInferenceAcceleratorList {
	var returns DataAwsLaunchTemplateElasticInferenceAcceleratorList
	_jsii_.Get(
		j,
		"elasticInferenceAccelerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) EnclaveOptions() DataAwsLaunchTemplateEnclaveOptionsList {
	var returns DataAwsLaunchTemplateEnclaveOptionsList
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Filter() DataAwsLaunchTemplateFilterList {
	var returns DataAwsLaunchTemplateFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) HibernationOptions() DataAwsLaunchTemplateHibernationOptionsList {
	var returns DataAwsLaunchTemplateHibernationOptionsList
	_jsii_.Get(
		j,
		"hibernationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) IamInstanceProfile() DataAwsLaunchTemplateIamInstanceProfileList {
	var returns DataAwsLaunchTemplateIamInstanceProfileList
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) InstanceInitiatedShutdownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) InstanceMarketOptions() DataAwsLaunchTemplateInstanceMarketOptionsList {
	var returns DataAwsLaunchTemplateInstanceMarketOptionsList
	_jsii_.Get(
		j,
		"instanceMarketOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) InstanceRequirements() DataAwsLaunchTemplateInstanceRequirementsList {
	var returns DataAwsLaunchTemplateInstanceRequirementsList
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) LatestVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) LicenseSpecification() DataAwsLaunchTemplateLicenseSpecificationList {
	var returns DataAwsLaunchTemplateLicenseSpecificationList
	_jsii_.Get(
		j,
		"licenseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) MaintenanceOptions() DataAwsLaunchTemplateMaintenanceOptionsList {
	var returns DataAwsLaunchTemplateMaintenanceOptionsList
	_jsii_.Get(
		j,
		"maintenanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) MetadataOptions() DataAwsLaunchTemplateMetadataOptionsList {
	var returns DataAwsLaunchTemplateMetadataOptionsList
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Monitoring() DataAwsLaunchTemplateMonitoringList {
	var returns DataAwsLaunchTemplateMonitoringList
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) NetworkInterfaces() DataAwsLaunchTemplateNetworkInterfacesList {
	var returns DataAwsLaunchTemplateNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Placement() DataAwsLaunchTemplatePlacementList {
	var returns DataAwsLaunchTemplatePlacementList
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) PrivateDnsNameOptions() DataAwsLaunchTemplatePrivateDnsNameOptionsList {
	var returns DataAwsLaunchTemplatePrivateDnsNameOptionsList
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) RamDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) SecurityGroupNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) TagSpecifications() DataAwsLaunchTemplateTagSpecificationsList {
	var returns DataAwsLaunchTemplateTagSpecificationsList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) Timeouts() DataAwsLaunchTemplateTimeoutsOutputReference {
	var returns DataAwsLaunchTemplateTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplate) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/launch_template aws_launch_template} Data Source.
func NewDataAwsLaunchTemplate(scope constructs.Construct, id *string, config *DataAwsLaunchTemplateConfig) DataAwsLaunchTemplate {
	_init_.Initialize()

	if err := validateNewDataAwsLaunchTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsLaunchTemplate{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsLaunchTemplate.DataAwsLaunchTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/launch_template aws_launch_template} Data Source.
func NewDataAwsLaunchTemplate_Override(d DataAwsLaunchTemplate, scope constructs.Construct, id *string, config *DataAwsLaunchTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsLaunchTemplate.DataAwsLaunchTemplate",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplate)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsLaunchTemplate resource upon running "cdktf plan <stack-name>".
func DataAwsLaunchTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsLaunchTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsLaunchTemplate.DataAwsLaunchTemplate",
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
func DataAwsLaunchTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsLaunchTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsLaunchTemplate.DataAwsLaunchTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsLaunchTemplate_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsLaunchTemplate_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsLaunchTemplate.DataAwsLaunchTemplate",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsLaunchTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsLaunchTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsLaunchTemplate.DataAwsLaunchTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLaunchTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsLaunchTemplate.DataAwsLaunchTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsLaunchTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsLaunchTemplate) PutFilter(value interface{}) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsLaunchTemplate) PutTimeouts(value *DataAwsLaunchTemplateTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsLaunchTemplate) ResetFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLaunchTemplate) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLaunchTemplate) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLaunchTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLaunchTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLaunchTemplate) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLaunchTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

