package services

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsService struct {
}

type itemsServiceInterface interface {
	GetItem()
}

func (s *itemsService) GetItem() {

}
