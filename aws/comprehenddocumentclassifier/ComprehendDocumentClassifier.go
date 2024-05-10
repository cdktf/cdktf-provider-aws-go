// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package comprehenddocumentclassifier

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/comprehenddocumentclassifier/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/comprehend_document_classifier aws_comprehend_document_classifier}.
type ComprehendDocumentClassifier interface {
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
	DataAccessRoleArn() *string
	SetDataAccessRoleArn(val *string)
	DataAccessRoleArnInput() *string
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
	InputDataConfig() ComprehendDocumentClassifierInputDataConfigOutputReference
	InputDataConfigInput() *ComprehendDocumentClassifierInputDataConfig
	LanguageCode() *string
	SetLanguageCode(val *string)
	LanguageCodeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	ModelKmsKeyId() *string
	SetModelKmsKeyId(val *string)
	ModelKmsKeyIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutputDataConfig() ComprehendDocumentClassifierOutputDataConfigOutputReference
	OutputDataConfigInput() *ComprehendDocumentClassifierOutputDataConfig
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
	Timeouts() ComprehendDocumentClassifierTimeoutsOutputReference
	TimeoutsInput() interface{}
	VersionName() *string
	SetVersionName(val *string)
	VersionNameInput() *string
	VersionNamePrefix() *string
	SetVersionNamePrefix(val *string)
	VersionNamePrefixInput() *string
	VolumeKmsKeyId() *string
	SetVolumeKmsKeyId(val *string)
	VolumeKmsKeyIdInput() *string
	VpcConfig() ComprehendDocumentClassifierVpcConfigOutputReference
	VpcConfigInput() *ComprehendDocumentClassifierVpcConfig
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
	PutInputDataConfig(value *ComprehendDocumentClassifierInputDataConfig)
	PutOutputDataConfig(value *ComprehendDocumentClassifierOutputDataConfig)
	PutTimeouts(value *ComprehendDocumentClassifierTimeouts)
	PutVpcConfig(value *ComprehendDocumentClassifierVpcConfig)
	ResetId()
	ResetMode()
	ResetModelKmsKeyId()
	ResetOutputDataConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetVersionName()
	ResetVersionNamePrefix()
	ResetVolumeKmsKeyId()
	ResetVpcConfig()
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

// The jsii proxy struct for ComprehendDocumentClassifier
type jsiiProxy_ComprehendDocumentClassifier struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComprehendDocumentClassifier) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) DataAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) DataAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) InputDataConfig() ComprehendDocumentClassifierInputDataConfigOutputReference {
	var returns ComprehendDocumentClassifierInputDataConfigOutputReference
	_jsii_.Get(
		j,
		"inputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) InputDataConfigInput() *ComprehendDocumentClassifierInputDataConfig {
	var returns *ComprehendDocumentClassifierInputDataConfig
	_jsii_.Get(
		j,
		"inputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) ModelKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) ModelKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) OutputDataConfig() ComprehendDocumentClassifierOutputDataConfigOutputReference {
	var returns ComprehendDocumentClassifierOutputDataConfigOutputReference
	_jsii_.Get(
		j,
		"outputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) OutputDataConfigInput() *ComprehendDocumentClassifierOutputDataConfig {
	var returns *ComprehendDocumentClassifierOutputDataConfig
	_jsii_.Get(
		j,
		"outputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) Timeouts() ComprehendDocumentClassifierTimeoutsOutputReference {
	var returns ComprehendDocumentClassifierTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) VersionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) VersionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) VersionNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) VersionNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) VolumeKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) VolumeKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) VpcConfig() ComprehendDocumentClassifierVpcConfigOutputReference {
	var returns ComprehendDocumentClassifierVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifier) VpcConfigInput() *ComprehendDocumentClassifierVpcConfig {
	var returns *ComprehendDocumentClassifierVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/comprehend_document_classifier aws_comprehend_document_classifier} Resource.
func NewComprehendDocumentClassifier(scope constructs.Construct, id *string, config *ComprehendDocumentClassifierConfig) ComprehendDocumentClassifier {
	_init_.Initialize()

	if err := validateNewComprehendDocumentClassifierParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComprehendDocumentClassifier{}

	_jsii_.Create(
		"@cdktf/provider-aws.comprehendDocumentClassifier.ComprehendDocumentClassifier",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/comprehend_document_classifier aws_comprehend_document_classifier} Resource.
func NewComprehendDocumentClassifier_Override(c ComprehendDocumentClassifier, scope constructs.Construct, id *string, config *ComprehendDocumentClassifierConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.comprehendDocumentClassifier.ComprehendDocumentClassifier",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetDataAccessRoleArn(val *string) {
	if err := j.validateSetDataAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetModelKmsKeyId(val *string) {
	if err := j.validateSetModelKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetVersionName(val *string) {
	if err := j.validateSetVersionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionName",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetVersionNamePrefix(val *string) {
	if err := j.validateSetVersionNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionNamePrefix",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifier)SetVolumeKmsKeyId(val *string) {
	if err := j.validateSetVolumeKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeKmsKeyId",
		val,
	)
}

// Generates CDKTF code for importing a ComprehendDocumentClassifier resource upon running "cdktf plan <stack-name>".
func ComprehendDocumentClassifier_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComprehendDocumentClassifier_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.comprehendDocumentClassifier.ComprehendDocumentClassifier",
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
func ComprehendDocumentClassifier_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComprehendDocumentClassifier_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.comprehendDocumentClassifier.ComprehendDocumentClassifier",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComprehendDocumentClassifier_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComprehendDocumentClassifier_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.comprehendDocumentClassifier.ComprehendDocumentClassifier",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComprehendDocumentClassifier_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComprehendDocumentClassifier_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.comprehendDocumentClassifier.ComprehendDocumentClassifier",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComprehendDocumentClassifier_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.comprehendDocumentClassifier.ComprehendDocumentClassifier",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComprehendDocumentClassifier) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComprehendDocumentClassifier) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComprehendDocumentClassifier) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComprehendDocumentClassifier) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComprehendDocumentClassifier) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComprehendDocumentClassifier) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComprehendDocumentClassifier) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComprehendDocumentClassifier) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComprehendDocumentClassifier) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComprehendDocumentClassifier) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComprehendDocumentClassifier) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) PutInputDataConfig(value *ComprehendDocumentClassifierInputDataConfig) {
	if err := c.validatePutInputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInputDataConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) PutOutputDataConfig(value *ComprehendDocumentClassifierOutputDataConfig) {
	if err := c.validatePutOutputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOutputDataConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) PutTimeouts(value *ComprehendDocumentClassifierTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) PutVpcConfig(value *ComprehendDocumentClassifierVpcConfig) {
	if err := c.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ResetMode() {
	_jsii_.InvokeVoid(
		c,
		"resetMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ResetModelKmsKeyId() {
	_jsii_.InvokeVoid(
		c,
		"resetModelKmsKeyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ResetOutputDataConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetOutputDataConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ResetVersionName() {
	_jsii_.InvokeVoid(
		c,
		"resetVersionName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ResetVersionNamePrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetVersionNamePrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ResetVolumeKmsKeyId() {
	_jsii_.InvokeVoid(
		c,
		"resetVolumeKmsKeyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifier) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendDocumentClassifier) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendDocumentClassifier) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

