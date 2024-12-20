// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsopensearchdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dataawsopensearchdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/data-sources/opensearch_domain aws_opensearch_domain}.
type DataAwsOpensearchDomain interface {
	cdktf.TerraformDataSource
	AccessPolicies() *string
	AdvancedOptions() cdktf.StringMap
	AdvancedSecurityOptions() DataAwsOpensearchDomainAdvancedSecurityOptionsList
	Arn() *string
	AutoTuneOptions() DataAwsOpensearchDomainAutoTuneOptionsList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterConfig() DataAwsOpensearchDomainClusterConfigList
	CognitoOptions() DataAwsOpensearchDomainCognitoOptionsList
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Created() cdktf.IResolvable
	DashboardEndpoint() *string
	DashboardEndpointV2() *string
	Deleted() cdktf.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DomainEndpointV2HostedZoneId() *string
	DomainId() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	EbsOptions() DataAwsOpensearchDomainEbsOptionsList
	EncryptionAtRest() DataAwsOpensearchDomainEncryptionAtRestList
	Endpoint() *string
	EndpointV2() *string
	EngineVersion() *string
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
	IpAddressType() *string
	KibanaEndpoint() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogPublishingOptions() DataAwsOpensearchDomainLogPublishingOptionsList
	// The tree node.
	Node() constructs.Node
	NodeToNodeEncryption() DataAwsOpensearchDomainNodeToNodeEncryptionList
	OffPeakWindowOptions() DataAwsOpensearchDomainOffPeakWindowOptionsOutputReference
	OffPeakWindowOptionsInput() *DataAwsOpensearchDomainOffPeakWindowOptions
	Processing() cdktf.IResolvable
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SnapshotOptions() DataAwsOpensearchDomainSnapshotOptionsList
	SoftwareUpdateOptions() DataAwsOpensearchDomainSoftwareUpdateOptionsList
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VpcOptions() DataAwsOpensearchDomainVpcOptionsList
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
	PutOffPeakWindowOptions(value *DataAwsOpensearchDomainOffPeakWindowOptions)
	ResetId()
	ResetOffPeakWindowOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
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

// The jsii proxy struct for DataAwsOpensearchDomain
type jsiiProxy_DataAwsOpensearchDomain struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsOpensearchDomain) AccessPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) AdvancedOptions() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"advancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) AdvancedSecurityOptions() DataAwsOpensearchDomainAdvancedSecurityOptionsList {
	var returns DataAwsOpensearchDomainAdvancedSecurityOptionsList
	_jsii_.Get(
		j,
		"advancedSecurityOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) AutoTuneOptions() DataAwsOpensearchDomainAutoTuneOptionsList {
	var returns DataAwsOpensearchDomainAutoTuneOptionsList
	_jsii_.Get(
		j,
		"autoTuneOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) ClusterConfig() DataAwsOpensearchDomainClusterConfigList {
	var returns DataAwsOpensearchDomainClusterConfigList
	_jsii_.Get(
		j,
		"clusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) CognitoOptions() DataAwsOpensearchDomainCognitoOptionsList {
	var returns DataAwsOpensearchDomainCognitoOptionsList
	_jsii_.Get(
		j,
		"cognitoOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) Created() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) DashboardEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) DashboardEndpointV2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardEndpointV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) Deleted() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"deleted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) DomainEndpointV2HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainEndpointV2HostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) EbsOptions() DataAwsOpensearchDomainEbsOptionsList {
	var returns DataAwsOpensearchDomainEbsOptionsList
	_jsii_.Get(
		j,
		"ebsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) EncryptionAtRest() DataAwsOpensearchDomainEncryptionAtRestList {
	var returns DataAwsOpensearchDomainEncryptionAtRestList
	_jsii_.Get(
		j,
		"encryptionAtRest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) EndpointV2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) KibanaEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kibanaEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) LogPublishingOptions() DataAwsOpensearchDomainLogPublishingOptionsList {
	var returns DataAwsOpensearchDomainLogPublishingOptionsList
	_jsii_.Get(
		j,
		"logPublishingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) NodeToNodeEncryption() DataAwsOpensearchDomainNodeToNodeEncryptionList {
	var returns DataAwsOpensearchDomainNodeToNodeEncryptionList
	_jsii_.Get(
		j,
		"nodeToNodeEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) OffPeakWindowOptions() DataAwsOpensearchDomainOffPeakWindowOptionsOutputReference {
	var returns DataAwsOpensearchDomainOffPeakWindowOptionsOutputReference
	_jsii_.Get(
		j,
		"offPeakWindowOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) OffPeakWindowOptionsInput() *DataAwsOpensearchDomainOffPeakWindowOptions {
	var returns *DataAwsOpensearchDomainOffPeakWindowOptions
	_jsii_.Get(
		j,
		"offPeakWindowOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) Processing() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"processing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) SnapshotOptions() DataAwsOpensearchDomainSnapshotOptionsList {
	var returns DataAwsOpensearchDomainSnapshotOptionsList
	_jsii_.Get(
		j,
		"snapshotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) SoftwareUpdateOptions() DataAwsOpensearchDomainSoftwareUpdateOptionsList {
	var returns DataAwsOpensearchDomainSoftwareUpdateOptionsList
	_jsii_.Get(
		j,
		"softwareUpdateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOpensearchDomain) VpcOptions() DataAwsOpensearchDomainVpcOptionsList {
	var returns DataAwsOpensearchDomainVpcOptionsList
	_jsii_.Get(
		j,
		"vpcOptions",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/data-sources/opensearch_domain aws_opensearch_domain} Data Source.
func NewDataAwsOpensearchDomain(scope constructs.Construct, id *string, config *DataAwsOpensearchDomainConfig) DataAwsOpensearchDomain {
	_init_.Initialize()

	if err := validateNewDataAwsOpensearchDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsOpensearchDomain{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsOpensearchDomain.DataAwsOpensearchDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/data-sources/opensearch_domain aws_opensearch_domain} Data Source.
func NewDataAwsOpensearchDomain_Override(d DataAwsOpensearchDomain, scope constructs.Construct, id *string, config *DataAwsOpensearchDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsOpensearchDomain.DataAwsOpensearchDomain",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsOpensearchDomain)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsOpensearchDomain)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsOpensearchDomain)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_DataAwsOpensearchDomain)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsOpensearchDomain)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsOpensearchDomain)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsOpensearchDomain)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsOpensearchDomain)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsOpensearchDomain resource upon running "cdktf plan <stack-name>".
func DataAwsOpensearchDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsOpensearchDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsOpensearchDomain.DataAwsOpensearchDomain",
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
func DataAwsOpensearchDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOpensearchDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsOpensearchDomain.DataAwsOpensearchDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsOpensearchDomain_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOpensearchDomain_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsOpensearchDomain.DataAwsOpensearchDomain",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsOpensearchDomain_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOpensearchDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsOpensearchDomain.DataAwsOpensearchDomain",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsOpensearchDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsOpensearchDomain.DataAwsOpensearchDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsOpensearchDomain) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsOpensearchDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsOpensearchDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsOpensearchDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsOpensearchDomain) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsOpensearchDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsOpensearchDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsOpensearchDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsOpensearchDomain) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsOpensearchDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsOpensearchDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsOpensearchDomain) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsOpensearchDomain) PutOffPeakWindowOptions(value *DataAwsOpensearchDomainOffPeakWindowOptions) {
	if err := d.validatePutOffPeakWindowOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOffPeakWindowOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsOpensearchDomain) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOpensearchDomain) ResetOffPeakWindowOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetOffPeakWindowOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOpensearchDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOpensearchDomain) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOpensearchDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOpensearchDomain) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOpensearchDomain) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOpensearchDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOpensearchDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOpensearchDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

