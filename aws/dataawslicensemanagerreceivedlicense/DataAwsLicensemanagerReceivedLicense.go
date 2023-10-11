// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslicensemanagerreceivedlicense

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/dataawslicensemanagerreceivedlicense/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/data-sources/licensemanager_received_license aws_licensemanager_received_license}.
type DataAwsLicensemanagerReceivedLicense interface {
	cdktf.TerraformDataSource
	Beneficiary() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsumptionConfiguration() DataAwsLicensemanagerReceivedLicenseConsumptionConfigurationList
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Entitlements() DataAwsLicensemanagerReceivedLicenseEntitlementsList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HomeRegion() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Issuer() DataAwsLicensemanagerReceivedLicenseIssuerList
	LicenseArn() *string
	SetLicenseArn(val *string)
	LicenseArnInput() *string
	LicenseMetadata() DataAwsLicensemanagerReceivedLicenseLicenseMetadataList
	LicenseName() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	ProductName() *string
	ProductSku() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ReceivedMetadata() DataAwsLicensemanagerReceivedLicenseReceivedMetadataList
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Validity() DataAwsLicensemanagerReceivedLicenseValidityList
	Version() *string
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
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsLicensemanagerReceivedLicense
type jsiiProxy_DataAwsLicensemanagerReceivedLicense struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) Beneficiary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beneficiary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) ConsumptionConfiguration() DataAwsLicensemanagerReceivedLicenseConsumptionConfigurationList {
	var returns DataAwsLicensemanagerReceivedLicenseConsumptionConfigurationList
	_jsii_.Get(
		j,
		"consumptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) Entitlements() DataAwsLicensemanagerReceivedLicenseEntitlementsList {
	var returns DataAwsLicensemanagerReceivedLicenseEntitlementsList
	_jsii_.Get(
		j,
		"entitlements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) HomeRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) Issuer() DataAwsLicensemanagerReceivedLicenseIssuerList {
	var returns DataAwsLicensemanagerReceivedLicenseIssuerList
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) LicenseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) LicenseArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) LicenseMetadata() DataAwsLicensemanagerReceivedLicenseLicenseMetadataList {
	var returns DataAwsLicensemanagerReceivedLicenseLicenseMetadataList
	_jsii_.Get(
		j,
		"licenseMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) LicenseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) ProductName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) ProductSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) ReceivedMetadata() DataAwsLicensemanagerReceivedLicenseReceivedMetadataList {
	var returns DataAwsLicensemanagerReceivedLicenseReceivedMetadataList
	_jsii_.Get(
		j,
		"receivedMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) Validity() DataAwsLicensemanagerReceivedLicenseValidityList {
	var returns DataAwsLicensemanagerReceivedLicenseValidityList
	_jsii_.Get(
		j,
		"validity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/data-sources/licensemanager_received_license aws_licensemanager_received_license} Data Source.
func NewDataAwsLicensemanagerReceivedLicense(scope constructs.Construct, id *string, config *DataAwsLicensemanagerReceivedLicenseConfig) DataAwsLicensemanagerReceivedLicense {
	_init_.Initialize()

	if err := validateNewDataAwsLicensemanagerReceivedLicenseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsLicensemanagerReceivedLicense{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsLicensemanagerReceivedLicense.DataAwsLicensemanagerReceivedLicense",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/data-sources/licensemanager_received_license aws_licensemanager_received_license} Data Source.
func NewDataAwsLicensemanagerReceivedLicense_Override(d DataAwsLicensemanagerReceivedLicense, scope constructs.Construct, id *string, config *DataAwsLicensemanagerReceivedLicenseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsLicensemanagerReceivedLicense.DataAwsLicensemanagerReceivedLicense",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense)SetLicenseArn(val *string) {
	if err := j.validateSetLicenseArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseArn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLicensemanagerReceivedLicense)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
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
func DataAwsLicensemanagerReceivedLicense_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsLicensemanagerReceivedLicense_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsLicensemanagerReceivedLicense.DataAwsLicensemanagerReceivedLicense",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsLicensemanagerReceivedLicense_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsLicensemanagerReceivedLicense_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsLicensemanagerReceivedLicense.DataAwsLicensemanagerReceivedLicense",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsLicensemanagerReceivedLicense_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsLicensemanagerReceivedLicense_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsLicensemanagerReceivedLicense.DataAwsLicensemanagerReceivedLicense",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLicensemanagerReceivedLicense_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsLicensemanagerReceivedLicense.DataAwsLicensemanagerReceivedLicense",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLicensemanagerReceivedLicense) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

