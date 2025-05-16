// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsresourcegroupstaggingapiresources

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dataawsresourcegroupstaggingapiresources/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/data-sources/resourcegroupstaggingapi_resources aws_resourcegroupstaggingapi_resources}.
type DataAwsResourcegroupstaggingapiResources interface {
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
	ExcludeCompliantResources() interface{}
	SetExcludeCompliantResources(val interface{})
	ExcludeCompliantResourcesInput() interface{}
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
	IncludeComplianceDetails() interface{}
	SetIncludeComplianceDetails(val interface{})
	IncludeComplianceDetailsInput() interface{}
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
	RawOverrides() interface{}
	ResourceArnList() *[]*string
	SetResourceArnList(val *[]*string)
	ResourceArnListInput() *[]*string
	ResourceTagMappingList() DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListStructList
	ResourceTypeFilters() *[]*string
	SetResourceTypeFilters(val *[]*string)
	ResourceTypeFiltersInput() *[]*string
	TagFilter() DataAwsResourcegroupstaggingapiResourcesTagFilterList
	TagFilterInput() interface{}
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
	PutTagFilter(value interface{})
	ResetExcludeCompliantResources()
	ResetId()
	ResetIncludeComplianceDetails()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceArnList()
	ResetResourceTypeFilters()
	ResetTagFilter()
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

// The jsii proxy struct for DataAwsResourcegroupstaggingapiResources
type jsiiProxy_DataAwsResourcegroupstaggingapiResources struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ExcludeCompliantResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCompliantResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ExcludeCompliantResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCompliantResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) IncludeComplianceDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeComplianceDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) IncludeComplianceDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeComplianceDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResourceArnList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResourceArnListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResourceTagMappingList() DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListStructList {
	var returns DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListStructList
	_jsii_.Get(
		j,
		"resourceTagMappingList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResourceTypeFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypeFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResourceTypeFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypeFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) TagFilter() DataAwsResourcegroupstaggingapiResourcesTagFilterList {
	var returns DataAwsResourcegroupstaggingapiResourcesTagFilterList
	_jsii_.Get(
		j,
		"tagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) TagFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/data-sources/resourcegroupstaggingapi_resources aws_resourcegroupstaggingapi_resources} Data Source.
func NewDataAwsResourcegroupstaggingapiResources(scope constructs.Construct, id *string, config *DataAwsResourcegroupstaggingapiResourcesConfig) DataAwsResourcegroupstaggingapiResources {
	_init_.Initialize()

	if err := validateNewDataAwsResourcegroupstaggingapiResourcesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsResourcegroupstaggingapiResources{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsResourcegroupstaggingapiResources.DataAwsResourcegroupstaggingapiResources",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/data-sources/resourcegroupstaggingapi_resources aws_resourcegroupstaggingapi_resources} Data Source.
func NewDataAwsResourcegroupstaggingapiResources_Override(d DataAwsResourcegroupstaggingapiResources, scope constructs.Construct, id *string, config *DataAwsResourcegroupstaggingapiResourcesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsResourcegroupstaggingapiResources.DataAwsResourcegroupstaggingapiResources",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources)SetExcludeCompliantResources(val interface{}) {
	if err := j.validateSetExcludeCompliantResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeCompliantResources",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources)SetIncludeComplianceDetails(val interface{}) {
	if err := j.validateSetIncludeComplianceDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeComplianceDetails",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources)SetResourceArnList(val *[]*string) {
	if err := j.validateSetResourceArnListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceArnList",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources)SetResourceTypeFilters(val *[]*string) {
	if err := j.validateSetResourceTypeFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypeFilters",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsResourcegroupstaggingapiResources resource upon running "cdktf plan <stack-name>".
func DataAwsResourcegroupstaggingapiResources_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsResourcegroupstaggingapiResources_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsResourcegroupstaggingapiResources.DataAwsResourcegroupstaggingapiResources",
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
func DataAwsResourcegroupstaggingapiResources_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsResourcegroupstaggingapiResources_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsResourcegroupstaggingapiResources.DataAwsResourcegroupstaggingapiResources",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsResourcegroupstaggingapiResources_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsResourcegroupstaggingapiResources_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsResourcegroupstaggingapiResources.DataAwsResourcegroupstaggingapiResources",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsResourcegroupstaggingapiResources_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsResourcegroupstaggingapiResources_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsResourcegroupstaggingapiResources.DataAwsResourcegroupstaggingapiResources",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsResourcegroupstaggingapiResources_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsResourcegroupstaggingapiResources.DataAwsResourcegroupstaggingapiResources",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) PutTagFilter(value interface{}) {
	if err := d.validatePutTagFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTagFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResetExcludeCompliantResources() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeCompliantResources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResetIncludeComplianceDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeComplianceDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResetResourceArnList() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceArnList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResetResourceTypeFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceTypeFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResetTagFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetTagFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

