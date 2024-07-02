package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/storage"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	torrentFilePath := filepath.Join(currentDir, "testtorrent.torrent")
	downloadDir := filepath.Join(currentDir, "downloads")

	err = os.MkdirAll(downloadDir, 0755)
	if err != nil {
		fmt.Printf("Error creating download directory: %v\n", err)
		os.Exit(1)
	}

	cfg := torrent.NewDefaultClientConfig()
	cfg.DefaultStorage = storage.NewFile(downloadDir)
	cfg.NoUpload = true 
	client, err := torrent.NewClient(cfg)
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	t, err := client.AddTorrentFromFile(torrentFilePath)
	if err != nil {
		fmt.Printf("Error adding torrent: %v\n", err)
		os.Exit(1)
	}

	<-t.GotInfo()

	t.DownloadAll()

	fmt.Printf("Downloading: %s\n", t.Name())
	fmt.Printf("Size: %d bytes\n", t.Length())

	for {
		stats := t.Stats()
		progress := float64(stats.BytesReadData.Int64()) / float64(t.Length()) * 100

		fmt.Printf("\rProgress: %.2f%% | Download Speed: %s/s | Peers: %d",
			progress,
			bytesToSize(stats.BytesReadData.Int64()-stats.BytesReadUsefulData.Int64()),
			stats.TotalPeers)

		if t.BytesMissing() == 0 {
			fmt.Println("\nDownload completed!")
			break
		}

		time.Sleep(time.Second)
	}
}

func bytesToSize(bytes int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}
	if bytes == 0 {
		return "0B"
	}
	i := 0
	for bytes >= 1024 && i < len(sizes)-1 {
		bytes /= 1024
		i++
	}
	return fmt.Sprintf("%d%s", bytes, sizes[i])
}
