// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeoptimizerrecommendationpreferences

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/computeoptimizerrecommendationpreferences/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/computeoptimizer_recommendation_preferences aws_computeoptimizer_recommendation_preferences}.
type ComputeoptimizerRecommendationPreferences interface {
	cdktf.TerraformResource
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
	EnhancedInfrastructureMetrics() *string
	SetEnhancedInfrastructureMetrics(val *string)
	EnhancedInfrastructureMetricsInput() *string
	ExternalMetricsPreference() ComputeoptimizerRecommendationPreferencesExternalMetricsPreferenceList
	ExternalMetricsPreferenceInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InferredWorkloadTypes() *string
	SetInferredWorkloadTypes(val *string)
	InferredWorkloadTypesInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LookBackPeriod() *string
	SetLookBackPeriod(val *string)
	LookBackPeriodInput() *string
	// The tree node.
	Node() constructs.Node
	PreferredResource() ComputeoptimizerRecommendationPreferencesPreferredResourceList
	PreferredResourceInput() interface{}
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
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	SavingsEstimationMode() *string
	SetSavingsEstimationMode(val *string)
	SavingsEstimationModeInput() *string
	Scope() ComputeoptimizerRecommendationPreferencesScopeList
	ScopeInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UtilizationPreference() ComputeoptimizerRecommendationPreferencesUtilizationPreferenceList
	UtilizationPreferenceInput() interface{}
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
	PutExternalMetricsPreference(value interface{})
	PutPreferredResource(value interface{})
	PutScope(value interface{})
	PutUtilizationPreference(value interface{})
	ResetEnhancedInfrastructureMetrics()
	ResetExternalMetricsPreference()
	ResetInferredWorkloadTypes()
	ResetLookBackPeriod()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreferredResource()
	ResetRegion()
	ResetSavingsEstimationMode()
	ResetScope()
	ResetUtilizationPreference()
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

// The jsii proxy struct for ComputeoptimizerRecommendationPreferences
type jsiiProxy_ComputeoptimizerRecommendationPreferences struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) EnhancedInfrastructureMetrics() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enhancedInfrastructureMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) EnhancedInfrastructureMetricsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enhancedInfrastructureMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) ExternalMetricsPreference() ComputeoptimizerRecommendationPreferencesExternalMetricsPreferenceList {
	var returns ComputeoptimizerRecommendationPreferencesExternalMetricsPreferenceList
	_jsii_.Get(
		j,
		"externalMetricsPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) ExternalMetricsPreferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalMetricsPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) InferredWorkloadTypes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferredWorkloadTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) InferredWorkloadTypesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferredWorkloadTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) LookBackPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookBackPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) LookBackPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookBackPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) PreferredResource() ComputeoptimizerRecommendationPreferencesPreferredResourceList {
	var returns ComputeoptimizerRecommendationPreferencesPreferredResourceList
	_jsii_.Get(
		j,
		"preferredResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) PreferredResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preferredResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) SavingsEstimationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsEstimationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) SavingsEstimationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsEstimationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) Scope() ComputeoptimizerRecommendationPreferencesScopeList {
	var returns ComputeoptimizerRecommendationPreferencesScopeList
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) ScopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) UtilizationPreference() ComputeoptimizerRecommendationPreferencesUtilizationPreferenceList {
	var returns ComputeoptimizerRecommendationPreferencesUtilizationPreferenceList
	_jsii_.Get(
		j,
		"utilizationPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences) UtilizationPreferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizationPreferenceInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/computeoptimizer_recommendation_preferences aws_computeoptimizer_recommendation_preferences} Resource.
func NewComputeoptimizerRecommendationPreferences(scope constructs.Construct, id *string, config *ComputeoptimizerRecommendationPreferencesConfig) ComputeoptimizerRecommendationPreferences {
	_init_.Initialize()

	if err := validateNewComputeoptimizerRecommendationPreferencesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeoptimizerRecommendationPreferences{}

	_jsii_.Create(
		"@cdktf/provider-aws.computeoptimizerRecommendationPreferences.ComputeoptimizerRecommendationPreferences",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/computeoptimizer_recommendation_preferences aws_computeoptimizer_recommendation_preferences} Resource.
func NewComputeoptimizerRecommendationPreferences_Override(c ComputeoptimizerRecommendationPreferences, scope constructs.Construct, id *string, config *ComputeoptimizerRecommendationPreferencesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.computeoptimizerRecommendationPreferences.ComputeoptimizerRecommendationPreferences",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences)SetEnhancedInfrastructureMetrics(val *string) {
	if err := j.validateSetEnhancedInfrastructureMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedInfrastructureMetrics",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences)SetInferredWorkloadTypes(val *string) {
	if err := j.validateSetInferredWorkloadTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferredWorkloadTypes",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences)SetLookBackPeriod(val *string) {
	if err := j.validateSetLookBackPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookBackPeriod",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerRecommendationPreferences)SetSavingsEstimationMode(val *string) {
	if err := j.validateSetSavingsEstimationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"savingsEstimationMode",
		val,
	)
}

// Generates CDKTF code for importing a ComputeoptimizerRecommendationPreferences resource upon running "cdktf plan <stack-name>".
func ComputeoptimizerRecommendationPreferences_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeoptimizerRecommendationPreferences_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.computeoptimizerRecommendationPreferences.ComputeoptimizerRecommendationPreferences",
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
func ComputeoptimizerRecommendationPreferences_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeoptimizerRecommendationPreferences_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.computeoptimizerRecommendationPreferences.ComputeoptimizerRecommendationPreferences",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeoptimizerRecommendationPreferences_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeoptimizerRecommendationPreferences_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.computeoptimizerRecommendationPreferences.ComputeoptimizerRecommendationPreferences",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeoptimizerRecommendationPreferences_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeoptimizerRecommendationPreferences_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.computeoptimizerRecommendationPreferences.ComputeoptimizerRecommendationPreferences",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeoptimizerRecommendationPreferences_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.computeoptimizerRecommendationPreferences.ComputeoptimizerRecommendationPreferences",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) PutExternalMetricsPreference(value interface{}) {
	if err := c.validatePutExternalMetricsPreferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExternalMetricsPreference",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) PutPreferredResource(value interface{}) {
	if err := c.validatePutPreferredResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPreferredResource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) PutScope(value interface{}) {
	if err := c.validatePutScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putScope",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) PutUtilizationPreference(value interface{}) {
	if err := c.validatePutUtilizationPreferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUtilizationPreference",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ResetEnhancedInfrastructureMetrics() {
	_jsii_.InvokeVoid(
		c,
		"resetEnhancedInfrastructureMetrics",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ResetExternalMetricsPreference() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalMetricsPreference",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ResetInferredWorkloadTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetInferredWorkloadTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ResetLookBackPeriod() {
	_jsii_.InvokeVoid(
		c,
		"resetLookBackPeriod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ResetPreferredResource() {
	_jsii_.InvokeVoid(
		c,
		"resetPreferredResource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ResetSavingsEstimationMode() {
	_jsii_.InvokeVoid(
		c,
		"resetSavingsEstimationMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ResetScope() {
	_jsii_.InvokeVoid(
		c,
		"resetScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ResetUtilizationPreference() {
	_jsii_.InvokeVoid(
		c,
		"resetUtilizationPreference",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeoptimizerRecommendationPreferences) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

