package timestreamwrite


type TimestreamwriteTableRetentionProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/timestreamwrite_table#magnetic_store_retention_period_in_days TimestreamwriteTable#magnetic_store_retention_period_in_days}.
	MagneticStoreRetentionPeriodInDays *float64 `field:"required" json:"magneticStoreRetentionPeriodInDays" yaml:"magneticStoreRetentionPeriodInDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/timestreamwrite_table#memory_store_retention_period_in_hours TimestreamwriteTable#memory_store_retention_period_in_hours}.
	MemoryStoreRetentionPeriodInHours *float64 `field:"required" json:"memoryStoreRetentionPeriodInHours" yaml:"memoryStoreRetentionPeriodInHours"`
}

