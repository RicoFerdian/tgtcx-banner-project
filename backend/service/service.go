package service

import (
	"errors"
	"time"

	"github.com/radityaqb/tgtc/backend/database"
	"github.com/radityaqb/tgtc/backend/dictionary"
)

var currentTime = time.Now()
var currentTimeFormat = currentTime.Format("YYYY-MM-DD")

func getBannerLocations(bannerID int64) ([]dictionary.Location, error) {
	db := database.GetDB()
	query := `
		SELECT l.location_id, l.location_name
		FROM locations l
		INNER JOIN bannerlocationtable bl ON bl.location_id=l.location_id
		WHERE bl.banner_id = $1
	`
	// Query execution, works just fine
	rows, err := db.Query(query, bannerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []dictionary.Location
	for rows.Next() {
		var data dictionary.Location
		rows.Scan(
			&data.ID, &data.Name,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}

func getBannerTiers(bannerID int64) ([]dictionary.Tier, error) {
	db := database.GetDB()
	query := `
		SELECT t.tier_id, t.tier_name
		FROM tiers t
		INNER JOIN bannertiertable bt ON bt.tier_id=t.tier_id
		WHERE bt.banner_id = $1
	`
	// Query execution, works just fine
	rows, err := db.Query(query, bannerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []dictionary.Tier
	for rows.Next() {
		var data dictionary.Tier
		rows.Scan(
			&data.ID, &data.Name,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}

func createBannerTier(tierID int64, bannerID int64) (bool, error) {
	db := database.GetDB()

	// construct query
	query := `
		INSERT INTO bannertiertable (banner_id, tier_id) VALUES
		($1, $2)
	`
	// actual query process
	result, err := db.Exec(query,
		bannerID,
		tierID,
	)
	if err != nil {
		return false, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if affected == 0 {
		return false, errors.New("no row created")
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func createBannerTiers(tiers []int64, bannerID int64) ([]dictionary.Tier, error) {
	var bulkresults []dictionary.Tier

	for _, tierID := range tiers {
		var data dictionary.Tier
		success, err := createBannerTier(tierID, bannerID)
		if err != nil || !success {
			return nil, err
		}
		bulkresults = append(bulkresults, data)
	}

	return bulkresults, nil
}

func deleteBannerTier(bannerID int64) (bool, error) {
	db := database.GetDB()
	query := `
		DELETE FROM bannertiertable
		WHERE banner_id = $1
	`
	// actual query process
	result, err := db.Exec(query,
		bannerID,
	)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if affected == 0 {
		return false, errors.New("no row updated")
	}
	return true, nil
}

func updateBannerTiers(tiers []int64, bannerID int64) ([]dictionary.Tier, error) {
	success, err := deleteBannerTier(bannerID)
	if err != nil || !success {
		return nil, err
	}
	return createBannerTiers(tiers, bannerID)
}

func createBannerLocation(locationID int64, bannerID int64) (bool, error) {
	db := database.GetDB()

	// construct query
	query := `
		INSERT INTO bannerlocationtable (banner_id, location_id) VALUES
		($1, $2)
	`
	// actual query process
	result, err := db.Exec(query,
		bannerID,
		locationID,
	)
	if err != nil {
		return false, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if affected == 0 {
		return false, errors.New("no row created")
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func createBannerLocations(locations []int64, bannerID int64) ([]dictionary.Location, error) {
	var bulkresults []dictionary.Location

	for _, locationID := range locations {
		var data dictionary.Location
		success, err := createBannerLocation(locationID, bannerID)
		if err != nil || !success {
			return nil, err
		}
		bulkresults = append(bulkresults, data)
	}

	return bulkresults, nil
}

func deleteBannerLocation(bannerID int64) (bool, error) {
	db := database.GetDB()
	query := `
		DELETE FROM bannerlocationtable
		WHERE banner_id = $1
	`
	// actual query process
	result, err := db.Exec(query,
		bannerID,
	)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if affected == 0 {
		return false, errors.New("no row updated")
	}
	return true, nil
}

func updateBannerLocations(locations []int64, bannerID int64) ([]dictionary.Location, error) {
	success, err := deleteBannerLocation(bannerID)
	if err != nil || !success {
		return nil, err
	}
	return createBannerLocations(locations, bannerID)
}

///////////////////////   Banners   /////////////////////////
// Get a single banner detail
// when user clicked the banner
// We need to return all the detailed value of a single banner
func GetBanner(paramID int) (*dictionary.Banner, error) {
	db := database.GetDB()
	query := `
		SELECT b.banner_id, b.banner_name, b.banner_expired, b.banner_start, b.banner_image_url,
		bc.banner_category_id, bc.banner_category_name
		FROM banners b 
		INNER JOIN bannerCategories bc ON bc.banner_category_id=b.banner_category
		WHERE banner_id = $1
	`
	row := db.QueryRow(query, paramID)
	// read query result, and assign to variable(s)
	var data dictionary.Banner
	err := row.Scan(
		&data.ID, &data.Name, &data.Expired, &data.Start, &data.ImageUrl, &data.Category.ID, &data.Category.Name,
	)
	if err != nil {
		return nil, err
	}
	data.Location, err = getBannerLocations(data.ID)
	if err != nil {
		return nil, err
	}
	data.Tier, err = getBannerTiers(data.ID)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func GetBanners() ([]dictionary.Banner, error) {
	db := database.GetDB()
	query := `
		SELECT b.banner_id, b.banner_name, b.banner_expired, b.banner_start, b.banner_image_url,
		bc.banner_category_id, bc.banner_category_name
		FROM banners b 
		INNER JOIN bannerCategories bc ON bc.banner_category_id=b.banner_category
		LIMIT 20
	`
	// Query execution, works just fine
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []dictionary.Banner
	for rows.Next() {
		var data dictionary.Banner
		rows.Scan(
			&data.ID, &data.Name, &data.Expired, &data.Start, &data.ImageUrl, &data.Category.ID, &data.Category.Name,
		)
		if err != nil {
			return nil, err
		}
		data.Location, err = getBannerLocations(data.ID)
		if err != nil {
			return nil, err
		}
		data.Tier, err = getBannerTiers(data.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}

func GetTierBanners(userID int) ([]dictionary.Banner, error) {
	db := database.GetDB()

	user, err := GetUser(userID)
	if err != nil {
		return nil, err
	}

	// Query for banners by user's tier and location, limited to 20 for now
	query := `
		SELECT b.banner_id, b.banner_name, b.banner_expired, b.banner_start, b.banner_image_url,
		bc.banner_category_id, bc.banner_category_name
		FROM banners b 
		INNER JOIN bannerCategories bc ON bc.banner_category_id=b.banner_category
		INNER JOIN bannertiertable bt ON bt.banner_id=b.banner_id
		INNER JOIN tiers t ON t.tier_id=bt.tier_id
		INNER JOIN bannerlocationtable bl ON bl.banner_id=b.banner_Id
		INNER JOIN locations l ON l.location_id=bl.location_id
		WHERE bt.tier_id = $1 AND bl.location_id = $2
		AND b.banner_start >= $3 AND b.banner_expired < $3
		LIMIT 20
	`
	// Query execution, works just fine
	rows, err := db.Query(query, user.Tier.ID, user.Location.ID, currentTimeFormat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []dictionary.Banner
	for rows.Next() {
		var data dictionary.Banner
		rows.Scan(
			&data.ID, &data.Name, &data.Expired, &data.Start, &data.Category.ID, &data.Category.Name,
		)
		if err != nil {
			return nil, err
		}
		data.Location, err = getBannerLocations(data.ID)
		if err != nil {
			return nil, err
		}
		data.Tier, err = getBannerTiers(data.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}
	return result, nil
}

func GetUserBanners(userID int) (*dictionary.User, error) {
	db := database.GetDB()

	user, err := GetUser(userID)
	if err != nil {
		return nil, err
	}

	// Query for banners by user's tier and location, limited to 20 for now
	query := `
		SELECT b.banner_id, b.banner_name, b.banner_expired, b.banner_start, b.banner_image_url,
		bc.banner_category_id, bc.banner_category_name
		FROM banners b 
		INNER JOIN bannerCategories bc ON bc.banner_category_id=b.banner_category
		INNER JOIN bannertiertable bt ON bt.banner_id=b.banner_id
		INNER JOIN tiers t ON t.tier_id=bt.tier_id
		INNER JOIN bannerlocationtable bl ON bl.banner_id=b.banner_Id
		INNER JOIN locations l ON l.location_id=bl.location_id
		WHERE bt.tier_id = $1 AND bl.location_id = $2
		AND b.banner_start >= $3 AND b.banner_expired < $3
		LIMIT 20
	`
	// Query execution, filter by user tier and location
	rows, err := db.Query(query, user.Tier.ID, user.Location.ID, currentTimeFormat)
	if err != nil {
		return user, err
	}
	defer rows.Close()
	var result []dictionary.Banner
	for rows.Next() {
		var data dictionary.Banner
		rows.Scan(
			&data.ID, &data.Name, &data.Expired, &data.Start, &data.ImageUrl, &data.Category.ID, &data.Category.Name,
		)
		if err != nil {
			return user, err
		}
		data.Location, err = getBannerLocations(data.ID)
		if err != nil {
			return nil, err
		}
		data.Tier, err = getBannerTiers(data.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}
	user.Banner = result
	return user, nil
}

// CreateBanner
func CreateBanner(data dictionary.Banner, tiers []int64, locations []int64) (*dictionary.Banner, error) {
	db := database.GetDB()
	// construct query
	query := `
	INSERT INTO banners (banner_name, banner_category, banner_expired, banner_start, banner_image_url) VALUES
		($1, $2, $3, $4, $5) RETURNING banner_id
	`
	// actual query process
	err := db.QueryRow(query,
		data.Name,
		data.Category.ID,
		data.Expired,
		data.Start,
		data.ImageUrl,
	).Scan(&data.ID)
	if err != nil {
		return nil, err
	}

	// Related tables
	data.Tier, err = createBannerTiers(tiers, data.ID)
	if err != nil {
		return nil, err
	}
	data.Location, err = createBannerLocations(locations, data.ID)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func UpdateBanner(data dictionary.Banner, tiers []int64, locations []int64) (*dictionary.Banner, error) {
	db := database.GetDB()
	query := `
	UPDATE 
		banners
	SET 
		banner_name = $2,
		banner_category = $3,
		banner_expired = $4,
		banner_start = $5,
		banner_image_url = $6
	WHERE
		banner_id = $1
	`
	// actual query process
	result, err := db.Exec(query,
		data.ID,
		data.Name,
		data.Category.ID,
		data.Expired,
		data.Start,
		data.ImageUrl,
	)
	if err != nil {
		return nil, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, errors.New("no row updated")
	}
	if err != nil {
		return nil, err
	}

	// Related tables
	data.Tier, err = updateBannerTiers(tiers, data.ID)
	if err != nil {
		return nil, err
	}
	data.Location, err = updateBannerLocations(locations, data.ID)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

///////////////////////   User   /////////////////////////
func GetUser(paramID int) (*dictionary.User, error) {
	db := database.GetDB()
	query := `
		SELECT u.user_id, u.user_name, l.location_id, l.location_name, t.tier_id, t.tier_name
		FROM users u
		INNER JOIN tiers t ON t.tier_id=u.user_tier
		INNER JOIN locations l ON l.location_id=u.user_location
		WHERE user_id = $1
	`
	row := db.QueryRow(query, paramID)
	// read query result, and assign to variable(s)
	var data dictionary.User
	err := row.Scan(
		&data.ID, &data.Name, &data.Location.ID, &data.Location.Name, &data.Tier.ID, &data.Tier.Name,
	)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetUsers() ([]dictionary.User, error) {
	db := database.GetDB()
	query := `
		SELECT u.user_id, u.user_name, l.location_id, l.location_name, t.tier_id, t.tier_name
		FROM users u
		INNER JOIN tiers t ON t.tier_id=u.user_tier
		INNER JOIN locations l ON l.location_id=u.user_location
		LIMIT 20
	`
	// Query execution, works just fine
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []dictionary.User
	for rows.Next() {
		var data dictionary.User
		rows.Scan(
			&data.ID, &data.Name, &data.Location.ID, &data.Location.Name, &data.Tier.ID, &data.Tier.Name,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}
	return result, nil
}

func CreateUser(data dictionary.User) (*dictionary.User, error) {
	db := database.GetDB()
	// construct query
	query := `
		INSERT INTO users (user_name, user_location, user_tier) VALUES
		($1, $2, $3)
	`
	// actual query process
	result, err := db.Exec(query,
		data.Name,
		data.Location.ID,
		data.Tier.ID,
	)
	if err != nil {
		return nil, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, errors.New("no row created")
	}
	data.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func UpdateUser(data dictionary.User) (*dictionary.User, error) {
	db := database.GetDB()
	query := `
	UPDATE 
		users
	SET 
		user_name = $2,
		user_location = $3,
		user_tier = $4,
	WHERE
		user_id = $1
	`
	// actual query process
	result, err := db.Exec(query,
		data.ID,
		data.Name,
		data.Location.ID,
		data.Tier.ID,
	)
	if err != nil {
		return nil, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, errors.New("no row updated")
	}
	data.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

///////////////////////   Location   /////////////////////////
// Buat ini Get All dan Get Details aja, simplifikasi aja
// Buat update dan create kayaknya 1 contoh saja cukup
func GetLocation(paramID int) (*dictionary.Location, error) {
	db := database.GetDB()
	query := `
		SELECT location_id, location_name
		FROM locations
		WHERE location_id = $1
	`
	row := db.QueryRow(query, paramID)
	// read query result, and assign to variable(s)
	var data dictionary.Location
	err := row.Scan(
		&data.ID, &data.Name,
	)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetLocations() ([]dictionary.Location, error) {
	db := database.GetDB()
	query := `
		SELECT location_id, location_name
		FROM locations LIMIT 20
	`
	// Query execution, works just fine
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []dictionary.Location
	for rows.Next() {
		var data dictionary.Location
		rows.Scan(
			&data.ID,
			&data.Name,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}
	return result, nil
}

///////////////////////   Tier   /////////////////////////
func GetTier(paramID int) (*dictionary.Tier, error) {
	db := database.GetDB()
	query := `
		SELECT tier_id, tier_name
		FROM tiers
		WHERE tier_id = $1
	`
	row := db.QueryRow(query, paramID)
	// read query result, and assign to variable(s)
	var data dictionary.Tier
	err := row.Scan(
		&data.ID, &data.Name,
	)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetTiers() ([]dictionary.Tier, error) {
	db := database.GetDB()
	query := `
		SELECT tier_id, tier_name
		FROM tiers LIMIT 20
	`
	// Query execution, works just fine
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []dictionary.Tier
	for rows.Next() {
		var data dictionary.Tier
		rows.Scan(
			&data.ID,
			&data.Name,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}
	return result, nil
}

///////////////////////   Banner Categories   /////////////////////////
func GetBannerCategory(paramID int) (*dictionary.BannerCategory, error) {
	db := database.GetDB()
	query := `
		SELECT banner_category_id, banner_category_name
		FROM bannercategories
		WHERE banner_category_id = $1
	`
	row := db.QueryRow(query, paramID)
	// read query result, and assign to variable(s)
	var data dictionary.BannerCategory
	err := row.Scan(
		&data.ID, &data.Name,
	)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetBannerCategories() ([]dictionary.BannerCategory, error) {
	db := database.GetDB()
	query := `
		SELECT banner_category_id, banner_category_name
		FROM bannercategories LIMIT 20
	`
	// Query execution, works just fine
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []dictionary.BannerCategory
	for rows.Next() {
		var data dictionary.BannerCategory
		rows.Scan(
			&data.ID,
			&data.Name,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}
	return result, nil
}
