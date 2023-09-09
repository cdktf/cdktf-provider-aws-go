// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2instancetype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/dataawsec2instancetype/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/data-sources/ec2_instance_type aws_ec2_instance_type}.
type DataAwsEc2InstanceType interface {
	cdktf.TerraformDataSource
	AutoRecoverySupported() cdktf.IResolvable
	BareMetal() cdktf.IResolvable
	BurstablePerformanceSupported() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CurrentGeneration() cdktf.IResolvable
	DedicatedHostsSupported() cdktf.IResolvable
	DefaultCores() *float64
	DefaultThreadsPerCore() *float64
	DefaultVcpus() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EbsEncryptionSupport() *string
	EbsNvmeSupport() *string
	EbsOptimizedSupport() *string
	EbsPerformanceBaselineBandwidth() *float64
	EbsPerformanceBaselineIops() *float64
	EbsPerformanceBaselineThroughput() *float64
	EbsPerformanceMaximumBandwidth() *float64
	EbsPerformanceMaximumIops() *float64
	EbsPerformanceMaximumThroughput() *float64
	EfaSupported() cdktf.IResolvable
	EnaSupport() *string
	EncryptionInTransitSupported() cdktf.IResolvable
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fpgas() DataAwsEc2InstanceTypeFpgasList
	// Experimental.
	Fqn() *string
	FreeTierEligible() cdktf.IResolvable
	// Experimental.
	FriendlyUniqueId() *string
	Gpus() DataAwsEc2InstanceTypeGpusList
	HibernationSupported() cdktf.IResolvable
	Hypervisor() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InferenceAccelerators() DataAwsEc2InstanceTypeInferenceAcceleratorsList
	InstanceDisks() DataAwsEc2InstanceTypeInstanceDisksList
	InstanceStorageSupported() cdktf.IResolvable
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	Ipv6Supported() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumIpv4AddressesPerInterface() *float64
	MaximumIpv6AddressesPerInterface() *float64
	MaximumNetworkInterfaces() *float64
	MemorySize() *float64
	NetworkPerformance() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SupportedArchitectures() *[]*string
	SupportedPlacementStrategies() *[]*string
	SupportedRootDeviceTypes() *[]*string
	SupportedUsagesClasses() *[]*string
	SupportedVirtualizationTypes() *[]*string
	SustainedClockSpeed() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAwsEc2InstanceTypeTimeoutsOutputReference
	TimeoutsInput() interface{}
	TotalFpgaMemory() *float64
	TotalGpuMemory() *float64
	TotalInstanceStorage() *float64
	ValidCores() *[]*float64
	ValidThreadsPerCore() *[]*float64
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
	PutTimeouts(value *DataAwsEc2InstanceTypeTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DataAwsEc2InstanceType
type jsiiProxy_DataAwsEc2InstanceType struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsEc2InstanceType) AutoRecoverySupported() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoRecoverySupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) BareMetal() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"bareMetal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) BurstablePerformanceSupported() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"burstablePerformanceSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) CurrentGeneration() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"currentGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) DedicatedHostsSupported() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"dedicatedHostsSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) DefaultCores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultCores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) DefaultThreadsPerCore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultThreadsPerCore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) DefaultVcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultVcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) EbsEncryptionSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) EbsNvmeSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsNvmeSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) EbsOptimizedSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsOptimizedSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) EbsPerformanceBaselineBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsPerformanceBaselineBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) EbsPerformanceBaselineIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsPerformanceBaselineIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) EbsPerformanceBaselineThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsPerformanceBaselineThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) EbsPerformanceMaximumBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsPerformanceMaximumBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) EbsPerformanceMaximumIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsPerformanceMaximumIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) EbsPerformanceMaximumThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsPerformanceMaximumThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) EfaSupported() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"efaSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) EnaSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enaSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) EncryptionInTransitSupported() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"encryptionInTransitSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) Fpgas() DataAwsEc2InstanceTypeFpgasList {
	var returns DataAwsEc2InstanceTypeFpgasList
	_jsii_.Get(
		j,
		"fpgas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) FreeTierEligible() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"freeTierEligible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) Gpus() DataAwsEc2InstanceTypeGpusList {
	var returns DataAwsEc2InstanceTypeGpusList
	_jsii_.Get(
		j,
		"gpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) HibernationSupported() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hibernationSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) Hypervisor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hypervisor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) InferenceAccelerators() DataAwsEc2InstanceTypeInferenceAcceleratorsList {
	var returns DataAwsEc2InstanceTypeInferenceAcceleratorsList
	_jsii_.Get(
		j,
		"inferenceAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) InstanceDisks() DataAwsEc2InstanceTypeInstanceDisksList {
	var returns DataAwsEc2InstanceTypeInstanceDisksList
	_jsii_.Get(
		j,
		"instanceDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) InstanceStorageSupported() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"instanceStorageSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) Ipv6Supported() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ipv6Supported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) MaximumIpv4AddressesPerInterface() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumIpv4AddressesPerInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) MaximumIpv6AddressesPerInterface() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumIpv6AddressesPerInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) MaximumNetworkInterfaces() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumNetworkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) NetworkPerformance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) SupportedArchitectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedArchitectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) SupportedPlacementStrategies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedPlacementStrategies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) SupportedRootDeviceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedRootDeviceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) SupportedUsagesClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedUsagesClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) SupportedVirtualizationTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedVirtualizationTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) SustainedClockSpeed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sustainedClockSpeed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) Timeouts() DataAwsEc2InstanceTypeTimeoutsOutputReference {
	var returns DataAwsEc2InstanceTypeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) TotalFpgaMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalFpgaMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) TotalGpuMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalGpuMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) TotalInstanceStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalInstanceStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) ValidCores() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"validCores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2InstanceType) ValidThreadsPerCore() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"validThreadsPerCore",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/data-sources/ec2_instance_type aws_ec2_instance_type} Data Source.
