//go:build !no_runtime_type_checking

package quicksightdataset

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validatePutCastColumnTypeOperationParameters(value *QuicksightDataSetLogicalTableMapDataTransformsCastColumnTypeOperation) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validatePutCreateColumnsOperationParameters(value *QuicksightDataSetLogicalTableMapDataTransformsCreateColumnsOperation) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validatePutFilterOperationParameters(value *QuicksightDataSetLogicalTableMapDataTransformsFilterOperation) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validatePutProjectOperationParameters(value *QuicksightDataSetLogicalTableMapDataTransformsProjectOperation) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validatePutRenameColumnOperationParameters(value *QuicksightDataSetLogicalTableMapDataTransformsRenameColumnOperation) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validatePutTagColumnOperationParameters(value *QuicksightDataSetLogicalTableMapDataTransformsTagColumnOperation) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validatePutUntagColumnOperationParameters(value *QuicksightDataSetLogicalTableMapDataTransformsUntagColumnOperation) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case *QuicksightDataSetLogicalTableMapDataTransforms:
		val := val.(*QuicksightDataSetLogicalTableMapDataTransforms)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case QuicksightDataSetLogicalTableMapDataTransforms:
		val_ := val.(QuicksightDataSetLogicalTableMapDataTransforms)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case cdktf.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *QuicksightDataSetLogicalTableMapDataTransforms, cdktf.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewQuicksightDataSetLogicalTableMapDataTransformsOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

