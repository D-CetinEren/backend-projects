# Unit Converter

A simple web application that allows users to convert between different units of measurement, such as length, weight, and temperature. The backend is built using Go, and the frontend is a minimal HTML interface without any CSS.

## Features
- Convert between various units of **length**, **weight**, and **temperature**.
- Simple and clean user interface.
- Backend logic written in Go with no external dependencies.
- No database required.

## Supported Units

### Length
- Millimeter (mm)
- Centimeter (cm)
- Meter (m)
- Kilometer (km)
- Inch (inch)
- Foot (ft)
- Yard (yard)
- Mile (mile)

### Weight
- Milligram (mg)
- Gram (g)
- Kilogram (kg)
- Ounce (oz)
- Pound (lb)

### Temperature
- Celsius (C)
- Fahrenheit (F)
- Kelvin (K)

## Requirements
- Go (1.23 or higher)

## Installation
1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd unit-converter
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. Open your browser and navigate to:
   ```
   http://localhost:8080/
   ```

## File Structure
```
unit-converter/
├── main.go          # Backend server in Go
├── templates/
│   ├── length.html  # Length conversion page
│   ├── weight.html  # Weight conversion page
│   └── temperature.html  # Temperature conversion page
```

## Usage
1. Select the type of conversion you want to perform (Length, Weight, or Temperature).
2. Input the value, the unit to convert from, and the unit to convert to.
3. Click **Convert** to see the result.
4. Use the navigation links to switch between different conversion types.

## Example
- **Length Conversion:** Convert `20 ft` to `cm`
  - Input: `20`
  - From Unit: `ft`
  - To Unit: `cm`
  - Result: `20 ft = 609.6 cm`

## License
This project is licensed under the MIT License. Feel free to use and modify it as needed.
