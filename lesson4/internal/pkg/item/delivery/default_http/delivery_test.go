package default_http

import (
	"backendOneLessons/lesson4/internal/pkg/item/usecase"
	"backendOneLessons/lesson4/internal/pkg/models"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDelivery_ServeHTTP(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/items/?max", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	items := usecase.NewInmemory()

	itemHandler := New(items, nil)

	itemHandler.ServeHTTP(rr, req)

	require.Equalf(t, http.StatusOK, rr.Code, "unexpected status")

	expectedItems, err := items.List(context.Background(), models.ItemFilter{})
	require.NoError(t, err)

	gotItems := make([]models.Item, 0, len(expectedItems))
	err = json.NewDecoder(rr.Body).Decode(&gotItems)

	require.Equal(t, expectedItems, gotItems)

}
