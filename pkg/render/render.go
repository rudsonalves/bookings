package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/rudsonalves/bookings/pkg/config"
	"github.com/rudsonalves/bookings/pkg/models"
)

const (
	TEMPL     = "./templates/"
	TEMPLBASE = "./templates/base.layout.tmpl"
)

var (
	app       *config.AppConfig
	functions = template.FuncMap{}
)

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData add defaults datas from page
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	// ...
	return td
}

// RenderTemplate renders templates using html/tamplate
func RenderTemplate(write http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		// create a new template cache
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatalf("Could not get template %q from tamplate cache", tmpl)
	}

	td = AddDefaultData(td)

	// render the template
	err := t.Execute(write, td)
	if err != nil {
		log.Printf("error parsing template: %s", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files name *.page.tmpl from ./templates
	pages, err := filepath.Glob(TEMPL + "*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range for all files in pages
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		log.Printf("create %s template", name)

		matches, err := filepath.Glob(TEMPL + "*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(TEMPL + "*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
