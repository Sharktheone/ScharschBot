package advancements

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

const langPath = "config/lang.json"

func GetLang() (lang map[string]interface{}) {
	var langJson map[string]interface{}
	jsonLang, err := os.ReadFile(langPath)
	if err != nil {
		log.Fatalf("Failed to get lang: %v", err)
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
