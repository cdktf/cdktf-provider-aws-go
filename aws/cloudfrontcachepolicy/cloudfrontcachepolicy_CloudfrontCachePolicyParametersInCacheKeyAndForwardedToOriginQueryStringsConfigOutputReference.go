package cloudfrontcachepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/cloudfrontcachepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig
	SetInternalValue(val *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig)
	QueryStringBehavior() *string
	SetQueryStringBehavior(val *string)
	QueryStringBehaviorInput() *string
	QueryStrings() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference
	QueryStringsInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings
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
	PutQueryStrings(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings)
	ResetQueryStrings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference
type jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) InternalValue() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig {
	var returns *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) QueryStringBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) QueryStringBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) QueryStrings() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference {
	var returns CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference
	_jsii_.Get(
		j,
		"queryStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) QueryStringsInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings {
	var returns *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings
	_jsii_.Get(
		j,
		"queryStringsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontCachePolicy.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference_Override(c CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontCachePolicy.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference)SetInternalValue(val *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference)SetQueryStringBehavior(val *string) {
	if err := j.validateSetQueryStringBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringBehavior",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) PutQueryStrings(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) {
	if err := c.validatePutQueryStringsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQueryStrings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) ResetQueryStrings() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStrings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

