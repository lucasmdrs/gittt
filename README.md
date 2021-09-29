# GITTT

This is a utility package to handle multiple github events a create multiple triggers in a single webhook endpoint.

## Getting started

### Using Predefined Actions and Conditions

```golang
import (
    "log"
    "net/http"

    "github.com/lucasmdrs/gittt"
)

func main() {
    g := gittt.Init()

    g.ListenAllEvents()

    always := g.ConditionBuilder(gittt.AnyEvent, g.ConditionAlways, nil)

    logPayload := g.ActionBuilder(g.ActionLogPayload, nil)

    always.AddAction(logPayload)

    g.AddConditions(always)

    http.HandleFunc("/webhook", g.Handler)

    log.Println("Wait for connections on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

```

### Creating Your Own Actions and Conditions

```golang
package main

import (
    "log"
    "net/http"
    "os/exec"
    "strings"

    "github.com/lucasmdrs/gittt"
)

func main() {
    g := gittt.Init()

    err := g.ListenForEvents(gittt.ReleaseEvent)
    if err != nil {
        log.Fatal(err.Error())
    }

    releaseNameContainsKeywords := func(data interface{}, keywords ...interface{}) bool {
        if release, ok := data.(gittt.Release); ok {
            for _, keyword := range keywords {
                return strings.Contains(release.ReleaseInfo.Name, keyword.(string))
            }
        }
        return false
    }

    onReleaseWithDeployInName := g.ConditionBuilder(gittt.ReleaseEvent, releaseNameContainsKeywords, "deploy")

    callScript := func(data interface{}, scriptNames ...interface{}) {
        for _, script := range scriptNames {
            cmd := exec.Command("/bin/bash", script.(string))
            err := cmd.Run()
            if err != nil {
                log.Fatal(err)
            }
        }

    }

    runMyScript := g.ActionBuilder(callScript, "my_script.sh")

    onReleaseWithDeployInName.AddAction(runMyScript)

    g.AddConditions(onReleaseWithDeployInName)

    http.HandleFunc("/webhook", g.Handler)

    log.Println("Wait for connections on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```
