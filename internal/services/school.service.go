package services

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/api-rest-go/internal/types"
	supabase "github.com/supabase-community/supabase-go"
)

func newSupabaseClient() (*supabase.Client, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		return nil, errors.New("missing SUPABASE env vars")
	}

	return supabase.NewClient(supabaseURL, supabaseKey, nil)
}

func GetAllSchoolsService() []types.School {
	client, err := newSupabaseClient()
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
	if err := json.Unmarshal(data, &schools); err != nil {
		return nil
	}

	return schools
}

func CreateSchoolsService(
	schools []types.CreateSchoolRequest,
) ([]types.School, error) {

	client, err := newSupabaseClient()
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

func DeleteSchoolService(schoolID string) error {
	client, err := newSupabaseClient()
	if err != nil {
		return err
	}

	_, _, err = client.
		From("schools").
		Delete("*", "").
		Eq("uuid", schoolID).
		Execute()

	return err
}
