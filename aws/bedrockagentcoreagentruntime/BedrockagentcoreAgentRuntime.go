// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreagentruntime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/bedrockagentcoreagentruntime/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/bedrockagentcore_agent_runtime aws_bedrockagentcore_agent_runtime}.
type BedrockagentcoreAgentRuntime interface {
	cdktf.TerraformResource
	AgentRuntimeArn() *string
	AgentRuntimeArtifact() BedrockagentcoreAgentRuntimeAgentRuntimeArtifactList
	AgentRuntimeArtifactInput() interface{}
	AgentRuntimeId() *string
	AgentRuntimeName() *string
	SetAgentRuntimeName(val *string)
	AgentRuntimeNameInput() *string
	AgentRuntimeVersion() *string
	AuthorizerConfiguration() BedrockagentcoreAgentRuntimeAuthorizerConfigurationList
	AuthorizerConfigurationInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifecycleConfiguration() BedrockagentcoreAgentRuntimeLifecycleConfigurationList
	LifecycleConfigurationInput() interface{}
	NetworkConfiguration() BedrockagentcoreAgentRuntimeNetworkConfigurationList
	NetworkConfigurationInput() interface{}
	// The tree node.
	Node() constructs.Node
	ProtocolConfiguration() BedrockagentcoreAgentRuntimeProtocolConfigurationList
	ProtocolConfigurationInput() interface{}
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RequestHeaderConfiguration() BedrockagentcoreAgentRuntimeRequestHeaderConfigurationList
	RequestHeaderConfigurationInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktf.StringMap
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BedrockagentcoreAgentRuntimeTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkloadIdentityDetails() BedrockagentcoreAgentRuntimeWorkloadIdentityDetailsList
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
	PutAgentRuntimeArtifact(value interface{})
	PutAuthorizerConfiguration(value interface{})
	PutLifecycleConfiguration(value interface{})
	PutNetworkConfiguration(value interface{})
	PutProtocolConfiguration(value interface{})
	PutRequestHeaderConfiguration(value interface{})
	PutTimeouts(value *BedrockagentcoreAgentRuntimeTimeouts)
	ResetAgentRuntimeArtifact()
	ResetAuthorizerConfiguration()
	ResetDescription()
	ResetEnvironmentVariables()
	ResetLifecycleConfiguration()
	ResetNetworkConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProtocolConfiguration()
	ResetRegion()
	ResetRequestHeaderConfiguration()
	ResetTags()
	ResetTimeouts()
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

// The jsii proxy struct for BedrockagentcoreAgentRuntime
type jsiiProxy_BedrockagentcoreAgentRuntime struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) AgentRuntimeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) AgentRuntimeArtifact() BedrockagentcoreAgentRuntimeAgentRuntimeArtifactList {
	var returns BedrockagentcoreAgentRuntimeAgentRuntimeArtifactList
	_jsii_.Get(
		j,
		"agentRuntimeArtifact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) AgentRuntimeArtifactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentRuntimeArtifactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) AgentRuntimeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) AgentRuntimeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) AgentRuntimeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) AgentRuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) AuthorizerConfiguration() BedrockagentcoreAgentRuntimeAuthorizerConfigurationList {
	var returns BedrockagentcoreAgentRuntimeAuthorizerConfigurationList
	_jsii_.Get(
		j,
		"authorizerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) AuthorizerConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) LifecycleConfiguration() BedrockagentcoreAgentRuntimeLifecycleConfigurationList {
	var returns BedrockagentcoreAgentRuntimeLifecycleConfigurationList
	_jsii_.Get(
		j,
		"lifecycleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) LifecycleConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) NetworkConfiguration() BedrockagentcoreAgentRuntimeNetworkConfigurationList {
	var returns BedrockagentcoreAgentRuntimeNetworkConfigurationList
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) NetworkConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) ProtocolConfiguration() BedrockagentcoreAgentRuntimeProtocolConfigurationList {
	var returns BedrockagentcoreAgentRuntimeProtocolConfigurationList
	_jsii_.Get(
		j,
		"protocolConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) ProtocolConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protocolConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) RequestHeaderConfiguration() BedrockagentcoreAgentRuntimeRequestHeaderConfigurationList {
	var returns BedrockagentcoreAgentRuntimeRequestHeaderConfigurationList
	_jsii_.Get(
		j,
		"requestHeaderConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) RequestHeaderConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeaderConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) Timeouts() BedrockagentcoreAgentRuntimeTimeoutsOutputReference {
	var returns BedrockagentcoreAgentRuntimeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime) WorkloadIdentityDetails() BedrockagentcoreAgentRuntimeWorkloadIdentityDetailsList {
	var returns BedrockagentcoreAgentRuntimeWorkloadIdentityDetailsList
	_jsii_.Get(
		j,
		"workloadIdentityDetails",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/bedrockagentcore_agent_runtime aws_bedrockagentcore_agent_runtime} Resource.
func NewBedrockagentcoreAgentRuntime(scope constructs.Construct, id *string, config *BedrockagentcoreAgentRuntimeConfig) BedrockagentcoreAgentRuntime {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreAgentRuntimeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreAgentRuntime{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentcoreAgentRuntime.BedrockagentcoreAgentRuntime",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/bedrockagentcore_agent_runtime aws_bedrockagentcore_agent_runtime} Resource.
func NewBedrockagentcoreAgentRuntime_Override(b BedrockagentcoreAgentRuntime, scope constructs.Construct, id *string, config *BedrockagentcoreAgentRuntimeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentcoreAgentRuntime.BedrockagentcoreAgentRuntime",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime)SetAgentRuntimeName(val *string) {
	if err := j.validateSetAgentRuntimeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentRuntimeName",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntime)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a BedrockagentcoreAgentRuntime resource upon running "cdktf plan <stack-name>".
func BedrockagentcoreAgentRuntime_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBedrockagentcoreAgentRuntime_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentcoreAgentRuntime.BedrockagentcoreAgentRuntime",
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
func BedrockagentcoreAgentRuntime_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentcoreAgentRuntime_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentcoreAgentRuntime.BedrockagentcoreAgentRuntime",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockagentcoreAgentRuntime_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentcoreAgentRuntime_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentcoreAgentRuntime.BedrockagentcoreAgentRuntime",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockagentcoreAgentRuntime_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentcoreAgentRuntime_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentcoreAgentRuntime.BedrockagentcoreAgentRuntime",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BedrockagentcoreAgentRuntime_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.bedrockagentcoreAgentRuntime.BedrockagentcoreAgentRuntime",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) PutAgentRuntimeArtifact(value interface{}) {
	if err := b.validatePutAgentRuntimeArtifactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAgentRuntimeArtifact",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) PutAuthorizerConfiguration(value interface{}) {
	if err := b.validatePutAuthorizerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAuthorizerConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) PutLifecycleConfiguration(value interface{}) {
	if err := b.validatePutLifecycleConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLifecycleConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) PutNetworkConfiguration(value interface{}) {
	if err := b.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) PutProtocolConfiguration(value interface{}) {
	if err := b.validatePutProtocolConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putProtocolConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) PutRequestHeaderConfiguration(value interface{}) {
	if err := b.validatePutRequestHeaderConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRequestHeaderConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) PutTimeouts(value *BedrockagentcoreAgentRuntimeTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ResetAgentRuntimeArtifact() {
	_jsii_.InvokeVoid(
		b,
		"resetAgentRuntimeArtifact",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ResetAuthorizerConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetAuthorizerConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		b,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ResetLifecycleConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetLifecycleConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ResetProtocolConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetProtocolConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ResetRegion() {
	_jsii_.InvokeVoid(
		b,
		"resetRegion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ResetRequestHeaderConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetRequestHeaderConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntime) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

