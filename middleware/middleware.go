package middleware

import (
	"myapp/data"

	"github.com/bahadorfarahani/celeritas"
)

type Middleware struct {
	App *celeritas.Celeritas
	Models data.Models
}
