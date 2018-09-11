package abstract_factory

import (
	"fmt"
)

// 订单主dao
type OrderMainDao interface {
	SaveOrderMain()
}

// 订单详情dao
type OrderDetailDao interface {
	SaveOrderDetail()
}

// 订单dao的factory
type DaoFactory interface {
	CreateOrderMainDAO() OrderMainDao
	CreateOrderDetailDAO() OrderDetailDao
}

// ====================使用数据库存储的订单的实现==================
// 实现OrderMainDao接口
type DbMainDao struct {
}

func (mainDao DbMainDao) SaveOrderMain() {
	fmt.Println("db main save")
}

// 实现OrderDetailDao接口
type DbDetailDao struct {
}

func (detailDao DbDetailDao) SaveOrderDetail() {
	fmt.Println("db detail save")
}

//实现DaoFactory接口，即DbDaoFactory是抽象工厂的实现
type DbDaoFactory struct {
}

func (dbDaoFactory DbDaoFactory) CreateOrderMainDAO() OrderMainDao {
	return DbMainDao{}
}
func (dbDaoFactory DbDaoFactory) CreateOrderDetailDAO() OrderDetailDao {
	return DbDetailDao{}
}

// ========================使用XML的订单的实现=============================
// todo
