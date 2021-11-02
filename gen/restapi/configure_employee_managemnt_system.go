// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/restapi/operations"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/restapi/operations/admin"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/restapi/operations/employee"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/restapi/operations/team_lead"
)

//go:generate swagger generate server --target ../../gen --name EmployeeManagemntSystem --spec ../../swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.EmployeeManagemntSystemAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.EmployeeManagemntSystemAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AdminAddEmployeeHandler == nil {
		api.AdminAddEmployeeHandler = admin.AddEmployeeHandlerFunc(func(params admin.AddEmployeeParams) middleware.Responder {
			return middleware.NotImplemented("operation admin.AddEmployee has not yet been implemented")
		})
	}
	if api.AdminDeleteEmployeeHandler == nil {
		api.AdminDeleteEmployeeHandler = admin.DeleteEmployeeHandlerFunc(func(params admin.DeleteEmployeeParams) middleware.Responder {
			return middleware.NotImplemented("operation admin.DeleteEmployee has not yet been implemented")
		})
	}
	if api.AdminShowEmployeeHandler == nil {
		api.AdminShowEmployeeHandler = admin.ShowEmployeeHandlerFunc(func(params admin.ShowEmployeeParams) middleware.Responder {
			return middleware.NotImplemented("operation admin.ShowEmployee has not yet been implemented")
		})
	}
	if api.EmployeeShowEmployeeSelfHandler == nil {
		api.EmployeeShowEmployeeSelfHandler = employee.ShowEmployeeSelfHandlerFunc(func(params employee.ShowEmployeeSelfParams) middleware.Responder {
			return middleware.NotImplemented("operation employee.ShowEmployeeSelf has not yet been implemented")
		})
	}
	if api.TeamLeadShowEmployeeTeamHandler == nil {
		api.TeamLeadShowEmployeeTeamHandler = team_lead.ShowEmployeeTeamHandlerFunc(func(params team_lead.ShowEmployeeTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation team_lead.ShowEmployeeTeam has not yet been implemented")
		})
	}
	if api.AdminUpdateEmployeeHandler == nil {
		api.AdminUpdateEmployeeHandler = admin.UpdateEmployeeHandlerFunc(func(params admin.UpdateEmployeeParams) middleware.Responder {
			return middleware.NotImplemented("operation admin.UpdateEmployee has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
