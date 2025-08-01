// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretsmanagersecretversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/secretsmanagersecretversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/secretsmanager_secret_version aws_secretsmanager_secret_version}.
type SecretsmanagerSecretVersion interface {
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HasSecretStringWo() cdktf.IResolvable
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SecretBinary() *string
	SetSecretBinary(val *string)
	SecretBinaryInput() *string
	SecretId() *string
	SetSecretId(val *string)
	SecretIdInput() *string
	SecretString() *string
	SetSecretString(val *string)
	SecretStringInput() *string
	SecretStringWo() *string
	SetSecretStringWo(val *string)
	SecretStringWoInput() *string
	SecretStringWoVersion() *float64
	SetSecretStringWoVersion(val *float64)
	SecretStringWoVersionInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VersionId() *string
	VersionStages() *[]*string
	SetVersionStages(val *[]*string)
	VersionStagesInput() *[]*string
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
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSecretBinary()
	ResetSecretString()
	ResetSecretStringWo()
	ResetSecretStringWoVersion()
	ResetVersionStages()
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

// The jsii proxy struct for SecretsmanagerSecretVersion
type jsiiProxy_SecretsmanagerSecretVersion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) HasSecretStringWo() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasSecretStringWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) SecretBinary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretBinary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) SecretBinaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretBinaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) SecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) SecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) SecretString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) SecretStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) SecretStringWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretStringWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) SecretStringWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretStringWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) SecretStringWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secretStringWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) SecretStringWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secretStringWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) VersionStages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versionStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretVersion) VersionStagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versionStagesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/secretsmanager_secret_version aws_secretsmanager_secret_version} Resource.
func NewSecretsmanagerSecretVersion(scope constructs.Construct, id *string, config *SecretsmanagerSecretVersionConfig) SecretsmanagerSecretVersion {
	_init_.Initialize()

	if err := validateNewSecretsmanagerSecretVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretsmanagerSecretVersion{}

	_jsii_.Create(
		"@cdktf/provider-aws.secretsmanagerSecretVersion.SecretsmanagerSecretVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/secretsmanager_secret_version aws_secretsmanager_secret_version} Resource.
func NewSecretsmanagerSecretVersion_Override(s SecretsmanagerSecretVersion, scope constructs.Construct, id *string, config *SecretsmanagerSecretVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.secretsmanagerSecretVersion.SecretsmanagerSecretVersion",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetSecretBinary(val *string) {
	if err := j.validateSetSecretBinaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretBinary",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetSecretId(val *string) {
	if err := j.validateSetSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretId",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetSecretString(val *string) {
	if err := j.validateSetSecretStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretString",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetSecretStringWo(val *string) {
	if err := j.validateSetSecretStringWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretStringWo",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetSecretStringWoVersion(val *float64) {
	if err := j.validateSetSecretStringWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretStringWoVersion",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretVersion)SetVersionStages(val *[]*string) {
	if err := j.validateSetVersionStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionStages",
		val,
	)
}

// Generates CDKTF code for importing a SecretsmanagerSecretVersion resource upon running "cdktf plan <stack-name>".
func SecretsmanagerSecretVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSecretsmanagerSecretVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.secretsmanagerSecretVersion.SecretsmanagerSecretVersion",
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
func SecretsmanagerSecretVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsmanagerSecretVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.secretsmanagerSecretVersion.SecretsmanagerSecretVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecretsmanagerSecretVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsmanagerSecretVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.secretsmanagerSecretVersion.SecretsmanagerSecretVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecretsmanagerSecretVersion_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsmanagerSecretVersion_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.secretsmanagerSecretVersion.SecretsmanagerSecretVersion",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecretsmanagerSecretVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.secretsmanagerSecretVersion.SecretsmanagerSecretVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecretsmanagerSecretVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecretsmanagerSecretVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecretsmanagerSecretVersion) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecretsmanagerSecretVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecretsmanagerSecretVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecretsmanagerSecretVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecretsmanagerSecretVersion) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecretsmanagerSecretVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecretsmanagerSecretVersion) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecretsmanagerSecretVersion) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) ResetRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) ResetSecretBinary() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretBinary",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) ResetSecretString() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretString",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) ResetSecretStringWo() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretStringWo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) ResetSecretStringWoVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretStringWoVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) ResetVersionStages() {
	_jsii_.InvokeVoid(
		s,
		"resetVersionStages",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsmanagerSecretVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

