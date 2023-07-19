package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/dmsendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/dms_endpoint aws_dms_endpoint}.
type DmsEndpoint interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ElasticsearchSettings() DmsEndpointElasticsearchSettingsOutputReference
	ElasticsearchSettingsInput() *DmsEndpointElasticsearchSettings
	EndpointArn() *string
	EndpointId() *string
	SetEndpointId(val *string)
	EndpointIdInput() *string
	EndpointType() *string
	SetEndpointType(val *string)
	EndpointTypeInput() *string
	EngineName() *string
	SetEngineName(val *string)
	EngineNameInput() *string
	ExtraConnectionAttributes() *string
	SetExtraConnectionAttributes(val *string)
	ExtraConnectionAttributesInput() *string
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
	KafkaSettings() DmsEndpointKafkaSettingsOutputReference
	KafkaSettingsInput() *DmsEndpointKafkaSettings
	KinesisSettings() DmsEndpointKinesisSettingsOutputReference
	KinesisSettingsInput() *DmsEndpointKinesisSettings
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MongodbSettings() DmsEndpointMongodbSettingsOutputReference
	MongodbSettingsInput() *DmsEndpointMongodbSettings
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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
	RedisSettings() DmsEndpointRedisSettingsOutputReference
	RedisSettingsInput() *DmsEndpointRedisSettings
	RedshiftSettings() DmsEndpointRedshiftSettingsOutputReference
	RedshiftSettingsInput() *DmsEndpointRedshiftSettings
	S3Settings() DmsEndpointS3SettingsOutputReference
	S3SettingsInput() *DmsEndpointS3Settings
	SecretsManagerAccessRoleArn() *string
	SetSecretsManagerAccessRoleArn(val *string)
	SecretsManagerAccessRoleArnInput() *string
	SecretsManagerArn() *string
	SetSecretsManagerArn(val *string)
	SecretsManagerArnInput() *string
	ServerName() *string
	SetServerName(val *string)
	ServerNameInput() *string
	ServiceAccessRole() *string
	SetServiceAccessRole(val *string)
	ServiceAccessRoleInput() *string
	SslMode() *string
	SetSslMode(val *string)
	SslModeInput() *string
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
	Timeouts() DmsEndpointTimeoutsOutputReference
	TimeoutsInput() interface{}
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	PutElasticsearchSettings(value *DmsEndpointElasticsearchSettings)
	PutKafkaSettings(value *DmsEndpointKafkaSettings)
	PutKinesisSettings(value *DmsEndpointKinesisSettings)
	PutMongodbSettings(value *DmsEndpointMongodbSettings)
	PutRedisSettings(value *DmsEndpointRedisSettings)
	PutRedshiftSettings(value *DmsEndpointRedshiftSettings)
	PutS3Settings(value *DmsEndpointS3Settings)
	PutTimeouts(value *DmsEndpointTimeouts)
	ResetCertificateArn()
	ResetDatabaseName()
	ResetElasticsearchSettings()
	ResetExtraConnectionAttributes()
	ResetId()
	ResetKafkaSettings()
	ResetKinesisSettings()
	ResetKmsKeyArn()
	ResetMongodbSettings()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPort()
	ResetRedisSettings()
	ResetRedshiftSettings()
	ResetS3Settings()
	ResetSecretsManagerAccessRoleArn()
	ResetSecretsManagerArn()
	ResetServerName()
	ResetServiceAccessRole()
	ResetSslMode()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetUsername()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DmsEndpoint
