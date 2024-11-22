// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2configurationset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/sesv2configurationset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/sesv2_configuration_set aws_sesv2_configuration_set}.
type Sesv2ConfigurationSet interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigurationSetName() *string
	SetConfigurationSetName(val *string)
	ConfigurationSetNameInput() *string
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
	DeliveryOptions() Sesv2ConfigurationSetDeliveryOptionsOutputReference
	DeliveryOptionsInput() *Sesv2ConfigurationSetDeliveryOptions
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
	SetId(val *string)
	IdInput() *string
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
	ReputationOptions() Sesv2ConfigurationSetReputationOptionsOutputReference
	ReputationOptionsInput() *Sesv2ConfigurationSetReputationOptions
	SendingOptions() Sesv2ConfigurationSetSendingOptionsOutputReference
	SendingOptionsInput() *Sesv2ConfigurationSetSendingOptions
	SuppressionOptions() Sesv2ConfigurationSetSuppressionOptionsOutputReference
	SuppressionOptionsInput() *Sesv2ConfigurationSetSuppressionOptions
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
	TrackingOptions() Sesv2ConfigurationSetTrackingOptionsOutputReference
	TrackingOptionsInput() *Sesv2ConfigurationSetTrackingOptions
	VdmOptions() Sesv2ConfigurationSetVdmOptionsOutputReference
	VdmOptionsInput() *Sesv2ConfigurationSetVdmOptions
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
	PutDeliveryOptions(value *Sesv2ConfigurationSetDeliveryOptions)
	PutReputationOptions(value *Sesv2ConfigurationSetReputationOptions)
	PutSendingOptions(value *Sesv2ConfigurationSetSendingOptions)
	PutSuppressionOptions(value *Sesv2ConfigurationSetSuppressionOptions)
	PutTrackingOptions(value *Sesv2ConfigurationSetTrackingOptions)
	PutVdmOptions(value *Sesv2ConfigurationSetVdmOptions)
	ResetDeliveryOptions()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReputationOptions()
	ResetSendingOptions()
	ResetSuppressionOptions()
	ResetTags()
	ResetTagsAll()
	ResetTrackingOptions()
	ResetVdmOptions()
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

