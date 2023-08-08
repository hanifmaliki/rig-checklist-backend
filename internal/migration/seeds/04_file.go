package seeds

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/pkg/minio"
	"gorm.io/gorm"
)

func SeedFile(db *gorm.DB) error {
	// Get a list of all files in the local disk folder
	localFolder := "./assets/seed-file"
	files, err := os.ReadDir(localFolder)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	// Loop through all the files in the local folder and upload them to MinIO
	for _, file := range files {
		// Only upload regular files (not directories)
		if !file.IsDir() {
			filePath := filepath.Join(localFolder, file.Name())

			// Upload the file to MinIO
			err = minio.FPutObject(file.Name(), filePath)
			if err != nil {
				log.Error().Msg(err.Error())
				return err
			}
		}
	}

	return nil
}
