// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configremediationconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/configremediationconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/config_remediation_configuration aws_config_remediation_configuration}.
type ConfigRemediationConfiguration interface {
	cdktf.TerraformResource
	Arn() *string
	Automatic() interface{}
	SetAutomatic(val interface{})
	AutomaticInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigRuleName() *string
	SetConfigRuleName(val *string)
	ConfigRuleNameInput() *string
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
	ExecutionControls() ConfigRemediationConfigurationExecutionControlsOutputReference
	ExecutionControlsInput() *ConfigRemediationConfigurationExecutionControls
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumAutomaticAttempts() *float64
	SetMaximumAutomaticAttempts(val *float64)
	MaximumAutomaticAttemptsInput() *float64
	// The tree node.
	Node() constructs.Node
	Parameter() ConfigRemediationConfigurationParameterList
	ParameterInput() interface{}
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
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	RetryAttemptSeconds() *float64
	SetRetryAttemptSeconds(val *float64)
	RetryAttemptSecondsInput() *float64
	TargetId() *string
	SetTargetId(val *string)
	TargetIdInput() *string
	TargetType() *string
	SetTargetType(val *string)
	TargetTypeInput() *string
	TargetVersion() *string
	SetTargetVersion(val *string)
	TargetVersionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutExecutionControls(value *ConfigRemediationConfigurationExecutionControls)
	PutParameter(value interface{})
	ResetAutomatic()
	ResetExecutionControls()
	ResetId()
	ResetMaximumAutomaticAttempts()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameter()
	ResetResourceType()
	ResetRetryAttemptSeconds()
	ResetTargetVersion()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ConfigRemediationConfiguration
type jsiiProxy_ConfigRemediationConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Automatic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automatic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) AutomaticInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ConfigRuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ExecutionControls() ConfigRemediationConfigurationExecutionControlsOutputReference {
	var returns ConfigRemediationConfigurationExecutionControlsOutputReference
	_jsii_.Get(
		j,
		"executionControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ExecutionControlsInput() *ConfigRemediationConfigurationExecutionControls {
	var returns *ConfigRemediationConfigurationExecutionControls
	_jsii_.Get(
		j,
		"executionControlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) MaximumAutomaticAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumAutomaticAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) MaximumAutomaticAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumAutomaticAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Parameter() ConfigRemediationConfigurationParameterList {
	var returns ConfigRemediationConfigurationParameterList
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) RetryAttemptSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryAttemptSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) RetryAttemptSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryAttemptSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TargetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TargetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TargetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TargetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/config_remediation_configuration aws_config_remediation_configuration} Resource.
func NewConfigRemediationConfiguration(scope constructs.Construct, id *string, config *ConfigRemediationConfigurationConfig) ConfigRemediationConfiguration {
	_init_.Initialize()

	if err := validateNewConfigRemediationConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigRemediationConfiguration{}

	_jsii_.Create(
		"@cdktf/provider-aws.configRemediationConfiguration.ConfigRemediationConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/config_remediation_configuration aws_config_remediation_configuration} Resource.
func NewConfigRemediationConfiguration_Override(c ConfigRemediationConfiguration, scope constructs.Construct, id *string, config *ConfigRemediationConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.configRemediationConfiguration.ConfigRemediationConfiguration",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetAutomatic(val interface{}) {
	if err := j.validateSetAutomaticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automatic",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetConfigRuleName(val *string) {
	if err := j.validateSetConfigRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configRuleName",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetMaximumAutomaticAttempts(val *float64) {
	if err := j.validateSetMaximumAutomaticAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumAutomaticAttempts",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetRetryAttemptSeconds(val *float64) {
	if err := j.validateSetRetryAttemptSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryAttemptSeconds",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetTargetId(val *string) {
	if err := j.validateSetTargetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetId",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetTargetType(val *string) {
	if err := j.validateSetTargetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration)SetTargetVersion(val *string) {
	if err := j.validateSetTargetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetVersion",
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
func ConfigRemediationConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConfigRemediationConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.configRemediationConfiguration.ConfigRemediationConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConfigRemediationConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConfigRemediationConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.configRemediationConfiguration.ConfigRemediationConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConfigRemediationConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConfigRemediationConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.configRemediationConfiguration.ConfigRemediationConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigRemediationConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.configRemediationConfiguration.ConfigRemediationConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) PutExecutionControls(value *ConfigRemediationConfigurationExecutionControls) {
	if err := c.validatePutExecutionControlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExecutionControls",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) PutParameter(value interface{}) {
	if err := c.validatePutParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putParameter",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ResetAutomatic() {
	_jsii_.InvokeVoid(
		c,
		"resetAutomatic",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ResetExecutionControls() {
	_jsii_.InvokeVoid(
		c,
		"resetExecutionControls",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ResetMaximumAutomaticAttempts() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximumAutomaticAttempts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ResetParameter() {
	_jsii_.InvokeVoid(
		c,
		"resetParameter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ResetResourceType() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ResetRetryAttemptSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetRetryAttemptSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ResetTargetVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

