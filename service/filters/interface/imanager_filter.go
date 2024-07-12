package filters_interface

// IManagerFilter defines the interface for manager-specific security matrix calculation filters.
// It embeds the SecurityMatrixCalculationFilter interface, meaning it inherits its methods.
// These filters create rules specific to manager access control.
type IManagerFilter interface {
	SecurityMatrixCalculationFilter
}
