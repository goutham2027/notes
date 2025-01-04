package urlshort

import (
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		UrlPath := r.URL.Path
		if v, ok := pathsToUrls[UrlPath]; ok {
			http.Redirect(rw, r, v, http.StatusPermanentRedirect)
			return
		}
		fallback.ServeHTTP(rw, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	yamlPaths, err := getYamlPaths(yml)
	if err != nil {
		return nil, err
	}

	return MapHandler(yamlPaths, fallback), nil
}

func getYamlPaths(yml []byte) (map[string]string, error) {
	type yamlPath struct {
		Path string `yaml:"path,omitempty"`
		URL  string `yaml:"url,omitempty"`
	}
	var yamlPaths []yamlPath

	err := yaml.Unmarshal(yml, &yamlPaths)
	if err != nil {
		log.Printf("cannot unmarshal data: %v", err)
		return nil, err
	}

	yamlPathsMap := make(map[string]string, len(yamlPaths))
	for _, y := range yamlPaths {
		yamlPathsMap[y.Path] = y.URL
	}
	return yamlPathsMap, err
}
