package acm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/acm/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority aws_acmpca_certificate_authority}.
type AcmpcaCertificateAuthority interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	CertificateAuthorityConfiguration() AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference
	CertificateAuthorityConfigurationInput() *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration
	CertificateChain() *string
	CertificateSigningRequest() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NotAfter() *string
	NotBefore() *string
	PermanentDeletionTimeInDays() *float64
	SetPermanentDeletionTimeInDays(val *float64)
	PermanentDeletionTimeInDaysInput() *float64
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
	RevocationConfiguration() AcmpcaCertificateAuthorityRevocationConfigurationOutputReference
	RevocationConfigurationInput() *AcmpcaCertificateAuthorityRevocationConfiguration
	Serial() *string
	Status() *string
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
	Timeouts() AcmpcaCertificateAuthorityTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutCertificateAuthorityConfiguration(value *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration)
	PutRevocationConfiguration(value *AcmpcaCertificateAuthorityRevocationConfiguration)
	PutTimeouts(value *AcmpcaCertificateAuthorityTimeouts)
	ResetEnabled()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPermanentDeletionTimeInDays()
	ResetRevocationConfiguration()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AcmpcaCertificateAuthority
type jsiiProxy_AcmpcaCertificateAuthority struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) CertificateAuthorityConfiguration() AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference {
	var returns AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference
	_jsii_.Get(
		j,
		"certificateAuthorityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) CertificateAuthorityConfigurationInput() *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration {
	var returns *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration
	_jsii_.Get(
		j,
		"certificateAuthorityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) CertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) CertificateSigningRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSigningRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) NotAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) NotBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) PermanentDeletionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"permanentDeletionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) PermanentDeletionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"permanentDeletionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) RevocationConfiguration() AcmpcaCertificateAuthorityRevocationConfigurationOutputReference {
	var returns AcmpcaCertificateAuthorityRevocationConfigurationOutputReference
	_jsii_.Get(
		j,
		"revocationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) RevocationConfigurationInput() *AcmpcaCertificateAuthorityRevocationConfiguration {
	var returns *AcmpcaCertificateAuthorityRevocationConfiguration
	_jsii_.Get(
		j,
		"revocationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Serial() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Timeouts() AcmpcaCertificateAuthorityTimeoutsOutputReference {
	var returns AcmpcaCertificateAuthorityTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority aws_acmpca_certificate_authority} Resource.
func NewAcmpcaCertificateAuthority(scope constructs.Construct, id *string, config *AcmpcaCertificateAuthorityConfig) AcmpcaCertificateAuthority {
	_init_.Initialize()

	if err := validateNewAcmpcaCertificateAuthorityParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AcmpcaCertificateAuthority{}

	_jsii_.Create(
		"@cdktf/provider-aws.acm.AcmpcaCertificateAuthority",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority aws_acmpca_certificate_authority} Resource.
func NewAcmpcaCertificateAuthority_Override(a AcmpcaCertificateAuthority, scope constructs.Construct, id *string, config *AcmpcaCertificateAuthorityConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.acm.AcmpcaCertificateAuthority",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority)SetPermanentDeletionTimeInDays(val *float64) {
	if err := j.validateSetPermanentDeletionTimeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permanentDeletionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
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
func AcmpcaCertificateAuthority_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAcmpcaCertificateAuthority_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.acm.AcmpcaCertificateAuthority",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AcmpcaCertificateAuthority_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.acm.AcmpcaCertificateAuthority",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) PutCertificateAuthorityConfiguration(value *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration) {
	if err := a.validatePutCertificateAuthorityConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCertificateAuthorityConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) PutRevocationConfiguration(value *AcmpcaCertificateAuthorityRevocationConfiguration) {
	if err := a.validatePutRevocationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRevocationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) PutTimeouts(value *AcmpcaCertificateAuthorityTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetPermanentDeletionTimeInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetPermanentDeletionTimeInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetRevocationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRevocationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
