package eksidentityproviderconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/eksidentityproviderconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EksIdentityProviderConfigOidcOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
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
	GroupsClaim() *string
	SetGroupsClaim(val *string)
	GroupsClaimInput() *string
	GroupsPrefix() *string
	SetGroupsPrefix(val *string)
	GroupsPrefixInput() *string
	IdentityProviderConfigName() *string
	SetIdentityProviderConfigName(val *string)
	IdentityProviderConfigNameInput() *string
	InternalValue() *EksIdentityProviderConfigOidc
	SetInternalValue(val *EksIdentityProviderConfigOidc)
	IssuerUrl() *string
	SetIssuerUrl(val *string)
	IssuerUrlInput() *string
	RequiredClaims() *map[string]*string
	SetRequiredClaims(val *map[string]*string)
	RequiredClaimsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsernameClaim() *string
	SetUsernameClaim(val *string)
	UsernameClaimInput() *string
	UsernamePrefix() *string
	SetUsernamePrefix(val *string)
	UsernamePrefixInput() *string
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
	ResetGroupsClaim()
	ResetGroupsPrefix()
	ResetRequiredClaims()
	ResetUsernameClaim()
	ResetUsernamePrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EksIdentityProviderConfigOidcOutputReference
type jsiiProxy_EksIdentityProviderConfigOidcOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) GroupsClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) GroupsClaimInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) GroupsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) GroupsPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) IdentityProviderConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) IdentityProviderConfigNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderConfigNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) InternalValue() *EksIdentityProviderConfigOidc {
	var returns *EksIdentityProviderConfigOidc
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) IssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) IssuerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) RequiredClaims() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requiredClaims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) RequiredClaimsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requiredClaimsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) UsernameClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) UsernameClaimInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) UsernamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) UsernamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernamePrefixInput",
		&returns,
	)
	return returns
}


func NewEksIdentityProviderConfigOidcOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EksIdentityProviderConfigOidcOutputReference {
	_init_.Initialize()

	if err := validateNewEksIdentityProviderConfigOidcOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EksIdentityProviderConfigOidcOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.eksIdentityProviderConfig.EksIdentityProviderConfigOidcOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEksIdentityProviderConfigOidcOutputReference_Override(e EksIdentityProviderConfigOidcOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.eksIdentityProviderConfig.EksIdentityProviderConfigOidcOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference)SetGroupsClaim(val *string) {
	if err := j.validateSetGroupsClaimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsClaim",
		val,
	)
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference)SetGroupsPrefix(val *string) {
	if err := j.validateSetGroupsPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsPrefix",
		val,
	)
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference)SetIdentityProviderConfigName(val *string) {
	if err := j.validateSetIdentityProviderConfigNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderConfigName",
		val,
	)
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference)SetInternalValue(val *EksIdentityProviderConfigOidc) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference)SetIssuerUrl(val *string) {
	if err := j.validateSetIssuerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerUrl",
		val,
	)
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference)SetRequiredClaims(val *map[string]*string) {
	if err := j.validateSetRequiredClaimsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredClaims",
		val,
	)
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference)SetUsernameClaim(val *string) {
	if err := j.validateSetUsernameClaimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameClaim",
		val,
	)
}

func (j *jsiiProxy_EksIdentityProviderConfigOidcOutputReference)SetUsernamePrefix(val *string) {
	if err := j.validateSetUsernamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernamePrefix",
		val,
	)
}

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) ResetGroupsClaim() {
	_jsii_.InvokeVoid(
		e,
		"resetGroupsClaim",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) ResetGroupsPrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetGroupsPrefix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) ResetRequiredClaims() {
	_jsii_.InvokeVoid(
		e,
		"resetRequiredClaims",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) ResetUsernameClaim() {
	_jsii_.InvokeVoid(
		e,
		"resetUsernameClaim",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) ResetUsernamePrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetUsernamePrefix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksIdentityProviderConfigOidcOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

