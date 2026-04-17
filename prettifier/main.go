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

func ReadInput(FileName string) string {
	file, err := os.Open(FileName)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	reIATA := regexp.MustCompile(" #[A-Z]{3}")
	reICAO := regexp.MustCompile("##[A-Z]{4}")
	var alldata []Data = GetName("airport-lookup.csv")
	var line string 
	for scanner.Scan(){
		line = scanner.Text()		
		var SingleRowIATA [] string = reIATA.FindAllString(line,-1)
		var SingleRowICAO [] string = reICAO.FindAllString(line,-1)

		//var str string = ""
		var i, j int

		for i = 0; i < len(SingleRowIATA); i++ {
			for k := 0; k < len(alldata); k++ {
				if(SingleRowIATA[i][2:] == alldata[k].iata_code){
					line = strings.Replace(line, SingleRowIATA[i][1:], alldata[k].name, 1)
				}
			}
		}

		for j = 0; j < len(SingleRowICAO); j++ {
			for l := 0; l < len(alldata); l++{
				if (SingleRowICAO[j][2:] == alldata[l].icao_code){
					line = strings.Replace(line, SingleRowICAO[j], alldata[l].name, 1)
				}
			}
		}
		
	}
	return line
}

func ReadDate (FileName string) string {
	reDate := regexp.MustCompile("D\\((.*?)T")
		var DateRow[] string = reDate.FindAllString(FileName,-1)

		for m := 0; m < len(DateRow) ; m++ {
			timestr := strings.TrimSpace(DateRow[m][2:12])
			layout := "2006-01-02"
			t, err := time.Parse(layout, timestr)
			if err != nil {
				fmt.Println(err)
			}
			FormatedDate := t.Format("02 Jan 2006")
			FileName = strings.Replace(FileName,DateRow[m][m:],FormatedDate,1)
		}
		fmt.Println(FileName)
		return FileName

	}
	
func ReadTime (FileName string) string {
	reTime := regexp.MustCompile("T(.*?)\\)")
		var TimeRow[] string = reTime.FindAllString(FileName,-1)
		
		for m := 0; m < len(TimeRow) ; m++ {
			timestr := strings.TrimSpace(TimeRow[m][1:11])
			layout := "15:04"
			t, err := time.Parse(layout, timestr)
			if err != nil {
				fmt.Println(err)
			}
			FormatedDate := t.Format("02 Jan 2006")
			FileName = strings.Replace(FileName,TimeRow[m][m:11],FormatedDate,1)
		}
		fmt.Println(FileName)
		return FileName
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
	OutPutMessage := ReadInput("input.txt")
	OutPutMessage = ReadDate(OutPutMessage)
	OutPutMessage = ReadTime(OutPutMessage)

	OutPut, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("error creating file")
		os.Exit(1)
	}
	defer file.Close()
	_, err = OutPut.WriteString(OutPutMessage)
	if err != nil {
		fmt.Println("error writing file")
		os.Exit(1)
	}
	

}
