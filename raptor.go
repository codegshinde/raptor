package raptor

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Raptor represents a simple configuration manager
var configManager *raptor

type raptor struct {
	config map[string]string
	file   string
	mutex  sync.Mutex
}

// SetConfigFile sets the configuration file for Raptor
func SetConfigFile(fileName string) error {
	// Get the root directory where the executable is located
	rootDir, err := getRootDir()
	if err != nil {
		return err
	}

	// Join the root directory and the provided file name to get the full path
	filePath := filepath.Join(rootDir, fileName)
	return configManager.setConfigFile(filePath)
}

// getRootDir returns the directory containing the executable
func getRootDir() (string, error) {
	// Get the executable's path
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	// Return the directory of the executable
	return filepath.Dir(ex), nil
}

// setConfigFile sets the configuration file for Raptor using the provided file path
func (r *raptor) setConfigFile(filePath string) error {
	// Lock the mutex to ensure thread safety
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Read the configuration from the specified file
	env, err := r.readConfigFile(filePath)
	if err != nil {
		fmt.Printf("Error reading config file '%s': %s\n", filePath, err)
		return err
	}

	// Update the configuration and file path
	r.config = env
	r.file = filePath
	return nil
}

// readConfigFile reads environment variables from a file and returns them as a map
func (r *raptor) readConfigFile(filePath string) (map[string]string, error) {
	env := make(map[string]string)

	// Open the specified file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split each line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			env[key] = value
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return env, nil
}

// Get retrieves the value associated with a key from the configuration
func Get(key string) (string, error) {
	return configManager.get(key)
}

// get retrieves the value associated with a key from the configuration
func (r *raptor) get(key string) (string, error) {
	// Lock the mutex to ensure thread safety
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Retrieve the value for the specified key
	value, exists := r.config[key]
	if !exists {
		return "", fmt.Errorf("key '%s' not found in configuration", key)
	}
	return value, nil
}

// Set sets a key-value pair in the configuration
func Set(key, value string) {
	configManager.set(key, value)
}

// set sets a key-value pair in the configuration
func (r *raptor) set(key, value string) {
	// Lock the mutex to ensure thread safety
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Update the configuration with the new key-value pair
	r.config[key] = value
}

// Delete removes a key from the configuration
func Delete(key string) {
	configManager.delete(key)
}

// delete removes a key from the configuration
func (r *raptor) delete(key string) {
	// Lock the mutex to ensure thread safety
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Delete the specified key from the configuration
	delete(r.config, key)
}

// GetAll returns the entire configuration map
func GetAll() map[string]string {
	return configManager.getAll()
}

// getAll returns the entire configuration map
func (r *raptor) getAll() map[string]string {
	// Lock the mutex to ensure thread safety
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Return a copy of the configuration map
	return r.config
}

// Initialize the package-level configManager
func init() {
	// Create a new instance of the raptor configuration manager
	configManager = &raptor{}
}