// The jsii proxy struct for Sesv2ConfigurationSet
type jsiiProxy_Sesv2ConfigurationSet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Sesv2ConfigurationSet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) ConfigurationSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) ConfigurationSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) DeliveryOptions() Sesv2ConfigurationSetDeliveryOptionsOutputReference {
	var returns Sesv2ConfigurationSetDeliveryOptionsOutputReference
	_jsii_.Get(
		j,
		"deliveryOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) DeliveryOptionsInput() *Sesv2ConfigurationSetDeliveryOptions {
	var returns *Sesv2ConfigurationSetDeliveryOptions
	_jsii_.Get(
		j,
		"deliveryOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) ReputationOptions() Sesv2ConfigurationSetReputationOptionsOutputReference {
	var returns Sesv2ConfigurationSetReputationOptionsOutputReference
	_jsii_.Get(
		j,
		"reputationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) ReputationOptionsInput() *Sesv2ConfigurationSetReputationOptions {
	var returns *Sesv2ConfigurationSetReputationOptions
	_jsii_.Get(
		j,
		"reputationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) SendingOptions() Sesv2ConfigurationSetSendingOptionsOutputReference {
	var returns Sesv2ConfigurationSetSendingOptionsOutputReference
	_jsii_.Get(
		j,
		"sendingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) SendingOptionsInput() *Sesv2ConfigurationSetSendingOptions {
	var returns *Sesv2ConfigurationSetSendingOptions
	_jsii_.Get(
		j,
		"sendingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) SuppressionOptions() Sesv2ConfigurationSetSuppressionOptionsOutputReference {
	var returns Sesv2ConfigurationSetSuppressionOptionsOutputReference
	_jsii_.Get(
		j,
		"suppressionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) SuppressionOptionsInput() *Sesv2ConfigurationSetSuppressionOptions {
	var returns *Sesv2ConfigurationSetSuppressionOptions
	_jsii_.Get(
		j,
		"suppressionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) TrackingOptions() Sesv2ConfigurationSetTrackingOptionsOutputReference {
	var returns Sesv2ConfigurationSetTrackingOptionsOutputReference
	_jsii_.Get(
		j,
		"trackingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) TrackingOptionsInput() *Sesv2ConfigurationSetTrackingOptions {
	var returns *Sesv2ConfigurationSetTrackingOptions
	_jsii_.Get(
		j,
		"trackingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) VdmOptions() Sesv2ConfigurationSetVdmOptionsOutputReference {
	var returns Sesv2ConfigurationSetVdmOptionsOutputReference
	_jsii_.Get(
		j,
		"vdmOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSet) VdmOptionsInput() *Sesv2ConfigurationSetVdmOptions {
	var returns *Sesv2ConfigurationSetVdmOptions
	_jsii_.Get(
		j,
		"vdmOptionsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/sesv2_configuration_set aws_sesv2_configuration_set} Resource.
func NewSesv2ConfigurationSet(scope constructs.Construct, id *string, config *Sesv2ConfigurationSetConfig) Sesv2ConfigurationSet {
	_init_.Initialize()

	if err := validateNewSesv2ConfigurationSetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Sesv2ConfigurationSet{}

	_jsii_.Create(
		"@cdktf/provider-aws.sesv2ConfigurationSet.Sesv2ConfigurationSet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/sesv2_configuration_set aws_sesv2_configuration_set} Resource.
func NewSesv2ConfigurationSet_Override(s Sesv2ConfigurationSet, scope constructs.Construct, id *string, config *Sesv2ConfigurationSetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sesv2ConfigurationSet.Sesv2ConfigurationSet",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSet)SetConfigurationSetName(val *string) {
	if err := j.validateSetConfigurationSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationSetName",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSet)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSet)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSet)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a Sesv2ConfigurationSet resource upon running "cdktf plan <stack-name>".
func Sesv2ConfigurationSet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSesv2ConfigurationSet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sesv2ConfigurationSet.Sesv2ConfigurationSet",
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
func Sesv2ConfigurationSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSesv2ConfigurationSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sesv2ConfigurationSet.Sesv2ConfigurationSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Sesv2ConfigurationSet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSesv2ConfigurationSet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sesv2ConfigurationSet.Sesv2ConfigurationSet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Sesv2ConfigurationSet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSesv2ConfigurationSet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sesv2ConfigurationSet.Sesv2ConfigurationSet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Sesv2ConfigurationSet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.sesv2ConfigurationSet.Sesv2ConfigurationSet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) PutDeliveryOptions(value *Sesv2ConfigurationSetDeliveryOptions) {
	if err := s.validatePutDeliveryOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDeliveryOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) PutReputationOptions(value *Sesv2ConfigurationSetReputationOptions) {
	if err := s.validatePutReputationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReputationOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) PutSendingOptions(value *Sesv2ConfigurationSetSendingOptions) {
	if err := s.validatePutSendingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSendingOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) PutSuppressionOptions(value *Sesv2ConfigurationSetSuppressionOptions) {
	if err := s.validatePutSuppressionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSuppressionOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) PutTrackingOptions(value *Sesv2ConfigurationSetTrackingOptions) {
	if err := s.validatePutTrackingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTrackingOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) PutVdmOptions(value *Sesv2ConfigurationSetVdmOptions) {
	if err := s.validatePutVdmOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVdmOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ResetDeliveryOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetDeliveryOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ResetReputationOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetReputationOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ResetSendingOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetSendingOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ResetSuppressionOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetSuppressionOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ResetTrackingOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetTrackingOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ResetVdmOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetVdmOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

