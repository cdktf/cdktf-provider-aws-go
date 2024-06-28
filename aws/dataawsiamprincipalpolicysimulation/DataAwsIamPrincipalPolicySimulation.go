// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsiamprincipalpolicysimulation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dataawsiamprincipalpolicysimulation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/data-sources/iam_principal_policy_simulation aws_iam_principal_policy_simulation}.
type DataAwsIamPrincipalPolicySimulation interface {
	cdktf.TerraformDataSource
	ActionNames() *[]*string
	SetActionNames(val *[]*string)
	ActionNamesInput() *[]*string
	AdditionalPoliciesJson() *[]*string
	SetAdditionalPoliciesJson(val *[]*string)
	AdditionalPoliciesJsonInput() *[]*string
	AllAllowed() cdktf.IResolvable
	CallerArn() *string
	SetCallerArn(val *string)
	CallerArnInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Context() DataAwsIamPrincipalPolicySimulationContextList
	ContextInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PermissionsBoundaryPoliciesJson() *[]*string
	SetPermissionsBoundaryPoliciesJson(val *[]*string)
	PermissionsBoundaryPoliciesJsonInput() *[]*string
	PolicySourceArn() *string
	SetPolicySourceArn(val *string)
	PolicySourceArnInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ResourceArns() *[]*string
	SetResourceArns(val *[]*string)
	ResourceArnsInput() *[]*string
	ResourceHandlingOption() *string
	SetResourceHandlingOption(val *string)
	ResourceHandlingOptionInput() *string
	ResourceOwnerAccountId() *string
	SetResourceOwnerAccountId(val *string)
	ResourceOwnerAccountIdInput() *string
	ResourcePolicyJson() *string
	SetResourcePolicyJson(val *string)
	ResourcePolicyJsonInput() *string
	Results() DataAwsIamPrincipalPolicySimulationResultsList
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
	PutContext(value interface{})
	ResetAdditionalPoliciesJson()
	ResetCallerArn()
	ResetContext()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPermissionsBoundaryPoliciesJson()
	ResetResourceArns()
	ResetResourceHandlingOption()
	ResetResourceOwnerAccountId()
	ResetResourcePolicyJson()
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

// The jsii proxy struct for DataAwsIamPrincipalPolicySimulation
type jsiiProxy_DataAwsIamPrincipalPolicySimulation struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ActionNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ActionNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) AdditionalPoliciesJson() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalPoliciesJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) AdditionalPoliciesJsonInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalPoliciesJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) AllAllowed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) CallerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) CallerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) Context() DataAwsIamPrincipalPolicySimulationContextList {
	var returns DataAwsIamPrincipalPolicySimulationContextList
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) PermissionsBoundaryPoliciesJson() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsBoundaryPoliciesJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) PermissionsBoundaryPoliciesJsonInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsBoundaryPoliciesJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) PolicySourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policySourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) PolicySourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policySourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResourceArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResourceArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResourceHandlingOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceHandlingOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResourceHandlingOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceHandlingOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResourceOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceOwnerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResourceOwnerAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceOwnerAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResourcePolicyJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePolicyJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResourcePolicyJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePolicyJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) Results() DataAwsIamPrincipalPolicySimulationResultsList {
	var returns DataAwsIamPrincipalPolicySimulationResultsList
	_jsii_.Get(
		j,
		"results",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/data-sources/iam_principal_policy_simulation aws_iam_principal_policy_simulation} Data Source.
func NewDataAwsIamPrincipalPolicySimulation(scope constructs.Construct, id *string, config *DataAwsIamPrincipalPolicySimulationConfig) DataAwsIamPrincipalPolicySimulation {
	_init_.Initialize()

	if err := validateNewDataAwsIamPrincipalPolicySimulationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsIamPrincipalPolicySimulation{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsIamPrincipalPolicySimulation.DataAwsIamPrincipalPolicySimulation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/data-sources/iam_principal_policy_simulation aws_iam_principal_policy_simulation} Data Source.
func NewDataAwsIamPrincipalPolicySimulation_Override(d DataAwsIamPrincipalPolicySimulation, scope constructs.Construct, id *string, config *DataAwsIamPrincipalPolicySimulationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsIamPrincipalPolicySimulation.DataAwsIamPrincipalPolicySimulation",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetActionNames(val *[]*string) {
	if err := j.validateSetActionNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionNames",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetAdditionalPoliciesJson(val *[]*string) {
	if err := j.validateSetAdditionalPoliciesJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalPoliciesJson",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetCallerArn(val *string) {
	if err := j.validateSetCallerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callerArn",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetPermissionsBoundaryPoliciesJson(val *[]*string) {
	if err := j.validateSetPermissionsBoundaryPoliciesJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissionsBoundaryPoliciesJson",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetPolicySourceArn(val *string) {
	if err := j.validateSetPolicySourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policySourceArn",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetResourceArns(val *[]*string) {
	if err := j.validateSetResourceArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceArns",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetResourceHandlingOption(val *string) {
	if err := j.validateSetResourceHandlingOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceHandlingOption",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetResourceOwnerAccountId(val *string) {
	if err := j.validateSetResourceOwnerAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceOwnerAccountId",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPrincipalPolicySimulation)SetResourcePolicyJson(val *string) {
	if err := j.validateSetResourcePolicyJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePolicyJson",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsIamPrincipalPolicySimulation resource upon running "cdktf plan <stack-name>".
func DataAwsIamPrincipalPolicySimulation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsIamPrincipalPolicySimulation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsIamPrincipalPolicySimulation.DataAwsIamPrincipalPolicySimulation",
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
func DataAwsIamPrincipalPolicySimulation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsIamPrincipalPolicySimulation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsIamPrincipalPolicySimulation.DataAwsIamPrincipalPolicySimulation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsIamPrincipalPolicySimulation_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsIamPrincipalPolicySimulation_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsIamPrincipalPolicySimulation.DataAwsIamPrincipalPolicySimulation",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsIamPrincipalPolicySimulation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsIamPrincipalPolicySimulation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsIamPrincipalPolicySimulation.DataAwsIamPrincipalPolicySimulation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsIamPrincipalPolicySimulation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsIamPrincipalPolicySimulation.DataAwsIamPrincipalPolicySimulation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) PutContext(value interface{}) {
	if err := d.validatePutContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContext",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResetAdditionalPoliciesJson() {
	_jsii_.InvokeVoid(
		d,
		"resetAdditionalPoliciesJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResetCallerArn() {
	_jsii_.InvokeVoid(
		d,
		"resetCallerArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResetContext() {
	_jsii_.InvokeVoid(
		d,
		"resetContext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResetPermissionsBoundaryPoliciesJson() {
	_jsii_.InvokeVoid(
		d,
		"resetPermissionsBoundaryPoliciesJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResetResourceArns() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceArns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResetResourceHandlingOption() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceHandlingOption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResetResourceOwnerAccountId() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceOwnerAccountId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ResetResourcePolicyJson() {
	_jsii_.InvokeVoid(
		d,
		"resetResourcePolicyJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPrincipalPolicySimulation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

