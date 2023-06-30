package fsxopenzfsfilesystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/fsxopenzfsfilesystem/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList
type jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList {
	_init_.Initialize()

	if err := validateNewFsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList{}

	_jsii_.Create(
		"@cdktf/provider-aws.fsxOpenzfsFileSystem.FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList_Override(f FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fsxOpenzfsFileSystem.FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList) Get(index *float64) FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasOutputReference {
	if err := f.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

