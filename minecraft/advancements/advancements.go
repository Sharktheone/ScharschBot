package advancements

import (
	"Scharsch-Bot/config"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var langPath = flag.String("minecraftLangPath", "internal", "Path to lang.json")

func init() {
	flag.Parse()
}

func GetLang() (lang map[string]interface{}) {
	var (
		langJson map[string]interface{}
		jsonLang []byte
		err      error
	)
	if *langPath == "internal" {
		jsonLang, err = config.MCLangJson.ReadFile("lang.json")
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
