package main

// Interfaces

type Product struct {
    ID int
    Name string
    Price float64
}

type ProductRepository interface {
    GetByID(id int) (*Product, error)
    GetAll() ([]*Product, error)
    Create(p *Product) error
    Update(p *Product) error
    Delete(id int) error
}

type ProductService struct {
    repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

func (s *ProductService) GetByID(id int) (*Product, error) {
    return s.repo.GetByID(id)
}

func (s *ProductService) GetAll() ([]*Product, error) {
    return s.repo.GetAll()
}

func (s *ProductService) Create(p *Product) error {
    return s.repo.Create(p)
}

func (s *ProductService) Update(p *Product) error {
    return s.repo.Update(p)
}

func (s *ProductService) Delete(id int) error {
    return s.repo.Delete(id)
}

// Implementations

type InMemoryProductRepository struct {
    products map[int]*Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
    return &InMemoryProductRepository{products: make(map[int]*Product)}
}

func (r *InMemoryProductRepository) GetByID(id int) (*Product, error) {
    p, ok := r.products[id]
    if !ok {
        return nil, fmt.Errorf("product not found")
    }
    return p, nil
}

func (r *InMemoryProductRepository) GetAll() ([]*Product, error) {
    products := make([]*Product, 0, len(r.products))
    for _, p := range r.products {
        products = append(products, p)
    }
    return products, nil
}

func (r *InMemoryProductRepository) Create(p *Product) error {
    _, ok := r.products[p.ID]
    if ok {
        return fmt.Errorf("product already exists")
    }
    r.products[p.ID] = p
    return nil
}

func (r *InMemoryProductRepository) Update(p *Product) error {
    _, ok := r.products[p.ID]
    if !ok {
        return fmt.Errorf("product not found")
    }
    r.products[p.ID] = p
    return nil
}

func (r *InMemoryProductRepository) Delete(id int) error {
    _, ok := r.products[id]
    if !ok {
        return fmt.Errorf("product not found")
    }
    delete(r.products, id)
    return nil
}
