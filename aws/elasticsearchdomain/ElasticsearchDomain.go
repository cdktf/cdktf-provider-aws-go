// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/elasticsearchdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/elasticsearch_domain aws_elasticsearch_domain}.
type ElasticsearchDomain interface {
	cdktf.TerraformResource
	AccessPolicies() *string
	SetAccessPolicies(val *string)
	AccessPoliciesInput() *string
	AdvancedOptions() *map[string]*string
	SetAdvancedOptions(val *map[string]*string)
	AdvancedOptionsInput() *map[string]*string
	AdvancedSecurityOptions() ElasticsearchDomainAdvancedSecurityOptionsOutputReference
	AdvancedSecurityOptionsInput() *ElasticsearchDomainAdvancedSecurityOptions
	Arn() *string
	AutoTuneOptions() ElasticsearchDomainAutoTuneOptionsOutputReference
	AutoTuneOptionsInput() *ElasticsearchDomainAutoTuneOptions
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterConfig() ElasticsearchDomainClusterConfigOutputReference
	ClusterConfigInput() *ElasticsearchDomainClusterConfig
	CognitoOptions() ElasticsearchDomainCognitoOptionsOutputReference
	CognitoOptionsInput() *ElasticsearchDomainCognitoOptions
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
	DomainEndpointOptions() ElasticsearchDomainDomainEndpointOptionsOutputReference
	DomainEndpointOptionsInput() *ElasticsearchDomainDomainEndpointOptions
	DomainId() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	EbsOptions() ElasticsearchDomainEbsOptionsOutputReference
	EbsOptionsInput() *ElasticsearchDomainEbsOptions
	ElasticsearchVersion() *string
	SetElasticsearchVersion(val *string)
	ElasticsearchVersionInput() *string
	EncryptAtRest() ElasticsearchDomainEncryptAtRestOutputReference
	EncryptAtRestInput() *ElasticsearchDomainEncryptAtRest
	Endpoint() *string
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
	KibanaEndpoint() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogPublishingOptions() ElasticsearchDomainLogPublishingOptionsList
	LogPublishingOptionsInput() interface{}
	// The tree node.
	Node() constructs.Node
	NodeToNodeEncryption() ElasticsearchDomainNodeToNodeEncryptionOutputReference
	NodeToNodeEncryptionInput() *ElasticsearchDomainNodeToNodeEncryption
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
	SnapshotOptions() ElasticsearchDomainSnapshotOptionsOutputReference
	SnapshotOptionsInput() *ElasticsearchDomainSnapshotOptions
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
	Timeouts() ElasticsearchDomainTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcOptions() ElasticsearchDomainVpcOptionsOutputReference
	VpcOptionsInput() *ElasticsearchDomainVpcOptions
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
	PutAdvancedSecurityOptions(value *ElasticsearchDomainAdvancedSecurityOptions)
	PutAutoTuneOptions(value *ElasticsearchDomainAutoTuneOptions)
	PutClusterConfig(value *ElasticsearchDomainClusterConfig)
	PutCognitoOptions(value *ElasticsearchDomainCognitoOptions)
	PutDomainEndpointOptions(value *ElasticsearchDomainDomainEndpointOptions)
	PutEbsOptions(value *ElasticsearchDomainEbsOptions)
	PutEncryptAtRest(value *ElasticsearchDomainEncryptAtRest)
	PutLogPublishingOptions(value interface{})
	PutNodeToNodeEncryption(value *ElasticsearchDomainNodeToNodeEncryption)
	PutSnapshotOptions(value *ElasticsearchDomainSnapshotOptions)
	PutTimeouts(value *ElasticsearchDomainTimeouts)
	PutVpcOptions(value *ElasticsearchDomainVpcOptions)
	ResetAccessPolicies()
	ResetAdvancedOptions()
	ResetAdvancedSecurityOptions()
	ResetAutoTuneOptions()
	ResetClusterConfig()
	ResetCognitoOptions()
	ResetDomainEndpointOptions()
	ResetEbsOptions()
	ResetElasticsearchVersion()
	ResetEncryptAtRest()
	ResetId()
	ResetLogPublishingOptions()
	ResetNodeToNodeEncryption()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSnapshotOptions()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetVpcOptions()
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

// The jsii proxy struct for ElasticsearchDomain
type jsiiProxy_ElasticsearchDomain struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElasticsearchDomain) AccessPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AccessPoliciesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AdvancedOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"advancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AdvancedOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"advancedOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AdvancedSecurityOptions() ElasticsearchDomainAdvancedSecurityOptionsOutputReference {
	var returns ElasticsearchDomainAdvancedSecurityOptionsOutputReference
	_jsii_.Get(
		j,
		"advancedSecurityOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AdvancedSecurityOptionsInput() *ElasticsearchDomainAdvancedSecurityOptions {
	var returns *ElasticsearchDomainAdvancedSecurityOptions
	_jsii_.Get(
		j,
		"advancedSecurityOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AutoTuneOptions() ElasticsearchDomainAutoTuneOptionsOutputReference {
	var returns ElasticsearchDomainAutoTuneOptionsOutputReference
	_jsii_.Get(
		j,
		"autoTuneOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AutoTuneOptionsInput() *ElasticsearchDomainAutoTuneOptions {
	var returns *ElasticsearchDomainAutoTuneOptions
	_jsii_.Get(
		j,
		"autoTuneOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) ClusterConfig() ElasticsearchDomainClusterConfigOutputReference {
	var returns ElasticsearchDomainClusterConfigOutputReference
	_jsii_.Get(
		j,
		"clusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) ClusterConfigInput() *ElasticsearchDomainClusterConfig {
	var returns *ElasticsearchDomainClusterConfig
	_jsii_.Get(
		j,
		"clusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) CognitoOptions() ElasticsearchDomainCognitoOptionsOutputReference {
	var returns ElasticsearchDomainCognitoOptionsOutputReference
	_jsii_.Get(
		j,
		"cognitoOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) CognitoOptionsInput() *ElasticsearchDomainCognitoOptions {
	var returns *ElasticsearchDomainCognitoOptions
	_jsii_.Get(
		j,
		"cognitoOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) DomainEndpointOptions() ElasticsearchDomainDomainEndpointOptionsOutputReference {
	var returns ElasticsearchDomainDomainEndpointOptionsOutputReference
	_jsii_.Get(
		j,
		"domainEndpointOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) DomainEndpointOptionsInput() *ElasticsearchDomainDomainEndpointOptions {
	var returns *ElasticsearchDomainDomainEndpointOptions
	_jsii_.Get(
		j,
		"domainEndpointOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) EbsOptions() ElasticsearchDomainEbsOptionsOutputReference {
	var returns ElasticsearchDomainEbsOptionsOutputReference
	_jsii_.Get(
		j,
		"ebsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) EbsOptionsInput() *ElasticsearchDomainEbsOptions {
	var returns *ElasticsearchDomainEbsOptions
	_jsii_.Get(
		j,
		"ebsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) ElasticsearchVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) ElasticsearchVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) EncryptAtRest() ElasticsearchDomainEncryptAtRestOutputReference {
	var returns ElasticsearchDomainEncryptAtRestOutputReference
	_jsii_.Get(
		j,
		"encryptAtRest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) EncryptAtRestInput() *ElasticsearchDomainEncryptAtRest {
	var returns *ElasticsearchDomainEncryptAtRest
	_jsii_.Get(
		j,
		"encryptAtRestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) KibanaEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kibanaEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) LogPublishingOptions() ElasticsearchDomainLogPublishingOptionsList {
	var returns ElasticsearchDomainLogPublishingOptionsList
	_jsii_.Get(
		j,
		"logPublishingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) LogPublishingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPublishingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) NodeToNodeEncryption() ElasticsearchDomainNodeToNodeEncryptionOutputReference {
	var returns ElasticsearchDomainNodeToNodeEncryptionOutputReference
	_jsii_.Get(
		j,
		"nodeToNodeEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) NodeToNodeEncryptionInput() *ElasticsearchDomainNodeToNodeEncryption {
	var returns *ElasticsearchDomainNodeToNodeEncryption
	_jsii_.Get(
		j,
		"nodeToNodeEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) SnapshotOptions() ElasticsearchDomainSnapshotOptionsOutputReference {
	var returns ElasticsearchDomainSnapshotOptionsOutputReference
	_jsii_.Get(
		j,
		"snapshotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) SnapshotOptionsInput() *ElasticsearchDomainSnapshotOptions {
	var returns *ElasticsearchDomainSnapshotOptions
	_jsii_.Get(
		j,
		"snapshotOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Timeouts() ElasticsearchDomainTimeoutsOutputReference {
	var returns ElasticsearchDomainTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) VpcOptions() ElasticsearchDomainVpcOptionsOutputReference {
	var returns ElasticsearchDomainVpcOptionsOutputReference
	_jsii_.Get(
		j,
		"vpcOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) VpcOptionsInput() *ElasticsearchDomainVpcOptions {
	var returns *ElasticsearchDomainVpcOptions
	_jsii_.Get(
		j,
		"vpcOptionsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/elasticsearch_domain aws_elasticsearch_domain} Resource.
func NewElasticsearchDomain(scope constructs.Construct, id *string, config *ElasticsearchDomainConfig) ElasticsearchDomain {
	_init_.Initialize()

	if err := validateNewElasticsearchDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticsearchDomain{}

	_jsii_.Create(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/elasticsearch_domain aws_elasticsearch_domain} Resource.
func NewElasticsearchDomain_Override(e ElasticsearchDomain, scope constructs.Construct, id *string, config *ElasticsearchDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomain",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetAccessPolicies(val *string) {
	if err := j.validateSetAccessPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessPolicies",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetAdvancedOptions(val *map[string]*string) {
	if err := j.validateSetAdvancedOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedOptions",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetElasticsearchVersion(val *string) {
	if err := j.validateSetElasticsearchVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchVersion",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a ElasticsearchDomain resource upon running "cdktf plan <stack-name>".
func ElasticsearchDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateElasticsearchDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomain",
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
func ElasticsearchDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticsearchDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticsearchDomain_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticsearchDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomain",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticsearchDomain_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticsearchDomain_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomain",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElasticsearchDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutAdvancedSecurityOptions(value *ElasticsearchDomainAdvancedSecurityOptions) {
	if err := e.validatePutAdvancedSecurityOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAdvancedSecurityOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutAutoTuneOptions(value *ElasticsearchDomainAutoTuneOptions) {
	if err := e.validatePutAutoTuneOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutoTuneOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutClusterConfig(value *ElasticsearchDomainClusterConfig) {
	if err := e.validatePutClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putClusterConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutCognitoOptions(value *ElasticsearchDomainCognitoOptions) {
	if err := e.validatePutCognitoOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCognitoOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutDomainEndpointOptions(value *ElasticsearchDomainDomainEndpointOptions) {
	if err := e.validatePutDomainEndpointOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDomainEndpointOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutEbsOptions(value *ElasticsearchDomainEbsOptions) {
	if err := e.validatePutEbsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEbsOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutEncryptAtRest(value *ElasticsearchDomainEncryptAtRest) {
	if err := e.validatePutEncryptAtRestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEncryptAtRest",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutLogPublishingOptions(value interface{}) {
	if err := e.validatePutLogPublishingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLogPublishingOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutNodeToNodeEncryption(value *ElasticsearchDomainNodeToNodeEncryption) {
	if err := e.validatePutNodeToNodeEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNodeToNodeEncryption",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutSnapshotOptions(value *ElasticsearchDomainSnapshotOptions) {
	if err := e.validatePutSnapshotOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSnapshotOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutTimeouts(value *ElasticsearchDomainTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutVpcOptions(value *ElasticsearchDomainVpcOptions) {
	if err := e.validatePutVpcOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVpcOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetAccessPolicies() {
	_jsii_.InvokeVoid(
		e,
		"resetAccessPolicies",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetAdvancedOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetAdvancedOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetAdvancedSecurityOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetAdvancedSecurityOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetAutoTuneOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoTuneOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetClusterConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetClusterConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetCognitoOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetCognitoOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetDomainEndpointOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetDomainEndpointOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetEbsOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetElasticsearchVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetElasticsearchVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetEncryptAtRest() {
	_jsii_.InvokeVoid(
		e,
		"resetEncryptAtRest",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetLogPublishingOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetLogPublishingOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetNodeToNodeEncryption() {
	_jsii_.InvokeVoid(
		e,
		"resetNodeToNodeEncryption",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetSnapshotOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetVpcOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetVpcOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

