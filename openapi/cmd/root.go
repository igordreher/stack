package cmd

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "openapi",
	RunE: func(cmd *cobra.Command, args []string) error {
		components, err := os.ReadDir("../components")
		if err != nil {
			return err
		}

		loader := &openapi3.Loader{
			Context:               cmd.Context(),
			IsExternalRefsAllowed: true,
		}

		finalSpec, err := loader.LoadFromFile("./base.yaml")
		if err != nil {
			return err
		}
		finalSpec.Components.Responses = openapi3.Responses{}
		finalSpec.Components.Parameters = openapi3.ParametersMap{}
		finalSpec.Components.RequestBodies = openapi3.RequestBodies{}

		for _, component := range components {
			files, err := os.ReadDir("../components/" + component.Name())
			if err != nil {
				return err
			}
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				var r = regexp.MustCompile("openapi(-(v[0-9]+))?.yaml")
				if r.MatchString(file.Name()) {
					matches := r.FindStringSubmatch(file.Name())
					version := matches[2]
					if version == "" {
						version = "v1"
					}

					doc, err := loader.LoadFromFile("../components/" + component.Name() + "/" + file.Name())
					if err != nil {
						return err
					}

					err = doc.Validate(cmd.Context())
					if err != nil {
						return err
					}

					for endpoint, path := range doc.Paths {
						for _, e := range []*openapi3.Operation{path.Put, path.Get, path.Post, path.Delete, path.Head, path.Patch, path.Options} {
							if e == nil {
								continue
							}
							e.Tags = append(e.Tags, fmt.Sprintf("%s.%s", component.Name(), version))
							if e.RequestBody != nil {
								e.RequestBody.Ref = rewriteRef(e.RequestBody.Ref, version)
								if value := e.RequestBody.Value; value != nil {
									for _, mediaType := range value.Content {
										if schema := mediaType.Schema; schema != nil {
											schema.Ref = rewriteRef(schema.Ref, version)
										}
									}
								}
							}
							for _, response := range e.Responses {
								response.Ref = rewriteRef(response.Ref, version)
								if value := response.Value; value != nil {
									for _, mediaType := range value.Content {
										if schema := mediaType.Schema; schema != nil {
											schema.Ref = rewriteRef(schema.Ref, version)
										}
									}
								}
							}
						}

						finalEndpoint := fmt.Sprintf("/api/%s", component.Name())
						if version != "" && version != "v1" && version != "v0" {
							finalEndpoint += "/" + version
						}
						finalEndpoint += endpoint

						finalSpec.Paths[finalEndpoint] = path
					}

					processSchemas(doc.Components.Schemas, version)
					for k, v := range doc.Components.Schemas {
						finalSpec.Components.Schemas[rewriteRef(k, version)] = v
					}
					processRequestBodies(doc.Components.RequestBodies, version)
					for k, v := range doc.Components.RequestBodies {
						finalSpec.Components.RequestBodies[rewriteRef(k, version)] = v
					}

					for k, v := range doc.Components.Responses {
						finalSpec.Components.Responses[rewriteRef(k, version)] = v
					}
					for k, v := range doc.Components.Parameters {
						finalSpec.Components.Parameters[rewriteRef(k, version)] = v
					}
				}
			}
		}

		outputFile, err := os.Create("./build/generate.yaml")
		if err != nil {
			return err
		}

		asJSON, err := finalSpec.MarshalJSON()
		if err != nil {
			return err
		}

		asMap := make(map[string]any)
		if err := json.Unmarshal(asJSON, &asMap); err != nil {
			panic(err)
		}

		return yaml.NewEncoder(outputFile).Encode(asMap)
	},
}

func rewriteRef(reference, version string) string {
	if reference == "" {
		return reference
	}
	if version == "v1" || version == "v0" {
		return reference
	}
	return reference + strings.ToUpper(version)
}

func processSchemaRefs(refs openapi3.SchemaRefs, version string) {
	for _, ref := range refs {
		ref.Ref = rewriteRef(ref.Ref, version)
		processSchema(ref.Value, version)
	}
}

var visited = map[*openapi3.Schema]any{}

func processSchema(schema *openapi3.Schema, version string) {
	if schema == nil {
		return
	}
	_, ok := visited[schema]
	if ok {
		return
	}
	visited[schema] = struct{}{}
	processSchemaRefs(schema.AllOf, version)
	processSchemaRefs(schema.OneOf, version)
	processSchemaRefs(schema.AnyOf, version)
	processSchemaRef(schema.Not, version)
	processSchemaRef(schema.Items, version)
	processSchemas(schema.Properties, version)
	processSchemaRef(schema.AdditionalProperties.Schema, version)
}

func processSchemaRef(schema *openapi3.SchemaRef, version string) {
	if schema == nil {
		return
	}
	schema.Ref = rewriteRef(schema.Ref, version)
	processSchema(schema.Value, version)
}

func processSchemas(schemas openapi3.Schemas, version string) {
	for _, schema := range schemas {
		processSchemaRef(schema, version)
	}
}

func processRequestBodies(schemas openapi3.RequestBodies, version string) {
	for _, schema := range schemas {
		processRequestBodyRef(schema, version)
	}
}

func processRequestBodyRef(schema *openapi3.RequestBodyRef, version string) {
	schema.Ref = rewriteRef(schema.Ref, version)
	if value := schema.Value; value != nil {
		processRequestBody(value, version)
	}
}

func processRequestBody(value *openapi3.RequestBody, version string) {
	for _, mediaType := range value.Content {
		processSchemaRef(mediaType.Schema, version)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
