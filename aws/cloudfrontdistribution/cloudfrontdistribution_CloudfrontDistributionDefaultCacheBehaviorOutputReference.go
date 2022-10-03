package cloudfrontdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/cloudfrontdistribution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontDistributionDefaultCacheBehaviorOutputReference interface {
	cdktf.ComplexObject
	AllowedMethods() *[]*string
	SetAllowedMethods(val *[]*string)
	AllowedMethodsInput() *[]*string
	CachedMethods() *[]*string
	SetCachedMethods(val *[]*string)
	CachedMethodsInput() *[]*string
	CachePolicyId() *string
	SetCachePolicyId(val *string)
	CachePolicyIdInput() *string
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
	Compress() interface{}
	SetCompress(val interface{})
	CompressInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultTtl() *float64
	SetDefaultTtl(val *float64)
	DefaultTtlInput() *float64
	FieldLevelEncryptionId() *string
	SetFieldLevelEncryptionId(val *string)
	FieldLevelEncryptionIdInput() *string
	ForwardedValues() CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference
	ForwardedValuesInput() *CloudfrontDistributionDefaultCacheBehaviorForwardedValues
	// Experimental.
	Fqn() *string
	FunctionAssociation() CloudfrontDistributionDefaultCacheBehaviorFunctionAssociationList
	FunctionAssociationInput() interface{}
	InternalValue() *CloudfrontDistributionDefaultCacheBehavior
	SetInternalValue(val *CloudfrontDistributionDefaultCacheBehavior)
	LambdaFunctionAssociation() CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociationList
	LambdaFunctionAssociationInput() interface{}
	MaxTtl() *float64
	SetMaxTtl(val *float64)
	MaxTtlInput() *float64
	MinTtl() *float64
	SetMinTtl(val *float64)
	MinTtlInput() *float64
	OriginRequestPolicyId() *string
	SetOriginRequestPolicyId(val *string)
	OriginRequestPolicyIdInput() *string
	RealtimeLogConfigArn() *string
	SetRealtimeLogConfigArn(val *string)
	RealtimeLogConfigArnInput() *string
	ResponseHeadersPolicyId() *string
	SetResponseHeadersPolicyId(val *string)
	ResponseHeadersPolicyIdInput() *string
	SmoothStreaming() interface{}
	SetSmoothStreaming(val interface{})
	SmoothStreamingInput() interface{}
	TargetOriginId() *string
	SetTargetOriginId(val *string)
	TargetOriginIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustedKeyGroups() *[]*string
	SetTrustedKeyGroups(val *[]*string)
	TrustedKeyGroupsInput() *[]*string
	TrustedSigners() *[]*string
	SetTrustedSigners(val *[]*string)
	TrustedSignersInput() *[]*string
	ViewerProtocolPolicy() *string
	SetViewerProtocolPolicy(val *string)
	ViewerProtocolPolicyInput() *string
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
	PutForwardedValues(value *CloudfrontDistributionDefaultCacheBehaviorForwardedValues)
	PutFunctionAssociation(value interface{})
	PutLambdaFunctionAssociation(value interface{})
	ResetCachePolicyId()
	ResetCompress()
	ResetDefaultTtl()
	ResetFieldLevelEncryptionId()
	ResetForwardedValues()
	ResetFunctionAssociation()
	ResetLambdaFunctionAssociation()
	ResetMaxTtl()
	ResetMinTtl()
	ResetOriginRequestPolicyId()
	ResetRealtimeLogConfigArn()
	ResetResponseHeadersPolicyId()
	ResetSmoothStreaming()
	ResetTrustedKeyGroups()
	ResetTrustedSigners()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontDistributionDefaultCacheBehaviorOutputReference
type jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) AllowedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) AllowedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) CachedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cachedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) CachedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cachedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) CachePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) CachePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) Compress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) CompressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) FieldLevelEncryptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) FieldLevelEncryptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ForwardedValues() CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference {
	var returns CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference
	_jsii_.Get(
		j,
		"forwardedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ForwardedValuesInput() *CloudfrontDistributionDefaultCacheBehaviorForwardedValues {
	var returns *CloudfrontDistributionDefaultCacheBehaviorForwardedValues
	_jsii_.Get(
		j,
		"forwardedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) FunctionAssociation() CloudfrontDistributionDefaultCacheBehaviorFunctionAssociationList {
	var returns CloudfrontDistributionDefaultCacheBehaviorFunctionAssociationList
	_jsii_.Get(
		j,
		"functionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) FunctionAssociationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) InternalValue() *CloudfrontDistributionDefaultCacheBehavior {
	var returns *CloudfrontDistributionDefaultCacheBehavior
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) LambdaFunctionAssociation() CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociationList {
	var returns CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociationList
	_jsii_.Get(
		j,
		"lambdaFunctionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) LambdaFunctionAssociationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) MaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) MaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) MinTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) MinTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) OriginRequestPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) OriginRequestPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) RealtimeLogConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) RealtimeLogConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResponseHeadersPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResponseHeadersPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SmoothStreaming() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreaming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SmoothStreamingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreamingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TargetOriginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TargetOriginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TrustedKeyGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TrustedKeyGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedKeyGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TrustedSigners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedSigners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TrustedSignersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedSignersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ViewerProtocolPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ViewerProtocolPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicyInput",
		&returns,
	)
	return returns
}


func NewCloudfrontDistributionDefaultCacheBehaviorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontDistributionDefaultCacheBehaviorOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontDistributionDefaultCacheBehaviorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionDefaultCacheBehaviorOutputReference_Override(c CloudfrontDistributionDefaultCacheBehaviorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetAllowedMethods(val *[]*string) {
	if err := j.validateSetAllowedMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedMethods",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetCachedMethods(val *[]*string) {
	if err := j.validateSetCachedMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachedMethods",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetCachePolicyId(val *string) {
	if err := j.validateSetCachePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachePolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetCompress(val interface{}) {
	if err := j.validateSetCompressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compress",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetFieldLevelEncryptionId(val *string) {
	if err := j.validateSetFieldLevelEncryptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLevelEncryptionId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetInternalValue(val *CloudfrontDistributionDefaultCacheBehavior) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetMaxTtl(val *float64) {
	if err := j.validateSetMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetMinTtl(val *float64) {
	if err := j.validateSetMinTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTtl",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetOriginRequestPolicyId(val *string) {
	if err := j.validateSetOriginRequestPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originRequestPolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetRealtimeLogConfigArn(val *string) {
	if err := j.validateSetRealtimeLogConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realtimeLogConfigArn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetResponseHeadersPolicyId(val *string) {
	if err := j.validateSetResponseHeadersPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseHeadersPolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetSmoothStreaming(val interface{}) {
	if err := j.validateSetSmoothStreamingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smoothStreaming",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetTargetOriginId(val *string) {
	if err := j.validateSetTargetOriginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetOriginId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetTrustedKeyGroups(val *[]*string) {
	if err := j.validateSetTrustedKeyGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedKeyGroups",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetTrustedSigners(val *[]*string) {
	if err := j.validateSetTrustedSignersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedSigners",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference)SetViewerProtocolPolicy(val *string) {
	if err := j.validateSetViewerProtocolPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewerProtocolPolicy",
		val,
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) PutForwardedValues(value *CloudfrontDistributionDefaultCacheBehaviorForwardedValues) {
	if err := c.validatePutForwardedValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putForwardedValues",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) PutFunctionAssociation(value interface{}) {
	if err := c.validatePutFunctionAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFunctionAssociation",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) PutLambdaFunctionAssociation(value interface{}) {
	if err := c.validatePutLambdaFunctionAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLambdaFunctionAssociation",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetCachePolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetCachePolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetCompress() {
	_jsii_.InvokeVoid(
		c,
		"resetCompress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetFieldLevelEncryptionId() {
	_jsii_.InvokeVoid(
		c,
		"resetFieldLevelEncryptionId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetForwardedValues() {
	_jsii_.InvokeVoid(
		c,
		"resetForwardedValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetFunctionAssociation() {
	_jsii_.InvokeVoid(
		c,
		"resetFunctionAssociation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetLambdaFunctionAssociation() {
	_jsii_.InvokeVoid(
		c,
		"resetLambdaFunctionAssociation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetMinTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetMinTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetOriginRequestPolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginRequestPolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetRealtimeLogConfigArn() {
	_jsii_.InvokeVoid(
		c,
		"resetRealtimeLogConfigArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetResponseHeadersPolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetResponseHeadersPolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetSmoothStreaming() {
	_jsii_.InvokeVoid(
		c,
		"resetSmoothStreaming",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetTrustedKeyGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetTrustedKeyGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetTrustedSigners() {
	_jsii_.InvokeVoid(
		c,
		"resetTrustedSigners",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

