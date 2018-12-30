package dry

import (
	"github.com/hbagdi/go-kong/kong"
	"github.com/kong/deck/crud"
	"github.com/kong/deck/diff"
	"github.com/kong/deck/print"
	"github.com/kong/deck/state"
	"github.com/kong/deck/utils"
)

// RouteCRUD implements Actions interface
// from the github.com/kong/crud package for the Route entitiy of Kong.
type RouteCRUD struct {
	// client    *kong.Client
	// callbacks []Callback // use this to update the current in-memory state
}

func routeFromStuct(arg diff.Event) *state.Route {
	route, ok := arg.Obj.(*state.Route)
	if !ok {
		panic("unexpected type, expected *state.Route")
	}

	return route
}

// Create creates a Route in Kong. TODO Doc
func (s *RouteCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	route := routeFromStuct(event)
	print.CreatePrintln("creating route ", *route.Name)
	route.ID = kong.String(utils.UUID())
	return route, nil
}

// Delete deletes a Route in Kong. TODO Doc
func (s *RouteCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	route := routeFromStuct(event)
	print.DeletePrintln("deleting route ", *route.Name)
	return route, nil
}

// Update updates a Route in Kong. TODO Doc
func (s *RouteCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	route := routeFromStuct(event)
	oldRoute, ok := event.OldObj.(*state.Route)
	if !ok {
		panic("unexpected type, expected *state.Route")
	}
	// TODO remove this hack
	oldRoute.CreatedAt = nil
	oldRoute.UpdatedAt = nil
	oldRoute.Service = &kong.Service{Name: oldRoute.Service.Name}
	oldRoute.ID = nil

	route.ID = nil
	route.Service = &kong.Service{Name: route.Service.Name}

	diff := getDiff(oldRoute.Route, route.Route)
	print.UpdatePrintln("updating route", *route.Name)
	print.UpdatePrintf("%s", diff)
	return route, nil
}