package main

import (
    "fmt"
    "os"
    "time"

    hook "github.com/robotn/gohook"
)

func main() {
    asynchronous_listener()
}

func asynchronous_listener() {
    fmt.Println("Press q to quit.")

    // Register the keydown event for 'q'
    hook.Register(hook.KeyDown, []string{"q"}, func(e hook.Event) {
        os.Exit(3)
    })

    // Start the hook in a separate goroutine
    go func() {
        s := hook.Start()
        <-hook.Process(s)
    }()

    // Main goroutine keeps printing
    for {
        fmt.Print(".")
        time.Sleep(500 * time.Millisecond)
    }
}

