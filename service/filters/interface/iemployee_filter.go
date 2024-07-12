package filters_interface

// IEmployeeFilter defines the interface for employee-specific security matrix calculation filters.
// It embeds the SecurityMatrixCalculationFilter interface, meaning it inherits its method(s).
// These filters create rules specific to employee access control.
type IEmployeeFilter interface {
	SecurityMatrixCalculationFilter
}
