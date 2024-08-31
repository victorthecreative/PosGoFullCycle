package tax

import "testing"

//func TestCalculateTax(t *testing.T) {
//	amount := 100.0
//	expected := 6.0
//
//	result := CaculateTax(amount)
//	if result != expected {
//		t.Errorf("CalculateTax returned %f, expected %f", result, expected)
//	}
//}

func TestCalulateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range table {
		result := CaculateTax(item.amount)
		if result != item.expected {
			t.Errorf("CalculateTax returned %f, expected %f", result, item.expected)
		}
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1,-4,263.0,4568.0,534.0,50.0,23.0,4.0}

	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CaculateTax(amount)
		if amount <= 0 && != result !=0 {
			t.Errorf("CalculateTax returned %f, expected %f", result, amount)
		}

	})
}
