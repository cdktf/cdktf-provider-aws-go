package ecstaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/ecstaskdefinition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference interface {
	cdktf.ComplexObject
	AuthorizationConfig() EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfigOutputReference
	AuthorizationConfigInput() *EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig
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
	FileSystemId() *string
	SetFileSystemId(val *string)
	FileSystemIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfiguration
	SetInternalValue(val *EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfiguration)
	RootDirectory() *string
	SetRootDirectory(val *string)
	RootDirectoryInput() *string
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
	PutAuthorizationConfig(value *EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference
type jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) AuthorizationConfig() EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfigOutputReference {
	var returns EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfigOutputReference
	_jsii_.Get(
		j,
		"authorizationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) AuthorizationConfigInput() *EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig {
	var returns *EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig
	_jsii_.Get(
		j,
		"authorizationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) FileSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) InternalValue() *EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfiguration {
	var returns *EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) RootDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) RootDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsTaskDefinition.EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference_Override(e EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsTaskDefinition.EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference)SetFileSystemId(val *string) {
	if err := j.validateSetFileSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemId",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference)SetInternalValue(val *EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference)SetRootDirectory(val *string) {
	if err := j.validateSetRootDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootDirectory",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) PutAuthorizationConfig(value *EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig) {
	if err := e.validatePutAuthorizationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAuthorizationConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

