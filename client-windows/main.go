package main

import (
    "log"

    "github.com/wailsapp/wails/v2"
    "github.com/wailsapp/wails/v2/pkg/options"
    "smart-vpn/client-windows/go"
)

func main() {
    app := go.NewApp()

    err := wails.Run(&options.App{
        Title:  "SmartVPN",
        Width:  600,
        Height: 500,
        Bind: []interface{}{
            app,
        },
    })

    if err != nil {
        log.Fatal(err)
    }
}