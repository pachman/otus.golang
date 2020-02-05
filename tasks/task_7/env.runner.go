package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ReadDir(dir string) (map[string]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := map[string]string{}

	for _, f := range files {
		file, err := os.Open(dir + f.Name())
		if err != nil {
			return nil, err
		}

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		name := strings.TrimSpace(f.Name())
		env[name] = strings.TrimSpace(string(bytes))
		file.Close()
	}

	return env, nil
}

func RunCmd(commands []string, env map[string]string) int {
	envs := make([]string, 0)

	for key, value := range env {
		if key != "" && value != "" {
			envs = append(envs, key+"="+value)
		}
	}

	cmdName := commands[0]
	cmd := exec.Command(cmdName)
	cmd.Args = commands
	cmd.Env = envs
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	return 0
}
