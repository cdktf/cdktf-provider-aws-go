package opensearchdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/opensearchdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain aws_opensearch_domain}.
type OpensearchDomain interface {
	cdktf.TerraformResource
	AccessPolicies() *string
	SetAccessPolicies(val *string)
	AccessPoliciesInput() *string
	AdvancedOptions() *map[string]*string
	SetAdvancedOptions(val *map[string]*string)
	AdvancedOptionsInput() *map[string]*string
	AdvancedSecurityOptions() OpensearchDomainAdvancedSecurityOptionsOutputReference
	AdvancedSecurityOptionsInput() *OpensearchDomainAdvancedSecurityOptions
	Arn() *string
	AutoTuneOptions() OpensearchDomainAutoTuneOptionsOutputReference
	AutoTuneOptionsInput() *OpensearchDomainAutoTuneOptions
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterConfig() OpensearchDomainClusterConfigOutputReference
	ClusterConfigInput() *OpensearchDomainClusterConfig
	CognitoOptions() OpensearchDomainCognitoOptionsOutputReference
	CognitoOptionsInput() *OpensearchDomainCognitoOptions
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DashboardEndpoint() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DomainEndpointOptions() OpensearchDomainDomainEndpointOptionsOutputReference
	DomainEndpointOptionsInput() *OpensearchDomainDomainEndpointOptions
	DomainId() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	EbsOptions() OpensearchDomainEbsOptionsOutputReference
	EbsOptionsInput() *OpensearchDomainEbsOptions
	EncryptAtRest() OpensearchDomainEncryptAtRestOutputReference
	EncryptAtRestInput() *OpensearchDomainEncryptAtRest
	Endpoint() *string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
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
	LogPublishingOptions() OpensearchDomainLogPublishingOptionsList
	LogPublishingOptionsInput() interface{}
	// The tree node.
	Node() constructs.Node
	NodeToNodeEncryption() OpensearchDomainNodeToNodeEncryptionOutputReference
	NodeToNodeEncryptionInput() *OpensearchDomainNodeToNodeEncryption
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
	SnapshotOptions() OpensearchDomainSnapshotOptionsOutputReference
	SnapshotOptionsInput() *OpensearchDomainSnapshotOptions
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
	Timeouts() OpensearchDomainTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcOptions() OpensearchDomainVpcOptionsOutputReference
	VpcOptionsInput() *OpensearchDomainVpcOptions
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
	PutAdvancedSecurityOptions(value *OpensearchDomainAdvancedSecurityOptions)
	PutAutoTuneOptions(value *OpensearchDomainAutoTuneOptions)
	PutClusterConfig(value *OpensearchDomainClusterConfig)
	PutCognitoOptions(value *OpensearchDomainCognitoOptions)
	PutDomainEndpointOptions(value *OpensearchDomainDomainEndpointOptions)
	PutEbsOptions(value *OpensearchDomainEbsOptions)
	PutEncryptAtRest(value *OpensearchDomainEncryptAtRest)
	PutLogPublishingOptions(value interface{})
	PutNodeToNodeEncryption(value *OpensearchDomainNodeToNodeEncryption)
	PutSnapshotOptions(value *OpensearchDomainSnapshotOptions)
	PutTimeouts(value *OpensearchDomainTimeouts)
	PutVpcOptions(value *OpensearchDomainVpcOptions)
	ResetAccessPolicies()
	ResetAdvancedOptions()
	ResetAdvancedSecurityOptions()
	ResetAutoTuneOptions()
	ResetClusterConfig()
	ResetCognitoOptions()
	ResetDomainEndpointOptions()
	ResetEbsOptions()
	ResetEncryptAtRest()
	ResetEngineVersion()
	ResetId()
	ResetLogPublishingOptions()
	ResetNodeToNodeEncryption()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSnapshotOptions()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetVpcOptions()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OpensearchDomain
type jsiiProxy_OpensearchDomain struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpensearchDomain) AccessPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) AccessPoliciesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) AdvancedOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"advancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) AdvancedOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"advancedOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) AdvancedSecurityOptions() OpensearchDomainAdvancedSecurityOptionsOutputReference {
	var returns OpensearchDomainAdvancedSecurityOptionsOutputReference
	_jsii_.Get(
		j,
		"advancedSecurityOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) AdvancedSecurityOptionsInput() *OpensearchDomainAdvancedSecurityOptions {
	var returns *OpensearchDomainAdvancedSecurityOptions
	_jsii_.Get(
		j,
		"advancedSecurityOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) AutoTuneOptions() OpensearchDomainAutoTuneOptionsOutputReference {
	var returns OpensearchDomainAutoTuneOptionsOutputReference
	_jsii_.Get(
		j,
		"autoTuneOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) AutoTuneOptionsInput() *OpensearchDomainAutoTuneOptions {
	var returns *OpensearchDomainAutoTuneOptions
	_jsii_.Get(
		j,
		"autoTuneOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) ClusterConfig() OpensearchDomainClusterConfigOutputReference {
	var returns OpensearchDomainClusterConfigOutputReference
	_jsii_.Get(
		j,
		"clusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) ClusterConfigInput() *OpensearchDomainClusterConfig {
	var returns *OpensearchDomainClusterConfig
	_jsii_.Get(
		j,
		"clusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) CognitoOptions() OpensearchDomainCognitoOptionsOutputReference {
	var returns OpensearchDomainCognitoOptionsOutputReference
	_jsii_.Get(
		j,
		"cognitoOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) CognitoOptionsInput() *OpensearchDomainCognitoOptions {
	var returns *OpensearchDomainCognitoOptions
	_jsii_.Get(
		j,
		"cognitoOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) DashboardEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) DomainEndpointOptions() OpensearchDomainDomainEndpointOptionsOutputReference {
	var returns OpensearchDomainDomainEndpointOptionsOutputReference
	_jsii_.Get(
		j,
		"domainEndpointOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) DomainEndpointOptionsInput() *OpensearchDomainDomainEndpointOptions {
	var returns *OpensearchDomainDomainEndpointOptions
	_jsii_.Get(
		j,
		"domainEndpointOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) EbsOptions() OpensearchDomainEbsOptionsOutputReference {
	var returns OpensearchDomainEbsOptionsOutputReference
	_jsii_.Get(
		j,
		"ebsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) EbsOptionsInput() *OpensearchDomainEbsOptions {
	var returns *OpensearchDomainEbsOptions
	_jsii_.Get(
		j,
		"ebsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) EncryptAtRest() OpensearchDomainEncryptAtRestOutputReference {
	var returns OpensearchDomainEncryptAtRestOutputReference
	_jsii_.Get(
		j,
		"encryptAtRest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) EncryptAtRestInput() *OpensearchDomainEncryptAtRest {
	var returns *OpensearchDomainEncryptAtRest
	_jsii_.Get(
		j,
		"encryptAtRestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) KibanaEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kibanaEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) LogPublishingOptions() OpensearchDomainLogPublishingOptionsList {
	var returns OpensearchDomainLogPublishingOptionsList
	_jsii_.Get(
		j,
		"logPublishingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) LogPublishingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPublishingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) NodeToNodeEncryption() OpensearchDomainNodeToNodeEncryptionOutputReference {
	var returns OpensearchDomainNodeToNodeEncryptionOutputReference
	_jsii_.Get(
		j,
		"nodeToNodeEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) NodeToNodeEncryptionInput() *OpensearchDomainNodeToNodeEncryption {
	var returns *OpensearchDomainNodeToNodeEncryption
	_jsii_.Get(
		j,
		"nodeToNodeEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) SnapshotOptions() OpensearchDomainSnapshotOptionsOutputReference {
	var returns OpensearchDomainSnapshotOptionsOutputReference
	_jsii_.Get(
		j,
		"snapshotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) SnapshotOptionsInput() *OpensearchDomainSnapshotOptions {
	var returns *OpensearchDomainSnapshotOptions
	_jsii_.Get(
		j,
		"snapshotOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) Timeouts() OpensearchDomainTimeoutsOutputReference {
	var returns OpensearchDomainTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) VpcOptions() OpensearchDomainVpcOptionsOutputReference {
	var returns OpensearchDomainVpcOptionsOutputReference
	_jsii_.Get(
		j,
		"vpcOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomain) VpcOptionsInput() *OpensearchDomainVpcOptions {
	var returns *OpensearchDomainVpcOptions
	_jsii_.Get(
		j,
		"vpcOptionsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain aws_opensearch_domain} Resource.
func NewOpensearchDomain(scope constructs.Construct, id *string, config *OpensearchDomainConfig) OpensearchDomain {
	_init_.Initialize()

	if err := validateNewOpensearchDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpensearchDomain{}

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain aws_opensearch_domain} Resource.
func NewOpensearchDomain_Override(o OpensearchDomain, scope constructs.Construct, id *string, config *OpensearchDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomain",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetAccessPolicies(val *string) {
	if err := j.validateSetAccessPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessPolicies",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetAdvancedOptions(val *map[string]*string) {
	if err := j.validateSetAdvancedOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedOptions",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomain)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
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
func OpensearchDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpensearchDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpensearchDomain_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpensearchDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomain",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpensearchDomain_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpensearchDomain_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomain",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpensearchDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OpensearchDomain) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OpensearchDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomain) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomain) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomain) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpensearchDomain) PutAdvancedSecurityOptions(value *OpensearchDomainAdvancedSecurityOptions) {
	if err := o.validatePutAdvancedSecurityOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAdvancedSecurityOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomain) PutAutoTuneOptions(value *OpensearchDomainAutoTuneOptions) {
	if err := o.validatePutAutoTuneOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoTuneOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomain) PutClusterConfig(value *OpensearchDomainClusterConfig) {
	if err := o.validatePutClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putClusterConfig",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomain) PutCognitoOptions(value *OpensearchDomainCognitoOptions) {
	if err := o.validatePutCognitoOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCognitoOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomain) PutDomainEndpointOptions(value *OpensearchDomainDomainEndpointOptions) {
	if err := o.validatePutDomainEndpointOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDomainEndpointOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomain) PutEbsOptions(value *OpensearchDomainEbsOptions) {
	if err := o.validatePutEbsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEbsOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomain) PutEncryptAtRest(value *OpensearchDomainEncryptAtRest) {
	if err := o.validatePutEncryptAtRestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEncryptAtRest",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomain) PutLogPublishingOptions(value interface{}) {
	if err := o.validatePutLogPublishingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLogPublishingOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomain) PutNodeToNodeEncryption(value *OpensearchDomainNodeToNodeEncryption) {
	if err := o.validatePutNodeToNodeEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNodeToNodeEncryption",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomain) PutSnapshotOptions(value *OpensearchDomainSnapshotOptions) {
	if err := o.validatePutSnapshotOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSnapshotOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomain) PutTimeouts(value *OpensearchDomainTimeouts) {
	if err := o.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomain) PutVpcOptions(value *OpensearchDomainVpcOptions) {
	if err := o.validatePutVpcOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putVpcOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetAccessPolicies() {
	_jsii_.InvokeVoid(
		o,
		"resetAccessPolicies",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetAdvancedOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetAdvancedOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetAdvancedSecurityOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetAdvancedSecurityOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetAutoTuneOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoTuneOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetClusterConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetClusterConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetCognitoOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetCognitoOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetDomainEndpointOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetDomainEndpointOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetEbsOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetEncryptAtRest() {
	_jsii_.InvokeVoid(
		o,
		"resetEncryptAtRest",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetLogPublishingOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetLogPublishingOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetNodeToNodeEncryption() {
	_jsii_.InvokeVoid(
		o,
		"resetNodeToNodeEncryption",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetSnapshotOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetSnapshotOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) ResetVpcOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetVpcOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

