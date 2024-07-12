package filters

import (
	filters_interface "github.com/nuno-bastos/gin-gonic-wire-api/service/filters/interface"

	"github.com/google/wire"
)

var AllowReportsFilterSet = wire.NewSet(
	NewAllowReportsFilter,
	// wire.Bind(new(filters_interface.IManagerFilter), new(*AllowReportsFilter)),
)

var DenyManagersFilterSet = wire.NewSet(
	NewDenyManagersFilter,
	// wire.Bind(new(filters_interface.IManagerFilter), new(*DenyManagersFilter)),
)

var DenyProfileLevelsFilterSet = wire.NewSet(
	NewDenyProfileLevelsFilter,
	// wire.Bind(new(filters_interface.IManagerFilter), new(*DenyProfileLevelsFilter)),
)

var AllowSelfFilterSet = wire.NewSet(
	NewAllowSelfFilter,
	// wire.Bind(new(filters_interface.IEmployeeFilter), new(*AllowSelfFilter)),
)

var DenySelfFilterSet = wire.NewSet(
	NewDenySelfFilter,
	// wire.Bind(new(filters_interface.IEmployeeFilter), new(*DenySelfFilter)),
)

// ProvideManagerFilters returns a slice of IManagerFilter instances.
// It aggregates the individual manager filters into a single slice.
//
// Parameters:
// - allowReports: an instance of AllowReportsFilter.
// - denyManagers: an instance of DenyManagersFilter.
// - denyProfileLevels: an instance of DenyProfileLevelsFilter.
//
// Returns:
// - []filters_interface.IManagerFilter: a slice of manager filters.
func ProvideManagerFilters(
	allowReports *AllowReportsFilter,
	denyManagers *DenyManagersFilter,
	denyProfileLevels *DenyProfileLevelsFilter,
) []filters_interface.IManagerFilter {
	return []filters_interface.IManagerFilter{
		allowReports,
		denyManagers,
		denyProfileLevels,
	}
}

// ProvideEmployeeFilters returns a slice of IEmployeeFilter instances.
// It aggregates the individual employee filters into a single slice.
//
// Parameters:
// - allowSelf: an instance of AllowSelfFilter.
// - denySelf: an instance of DenySelfFilter.
//
// Returns:
// - []filters_interface.IEmployeeFilter: a slice of employee filters.
func ProvideEmployeeFilters(
	allowSelf *AllowSelfFilter,
	denySelf *DenySelfFilter,
) []filters_interface.IEmployeeFilter {
	return []filters_interface.IEmployeeFilter{
		allowSelf,
		denySelf,
	}
}

var ManagerFiltersSet = wire.NewSet(
	AllowReportsFilterSet,
	DenyManagersFilterSet,
	DenyProfileLevelsFilterSet,
	ProvideManagerFilters,
)

var EmployeeFiltersSet = wire.NewSet(
	AllowSelfFilterSet,
	DenySelfFilterSet,
	ProvideEmployeeFilters,
)

var SecurityFiltersSet = wire.NewSet(
	ManagerFiltersSet,
	EmployeeFiltersSet,
)
