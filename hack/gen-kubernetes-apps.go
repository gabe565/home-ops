package main

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v3"
)

var (
	file               = "README.md"
	root               = "kubernetes/apps"
	startTag           = "<!-- Begin apps section -->"
	endTag             = "<!-- End apps section -->"
	supportingServices = []string{
		"borgmatic",
		"postgresql",
		"redis",
		"mariadb",
	}
	yamlRe = regexp.MustCompile(`\.ya?ml$`)
)

//go:embed apps.html.tmpl
var appsTemplate string

type Namespace struct {
	Name       string
	Services   map[string]Match
	Supporting map[string]Match
}

type Match struct {
	Path      string
	Name      string
	Namespace string
}

func main() {
	var namespaces map[string]Namespace

	var group errgroup.Group
	matchCh := make(chan Match)

	group.Go(func() error {
		defer close(matchCh)
		return filepath.Walk(root, walkFunc(matchCh))
	})

	group.Go(func() error {
		namespaces = prepareMatches(matchCh)
		return nil
	})

	if err := group.Wait(); err != nil {
		log.Fatal(err)
	}

	if err := templateOutput(namespaces); err != nil {
		log.Fatal(err)
	}
}

func walkFunc(matchCh chan Match) filepath.WalkFunc {
	outputSubdirCount := strings.Count(file, string(os.PathSeparator))
	outputPathPrefix := strings.Repeat(".."+string(os.PathSeparator), outputSubdirCount)

	return func(path string, info fs.FileInfo, err error) error {
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
			var data any
			if err := decoder.Decode(&data); err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return fmt.Errorf("unmarshal failed for %q: %w", path, err)
			}

			if data, ok := data.(map[string]any); ok {
				apiVersion, _ := data["apiVersion"].(string)
				kind, _ := data["kind"].(string)
				metadata, _ := data["metadata"].(map[string]any)
				name, _ := metadata["name"].(string)

				switch {
				case strings.HasPrefix(apiVersion, "helm.toolkit.fluxcd.io") && kind == "HelmRelease":
				case apiVersion == "postgresql.cnpg.io/v1" && kind == "Cluster":
				default:
					continue
				}

				namespace, _ := metadata["namespace"].(string)
				path = filepath.Join(outputPathPrefix, path)

				matchCh <- Match{
					Path:      path,
					Name:      name,
					Namespace: namespace,
				}
			}
		}
	}
}

func prepareMatches(matches chan Match) map[string]Namespace {
	namespaces := make(map[string]Namespace)

	for service := range matches {
		namespace, ok := namespaces[service.Namespace]
		if !ok {
			namespace = Namespace{
				Name:       service.Namespace,
				Services:   make(map[string]Match),
				Supporting: make(map[string]Match),
			}
			namespaces[namespace.Name] = namespace
		}

		var supportingService bool
		for _, supportingName := range supportingServices {
			if service.Name == supportingName {
				supportingService = true
				break
			}
		}
		if supportingService {
			namespace.Supporting[service.Name] = service
		} else {
			namespace.Services[service.Name] = service
		}
	}

	return namespaces
}

func templateOutput(clusters map[string]Namespace) error {
	tmpl, err := template.New("").Parse(appsTemplate)
	if err != nil {
		return err
	}

	src, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	startIdx := bytes.Index(src, []byte(startTag))
	if startIdx == -1 {
		return fmt.Errorf("no start tag %q in %q", startTag, file)
	}

	endIdx := bytes.Index(src, []byte(endTag))
	if endIdx == -1 {
		return fmt.Errorf("no end tag %q in %q", endTag, file)
	}

	buf := bytes.NewBuffer(make([]byte, 0, endIdx-startIdx))
	buf.Write(src[:startIdx+len(startTag)+1])
	if err := tmpl.Execute(buf, clusters); err != nil {
		return err
	}
	buf.Write(src[endIdx:])

	if err := os.WriteFile(file, buf.Bytes(), 0o644); err != nil {
		return err
	}

	return nil
}
