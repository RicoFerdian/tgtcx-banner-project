package dictionary

type Banner struct {
	ID       int64          `json:"id"`
	Name     string         `json:"name"`
	Category BannerCategory `json:"category"`         // One banner have one category
	Location []Location     `json:"banner_locations"` // One banner have multiple target location
	Tier     []Tier         `json:"banner_tiers"`     // One banner can have multiple tiers
	ImageUrl string         `json:"image_url"`
	Start    string         `json:"start"`
	Expired  string         `json:"expired"`
}

type User struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"` // One user have one location
	Tier     Tier     `json:"tier"`
	Banner   []Banner `json:"personalized_banners"` // One user can have multiple personalized banner
}

type Location struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Tier struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type BannerCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
