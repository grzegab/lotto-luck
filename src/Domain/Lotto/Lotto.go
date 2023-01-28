package Lotto

type WithPrediction struct {
	LottoID          int
	Date             string
	Numbers          []int
	PredictionYearly []int
	PredictionWhole  []int
	OwnNumbers       []int
	HitsYear         int
	HitsWhole        int
	HitsOwn          int
}

func (w *WithPrediction) CalculateHits() {
	if w.Numbers == nil || w.PredictionWhole == nil {
		w.HitsWhole = 0
	} else {
		w.HitsWhole = checkHits(w.Numbers, w.PredictionWhole)
	}

	if w.Numbers == nil || w.PredictionYearly == nil {
		w.HitsYear = 0
	} else {
		w.HitsYear = checkHits(w.Numbers, w.PredictionYearly)
	}

	if w.OwnNumbers != nil {
		w.HitsOwn = checkHits(w.Numbers, w.OwnNumbers)
	}
}

func checkHits(a []int, b []int) int {
	count := 0
	for _, an := range a {
		for _, bn := range b {
			if an == bn {
				count++
				break
			}
		}
	}

	return count
}
