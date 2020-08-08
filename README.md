# Plex Hooks

Library to listen [Plex Webhook Events](https://support.plex.tv/articles/115002267687-webhooks/) and execute actions.

**Plex Webhooks require a Plex Pass subscription**

## Installation

Use the library in your own app to register your actions.
```bash
go get github.com/acamilleri/go-plexhooks
```

## Usage

```golang
package main

import (
    "net"

    "github.com/sirupsen/logrus"
    "github.com/acamilleri/go-plexhooks"
    "github.com/acamilleri/go-plexhooks/plex"
)

type MyActionOnMediaPlay struct{}

func (a *MyActionOnMediaPlay) Name() string {
	return "MyActionOnMediaPlay"
}

func (a *MyActionOnMediaPlay) Execute(event plex.Event) error {
	fmt.Printf("I'm print a message when i received a MediaPlay event")
	return nil
}

func main() {
	listenAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1")
	if err != nil {
		panic(err)
	}

	actions := plexhooks.NewActions()
	actions.Add(plex.MediaPlay, &MyActionOnMediaPlay{})

	app := plexhooks.New(plexhooks.Definition{
		ListenAddr: listenAddr,
		Actions:    actions,
		Logger: plexhooks.LoggerDefinition{
			Level:     logrus.InfoLevel,
			Formatter: &logrus.JSONFormatter{},
		},
	})
    
    err := app.Run()
    if err != nil {
        panic(err)
    }
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
