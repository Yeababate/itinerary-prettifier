# Itinerary Formatter (Go)

## Overview
This Go program processes an input text file containing encoded itinerary information and converts it into a human-readable format. It replaces airport codes with names, formats dates and times, and cleans up spacing.

---

## Features
- ✈️ Replace **IATA codes** (`#XXX`) with airport names  
- 🛫 Replace **ICAO codes** (`##XXXX`) with airport names  
- 🏙️ Replace **city codes** (`*#XXX`) with municipality names  
- 📅 Format dates from `D(YYYY-MM-DD)` → `DD Mon YYYY`  
- ⏰ Format time:
  - `T12(...)` → 12-hour format  
  - `T24(...)` → 24-hour format  
  - Supports both timezone offsets and `Z` (UTC)  
- 🧹 Normalize vertical spacing (remove excessive blank lines)

---

## Project Structure
├── main.go
├── input.txt
├── output.txt
└── airport-lookup.csv


---

## CSV Format
The airport lookup file must follow this structure:

Example:

---

## Usage

### Run the program: 
go run . ./input.txt ./output.txt ./airport-lookup.csv

### Help: 
go run . -h

---

## Input Format Examples

### Airport Codes
#HEL → Helsinki Vantaa Airport
##EFHK → Helsinki Vantaa Airport
*#HEL → Helsinki

### Date 
D(2024-03-15) → 15 Mar 2024

### Time 
T12(2024-03-15T14:30+02:00) → 02:30PM (+02:00)
T24(2024-03-15T14:30+02:00) → 14:30 (+02:00)
T12(2024-03-15T14:30Z) → 02:30PM (+00:00)
T24(2024-03-15T14:30Z) → 14:30 (+00:00)

---

## Error Handling
- Missing input file → `"Input not found."`
- Missing CSV file → `"Airport lookup not found."`
- Malformed airport data → `"Airport lookup malformed"`

---

## Notes
- The program assumes valid CSV formatting (comma-separated, no quoted commas).
- All fields in the CSV must be non-empty.
- Time parsing expects strict ISO 8601 format.

---

## Example Workflow

1. Prepare `input.txt` with encoded itinerary.
2. Provide `airport-lookup.csv`.
3. Run the program.
4. Check formatted output in `output.txt`.

---

## License
This project is open-source and free to use.