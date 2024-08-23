// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devicefarmnetworkprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/devicefarmnetworkprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/devicefarm_network_profile aws_devicefarm_network_profile}.
type DevicefarmNetworkProfile interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DownlinkBandwidthBits() *float64
	SetDownlinkBandwidthBits(val *float64)
	DownlinkBandwidthBitsInput() *float64
	DownlinkDelayMs() *float64
	SetDownlinkDelayMs(val *float64)
	DownlinkDelayMsInput() *float64
	DownlinkJitterMs() *float64
	SetDownlinkJitterMs(val *float64)
	DownlinkJitterMsInput() *float64
	DownlinkLossPercent() *float64
	SetDownlinkLossPercent(val *float64)
	DownlinkLossPercentInput() *float64
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProjectArn() *string
	SetProjectArn(val *string)
	ProjectArnInput() *string
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
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UplinkBandwidthBits() *float64
	SetUplinkBandwidthBits(val *float64)
	UplinkBandwidthBitsInput() *float64
	UplinkDelayMs() *float64
	SetUplinkDelayMs(val *float64)
	UplinkDelayMsInput() *float64
	UplinkJitterMs() *float64
	SetUplinkJitterMs(val *float64)
	UplinkJitterMsInput() *float64
	UplinkLossPercent() *float64
	SetUplinkLossPercent(val *float64)
	UplinkLossPercentInput() *float64
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
	ResetDescription()
	ResetDownlinkBandwidthBits()
	ResetDownlinkDelayMs()
	ResetDownlinkJitterMs()
	ResetDownlinkLossPercent()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetType()
	ResetUplinkBandwidthBits()
	ResetUplinkDelayMs()
	ResetUplinkJitterMs()
	ResetUplinkLossPercent()
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

// The jsii proxy struct for DevicefarmNetworkProfile
type jsiiProxy_DevicefarmNetworkProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DevicefarmNetworkProfile) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) DownlinkBandwidthBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downlinkBandwidthBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) DownlinkBandwidthBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downlinkBandwidthBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) DownlinkDelayMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downlinkDelayMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) DownlinkDelayMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downlinkDelayMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) DownlinkJitterMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downlinkJitterMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) DownlinkJitterMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downlinkJitterMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) DownlinkLossPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downlinkLossPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) DownlinkLossPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downlinkLossPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) ProjectArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) ProjectArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) UplinkBandwidthBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uplinkBandwidthBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) UplinkBandwidthBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uplinkBandwidthBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) UplinkDelayMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uplinkDelayMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) UplinkDelayMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uplinkDelayMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) UplinkJitterMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uplinkJitterMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) UplinkJitterMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uplinkJitterMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) UplinkLossPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uplinkLossPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicefarmNetworkProfile) UplinkLossPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uplinkLossPercentInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/devicefarm_network_profile aws_devicefarm_network_profile} Resource.
func NewDevicefarmNetworkProfile(scope constructs.Construct, id *string, config *DevicefarmNetworkProfileConfig) DevicefarmNetworkProfile {
	_init_.Initialize()

	if err := validateNewDevicefarmNetworkProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DevicefarmNetworkProfile{}

	_jsii_.Create(
		"@cdktf/provider-aws.devicefarmNetworkProfile.DevicefarmNetworkProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/devicefarm_network_profile aws_devicefarm_network_profile} Resource.
func NewDevicefarmNetworkProfile_Override(d DevicefarmNetworkProfile, scope constructs.Construct, id *string, config *DevicefarmNetworkProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.devicefarmNetworkProfile.DevicefarmNetworkProfile",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetDownlinkBandwidthBits(val *float64) {
	if err := j.validateSetDownlinkBandwidthBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"downlinkBandwidthBits",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetDownlinkDelayMs(val *float64) {
	if err := j.validateSetDownlinkDelayMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"downlinkDelayMs",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetDownlinkJitterMs(val *float64) {
	if err := j.validateSetDownlinkJitterMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"downlinkJitterMs",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetDownlinkLossPercent(val *float64) {
	if err := j.validateSetDownlinkLossPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"downlinkLossPercent",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetProjectArn(val *string) {
	if err := j.validateSetProjectArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectArn",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetUplinkBandwidthBits(val *float64) {
	if err := j.validateSetUplinkBandwidthBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uplinkBandwidthBits",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetUplinkDelayMs(val *float64) {
	if err := j.validateSetUplinkDelayMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uplinkDelayMs",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetUplinkJitterMs(val *float64) {
	if err := j.validateSetUplinkJitterMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uplinkJitterMs",
		val,
	)
}

func (j *jsiiProxy_DevicefarmNetworkProfile)SetUplinkLossPercent(val *float64) {
	if err := j.validateSetUplinkLossPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uplinkLossPercent",
		val,
	)
}

// Generates CDKTF code for importing a DevicefarmNetworkProfile resource upon running "cdktf plan <stack-name>".
func DevicefarmNetworkProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDevicefarmNetworkProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.devicefarmNetworkProfile.DevicefarmNetworkProfile",
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
func DevicefarmNetworkProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevicefarmNetworkProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.devicefarmNetworkProfile.DevicefarmNetworkProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DevicefarmNetworkProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevicefarmNetworkProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.devicefarmNetworkProfile.DevicefarmNetworkProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DevicefarmNetworkProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevicefarmNetworkProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.devicefarmNetworkProfile.DevicefarmNetworkProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DevicefarmNetworkProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.devicefarmNetworkProfile.DevicefarmNetworkProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DevicefarmNetworkProfile) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DevicefarmNetworkProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DevicefarmNetworkProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DevicefarmNetworkProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DevicefarmNetworkProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DevicefarmNetworkProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DevicefarmNetworkProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DevicefarmNetworkProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DevicefarmNetworkProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DevicefarmNetworkProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DevicefarmNetworkProfile) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetDownlinkBandwidthBits() {
	_jsii_.InvokeVoid(
		d,
		"resetDownlinkBandwidthBits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetDownlinkDelayMs() {
	_jsii_.InvokeVoid(
		d,
		"resetDownlinkDelayMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetDownlinkJitterMs() {
	_jsii_.InvokeVoid(
		d,
		"resetDownlinkJitterMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetDownlinkLossPercent() {
	_jsii_.InvokeVoid(
		d,
		"resetDownlinkLossPercent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetUplinkBandwidthBits() {
	_jsii_.InvokeVoid(
		d,
		"resetUplinkBandwidthBits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetUplinkDelayMs() {
	_jsii_.InvokeVoid(
		d,
		"resetUplinkDelayMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetUplinkJitterMs() {
	_jsii_.InvokeVoid(
		d,
		"resetUplinkJitterMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ResetUplinkLossPercent() {
	_jsii_.InvokeVoid(
		d,
		"resetUplinkLossPercent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicefarmNetworkProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevicefarmNetworkProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevicefarmNetworkProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

