package gqlserver

import "github.com/graphql-go/graphql"

var BannerCategoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "BannerCategory",
		Description: "Detail of the banner category",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var BannerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Banner",
		Description: "Detail of the banner",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"category": &graphql.Field{
				Type: BannerCategoryType,
			},
			"location": &graphql.Field{
				Type: graphql.NewList(LocationType),
			},
			"tier": &graphql.Field{
				Type: graphql.NewList(TierType),
			},
			"expired": &graphql.Field{
				Type: graphql.String,
			},
			"image_url": &graphql.Field{
				Type: graphql.String,
			},
			"start": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "User",
		Description: "Detail of the user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"location": &graphql.Field{
				Type: LocationType,
			},
			"tier": &graphql.Field{
				Type: TierType,
			},
			"personalized_banners": &graphql.Field{
				Type: graphql.NewList(BannerType),
			},
		},
	},
)

var TierType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Tier",
		Description: "Detail of the tier",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var LocationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Location",
		Description: "Detail of the location",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
