package book

type Service interface {
	AddBook(input BookInput) (*Book, error)
	GetAllBooks() ([]Book, error)
	GetBook(id uint) (*Book, error)
	UpdateBook(id uint, input BookInput) (*Book, error)
	DeleteBook(id uint) error
}
type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) AddBook(input BookInput) (*Book, error) {
	book := &Book{
		Title:  input.Title,
		Author: input.Author,
		Year:   input.Year,
		Stock:  input.Stock,
	}

	err := s.repo.Create(book)

	return book, err
}

func (s *service) GetAllBooks() ([]Book, error) {
	return s.repo.FindAll()
}

func (s *service) GetBook(id uint) (*Book, error) {
	return s.repo.FindByID(id)
}

func (s *service) UpdateBook(id uint, input BookInput) (*Book, error) {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	book.Title = input.Title
	book.Author = input.Author
	book.Year = input.Year
	book.Stock = input.Stock

	err = s.repo.Update(book)

	return book, err
}

func (s *service) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}
