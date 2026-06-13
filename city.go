package main

type City struct {
	Name    string
	Region  string
	Lat     float64
	Long    float64
	Unit    string
	Station string
	Emcwf   map[string]float64
	Hrrr    map[string]float64
	Metar   map[string]float64
}

var (
	nyc         = &City{Name: "New York City", Region: "us", Lat: 40.7772, Long: -73.8726, Unit: "F", Station: "KLGA"}
	chicago     = &City{Name: "Chicago", Region: "us", Lat: 41.9742, Long: -87.9073, Unit: "F", Station: "KORD"}
	miami       = &City{Name: "Miami", Region: "us", Lat: 25.7959, Long: -80.2870, Unit: "F", Station: "KMIA"}
	dallas      = &City{Name: "Dallas", Region: "us", Lat: 32.8471, Long: -96.8518, Unit: "F", Station: "KDAL"}
	seattle     = &City{Name: "Seattle", Region: "us", Lat: 47.4502, Long: -122.3088, Unit: "F", Station: "KSEA"}
	atlanta     = &City{Name: "Atlanta", Region: "us", Lat: 33.6407, Long: -84.4277, Unit: "F", Station: "KATL"}
	london      = &City{Name: "London", Region: "eu", Lat: 51.5048, Long: 0.0495, Unit: "C", Station: "EGLC"}
	paris       = &City{Name: "Paris", Region: "eu", Lat: 48.9962, Long: 2.5979, Unit: "C", Station: "LFPG"}
	munich      = &City{Name: "Munich", Region: "eu", Lat: 48.3537, Long: 11.7750, Unit: "C", Station: "EDDM"}
	ankara      = &City{Name: "Ankara", Region: "eu", Lat: 40.1281, Long: 32.9951, Unit: "C", Station: "LTAC"}
	seoul       = &City{Name: "Seoul", Region: "asia", Lat: 37.4691, Long: 126.4505, Unit: "C", Station: "RKSI"}
	tokyo       = &City{Name: "Tokyo", Region: "asia", Lat: 35.7647, Long: 140.3864, Unit: "C", Station: "RJTT"}
	shanghai    = &City{Name: "Shanghai", Region: "asia", Lat: 31.1443, Long: 121.8083, Unit: "C", Station: "ZSPD"}
	singapore   = &City{Name: "Singapore", Region: "asia", Lat: 1.3502, Long: 103.9940, Unit: "C", Station: "WSSS"}
	lucknow     = &City{Name: "Lucknow", Region: "asia", Lat: 26.7606, Long: 80.8893, Unit: "C", Station: "VILK"}
	telAviv     = &City{Name: "Tel Aviv", Region: "asia", Lat: 32.0114, Long: 34.8867, Unit: "C", Station: "LLBG"}
	toronto     = &City{Name: "Toronto", Region: "ca", Lat: 43.6772, Long: -79.6306, Unit: "C", Station: "CYYZ"}
	saoPaulo    = &City{Name: "SÃ£o Paulo", Region: "sa", Lat: -23.4356, Long: -46.4731, Unit: "C", Station: "SBGR"}
	buenosAires = &City{Name: "Buenos Aires", Region: "sa", Lat: -34.8222, Long: -58.5358, Unit: "C", Station: "SAEZ"}
	wellington  = &City{Name: "Wellington", Region: "oc", Lat: -41.3272, Long: 174.8052, Unit: "C", Station: "NZWN"}
)
