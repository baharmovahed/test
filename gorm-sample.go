package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)


type Product struct {
  gorm.Model
  Code string
  Price uint
}

func main() {
  db, err := gorm.Open("mysql", "root:@/iotree?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()

  // Migrate the schema
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "12", Price: 200})

//   // Read
//   var product Product
//   db.First(&product, 1) // find product with id 1
//   db.First(&product, "code = ?", "L1212") // find product with code l1212

//   // Update - update product's price to 2000
//   db.Model(&product).Update("Price", 2000)

//   // Delete - delete product
//   db.Delete(&product)
}