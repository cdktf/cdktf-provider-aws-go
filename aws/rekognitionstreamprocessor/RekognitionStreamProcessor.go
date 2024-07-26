// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/rekognitionstreamprocessor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/rekognition_stream_processor aws_rekognition_stream_processor}.
type RekognitionStreamProcessor interface {
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
	DataSharingPreference() RekognitionStreamProcessorDataSharingPreferenceList
	DataSharingPreferenceInput() interface{}
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
	Input() RekognitionStreamProcessorInputList
	InputInput() interface{}
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotificationChannel() RekognitionStreamProcessorNotificationChannelList
	NotificationChannelInput() interface{}
	Output() RekognitionStreamProcessorOutputList
	OutputInput() interface{}
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
	RegionsOfInterest() RekognitionStreamProcessorRegionsOfInterestList
	RegionsOfInterestInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Settings() RekognitionStreamProcessorSettingsList
	SettingsInput() interface{}
	StreamProcessorArn() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktf.StringMap
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() RekognitionStreamProcessorTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutDataSharingPreference(value interface{})
	PutInput(value interface{})
	PutNotificationChannel(value interface{})
	PutOutput(value interface{})
	PutRegionsOfInterest(value interface{})
	PutSettings(value interface{})
	PutTimeouts(value *RekognitionStreamProcessorTimeouts)
	ResetDataSharingPreference()
	ResetInput()
	ResetKmsKeyId()
	ResetNotificationChannel()
	ResetOutput()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegionsOfInterest()
	ResetSettings()
	ResetTags()
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

// The jsii proxy struct for RekognitionStreamProcessor
type jsiiProxy_RekognitionStreamProcessor struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RekognitionStreamProcessor) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) DataSharingPreference() RekognitionStreamProcessorDataSharingPreferenceList {
	var returns RekognitionStreamProcessorDataSharingPreferenceList
	_jsii_.Get(
		j,
		"dataSharingPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) DataSharingPreferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSharingPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) Input() RekognitionStreamProcessorInputList {
	var returns RekognitionStreamProcessorInputList
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) InputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) NotificationChannel() RekognitionStreamProcessorNotificationChannelList {
	var returns RekognitionStreamProcessorNotificationChannelList
	_jsii_.Get(
		j,
		"notificationChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) NotificationChannelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) Output() RekognitionStreamProcessorOutputList {
	var returns RekognitionStreamProcessorOutputList
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) OutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) RegionsOfInterest() RekognitionStreamProcessorRegionsOfInterestList {
	var returns RekognitionStreamProcessorRegionsOfInterestList
	_jsii_.Get(
		j,
		"regionsOfInterest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) RegionsOfInterestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionsOfInterestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) Settings() RekognitionStreamProcessorSettingsList {
	var returns RekognitionStreamProcessorSettingsList
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) SettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) StreamProcessorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamProcessorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) Timeouts() RekognitionStreamProcessorTimeoutsOutputReference {
	var returns RekognitionStreamProcessorTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessor) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/rekognition_stream_processor aws_rekognition_stream_processor} Resource.
func NewRekognitionStreamProcessor(scope constructs.Construct, id *string, config *RekognitionStreamProcessorConfig) RekognitionStreamProcessor {
	_init_.Initialize()

	if err := validateNewRekognitionStreamProcessorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RekognitionStreamProcessor{}

	_jsii_.Create(
		"@cdktf/provider-aws.rekognitionStreamProcessor.RekognitionStreamProcessor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/rekognition_stream_processor aws_rekognition_stream_processor} Resource.
func NewRekognitionStreamProcessor_Override(r RekognitionStreamProcessor, scope constructs.Construct, id *string, config *RekognitionStreamProcessorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.rekognitionStreamProcessor.RekognitionStreamProcessor",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessor)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessor)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessor)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessor)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessor)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessor)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessor)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessor)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessor)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessor)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessor)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a RekognitionStreamProcessor resource upon running "cdktf plan <stack-name>".
func RekognitionStreamProcessor_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRekognitionStreamProcessor_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.rekognitionStreamProcessor.RekognitionStreamProcessor",
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
func RekognitionStreamProcessor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRekognitionStreamProcessor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.rekognitionStreamProcessor.RekognitionStreamProcessor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RekognitionStreamProcessor_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRekognitionStreamProcessor_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.rekognitionStreamProcessor.RekognitionStreamProcessor",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RekognitionStreamProcessor_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRekognitionStreamProcessor_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.rekognitionStreamProcessor.RekognitionStreamProcessor",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RekognitionStreamProcessor_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.rekognitionStreamProcessor.RekognitionStreamProcessor",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessor) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RekognitionStreamProcessor) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RekognitionStreamProcessor) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RekognitionStreamProcessor) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RekognitionStreamProcessor) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RekognitionStreamProcessor) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RekognitionStreamProcessor) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RekognitionStreamProcessor) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RekognitionStreamProcessor) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RekognitionStreamProcessor) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessor) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RekognitionStreamProcessor) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) PutDataSharingPreference(value interface{}) {
	if err := r.validatePutDataSharingPreferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putDataSharingPreference",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) PutInput(value interface{}) {
	if err := r.validatePutInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putInput",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) PutNotificationChannel(value interface{}) {
	if err := r.validatePutNotificationChannelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putNotificationChannel",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) PutOutput(value interface{}) {
	if err := r.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putOutput",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) PutRegionsOfInterest(value interface{}) {
	if err := r.validatePutRegionsOfInterestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRegionsOfInterest",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) PutSettings(value interface{}) {
	if err := r.validatePutSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSettings",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) PutTimeouts(value *RekognitionStreamProcessorTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) ResetDataSharingPreference() {
	_jsii_.InvokeVoid(
		r,
		"resetDataSharingPreference",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) ResetInput() {
	_jsii_.InvokeVoid(
		r,
		"resetInput",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) ResetNotificationChannel() {
	_jsii_.InvokeVoid(
		r,
		"resetNotificationChannel",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) ResetOutput() {
	_jsii_.InvokeVoid(
		r,
		"resetOutput",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) ResetRegionsOfInterest() {
	_jsii_.InvokeVoid(
		r,
		"resetRegionsOfInterest",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) ResetSettings() {
	_jsii_.InvokeVoid(
		r,
		"resetSettings",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessor) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessor) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessor) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessor) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessor) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

