// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsiampolicydocument

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dataawsiampolicydocument/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/data-sources/iam_policy_document aws_iam_policy_document}.
type DataAwsIamPolicyDocument interface {
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
	Json() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MinifiedJson() *string
	// The tree node.
	Node() constructs.Node
	OverrideJson() *string
	SetOverrideJson(val *string)
	OverrideJsonInput() *string
	OverridePolicyDocuments() *[]*string
	SetOverridePolicyDocuments(val *[]*string)
	OverridePolicyDocumentsInput() *[]*string
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SourceJson() *string
	SetSourceJson(val *string)
	SourceJsonInput() *string
	SourcePolicyDocuments() *[]*string
	SetSourcePolicyDocuments(val *[]*string)
	SourcePolicyDocumentsInput() *[]*string
	Statement() DataAwsIamPolicyDocumentStatementList
	StatementInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutStatement(value interface{})
	ResetId()
	ResetOverrideJson()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOverridePolicyDocuments()
	ResetPolicyId()
	ResetSourceJson()
	ResetSourcePolicyDocuments()
	ResetStatement()
	ResetVersion()
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

// The jsii proxy struct for DataAwsIamPolicyDocument
type jsiiProxy_DataAwsIamPolicyDocument struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) Json() *string {
	var returns *string
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) MinifiedJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minifiedJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) OverrideJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) OverrideJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) OverridePolicyDocuments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overridePolicyDocuments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) OverridePolicyDocumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overridePolicyDocumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) SourceJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) SourceJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) SourcePolicyDocuments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePolicyDocuments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) SourcePolicyDocumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePolicyDocumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) Statement() DataAwsIamPolicyDocumentStatementList {
	var returns DataAwsIamPolicyDocumentStatementList
	_jsii_.Get(
		j,
		"statement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) StatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocument) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/data-sources/iam_policy_document aws_iam_policy_document} Data Source.
func NewDataAwsIamPolicyDocument(scope constructs.Construct, id *string, config *DataAwsIamPolicyDocumentConfig) DataAwsIamPolicyDocument {
	_init_.Initialize()

	if err := validateNewDataAwsIamPolicyDocumentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsIamPolicyDocument{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/data-sources/iam_policy_document aws_iam_policy_document} Data Source.
func NewDataAwsIamPolicyDocument_Override(d DataAwsIamPolicyDocument, scope constructs.Construct, id *string, config *DataAwsIamPolicyDocumentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocument)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocument)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocument)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocument)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocument)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocument)SetOverrideJson(val *string) {
	if err := j.validateSetOverrideJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideJson",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocument)SetOverridePolicyDocuments(val *[]*string) {
	if err := j.validateSetOverridePolicyDocumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overridePolicyDocuments",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocument)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocument)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocument)SetSourceJson(val *string) {
	if err := j.validateSetSourceJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceJson",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocument)SetSourcePolicyDocuments(val *[]*string) {
	if err := j.validateSetSourcePolicyDocumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePolicyDocuments",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocument)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsIamPolicyDocument resource upon running "cdktf plan <stack-name>".
func DataAwsIamPolicyDocument_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsIamPolicyDocument_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument",
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
func DataAwsIamPolicyDocument_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsIamPolicyDocument_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsIamPolicyDocument_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsIamPolicyDocument_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsIamPolicyDocument_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsIamPolicyDocument_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsIamPolicyDocument_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsIamPolicyDocument) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsIamPolicyDocument) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsIamPolicyDocument) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsIamPolicyDocument) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsIamPolicyDocument) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsIamPolicyDocument) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsIamPolicyDocument) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsIamPolicyDocument) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsIamPolicyDocument) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsIamPolicyDocument) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) PutStatement(value interface{}) {
	if err := d.validatePutStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStatement",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) ResetOverrideJson() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) ResetOverridePolicyDocuments() {
	_jsii_.InvokeVoid(
		d,
		"resetOverridePolicyDocuments",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) ResetPolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) ResetSourceJson() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) ResetSourcePolicyDocuments() {
	_jsii_.InvokeVoid(
		d,
		"resetSourcePolicyDocuments",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) ResetStatement() {
	_jsii_.InvokeVoid(
		d,
		"resetStatement",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocument) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