func NewDataAwsEc2InstanceType(scope constructs.Construct, id *string, config *DataAwsEc2InstanceTypeConfig) DataAwsEc2InstanceType {
	_init_.Initialize()

	if err := validateNewDataAwsEc2InstanceTypeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsEc2InstanceType{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEc2InstanceType.DataAwsEc2InstanceType",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/data-sources/ec2_instance_type aws_ec2_instance_type} Data Source.
func NewDataAwsEc2InstanceType_Override(d DataAwsEc2InstanceType, scope constructs.Construct, id *string, config *DataAwsEc2InstanceTypeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEc2InstanceType.DataAwsEc2InstanceType",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsEc2InstanceType)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2InstanceType)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2InstanceType)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2InstanceType)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2InstanceType)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2InstanceType)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2InstanceType)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
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
func DataAwsEc2InstanceType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsEc2InstanceType_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsEc2InstanceType.DataAwsEc2InstanceType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsEc2InstanceType_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsEc2InstanceType_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsEc2InstanceType.DataAwsEc2InstanceType",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsEc2InstanceType_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsEc2InstanceType_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsEc2InstanceType.DataAwsEc2InstanceType",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsEc2InstanceType_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsEc2InstanceType.DataAwsEc2InstanceType",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsEc2InstanceType) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsEc2InstanceType) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsEc2InstanceType) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsEc2InstanceType) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsEc2InstanceType) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsEc2InstanceType) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsEc2InstanceType) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsEc2InstanceType) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsEc2InstanceType) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsEc2InstanceType) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsEc2InstanceType) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsEc2InstanceType) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsEc2InstanceType) PutTimeouts(value *DataAwsEc2InstanceTypeTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsEc2InstanceType) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEc2InstanceType) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEc2InstanceType) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEc2InstanceType) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2InstanceType) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2InstanceType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2InstanceType) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

