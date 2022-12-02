package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetChatterinoConfigURL() string {
	var configMode int

	var os string = runtime.GOOS

	switch os {
	case "windows":
		configMode = 0
	case "darwin":
		configMode = 1
	case "linux":
		configMode = 2
	default:
		fmt.Printf("%s.\n", os)
	}

	var configUrls [3]string = [3]string{"/AppData/Roaming/Chatterino2", "/Library/Application", "/.local/share/chatterino"}
	return configUrls[configMode]
}

func GetChatterinoConfigJSON() string {
	fmt.Print("Opening client dir: ")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	path := homeDir + GetChatterinoConfigURL()
	fmt.Println(path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal("Directory does not exist")
	}

	fmt.Println("Retreiving Config: ")

	files, fpErr := filepath.Glob(path + "/*")
	if fpErr != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		fmt.Println("* file - " + file)
	}

	// parse config
	settingsPath := path + "/Settings/settings.json"

	fileContent, ioErr := ioutil.ReadFile(settingsPath)
	if ioErr != nil {
		log.Fatal(ioErr)
	}

	jsonString := string(fileContent)
	return jsonString
}

func ParseChatterinoConfigJSON(jsonString string) ChatterinoConfig {

	jsonDataReader := strings.NewReader(jsonString)
	decoder := json.NewDecoder(jsonDataReader)
	var chatterinoConfig ChatterinoConfig
	decoderErr := decoder.Decode(&chatterinoConfig)
	if decoderErr != nil {
		panic(decoderErr)
	}
	/*for {
		err := decoder.Decode(&chatterinoConfig)
		if err != nil {
			panic(err)
		}
		if err == io.EOF {
			break
		}
	}*/
	return chatterinoConfig
}

func PushChatterinoConfigChanges(configUpdates ChatterinoConfig) {

}
