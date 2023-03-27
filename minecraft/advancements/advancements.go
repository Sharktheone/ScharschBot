package advancements

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var langPath = flag.String("minecraftLangPath", "config/lang.json", "Path to lang.json")

func init() {
	flag.Parse()
}

func GetLang() (lang map[string]interface{}) {
	var langJson map[string]interface{}
	jsonLang, err := os.ReadFile(*langPath)
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
