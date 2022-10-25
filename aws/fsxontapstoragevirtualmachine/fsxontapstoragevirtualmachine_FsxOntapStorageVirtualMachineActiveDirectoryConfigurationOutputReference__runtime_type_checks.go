//go:build !no_runtime_type_checking

package fsxontapstoragevirtualmachine

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validatePutSelfManagedActiveDirectoryConfigurationParameters(value *FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateSetInternalValueParameters(val *FsxOntapStorageVirtualMachineActiveDirectoryConfiguration) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateSetNetbiosNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewFsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

