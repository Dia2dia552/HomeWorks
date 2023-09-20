package SliceFloat

import "fmt"

func CountAverageRate(rates []float64) float64 {
	sum := 0.0
	for _, rate := range rates {
		sum += rate
	}
	average := sum / float64(len(rates))
	return average

}
func StartCountRates() {
	rates := []float64{8.0, 10.0, 5.5, 7.3, 8.5}
	averageRate := CountAverageRate(rates)
	fmt.Printf("Середня оцінка з предмету : %2f\n", averageRate)
}
