// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/batchjobdefinition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference interface {
	cdktf.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	Env() BatchJobDefinitionEksPropertiesPodPropertiesContainersEnvList
	EnvInput() interface{}
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	ImagePullPolicy() *string
	SetImagePullPolicy(val *string)
	ImagePullPolicyInput() *string
	InternalValue() *BatchJobDefinitionEksPropertiesPodPropertiesContainers
	SetInternalValue(val *BatchJobDefinitionEksPropertiesPodPropertiesContainers)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Resources() BatchJobDefinitionEksPropertiesPodPropertiesContainersResourcesOutputReference
	ResourcesInput() *BatchJobDefinitionEksPropertiesPodPropertiesContainersResources
	SecurityContext() BatchJobDefinitionEksPropertiesPodPropertiesContainersSecurityContextOutputReference
	SecurityContextInput() *BatchJobDefinitionEksPropertiesPodPropertiesContainersSecurityContext
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VolumeMounts() BatchJobDefinitionEksPropertiesPodPropertiesContainersVolumeMountsList
	VolumeMountsInput() interface{}
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
	PutEnv(value interface{})
	PutResources(value *BatchJobDefinitionEksPropertiesPodPropertiesContainersResources)
	PutSecurityContext(value *BatchJobDefinitionEksPropertiesPodPropertiesContainersSecurityContext)
	PutVolumeMounts(value interface{})
	ResetArgs()
	ResetCommand()
	ResetEnv()
	ResetImagePullPolicy()
	ResetName()
	ResetResources()
	ResetSecurityContext()
	ResetVolumeMounts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference
type jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) Env() BatchJobDefinitionEksPropertiesPodPropertiesContainersEnvList {
	var returns BatchJobDefinitionEksPropertiesPodPropertiesContainersEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ImagePullPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ImagePullPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) InternalValue() *BatchJobDefinitionEksPropertiesPodPropertiesContainers {
	var returns *BatchJobDefinitionEksPropertiesPodPropertiesContainers
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) Resources() BatchJobDefinitionEksPropertiesPodPropertiesContainersResourcesOutputReference {
	var returns BatchJobDefinitionEksPropertiesPodPropertiesContainersResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ResourcesInput() *BatchJobDefinitionEksPropertiesPodPropertiesContainersResources {
	var returns *BatchJobDefinitionEksPropertiesPodPropertiesContainersResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) SecurityContext() BatchJobDefinitionEksPropertiesPodPropertiesContainersSecurityContextOutputReference {
	var returns BatchJobDefinitionEksPropertiesPodPropertiesContainersSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) SecurityContextInput() *BatchJobDefinitionEksPropertiesPodPropertiesContainersSecurityContext {
	var returns *BatchJobDefinitionEksPropertiesPodPropertiesContainersSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) VolumeMounts() BatchJobDefinitionEksPropertiesPodPropertiesContainersVolumeMountsList {
	var returns BatchJobDefinitionEksPropertiesPodPropertiesContainersVolumeMountsList
	_jsii_.Get(
		j,
		"volumeMounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) VolumeMountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeMountsInput",
		&returns,
	)
	return returns
}


func NewBatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference {
	_init_.Initialize()

	if err := validateNewBatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.batchJobDefinition.BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference_Override(b BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.batchJobDefinition.BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference)SetImagePullPolicy(val *string) {
	if err := j.validateSetImagePullPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullPolicy",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference)SetInternalValue(val *BatchJobDefinitionEksPropertiesPodPropertiesContainers) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) PutEnv(value interface{}) {
	if err := b.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEnv",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) PutResources(value *BatchJobDefinitionEksPropertiesPodPropertiesContainersResources) {
	if err := b.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putResources",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) PutSecurityContext(value *BatchJobDefinitionEksPropertiesPodPropertiesContainersSecurityContext) {
	if err := b.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSecurityContext",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) PutVolumeMounts(value interface{}) {
	if err := b.validatePutVolumeMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putVolumeMounts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		b,
		"resetArgs",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		b,
		"resetCommand",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		b,
		"resetEnv",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ResetImagePullPolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetImagePullPolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		b,
		"resetName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		b,
		"resetResources",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		b,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ResetVolumeMounts() {
	_jsii_.InvokeVoid(
		b,
		"resetVolumeMounts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

