package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestDownloadFile проверяет скачивание и сохранение файла
func TestDownloadFile(t *testing.T) {
	// Создаем временный HTTP-сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test file content"))
	}))
	defer server.Close()

	// Создаем временный файл для сохранения
	tempDir := t.TempDir()
	outputPath := filepath.Join(tempDir, "testfile.txt")

	// Тестируем функцию
	err := downloadFile(server.URL, outputPath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Проверяем содержимое файла
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read downloaded file: %v", err)
	}

	if string(content) != "Test file content" {
		t.Errorf("Expected file content 'Test file content', got '%s'", string(content))
	}
}

// TestParseLinks проверяет извлечение ссылок из HTML
func TestParseLinks(t *testing.T) {
	// Пример HTML-контента
	htmlContent := `<html>
		<body>
			<a href="https://example.com/page1">Page 1</a>
			<a href="/page2">Page 2</a>
		</body>
	</html>`

	// Создаем ридер
	reader := strings.NewReader(htmlContent)
	baseURL := "https://example.com"

	links, err := parseLinks(reader, baseURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedLinks := []string{
		"https://example.com/page1",
		"https://example.com/page2",
	}

	if len(links) != len(expectedLinks) {
		t.Fatalf("Expected %d links, got %d", len(expectedLinks), len(links))
	}

	for i, link := range links {
		if link != expectedLinks[i] {
			t.Errorf("Expected link %s, got %s", expectedLinks[i], link)
		}
	}
}

// TestDownloadSite проверяет рекурсивное скачивание сайта
func TestDownloadSite(t *testing.T) {
	// Создаем временный HTTP-сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<html><body><a href="/page1">Page 1</a></body></html>`))
		case "/page1":
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`Page 1 Content`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	// Параметры теста
	outputDir := t.TempDir()

	err := downloadSite(server.URL, outputDir, 2)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Формируем корректные пути
	host := strings.ReplaceAll(strings.Split(server.URL, "://")[1], ":", "_")
	indexPath := filepath.Join(outputDir, host, "index.html")
	page1Path := filepath.Join(outputDir, host, "page1")

	// Проверяем существование файлов
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Errorf("Expected index.html to exist at %s, but it does not", indexPath)
	}

	if _, err := os.Stat(page1Path); os.IsNotExist(err) {
		t.Errorf("Expected page1 to exist at %s, but it does not", page1Path)
	}

	// Проверяем содержимое файлов
	indexContent, _ := os.ReadFile(indexPath)
	if !strings.Contains(string(indexContent), "Page 1") {
		t.Errorf("Expected index.html to contain 'Page 1', got %s", string(indexContent))
	}

	page1Content, _ := os.ReadFile(page1Path)
	if !strings.Contains(string(page1Content), "Page 1 Content") {
		t.Errorf("Expected page1 to contain 'Page 1 Content', got %s", string(page1Content))
	}
}
