// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclAssociationConfigRequestBodyOutputReference interface {
	cdktf.ComplexObject
	ApiGateway() Wafv2WebAclAssociationConfigRequestBodyApiGatewayOutputReference
	ApiGatewayInput() *Wafv2WebAclAssociationConfigRequestBodyApiGateway
	AppRunnerService() Wafv2WebAclAssociationConfigRequestBodyAppRunnerServiceOutputReference
	AppRunnerServiceInput() *Wafv2WebAclAssociationConfigRequestBodyAppRunnerService
	Cloudfront() Wafv2WebAclAssociationConfigRequestBodyCloudfrontOutputReference
	CloudfrontInput() *Wafv2WebAclAssociationConfigRequestBodyCloudfront
	CognitoUserPool() Wafv2WebAclAssociationConfigRequestBodyCognitoUserPoolOutputReference
	CognitoUserPoolInput() *Wafv2WebAclAssociationConfigRequestBodyCognitoUserPool
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VerifiedAccessInstance() Wafv2WebAclAssociationConfigRequestBodyVerifiedAccessInstanceOutputReference
	VerifiedAccessInstanceInput() *Wafv2WebAclAssociationConfigRequestBodyVerifiedAccessInstance
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
	PutApiGateway(value *Wafv2WebAclAssociationConfigRequestBodyApiGateway)
	PutAppRunnerService(value *Wafv2WebAclAssociationConfigRequestBodyAppRunnerService)
	PutCloudfront(value *Wafv2WebAclAssociationConfigRequestBodyCloudfront)
	PutCognitoUserPool(value *Wafv2WebAclAssociationConfigRequestBodyCognitoUserPool)
	PutVerifiedAccessInstance(value *Wafv2WebAclAssociationConfigRequestBodyVerifiedAccessInstance)
	ResetApiGateway()
	ResetAppRunnerService()
	ResetCloudfront()
	ResetCognitoUserPool()
	ResetVerifiedAccessInstance()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2WebAclAssociationConfigRequestBodyOutputReference
type jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) ApiGateway() Wafv2WebAclAssociationConfigRequestBodyApiGatewayOutputReference {
	var returns Wafv2WebAclAssociationConfigRequestBodyApiGatewayOutputReference
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) ApiGatewayInput() *Wafv2WebAclAssociationConfigRequestBodyApiGateway {
	var returns *Wafv2WebAclAssociationConfigRequestBodyApiGateway
	_jsii_.Get(
		j,
		"apiGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) AppRunnerService() Wafv2WebAclAssociationConfigRequestBodyAppRunnerServiceOutputReference {
	var returns Wafv2WebAclAssociationConfigRequestBodyAppRunnerServiceOutputReference
	_jsii_.Get(
		j,
		"appRunnerService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) AppRunnerServiceInput() *Wafv2WebAclAssociationConfigRequestBodyAppRunnerService {
	var returns *Wafv2WebAclAssociationConfigRequestBodyAppRunnerService
	_jsii_.Get(
		j,
		"appRunnerServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) Cloudfront() Wafv2WebAclAssociationConfigRequestBodyCloudfrontOutputReference {
	var returns Wafv2WebAclAssociationConfigRequestBodyCloudfrontOutputReference
	_jsii_.Get(
		j,
		"cloudfront",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) CloudfrontInput() *Wafv2WebAclAssociationConfigRequestBodyCloudfront {
	var returns *Wafv2WebAclAssociationConfigRequestBodyCloudfront
	_jsii_.Get(
		j,
		"cloudfrontInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) CognitoUserPool() Wafv2WebAclAssociationConfigRequestBodyCognitoUserPoolOutputReference {
	var returns Wafv2WebAclAssociationConfigRequestBodyCognitoUserPoolOutputReference
	_jsii_.Get(
		j,
		"cognitoUserPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) CognitoUserPoolInput() *Wafv2WebAclAssociationConfigRequestBodyCognitoUserPool {
	var returns *Wafv2WebAclAssociationConfigRequestBodyCognitoUserPool
	_jsii_.Get(
		j,
		"cognitoUserPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) VerifiedAccessInstance() Wafv2WebAclAssociationConfigRequestBodyVerifiedAccessInstanceOutputReference {
	var returns Wafv2WebAclAssociationConfigRequestBodyVerifiedAccessInstanceOutputReference
	_jsii_.Get(
		j,
		"verifiedAccessInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) VerifiedAccessInstanceInput() *Wafv2WebAclAssociationConfigRequestBodyVerifiedAccessInstance {
	var returns *Wafv2WebAclAssociationConfigRequestBodyVerifiedAccessInstance
	_jsii_.Get(
		j,
		"verifiedAccessInstanceInput",
		&returns,
	)
	return returns
}


func NewWafv2WebAclAssociationConfigRequestBodyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Wafv2WebAclAssociationConfigRequestBodyOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclAssociationConfigRequestBodyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclAssociationConfigRequestBodyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafv2WebAclAssociationConfigRequestBodyOutputReference_Override(w Wafv2WebAclAssociationConfigRequestBodyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclAssociationConfigRequestBodyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) PutApiGateway(value *Wafv2WebAclAssociationConfigRequestBodyApiGateway) {
	if err := w.validatePutApiGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putApiGateway",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) PutAppRunnerService(value *Wafv2WebAclAssociationConfigRequestBodyAppRunnerService) {
	if err := w.validatePutAppRunnerServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAppRunnerService",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) PutCloudfront(value *Wafv2WebAclAssociationConfigRequestBodyCloudfront) {
	if err := w.validatePutCloudfrontParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCloudfront",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) PutCognitoUserPool(value *Wafv2WebAclAssociationConfigRequestBodyCognitoUserPool) {
	if err := w.validatePutCognitoUserPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCognitoUserPool",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) PutVerifiedAccessInstance(value *Wafv2WebAclAssociationConfigRequestBodyVerifiedAccessInstance) {
	if err := w.validatePutVerifiedAccessInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putVerifiedAccessInstance",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) ResetApiGateway() {
	_jsii_.InvokeVoid(
		w,
		"resetApiGateway",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) ResetAppRunnerService() {
	_jsii_.InvokeVoid(
		w,
		"resetAppRunnerService",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) ResetCloudfront() {
	_jsii_.InvokeVoid(
		w,
		"resetCloudfront",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) ResetCognitoUserPool() {
	_jsii_.InvokeVoid(
		w,
		"resetCognitoUserPool",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) ResetVerifiedAccessInstance() {
	_jsii_.InvokeVoid(
		w,
		"resetVerifiedAccessInstance",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclAssociationConfigRequestBodyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

