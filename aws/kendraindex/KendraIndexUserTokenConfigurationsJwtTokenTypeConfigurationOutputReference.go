package kendraindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v15/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v15/kendraindex/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference interface {
	cdktf.ComplexObject
	ClaimRegex() *string
	SetClaimRegex(val *string)
	ClaimRegexInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GroupAttributeField() *string
	SetGroupAttributeField(val *string)
	GroupAttributeFieldInput() *string
	InternalValue() *KendraIndexUserTokenConfigurationsJwtTokenTypeConfiguration
	SetInternalValue(val *KendraIndexUserTokenConfigurationsJwtTokenTypeConfiguration)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	KeyLocation() *string
	SetKeyLocation(val *string)
	KeyLocationInput() *string
	SecretsManagerArn() *string
	SetSecretsManagerArn(val *string)
	SecretsManagerArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	UserNameAttributeField() *string
	SetUserNameAttributeField(val *string)
	UserNameAttributeFieldInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetClaimRegex()
	ResetGroupAttributeField()
	ResetIssuer()
	ResetSecretsManagerArn()
	ResetUrl()
	ResetUserNameAttributeField()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference
type jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) ClaimRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"claimRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) ClaimRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"claimRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) GroupAttributeField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupAttributeField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) GroupAttributeFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupAttributeFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) InternalValue() *KendraIndexUserTokenConfigurationsJwtTokenTypeConfiguration {
	var returns *KendraIndexUserTokenConfigurationsJwtTokenTypeConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) KeyLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) KeyLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) SecretsManagerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) SecretsManagerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) UserNameAttributeField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameAttributeField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) UserNameAttributeFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameAttributeFieldInput",
		&returns,
	)
	return returns
}


func NewKendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kendraIndex.KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference_Override(k KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kendraIndex.KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference)SetClaimRegex(val *string) {
	if err := j.validateSetClaimRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"claimRegex",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference)SetGroupAttributeField(val *string) {
	if err := j.validateSetGroupAttributeFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupAttributeField",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference)SetInternalValue(val *KendraIndexUserTokenConfigurationsJwtTokenTypeConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference)SetKeyLocation(val *string) {
	if err := j.validateSetKeyLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyLocation",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference)SetSecretsManagerArn(val *string) {
	if err := j.validateSetSecretsManagerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerArn",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference)SetUserNameAttributeField(val *string) {
	if err := j.validateSetUserNameAttributeFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameAttributeField",
		val,
	)
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) ResetClaimRegex() {
	_jsii_.InvokeVoid(
		k,
		"resetClaimRegex",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) ResetGroupAttributeField() {
	_jsii_.InvokeVoid(
		k,
		"resetGroupAttributeField",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		k,
		"resetIssuer",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) ResetSecretsManagerArn() {
	_jsii_.InvokeVoid(
		k,
		"resetSecretsManagerArn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		k,
		"resetUrl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) ResetUserNameAttributeField() {
	_jsii_.InvokeVoid(
		k,
		"resetUserNameAttributeField",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

