package util

import (
	"bmc-test-golang-service/model"
	"bufio"
	"encoding/csv"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func InsertBatchIntoDB(db *gorm.DB, batch [][]string) {
	tx := db.Begin()
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	passengers := make([]model.PassengerInfo, 0, len(batch))

	for _, passengerData := range batch {
		passenger := model.PassengerInfo{
			PassengerID: passengerData[0],
			Survived:    ParseInt(passengerData[1]),
			Pclass:      ParseInt(passengerData[2]),
			Name:        passengerData[3],
			Sex:         passengerData[4],
			Age:         ParseFloat32(passengerData[5]),
			SibSb:       ParseInt(passengerData[6]),
			Parch:       ParseInt(passengerData[7]),
			Ticket:      passengerData[8],
			Fare:        ParseFloat64(passengerData[9]),
			Cabin:       passengerData[10],
			Embarked:    passengerData[11],
		}
		passengers = append(passengers, passenger)
	}

	if err := tx.CreateInBatches(passengers, len(passengers)).Error; err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
}

func ParseInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return value
}

func ParseFloat64(s string) float64 {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return float64(value)
}

func ParseFloat32(s string) float32 {
	value, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0.0
	}
	return float32(value)
}

func ValidateAttributes(attributes []string) bool {
	fieldMap := make(map[string]struct{})
	passengerInfoType := reflect.TypeOf(model.PassengerInfo{})
	for i := 0; i < passengerInfoType.NumField(); i++ {
		field := passengerInfoType.Field(i)
		fieldMap[strings.ToLower(field.Name)] = struct{}{}
	}

	for _, attr := range attributes {
		if _, exists := fieldMap[attr]; !exists {
			return false
		}
	}

	return true
}

func ImportCsvFileIntoDatabase(db *gorm.DB) {
	file, err := os.Open("titanic.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	batchSize := 20
	batch := make([][]string, 0, batchSize)

	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		passengerData := []string{
			record[0], record[1], record[2], record[3], record[4],
			record[5], record[6], record[7], record[8], record[9],
			record[10], record[11],
		}

		batch = append(batch, passengerData)

		if len(batch) == batchSize {
			InsertBatchIntoDB(db, batch)
			batch = batch[:0] // Clear the batch.
		}
	}

	if len(batch) > 0 {
		InsertBatchIntoDB(db, batch)
	}
	log.Println("CSV data imported into the database.")
}

func CreatePassengerInfoUsingPassedAttributes(passenger model.PassengerInfo, attributeNames []string) model.PassengerInfoDTO {
	passengerDTO := model.PassengerInfoDTO{
		PassengerID: passenger.PassengerID,
	}

	for _, attributeName := range attributeNames {
		switch attributeName {
		case "survived":
			passengerDTO.Survived = passenger.Survived
		case "pclass":
			passengerDTO.Pclass = passenger.Pclass
		case "name":
			passengerDTO.Name = passenger.Name
		case "sex":
			passengerDTO.Sex = passenger.Sex
		case "age":
			passengerDTO.Age = passenger.Age
		case "sibsb":
			passengerDTO.SibSb = passenger.SibSb
		case "parch":
			passengerDTO.Parch = passenger.Parch
		case "ticket":
			passengerDTO.Ticket = passenger.Ticket
		case "fare":
			passengerDTO.Fare = passenger.Fare
		case "cabin":
			passengerDTO.Cabin = passenger.Cabin
		case "embarked":
			passengerDTO.Embarked = passenger.Embarked
		}
	}

	return passengerDTO
}

func CreateHistogramHtmlTemplate() string {
	return `
<!DOCTYPE html>
<html>
<head>
    <title>Fare Price Histogram</title>
    <!-- Include Chart.js library -->
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.7.0/chart.min.js"></script>
</head>
<body>
    <h1>Fare Price Histogram</h1>
    <div style="width: 80%; margin: 0 auto;">
        <!-- Create a canvas element for the chart -->
        <canvas id="histogramChart"></canvas>
    </div>

    <script>
        // Data for the chart
        var labels = {{.Labels}};
        var data = {{.Data}};

        // Get the canvas element
        var ctx = document.getElementById('histogramChart').getContext('2d');

        // Create a new histogram chart
        var chart = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: labels,
                datasets: [{
                    label: 'Count',
                    data: data,
                    backgroundColor: 'rgba(75, 192, 192, 0.2)', // Adjust color as needed
                    borderColor: 'rgba(75, 192, 192, 1)', // Adjust color as needed
                    borderWidth: 1
                }]
            },
            options: {
                scales: {
                    y: {
                        beginAtZero: true,
                        title: {
                            display: true,
                            text: 'Count'
                        }
                    }
                },
                responsive: true,
                maintainAspectRatio: false
            }
        });
    </script>
</body>
</html>
`
}
