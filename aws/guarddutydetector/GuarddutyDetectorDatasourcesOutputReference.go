// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guarddutydetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/guarddutydetector/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GuarddutyDetectorDatasourcesOutputReference interface {
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
	InternalValue() *GuarddutyDetectorDatasources
	SetInternalValue(val *GuarddutyDetectorDatasources)
	Kubernetes() GuarddutyDetectorDatasourcesKubernetesOutputReference
	KubernetesInput() *GuarddutyDetectorDatasourcesKubernetes
	MalwareProtection() GuarddutyDetectorDatasourcesMalwareProtectionOutputReference
	MalwareProtectionInput() *GuarddutyDetectorDatasourcesMalwareProtection
	S3Logs() GuarddutyDetectorDatasourcesS3LogsOutputReference
	S3LogsInput() *GuarddutyDetectorDatasourcesS3Logs
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
	PutKubernetes(value *GuarddutyDetectorDatasourcesKubernetes)
	PutMalwareProtection(value *GuarddutyDetectorDatasourcesMalwareProtection)
	PutS3Logs(value *GuarddutyDetectorDatasourcesS3Logs)
	ResetKubernetes()
	ResetMalwareProtection()
	ResetS3Logs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GuarddutyDetectorDatasourcesOutputReference
type jsiiProxy_GuarddutyDetectorDatasourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) InternalValue() *GuarddutyDetectorDatasources {
	var returns *GuarddutyDetectorDatasources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) Kubernetes() GuarddutyDetectorDatasourcesKubernetesOutputReference {
	var returns GuarddutyDetectorDatasourcesKubernetesOutputReference
	_jsii_.Get(
		j,
		"kubernetes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) KubernetesInput() *GuarddutyDetectorDatasourcesKubernetes {
	var returns *GuarddutyDetectorDatasourcesKubernetes
	_jsii_.Get(
		j,
		"kubernetesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) MalwareProtection() GuarddutyDetectorDatasourcesMalwareProtectionOutputReference {
	var returns GuarddutyDetectorDatasourcesMalwareProtectionOutputReference
	_jsii_.Get(
		j,
		"malwareProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) MalwareProtectionInput() *GuarddutyDetectorDatasourcesMalwareProtection {
	var returns *GuarddutyDetectorDatasourcesMalwareProtection
	_jsii_.Get(
		j,
		"malwareProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) S3Logs() GuarddutyDetectorDatasourcesS3LogsOutputReference {
	var returns GuarddutyDetectorDatasourcesS3LogsOutputReference
	_jsii_.Get(
		j,
		"s3Logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) S3LogsInput() *GuarddutyDetectorDatasourcesS3Logs {
	var returns *GuarddutyDetectorDatasourcesS3Logs
	_jsii_.Get(
		j,
		"s3LogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGuarddutyDetectorDatasourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GuarddutyDetectorDatasourcesOutputReference {
	_init_.Initialize()

	if err := validateNewGuarddutyDetectorDatasourcesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GuarddutyDetectorDatasourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.guarddutyDetector.GuarddutyDetectorDatasourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGuarddutyDetectorDatasourcesOutputReference_Override(g GuarddutyDetectorDatasourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.guarddutyDetector.GuarddutyDetectorDatasourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference)SetInternalValue(val *GuarddutyDetectorDatasources) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) PutKubernetes(value *GuarddutyDetectorDatasourcesKubernetes) {
	if err := g.validatePutKubernetesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKubernetes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) PutMalwareProtection(value *GuarddutyDetectorDatasourcesMalwareProtection) {
	if err := g.validatePutMalwareProtectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMalwareProtection",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) PutS3Logs(value *GuarddutyDetectorDatasourcesS3Logs) {
	if err := g.validatePutS3LogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putS3Logs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) ResetKubernetes() {
	_jsii_.InvokeVoid(
		g,
		"resetKubernetes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) ResetMalwareProtection() {
	_jsii_.InvokeVoid(
		g,
		"resetMalwareProtection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) ResetS3Logs() {
	_jsii_.InvokeVoid(
		g,
		"resetS3Logs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyDetectorDatasourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

