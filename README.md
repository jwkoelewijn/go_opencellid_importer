# go OpenCellID importer
Transforms a OpenCellID CSV to a new CSV with only towers from selected countries

## Usage
1. Download CSV file from [opencellid.org](http://www.opencellid.org)
2. go run importer.go (or importer when downloaded as a binary)

This will filter the input CSV (_cell\_towers.csv_) file to only accept cell towers with the Mobile Country Codes 204, 206 and 262 (the Netherlands, Belgium and Germany) and will output them to the file '_selected\_cell\_towers.csv_'

## Advanced Usage
**--help** Display usage information

**--inputFile** Specify the file to use as input

**--outputFile** Specify to which file the output should be written

**--countryCodes** Comma separated list of MCC's (Mobile Country Code's) to filter from the input file. For a list of MCC's used, see [Wikipedia](http://en.wikipedia.org/wiki/Mobile_country_code)

## About
This little program was written in Go to easily select all mobile tower information from certain countries to avoid importing all towers from all countries.