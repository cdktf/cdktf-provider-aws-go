package cognitouserpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/cognitouserpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/cognito_user_pool aws_cognito_user_pool}.
type CognitoUserPool interface {
	cdktf.TerraformResource
	AccountRecoverySetting() CognitoUserPoolAccountRecoverySettingOutputReference
	AccountRecoverySettingInput() *CognitoUserPoolAccountRecoverySetting
	AdminCreateUserConfig() CognitoUserPoolAdminCreateUserConfigOutputReference
	AdminCreateUserConfigInput() *CognitoUserPoolAdminCreateUserConfig
	AliasAttributes() *[]*string
	SetAliasAttributes(val *[]*string)
	AliasAttributesInput() *[]*string
	Arn() *string
	AutoVerifiedAttributes() *[]*string
	SetAutoVerifiedAttributes(val *[]*string)
	AutoVerifiedAttributesInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreationDate() *string
	CustomDomain() *string
	DeletionProtection() *string
	SetDeletionProtection(val *string)
	DeletionProtectionInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeviceConfiguration() CognitoUserPoolDeviceConfigurationOutputReference
	DeviceConfigurationInput() *CognitoUserPoolDeviceConfiguration
	Domain() *string
	EmailConfiguration() CognitoUserPoolEmailConfigurationOutputReference
	EmailConfigurationInput() *CognitoUserPoolEmailConfiguration
	EmailVerificationMessage() *string
	SetEmailVerificationMessage(val *string)
	EmailVerificationMessageInput() *string
	EmailVerificationSubject() *string
	SetEmailVerificationSubject(val *string)
	EmailVerificationSubjectInput() *string
	Endpoint() *string
	EstimatedNumberOfUsers() *float64
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
	LambdaConfig() CognitoUserPoolLambdaConfigOutputReference
	LambdaConfigInput() *CognitoUserPoolLambdaConfig
	LastModifiedDate() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MfaConfiguration() *string
	SetMfaConfiguration(val *string)
	MfaConfigurationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PasswordPolicy() CognitoUserPoolPasswordPolicyOutputReference
	PasswordPolicyInput() *CognitoUserPoolPasswordPolicy
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
	Schema() CognitoUserPoolSchemaList
	SchemaInput() interface{}
	SmsAuthenticationMessage() *string
	SetSmsAuthenticationMessage(val *string)
	SmsAuthenticationMessageInput() *string
	SmsConfiguration() CognitoUserPoolSmsConfigurationOutputReference
	SmsConfigurationInput() *CognitoUserPoolSmsConfiguration
	SmsVerificationMessage() *string
	SetSmsVerificationMessage(val *string)
	SmsVerificationMessageInput() *string
	SoftwareTokenMfaConfiguration() CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference
	SoftwareTokenMfaConfigurationInput() *CognitoUserPoolSoftwareTokenMfaConfiguration
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
	UserAttributeUpdateSettings() CognitoUserPoolUserAttributeUpdateSettingsOutputReference
	UserAttributeUpdateSettingsInput() *CognitoUserPoolUserAttributeUpdateSettings
	UsernameAttributes() *[]*string
	SetUsernameAttributes(val *[]*string)
	UsernameAttributesInput() *[]*string
	UsernameConfiguration() CognitoUserPoolUsernameConfigurationOutputReference
	UsernameConfigurationInput() *CognitoUserPoolUsernameConfiguration
	UserPoolAddOns() CognitoUserPoolUserPoolAddOnsOutputReference
	UserPoolAddOnsInput() *CognitoUserPoolUserPoolAddOns
	VerificationMessageTemplate() CognitoUserPoolVerificationMessageTemplateOutputReference
	VerificationMessageTemplateInput() *CognitoUserPoolVerificationMessageTemplate
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
	PutAccountRecoverySetting(value *CognitoUserPoolAccountRecoverySetting)
	PutAdminCreateUserConfig(value *CognitoUserPoolAdminCreateUserConfig)
	PutDeviceConfiguration(value *CognitoUserPoolDeviceConfiguration)
	PutEmailConfiguration(value *CognitoUserPoolEmailConfiguration)
	PutLambdaConfig(value *CognitoUserPoolLambdaConfig)
	PutPasswordPolicy(value *CognitoUserPoolPasswordPolicy)
	PutSchema(value interface{})
	PutSmsConfiguration(value *CognitoUserPoolSmsConfiguration)
	PutSoftwareTokenMfaConfiguration(value *CognitoUserPoolSoftwareTokenMfaConfiguration)
	PutUserAttributeUpdateSettings(value *CognitoUserPoolUserAttributeUpdateSettings)
	PutUsernameConfiguration(value *CognitoUserPoolUsernameConfiguration)
	PutUserPoolAddOns(value *CognitoUserPoolUserPoolAddOns)
	PutVerificationMessageTemplate(value *CognitoUserPoolVerificationMessageTemplate)
	ResetAccountRecoverySetting()
	ResetAdminCreateUserConfig()
	ResetAliasAttributes()
	ResetAutoVerifiedAttributes()
	ResetDeletionProtection()
	ResetDeviceConfiguration()
	ResetEmailConfiguration()
	ResetEmailVerificationMessage()
	ResetEmailVerificationSubject()
	ResetId()
	ResetLambdaConfig()
	ResetMfaConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasswordPolicy()
	ResetSchema()
	ResetSmsAuthenticationMessage()
	ResetSmsConfiguration()
	ResetSmsVerificationMessage()
	ResetSoftwareTokenMfaConfiguration()
	ResetTags()
	ResetTagsAll()
	ResetUserAttributeUpdateSettings()
	ResetUsernameAttributes()
	ResetUsernameConfiguration()
	ResetUserPoolAddOns()
	ResetVerificationMessageTemplate()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CognitoUserPool
type jsiiProxy_CognitoUserPool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoUserPool) AccountRecoverySetting() CognitoUserPoolAccountRecoverySettingOutputReference {
	var returns CognitoUserPoolAccountRecoverySettingOutputReference
	_jsii_.Get(
		j,
		"accountRecoverySetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AccountRecoverySettingInput() *CognitoUserPoolAccountRecoverySetting {
	var returns *CognitoUserPoolAccountRecoverySetting
	_jsii_.Get(
		j,
		"accountRecoverySettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AdminCreateUserConfig() CognitoUserPoolAdminCreateUserConfigOutputReference {
	var returns CognitoUserPoolAdminCreateUserConfigOutputReference
	_jsii_.Get(
		j,
		"adminCreateUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AdminCreateUserConfigInput() *CognitoUserPoolAdminCreateUserConfig {
	var returns *CognitoUserPoolAdminCreateUserConfig
	_jsii_.Get(
		j,
		"adminCreateUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AliasAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AliasAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AutoVerifiedAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoVerifiedAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AutoVerifiedAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoVerifiedAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) CustomDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) DeletionProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) DeletionProtectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) DeviceConfiguration() CognitoUserPoolDeviceConfigurationOutputReference {
	var returns CognitoUserPoolDeviceConfigurationOutputReference
	_jsii_.Get(
		j,
		"deviceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) DeviceConfigurationInput() *CognitoUserPoolDeviceConfiguration {
	var returns *CognitoUserPoolDeviceConfiguration
	_jsii_.Get(
		j,
		"deviceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EmailConfiguration() CognitoUserPoolEmailConfigurationOutputReference {
	var returns CognitoUserPoolEmailConfigurationOutputReference
	_jsii_.Get(
		j,
		"emailConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EmailConfigurationInput() *CognitoUserPoolEmailConfiguration {
	var returns *CognitoUserPoolEmailConfiguration
	_jsii_.Get(
		j,
		"emailConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EmailVerificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EmailVerificationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EmailVerificationSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EmailVerificationSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EstimatedNumberOfUsers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"estimatedNumberOfUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) LambdaConfig() CognitoUserPoolLambdaConfigOutputReference {
	var returns CognitoUserPoolLambdaConfigOutputReference
	_jsii_.Get(
		j,
		"lambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) LambdaConfigInput() *CognitoUserPoolLambdaConfig {
	var returns *CognitoUserPoolLambdaConfig
	_jsii_.Get(
		j,
		"lambdaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) LastModifiedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) MfaConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) MfaConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) PasswordPolicy() CognitoUserPoolPasswordPolicyOutputReference {
	var returns CognitoUserPoolPasswordPolicyOutputReference
	_jsii_.Get(
		j,
		"passwordPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) PasswordPolicyInput() *CognitoUserPoolPasswordPolicy {
	var returns *CognitoUserPoolPasswordPolicy
	_jsii_.Get(
		j,
		"passwordPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Schema() CognitoUserPoolSchemaList {
	var returns CognitoUserPoolSchemaList
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SmsAuthenticationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsAuthenticationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SmsAuthenticationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsAuthenticationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SmsConfiguration() CognitoUserPoolSmsConfigurationOutputReference {
	var returns CognitoUserPoolSmsConfigurationOutputReference
	_jsii_.Get(
		j,
		"smsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SmsConfigurationInput() *CognitoUserPoolSmsConfiguration {
	var returns *CognitoUserPoolSmsConfiguration
	_jsii_.Get(
		j,
		"smsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SmsVerificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsVerificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SmsVerificationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsVerificationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SoftwareTokenMfaConfiguration() CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference {
	var returns CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference
	_jsii_.Get(
		j,
		"softwareTokenMfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SoftwareTokenMfaConfigurationInput() *CognitoUserPoolSoftwareTokenMfaConfiguration {
	var returns *CognitoUserPoolSoftwareTokenMfaConfiguration
	_jsii_.Get(
		j,
		"softwareTokenMfaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UserAttributeUpdateSettings() CognitoUserPoolUserAttributeUpdateSettingsOutputReference {
	var returns CognitoUserPoolUserAttributeUpdateSettingsOutputReference
	_jsii_.Get(
		j,
		"userAttributeUpdateSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UserAttributeUpdateSettingsInput() *CognitoUserPoolUserAttributeUpdateSettings {
	var returns *CognitoUserPoolUserAttributeUpdateSettings
	_jsii_.Get(
		j,
		"userAttributeUpdateSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UsernameAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usernameAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UsernameAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usernameAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UsernameConfiguration() CognitoUserPoolUsernameConfigurationOutputReference {
	var returns CognitoUserPoolUsernameConfigurationOutputReference
	_jsii_.Get(
		j,
		"usernameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UsernameConfigurationInput() *CognitoUserPoolUsernameConfiguration {
	var returns *CognitoUserPoolUsernameConfiguration
	_jsii_.Get(
		j,
		"usernameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UserPoolAddOns() CognitoUserPoolUserPoolAddOnsOutputReference {
	var returns CognitoUserPoolUserPoolAddOnsOutputReference
	_jsii_.Get(
		j,
		"userPoolAddOns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UserPoolAddOnsInput() *CognitoUserPoolUserPoolAddOns {
	var returns *CognitoUserPoolUserPoolAddOns
	_jsii_.Get(
		j,
		"userPoolAddOnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) VerificationMessageTemplate() CognitoUserPoolVerificationMessageTemplateOutputReference {
	var returns CognitoUserPoolVerificationMessageTemplateOutputReference
	_jsii_.Get(
		j,
		"verificationMessageTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) VerificationMessageTemplateInput() *CognitoUserPoolVerificationMessageTemplate {
	var returns *CognitoUserPoolVerificationMessageTemplate
	_jsii_.Get(
		j,
		"verificationMessageTemplateInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/cognito_user_pool aws_cognito_user_pool} Resource.
func NewCognitoUserPool(scope constructs.Construct, id *string, config *CognitoUserPoolConfig) CognitoUserPool {
	_init_.Initialize()

	if err := validateNewCognitoUserPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUserPool{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/cognito_user_pool aws_cognito_user_pool} Resource.
func NewCognitoUserPool_Override(c CognitoUserPool, scope constructs.Construct, id *string, config *CognitoUserPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPool",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetAliasAttributes(val *[]*string) {
	if err := j.validateSetAliasAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasAttributes",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetAutoVerifiedAttributes(val *[]*string) {
	if err := j.validateSetAutoVerifiedAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoVerifiedAttributes",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetDeletionProtection(val *string) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetEmailVerificationMessage(val *string) {
	if err := j.validateSetEmailVerificationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetEmailVerificationSubject(val *string) {
	if err := j.validateSetEmailVerificationSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailVerificationSubject",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetMfaConfiguration(val *string) {
	if err := j.validateSetMfaConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mfaConfiguration",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetSmsAuthenticationMessage(val *string) {
	if err := j.validateSetSmsAuthenticationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsAuthenticationMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetSmsVerificationMessage(val *string) {
	if err := j.validateSetSmsVerificationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool)SetUsernameAttributes(val *[]*string) {
	if err := j.validateSetUsernameAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameAttributes",
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
func CognitoUserPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoUserPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitoUserPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoUserPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitoUserPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoUserPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoUserPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CognitoUserPool) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CognitoUserPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPool) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPool) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutAccountRecoverySetting(value *CognitoUserPoolAccountRecoverySetting) {
	if err := c.validatePutAccountRecoverySettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAccountRecoverySetting",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutAdminCreateUserConfig(value *CognitoUserPoolAdminCreateUserConfig) {
	if err := c.validatePutAdminCreateUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAdminCreateUserConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutDeviceConfiguration(value *CognitoUserPoolDeviceConfiguration) {
	if err := c.validatePutDeviceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDeviceConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutEmailConfiguration(value *CognitoUserPoolEmailConfiguration) {
	if err := c.validatePutEmailConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEmailConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutLambdaConfig(value *CognitoUserPoolLambdaConfig) {
	if err := c.validatePutLambdaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLambdaConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutPasswordPolicy(value *CognitoUserPoolPasswordPolicy) {
	if err := c.validatePutPasswordPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPasswordPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutSchema(value interface{}) {
	if err := c.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSchema",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutSmsConfiguration(value *CognitoUserPoolSmsConfiguration) {
	if err := c.validatePutSmsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSmsConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutSoftwareTokenMfaConfiguration(value *CognitoUserPoolSoftwareTokenMfaConfiguration) {
	if err := c.validatePutSoftwareTokenMfaConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSoftwareTokenMfaConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutUserAttributeUpdateSettings(value *CognitoUserPoolUserAttributeUpdateSettings) {
	if err := c.validatePutUserAttributeUpdateSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUserAttributeUpdateSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutUsernameConfiguration(value *CognitoUserPoolUsernameConfiguration) {
	if err := c.validatePutUsernameConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUsernameConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutUserPoolAddOns(value *CognitoUserPoolUserPoolAddOns) {
	if err := c.validatePutUserPoolAddOnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUserPoolAddOns",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutVerificationMessageTemplate(value *CognitoUserPoolVerificationMessageTemplate) {
	if err := c.validatePutVerificationMessageTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVerificationMessageTemplate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetAccountRecoverySetting() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountRecoverySetting",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetAdminCreateUserConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAdminCreateUserConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetAliasAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetAliasAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetAutoVerifiedAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoVerifiedAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		c,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetDeviceConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetDeviceConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetEmailConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetEmailVerificationMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailVerificationMessage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetEmailVerificationSubject() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailVerificationSubject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetLambdaConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLambdaConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetMfaConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetMfaConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetPasswordPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetPasswordPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetSchema() {
	_jsii_.InvokeVoid(
		c,
		"resetSchema",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetSmsAuthenticationMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetSmsAuthenticationMessage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetSmsConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetSmsConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetSmsVerificationMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetSmsVerificationMessage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetSoftwareTokenMfaConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetSoftwareTokenMfaConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetUserAttributeUpdateSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetUserAttributeUpdateSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetUsernameAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetUsernameAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetUsernameConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetUsernameConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetUserPoolAddOns() {
	_jsii_.InvokeVoid(
		c,
		"resetUserPoolAddOns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetVerificationMessageTemplate() {
	_jsii_.InvokeVoid(
		c,
		"resetVerificationMessageTemplate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

