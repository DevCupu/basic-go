package book

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" gorm:"not null"`
	Author string `json:"author" gorm:"not null"`
	Year   int    `json:"year"`
	Stock  int    `json:"stock" gorm:"default:0"`
}

// Method helper: Cek apakah buku masih bisa dipinjam
func (b *Book) IsAvailable() bool {
	return b.Stock > 0
}
