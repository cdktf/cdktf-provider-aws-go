package gluecrawler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/gluecrawler/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler aws_glue_crawler}.
type GlueCrawler interface {
	cdktf.TerraformResource
	Arn() *string
	CatalogTarget() GlueCrawlerCatalogTargetList
	CatalogTargetInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Classifiers() *[]*string
	SetClassifiers(val *[]*string)
	ClassifiersInput() *[]*string
	Configuration() *string
	SetConfiguration(val *string)
	ConfigurationInput() *string
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DeltaTarget() GlueCrawlerDeltaTargetList
	DeltaTargetInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DynamodbTarget() GlueCrawlerDynamodbTargetList
	DynamodbTargetInput() interface{}
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
	JdbcTarget() GlueCrawlerJdbcTargetList
	JdbcTargetInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LineageConfiguration() GlueCrawlerLineageConfigurationOutputReference
	LineageConfigurationInput() *GlueCrawlerLineageConfiguration
	MongodbTarget() GlueCrawlerMongodbTargetList
	MongodbTargetInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	RecrawlPolicy() GlueCrawlerRecrawlPolicyOutputReference
	RecrawlPolicyInput() *GlueCrawlerRecrawlPolicy
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	S3Target() GlueCrawlerS3TargetList
	S3TargetInput() interface{}
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
	SchemaChangePolicy() GlueCrawlerSchemaChangePolicyOutputReference
	SchemaChangePolicyInput() *GlueCrawlerSchemaChangePolicy
	SecurityConfiguration() *string
	SetSecurityConfiguration(val *string)
	SecurityConfigurationInput() *string
	TablePrefix() *string
	SetTablePrefix(val *string)
	TablePrefixInput() *string
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
	PutCatalogTarget(value interface{})
	PutDeltaTarget(value interface{})
	PutDynamodbTarget(value interface{})
	PutJdbcTarget(value interface{})
	PutLineageConfiguration(value *GlueCrawlerLineageConfiguration)
	PutMongodbTarget(value interface{})
	PutRecrawlPolicy(value *GlueCrawlerRecrawlPolicy)
	PutS3Target(value interface{})
	PutSchemaChangePolicy(value *GlueCrawlerSchemaChangePolicy)
	ResetCatalogTarget()
	ResetClassifiers()
	ResetConfiguration()
	ResetDeltaTarget()
	ResetDescription()
	ResetDynamodbTarget()
	ResetId()
	ResetJdbcTarget()
	ResetLineageConfiguration()
	ResetMongodbTarget()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRecrawlPolicy()
	ResetS3Target()
	ResetSchedule()
	ResetSchemaChangePolicy()
	ResetSecurityConfiguration()
	ResetTablePrefix()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueCrawler
type jsiiProxy_GlueCrawler struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueCrawler) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) CatalogTarget() GlueCrawlerCatalogTargetList {
	var returns GlueCrawlerCatalogTargetList
	_jsii_.Get(
		j,
		"catalogTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) CatalogTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Classifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) ClassifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Configuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) ConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DeltaTarget() GlueCrawlerDeltaTargetList {
	var returns GlueCrawlerDeltaTargetList
	_jsii_.Get(
		j,
		"deltaTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DeltaTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deltaTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DynamodbTarget() GlueCrawlerDynamodbTargetList {
	var returns GlueCrawlerDynamodbTargetList
	_jsii_.Get(
		j,
		"dynamodbTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DynamodbTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamodbTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) JdbcTarget() GlueCrawlerJdbcTargetList {
	var returns GlueCrawlerJdbcTargetList
	_jsii_.Get(
		j,
		"jdbcTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) JdbcTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) LineageConfiguration() GlueCrawlerLineageConfigurationOutputReference {
	var returns GlueCrawlerLineageConfigurationOutputReference
	_jsii_.Get(
		j,
		"lineageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) LineageConfigurationInput() *GlueCrawlerLineageConfiguration {
	var returns *GlueCrawlerLineageConfiguration
	_jsii_.Get(
		j,
		"lineageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) MongodbTarget() GlueCrawlerMongodbTargetList {
	var returns GlueCrawlerMongodbTargetList
	_jsii_.Get(
		j,
		"mongodbTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) MongodbTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongodbTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) RecrawlPolicy() GlueCrawlerRecrawlPolicyOutputReference {
	var returns GlueCrawlerRecrawlPolicyOutputReference
	_jsii_.Get(
		j,
		"recrawlPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) RecrawlPolicyInput() *GlueCrawlerRecrawlPolicy {
	var returns *GlueCrawlerRecrawlPolicy
	_jsii_.Get(
		j,
		"recrawlPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) S3Target() GlueCrawlerS3TargetList {
	var returns GlueCrawlerS3TargetList
	_jsii_.Get(
		j,
		"s3Target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) S3TargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3TargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) SchemaChangePolicy() GlueCrawlerSchemaChangePolicyOutputReference {
	var returns GlueCrawlerSchemaChangePolicyOutputReference
	_jsii_.Get(
		j,
		"schemaChangePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) SchemaChangePolicyInput() *GlueCrawlerSchemaChangePolicy {
	var returns *GlueCrawlerSchemaChangePolicy
	_jsii_.Get(
		j,
		"schemaChangePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TablePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TablePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler aws_glue_crawler} Resource.
func NewGlueCrawler(scope constructs.Construct, id *string, config *GlueCrawlerConfig) GlueCrawler {
	_init_.Initialize()

	if err := validateNewGlueCrawlerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCrawler{}

	_jsii_.Create(
		"@cdktf/provider-aws.glueCrawler.GlueCrawler",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler aws_glue_crawler} Resource.
func NewGlueCrawler_Override(g GlueCrawler, scope constructs.Construct, id *string, config *GlueCrawlerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.glueCrawler.GlueCrawler",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueCrawler)SetClassifiers(val *[]*string) {
	if err := j.validateSetClassifiersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classifiers",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetConfiguration(val *string) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetSecurityConfiguration(val *string) {
	if err := j.validateSetSecurityConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetTablePrefix(val *string) {
	if err := j.validateSetTablePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tablePrefix",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler)SetTagsAll(val *map[string]*string) {
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
func GlueCrawler_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlueCrawler_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.glueCrawler.GlueCrawler",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GlueCrawler_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlueCrawler_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.glueCrawler.GlueCrawler",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GlueCrawler_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlueCrawler_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.glueCrawler.GlueCrawler",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueCrawler_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.glueCrawler.GlueCrawler",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GlueCrawler) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GlueCrawler) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawler) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawler) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawler) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawler) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawler) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawler) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawler) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawler) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawler) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawler) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueCrawler) PutCatalogTarget(value interface{}) {
	if err := g.validatePutCatalogTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCatalogTarget",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawler) PutDeltaTarget(value interface{}) {
	if err := g.validatePutDeltaTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeltaTarget",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawler) PutDynamodbTarget(value interface{}) {
	if err := g.validatePutDynamodbTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDynamodbTarget",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawler) PutJdbcTarget(value interface{}) {
	if err := g.validatePutJdbcTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJdbcTarget",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawler) PutLineageConfiguration(value *GlueCrawlerLineageConfiguration) {
	if err := g.validatePutLineageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLineageConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawler) PutMongodbTarget(value interface{}) {
	if err := g.validatePutMongodbTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMongodbTarget",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawler) PutRecrawlPolicy(value *GlueCrawlerRecrawlPolicy) {
	if err := g.validatePutRecrawlPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRecrawlPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawler) PutS3Target(value interface{}) {
	if err := g.validatePutS3TargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putS3Target",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawler) PutSchemaChangePolicy(value *GlueCrawlerSchemaChangePolicy) {
	if err := g.validatePutSchemaChangePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSchemaChangePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawler) ResetCatalogTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalogTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetClassifiers() {
	_jsii_.InvokeVoid(
		g,
		"resetClassifiers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetDeltaTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetDeltaTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetDynamodbTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetDynamodbTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetJdbcTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetJdbcTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetLineageConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetLineageConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetMongodbTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetMongodbTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetRecrawlPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRecrawlPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetS3Target() {
	_jsii_.InvokeVoid(
		g,
		"resetS3Target",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetSchemaChangePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaChangePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetTablePrefix() {
	_jsii_.InvokeVoid(
		g,
		"resetTablePrefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetTagsAll() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawler) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawler) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawler) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

