// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package securityhubautomationrule

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validatePutNoteParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SecurityhubAutomationRuleActionsFindingFieldsUpdateNote:
		value := value.(*[]*SecurityhubAutomationRuleActionsFindingFieldsUpdateNote)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubAutomationRuleActionsFindingFieldsUpdateNote:
		value_ := value.([]*SecurityhubAutomationRuleActionsFindingFieldsUpdateNote)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SecurityhubAutomationRuleActionsFindingFieldsUpdateNote; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validatePutRelatedFindingsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SecurityhubAutomationRuleActionsFindingFieldsUpdateRelatedFindings:
		value := value.(*[]*SecurityhubAutomationRuleActionsFindingFieldsUpdateRelatedFindings)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubAutomationRuleActionsFindingFieldsUpdateRelatedFindings:
		value_ := value.([]*SecurityhubAutomationRuleActionsFindingFieldsUpdateRelatedFindings)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SecurityhubAutomationRuleActionsFindingFieldsUpdateRelatedFindings; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validatePutSeverityParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SecurityhubAutomationRuleActionsFindingFieldsUpdateSeverity:
		value := value.(*[]*SecurityhubAutomationRuleActionsFindingFieldsUpdateSeverity)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubAutomationRuleActionsFindingFieldsUpdateSeverity:
		value_ := value.([]*SecurityhubAutomationRuleActionsFindingFieldsUpdateSeverity)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SecurityhubAutomationRuleActionsFindingFieldsUpdateSeverity; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validatePutWorkflowParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SecurityhubAutomationRuleActionsFindingFieldsUpdateWorkflow:
		value := value.(*[]*SecurityhubAutomationRuleActionsFindingFieldsUpdateWorkflow)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubAutomationRuleActionsFindingFieldsUpdateWorkflow:
		value_ := value.([]*SecurityhubAutomationRuleActionsFindingFieldsUpdateWorkflow)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SecurityhubAutomationRuleActionsFindingFieldsUpdateWorkflow; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateResolveParameters(context cdktf.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateSetConfidenceParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateSetCriticalityParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *SecurityhubAutomationRuleActionsFindingFieldsUpdate:
		val := val.(*SecurityhubAutomationRuleActionsFindingFieldsUpdate)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case SecurityhubAutomationRuleActionsFindingFieldsUpdate:
		val_ := val.(SecurityhubAutomationRuleActionsFindingFieldsUpdate)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *SecurityhubAutomationRuleActionsFindingFieldsUpdate; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateSetTypesParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateSetUserDefinedFieldsParameters(val *map[string]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) validateSetVerificationStateParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewSecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

