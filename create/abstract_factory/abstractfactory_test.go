package abstract_factory

import (
	"testing"
)

func TestDbDaoFactory(t *testing.T) {
	var factory DaoFactory
	factory = DbDaoFactory{}
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}
