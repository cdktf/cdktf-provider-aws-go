// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryservicetrust

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/directoryservicetrust/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/directory_service_trust aws_directory_service_trust}.
type DirectoryServiceTrust interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConditionalForwarderIpAddrs() *[]*string
	SetConditionalForwarderIpAddrs(val *[]*string)
	ConditionalForwarderIpAddrsInput() *[]*string
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
	CreatedDateTime() *string
	DeleteAssociatedConditionalForwarder() interface{}
	SetDeleteAssociatedConditionalForwarder(val interface{})
	DeleteAssociatedConditionalForwarderInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DirectoryId() *string
	SetDirectoryId(val *string)
	DirectoryIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	LastUpdatedDateTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	RemoteDomainName() *string
	SetRemoteDomainName(val *string)
	RemoteDomainNameInput() *string
	SelectiveAuth() *string
	SetSelectiveAuth(val *string)
	SelectiveAuthInput() *string
	StateLastUpdatedDateTime() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrustDirection() *string
	SetTrustDirection(val *string)
	TrustDirectionInput() *string
	TrustPassword() *string
	SetTrustPassword(val *string)
	TrustPasswordInput() *string
	TrustState() *string
	TrustStateReason() *string
	TrustType() *string
	SetTrustType(val *string)
	TrustTypeInput() *string
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
	ResetConditionalForwarderIpAddrs()
	ResetDeleteAssociatedConditionalForwarder()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSelectiveAuth()
	ResetTrustType()
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

// The jsii proxy struct for DirectoryServiceTrust
type jsiiProxy_DirectoryServiceTrust struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DirectoryServiceTrust) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) ConditionalForwarderIpAddrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"conditionalForwarderIpAddrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) ConditionalForwarderIpAddrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"conditionalForwarderIpAddrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) CreatedDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) DeleteAssociatedConditionalForwarder() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAssociatedConditionalForwarder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) DeleteAssociatedConditionalForwarderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAssociatedConditionalForwarderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) LastUpdatedDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) RemoteDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) RemoteDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) SelectiveAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectiveAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) SelectiveAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectiveAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) StateLastUpdatedDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateLastUpdatedDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) TrustDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) TrustDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) TrustPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) TrustPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) TrustState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) TrustStateReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStateReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) TrustType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceTrust) TrustTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/directory_service_trust aws_directory_service_trust} Resource.
func NewDirectoryServiceTrust(scope constructs.Construct, id *string, config *DirectoryServiceTrustConfig) DirectoryServiceTrust {
	_init_.Initialize()

	if err := validateNewDirectoryServiceTrustParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DirectoryServiceTrust{}

	_jsii_.Create(
		"@cdktf/provider-aws.directoryServiceTrust.DirectoryServiceTrust",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/directory_service_trust aws_directory_service_trust} Resource.
func NewDirectoryServiceTrust_Override(d DirectoryServiceTrust, scope constructs.Construct, id *string, config *DirectoryServiceTrustConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.directoryServiceTrust.DirectoryServiceTrust",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetConditionalForwarderIpAddrs(val *[]*string) {
	if err := j.validateSetConditionalForwarderIpAddrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conditionalForwarderIpAddrs",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetDeleteAssociatedConditionalForwarder(val interface{}) {
	if err := j.validateSetDeleteAssociatedConditionalForwarderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAssociatedConditionalForwarder",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetDirectoryId(val *string) {
	if err := j.validateSetDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetRemoteDomainName(val *string) {
	if err := j.validateSetRemoteDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteDomainName",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetSelectiveAuth(val *string) {
	if err := j.validateSetSelectiveAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectiveAuth",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetTrustDirection(val *string) {
	if err := j.validateSetTrustDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustDirection",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetTrustPassword(val *string) {
	if err := j.validateSetTrustPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustPassword",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceTrust)SetTrustType(val *string) {
	if err := j.validateSetTrustTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustType",
		val,
	)
}

// Generates CDKTF code for importing a DirectoryServiceTrust resource upon running "cdktf plan <stack-name>".
func DirectoryServiceTrust_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDirectoryServiceTrust_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.directoryServiceTrust.DirectoryServiceTrust",
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
func DirectoryServiceTrust_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectoryServiceTrust_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.directoryServiceTrust.DirectoryServiceTrust",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DirectoryServiceTrust_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectoryServiceTrust_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.directoryServiceTrust.DirectoryServiceTrust",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DirectoryServiceTrust_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectoryServiceTrust_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.directoryServiceTrust.DirectoryServiceTrust",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DirectoryServiceTrust_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.directoryServiceTrust.DirectoryServiceTrust",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DirectoryServiceTrust) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DirectoryServiceTrust) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DirectoryServiceTrust) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DirectoryServiceTrust) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DirectoryServiceTrust) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DirectoryServiceTrust) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DirectoryServiceTrust) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DirectoryServiceTrust) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DirectoryServiceTrust) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DirectoryServiceTrust) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DirectoryServiceTrust) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DirectoryServiceTrust) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceTrust) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DirectoryServiceTrust) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DirectoryServiceTrust) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DirectoryServiceTrust) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DirectoryServiceTrust) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DirectoryServiceTrust) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DirectoryServiceTrust) ResetConditionalForwarderIpAddrs() {
	_jsii_.InvokeVoid(
		d,
		"resetConditionalForwarderIpAddrs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryServiceTrust) ResetDeleteAssociatedConditionalForwarder() {
	_jsii_.InvokeVoid(
		d,
		"resetDeleteAssociatedConditionalForwarder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryServiceTrust) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryServiceTrust) ResetSelectiveAuth() {
	_jsii_.InvokeVoid(
		d,
		"resetSelectiveAuth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryServiceTrust) ResetTrustType() {
	_jsii_.InvokeVoid(
		d,
		"resetTrustType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryServiceTrust) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceTrust) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceTrust) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceTrust) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceTrust) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceTrust) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

