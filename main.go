package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

const (
	dataDir    = "data"
	dateLayout = "2006-01-02"
)

var brackets = []int{200, 500, 1000, 2000, 5000, 10000, 20000, 50000, 100000, 200000, 500000, 1000000}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "radar: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	date := time.Now().Format(dateLayout)
	outputRoot := dataDir

	outputDir := filepath.Join(outputRoot, date)
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("create output directory %s: %w", outputDir, err)
	}

	apiToken, err := cloudflareAPIToken(".env")
	if err != nil {
		return err
	}

	client := cloudflare.NewClient(option.WithAPIToken(apiToken))
	for _, bracket := range brackets {
		if err := downloadTopDomains(ctx, client, outputDir, bracket); err != nil {
			return err
		}
	}

	return nil
}

func downloadTopDomains(ctx context.Context, client *cloudflare.Client, outputDir string, bracket int) error {
	alias := fmt.Sprintf("ranking_top_%d", bracket)
	body, err := client.Radar.Datasets.Get(ctx, alias)
	if err != nil {
		return fmt.Errorf("download %s: %w", alias, err)
	}
	if body == nil {
		return fmt.Errorf("download %s: empty response", alias)
	}

	filename := fmt.Sprintf("top-%d-domains.csv", bracket)
	path := filepath.Join(outputDir, filename)
	if err := os.WriteFile(path, []byte(withoutHeader(*body)), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	fmt.Printf("downloaded %s\n", path)
	return nil
}

func withoutHeader(csv string) string {
	_, rest, _ := strings.Cut(csv, "\n")
	return rest
}

func cloudflareAPIToken(path string) (string, error) {
	const key = "CLOUDFLARE_API_TOKEN="
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", path, err)
	}
	raw := strings.TrimSpace(string(data))
	if !strings.HasPrefix(raw, key) {
		return "", fmt.Errorf("%s must contain %s...", path, key)
	}
	token := strings.TrimPrefix(raw, key)
	if token == "" {
		return "", fmt.Errorf("%s must contain %s...", path, key)
	}
	return token, nil
}
