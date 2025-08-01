// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configorganizationcustompolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/configorganizationcustompolicyrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/config_organization_custom_policy_rule aws_config_organization_custom_policy_rule}.
type ConfigOrganizationCustomPolicyRule interface {
	cdktf.TerraformResource
	Arn() *string
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
	DebugLogDeliveryAccounts() *[]*string
	SetDebugLogDeliveryAccounts(val *[]*string)
	DebugLogDeliveryAccountsInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExcludedAccounts() *[]*string
	SetExcludedAccounts(val *[]*string)
	ExcludedAccountsInput() *[]*string
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
	InputParameters() *string
	SetInputParameters(val *string)
	InputParametersInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumExecutionFrequency() *string
	SetMaximumExecutionFrequency(val *string)
	MaximumExecutionFrequencyInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PolicyRuntime() *string
	SetPolicyRuntime(val *string)
	PolicyRuntimeInput() *string
	PolicyText() *string
	SetPolicyText(val *string)
	PolicyTextInput() *string
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
	ResourceIdScope() *string
	SetResourceIdScope(val *string)
	ResourceIdScopeInput() *string
	ResourceTypesScope() *[]*string
	SetResourceTypesScope(val *[]*string)
	ResourceTypesScopeInput() *[]*string
	TagKeyScope() *string
	SetTagKeyScope(val *string)
	TagKeyScopeInput() *string
	TagValueScope() *string
	SetTagValueScope(val *string)
	TagValueScopeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ConfigOrganizationCustomPolicyRuleTimeoutsOutputReference
	TimeoutsInput() interface{}
	TriggerTypes() *[]*string
	SetTriggerTypes(val *[]*string)
	TriggerTypesInput() *[]*string
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
	PutTimeouts(value *ConfigOrganizationCustomPolicyRuleTimeouts)
	ResetDebugLogDeliveryAccounts()
	ResetDescription()
	ResetExcludedAccounts()
	ResetId()
	ResetInputParameters()
	ResetMaximumExecutionFrequency()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetResourceIdScope()
	ResetResourceTypesScope()
	ResetTagKeyScope()
	ResetTagValueScope()
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

// The jsii proxy struct for ConfigOrganizationCustomPolicyRule
type jsiiProxy_ConfigOrganizationCustomPolicyRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) DebugLogDeliveryAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"debugLogDeliveryAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) DebugLogDeliveryAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"debugLogDeliveryAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) ExcludedAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) ExcludedAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) InputParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) InputParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) MaximumExecutionFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumExecutionFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) MaximumExecutionFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumExecutionFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) PolicyRuntime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) PolicyRuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) PolicyText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) PolicyTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResourceIdScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResourceIdScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResourceTypesScope() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResourceTypesScopeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) TagKeyScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) TagKeyScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) TagValueScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) TagValueScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) Timeouts() ConfigOrganizationCustomPolicyRuleTimeoutsOutputReference {
	var returns ConfigOrganizationCustomPolicyRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) TriggerTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"triggerTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule) TriggerTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"triggerTypesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/config_organization_custom_policy_rule aws_config_organization_custom_policy_rule} Resource.
func NewConfigOrganizationCustomPolicyRule(scope constructs.Construct, id *string, config *ConfigOrganizationCustomPolicyRuleConfig) ConfigOrganizationCustomPolicyRule {
	_init_.Initialize()

	if err := validateNewConfigOrganizationCustomPolicyRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigOrganizationCustomPolicyRule{}

	_jsii_.Create(
		"@cdktf/provider-aws.configOrganizationCustomPolicyRule.ConfigOrganizationCustomPolicyRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/config_organization_custom_policy_rule aws_config_organization_custom_policy_rule} Resource.
func NewConfigOrganizationCustomPolicyRule_Override(c ConfigOrganizationCustomPolicyRule, scope constructs.Construct, id *string, config *ConfigOrganizationCustomPolicyRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.configOrganizationCustomPolicyRule.ConfigOrganizationCustomPolicyRule",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetDebugLogDeliveryAccounts(val *[]*string) {
	if err := j.validateSetDebugLogDeliveryAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"debugLogDeliveryAccounts",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetExcludedAccounts(val *[]*string) {
	if err := j.validateSetExcludedAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedAccounts",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetInputParameters(val *string) {
	if err := j.validateSetInputParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputParameters",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetMaximumExecutionFrequency(val *string) {
	if err := j.validateSetMaximumExecutionFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumExecutionFrequency",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetPolicyRuntime(val *string) {
	if err := j.validateSetPolicyRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyRuntime",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetPolicyText(val *string) {
	if err := j.validateSetPolicyTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyText",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetResourceIdScope(val *string) {
	if err := j.validateSetResourceIdScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceIdScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetResourceTypesScope(val *[]*string) {
	if err := j.validateSetResourceTypesScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypesScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetTagKeyScope(val *string) {
	if err := j.validateSetTagKeyScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagKeyScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetTagValueScope(val *string) {
	if err := j.validateSetTagValueScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagValueScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomPolicyRule)SetTriggerTypes(val *[]*string) {
	if err := j.validateSetTriggerTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerTypes",
		val,
	)
}

// Generates CDKTF code for importing a ConfigOrganizationCustomPolicyRule resource upon running "cdktf plan <stack-name>".
func ConfigOrganizationCustomPolicyRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateConfigOrganizationCustomPolicyRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.configOrganizationCustomPolicyRule.ConfigOrganizationCustomPolicyRule",
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
func ConfigOrganizationCustomPolicyRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConfigOrganizationCustomPolicyRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.configOrganizationCustomPolicyRule.ConfigOrganizationCustomPolicyRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConfigOrganizationCustomPolicyRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConfigOrganizationCustomPolicyRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.configOrganizationCustomPolicyRule.ConfigOrganizationCustomPolicyRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConfigOrganizationCustomPolicyRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConfigOrganizationCustomPolicyRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.configOrganizationCustomPolicyRule.ConfigOrganizationCustomPolicyRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigOrganizationCustomPolicyRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.configOrganizationCustomPolicyRule.ConfigOrganizationCustomPolicyRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) PutTimeouts(value *ConfigOrganizationCustomPolicyRuleTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResetDebugLogDeliveryAccounts() {
	_jsii_.InvokeVoid(
		c,
		"resetDebugLogDeliveryAccounts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResetExcludedAccounts() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludedAccounts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResetInputParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetInputParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResetMaximumExecutionFrequency() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximumExecutionFrequency",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResetResourceIdScope() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceIdScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResetResourceTypesScope() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceTypesScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResetTagKeyScope() {
	_jsii_.InvokeVoid(
		c,
		"resetTagKeyScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResetTagValueScope() {
	_jsii_.InvokeVoid(
		c,
		"resetTagValueScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationCustomPolicyRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

