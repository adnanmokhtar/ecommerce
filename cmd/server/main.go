package main

import (
	"net/http"
	"os"
	"path/filepath"
	"reflect"

	"github.com/adnanmokhtar/ecommerce/internal/modules"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// Dynamically load routes from all modules
	loadModuleRoutes(r)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// loadModuleRoutes dynamically loads routes from all modules
func loadModuleRoutes(r *gin.Engine) {
	modulesDir := "./internal/modules"

	// Find all subdirectories in the modules directory
	subDirs, err := findSubdirectories(modulesDir)
	if err != nil {
		panic(err)
	}

	for _, subDir := range subDirs {
		modulePath := filepath.Join(modulesDir, subDir)

		// Check if the module follows a specific directory structure
		if isModuleDirectory(modulePath) {
			// Load routes for the module
			loadRoutesForModule(r, modulePath)
		}
	}
}

// findSubdirectories returns a list of subdirectories in a given directory
func findSubdirectories(dir string) ([]string, error) {
	var subDirs []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != dir {
			subDirs = append(subDirs, filepath.Base(path))
		}
		return nil
	})
	return subDirs, err
}

// isModuleDirectory checks if a directory follows a specific module structure
func isModuleDirectory(modulePath string) bool {
	// Customize this function to check if the modulePath contains the expected structure
	// For example, check for the presence of a "presentation/http/routes.go" file.
	routesFilePath := filepath.Join(modulePath, "presentation/http/routes.go")
	_, err := os.Stat(routesFilePath)
	return !os.IsNotExist(err)
}

// loadRoutesForModule loads routes for a specific module
func loadRoutesForModule(r *gin.Engine, modulePath string) {
	// Use reflection or other mechanisms to load routes from the module
	// You can follow the previous example to load routes dynamically.
	// Here, we assume each module has a LoadRoutes function.
	module := reflect.New(reflect.TypeOf(modules.Module{})).Interface()

	if routesLoader, ok := module.(func(*gin.RouterGroup)); ok {
		// Create a new router group for the module
		moduleRouter := r.Group("/")
		routesLoader(moduleRouter)
	}
}
