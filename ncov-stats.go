package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Kamva/mgm/v2"
	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version    string
	commitHash string
	buildDate  string
)

func SetupDB(uri, name string) error {
	err := mgm.SetDefaultConfig(nil, name, options.Client().ApplyURI(uri))
	return err
}

func fetchData() error {
	http := resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(20 * time.Second).
		SetHostURL("http://lab.isaaclin.cn/nCoV/api")

	rsp, err := http.R().
		SetResult(&Result{}).
		Get("/area?latest=1")
	if err != nil {
		return err
	}
	result := rsp.Result().(*Result)
	if result == nil || !result.Success {
		return fmt.Errorf("invalid result: %s\n", rsp.String())
	}

	date := time.Now().Add(-24 * time.Hour).Format("2006-01-02")

	var infos []NCovInfo
	if err := mgm.Coll(&NCovInfo{}).SimpleFind(&infos, bson.M{"date": date}); err != nil {
		return err
	}

	infoMap := make(map[string]*NCovInfo)
	for _, info := range infos {
		infoMap[info.Region] = &info
	}

	for _, result := range result.Results {
		if result.CountryName == "中国" && infoMap[result.ProvinceShortName] == nil {
			var cities []*CityInfo
			for _, city := range result.Cities {
				cities = append(cities, &CityInfo{
					Name:               city.CityName,
					LocID:              city.LocationId,
					Dead:               city.DeadCount,
					Confirmed:          city.ConfirmedCount,
					Suspected:          city.SuspectedCount,
					Cured:              city.CuredCount,
					RemainingConfirmed: city.CurrentConfirmedCount,
				})
			}
			info := NCovInfo{
				Region:             result.ProvinceShortName,
				LocID:              result.LocationId,
				Date:               date,
				Dead:               result.DeadCount,
				Confirmed:          result.ConfirmedCount,
				Suspected:          result.SuspectedCount,
				Cured:              result.CuredCount,
				RemainingConfirmed: result.CurrentConfirmedCount,
				Cities:             cities,
			}
			_ = mgm.Coll(&NCovInfo{}).Create(&info)
			fmt.Println("update", result.ProvinceShortName, date)
		}
	}
	fmt.Println("done.")

	return nil
}

func main() {
	v, p := viper.New(), pflag.NewFlagSet(friendlyAppName, pflag.ExitOnError)

	configure(v, p)

	p.String("config", "", "Configuration file")
	p.Bool("version", false, "Show version information")

	_ = p.Parse(os.Args[1:])

	if v, _ := p.GetBool("version"); v {
		fmt.Printf("%s version %s (%s) built on %s\n", friendlyAppName, version, commitHash, buildDate)

		os.Exit(0)
	}

	if c, _ := p.GetString("config"); c != "" {
		v.SetConfigFile(c)
	}

	err := v.ReadInConfig()
	_, configFileNotFound := err.(viper.ConfigFileNotFoundError)
	if !configFileNotFound {
		log.Panic("failed to read configuration", err)
	}

	var config configuration
	err = v.Unmarshal(&config)
	if err != nil {
		log.Panic("failed to unmarshal configuration", err)
	}

	if configFileNotFound {
		log.Println("configuration file not found")
	}

	fmt.Printf("%+v\n", config)

	if err := SetupDB(config.DB.URI, config.DB.DBName); err != nil {
		log.Fatalf("mongodb error: %v\n", err)
	}

	if err := fetchData(); err != nil {
		time.Sleep(time.Second * 30)
		log.Fatalf("fetchData error: %v\n", err)
	}
}
