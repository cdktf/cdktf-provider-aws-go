//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package gluecrawler

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueCrawlerJdbcTargetList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueCrawlerJdbcTargetList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueCrawlerJdbcTargetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueCrawlerJdbcTargetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueCrawlerJdbcTargetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueCrawlerJdbcTargetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueCrawlerJdbcTargetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

