
# Fluxfile

  

**Fluxfile** is a command-line application to sync files of local-first applications with Git and other locations.

## Usage

**Fluxfile** provides several commands to manage and sync your projects. Below are the available commands:

  

### Init

Initialize a directory for syncing with a Git repository.
``` bash
fluxfile init
```
### Sync

Sync the local directory with the specified GitHub repository.

  
``` bash
fluxfile sync [identifier]
```
  

>Use the --all flag to sync all projects.

  
``` bash
fluxfile sync --all
```
  

### List

List all initialized projects.

  
``` bash
fluxfile list
```
  

### Edit

Edit details of a specific project.

  
``` bash
fluxfile edit [identifier]
```
  

### Remove

Remove a specific project.

  
``` bash
fluxfile remove [identifier]
```

## Configuration

Fluxfile uses a configuration file located at ~/.fluxfile/config.yaml (or %APPDATA%\fluxfile\config.yaml on Windows). This file stores the details of all projects.

  

Example Configuration

  
```yml
projects:
- project_name: myproject
  directory: /path/to/myproject
  git_url: https://github.com/username/myproject.git
  branch_name: main


ssh_enabled: true

```
  
  

## Contributing

We welcome contributions to Fluxfile! Please fork the repository and submit pull requests

