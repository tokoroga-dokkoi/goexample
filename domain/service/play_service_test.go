package domain_service

import (
	"testing"

	. "github.com/MikiWaraMiki/goexample/domain/model"
)

func TestFetchByPlayIdWhenExist(t *testing.T) {
	plays := []Play{
		{
			PlayID:   "test",
			Name:     "test",
			TypeName: "tragedy",
		},
	}

	play_service := PlayService{
		plays: plays,
	}

	result, err := play_service.fetchByPlayId("test")

	if err != nil {
		t.Fatal("fetchByPlayId() is nil return")
	}

	if result.PlayID != "test" {
		t.Errorf("fetchByPlayId() = %v, want %v", result.PlayID, "test")
	}
}
