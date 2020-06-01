 
package logger

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"

	"log"
	"os"
	"github.com/pkg/errors"
)

const channelName = "github.com/jWinterDay/desktop_logger"

type LoggerFlutterPlugin struct {
	channel *plugin.MethodChannel
}

var _ flutter.Plugin = &LoggerFlutterPlugin{}

func (p *LoggerFlutterPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	p.channel = plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	p.channel.HandleFunc("log", handleLog)

	return nil
}

func handleLog(arguments interface{}) (reply interface{}, err error) {
	var ok bool
	var argsMap map[interface{}]interface{}
	if argsMap, ok = arguments.(map[interface{}]interface{}); !ok {
		return nil, errors.New("invalid arguments")
	}

	// params
	var pathParam string
	var prefixParam string
	var textParam string

	if path, ok := argsMap["path"]; ok {
		pathParam = path.(string)
	}

	if prefix, ok := argsMap["prefix"]; ok {
		prefixParam = prefix.(string)
	}

	if text, ok := argsMap["text"]; ok {
		textParam = text.(string)
	}

	f, err := os.OpenFile(pathParam, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	logger := log.New(f, prefixParam, log.LstdFlags)

	logger.Println(textParam)

	return nil, nil
}
