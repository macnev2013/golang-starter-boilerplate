package render

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/intrigues/golang-starter-boilerplate/core"
	"github.com/intrigues/golang-starter-boilerplate/internal/config"
	"github.com/intrigues/golang-starter-boilerplate/internal/global"
)

type render struct {
	cfg           *config.Config
	logger        core.Clog
	templateCache TemplateCacheType
}

// TemplateCacheType defines type of template cache
type TemplateCacheType map[string]*template.Template

func New(cfg *config.Config, logger core.Clog) core.Render {
	tc, _ := GenearteTemplateCache()
	return &render{
		cfg:           cfg,
		logger:        logger,
		templateCache: tc,
	}
}

// GenearteTemplateCache creates template cache map
func GenearteTemplateCache() (TemplateCacheType, error) {
	cache := TemplateCacheType{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/%s", global.TEMPLATE_DIR, global.PAGE_TEMPLATE))
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/%s", global.TEMPLATE_DIR, global.LAYOUT_TEMPLATE))
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/%s", global.TEMPLATE_DIR, global.LAYOUT_TEMPLATE))
			if err != nil {
				return cache, err
			}
		}
		cache[name] = ts
	}

	return cache, nil
}
