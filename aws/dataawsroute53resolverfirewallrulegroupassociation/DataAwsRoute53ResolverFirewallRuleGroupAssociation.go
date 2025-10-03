// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsroute53resolverfirewallrulegroupassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dataawsroute53resolverfirewallrulegroupassociation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/data-sources/route53_resolver_firewall_rule_group_association aws_route53_resolver_firewall_rule_group_association}.
type DataAwsRoute53ResolverFirewallRuleGroupAssociation interface {
	cdktf.TerraformDataSource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *string
	CreatorRequestId() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FirewallRuleGroupAssociationId() *string
	SetFirewallRuleGroupAssociationId(val *string)
	FirewallRuleGroupAssociationIdInput() *string
	FirewallRuleGroupId() *string
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
	ManagedOwnerName() *string
	ModificationTime() *string
	MutationProtection() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Priority() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Status() *string
	StatusMessage() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VpcId() *string
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
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
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

// The jsii proxy struct for DataAwsRoute53ResolverFirewallRuleGroupAssociation
type jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) CreatorRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creatorRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) FirewallRuleGroupAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) FirewallRuleGroupAssociationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) FirewallRuleGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) ManagedOwnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) ModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) MutationProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mutationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) StatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/data-sources/route53_resolver_firewall_rule_group_association aws_route53_resolver_firewall_rule_group_association} Data Source.
func NewDataAwsRoute53ResolverFirewallRuleGroupAssociation(scope constructs.Construct, id *string, config *DataAwsRoute53ResolverFirewallRuleGroupAssociationConfig) DataAwsRoute53ResolverFirewallRuleGroupAssociation {
	_init_.Initialize()

	if err := validateNewDataAwsRoute53ResolverFirewallRuleGroupAssociationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsRoute53ResolverFirewallRuleGroupAssociation.DataAwsRoute53ResolverFirewallRuleGroupAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/data-sources/route53_resolver_firewall_rule_group_association aws_route53_resolver_firewall_rule_group_association} Data Source.
func NewDataAwsRoute53ResolverFirewallRuleGroupAssociation_Override(d DataAwsRoute53ResolverFirewallRuleGroupAssociation, scope constructs.Construct, id *string, config *DataAwsRoute53ResolverFirewallRuleGroupAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsRoute53ResolverFirewallRuleGroupAssociation.DataAwsRoute53ResolverFirewallRuleGroupAssociation",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation)SetFirewallRuleGroupAssociationId(val *string) {
	if err := j.validateSetFirewallRuleGroupAssociationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallRuleGroupAssociationId",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsRoute53ResolverFirewallRuleGroupAssociation resource upon running "cdktf plan <stack-name>".
func DataAwsRoute53ResolverFirewallRuleGroupAssociation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsRoute53ResolverFirewallRuleGroupAssociation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRoute53ResolverFirewallRuleGroupAssociation.DataAwsRoute53ResolverFirewallRuleGroupAssociation",
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
func DataAwsRoute53ResolverFirewallRuleGroupAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsRoute53ResolverFirewallRuleGroupAssociation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRoute53ResolverFirewallRuleGroupAssociation.DataAwsRoute53ResolverFirewallRuleGroupAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsRoute53ResolverFirewallRuleGroupAssociation_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsRoute53ResolverFirewallRuleGroupAssociation_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRoute53ResolverFirewallRuleGroupAssociation.DataAwsRoute53ResolverFirewallRuleGroupAssociation",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsRoute53ResolverFirewallRuleGroupAssociation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsRoute53ResolverFirewallRuleGroupAssociation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRoute53ResolverFirewallRuleGroupAssociation.DataAwsRoute53ResolverFirewallRuleGroupAssociation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsRoute53ResolverFirewallRuleGroupAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsRoute53ResolverFirewallRuleGroupAssociation.DataAwsRoute53ResolverFirewallRuleGroupAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53ResolverFirewallRuleGroupAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

