//package main
//
//import (
//	"flaxxed/function"
//	"github.com/GoogleCloudPlatform/function-framework-go/funcframework"
//	"log"
//	"os"
//)

////
//import (
//	_ "cloud.google.com/go/cache/apiv1"
//	_ "github.com/jinzhu/gorm/dialects/mysql"
//)
//

//
//func main() {
//	tagDao:=dao.TagsDaoImpl{}
//	 res, _:=tagDao.FindByString("broiled deer game meat")
//	 for _, t := range res{
//	 	fmt.Println(t.Tag)
//	 }
//
//	parRes, _:=tagDao.FindParentTags("broiled cooked deer game meat")
//	for _, t := range parRes{
//		fmt.Println(t.Tag)
//	}

//db, err := gorm.Open("mysql", "root:flaxxed@/flaxxed?charset=utf8&parseTime=True&loc=Local")
//if err != nil {
//	fmt.Print(err)
//	os.Exit(1)
//}
////vegetarian
////var food data.Food
//db.LogMode(true)
//var tag data.Tag
//var param = "apple"
//db.Set("gorm:auto_preload", true).Where("tag = ?", param).First(&tag)
////db.Set("gorm:auto_preload", true).Debug().First(&food, 1)
////fmt.Println(food.Name)
////var another data.Food
////
////db.Set("gorm:auto_preload", true).Debug().First(&another, 2)
//fmt.Println(len(tag.TaggedFood))
//fmt.Println(tag.TaggedFood[0])
//fmt.Println(tag.FoodCategory.Name)
//
////fmt.Println(another.FoodCategory.Name)
//defer db.Close()
//}
package main

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"log"
	"os"
	console "redis_client/dist/redis"
)

func main() {

	//ctx:=context.Background()
	//var client *firestore.Client
	//var err error
	//
	//client, err = firestore.NewClient(ctx, "flaxxed")
	//if err != nil {
	//	log.Fatal("can't initialize the firestore client", err)
	//}
	////colRef:=client.Collection("food")
	//sl:=make([]*firestore.DocumentRef, 0)
	//start := time.Now()
	//for i:=1; i< 8791; i++{
	//	docRef, err:=client.Collection("food").Doc(strconv.Itoa(i)).Get(ctx)
	//	if err!= nil{
	//		fmt.Println(err, "-----", i)
	//	}
	//	data:=docRef.Data()
	//	if data==nil {
	//		fmt.Println( "-----", i)
	//	}
	//}
	//generationEnd:=time.Now().Nanosecond()
	//fmt.Println(generationEnd - start.Nanosecond())
	//
	//food, err:=client.GetAll(ctx, sl)
	//if err!= nil{
	//	fmt.Println(err)
	//}else {
	//	for _,d:=range food{
	//		fmt.Println(d.Data()["Name"])
	//	}
	//}
	//fmt.Println(time.Now().Nanosecond() - start.Nanosecond())
	funcframework.RegisterHTTPFunction("/files", console.Execute)
	//funcframework.RegisterHTTPFunction("/selection", selection.CalculateAndSave)

	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
