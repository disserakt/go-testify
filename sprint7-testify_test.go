package main

import (
    "net/http"
    "net/http/httptest"
    "strconv"
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

var cafeList = map[string][]string{
    "moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
    countStr := req.URL.Query().Get("count")
    if countStr == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("count missing"))
        return
    }

    count, err := strconv.Atoi(countStr)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("wrong count value"))
        return
    }

    city := req.URL.Query().Get("city")

    cafe, ok := cafeList[city]
    if !ok {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("wrong city value"))
        return
    }

    if count > len(cafe) {
        count = len(cafe)
    }

    answer := strings.Join(cafe[:count], ",")

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(answer))
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    // здесь нужно создать запрос к сервису
    req := httptest.NewRequest("GET", "/cafe?count=6&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    // здесь нужно добавить необходимые проверки
    
    if status := responseRecorder.Code; status != http.StatusOK {
        t.Errorf("expected status code: %d, got %d", http.StatusOK, status)
    }
    
    if status := responseRecorder.Code; status != http.StatusBadRequest {
        t.Errorf("expected status code: %d, got %d", http.StatusBadRequest, status)
    }

    expected := `wrong city value`
    if responseRecorder.Body.String() != expected {
        t.Errorf("expected body: %s, got %s", expected, responseRecorder.Body.String())
    }
    
    countStr := req.URL.Query().Get("count")
    count, _ := strconv.Atoi(countStr)
    if count >= totalCount {
        t.Errorf("expected body: %s, got %s", expected, responseRecorder.Body.String())
    }
}
