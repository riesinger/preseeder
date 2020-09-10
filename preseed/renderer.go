package preseed

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

const (
	defaultPreseedPath = "./preseeds"
	// DefaultPreseedName is the name of the default template
	// It will be used when no template for the given hostname is found
	DefaultPreseedName = "_default_"
)

// Renderer for preseeds. This handles loading the respective templates and populating their data
type Renderer struct {
	// BaseDirectory containing the preseed files
	// If this is not defined when creating a renderer, it will fall back to the './preseeds' directory
	BaseDirectory string
}

func (r *Renderer) preseedDirectory() string {
	if r.BaseDirectory != "" {
		return r.BaseDirectory
	}
	return defaultPreseedPath
}

func (r *Renderer) templateForHostname(hostname string) (*os.File, error) {
	file, err := os.Open(filepath.Join(r.preseedDirectory(), hostname))
	if err == nil {
		return file, nil
	} else if os.IsNotExist(err) {
		return os.Open(r.DefaultPreseedPath())
	} else {
		return nil, err
	}
}

// GetForHostname is used to render a preseed file for a given hostname
func (r *Renderer) GetForHostname(hostname string) ([]byte, error) {
	log.Printf("Getting preseed for hostname '%s'", hostname)
	file, err := r.templateForHostname(hostname)
	if err != nil {
		return nil, fmt.Errorf("renderer: could not get template: %w", err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("renderer: cannot read template file: %w", err)
	}
	tmpl, err := template.New(hostname).Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("renderer: cannot parse template: %w", err)
	}
	sb := strings.Builder{}
	if err := tmpl.Execute(&sb, Data{
		Hostname:   hostname,
		RenderedAt: time.Now().UTC(),
	}); err != nil {
		return nil, fmt.Errorf("Error rendering template: %w", err)
	}
	return []byte(sb.String()), nil
}

// DefaultPreseedExists checks whether the default preseed file exists and is readable
func (r *Renderer) DefaultPreseedExists() bool {
	if _, err := os.Stat(r.DefaultPreseedPath()); err != nil {
		return false
	}
	return true
}

// DefaultPreseedPath returns the name (and directory) of the default preseed file
func (r *Renderer) DefaultPreseedPath() string {
	return filepath.Join(r.preseedDirectory(), DefaultPreseedName)
}
