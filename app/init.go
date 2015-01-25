// @APIVersion 1.0.0
// @APITitle My Cool API
// @APIDescription My API usually works as expected. But sometimes it's not true
// @Contact api@contact.me
// @TermsOfServiceUrl http://google.com/
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause

package app

import (
	"os"
	"github.com/robfig/revel"
	"licorne/utilities/helper"
	"licorne/utilities/mongo"
	"licorne/utilities/tracelog"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	revel.OnAppStart(initApp)
	// revel.OnAppStart(FillCache())
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

// initApp contains all application level initialization
func initApp() {
	// Capture the global email settings
	// tracelog.EmailHost = revel.Config.StringDefault("email.host", "")
	// tracelog.EmailPort = revel.Config.IntDefault("email.port", 0)
	// tracelog.EmailUserName = revel.Config.StringDefault("email.username", "")
	// tracelog.EmailPassword = revel.Config.StringDefault("email.password", "")
	// tracelog.EmailTo = revel.Config.StringDefault("email.to", "")
	// tracelog.EmailAlertSubject = revel.Config.StringDefault("email.alert_subject", "")

	// MongoDB Settings
	helper.MONGO_DATABASE = revel.Config.StringDefault("mgo.use_database", "")

	// Init mongo
	err := mongo.Startup(helper.MAIN_GO_ROUTINE)
	if err != nil {
		tracelog.COMPLETED_ERROR(err, helper.MAIN_GO_ROUTINE, "initApp")
		os.Exit(1)
	}
}
