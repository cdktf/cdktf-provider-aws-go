// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftdatashareconsumerassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/redshiftdatashareconsumerassociation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/redshift_data_share_consumer_association aws_redshift_data_share_consumer_association}.
type RedshiftDataShareConsumerAssociation interface {
	cdktf.TerraformResource
	AllowWrites() interface{}
	SetAllowWrites(val interface{})
	AllowWritesInput() interface{}
	AssociateEntireAccount() interface{}
	SetAssociateEntireAccount(val interface{})
	AssociateEntireAccountInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsumerArn() *string
	SetConsumerArn(val *string)
	ConsumerArnInput() *string
	ConsumerRegion() *string
	SetConsumerRegion(val *string)
	ConsumerRegionInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DataShareArn() *string
	SetDataShareArn(val *string)
	DataShareArnInput() *string
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
	ManagedBy() *string
	// The tree node.
	Node() constructs.Node
	ProducerArn() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAllowWrites()
	ResetAssociateEntireAccount()
	ResetConsumerArn()
	ResetConsumerRegion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
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

// The jsii proxy struct for RedshiftDataShareConsumerAssociation
type jsiiProxy_RedshiftDataShareConsumerAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) AllowWrites() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowWrites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) AllowWritesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowWritesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) AssociateEntireAccount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateEntireAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) AssociateEntireAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateEntireAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) ConsumerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) ConsumerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) ConsumerRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) ConsumerRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) DataShareArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataShareArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) DataShareArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataShareArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) ManagedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) ProducerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"producerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/redshift_data_share_consumer_association aws_redshift_data_share_consumer_association} Resource.
func NewRedshiftDataShareConsumerAssociation(scope constructs.Construct, id *string, config *RedshiftDataShareConsumerAssociationConfig) RedshiftDataShareConsumerAssociation {
	_init_.Initialize()

	if err := validateNewRedshiftDataShareConsumerAssociationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftDataShareConsumerAssociation{}

	_jsii_.Create(
		"@cdktf/provider-aws.redshiftDataShareConsumerAssociation.RedshiftDataShareConsumerAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/redshift_data_share_consumer_association aws_redshift_data_share_consumer_association} Resource.
func NewRedshiftDataShareConsumerAssociation_Override(r RedshiftDataShareConsumerAssociation, scope constructs.Construct, id *string, config *RedshiftDataShareConsumerAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.redshiftDataShareConsumerAssociation.RedshiftDataShareConsumerAssociation",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation)SetAllowWrites(val interface{}) {
	if err := j.validateSetAllowWritesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowWrites",
		val,
	)
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation)SetAssociateEntireAccount(val interface{}) {
	if err := j.validateSetAssociateEntireAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateEntireAccount",
		val,
	)
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation)SetConsumerArn(val *string) {
	if err := j.validateSetConsumerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerArn",
		val,
	)
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation)SetConsumerRegion(val *string) {
	if err := j.validateSetConsumerRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerRegion",
		val,
	)
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation)SetDataShareArn(val *string) {
	if err := j.validateSetDataShareArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataShareArn",
		val,
	)
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RedshiftDataShareConsumerAssociation)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTF code for importing a RedshiftDataShareConsumerAssociation resource upon running "cdktf plan <stack-name>".
func RedshiftDataShareConsumerAssociation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRedshiftDataShareConsumerAssociation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.redshiftDataShareConsumerAssociation.RedshiftDataShareConsumerAssociation",
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
func RedshiftDataShareConsumerAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftDataShareConsumerAssociation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.redshiftDataShareConsumerAssociation.RedshiftDataShareConsumerAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedshiftDataShareConsumerAssociation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftDataShareConsumerAssociation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.redshiftDataShareConsumerAssociation.RedshiftDataShareConsumerAssociation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedshiftDataShareConsumerAssociation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftDataShareConsumerAssociation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.redshiftDataShareConsumerAssociation.RedshiftDataShareConsumerAssociation",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RedshiftDataShareConsumerAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.redshiftDataShareConsumerAssociation.RedshiftDataShareConsumerAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) ResetAllowWrites() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowWrites",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) ResetAssociateEntireAccount() {
	_jsii_.InvokeVoid(
		r,
		"resetAssociateEntireAccount",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) ResetConsumerArn() {
	_jsii_.InvokeVoid(
		r,
		"resetConsumerArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) ResetConsumerRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetConsumerRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) ResetRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftDataShareConsumerAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

