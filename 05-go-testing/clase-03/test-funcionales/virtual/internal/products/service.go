package products
type Service interface {
	Store(nombre string, tipo string, cantidad int, precio float64) (Product, error)
 }
 type service struct {
	repository Repository
 }
 func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
 }
func (s *service) Store(nombre, tipo string, cantidad int, precio float64) (Product, error) {

	producto, err := s.repository.Store(nombre, tipo, cantidad, precio)
	if err != nil{
		return Product{}, err
	}
 
	return producto, nil
 }
 