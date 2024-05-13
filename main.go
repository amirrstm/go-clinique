package main

import (
	"github.com/amirrstm/go-clinique/cmd/server"
	"github.com/amirrstm/go-clinique/pkg/config"

	_ "github.com/amirrstm/go-clinique/docs" // load API Docs files (Swagger)
)

//	@title			Fiber Go Auth
//	@version		1.0
//	@description	Fiber go web framework based REST API boilerplate
//	@contact.name	A. Rostami
//	@contact.email	amr.rostam@gmail.com
//	@termsOfService
//	@license.name				MIT
//	@license.url				https://opensource.org/licenses/MIT
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@host						localhost:9000
//	@BasePath					/api
func main() {

	// setup various configuration for app
	config.LoadAllConfigs(".env")

	server.Serve()

}
