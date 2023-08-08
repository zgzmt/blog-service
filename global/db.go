package global

import "github.com/jinzhu/gorm"

var(
	DBEngine *gorm.DB
)

func test(){
	fmt.Print("test")
}
