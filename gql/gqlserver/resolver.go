package gqlserver

import (
	"github.com/graphql-go/graphql"
	"github.com/radityaqb/tgtc/backend/dictionary"
	"github.com/radityaqb/tgtc/backend/service"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

///////////////////////   Banner Resolver   /////////////////////////
func (r *Resolver) GetBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["banner_id"].(int)
		return service.GetBanner(id)
	}
}

func (r *Resolver) GetBanners() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return service.GetBanners()
	}
}
func (r *Resolver) GetTierBanners() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["user_id"].(int)
		return service.GetTierBanners(id)
	}
}
func (r *Resolver) GetUserBanners() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["user_id"].(int)
		return service.GetUserBanners(id)
	}
}

func (r *Resolver) CreateBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["banner_name"].(string)
		category_id, _ := p.Args["banner_category"].(int64)
		expired, _ := p.Args["banner_expired"].(string)
		start, _ := p.Args["banner_start"].(string)
		image_url, _ := p.Args["banner_image_url"].(string)

		// Array input for updating related table
		tiers, _ := p.Args["banner_tiers"].([]int64)
		locations, _ := p.Args["banner_locations"].([]int64)

		req := dictionary.Banner{
			Name:     name,
			Category: dictionary.BannerCategory{ID: category_id},
			Expired:  expired,
			Start:    start,
			ImageUrl: image_url,
		}
		return service.CreateBanner(req, tiers, locations)
	}
}

func (r *Resolver) UpdateBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["banner_id"].(int)
		name, _ := p.Args["banner_name"].(string)
		category_id, _ := p.Args["banner_category"].(int64)
		expired, _ := p.Args["banner_expired"].(string)
		start, _ := p.Args["banner_start"].(string)
		image_url, _ := p.Args["banner_image_url"].(string)

		// Array input for updating related table
		tiers, _ := p.Args["banner_tiers"].([]int64)
		locations, _ := p.Args["banner_locations"].([]int64)

		req := dictionary.Banner{
			ID:       int64(id),
			Name:     name,
			Category: dictionary.BannerCategory{ID: category_id},
			Expired:  expired,
			Start:    start,
			ImageUrl: image_url,
		}

		return service.UpdateBanner(req, tiers, locations)
	}
}

///////////////////////   User Resolver   /////////////////////////
func (r *Resolver) GetUser() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["user_id"].(int)
		return service.GetUser(id)
	}
}

func (r *Resolver) GetUsers() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return service.GetUsers()
	}
}

func (r *Resolver) CreateUser() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["user_name"].(string)
		location, _ := p.Args["user_location"].(int64)
		tier, _ := p.Args["user_tier"].(int64)

		req := dictionary.User{
			Name:     name,
			Location: dictionary.Location{ID: location},
			Tier:     dictionary.Tier{ID: tier},
		}
		_, err := service.CreateUser(req)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func (r *Resolver) UpdateUser() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["user_id"].(int)
		name, _ := p.Args["user_name"].(string)
		location, _ := p.Args["user_location"].(int64)
		tier, _ := p.Args["user_tier"].(int64)

		req := dictionary.User{
			ID:       int64(id),
			Name:     name,
			Location: dictionary.Location{ID: location},
			Tier:     dictionary.Tier{ID: tier},
		}
		_, err := service.UpdateUser(req)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

///////////////////////   Tier Resolver   /////////////////////////
func (r *Resolver) GetTier() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["tier_id"].(int)
		return service.GetTier(id)
	}
}

func (r *Resolver) GetTiers() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return service.GetTiers()
	}
}

///////////////////////   Locations Resolver   /////////////////////////
func (r *Resolver) GetLocation() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["location_id"].(int)
		return service.GetLocation(id)
	}
}

func (r *Resolver) GetLocations() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return service.GetLocations()
	}
}

///////////////////////   Banner Category Resolver   /////////////////////////
func (r *Resolver) GetBannerCategory() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["banner_category_id"].(int)
		return service.GetBannerCategory(id)
	}
}

func (r *Resolver) GetBannerCategories() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return service.GetBannerCategories()
	}
}
