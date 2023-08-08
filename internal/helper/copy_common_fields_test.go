package helper

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func TestCopyCommonFields(t *testing.T) {
	type userSrc struct {
		Name string
		Age  int
	}
	type userDest struct {
		Name string
		Age  int
	}

	src := userSrc{
		Name: "Test",
		Age:  25,
	}
	dest := userDest{}

	CopyCommonFields(&dest, &src)

	if dest.Name != "Test" || dest.Age != 25 {
		log.Fatal().Msg("Src and dest is not same")
	}
}
