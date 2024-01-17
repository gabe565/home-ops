package main

import (
	"bytes"
	_ "embed"
	"errors"
	"flag"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed apps.html.tmpl
var appsTemplate string

type Namespace struct {
	Name       string
	Services   []Match
	Supporting []Match
}

type Match struct {
	Path      string
	Name      string
	Namespace string
}

var (
	file               string
	root               string
	startTag           string
	endTag             string
	supportingServices []string
)

func init() {
	flag.StringVar(&file, "file", "docs/kubernetes/apps.md", "Destination markdown file")
	flag.StringVar(&root, "root", "kubernetes/apps", "Root directory to search")
	flag.StringVar(&startTag, "start-tag", "<!-- Begin apps table -->", "Tag to start replacement")
	flag.StringVar(&endTag, "end-tag", "<!-- End apps table -->", "Tag to end replacment")
	supportingServicesStr := flag.String("exclude-names", "borgmatic,postgresql,redis", "Comma-separated list of manifest names to exclude")
	flag.Parse()
	supportingServices = strings.Split(*supportingServicesStr, ",")
}

func main() {
	yamlRe := regexp.MustCompile(`\.ya?ml$`)

	matches := make(map[string][]Match)
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !yamlRe.MatchString(path) {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		decoder := yaml.NewDecoder(f)
		for {
			data := make(map[string]any)
			if err := decoder.Decode(&data); err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return err
			}

			if match, err := extract(data); err != nil {
				return err
			} else if match != nil {
				match.Path = path
				matches[match.Namespace] = append(matches[match.Namespace], *match)
			}
		}
	})
	if err != nil {
		log.Fatalf("Failed to walk root: %s", err)
	}

	namespaces := prepareMatches(matches)
	if err := templateOutput(namespaces); err != nil {
		log.Fatal(err)
	}
}

func shieldsParam(s string) string {
	return strings.ReplaceAll(s, "-", "--")
}

func extract(data map[string]any) (*Match, error) {
	apiVersion, _ := data["apiVersion"].(string)
	kind, _ := data["kind"].(string)
	metadata, _ := data["metadata"].(map[string]any)
	namespace, _ := metadata["namespace"].(string)
	name, _ := metadata["name"].(string)

	match := Match{
		Name:      name,
		Namespace: namespace,
	}

	switch {
	case strings.HasPrefix(apiVersion, "helm.toolkit.fluxcd.io") && kind == "HelmRelease":
	case apiVersion == "postgresql.cnpg.io/v1" && kind == "Cluster":
	default:
		return nil, nil
	}
	return &match, nil
}

func prepareMatches(matches map[string][]Match) []Namespace {
	namespaces := make([]Namespace, 0, len(matches))

	for namespace, match := range matches {
		slices.SortStableFunc(match, func(a, b Match) int {
			return strings.Compare(a.Name, b.Name)
		})
		services := make([]Match, 0, len(match))
		supporting := make([]Match, 0, len(match))
		for _, service := range match {
			var supportingService bool
			for _, supportingName := range supportingServices {
				if service.Name == supportingName {
					supportingService = true
					break
				}
			}
			if supportingService {
				supporting = append(supporting, service)
			} else {
				services = append(services, service)
			}
		}
		namespaces = append(namespaces, Namespace{
			Name:       namespace,
			Services:   services,
			Supporting: supporting,
		})
	}
	slices.SortStableFunc(namespaces, func(a, b Namespace) int {
		return strings.Compare(a.Name, b.Name)
	})
	return namespaces
}

func templateOutput(namespaces []Namespace) error {
	tmpl, err := template.New("").
		Funcs(template.FuncMap{"shieldsParam": shieldsParam}).
		Parse(appsTemplate)
	if err != nil {
		log.Fatalf("Failed to load template: %s", err)
	}

	src, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to load %q: %s", file, err)
	}

	startIdx := bytes.Index(src, []byte(startTag))
	if startIdx == -1 {
		log.Fatalf("Could not find start tag: %q", startTag)
	}

	endIdx := bytes.Index(src, []byte(endTag))
	if endIdx == -1 {
		log.Fatalf("Could not find end tag: %q", endTag)
	}

	buf := bytes.NewBuffer(make([]byte, 0, endIdx-startIdx))
	buf.Write(src[:startIdx+len(startTag)+1])
	if err := tmpl.Execute(buf, namespaces); err != nil {
		log.Fatalf("Failed to execute template: %s", err)
	}
	buf.Write(src[endIdx:])

	if err := os.WriteFile(file, buf.Bytes(), 0644); err != nil {
		log.Fatalf("Failed to write %q: %s", file, err)
	}

	return nil
}
