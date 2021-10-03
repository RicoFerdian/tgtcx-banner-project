package gqlserver

import (
	"github.com/graphql-go/graphql"
)

type SchemaWrapper struct {
	dataResolver *Resolver
	Schema       graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) WithProductResolver(pr *Resolver) *SchemaWrapper {
	s.dataResolver = pr
	return s
}

func (s *SchemaWrapper) Init() error {
	// Banner schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "BannerGetter",
			Description: "All query related to getting product data",
			Fields: graphql.Fields{
				// Banner related
				"BannerDetail": &graphql.Field{
					Type:        BannerType,
					Description: "Get single banner detail by banner id",
					Args: graphql.FieldConfigArgument{
						"banner_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.dataResolver.GetBanner(),
				},
				"Banners": &graphql.Field{
					Type:        graphql.NewList(BannerType),
					Description: "Get all banners",
					Resolve:     s.dataResolver.GetBanners(),
				},
				// User related
				"User": &graphql.Field{
					Type:        UserType,
					Description: "Get single user detail based on user id",
					Args: graphql.FieldConfigArgument{
						"user_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.dataResolver.GetUser(),
				},
				"Users": &graphql.Field{
					Type:        graphql.NewList(UserType),
					Description: "Get all users",
					Resolve:     s.dataResolver.GetUsers(),
				},
				"UserPersonalizedBanners": &graphql.Field{
					Type:        UserType,
					Description: "Get user with personalized banner by user id",
					Args: graphql.FieldConfigArgument{
						"user_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.dataResolver.GetUserBanners(),
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:        "BannerSetter",
			Description: "All query related to modify banner data",
			Fields: graphql.Fields{
				// Create new banner
				"CreateBanner": &graphql.Field{
					Type:        BannerType,
					Description: "Create banner",
					Args: graphql.FieldConfigArgument{
						"banner_name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"banner_category": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"banner_expired": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"banner_start": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"banner_image_url": &graphql.ArgumentConfig{
							Type: graphql.String,
						},

						// Related table parameters
						"banner_tiers": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.Int),
						},
						"banner_locations": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.Int),
						},
					},
					Resolve: s.dataResolver.CreateBanner(),
				},
				// Update a banner by its id
				"UpdateBanner": &graphql.Field{
					Type:        BannerType,
					Description: "Update banner by ID",
					Args: graphql.FieldConfigArgument{
						"banner_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"banner_name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"banner_category": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"banner_expired": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"banner_start": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"banner_image_url": &graphql.ArgumentConfig{
							Type: graphql.String,
						},

						// Related table parameters
						"banner_tiers": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.Int),
						},
						"banner_locations": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.Int),
						},
					},
					Resolve: s.dataResolver.UpdateBanner(),
				},
				// TODO : Update banner tier
				// TODO : Update banner location
			},
		}),
	})

	if err != nil {
		return err
	}

	s.Schema = schema

	return nil
}
