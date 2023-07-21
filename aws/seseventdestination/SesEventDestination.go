package seseventdestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/seseventdestination/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/ses_event_destination aws_ses_event_destination}.
type SesEventDestination interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudwatchDestination() SesEventDestinationCloudwatchDestinationList
	CloudwatchDestinationInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	KinesisDestination() SesEventDestinationKinesisDestinationOutputReference
	KinesisDestinationInput() *SesEventDestinationKinesisDestination
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MatchingTypes() *[]*string
	SetMatchingTypes(val *[]*string)
	MatchingTypesInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	SnsDestination() SesEventDestinationSnsDestinationOutputReference
	SnsDestinationInput() *SesEventDestinationSnsDestination
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
	PutCloudwatchDestination(value interface{})
	PutKinesisDestination(value *SesEventDestinationKinesisDestination)
	PutSnsDestination(value *SesEventDestinationSnsDestination)
	ResetCloudwatchDestination()
	ResetEnabled()
	ResetId()
	ResetKinesisDestination()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSnsDestination()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SesEventDestination
type jsiiProxy_SesEventDestination struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SesEventDestination) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) CloudwatchDestination() SesEventDestinationCloudwatchDestinationList {
	var returns SesEventDestinationCloudwatchDestinationList
	_jsii_.Get(
		j,
		"cloudwatchDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) CloudwatchDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) ConfigurationSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) ConfigurationSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) KinesisDestination() SesEventDestinationKinesisDestinationOutputReference {
	var returns SesEventDestinationKinesisDestinationOutputReference
	_jsii_.Get(
		j,
		"kinesisDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) KinesisDestinationInput() *SesEventDestinationKinesisDestination {
	var returns *SesEventDestinationKinesisDestination
	_jsii_.Get(
		j,
		"kinesisDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) MatchingTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchingTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) MatchingTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchingTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) SnsDestination() SesEventDestinationSnsDestinationOutputReference {
	var returns SesEventDestinationSnsDestinationOutputReference
	_jsii_.Get(
		j,
		"snsDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) SnsDestinationInput() *SesEventDestinationSnsDestination {
	var returns *SesEventDestinationSnsDestination
	_jsii_.Get(
		j,
		"snsDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEventDestination) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/ses_event_destination aws_ses_event_destination} Resource.
func NewSesEventDestination(scope constructs.Construct, id *string, config *SesEventDestinationConfig) SesEventDestination {
	_init_.Initialize()

	if err := validateNewSesEventDestinationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SesEventDestination{}

	_jsii_.Create(
		"@cdktf/provider-aws.sesEventDestination.SesEventDestination",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/ses_event_destination aws_ses_event_destination} Resource.
func NewSesEventDestination_Override(s SesEventDestination, scope constructs.Construct, id *string, config *SesEventDestinationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sesEventDestination.SesEventDestination",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SesEventDestination)SetConfigurationSetName(val *string) {
	if err := j.validateSetConfigurationSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationSetName",
		val,
	)
}

func (j *jsiiProxy_SesEventDestination)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SesEventDestination)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SesEventDestination)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SesEventDestination)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SesEventDestination)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SesEventDestination)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SesEventDestination)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SesEventDestination)SetMatchingTypes(val *[]*string) {
	if err := j.validateSetMatchingTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchingTypes",
		val,
	)
}

func (j *jsiiProxy_SesEventDestination)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SesEventDestination)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SesEventDestination)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func SesEventDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSesEventDestination_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sesEventDestination.SesEventDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SesEventDestination_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSesEventDestination_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sesEventDestination.SesEventDestination",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SesEventDestination_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSesEventDestination_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sesEventDestination.SesEventDestination",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SesEventDestination_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.sesEventDestination.SesEventDestination",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SesEventDestination) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SesEventDestination) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SesEventDestination) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SesEventDestination) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SesEventDestination) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SesEventDestination) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SesEventDestination) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SesEventDestination) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SesEventDestination) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SesEventDestination) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SesEventDestination) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SesEventDestination) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SesEventDestination) PutCloudwatchDestination(value interface{}) {
	if err := s.validatePutCloudwatchDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCloudwatchDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesEventDestination) PutKinesisDestination(value *SesEventDestinationKinesisDestination) {
	if err := s.validatePutKinesisDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKinesisDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesEventDestination) PutSnsDestination(value *SesEventDestinationSnsDestination) {
	if err := s.validatePutSnsDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSnsDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesEventDestination) ResetCloudwatchDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudwatchDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEventDestination) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEventDestination) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEventDestination) ResetKinesisDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetKinesisDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEventDestination) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEventDestination) ResetSnsDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetSnsDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEventDestination) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEventDestination) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEventDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEventDestination) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

