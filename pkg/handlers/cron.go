package handlers

import (
	"context"
	"encoding/json"
	"evendo-viator/conf"
	"evendo-viator/pkg/handlers/api"
	modelAPI "evendo-viator/pkg/model/api/destination"
	model "evendo-viator/pkg/model/api/products"
	"evendo-viator/pkg/repo"
	"evendo-viator/pkg/utils"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/carlescere/scheduler"
	"gitlab.com/jfcore/common/ginext"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CronJobHandler struct {
	CronJobInfo      map[string]map[string]interface{}
	CronJob          map[string]scheduler.Job
	CronScheduleInfo map[string]map[string]interface{}
	CronSchedule     map[string]*time.Timer
	ViatorAPIHandler *api.ViatorAPIHandlers
	ProductMongoDB   *repo.ProductMongo
}

func NewCronJobHandlers(
	ViatorAPIHandler *api.ViatorAPIHandlers,
	ProductMongoDB *repo.ProductMongo,
) *CronJobHandler {
	return &CronJobHandler{
		ViatorAPIHandler: ViatorAPIHandler,
		ProductMongoDB:   ProductMongoDB,
	}
}

//func (h *CronJobHandler) Test(r *ginext.Request) *ginext.Response {
//	productDetail, err := h.ViatorAPIHandler.GetProductByCode("67431P9")
//	log.Println(err)
//	log.Println(productDetail)
//	return ginext.NewResponseData(common.CODE_SUCCESS, "Success")
//}

func (h *CronJobHandler) StartCron() {
	Job := make(map[string]scheduler.Job)
	JobDetail := make(map[string]map[string]interface{})

	JobTask := func() {
		//var err error
		log.Println("================= Start the job ===================")
		h.GetProductFromDestinationList()
		log.Println("================= End the job ===================")
	}

	job, err := scheduler.Every(24).Hours().Run(JobTask)
	if err != nil {
		log.Println("While starting cron got error : " + err.Error())
	}
	Job["ADMIN"] = *job
	CronJobDetail := make(map[string]interface{})
	CronJobDetail["time_start"] = time.Now().Format("2006-01-02 15:04:05")
	CronJobDetail["frequency"] = "Every 8 hours"
	CronJobDetail["description"] = "Call viator to get products"

	JobDetail["ADMIN"] = CronJobDetail
	h.CronJob = Job
	h.CronJobInfo = JobDetail

}

func (h *CronJobHandler) GetProductFromDestinationList() (err error) {
	ctx := context.Background()
	mongoClient, _ := h.ConnectMongoDB(ctx)
	destinationList, err := h.ViatorAPIHandler.GetDestination()
	destinationListDetail := destinationList.Data
	var destinationListDetailCountry []modelAPI.APIDestinationDetail
	for _, destination := range destinationListDetail {
		if utils.IntContains([]int{13, 67}, destination.DestinationID) {
			destinationListDetailCountry = append(destinationListDetailCountry, destination)
		}
	}

	numberOfThread, err := strconv.Atoi(conf.GetConfig().CronJobThread)
	if err != nil {
		utils.ExitOnErr(err)
	}
	valuePerThread := len(destinationListDetailCountry) / numberOfThread
	wg := &sync.WaitGroup{}
	var finalMaxRecord int64
	wg.Add(numberOfThread)
	for i := 0; i < numberOfThread; i++ {
		var limit int64
		var offset int64
		offset = int64(i) * int64(valuePerThread)

		if i+1 == numberOfThread {
			limit = int64(len(destinationListDetailCountry)) - int64(valuePerThread)*int64(i)
		} else {
			limit = int64(valuePerThread)
		}
		finalMaxRecord = limit + offset
		go func(destinationListDetailInfo []modelAPI.APIDestinationDetail, mongoClientInfo *mongo.Client, wgInfo *sync.WaitGroup) {
			defer wgInfo.Done()
			h.GetDataBasedOnDestinationList(destinationListDetailInfo, mongoClientInfo)
		}(destinationListDetailCountry[offset:limit+offset], mongoClient, wg)
	}
	if int(finalMaxRecord) < len(destinationListDetailCountry) {
		go func(destinationListDetailInfo []modelAPI.APIDestinationDetail, mongoClientInfo *mongo.Client, wgInfo *sync.WaitGroup) {
			defer wgInfo.Done()
			h.GetDataBasedOnDestinationList(destinationListDetailInfo, mongoClientInfo)
		}(destinationListDetailCountry[finalMaxRecord:len(destinationListDetailCountry)], mongoClient, wg)
	}
	wg.Wait()
	mongoClient.Disconnect(ctx)
	return nil
}

func (h *CronJobHandler) GetDataBasedOnDestinationList(destinationListDetail []modelAPI.APIDestinationDetail, mongoClient *mongo.Client) {
	for _, destination := range destinationListDetail {
		log.Println("====================> " + fmt.Sprint(destination.DestinationID))
		numberOfProductPerThread := conf.GetConfig().NumberOfProductPerThread
		numberOfProductPerThreadParsed, err := strconv.Atoi(numberOfProductPerThread)
		if err != nil {
			numberOfProductPerThreadParsed = 5
		}
		var threshold float64
		threshold = 5.0
		startAt := 1
		for i := 1; i <= int(threshold); i++ {
			searchRequest := model.SearchProductRequest{
				Filtering: model.FilteringDetail{
					Destination: strconv.Itoa(destination.DestinationID),
					Tags:        []int{},
					Flags:       []string{},
				},
				Pagination: model.PaginationDetail{
					Start: startAt,
					Count: numberOfProductPerThreadParsed,
				},
				Currency: "USD",
			}
			if i == 1 || (i == int(threshold)) {
				productList, err := h.ViatorAPIHandler.ProductSearch(searchRequest)
				if err != nil {
					continue
				}

				threshold = math.Ceil(float64(productList.TotalCount) / float64(numberOfProductPerThreadParsed))
				startAt = i * numberOfProductPerThreadParsed
				if i == int(threshold) {
					h.InsertProduct(mongoClient, productList, "USD", destination)
				}
			}
			productList, err := h.ViatorAPIHandler.ProductSearch(searchRequest)
			if err != nil {
				continue
			}
			threshold = math.Ceil(float64(productList.TotalCount) / float64(numberOfProductPerThreadParsed))
			startAt = i * numberOfProductPerThreadParsed
			h.InsertProduct(mongoClient, productList, "USD", destination)
		}
		f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(fmt.Sprint(destination.DestinationID) + ","); err != nil {
			panic(err)
		}
		log.Println("COMPLETING : " + fmt.Sprint(destination.DestinationID))

	}

}

func (r *CronJobHandler) ConnectMongoDB(ctx context.Context) (*mongo.Client, error) {
	mongoHost := conf.GetConfig().MongoHost
	mongoPort := conf.GetConfig().MongoPort
	credential := options.Credential{
		Username: conf.GetConfig().MongoUsername,
		Password: conf.GetConfig().MongoPassword,
	}

	clientOptions := options.Client().ApplyURI("mongodb://" + mongoHost + ":" + mongoPort).SetAuth(credential)
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (h *CronJobHandler) RetrieveLanguageElement(data model.ProductDetailByCode) map[string]interface{} {
	output := make(map[string]interface{})
	output["title"] = data.Title
	output["ticketInfo"] = data.TicketInfo
	output["description"] = data.Description
	output["product_option"] = data.ProductOptions

	var listImageCaption []string
	for _, image := range data.Images {
		listImageCaption = append(listImageCaption, image.Caption)
	}
	if len(listImageCaption) > 0 {
		output["caption"] = listImageCaption
	}
	return output
}

func (h *CronJobHandler) RetrieveLanguageInformation(language string, productID string) (productLanguageInformation map[string]interface{}) {
	productDetail, err := h.ViatorAPIHandler.GetProductByCodeAndLanguage(productID, language)
	if err != nil {
		return map[string]interface{}{}
	}
	productLanguageInformation = h.RetrieveLanguageElement(*productDetail)
	return productLanguageInformation
}

func (h *CronJobHandler) InsertProduct(mongoClient *mongo.Client, productList *model.SearchProductResponse, currency string, destination modelAPI.APIDestinationDetail) {
	//var wg sync.WaitGroup
	//wg.Add(len(productList.Products))
	for _, product := range productList.Products {
		if h.ProductMongoDB.GetOne(mongoClient, product.ProductCode, fmt.Sprint(destination.DestinationID)) {
			continue
			//productDetail, err := h.ViatorAPIHandler.GetProductByCode(product.ProductCode)
			//if err != nil {
			//	log.Println(err)
			//	continue
			//}
			//locationDetail := h.HandlerGetListOfLocationV2(*productDetail)
			//filter := bson.D{{"product_code", product.ProductCode}}
			//update := bson.D{{"$set", bson.D{{"location", locationDetail}, {"detail", productDetail}, {"created_at", time.Now().Unix()}}}}
			//if _, err := mongoClient.Database("viator").Collection("product").UpdateOne(context.Background(), filter, update); err != nil {
			//	log.Println(err)
			//}
		}

		productDetail, err := h.ViatorAPIHandler.GetProductByCode(product.ProductCode)
		if err != nil {
			log.Println(err)
			return
		}
		locationDetail := h.HandlerGetListOfLocationV2(*productDetail)
		listLanguage := map[string]interface{}{"en": h.RetrieveLanguageElement(*productDetail)}
		productAvailability, _ := h.ViatorAPIHandler.AvailabilitySchedulesByProductCode(product.ProductCode)
		productSupplier, _ := h.ViatorAPIHandler.GetSupplier([]string{product.ProductCode})

		var productGeneral map[string]interface{}
		inrec, _ := json.Marshal(product)
		json.Unmarshal(inrec, &productGeneral)
		insertRequest := map[string]interface{}{
			"general":      productGeneral,
			"created_at":   time.Now().Unix(),
			"detail":       productDetail,
			"destination":  destination,
			"availability": productAvailability,
			"language":     listLanguage,
			"supplier":     productSupplier,
			"location":     locationDetail,
		}
		go h.ProductMongoDB.Create(mongoClient, product.ProductCode, insertRequest, fmt.Sprint(destination.DestinationID))
	}
}

type ReferenceTravelerPickup struct {
	Type      string      `json:"type"`
	Reference string      `json:"reference"`
	Detail    interface{} `json:"detail"`
}

type ReferenceItem struct {
	Reference string      `json:"reference"`
	Item      interface{} `json:"item"`
	Detail    interface{} `json:"detail"`
}

type ReferenceDay struct {
	Reference string      `json:"reference"`
	Day       interface{} `json:"day"`
	Detail    interface{} `json:"detail"`
}

func (h *CronJobHandler) HandlerGetListOfLocationV2(productDetail model.ProductDetailByCode) map[string]interface{} {
	locationDetailResult := make(map[string]interface{})
	var referenceList []string
	detail := productDetail.Logistics
	if len(productDetail.Itinerary.Routes) > 0 {
		for _, route := range productDetail.Itinerary.Routes {
			for _, pointOfInterest := range route.PointOfInterest {
				referenceList = append(referenceList, pointOfInterest.Location.Ref)

			}
			for _, stop := range route.Stops {
				referenceList = append(referenceList, stop.StopLocation.Ref)
			}

		}
	}

	if len(detail.Start) >= 0 {
		for _, location := range detail.Start {
			referenceList = append(referenceList, location.Location.Ref)
		}
	}

	if len(detail.End) >= 0 {
		for _, location := range detail.End {
			referenceList = append(referenceList, location.Location.Ref)
		}
	}

	if len(productDetail.Itinerary.Days) > 0 {
		for _, day := range productDetail.Itinerary.Days {
			for _, item := range day.Items {
				referenceList = append(referenceList, item.PointOfInterestLocation.Location.Ref)
			}

		}
	}

	if len(productDetail.Itinerary.ItineraryItems) > 0 {
		for _, item := range productDetail.Itinerary.ItineraryItems {
			referenceList = append(referenceList, item.PointOfInterestLocation.Location.Ref)
		}
	}

	if len(detail.TravelerPickup.Locations) > 0 {
		for _, travelerPickup := range detail.TravelerPickup.Locations {
			referenceList = append(referenceList, travelerPickup.Location.Ref)
		}
	}

	if len(productDetail.Itinerary.PointOfInterestLocations) > 0 {
		for _, point := range productDetail.Itinerary.PointOfInterestLocations {
			referenceList = append(referenceList, point.Location.Ref)
		}
	}

	if productDetail.Itinerary.ActivityInfo.Location.Ref != "" {
		referenceList = append(referenceList, productDetail.Itinerary.ActivityInfo.Location.Ref)
	}
	if len(referenceList) > 500 {
		var numberOfThread float64
		numberOfThread = float64(len(referenceList)) / 500.0
		numberOfThread = math.Ceil(numberOfThread)
		for i := 0; i < int(numberOfThread); i++ {
			end := (i * 500) + 500
			if end > len(referenceList) {
				end = len(referenceList)
			}
			result, err := h.ViatorAPIHandler.GetLocationBulk(referenceList[i*500 : end])
			if err == nil {
				if len(result.Locations) > 0 {
					for _, locationDetail := range result.Locations {
						locationDetailResult[locationDetail.Reference] = locationDetail
					}
				}
			}
		}

	} else {
		result, err := h.ViatorAPIHandler.GetLocationBulk(referenceList)
		if err == nil {
			if len(result.Locations) > 0 {
				for _, locationDetail := range result.Locations {
					locationDetailResult[locationDetail.Reference] = locationDetail
				}
			}
		}
	}

	return locationDetailResult
}

func (h *CronJobHandler) HandlerGetListOfLocation(productDetail model.ProductDetailByCode) map[string]interface{} {
	var locationDetailResult map[string]interface{}
	var referenceList []string
	var referenceStart []string
	var referenceEnd []string
	var referenceTravelerPickup []ReferenceTravelerPickup
	var referenceDays []ReferenceDay
	var referenceItems []ReferenceItem
	detail := productDetail.Logistics
	if len(productDetail.Itinerary.Routes) > 0 {
		for _, route := range productDetail.Itinerary.Routes {
			for _, pointOfInterest := range route.PointOfInterest {
				referenceItems = append(referenceItems, ReferenceItem{
					Item:      pointOfInterest,
					Reference: pointOfInterest.Location.Ref,
				})
				referenceList = append(referenceList, pointOfInterest.Location.Ref)

			}

		}
	}

	if len(detail.Start) >= 0 {
		for _, location := range detail.Start {
			referenceStart = append(referenceStart, location.Location.Ref)
			referenceList = append(referenceList, location.Location.Ref)
		}
	}

	if len(detail.End) >= 0 {
		for _, location := range detail.End {
			referenceEnd = append(referenceEnd, location.Location.Ref)
			referenceList = append(referenceList, location.Location.Ref)
		}
	}

	if len(productDetail.Itinerary.Days) > 0 {
		for _, day := range productDetail.Itinerary.Days {
			for _, item := range day.Items {
				referenceDays = append(referenceDays, ReferenceDay{
					Reference: item.PointOfInterestLocation.Location.Ref,
					Day:       day,
				})
				referenceList = append(referenceList, item.PointOfInterestLocation.Location.Ref)
			}

		}
	}

	if len(productDetail.Itinerary.ItineraryItems) > 0 {
		for _, item := range productDetail.Itinerary.ItineraryItems {
			referenceItems = append(referenceItems, ReferenceItem{
				Item:      item,
				Reference: item.PointOfInterestLocation.Location.Ref,
			})
			referenceList = append(referenceList, item.PointOfInterestLocation.Location.Ref)
		}
	}

	if len(detail.TravelerPickup.Locations) > 0 {
		for _, travelerPickup := range detail.TravelerPickup.Locations {
			referenceTravelerPickup = append(referenceTravelerPickup, ReferenceTravelerPickup{
				Type:      travelerPickup.PickupType,
				Reference: travelerPickup.Location.Ref,
			})
			referenceList = append(referenceList, travelerPickup.Location.Ref)
		}
	}

	if len(productDetail.Itinerary.PointOfInterestLocations) > 0 {
		for _, point := range productDetail.Itinerary.PointOfInterestLocations {
			referenceItems = append(referenceItems, ReferenceItem{
				Item:      point,
				Reference: point.Location.Ref,
			})
			referenceList = append(referenceList, point.Location.Ref)
		}
	}

	if productDetail.Itinerary.ActivityInfo.Location.Ref != "" {
		referenceItems = append(referenceItems, ReferenceItem{
			Item:      productDetail.Itinerary.ActivityInfo,
			Reference: productDetail.Itinerary.ActivityInfo.Location.Ref,
		})
		referenceList = append(referenceList, productDetail.Itinerary.ActivityInfo.Location.Ref)
	}

	result, err := h.ViatorAPIHandler.GetLocationBulk(referenceList)
	if err == nil {
		var listReferenceStart []interface{}
		var listReferenceEnd []interface{}
		var listReferenceTravelerPickup []interface{}
		var listReferenceDays []interface{}
		var listReferenceItems []interface{}

		for _, location := range result.Locations {
			if utils.StringContains(referenceStart, location.Reference) {
				listReferenceStart = append(listReferenceStart, location)
			} else if utils.StringContains(referenceEnd, location.Reference) {
				listReferenceEnd = append(listReferenceEnd, location)
			}
			for _, item := range referenceItems {
				if item.Reference == location.Reference {
					item.Detail = location
					listReferenceItems = append(listReferenceItems, item)
				}
			}

			for _, pickup := range referenceTravelerPickup {
				if pickup.Reference == location.Reference {
					pickup.Detail = location
					listReferenceTravelerPickup = append(listReferenceTravelerPickup, pickup)
				}
			}

			for _, day := range referenceDays {
				if day.Reference == location.Reference {
					day.Detail = location
					listReferenceDays = append(listReferenceDays, day)
				}
			}
		}
		locationDetailResult = map[string]interface{}{
			"start":  listReferenceStart,
			"end":    listReferenceEnd,
			"pickup": listReferenceTravelerPickup,
			"day":    listReferenceDays,
			"items":  listReferenceItems,
		}
	}
	return locationDetailResult
}

func (h *CronJobHandler) ListCronJobInfo(r *ginext.Request) *ginext.Response {
	return ginext.NewResponseData(http.StatusOK, h.CronJobInfo)
}

func (h *CronJobHandler) ListCronScheduleInfo(r *ginext.Request) *ginext.Response {
	return ginext.NewResponseData(http.StatusOK, h.CronSchedule)
}

func (h *CronJobHandler) ClearAllCronFunction() {
	for key, element := range h.CronJob {
		if key != "ADMIN" {
			element.Quit <- true
			delete(h.CronJob, key)
			delete(h.CronJobInfo, key)
		}
	}
}

func (h *CronJobHandler) ClearAllScheduleFunction() {
	for key, element := range h.CronSchedule {
		if key != "ADMIN" {
			element.Stop()
			delete(h.CronSchedule, key)
			delete(h.CronScheduleInfo, key)
		}
	}
}
