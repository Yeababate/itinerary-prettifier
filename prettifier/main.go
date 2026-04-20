package main
import ("fmt" 
"bufio" 
"os"
"strings"
"regexp"
"time"
)

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
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	defer file.Close()

	DataStorage := []Data{}
	EachData := Data{}
	EachLine := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		EachLine = scanner.Text()
		Column := strings.Split(EachLine, ",")
		EachData.name = Column[0] 
		EachData.iso_country = Column[1]
		EachData.municipality = Column[2]
		EachData.icao_code = Column[3]
		EachData.iata_code = Column[4]
		EachData.coordinates = Column[5]
		DataStorage = append(DataStorage, EachData)
	}
	return DataStorage

}

func ReadInput(inputFile string) string {
	line, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Input not found")
		os.Exit(0)
	}
	return string(line)
}

func CheckExists(input string, lookup string) bool{  // ?
	_,err := os.Stat(input)
	if os.IsNotExist(err){
		fmt.Println("Input not found.")
	}
	_, err = os.Stat(lookup)
	if os.IsNotExist(err){
		fmt.Println("Airport lookup not found.")
	}
	return true
}

func GetIATACode(FileName string, Csv string) string {
	reIATA := regexp.MustCompile("#[A-Z]{3}")
	SingleRowIATA := reIATA.FindAllString(FileName,-1)
	alldata := GetName(Csv)
	for i := 0; i < len(SingleRowIATA); i++ {
		for k := 0; k < len(alldata); k++ {
			if(SingleRowIATA[i][1:] == alldata[k].iata_code){
				FileName = strings.Replace(FileName, SingleRowIATA[i], alldata[k].name, 1)
			}
		}
	}
	return FileName
}

func GetICAOCode(FileName string, Csv string) string {
	reICAO := regexp.MustCompile("##[A-Z]{4}")
	SingleRowICAO := reICAO.FindAllString(FileName,-1)
	alldata := GetName(Csv)
	for i := 0; i < len(SingleRowICAO); i++ {
		for k := 0; k < len(alldata); k++ {
			if(SingleRowICAO[i][2:] == alldata[k].icao_code){
				FileName = strings.Replace(FileName, SingleRowICAO[i], alldata[k].name, 1)
			}
		}
	}
	return FileName
}

func ReadDate (OutPut string) string{
	reDate := regexp.MustCompile("D\\((.*?)\\)")
	Date := reDate.FindAllString(OutPut, -1)
	for i := 0; i < len(Date); i ++ {
		d, err := time.Parse("2006-01-02", Date[i][2:12])
		if err != nil {
			continue
		}
		FormaredDate := d.Format("02 Jan 2006")
		OutPut = strings.Replace(OutPut, Date[i], FormaredDate, 1)
	}
	return OutPut
}

func Read12hrTime (OutPut string) string{
	reTime12 := regexp.MustCompile("T12\\((.*?)\\)")
	Time12hr := reTime12.FindAllString(OutPut, -1)
	for i := 0; i < len(Time12hr); i++ {
		t, err := time.Parse("2006-01-02T15:04-07:00", Time12hr[i][4:26])
		if err != nil {
			continue
		}
		FormaredTime := t.Format("03:04PM (-07:00)")
		OutPut = strings.Replace(OutPut, Time12hr[i], FormaredTime, 1)
	}
	return OutPut
}

func Read24hrTime (OutPut string) string{
	reTime24 := regexp.MustCompile(`T24\((.*\d)\)`)
	Time24hr := reTime24.FindAllString(OutPut, -1)
	for i := 0; i < len(Time24hr); i++ {
		t, err := time.Parse("2006-01-02T15:04-07:00", Time24hr[i][4:26])
		if err != nil {
			continue
		}
		FormaredTime := t.Format("15:04 (-07:00)")
		OutPut = strings.Replace(OutPut,Time24hr[i], FormaredTime , 1)
	}
	return OutPut
}

func Read12ZTime (OutPut string) string{
	reTime12z := regexp.MustCompile("T12\\((.*[Zz])\\)")
	Time12Z := reTime12z.FindAllString(OutPut, -1)
	for i := 0; i < len(Time12Z); i++ {
		t, err := time.Parse("2006-01-02T15:04Z", Time12Z[i][4:21])
		if err != nil {
			continue
		}
		FormaredTime := t.Format("03:04PM (+00:00)")
		OutPut = strings.Replace(OutPut, Time12Z[i], FormaredTime , 1)
	}
	return OutPut
}

func Read24ZTime (OutPut string) string{
	reTime24z := regexp.MustCompile("T24\\((.*[Zz])\\)")
	Time24Z := reTime24z.FindAllString(OutPut, -1)
	for i := 0; i < len(Time24Z); i++ {
		t, err := time.Parse("2006-01-02T15:04Z", Time24Z[i][4:21])
		if err != nil {
			continue
		}
		FormaredTime := t.Format("15:04 (+00:00)")
		OutPut = strings.Replace(OutPut, Time24Z[i], FormaredTime , 1)
	}
	return OutPut
}

func WriteToOutput(name string, output string){
	file,_ := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	_, err := file.WriteString(strings.TrimSpace(output))
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func main(){

	CheckExists(os.Args[1],os.Args[2])
	if len(os.Args) != 3{
		fmt.Println("Input the right amount of arguments")
	}

	file, err := os.Open("airport-lookup.csv")
	if err != nil {
		fmt.Println("Error opening airport-lookup file")
		os.Exit(1)
	}
	defer file.Close()
	input := os.Args[1]
	Csv := os.Args[2]
	OutPutMessage := ReadInput(input)
	OutPutMessage = GetICAOCode(OutPutMessage, Csv)
	OutPutMessage = GetIATACode(OutPutMessage, Csv)
	OutPutMessage = ReadDate(OutPutMessage)
	OutPutMessage = Read12ZTime(OutPutMessage)
	OutPutMessage = Read24ZTime(OutPutMessage)
	OutPutMessage = Read12hrTime(OutPutMessage)
	OutPutMessage = Read24hrTime(OutPutMessage)
	WriteToOutput("output.txt", OutPutMessage)
}
