package database

import (
"fmt"
"log"
"regexp"
"testing"

"github.com/jinzhu/gorm"
"github.com/stretchr/testify/assert"
"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestContains(t *testing.T) {
	testSlice := []string{"test string", "test string 2"}
	presentedSring := "test string"
	notPresentedSring := "not presented strind"
	switch {
	case !contains(testSlice, presentedSring):
		t.Errorf("The slice \"%s\" contains \"%s\", but the function returns false.", testSlice, presentedSring)
	case contains(testSlice, notPresentedSring):
		t.Errorf("The slice \"%s\" does not contain \"%s\", but the function returns true.", testSlice, notPresentedSring)
	case contains(testSlice, ""):
		t.Errorf("The slice \"%s\" does not contain empty string, but the function returns true.", testSlice)
	}
}

func newDB() (sqlmock.Sqlmock, *gorm.DB, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
		return nil, nil, err
	}
	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		log.Fatalf("can't open gorm connection: %s", err)
		return nil, nil, err
	}
	gormDB.LogMode(true)

	return mock, gormDB.Set("gorm:update_column", true), nil
}

func formatRequest(s string) string {
	return fmt.Sprintf("^%s$", regexp.QuoteMeta(s))
}

func TestAllProducts(t *testing.T) {
	m, db, _ := newDB()
	var productFieldNames = []string{"id", "name", "image", "shelfLife", "units"}
	rows := sqlmock.NewRows(productFieldNames)
	p := Product{ID: "id", Name: "name", Image: "image", ShelfLife: 2, Units: "units"}
	rows = rows.AddRow(p.ID, p.Name, p.Image, p.ShelfLife, p.Units)
	expProducts := []Product{p}
	m.ExpectQuery(formatRequest("SELECT products.id, products.name, products.image, products.shelf_life, m_units.unit FROM \"products\" LEFT JOIN m_units on m_units.id = products.units")).
		WillReturnRows(rows)
	products, _ := AllProducts(db)
	assert.Equal(t, products, expProducts)
}

func TestFindProductByName(t *testing.T) {
	m, db, err := newDB()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}
	var productFieldNames = []string{"id", "testProduct", "image", "shelfLife", "units"}
	rows := sqlmock.NewRows(productFieldNames)
	p := Product{ID: "id", Name: "testProduct", Image: "image", ShelfLife: 3, Units: "units"}
	rows = rows.AddRow(p.ID, p.Name, p.Image, p.ShelfLife, p.Units)
	m.ExpectQuery(formatRequest("SELECT products.id, products.name, products.image, products.shelf_life, m_units.unit FROM \"products\" LEFT JOIN m_units on m_units.id = products.units WHERE (name = $1)")).WithArgs("testproduct").
		WillReturnRows(rows)
	product, err := FindProductByName("testProduct", db)
	if err != nil {
		log.Fatalf("can't find product by name: %s", err)
	}
	assert.EqualValues(t, product, &p)
}

func TestFindProductByID(t *testing.T) {
	m, db, err := newDB()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}
	var productFieldNames = []string{"testID", "testProduct", "image", "shelfLife", "units"}
	rows := sqlmock.NewRows(productFieldNames)
	p := Product{ID: "id", Name: "testProduct", Image: "image", ShelfLife: 3, Units: "units"}
	rows = rows.AddRow(p.ID, p.Name, p.Image, p.ShelfLife, p.Units)
	m.ExpectQuery(formatRequest("SELECT products.id, products.name, products.image, products.shelf_life, m_units.unit FROM \"products\" LEFT JOIN m_units on m_units.id = products.units WHERE (products.id = $1)")).WithArgs("testID").
		WillReturnRows(rows)
	product, err := FindProductByID("testID", db)
	if err != nil {
		log.Fatalf("can't find product by ID: %s", err)
	}
	assert.EqualValues(t, product, &p)
}

func TestCheckProductName(t *testing.T) {
	m, db, err := newDB()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}
	var productFieldNames = []string{"testID"}
	rows := sqlmock.NewRows(productFieldNames)
	p := Product{ID: "testID", Name: "testProduct", Image: "image", ShelfLife: 3, Units: "units"}
	rows = rows.AddRow(p.ID)
	m.ExpectQuery(formatRequest("select products.ID from products where products.name = $1")).WithArgs("testproduct").
		WillReturnRows(rows)
	err = CheckProductName("testProduct", db)
	if err != nil {
		log.Fatalf("can't find product id by name: %s", err)
	}
}

func TestDeleteProductByID(t *testing.T) {
	m, db, err := newDB()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}
	id := "testID"
	m.ExpectExec(formatRequest("DELETE FROM \"products\" WHERE (id = $1)")).
		WithArgs("testID").
		WillReturnResult(sqlmock.NewResult(0, 1))
	assert.Nil(t, DeleteProductByID(db, id))
}

func TestUpdateProduct(t *testing.T) {
	m, db, err := newDB()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}
	var productFieldNames = []string{"testID", "testProduct", "image", "shelfLife", "units"}
	rows := sqlmock.NewRows(productFieldNames)
	p := Product{ID: "id", Name: "testProduct", Image: "image", ShelfLife: 3, Units: "units"}
	rows = rows.AddRow(p.ID, p.Name, p.Image, p.ShelfLife, p.Units)
	m.ExpectQuery(formatRequest("SELECT products.id, products.name, products.image, products.shelf_life, m_units.unit FROM \"products\" LEFT JOIN m_units on m_units.id = products.units WHERE (products.id = $1)")).WithArgs("testID").
		WillReturnRows(rows)

	var mUnitFieldNames = []string{"id", "unit"}
	rowsUnits := sqlmock.NewRows(mUnitFieldNames)
	mu := MUnit{ID: "testUnitID", Unit: "testUnit"}
	rowsUnits = rowsUnits.AddRow(mu.ID, mu.Unit)
	m.ExpectQuery(formatRequest("SELECT * FROM \"m_units\"  WHERE (unit = 'ml') ORDER BY \"m_units\".\"id\" ASC LIMIT 1")).WithArgs("testUnit").
		WillReturnRows(rowsUnits)

	m.ExpectQuery(formatRequest("INSERT INTO \"products\" (\"name\",\"shelf_life\",\"units\",\"image\") VALUES ('name','2','a33bbf4b-dc4a-427c-8ac8-628945e2efa5','image') RETURNING \"products\".\"id\"")).WithArgs("testUnit").
		WillReturnRows(rowsUnits)

	product, err := FindProductByID("testID", db)
	if err != nil {
		log.Fatalf("can't find product by ID: %s", err)
	}
	assert.EqualValues(t, product, &p)
}
