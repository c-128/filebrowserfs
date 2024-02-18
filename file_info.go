package filebrowserfs

type FileInfo struct {
	Path      string `json:"path"`
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Extension string `json:"extension"`
	Modified  string `json:"modified"`
	FileMode  int    `json:"mode"`
	IsDir     bool   `json:"IsDir"`
	IsSymlink bool   `json:"IsSymlink"`
	Type      string `json:"type"`
}
