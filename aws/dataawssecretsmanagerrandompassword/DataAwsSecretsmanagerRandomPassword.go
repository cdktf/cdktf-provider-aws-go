// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawssecretsmanagerrandompassword

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dataawssecretsmanagerrandompassword/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/data-sources/secretsmanager_random_password aws_secretsmanager_random_password}.
type DataAwsSecretsmanagerRandomPassword interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	ExcludeCharacters() *string
	SetExcludeCharacters(val *string)
	ExcludeCharactersInput() *string
	ExcludeLowercase() interface{}
	SetExcludeLowercase(val interface{})
	ExcludeLowercaseInput() interface{}
	ExcludeNumbers() interface{}
	SetExcludeNumbers(val interface{})
	ExcludeNumbersInput() interface{}
	ExcludePunctuation() interface{}
	SetExcludePunctuation(val interface{})
	ExcludePunctuationInput() interface{}
	ExcludeUppercase() interface{}
	SetExcludeUppercase(val interface{})
	ExcludeUppercaseInput() interface{}
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
	IncludeSpace() interface{}
	SetIncludeSpace(val interface{})
	IncludeSpaceInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PasswordLength() *float64
	SetPasswordLength(val *float64)
	PasswordLengthInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	RandomPassword() *string
	// Experimental.
	RawOverrides() interface{}
	RequireEachIncludedType() interface{}
	SetRequireEachIncludedType(val interface{})
	RequireEachIncludedTypeInput() interface{}
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
	ResetExcludeCharacters()
	ResetExcludeLowercase()
	ResetExcludeNumbers()
	ResetExcludePunctuation()
	ResetExcludeUppercase()
	ResetId()
	ResetIncludeSpace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasswordLength()
	ResetRequireEachIncludedType()
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

// The jsii proxy struct for DataAwsSecretsmanagerRandomPassword
type jsiiProxy_DataAwsSecretsmanagerRandomPassword struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ExcludeCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ExcludeCharactersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ExcludeLowercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeLowercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ExcludeLowercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeLowercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ExcludeNumbers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ExcludeNumbersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ExcludePunctuation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludePunctuation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ExcludePunctuationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludePunctuationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ExcludeUppercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeUppercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ExcludeUppercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeUppercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) IncludeSpace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) IncludeSpaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) PasswordLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) PasswordLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) RandomPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"randomPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) RequireEachIncludedType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEachIncludedType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) RequireEachIncludedTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEachIncludedTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/data-sources/secretsmanager_random_password aws_secretsmanager_random_password} Data Source.
func NewDataAwsSecretsmanagerRandomPassword(scope constructs.Construct, id *string, config *DataAwsSecretsmanagerRandomPasswordConfig) DataAwsSecretsmanagerRandomPassword {
	_init_.Initialize()

	if err := validateNewDataAwsSecretsmanagerRandomPasswordParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsSecretsmanagerRandomPassword{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsSecretsmanagerRandomPassword.DataAwsSecretsmanagerRandomPassword",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/data-sources/secretsmanager_random_password aws_secretsmanager_random_password} Data Source.
func NewDataAwsSecretsmanagerRandomPassword_Override(d DataAwsSecretsmanagerRandomPassword, scope constructs.Construct, id *string, config *DataAwsSecretsmanagerRandomPasswordConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsSecretsmanagerRandomPassword.DataAwsSecretsmanagerRandomPassword",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetExcludeCharacters(val *string) {
	if err := j.validateSetExcludeCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeCharacters",
		val,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetExcludeLowercase(val interface{}) {
	if err := j.validateSetExcludeLowercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeLowercase",
		val,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetExcludeNumbers(val interface{}) {
	if err := j.validateSetExcludeNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeNumbers",
		val,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetExcludePunctuation(val interface{}) {
	if err := j.validateSetExcludePunctuationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePunctuation",
		val,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetExcludeUppercase(val interface{}) {
	if err := j.validateSetExcludeUppercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeUppercase",
		val,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetIncludeSpace(val interface{}) {
	if err := j.validateSetIncludeSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSpace",
		val,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetPasswordLength(val *float64) {
	if err := j.validateSetPasswordLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordLength",
		val,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsSecretsmanagerRandomPassword)SetRequireEachIncludedType(val interface{}) {
	if err := j.validateSetRequireEachIncludedTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireEachIncludedType",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsSecretsmanagerRandomPassword resource upon running "cdktf plan <stack-name>".
func DataAwsSecretsmanagerRandomPassword_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsSecretsmanagerRandomPassword_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsSecretsmanagerRandomPassword.DataAwsSecretsmanagerRandomPassword",
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
func DataAwsSecretsmanagerRandomPassword_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsSecretsmanagerRandomPassword_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsSecretsmanagerRandomPassword.DataAwsSecretsmanagerRandomPassword",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsSecretsmanagerRandomPassword_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsSecretsmanagerRandomPassword_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsSecretsmanagerRandomPassword.DataAwsSecretsmanagerRandomPassword",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsSecretsmanagerRandomPassword_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsSecretsmanagerRandomPassword_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsSecretsmanagerRandomPassword.DataAwsSecretsmanagerRandomPassword",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsSecretsmanagerRandomPassword_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsSecretsmanagerRandomPassword.DataAwsSecretsmanagerRandomPassword",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ResetExcludeCharacters() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeCharacters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ResetExcludeLowercase() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeLowercase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ResetExcludeNumbers() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeNumbers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ResetExcludePunctuation() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludePunctuation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ResetExcludeUppercase() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeUppercase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ResetIncludeSpace() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeSpace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ResetPasswordLength() {
	_jsii_.InvokeVoid(
		d,
		"resetPasswordLength",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ResetRequireEachIncludedType() {
	_jsii_.InvokeVoid(
		d,
		"resetRequireEachIncludedType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSecretsmanagerRandomPassword) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

