package raptor

import (
	"bufio"
	"fmt"
	"os"
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
func SetConfigFile(filePath string) error {
	return configManager.setConfigFile(filePath)
}

func (r *raptor) setConfigFile(filePath string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	env, err := r.readConfigFile(filePath)
	if err != nil {
		fmt.Printf("Error reading config file '%s': %s\n", filePath, err)
		return err
	}
	r.config = env
	r.file = filePath
	return nil
}

// ReadConfigFile reads environment variables from a file and returns them as a map.
func (r *raptor) readConfigFile(filePath string) (map[string]string, error) {
	env := make(map[string]string)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			env[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return env, nil
}

// Get retrieves the value associated with a key from the configuration
func Get(key string) (string, error) {
	return configManager.get(key)
}

func (r *raptor) get(key string) (string, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

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

func (r *raptor) set(key, value string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.config[key] = value
}

// Delete removes a key from the configuration
func Delete(key string) {
	configManager.delete(key)
}

func (r *raptor) delete(key string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	delete(r.config, key)
}

// GetAll returns the entire configuration map
func GetAll() map[string]string {
	return configManager.getAll()
}

func (r *raptor) getAll() map[string]string {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.config
}

// Initialize the package-level configManager
func init() {
	configManager = &raptor{}
}
