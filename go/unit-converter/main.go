package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// Length conversion
func convertLength(value float64, fromUnit, toUnit string) float64 {
	units := map[string]float64{
		"mm": 0.001, "cm": 0.01, "m": 1, "km": 1000,
		"inch": 0.0254, "ft": 0.3048, "yard": 0.9144, "mile": 1609.34,
	}
	return value * units[fromUnit] / units[toUnit]
}

// Weight conversion
func convertWeight(value float64, fromUnit, toUnit string) float64 {
	units := map[string]float64{
		"mg": 0.000001, "g": 0.001, "kg": 1, "oz": 0.0283495, "lb": 0.453592,
	}
	return value * units[fromUnit] / units[toUnit]
}

// Temperature conversion
func convertTemperature(value float64, fromUnit, toUnit string) float64 {
	switch fromUnit + "_" + toUnit {
	case "C_F":
		return value*9/5 + 32
	case "C_K":
		return value + 273.15
	case "F_C":
		return (value - 32) * 5 / 9
	case "F_K":
		return (value-32)*5/9 + 273.15
	case "K_C":
		return value - 273.15
	case "K_F":
		return (value-273.15)*9/5 + 32
	default:
		return value
	}
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		value, _ := strconv.ParseFloat(r.FormValue("value"), 64)
		fromUnit := r.FormValue("from_unit")
		toUnit := r.FormValue("to_unit")
		unitType := r.FormValue("unit_type")
		var result float64

		switch unitType {
		case "length":
			result = convertLength(value, fromUnit, toUnit)
		case "weight":
			result = convertWeight(value, fromUnit, toUnit)
		case "temperature":
			result = convertTemperature(value, fromUnit, toUnit)
		}

		templates.ExecuteTemplate(w, unitType+".html", map[string]interface{}{
			"Result":   result,
			"FromUnit": fromUnit,
			"ToUnit":   toUnit,
			"Value":    value,
		})
	} else {
		unitType := r.URL.Query().Get("unit")
		if unitType == "" {
			unitType = "length"
		}
		templates.ExecuteTemplate(w, unitType+".html", nil)
	}
}

func main() {
	http.HandleFunc("/", convertHandler)
	http.ListenAndServe(":8080", nil)
}
