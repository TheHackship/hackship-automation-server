package main

import (
	"encoding/json"
	"log"
	"os"
	"script_trigger_server/runner"
	"script_trigger_server/server"
)

func configFileParser(filePath string) (map[string]map[string]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	serviceDir := make(map[string]map[string]string)

	if err := json.Unmarshal(data, &serviceDir); err != nil {
		return nil, err
	}

	return serviceDir, nil
}

func main() {
	scriptChan := make(chan string)

	// Parse config file
	configMap, err := configFileParser("./config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Start script runner
	r := runner.NewRunner(scriptChan)
	go r.Start()

	// Start HTTP server in goroutine
	s := server.NewServer(configMap, scriptChan)
	go s.Start()

	select {}
}
