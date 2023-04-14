package cloudfrontresponseheaderspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v13/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v13/cloudfrontresponseheaderspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference interface {
	cdktf.ComplexObject
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
	ContentSecurityPolicy() CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference
	ContentSecurityPolicyInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy
	ContentTypeOptions() CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference
	ContentTypeOptionsInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	FrameOptions() CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference
	FrameOptionsInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions
	InternalValue() *CloudfrontResponseHeadersPolicySecurityHeadersConfig
	SetInternalValue(val *CloudfrontResponseHeadersPolicySecurityHeadersConfig)
	ReferrerPolicy() CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference
	ReferrerPolicyInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy
	StrictTransportSecurity() CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference
	StrictTransportSecurityInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	XssProtection() CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference
	XssProtectionInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection
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
	PutContentSecurityPolicy(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy)
	PutContentTypeOptions(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions)
	PutFrameOptions(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions)
	PutReferrerPolicy(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy)
	PutStrictTransportSecurity(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity)
	PutXssProtection(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection)
	ResetContentSecurityPolicy()
	ResetContentTypeOptions()
	ResetFrameOptions()
	ResetReferrerPolicy()
	ResetStrictTransportSecurity()
	ResetXssProtection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ContentSecurityPolicy() CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference
	_jsii_.Get(
		j,
		"contentSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ContentSecurityPolicyInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy
	_jsii_.Get(
		j,
		"contentSecurityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ContentTypeOptions() CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"contentTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ContentTypeOptionsInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions
	_jsii_.Get(
		j,
		"contentTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) FrameOptions() CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference
	_jsii_.Get(
		j,
		"frameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) FrameOptionsInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions
	_jsii_.Get(
		j,
		"frameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) InternalValue() *CloudfrontResponseHeadersPolicySecurityHeadersConfig {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ReferrerPolicy() CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference
	_jsii_.Get(
		j,
		"referrerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ReferrerPolicyInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy
	_jsii_.Get(
		j,
		"referrerPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) StrictTransportSecurity() CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference
	_jsii_.Get(
		j,
		"strictTransportSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) StrictTransportSecurityInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity
	_jsii_.Get(
		j,
		"strictTransportSecurityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) XssProtection() CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference
	_jsii_.Get(
		j,
		"xssProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) XssProtectionInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection
	_jsii_.Get(
		j,
		"xssProtectionInput",
		&returns,
	)
	return returns
}


func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference_Override(c CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference)SetInternalValue(val *CloudfrontResponseHeadersPolicySecurityHeadersConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) PutContentSecurityPolicy(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) {
	if err := c.validatePutContentSecurityPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putContentSecurityPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) PutContentTypeOptions(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) {
	if err := c.validatePutContentTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putContentTypeOptions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) PutFrameOptions(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) {
	if err := c.validatePutFrameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFrameOptions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) PutReferrerPolicy(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) {
	if err := c.validatePutReferrerPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReferrerPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) PutStrictTransportSecurity(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) {
	if err := c.validatePutStrictTransportSecurityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStrictTransportSecurity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) PutXssProtection(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) {
	if err := c.validatePutXssProtectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putXssProtection",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ResetContentSecurityPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetContentSecurityPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ResetContentTypeOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetContentTypeOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ResetFrameOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetFrameOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ResetReferrerPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetReferrerPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ResetStrictTransportSecurity() {
	_jsii_.InvokeVoid(
		c,
		"resetStrictTransportSecurity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ResetXssProtection() {
	_jsii_.InvokeVoid(
		c,
		"resetXssProtection",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

