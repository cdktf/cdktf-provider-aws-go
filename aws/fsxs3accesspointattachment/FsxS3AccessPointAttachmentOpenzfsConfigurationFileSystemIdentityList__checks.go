// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package fsxs3accesspointattachment

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityList) validateResolveParameters(context cdktf.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentity:
		val := val.(*[]*FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentity)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentity:
		val_ := val.([]*FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentity)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentity; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewFsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

