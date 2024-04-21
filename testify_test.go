package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusOK, responseRecorder.Code)           // Проверяем код ответа
	assert.NotEmpty(t, responseRecorder.Body.String())               // Проверяем, что тело ответа не пустое
	assert.Equal(t, totalCount, len(responseRecorder.Body.String())) //Сверяем количество кафе
	assert.Equal(t, cafeList, responseRecorder.Body.String())        // Проверяем все доступные кафе
}

func TestMainHandlerWhenCityIsWrong(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=kazan", nil) // Создаем запрос с неподдерживаемым городом kazan

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code) // Проверяем код ответа

	assert.Contains(t, responseRecorder.Body.String(), "wrong city value") // Проверяем сообщение об ошибке для неподдерживаемого города
}

func TestMainHandlerIsRequestCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code) // Проверяем код ответа
	assert.NotEmpty(t, responseRecorder.Body.String())     //Проверяем, что тело ответа не пустое
}
