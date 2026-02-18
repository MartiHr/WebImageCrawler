package storage

import (
	"database/sql"
)

type ImageRecord struct {
	ID            int
	URL           string
	Filename      string
	Alt           string
	Title         string
	Width         int
	Height        int
	Format        string
	ThumbnailPath string
}

type ImageRepository struct {
	db *sql.DB
}

func NewImageRepository(db *sql.DB) *ImageRepository {
	return &ImageRepository{db: db}
}

func (r *ImageRepository) Insert(img ImageRecord) error {
	_, err := r.db.Exec(`
       INSERT IGNORE INTO images
       (url, filename, alt, title, width, height, format, thumbnail_path)
       VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		img.URL,
		img.Filename,
		img.Alt,
		img.Title,
		img.Width,
		img.Height,
		img.Format,
		img.ThumbnailPath,
	)
	return err
}

func (r *ImageRepository) SearchAll(filters map[string]string) ([]ImageRecord, error) {
	query := `
			SELECT id, url, filename, alt, title, width, height, format, thumbnail_path
				FROM images
			WHERE 1=1
	`

	args := []any{}

	if v := filters["url"]; v != "" {
		query += " AND url LIKE ?"
		args = append(args, "%"+v+"%")
	}
	if v := filters["filename"]; v != "" {
		query += " AND filename LIKE ?"
		args = append(args, "%"+v+"%")
	}
	if v := filters["alt"]; v != "" {
		query += " AND alt LIKE ?"
		args = append(args, "%"+v+"%")
	}
	if v := filters["title"]; v != "" {
		query += " AND title LIKE ?"
		args = append(args, "%"+v+"%")
	}
	if v := filters["format"]; v != "" {
		query += " AND format = ?"
		args = append(args, v)
	}
	if v := filters["min_width"]; v != "" {
		query += " AND width >= ?"
		args = append(args, v)
	}
	if v := filters["max_width"]; v != "" {
		query += " AND width <= ?"
		args = append(args, v)
	}
	if v := filters["min_height"]; v != "" {
		query += " AND height >= ?"
		args = append(args, v)
	}
	if v := filters["max_height"]; v != "" {
		query += " AND height <= ?"
		args = append(args, v)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []ImageRecord
	for rows.Next() {
		var img ImageRecord
		rows.Scan(
			&img.ID,
			&img.URL,
			&img.Filename,
			&img.Alt,
			&img.Title,
			&img.Width,
			&img.Height,
			&img.Format,
			&img.ThumbnailPath,
		)
		results = append(results, img)
	}

	return results, nil
}
