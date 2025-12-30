package services

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/api-rest-go/internal/types"
	supabase "github.com/supabase-community/supabase-go"
)

func GetAllSchoolsService() []types.School {

	err := godotenv.Load()

	if err != nil {
		log.Println("No se encontró archivo .env")
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_key")

	client, err := supabase.NewClient(supabaseUrl, supabaseKey, nil)
	if err != nil {
		return nil
	}

	data, _, err := client.
		From("schools").
		Select("*", "", false).
		Execute()

	if err != nil {
		return nil
	}

	var schools []types.School
	err = json.Unmarshal(data, &schools)
	if err != nil {
		return nil
	}

	return schools
}

func CreateSchoolsService(schools []types.CreateSchoolRequest) ([]types.School, error) {
	err := godotenv.Load()

	if err != nil {
		log.Println("No se encontró archivo .env")
	}

	client, err := supabase.NewClient(
		os.Getenv("SUPABASE_URL"),
		os.Getenv("SUPABASE_key"),
		nil,
	)

	if err != nil {
		return nil, err
	}

	data, _, err := client.
		From("schools").
		Insert(schools, false, "", "", "").
		Execute()

	if err != nil {
		return nil, err
	}

	var created []types.School
	if err := json.Unmarshal(data, &created); err != nil {
		return nil, err
	}

	return created, nil
}
