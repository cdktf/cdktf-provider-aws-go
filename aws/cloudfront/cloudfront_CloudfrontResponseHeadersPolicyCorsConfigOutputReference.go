package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/cloudfront/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontResponseHeadersPolicyCorsConfigOutputReference interface {
	cdktf.ComplexObject
	AccessControlAllowCredentials() interface{}
	SetAccessControlAllowCredentials(val interface{})
	AccessControlAllowCredentialsInput() interface{}
	AccessControlAllowHeaders() CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference
	AccessControlAllowHeadersInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders
	AccessControlAllowMethods() CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference
	AccessControlAllowMethodsInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods
	AccessControlAllowOrigins() CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference
	AccessControlAllowOriginsInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins
	AccessControlExposeHeaders() CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference
	AccessControlExposeHeadersInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders
	AccessControlMaxAgeSec() *float64
	SetAccessControlMaxAgeSec(val *float64)
	AccessControlMaxAgeSecInput() *float64
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
	InternalValue() *CloudfrontResponseHeadersPolicyCorsConfig
	SetInternalValue(val *CloudfrontResponseHeadersPolicyCorsConfig)
	OriginOverride() interface{}
	SetOriginOverride(val interface{})
	OriginOverrideInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutAccessControlAllowHeaders(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders)
	PutAccessControlAllowMethods(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods)
	PutAccessControlAllowOrigins(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins)
	PutAccessControlExposeHeaders(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders)
	ResetAccessControlExposeHeaders()
	ResetAccessControlMaxAgeSec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicyCorsConfigOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowHeaders() CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference {
	var returns CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowHeadersInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders {
	var returns *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders
	_jsii_.Get(
		j,
		"accessControlAllowHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowMethods() CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference {
	var returns CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowMethodsInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods {
	var returns *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods
	_jsii_.Get(
		j,
		"accessControlAllowMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowOrigins() CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference {
	var returns CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowOriginsInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins {
	var returns *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins
	_jsii_.Get(
		j,
		"accessControlAllowOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlExposeHeaders() CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference {
	var returns CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference
	_jsii_.Get(
		j,
		"accessControlExposeHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlExposeHeadersInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders {
	var returns *CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders
	_jsii_.Get(
		j,
		"accessControlExposeHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlMaxAgeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlMaxAgeSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) InternalValue() *CloudfrontResponseHeadersPolicyCorsConfig {
	var returns *CloudfrontResponseHeadersPolicyCorsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) OriginOverride() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) OriginOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontResponseHeadersPolicyCorsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontResponseHeadersPolicyCorsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontResponseHeadersPolicyCorsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfront.CloudfrontResponseHeadersPolicyCorsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicyCorsConfigOutputReference_Override(c CloudfrontResponseHeadersPolicyCorsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfront.CloudfrontResponseHeadersPolicyCorsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference)SetAccessControlAllowCredentials(val interface{}) {
	if err := j.validateSetAccessControlAllowCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessControlAllowCredentials",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference)SetAccessControlMaxAgeSec(val *float64) {
	if err := j.validateSetAccessControlMaxAgeSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessControlMaxAgeSec",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference)SetInternalValue(val *CloudfrontResponseHeadersPolicyCorsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference)SetOriginOverride(val interface{}) {
	if err := j.validateSetOriginOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originOverride",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) PutAccessControlAllowHeaders(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) {
	if err := c.validatePutAccessControlAllowHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAccessControlAllowHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) PutAccessControlAllowMethods(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) {
	if err := c.validatePutAccessControlAllowMethodsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAccessControlAllowMethods",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) PutAccessControlAllowOrigins(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) {
	if err := c.validatePutAccessControlAllowOriginsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAccessControlAllowOrigins",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) PutAccessControlExposeHeaders(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) {
	if err := c.validatePutAccessControlExposeHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAccessControlExposeHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) ResetAccessControlExposeHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessControlExposeHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) ResetAccessControlMaxAgeSec() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessControlMaxAgeSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

