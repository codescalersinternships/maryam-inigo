INI Parser 

Go package created to manipulate a .ini file.


FEATURES 

- load input from string or file
- save parsed data to a new external file
- ensure correct syntax of INI
- set values of the pair (key,value) in a specific section
- get all section names
- find the value of a key in a section
- get specific section by name
- set section by name
- parse INI file


FUNCTIONS

- getDataFromString() to parse data from string
- getDataFromFile() to parse data from external file
- saveToFile() to output parsed data to external file
- checkSectionName() to ensure section name is between 2 square brackets
- setValues() of section, key and value in the ini map of the parser
- getSectionNames() to retrieve all sections in parsed ini
- getValue() to retrieve a value of a specific key in a section
- getSection() retrieves the section given its name
- setSection() to create a new section by name
- parse() to parse INI file into a map given a string input


