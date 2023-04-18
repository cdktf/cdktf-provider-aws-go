package fsxfilecache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/fsxfilecache/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxFileCacheLustreConfigurationMetadataConfigurationList interface {
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
	Get(index *float64) FsxFileCacheLustreConfigurationMetadataConfigurationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxFileCacheLustreConfigurationMetadataConfigurationList
type jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFsxFileCacheLustreConfigurationMetadataConfigurationList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FsxFileCacheLustreConfigurationMetadataConfigurationList {
	_init_.Initialize()

	if err := validateNewFsxFileCacheLustreConfigurationMetadataConfigurationListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList{}

	_jsii_.Create(
		"@cdktf/provider-aws.fsxFileCache.FsxFileCacheLustreConfigurationMetadataConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFsxFileCacheLustreConfigurationMetadataConfigurationList_Override(f FsxFileCacheLustreConfigurationMetadataConfigurationList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fsxFileCache.FsxFileCacheLustreConfigurationMetadataConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList) Get(index *float64) FsxFileCacheLustreConfigurationMetadataConfigurationOutputReference {
	if err := f.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns FsxFileCacheLustreConfigurationMetadataConfigurationOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (f *jsiiProxy_FsxFileCacheLustreConfigurationMetadataConfigurationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

