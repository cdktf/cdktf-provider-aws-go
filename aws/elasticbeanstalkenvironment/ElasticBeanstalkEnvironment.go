package elasticbeanstalkenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/elasticbeanstalkenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/elastic_beanstalk_environment aws_elastic_beanstalk_environment}.
type ElasticBeanstalkEnvironment interface {
	cdktf.TerraformResource
	AllSettings() ElasticBeanstalkEnvironmentAllSettingsList
	Application() *string
	SetApplication(val *string)
	ApplicationInput() *string
	Arn() *string
	AutoscalingGroups() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cname() *string
	CnamePrefix() *string
	SetCnamePrefix(val *string)
	CnamePrefixInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EndpointUrl() *string
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
	Instances() *[]*string
	LaunchConfigurations() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancers() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PlatformArn() *string
	SetPlatformArn(val *string)
	PlatformArnInput() *string
	PollInterval() *string
	SetPollInterval(val *string)
	PollIntervalInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Queues() *[]*string
	// Experimental.
	RawOverrides() interface{}
	Setting() ElasticBeanstalkEnvironmentSettingList
	SettingInput() interface{}
	SolutionStackName() *string
	SetSolutionStackName(val *string)
	SolutionStackNameInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TemplateName() *string
	SetTemplateName(val *string)
	TemplateNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	Triggers() *[]*string
	VersionLabel() *string
	SetVersionLabel(val *string)
	VersionLabelInput() *string
	WaitForReadyTimeout() *string
	SetWaitForReadyTimeout(val *string)
	WaitForReadyTimeoutInput() *string
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
	PutSetting(value interface{})
	ResetCnamePrefix()
	ResetDescription()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatformArn()
	ResetPollInterval()
	ResetSetting()
	ResetSolutionStackName()
	ResetTags()
	ResetTagsAll()
	ResetTemplateName()
	ResetTier()
	ResetVersionLabel()
	ResetWaitForReadyTimeout()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ElasticBeanstalkEnvironment
type jsiiProxy_ElasticBeanstalkEnvironment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) AllSettings() ElasticBeanstalkEnvironmentAllSettingsList {
	var returns ElasticBeanstalkEnvironmentAllSettingsList
	_jsii_.Get(
		j,
		"allSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Application() *string {
	var returns *string
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) ApplicationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) AutoscalingGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoscalingGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Cname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) CnamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cnamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) CnamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cnamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) EndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Instances() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) LaunchConfigurations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"launchConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) LoadBalancers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) PlatformArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) PlatformArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) PollInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) PollIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Queues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Setting() ElasticBeanstalkEnvironmentSettingList {
	var returns ElasticBeanstalkEnvironmentSettingList
	_jsii_.Get(
		j,
		"setting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) SettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"settingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) SolutionStackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"solutionStackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) SolutionStackNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"solutionStackNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) TemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) TemplateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) Triggers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) VersionLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) VersionLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) WaitForReadyTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForReadyTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment) WaitForReadyTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForReadyTimeoutInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/elastic_beanstalk_environment aws_elastic_beanstalk_environment} Resource.
func NewElasticBeanstalkEnvironment(scope constructs.Construct, id *string, config *ElasticBeanstalkEnvironmentConfig) ElasticBeanstalkEnvironment {
	_init_.Initialize()

	if err := validateNewElasticBeanstalkEnvironmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticBeanstalkEnvironment{}

	_jsii_.Create(
		"@cdktf/provider-aws.elasticBeanstalkEnvironment.ElasticBeanstalkEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/elastic_beanstalk_environment aws_elastic_beanstalk_environment} Resource.
func NewElasticBeanstalkEnvironment_Override(e ElasticBeanstalkEnvironment, scope constructs.Construct, id *string, config *ElasticBeanstalkEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.elasticBeanstalkEnvironment.ElasticBeanstalkEnvironment",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetApplication(val *string) {
	if err := j.validateSetApplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"application",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetCnamePrefix(val *string) {
	if err := j.validateSetCnamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cnamePrefix",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetPlatformArn(val *string) {
	if err := j.validateSetPlatformArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformArn",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetPollInterval(val *string) {
	if err := j.validateSetPollIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pollInterval",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetSolutionStackName(val *string) {
	if err := j.validateSetSolutionStackNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"solutionStackName",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetTemplateName(val *string) {
	if err := j.validateSetTemplateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateName",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetVersionLabel(val *string) {
	if err := j.validateSetVersionLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionLabel",
		val,
	)
}

func (j *jsiiProxy_ElasticBeanstalkEnvironment)SetWaitForReadyTimeout(val *string) {
	if err := j.validateSetWaitForReadyTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForReadyTimeout",
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
func ElasticBeanstalkEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticBeanstalkEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elasticBeanstalkEnvironment.ElasticBeanstalkEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticBeanstalkEnvironment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticBeanstalkEnvironment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elasticBeanstalkEnvironment.ElasticBeanstalkEnvironment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticBeanstalkEnvironment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticBeanstalkEnvironment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elasticBeanstalkEnvironment.ElasticBeanstalkEnvironment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElasticBeanstalkEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.elasticBeanstalkEnvironment.ElasticBeanstalkEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElasticBeanstalkEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElasticBeanstalkEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElasticBeanstalkEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElasticBeanstalkEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElasticBeanstalkEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElasticBeanstalkEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElasticBeanstalkEnvironment) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElasticBeanstalkEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElasticBeanstalkEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElasticBeanstalkEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) PutSetting(value interface{}) {
	if err := e.validatePutSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSetting",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetCnamePrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetCnamePrefix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetPlatformArn() {
	_jsii_.InvokeVoid(
		e,
		"resetPlatformArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetPollInterval() {
	_jsii_.InvokeVoid(
		e,
		"resetPollInterval",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetSetting() {
	_jsii_.InvokeVoid(
		e,
		"resetSetting",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetSolutionStackName() {
	_jsii_.InvokeVoid(
		e,
		"resetSolutionStackName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetTemplateName() {
	_jsii_.InvokeVoid(
		e,
		"resetTemplateName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetTier() {
	_jsii_.InvokeVoid(
		e,
		"resetTier",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetVersionLabel() {
	_jsii_.InvokeVoid(
		e,
		"resetVersionLabel",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ResetWaitForReadyTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetWaitForReadyTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticBeanstalkEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

