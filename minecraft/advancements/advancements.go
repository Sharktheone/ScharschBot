package advancements

import (
	"Scharsch-Bot/config"
	"Scharsch-Bot/flags"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var langPath = flags.String("minecraftLangPath")

func GetLang() (lang map[string]interface{}) {
	var (
		langJson map[string]interface{}
		jsonLang []byte
		err      error
	)
	if *langPath == "internal" {
		jsonLang, err = config.GetLang()
		if err != nil {
			log.Fatalf("Failed to get lang: %v", err)
		}
	} else {
		jsonLang, err = os.ReadFile(*langPath)
		if err != nil {
			log.Fatalf("Failed to get lang: %v", err)
		}
	}

	decoder := json.NewDecoder(strings.NewReader(string(jsonLang)))
	err = decoder.Decode(&langJson)
	if err != nil {
		log.Fatalf("Failed to decode Lang json: %v", err)
	}

	return langJson
}

func Decode(path string) (name string) {
	lang := GetLang()
	path = strings.ReplaceAll(path, "/", ".")
	fullPath := fmt.Sprintf("advancements.%s.title", path)
	advancementName := lang[fullPath].(string)
	return advancementName
}
