package main
import ("fmt" 
// "bufio" 
"os"
"strings"
"regexp"
"time"
"encoding/csv"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Magenta = "\033[35m"
var Italic = "\033[3m"
var Bold = "\033[1m"

type Data struct {
	name string
	iso_country string
	municipality string
	icao_code string
	iata_code string
	coordinates string
}

func GetName(FileName string)[]Data{
	file, err := os.Open(FileName)
	if err != nil {
		fmt.Println(Red + "Error opening file" + Reset)
		os.Exit(0)
	}
	reader:= csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(Red + Bold + "Airport lookup malformed" + Reset)
		os.Exit(0)
	}

	DataStorage := []Data{}
	for _, Column := range records {
		if len(Column) != 6 {
    		fmt.Println(Red + Bold + "Airport lookup malformed" + Reset)
    		os.Exit(0)
		}
		EachData := Data{
			name: Column[0],
			iso_country: Column[1],
			municipality: Column[2],
			icao_code: Column[3],
			iata_code: Column[4],
			coordinates: Column[5],
		}
		if EachData.name == "" || EachData.iso_country == "" || EachData.municipality == "" || EachData.icao_code == "" || EachData.iata_code == "" || EachData.coordinates == ""{
			fmt.Println(Red + Bold + "Airport lookup malformed" + Reset)
			os.Exit(0)
		} 
		DataStorage = append(DataStorage, EachData)
	}
	return DataStorage
}

func ReadInput(inputFile string) string {
	line, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(Red + "Error Reading" + Reset)
		os.Exit(0)
	}
	return strings.TrimSpace(string(line))
}

func Malformed(Column Data) bool{
	name := strings.TrimSpace(Column.name)
	iso_country := strings.TrimSpace(Column.iso_country)
	municipality := strings.TrimSpace(Column.municipality)
	icao_code := strings.TrimSpace(Column.icao_code)
	iata_code := strings.TrimSpace(Column.iata_code)
	coordinates := strings.TrimSpace(Column.coordinates)
	var data = []string{name, iso_country, municipality, icao_code, iata_code, coordinates}

	for _, val := range data{
		if val == "" {
			return true
		}
		if strings.ContainsRune(val, '�') {
			return true
		}
	}
	return false
}

func GetIATACode(FileName string, Csv string) string {
	reIATA := regexp.MustCompile(`#[A-Z]{3}\b`)
	SingleRowIATA := reIATA.FindAllString(FileName,-1)
	alldata := GetName(Csv)
	for i := 0; i < len(SingleRowIATA); i++ {
		match := SingleRowIATA[i]
		if strings.Contains(FileName, "#" + match){
			continue
		}
		for k := 0; k < len(alldata); k++ {
			if(SingleRowIATA[i][1:] == alldata[k].iata_code){
				if Malformed(alldata[k]) == true{
					fmt.Println(Red + Bold +"Airport lookup malformed" + Reset)
					os.Exit(0)
				}
				FileName = strings.Replace(FileName, SingleRowIATA[i], alldata[k].name, 1)
			}
		}
	}
	return FileName
}

func GetICAOCode(FileName string, Csv string) string {
	reICAO := regexp.MustCompile(`##[A-Z]{4}\b`)
	SingleRowICAO := reICAO.FindAllString(FileName,-1)
	alldata := GetName(Csv)
	for i := 0; i < len(SingleRowICAO); i++ {
		for k := 0; k < len(alldata); k++ {
			if(SingleRowICAO[i][2:] == alldata[k].icao_code){
				if Malformed(alldata[k]) == true{
				fmt.Println(Red + Bold + "Airport lookup malformed" + Reset)
				os.Exit(0)
			}
				FileName = strings.Replace(FileName, SingleRowICAO[i], alldata[k].name, 1)
			}
		}
	}
	return FileName
}

func GetCityName(FileName string, Csv string) string {
	reCity := regexp.MustCompile(`\*#[A-Z]{3}\b|\*##[A-Z]{4}\b`)
	SingleRowCity := reCity.FindAllString(FileName,-1)
	alldata := GetName(Csv)
	for i := 0; i < len(SingleRowCity); i++ {
		for k := 0; k < len(alldata); k++ {
			if strings.HasPrefix(SingleRowCity[i], "*#"){
				if(SingleRowCity[i][2:] == alldata[k].iata_code) {
					FileName = strings.Replace(FileName, SingleRowCity[i], alldata[k].municipality, 1)
				}
			}
			if strings.HasPrefix(SingleRowCity[i], "*##"){
				if(SingleRowCity[i][3:] == alldata[k].icao_code) {
					FileName = strings.Replace(FileName, SingleRowCity[i], alldata[k].municipality, 1)
				}
			}
		}
	}
	return FileName
}

