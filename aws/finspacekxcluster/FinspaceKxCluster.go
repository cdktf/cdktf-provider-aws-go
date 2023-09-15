// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/finspacekxcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.17.0/docs/resources/finspace_kx_cluster aws_finspace_kx_cluster}.
type FinspaceKxCluster interface {
	cdktf.TerraformResource
	Arn() *string
	AutoScalingConfiguration() FinspaceKxClusterAutoScalingConfigurationOutputReference
	AutoScalingConfigurationInput() *FinspaceKxClusterAutoScalingConfiguration
	AvailabilityZoneId() *string
	SetAvailabilityZoneId(val *string)
	AvailabilityZoneIdInput() *string
	AzMode() *string
	SetAzMode(val *string)
	AzModeInput() *string
	CacheStorageConfigurations() FinspaceKxClusterCacheStorageConfigurationsList
	CacheStorageConfigurationsInput() interface{}
	CapacityConfiguration() FinspaceKxClusterCapacityConfigurationOutputReference
	CapacityConfigurationInput() *FinspaceKxClusterCapacityConfiguration
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Code() FinspaceKxClusterCodeOutputReference
	CodeInput() *FinspaceKxClusterCode
	CommandLineArguments() *map[string]*string
	SetCommandLineArguments(val *map[string]*string)
	CommandLineArgumentsInput() *map[string]*string
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
	CreatedTimestamp() *string
	Database() FinspaceKxClusterDatabaseList
	DatabaseInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnvironmentId() *string
	SetEnvironmentId(val *string)
	EnvironmentIdInput() *string
	ExecutionRole() *string
	SetExecutionRole(val *string)
	ExecutionRoleInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InitializationScript() *string
	SetInitializationScript(val *string)
	InitializationScriptInput() *string
	LastModifiedTimestamp() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	ReleaseLabel() *string
	SetReleaseLabel(val *string)
	ReleaseLabelInput() *string
	SavedownStorageConfiguration() FinspaceKxClusterSavedownStorageConfigurationOutputReference
	SavedownStorageConfigurationInput() *FinspaceKxClusterSavedownStorageConfiguration
	Status() *string
	StatusReason() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() FinspaceKxClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VpcConfiguration() FinspaceKxClusterVpcConfigurationOutputReference
	VpcConfigurationInput() *FinspaceKxClusterVpcConfiguration
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
	PutAutoScalingConfiguration(value *FinspaceKxClusterAutoScalingConfiguration)
	PutCacheStorageConfigurations(value interface{})
	PutCapacityConfiguration(value *FinspaceKxClusterCapacityConfiguration)
	PutCode(value *FinspaceKxClusterCode)
	PutDatabase(value interface{})
	PutSavedownStorageConfiguration(value *FinspaceKxClusterSavedownStorageConfiguration)
	PutTimeouts(value *FinspaceKxClusterTimeouts)
	PutVpcConfiguration(value *FinspaceKxClusterVpcConfiguration)
	ResetAutoScalingConfiguration()
	ResetAvailabilityZoneId()
	ResetCacheStorageConfigurations()
	ResetCode()
	ResetCommandLineArguments()
	ResetDatabase()
	ResetDescription()
	ResetExecutionRole()
	ResetId()
	ResetInitializationScript()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSavedownStorageConfiguration()
	ResetTags()
	ResetTagsAll()
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

// The jsii proxy struct for FinspaceKxCluster
type jsiiProxy_FinspaceKxCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FinspaceKxCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) AutoScalingConfiguration() FinspaceKxClusterAutoScalingConfigurationOutputReference {
	var returns FinspaceKxClusterAutoScalingConfigurationOutputReference
	_jsii_.Get(
		j,
		"autoScalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) AutoScalingConfigurationInput() *FinspaceKxClusterAutoScalingConfiguration {
	var returns *FinspaceKxClusterAutoScalingConfiguration
	_jsii_.Get(
		j,
		"autoScalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) AvailabilityZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) AzMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) AzModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) CacheStorageConfigurations() FinspaceKxClusterCacheStorageConfigurationsList {
	var returns FinspaceKxClusterCacheStorageConfigurationsList
	_jsii_.Get(
		j,
		"cacheStorageConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) CacheStorageConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheStorageConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) CapacityConfiguration() FinspaceKxClusterCapacityConfigurationOutputReference {
	var returns FinspaceKxClusterCapacityConfigurationOutputReference
	_jsii_.Get(
		j,
		"capacityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) CapacityConfigurationInput() *FinspaceKxClusterCapacityConfiguration {
	var returns *FinspaceKxClusterCapacityConfiguration
	_jsii_.Get(
		j,
		"capacityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Code() FinspaceKxClusterCodeOutputReference {
	var returns FinspaceKxClusterCodeOutputReference
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) CodeInput() *FinspaceKxClusterCode {
	var returns *FinspaceKxClusterCode
	_jsii_.Get(
		j,
		"codeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) CommandLineArguments() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"commandLineArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) CommandLineArgumentsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"commandLineArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) CreatedTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Database() FinspaceKxClusterDatabaseList {
	var returns FinspaceKxClusterDatabaseList
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) EnvironmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) InitializationScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initializationScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) InitializationScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initializationScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) LastModifiedTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) ReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) SavedownStorageConfiguration() FinspaceKxClusterSavedownStorageConfigurationOutputReference {
	var returns FinspaceKxClusterSavedownStorageConfigurationOutputReference
	_jsii_.Get(
		j,
		"savedownStorageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) SavedownStorageConfigurationInput() *FinspaceKxClusterSavedownStorageConfiguration {
	var returns *FinspaceKxClusterSavedownStorageConfiguration
	_jsii_.Get(
		j,
		"savedownStorageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Timeouts() FinspaceKxClusterTimeoutsOutputReference {
	var returns FinspaceKxClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) VpcConfiguration() FinspaceKxClusterVpcConfigurationOutputReference {
	var returns FinspaceKxClusterVpcConfigurationOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxCluster) VpcConfigurationInput() *FinspaceKxClusterVpcConfiguration {
	var returns *FinspaceKxClusterVpcConfiguration
	_jsii_.Get(
		j,
		"vpcConfigurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.17.0/docs/resources/finspace_kx_cluster aws_finspace_kx_cluster} Resource.
func NewFinspaceKxCluster(scope constructs.Construct, id *string, config *FinspaceKxClusterConfig) FinspaceKxCluster {
	_init_.Initialize()

	if err := validateNewFinspaceKxClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FinspaceKxCluster{}

	_jsii_.Create(
		"@cdktf/provider-aws.finspaceKxCluster.FinspaceKxCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.17.0/docs/resources/finspace_kx_cluster aws_finspace_kx_cluster} Resource.
func NewFinspaceKxCluster_Override(f FinspaceKxCluster, scope constructs.Construct, id *string, config *FinspaceKxClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.finspaceKxCluster.FinspaceKxCluster",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetAvailabilityZoneId(val *string) {
	if err := j.validateSetAvailabilityZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneId",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetAzMode(val *string) {
	if err := j.validateSetAzModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azMode",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetCommandLineArguments(val *map[string]*string) {
	if err := j.validateSetCommandLineArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commandLineArguments",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetEnvironmentId(val *string) {
	if err := j.validateSetEnvironmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentId",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetInitializationScript(val *string) {
	if err := j.validateSetInitializationScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initializationScript",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetReleaseLabel(val *string) {
	if err := j.validateSetReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxCluster)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
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
func FinspaceKxCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFinspaceKxCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.finspaceKxCluster.FinspaceKxCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FinspaceKxCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFinspaceKxCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.finspaceKxCluster.FinspaceKxCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FinspaceKxCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFinspaceKxCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.finspaceKxCluster.FinspaceKxCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FinspaceKxCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.finspaceKxCluster.FinspaceKxCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FinspaceKxCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FinspaceKxCluster) PutAutoScalingConfiguration(value *FinspaceKxClusterAutoScalingConfiguration) {
	if err := f.validatePutAutoScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAutoScalingConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FinspaceKxCluster) PutCacheStorageConfigurations(value interface{}) {
	if err := f.validatePutCacheStorageConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putCacheStorageConfigurations",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FinspaceKxCluster) PutCapacityConfiguration(value *FinspaceKxClusterCapacityConfiguration) {
	if err := f.validatePutCapacityConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putCapacityConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FinspaceKxCluster) PutCode(value *FinspaceKxClusterCode) {
	if err := f.validatePutCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putCode",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FinspaceKxCluster) PutDatabase(value interface{}) {
	if err := f.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putDatabase",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FinspaceKxCluster) PutSavedownStorageConfiguration(value *FinspaceKxClusterSavedownStorageConfiguration) {
	if err := f.validatePutSavedownStorageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSavedownStorageConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FinspaceKxCluster) PutTimeouts(value *FinspaceKxClusterTimeouts) {
	if err := f.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FinspaceKxCluster) PutVpcConfiguration(value *FinspaceKxClusterVpcConfiguration) {
	if err := f.validatePutVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putVpcConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetAutoScalingConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetAutoScalingConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetAvailabilityZoneId() {
	_jsii_.InvokeVoid(
		f,
		"resetAvailabilityZoneId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetCacheStorageConfigurations() {
	_jsii_.InvokeVoid(
		f,
		"resetCacheStorageConfigurations",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetCode() {
	_jsii_.InvokeVoid(
		f,
		"resetCode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetCommandLineArguments() {
	_jsii_.InvokeVoid(
		f,
		"resetCommandLineArguments",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetDatabase() {
	_jsii_.InvokeVoid(
		f,
		"resetDatabase",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetExecutionRole() {
	_jsii_.InvokeVoid(
		f,
		"resetExecutionRole",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetInitializationScript() {
	_jsii_.InvokeVoid(
		f,
		"resetInitializationScript",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetSavedownStorageConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetSavedownStorageConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

