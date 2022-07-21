package parser

import "bufio"
import "os"
import "strings"
import "errors"

// struct to hold the map of map for section, key and value
type Parser struct {
	ini map[string]map[string]string
}

// returns pointer to the struct
func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) GetDataFromString(data string) (e error) {
	e = p.parse(data)
	return e
}

func (p *Parser) GetDataFromFile(fileName string) (e error) {
	file, e := os.ReadFile(fileName)
	if e != nil {
		return e
	}
	data := string(file)
	e = p.parse(data)
	return e
}

func (p *Parser) SaveToFile(fileName string, data map[string]map[string]string) error {

	f, fe := os.Create(fileName)
	defer f.Close()
	
	if fe != nil {
		return errors.New("Could not open file")
	}

	for k := range data {
		_, e := f.WriteString(k + "\n")
		for key, value := range data[k] {
			f.WriteString(key + "=" + value + "\n")
		}
		f.WriteString("\n")
		if e != nil {
			return errors.New("Encountered problem while writing to file")
		}
	}
	return nil
}

// look for open and close square brackets around section name
func checkSectionName(line string) bool {
	if (len(line) > 0) && (line[0] == '[') && (strings.Count(line, "]") == 1) &&
		(strings.Count(line, "[") == 1) && (line[len(line)-1] == ']') {
		return true
	}
	return false
	
}

// assign value to specific sectiona and key 
func (p *Parser) setValues(section, key, value string) error{
	_, ok := p.ini[section]
	
	if !ok {
		return errors.New("Cannot add value, section or key not found")
	}
	p.ini[section][key] = value
	return nil
	
}

// retrieve array of strings with all sections 
func (p *Parser) GetSectionNames() []string {

	keys := make([]string, 0, len(p.ini))
	for k := range p.ini {
		keys = append(keys, k)
	}
	return keys
}

// get value of key in a section
func (p *Parser) GetValue(section string, key string) (string, error) {
	_, ok := p.ini[section]
	if !ok {
		return "", errors.New("Section does not exist")
	}
	_, ok = p.ini[section][key]
	if !ok {
		return "", errors.New("Key does not exist")
	}

	return p.ini[section][key], nil
}

// get contents of a section
func (p *Parser) GetSection(section string) (map[string]string, error) {
	data, ok := p.ini[section]

	if !ok {
		return nil , errors.New("Section does not exist")
	}
	return data, nil
}

func (p *Parser) setSection(section string) error {
	_, ok := p.ini[section]
	if !ok {
		
		if p.ini[section] == nil {
			p.ini[section] = make(map[string]string)
		}
		
		return nil
	}
	return errors.New("Cannot create section")
}

func (p *Parser) parse(content string) error {
	
	scanner := bufio.NewScanner(strings.NewReader(content))

	p.ini = make(map[string]map[string]string)

	var key string
	var value string
	var section string
	
	sectionFound := false

	for scanner.Scan() {
		line := scanner.Text()

		items := strings.Split(line, " ")

		// check if line is a comment
		if items[0] == ";" {
			continue
		} else if checkSectionName(items[0]) == true { // look for section 
			x := string(items[0])
			e := p.setSection(x[1:len(x)-1])
			section = x[1:len(x)-1] // remove square brackets
			sectionFound = true 

			if e != nil {
				return e
			}
		} else if len(items) == 1 {
			if items[0] == "" { // empty line
				continue
			} else { // invalid type of line
				return errors.New("Syntax incorrect")
			}
		} else if sectionFound == true { // key and values can be retrieved now
			keyValuePair := strings.Split(line, "=")

			if len(keyValuePair) == 2 {

				key = keyValuePair[0]
				value = keyValuePair[1]
				p.setValues(section, key, value)
			} else { // no equal sign
				return errors.New("Key value pair incorrect")
			}
		} else { 
			return errors.New("Failed to parse the line")
		}

	}
	p.SaveToFile("output.txt", p.ini)
	return nil
}