func ReadDate (OutPut string) string{
	reDate := regexp.MustCompile(`D\((.*?)\)`)
	Date := reDate.FindAllString(OutPut, -1)
	for i := 0; i < len(Date); i ++ {
		d, err := time.Parse("2006-01-02", Date[i][2:12])
		if err != nil {
			continue
		}
		FormatedDate := d.Format("02 Jan 2006")
		OutPut = strings.Replace(OutPut, Date[i], FormatedDate, 1)
	}
	return OutPut
}

func Read12hrTime (OutPut string) string{
	reTime12 := regexp.MustCompile(`T12\((.*?)\)`)
	Time12hr := reTime12.FindAllString(OutPut, -1)
	for i := 0; i < len(Time12hr); i++ {
		if len(Time12hr[i]) < 26 {
			continue
		}
		t, err := time.Parse("2006-01-02T15:04-07:00", Time12hr[i][4:26])
		if err != nil {
			continue
		}
		FormatedTime := t.Format("03:04PM (-07:00)")
		OutPut = strings.Replace(OutPut, Time12hr[i], FormatedTime, 1)
	}
	return OutPut
}

func Read24hrTime (OutPut string) string{
	reTime24 := regexp.MustCompile(`T24\((.*\d)\)`)
	Time24hr := reTime24.FindAllString(OutPut, -1)
	for i := 0; i < len(Time24hr); i++ {
		if len(Time24hr[i]) < 26 {
			continue
		}
		t, err := time.Parse("2006-01-02T15:04-07:00", Time24hr[i][4:26])
		if err != nil {
			continue
		}
		FormatedTime := t.Format("15:04 (-07:00)")
		OutPut = strings.Replace(OutPut,Time24hr[i], FormatedTime , 1)
	}
	return OutPut
}

func Read12ZTime (OutPut string) string{
	reTime12z := regexp.MustCompile("T12\\((.*[Zz])\\)")
	Time12Z := reTime12z.FindAllString(OutPut, -1)
	for i := 0; i < len(Time12Z); i++ {
		if len(Time12Z[i]) < 21 {
			continue
		}
		t, err := time.Parse("2006-01-02T15:04Z", Time12Z[i][4:21])
		if err != nil {
			continue
		}
		FormatedTime := t.Format("03:04PM (+00:00)")
		OutPut = strings.Replace(OutPut, Time12Z[i], FormatedTime , 1)
	}
	return OutPut
}

func Read24ZTime (OutPut string) string{
	reTime24z := regexp.MustCompile("T24\\((.*[Zz])\\)")
	Time24Z := reTime24z.FindAllString(OutPut, -1)
	for i := 0; i < len(Time24Z); i++ {
		if len(Time24Z[i]) < 21 {
			continue
		}
		t, err := time.Parse("2006-01-02T15:04Z", Time24Z[i][4:21])
		if err != nil {
			continue
		}
		FormatedTime := t.Format("15:04 (+00:00)")
		OutPut = strings.Replace(OutPut, Time24Z[i], FormatedTime , 1)
	}
	return OutPut
}

func VerticalSpaces (output string) string {
	reVertical := regexp.MustCompile(`\r\n|\r|\v|\f`)
	output = reVertical.ReplaceAllString(output, "\n")

	reLimit := regexp.MustCompile(`\n{3,}`)
	return reLimit.ReplaceAllString(output, "\n\n")
}

func WriteToOutput(name string, output string){
	file,_ := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	_, err := file.WriteString(strings.TrimSpace(output))
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func CheckExists(input string, csv string) bool{
	_,err := os.Stat(input)
	if os.IsNotExist(err){
		fmt.Println(Red + Bold + "Input not found." + Reset)
		return false
	}
	_, err = os.Stat(csv)
	if os.IsNotExist(err){
		fmt.Println(Red + Bold + "Airport lookup not found." + Reset)
		return false
	}
	return true
}

func main(){
	if len(os.Args) != 4 || os.Args[1] == "-h" || os.Args[1] == "-H" {
		fmt.Println(Green + "itinerary usage:\ngo run . ./input.txt ./output.txt ./airport-lookup.csv" + Reset)
		os.Exit(0)
	}

	if CheckExists(os.Args[1],os.Args[3]) == false {
		os.Exit(0)
	}

	input := os.Args[1]
	Csv := os.Args[3]
	output := os.Args[2]
	OutPutMessage := ReadInput(input)
	OutPutMessage = VerticalSpaces(OutPutMessage)
	OutPutMessage = GetCityName(OutPutMessage, Csv)
	OutPutMessage = GetICAOCode(OutPutMessage, Csv)
	OutPutMessage = GetIATACode(OutPutMessage, Csv)
	OutPutMessage = ReadDate(OutPutMessage)
	OutPutMessage = Read12ZTime(OutPutMessage)
	OutPutMessage = Read24ZTime(OutPutMessage)
	OutPutMessage = Read12hrTime(OutPutMessage)
	OutPutMessage = Read24hrTime(OutPutMessage)
	WriteToOutput(output, OutPutMessage)
	fmt.Println(Magenta + Italic + OutPutMessage + Reset)
}
