package handlers

import (
	"context"
	modelAPI "evendo-viator/pkg/model/api/destination"
	model "evendo-viator/pkg/model/api/products"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type SimpleFindProduct struct {
	ProductCode string `json:"product_code" bson:"product_code"`
}

func (h *CronJobHandler) FormatOldDatabase(client *mongo.Client, offset int64) {
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	for i := 1; i < 1000; i++ {
		log.Println(i)
		cur, err := h.ProductMongoDB.GetMissingMany(client, offset)
		if err != nil {
			log.Println(err)
			continue
		}

		for cur.Next(context.Background()) {
			t := SimpleFindProduct{}
			err = cur.Decode(&t)
			if err != nil {
				log.Println(err)
				continue
			}

			productDetail, err := h.ViatorAPIHandler.GetProductByCode(t.ProductCode)
			if err != nil {
				f.WriteString(t.ProductCode + " got error " + err.Error())
				log.Println(err)
				if strings.Contains(err.Error(), "Body Too Many Requests") {
					time.Sleep(time.Second * 120)
				}

				continue
			}
			locationDetail := h.HandlerGetListOfLocationV2(*productDetail)
			filter := bson.D{{"product_code", t.ProductCode}}
			update := bson.D{{"$set", bson.D{{"location", locationDetail}, {"detail", productDetail}, {"created_at", time.Now().Unix()}}}}
			if _, err := client.Database("viator").Collection("product").UpdateOne(context.Background(), filter, update); err != nil {
				log.Println(err)
			}
		}
	}

}

func (h *CronJobHandler) CustomQuickQuery(client *mongo.Client) {

	wg := &sync.WaitGroup{}
	wg.Add(3)
	var SORT_TYPE_PRICE = "PRICE"
	var SORT_ORDER_ASCENDING = "ASCENDING"
	destination := modelAPI.APIDestinationDetail{
		SortOrder:           3094,
		Selectable:          true,
		DestinationURLName:  "Spain",
		DefaultCurrencyCode: "EUR",
		LookupID:            "6.67",
		ParentID:            6,
		Timezone:            "Europe/Madrid",
		IataCode:            "",
		DestinationName:     "Spain",
		DestinationID:       67,
		DestinationType:     "COUNTRY",
		Latitude:            40.463667,
		Longitude:           -3.74922,
	}
	for i := 7; i <= 9; i++ {
		go func(index int) {
			defer wg.Done()
			for _, start := range []int{1, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000} {
				log.Println("WORKING")

				start = start + (index * 1000)
				log.Println(start)
				if start == 10000 {
					return
				}

				searchRequest := model.SearchProductRequest{
					Filtering: model.FilteringDetail{
						Destination: strconv.Itoa(destination.DestinationID),
						Tags:        []int{},
						Flags:       []string{},
					},
					Sorting: model.SortingDetail{
						Sort:  &SORT_TYPE_PRICE,
						Order: &SORT_ORDER_ASCENDING,
					},
					Pagination: model.PaginationDetail{
						Start: start,
						Count: 100,
					},
					Currency: "USD",
				}
				productList, err := h.ViatorAPIHandler.ProductSearch(searchRequest)
				if err != nil {
					log.Println(err)
					continue
				}
				h.InsertProduct(client, productList, "USD", destination)
			}
		}(i)
	}
	wg.Wait()
	log.Println("ALL DONE")
}
