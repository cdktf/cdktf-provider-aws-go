// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/gameliftfleet/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/gamelift_fleet aws_gamelift_fleet}.
type GameliftFleet interface {
	cdktf.TerraformResource
	Arn() *string
	BuildArn() *string
	BuildId() *string
	SetBuildId(val *string)
	BuildIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateConfiguration() GameliftFleetCertificateConfigurationOutputReference
	CertificateConfigurationInput() *GameliftFleetCertificateConfiguration
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Ec2InboundPermission() GameliftFleetEc2InboundPermissionList
	Ec2InboundPermissionInput() interface{}
	Ec2InstanceType() *string
	SetEc2InstanceType(val *string)
	Ec2InstanceTypeInput() *string
	FleetType() *string
	SetFleetType(val *string)
	FleetTypeInput() *string
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
	InstanceRoleArn() *string
	SetInstanceRoleArn(val *string)
	InstanceRoleArnInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogPaths() *[]*string
	MetricGroups() *[]*string
	SetMetricGroups(val *[]*string)
	MetricGroupsInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NewGameSessionProtectionPolicy() *string
	SetNewGameSessionProtectionPolicy(val *string)
	NewGameSessionProtectionPolicyInput() *string
	// The tree node.
	Node() constructs.Node
	OperatingSystem() *string
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
	ResourceCreationLimitPolicy() GameliftFleetResourceCreationLimitPolicyOutputReference
	ResourceCreationLimitPolicyInput() *GameliftFleetResourceCreationLimitPolicy
	RuntimeConfiguration() GameliftFleetRuntimeConfigurationOutputReference
	RuntimeConfigurationInput() *GameliftFleetRuntimeConfiguration
	ScriptArn() *string
	ScriptId() *string
	SetScriptId(val *string)
	ScriptIdInput() *string
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
	Timeouts() GameliftFleetTimeoutsOutputReference
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
	PutCertificateConfiguration(value *GameliftFleetCertificateConfiguration)
	PutEc2InboundPermission(value interface{})
	PutResourceCreationLimitPolicy(value *GameliftFleetResourceCreationLimitPolicy)
	PutRuntimeConfiguration(value *GameliftFleetRuntimeConfiguration)
	PutTimeouts(value *GameliftFleetTimeouts)
	ResetBuildId()
	ResetCertificateConfiguration()
	ResetDescription()
	ResetEc2InboundPermission()
	ResetFleetType()
	ResetId()
	ResetInstanceRoleArn()
	ResetMetricGroups()
	ResetNewGameSessionProtectionPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceCreationLimitPolicy()
	ResetRuntimeConfiguration()
	ResetScriptId()
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

// The jsii proxy struct for GameliftFleet
type jsiiProxy_GameliftFleet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GameliftFleet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) BuildArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) BuildId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) BuildIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) CertificateConfiguration() GameliftFleetCertificateConfigurationOutputReference {
	var returns GameliftFleetCertificateConfigurationOutputReference
	_jsii_.Get(
		j,
		"certificateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) CertificateConfigurationInput() *GameliftFleetCertificateConfiguration {
	var returns *GameliftFleetCertificateConfiguration
	_jsii_.Get(
		j,
		"certificateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Ec2InboundPermission() GameliftFleetEc2InboundPermissionList {
	var returns GameliftFleetEc2InboundPermissionList
	_jsii_.Get(
		j,
		"ec2InboundPermission",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Ec2InboundPermissionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InboundPermissionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Ec2InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Ec2InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) FleetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) FleetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) InstanceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) InstanceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) LogPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) MetricGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) MetricGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) NewGameSessionProtectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newGameSessionProtectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) NewGameSessionProtectionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newGameSessionProtectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ResourceCreationLimitPolicy() GameliftFleetResourceCreationLimitPolicyOutputReference {
	var returns GameliftFleetResourceCreationLimitPolicyOutputReference
	_jsii_.Get(
		j,
		"resourceCreationLimitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ResourceCreationLimitPolicyInput() *GameliftFleetResourceCreationLimitPolicy {
	var returns *GameliftFleetResourceCreationLimitPolicy
	_jsii_.Get(
		j,
		"resourceCreationLimitPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) RuntimeConfiguration() GameliftFleetRuntimeConfigurationOutputReference {
	var returns GameliftFleetRuntimeConfigurationOutputReference
	_jsii_.Get(
		j,
		"runtimeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) RuntimeConfigurationInput() *GameliftFleetRuntimeConfiguration {
	var returns *GameliftFleetRuntimeConfiguration
	_jsii_.Get(
		j,
		"runtimeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ScriptArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ScriptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ScriptIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Timeouts() GameliftFleetTimeoutsOutputReference {
	var returns GameliftFleetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/gamelift_fleet aws_gamelift_fleet} Resource.
func NewGameliftFleet(scope constructs.Construct, id *string, config *GameliftFleetConfig) GameliftFleet {
	_init_.Initialize()

	if err := validateNewGameliftFleetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftFleet{}

	_jsii_.Create(
		"@cdktf/provider-aws.gameliftFleet.GameliftFleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/gamelift_fleet aws_gamelift_fleet} Resource.
func NewGameliftFleet_Override(g GameliftFleet, scope constructs.Construct, id *string, config *GameliftFleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.gameliftFleet.GameliftFleet",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GameliftFleet)SetBuildId(val *string) {
	if err := j.validateSetBuildIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildId",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetEc2InstanceType(val *string) {
	if err := j.validateSetEc2InstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2InstanceType",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetFleetType(val *string) {
	if err := j.validateSetFleetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fleetType",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetInstanceRoleArn(val *string) {
	if err := j.validateSetInstanceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceRoleArn",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetMetricGroups(val *[]*string) {
	if err := j.validateSetMetricGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricGroups",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetNewGameSessionProtectionPolicy(val *string) {
	if err := j.validateSetNewGameSessionProtectionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newGameSessionProtectionPolicy",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetScriptId(val *string) {
	if err := j.validateSetScriptIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptId",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
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
func GameliftFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftFleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.gameliftFleet.GameliftFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftFleet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftFleet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.gameliftFleet.GameliftFleet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftFleet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftFleet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.gameliftFleet.GameliftFleet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GameliftFleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.gameliftFleet.GameliftFleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GameliftFleet) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GameliftFleet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GameliftFleet) PutCertificateConfiguration(value *GameliftFleetCertificateConfiguration) {
	if err := g.validatePutCertificateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCertificateConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutEc2InboundPermission(value interface{}) {
	if err := g.validatePutEc2InboundPermissionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEc2InboundPermission",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutResourceCreationLimitPolicy(value *GameliftFleetResourceCreationLimitPolicy) {
	if err := g.validatePutResourceCreationLimitPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResourceCreationLimitPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutRuntimeConfiguration(value *GameliftFleetRuntimeConfiguration) {
	if err := g.validatePutRuntimeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRuntimeConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutTimeouts(value *GameliftFleetTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) ResetBuildId() {
	_jsii_.InvokeVoid(
		g,
		"resetBuildId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetCertificateConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetCertificateConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetEc2InboundPermission() {
	_jsii_.InvokeVoid(
		g,
		"resetEc2InboundPermission",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetFleetType() {
	_jsii_.InvokeVoid(
		g,
		"resetFleetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetInstanceRoleArn() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceRoleArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetMetricGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetMetricGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetNewGameSessionProtectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetNewGameSessionProtectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetResourceCreationLimitPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceCreationLimitPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetRuntimeConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetScriptId() {
	_jsii_.InvokeVoid(
		g,
		"resetScriptId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetTagsAll() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

