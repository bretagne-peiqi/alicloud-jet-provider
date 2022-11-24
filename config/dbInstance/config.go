package dbInstance

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_db_instance", func(r *config.Resource) {

		//r.ShortGroup = "db"
		r.Kind = "Mysql"
		r.ExternalName = config.IdentifierFromProvider
		r.UseAsync = true
	})
}
