package models

import "fmt"

var tablePrefix = "md_"

func init() {

}

func TNCategory() string {
	return fmt.Sprintf("%s%s", tablePrefix, "category")
}

func TNBook() string {
	return fmt.Sprintf("%s%s", tablePrefix, "books")
}

func TNBookCategory() string {
	return fmt.Sprintf("%s%s", tablePrefix, "book_category")
}
