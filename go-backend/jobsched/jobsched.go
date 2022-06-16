package main

type GithubRemote struct {
	ID          string `gorm:"column:id;primary_key" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Url         string `gorm:"column:url" json:"url"`
	HomepageURL string `gorm:"column:homepageurl" json:"homepageUrl"`
	Description string `gorm:"column:description" json:"description"`
	Diskusage   int    `gorm:"column:diskusage" json:"diskUsage"`
}

func main() {
	querySync()
}
