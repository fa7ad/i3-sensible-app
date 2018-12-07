package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var conf []byte

type workspace struct {
	Name    string
	Focused bool
}

func getEnv(env, def string) string {
	val, ok := os.LookupEnv(env)
	if ok {
		return val
	}
	return def
}

func init() {
	fback := path.Join(getEnv("HOME", "~"), ".config")
	file := path.Join(getEnv("XDG_CONFIG_HOME", fback), "i3", "defaults.json")
	cfile, err := filepath.Abs(file)
	if err != nil {
		log.Fatalf("Config file %s doesn't exist", file)
	}
	conf, err = ioutil.ReadFile(cfile)
	if err != nil {
		log.Fatalf("Failed to read config file %s!", cfile)
	}
}

func getFocusedWs(wss []workspace) (workspace string, err error) {
	for _, ws := range wss {
		if ws.Focused {
			workspace = ws.Name
		}
	}
	if len(workspace) < 1 {
		err = errors.New("no active workspace found")
	}
	return
}

func main() {
	var config map[string]([]string)
	err := json.Unmarshal(conf, &config)
	if err != nil {
		log.Fatal(err)
	}
	out, err := exec.Command("i3-msg", "-t", "get_workspaces").Output()
	if err != nil {
		log.Fatal(err)
	}
	var output []workspace
	json.Unmarshal(out, &output)
	ws, err := getFocusedWs(output)
	if err != nil {
		log.Fatalln(err)
	}

	args := config[ws]
	cmd := exec.Command(args[0], args[1:]...)
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
