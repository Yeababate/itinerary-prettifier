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

## CSV Format
The airport lookup file must follow this structure:
name,iso_country,municipality,icao_code,iata_code,coordinates

Example:
Hongyuan Airport,CN,Aba,ZUHY,AHJ,"102.35224, 32.53154"

---

## Usage

### Run the program: 
go run . ./input.txt ./output.txt ./airport-lookup.csv

### Help: 
go run . -h/-H

---

## Input Format Examples

### Airport Codes
#HEL → IATA code → Hongyuan Airport

##EFHK → ICAO code → Hongyuan Airport

*#HEL → Aba

### Date 
D(2026-04-27) → 27 Apr 2026

### Time 
- D(2022-05-09T08:07Z) → 09 May 2022
- T12(2069-04-24T19:18-02:00) → 07:18PM (-02:00)
- T12(2080-05-04T14:54Z) → 02:54PM (+00:00)
- T24(2032-07-17T04:08+13:00) → 04:08 (+13:00)

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
