package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	
	"github.com/zeze322/wt-guided-weaponry/internal/api"
	"github.com/zeze322/wt-guided-weaponry/internal/db/mongodb"
	"github.com/zeze322/wt-guided-weaponry/internal/db/postgresdb"
)

func TestHandleCategories(t *testing.T) {
	postgres, err := postgresdb.New(context.Background(), os.Getenv("POSTGRES_URL"))
	require.NoError(t, err)

	mongo, err := mongodb.New(context.Background(), os.Getenv("MONGO_URL"), os.Getenv("MONGODB_DATABASE"), os.Getenv("MONGODB_COLLECTION"))
	require.NoError(t, err)

	s := api.NewServer(logrus.New(), os.Getenv("PORT"), postgres, mongo)

	rr := httptest.NewRecorder()
	defer rr.Result().Body.Close()

	req, err := http.NewRequest("GET", "/categories", nil)
	require.NoError(t, err)
	defer req.Body.Close()

	err = s.HandleCategories(rr, req)
	require.NoError(t, err)

	var categories api.CategoriesResponse

	err = json.NewDecoder(rr.Body).Decode(&categories)
	require.NoError(t, err)
	require.NotEmpty(t, categories)

	fmt.Println(categories)
}
