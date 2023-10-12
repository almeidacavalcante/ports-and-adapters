package db

import (
	"database/sql"
	"github.com/almeidacavalcante/ports-and-adapters/application"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, status, price from products where id = $1")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Status, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	err := p.db.QueryRow("select count(id) from products where id = $1", product.GetID()).Scan(&rows)
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(`insert into products(id, name, status, price) values($1, $2, $3, $4)`)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetStatus(), product.GetPrice())
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(`update products set name = $1, status = $2, price = $3 where id = $4`)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetName(), product.GetStatus(), product.GetPrice(), product.GetID())
	if err != nil {
		return nil, err
	}
	return product, nil
}
