package filebrowserfs

type DirectoryListing struct {
	Files      []FileInfo `json:"items"`
	NumDirs    int        `json:"numDirs"`
	NumFiles   int        `json:"numFiles"`
	Path       string     `json:"path"`
	Name       string     `json:"name"`
	Size       int        `json:"size"`
	Extension  string     `json:"extension"`
	Modified   string     `json:"modified"`
	Mode       uint64     `json:"mode"`
	IsDir      bool       `json:"isDir"`
	IsSymlinkg bool       `json:"isSymlink"`
	Type       string     `json:"type"`
}
