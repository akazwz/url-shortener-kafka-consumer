package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/akazwz/url-shortener-kafka/global"
	"github.com/akazwz/url-shortener-kafka/initialize"
	"github.com/akazwz/url-shortener-kafka/model"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func consumeVisitsLog() {
	mechanism, err := scram.Mechanism(scram.SHA256, os.Getenv("KAFKA_USERNAME"), os.Getenv("KAFKA_PASSWORD"))
	if err != nil {
		log.Fatalln("mechain error:", err)
	}
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	r := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers: []string{os.Getenv("KAFKA_ADDR")},
			Topic:   "visits-log",
			Dialer:  dialer,
			GroupID: "group_3",
		},
	)
	for {
		message, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln("fetch message error:", err)
			return
		}
		writeToDB(message)
	}
}

func writeToDB(message kafka.Message) {
	visitsLog := model.Visits{}
	err := json.Unmarshal(message.Value, &visitsLog)
	if err != nil {
		log.Println("json unmarshal error:", err)
	}
	visits := model.VisitsLog{
		UUID:            uuid.NewV4().String(),
		Short:           visitsLog.Short,
		Url:             visitsLog.Url,
		Ip:              visitsLog.Ip,
		Region:          visitsLog.Region,
		Country:         visitsLog.Country,
		City:            visitsLog.City,
		Longitude:       visitsLog.Longitude,
		Latitude:        visitsLog.Latitude,
		UA:              visitsLog.UA.UA,
		BrowserName:     visitsLog.UA.Browser.Name,
		BrowserVersion:  visitsLog.UA.Browser.Version,
		BrowserMajor:    visitsLog.UA.Browser.Major,
		EngineName:      visitsLog.UA.Engine.Name,
		EngineVersion:   visitsLog.UA.Engine.Version,
		OSName:          visitsLog.UA.OS.Name,
		OSVersion:       visitsLog.UA.OS.Version,
		DeviceModel:     visitsLog.UA.Device.Model,
		DeviceType:      visitsLog.UA.Device.Type,
		DeviceVendor:    visitsLog.UA.Device.Vendor,
		CPUArchitecture: visitsLog.UA.CPU.Architecture,
		Time:            message.Time.UnixMilli(),
	}
	err = global.GDB.Create(&visits).Error
	if err != nil {
		log.Println("新增访问记录失败:", err)
	}
}

func main() {
	consumeVisitsLog()
}

func init() {
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatalln("读取配置文件失败")
		}
	}
	global.GDB = initialize.InitGorm()
	initialize.RegisterTables(global.GDB)
}
