// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccesstrustprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/verifiedaccesstrustprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/verifiedaccess_trust_provider aws_verifiedaccess_trust_provider}.
type VerifiedaccessTrustProvider interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DeviceOptions() VerifiedaccessTrustProviderDeviceOptionsOutputReference
	DeviceOptionsInput() *VerifiedaccessTrustProviderDeviceOptions
	DeviceTrustProviderType() *string
	SetDeviceTrustProviderType(val *string)
	DeviceTrustProviderTypeInput() *string
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
	NativeApplicationOidcOptions() VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference
	NativeApplicationOidcOptionsInput() *VerifiedaccessTrustProviderNativeApplicationOidcOptions
	// The tree node.
	Node() constructs.Node
	OidcOptions() VerifiedaccessTrustProviderOidcOptionsOutputReference
	OidcOptionsInput() *VerifiedaccessTrustProviderOidcOptions
	PolicyReferenceName() *string
	SetPolicyReferenceName(val *string)
	PolicyReferenceNameInput() *string
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
	SseSpecification() VerifiedaccessTrustProviderSseSpecificationOutputReference
	SseSpecificationInput() *VerifiedaccessTrustProviderSseSpecification
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
	Timeouts() VerifiedaccessTrustProviderTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrustProviderType() *string
	SetTrustProviderType(val *string)
	TrustProviderTypeInput() *string
	UserTrustProviderType() *string
	SetUserTrustProviderType(val *string)
	UserTrustProviderTypeInput() *string
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
	PutDeviceOptions(value *VerifiedaccessTrustProviderDeviceOptions)
	PutNativeApplicationOidcOptions(value *VerifiedaccessTrustProviderNativeApplicationOidcOptions)
	PutOidcOptions(value *VerifiedaccessTrustProviderOidcOptions)
	PutSseSpecification(value *VerifiedaccessTrustProviderSseSpecification)
	PutTimeouts(value *VerifiedaccessTrustProviderTimeouts)
	ResetDescription()
	ResetDeviceOptions()
	ResetDeviceTrustProviderType()
	ResetId()
	ResetNativeApplicationOidcOptions()
	ResetOidcOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSseSpecification()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetUserTrustProviderType()
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

