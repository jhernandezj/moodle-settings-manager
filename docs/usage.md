# Usage Guide

## Overview

The Moodle Settings Manager is designed to provide a simple and efficient way to manage Moodle settings programmatically. This document outlines the basic usage and features of the module.

## Basic Usage

### Initializing the Manager

```go
import "github.com/jhernandezj/moodle-settings-manager/pkg/settings"

config := // your configuration
manager := settings.New(config)
```

### Getting Settings

```go
value, err := manager.GetSetting("setting_name")
if err != nil {
    // handle error
}
```

### Updating Settings

```go
err := manager.UpdateSetting("setting_name", "new_value")
if err != nil {
    // handle error
}
```

## Configuration

The module requires basic configuration for database connectivity. See the configuration guide for more details.

## Best Practices

1. Always handle errors appropriately
2. Use environment variables for sensitive configuration
3. Implement proper logging in production environments
