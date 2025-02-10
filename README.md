# TDWS Demo module in Go

This is a demo module for TDWS written in Go.

## Information

All workflows and activities needs to be defined in the `main.go`, `metadata.json`, `go.mod` and `go.sum` files. This is becuase TDWS imports the module as a go plugin.

All modules need to have the function called `TdwsRegister`, which is called by TDWS to register the workflows and activities.
The TdwsRegister function should take temporal worker as an argument and register the workflows and activities to it.

The `metadata.json` file should be located in the root directory. This file should contain the following fields:

- name - the name of the module
- description - a short description of the module
- creator - the creator of the module
- version - the version of the module

### Example metadata.json

```json
{
  "name": "TDWS Demo Module",
  "description": "This is a demo module for TDWS",
  "creator": "TDWS",
  "version": "1.0.0"
}
```

### Example TdwsRegister function

```go
package main

import (
	"go.temporal.io/sdk/worker"
)

func TdwsRegister(worker worker.Worker) {
    worker.RegisterWorkflow(ExampleWorkflow)
    worker.RegisterActivity(ExampleActivity)
}
```
