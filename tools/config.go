package tools

import (
    "errors"
    "gopkg.in/yaml.v2"
    "io"
    "os"
    "path/filepath"
    "runtime"
)

type Project struct {
    ProjectName string `yaml:"project_name"`
    Directory   string `yaml:"directory"`
    GitURL      string `yaml:"git_url"`
    BranchName  string `yaml:"branch_name"`
}

type Config struct {
    Projects []Project `yaml:"projects"`
}

var configFilePath string

func getConfigFilePath() string {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        panic("Unable to find the home directory")
    }

    switch runtime.GOOS {
    case "linux", "darwin":
        return filepath.Join(homeDir, ".fluxfile", "config.yaml")
    case "windows":
        appData := os.Getenv("APPDATA")
        if appData == "" {
            panic("APPDATA environment variable is not set on Windows")
        }
        return filepath.Join(appData, "fluxfile", "config.yaml")
    default:
        panic("Unsupported OS")
    }
}

func init() {
    configFilePath = getConfigFilePath()
}

func LoadConfig() (*Config, error) {
    var config Config
    if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
        return &config, nil 
    }
    
    file, err := os.Open(configFilePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }

    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }
    
    return &config, nil
}

func SaveConfig(config *Config) error {
    data, err := yaml.Marshal(config)
    if err != nil {
        return err
    }

    configDir := filepath.Dir(configFilePath)
    // Create the directory if directory is missing!!!
    if _, err := os.Stat(configDir); os.IsNotExist(err) {
        os.Mkdir(configDir, os.ModePerm) 
    }

    file, err := os.Create(configFilePath)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.Write(data)
    return err
}

func FindProjectByIdentifier(config *Config, identifier string) (*Project, int) {
    for index, project := range config.Projects {
        if project.ProjectName == identifier {
            return &project, index
        }
    }
    return nil, -1
}

func RemoveProject(config *Config, identifier string) error {
    _, index := FindProjectByIdentifier(config, identifier)
    if index == -1 {
        return errors.New("project not found")
    }
    
    config.Projects = append(config.Projects[:index], config.Projects[index+1:]...)
    return SaveConfig(config)
}