// The jsii proxy struct for VerifiedaccessTrustProvider
type jsiiProxy_VerifiedaccessTrustProvider struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) DeviceOptions() VerifiedaccessTrustProviderDeviceOptionsOutputReference {
	var returns VerifiedaccessTrustProviderDeviceOptionsOutputReference
	_jsii_.Get(
		j,
		"deviceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) DeviceOptionsInput() *VerifiedaccessTrustProviderDeviceOptions {
	var returns *VerifiedaccessTrustProviderDeviceOptions
	_jsii_.Get(
		j,
		"deviceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) DeviceTrustProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTrustProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) DeviceTrustProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTrustProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) NativeApplicationOidcOptions() VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference {
	var returns VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference
	_jsii_.Get(
		j,
		"nativeApplicationOidcOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) NativeApplicationOidcOptionsInput() *VerifiedaccessTrustProviderNativeApplicationOidcOptions {
	var returns *VerifiedaccessTrustProviderNativeApplicationOidcOptions
	_jsii_.Get(
		j,
		"nativeApplicationOidcOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) OidcOptions() VerifiedaccessTrustProviderOidcOptionsOutputReference {
	var returns VerifiedaccessTrustProviderOidcOptionsOutputReference
	_jsii_.Get(
		j,
		"oidcOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) OidcOptionsInput() *VerifiedaccessTrustProviderOidcOptions {
	var returns *VerifiedaccessTrustProviderOidcOptions
	_jsii_.Get(
		j,
		"oidcOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) PolicyReferenceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyReferenceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) PolicyReferenceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyReferenceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) SseSpecification() VerifiedaccessTrustProviderSseSpecificationOutputReference {
	var returns VerifiedaccessTrustProviderSseSpecificationOutputReference
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) SseSpecificationInput() *VerifiedaccessTrustProviderSseSpecification {
	var returns *VerifiedaccessTrustProviderSseSpecification
	_jsii_.Get(
		j,
		"sseSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) Timeouts() VerifiedaccessTrustProviderTimeoutsOutputReference {
	var returns VerifiedaccessTrustProviderTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) TrustProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) TrustProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) UserTrustProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTrustProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProvider) UserTrustProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTrustProviderTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/verifiedaccess_trust_provider aws_verifiedaccess_trust_provider} Resource.
func NewVerifiedaccessTrustProvider(scope constructs.Construct, id *string, config *VerifiedaccessTrustProviderConfig) VerifiedaccessTrustProvider {
	_init_.Initialize()

	if err := validateNewVerifiedaccessTrustProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VerifiedaccessTrustProvider{}

	_jsii_.Create(
		"@cdktf/provider-aws.verifiedaccessTrustProvider.VerifiedaccessTrustProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/verifiedaccess_trust_provider aws_verifiedaccess_trust_provider} Resource.
func NewVerifiedaccessTrustProvider_Override(v VerifiedaccessTrustProvider, scope constructs.Construct, id *string, config *VerifiedaccessTrustProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.verifiedaccessTrustProvider.VerifiedaccessTrustProvider",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetDeviceTrustProviderType(val *string) {
	if err := j.validateSetDeviceTrustProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTrustProviderType",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetPolicyReferenceName(val *string) {
	if err := j.validateSetPolicyReferenceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyReferenceName",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetTrustProviderType(val *string) {
	if err := j.validateSetTrustProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustProviderType",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProvider)SetUserTrustProviderType(val *string) {
	if err := j.validateSetUserTrustProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTrustProviderType",
		val,
	)
}

// Generates CDKTF code for importing a VerifiedaccessTrustProvider resource upon running "cdktf plan <stack-name>".
func VerifiedaccessTrustProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVerifiedaccessTrustProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.verifiedaccessTrustProvider.VerifiedaccessTrustProvider",
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
func VerifiedaccessTrustProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVerifiedaccessTrustProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.verifiedaccessTrustProvider.VerifiedaccessTrustProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VerifiedaccessTrustProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVerifiedaccessTrustProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.verifiedaccessTrustProvider.VerifiedaccessTrustProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VerifiedaccessTrustProvider_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVerifiedaccessTrustProvider_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.verifiedaccessTrustProvider.VerifiedaccessTrustProvider",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VerifiedaccessTrustProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.verifiedaccessTrustProvider.VerifiedaccessTrustProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) PutDeviceOptions(value *VerifiedaccessTrustProviderDeviceOptions) {
	if err := v.validatePutDeviceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDeviceOptions",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) PutNativeApplicationOidcOptions(value *VerifiedaccessTrustProviderNativeApplicationOidcOptions) {
	if err := v.validatePutNativeApplicationOidcOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNativeApplicationOidcOptions",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) PutOidcOptions(value *VerifiedaccessTrustProviderOidcOptions) {
	if err := v.validatePutOidcOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOidcOptions",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) PutSseSpecification(value *VerifiedaccessTrustProviderSseSpecification) {
	if err := v.validatePutSseSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putSseSpecification",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) PutTimeouts(value *VerifiedaccessTrustProviderTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ResetDeviceOptions() {
	_jsii_.InvokeVoid(
		v,
		"resetDeviceOptions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ResetDeviceTrustProviderType() {
	_jsii_.InvokeVoid(
		v,
		"resetDeviceTrustProviderType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ResetNativeApplicationOidcOptions() {
	_jsii_.InvokeVoid(
		v,
		"resetNativeApplicationOidcOptions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ResetOidcOptions() {
	_jsii_.InvokeVoid(
		v,
		"resetOidcOptions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ResetSseSpecification() {
	_jsii_.InvokeVoid(
		v,
		"resetSseSpecification",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ResetTagsAll() {
	_jsii_.InvokeVoid(
		v,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ResetUserTrustProviderType() {
	_jsii_.InvokeVoid(
		v,
		"resetUserTrustProviderType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