type jsiiProxy_DmsEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DmsEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ElasticsearchSettings() DmsEndpointElasticsearchSettingsOutputReference {
	var returns DmsEndpointElasticsearchSettingsOutputReference
	_jsii_.Get(
		j,
		"elasticsearchSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ElasticsearchSettingsInput() *DmsEndpointElasticsearchSettings {
	var returns *DmsEndpointElasticsearchSettings
	_jsii_.Get(
		j,
		"elasticsearchSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EngineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EngineNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ExtraConnectionAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraConnectionAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ExtraConnectionAttributesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraConnectionAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KafkaSettings() DmsEndpointKafkaSettingsOutputReference {
	var returns DmsEndpointKafkaSettingsOutputReference
	_jsii_.Get(
		j,
		"kafkaSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KafkaSettingsInput() *DmsEndpointKafkaSettings {
	var returns *DmsEndpointKafkaSettings
	_jsii_.Get(
		j,
		"kafkaSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KinesisSettings() DmsEndpointKinesisSettingsOutputReference {
	var returns DmsEndpointKinesisSettingsOutputReference
	_jsii_.Get(
		j,
		"kinesisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KinesisSettingsInput() *DmsEndpointKinesisSettings {
	var returns *DmsEndpointKinesisSettings
	_jsii_.Get(
		j,
		"kinesisSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) MongodbSettings() DmsEndpointMongodbSettingsOutputReference {
	var returns DmsEndpointMongodbSettingsOutputReference
	_jsii_.Get(
		j,
		"mongodbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) MongodbSettingsInput() *DmsEndpointMongodbSettings {
	var returns *DmsEndpointMongodbSettings
	_jsii_.Get(
		j,
		"mongodbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) RedisSettings() DmsEndpointRedisSettingsOutputReference {
	var returns DmsEndpointRedisSettingsOutputReference
	_jsii_.Get(
		j,
		"redisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) RedisSettingsInput() *DmsEndpointRedisSettings {
	var returns *DmsEndpointRedisSettings
	_jsii_.Get(
		j,
		"redisSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) RedshiftSettings() DmsEndpointRedshiftSettingsOutputReference {
	var returns DmsEndpointRedshiftSettingsOutputReference
	_jsii_.Get(
		j,
		"redshiftSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) RedshiftSettingsInput() *DmsEndpointRedshiftSettings {
	var returns *DmsEndpointRedshiftSettings
	_jsii_.Get(
		j,
		"redshiftSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) S3Settings() DmsEndpointS3SettingsOutputReference {
	var returns DmsEndpointS3SettingsOutputReference
	_jsii_.Get(
		j,
		"s3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) S3SettingsInput() *DmsEndpointS3Settings {
	var returns *DmsEndpointS3Settings
	_jsii_.Get(
		j,
		"s3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) SecretsManagerAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) SecretsManagerAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) SecretsManagerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) SecretsManagerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ServiceAccessRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ServiceAccessRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) SslMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) SslModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Timeouts() DmsEndpointTimeoutsOutputReference {
	var returns DmsEndpointTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/dms_endpoint aws_dms_endpoint} Resource.
func NewDmsEndpoint(scope constructs.Construct, id *string, config *DmsEndpointConfig) DmsEndpoint {
	_init_.Initialize()

	if err := validateNewDmsEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/dms_endpoint aws_dms_endpoint} Resource.
func NewDmsEndpoint_Override(d DmsEndpoint, scope constructs.Construct, id *string, config *DmsEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpoint",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetEndpointId(val *string) {
	if err := j.validateSetEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetEngineName(val *string) {
	if err := j.validateSetEngineNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetExtraConnectionAttributes(val *string) {
	if err := j.validateSetExtraConnectionAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraConnectionAttributes",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetSecretsManagerAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetSecretsManagerArn(val *string) {
	if err := j.validateSetSecretsManagerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetServerName(val *string) {
	if err := j.validateSetServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetServiceAccessRole(val *string) {
	if err := j.validateSetServiceAccessRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRole",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetSslMode(val *string) {
	if err := j.validateSetSslModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslMode",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
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
func DmsEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DmsEndpoint) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DmsEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutElasticsearchSettings(value *DmsEndpointElasticsearchSettings) {
	if err := d.validatePutElasticsearchSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putElasticsearchSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutKafkaSettings(value *DmsEndpointKafkaSettings) {
	if err := d.validatePutKafkaSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKafkaSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutKinesisSettings(value *DmsEndpointKinesisSettings) {
	if err := d.validatePutKinesisSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKinesisSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutMongodbSettings(value *DmsEndpointMongodbSettings) {
	if err := d.validatePutMongodbSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMongodbSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutRedisSettings(value *DmsEndpointRedisSettings) {
	if err := d.validatePutRedisSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRedisSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutRedshiftSettings(value *DmsEndpointRedshiftSettings) {
	if err := d.validatePutRedshiftSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRedshiftSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutS3Settings(value *DmsEndpointS3Settings) {
	if err := d.validatePutS3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putS3Settings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutTimeouts(value *DmsEndpointTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		d,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetElasticsearchSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetElasticsearchSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetExtraConnectionAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetExtraConnectionAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetKafkaSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetKafkaSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetKinesisSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetKinesisSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetMongodbSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetMongodbSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetRedisSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetRedisSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetRedshiftSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetRedshiftSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetS3Settings() {
	_jsii_.InvokeVoid(
		d,
		"resetS3Settings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetSecretsManagerAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetSecretsManagerArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetServerName() {
	_jsii_.InvokeVoid(
		d,
		"resetServerName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetServiceAccessRole() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccessRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetSslMode() {
	_jsii_.InvokeVoid(
		d,
		"resetSslMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

