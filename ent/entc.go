//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/contrib/entoas/spec"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entoas.NewExtension(
		entoas.SpecTitle("API Tutor 42"),
		entoas.SpecDescription("API to manage tutor events."),
		entoas.SpecVersion("0.0.1"),
		entoas.SimpleModels(),
		entoas.Mutations(
			alterRoutes,
			createRoutes,
		),
	)
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ex))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

func alterRoutes(g *gen.Graph, s *spec.Spec) error {
	s.Paths["/events"].Get.Parameters = nil
	s.Paths["/users"].Get.Parameters = nil
	return nil
}

func createRoutes(g *gen.Graph, s *spec.Spec) error {
	createAuthRoutes()
	createEventRoutes()
	return nil
}

func createAuthRoutes(s *spec.Spec) {
	s.Paths["/auth/login"] = &spec.Path{}
	s.Paths["/auth/token"] = &spec.Path{}
	s.Paths["/auth/token/info"] = &spec.Path{}
}

func createEventRoutes(s *spec.Spec) {
	s.Paths["/events/{id}/users/{user_id}"] = &spec.Path{
		Post: &spec.Operation{
			Summary:     "Add user to an event",
			Description: "Subscribe a user to an incoming event.",
			Tags:        []string{"Event"},
			OperationID: "subscribeEvent",
			Parameters: []*spec.Parameter{{
				Name:        "id",
				In:          spec.InPath,
				Description: "ID of the event",
				Required:    true,
				Schema: &spec.Type{
					Type: "string",
				},
			}, {
				Name:        "user_id",
				In:          spec.InPath,
				Description: "ID of the user",
				Required:    true,
				Schema: &spec.Type{
					Type: "string",
				},
			}},
			Responses: map[string]*spec.OperationResponse{
				"200": {
					Response: &spec.Response{
						Description: "User subscribed",
						Content: spec.Content{
							spec.JSON: &spec.MediaTypeObject{
								Unique: true,
								Ref: &spec.Schema{
									Name: "Event",
								},
							},
						},
					},
				},
				"400": {Ref: &spec.Response{Name: "400"}},
				"404": {Ref: &spec.Response{Name: "404"}},
				"500": {Ref: &spec.Response{Name: "500"}},
			},
		},
		Delete: &spec.Operation{
			Summary:     "Remove user from an event",
			Description: "Unsubscribe a user to an incoming event.",
			Tags:        []string{"Event"},
			OperationID: "unsubscribeEvent",
			Parameters: []*spec.Parameter{{
				Name:        "id",
				In:          spec.InPath,
				Description: "ID of the event",
				Required:    true,
				Schema: &spec.Type{
					Type: "string",
				},
			}, {
				Name:        "user_id",
				In:          spec.InPath,
				Description: "ID of the user",
				Required:    true,
				Schema: &spec.Type{
					Type: "string",
				},
			}},
			Responses: map[string]*spec.OperationResponse{
				"200": {
					Response: &spec.Response{
						Description: "User unsubscribed",
						Content: spec.Content{
							spec.JSON: &spec.MediaTypeObject{
								Unique: true,
								Ref: &spec.Schema{
									Name: "Event",
								},
							},
						},
					},
				},
				"400": {Ref: &spec.Response{Name: "400"}},
				"404": {Ref: &spec.Response{Name: "404"}},
				"500": {Ref: &spec.Response{Name: "500"}},
			},
		},
	}
}
