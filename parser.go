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

func (p *Parser) LoadFromString(data string) (e error) {
	e = p.parse(data)
	return e
}

func (p *Parser) LoadFromFile(fileName string) (e error) {
	f, e := os.ReadFile(fileName)
	if e != nil {
		return e
	}
	data := string(f)
	e = p.parse(data)
	return e
}

func (p *Parser) String(data map[string]map[string]string) (string, error) { 
	var output string
	for k := range data {
		output := "[" + k + "]\n"
		for key, value := range data[k] {
			output += key 
			output += "=" 
			output += value 
			output += "\n"
		}
		output += "\n"
	}
	if output == "" {
		return output, errors.New("failed to convert map to string")
	}
	return output, nil
}

func (p *Parser) SaveToFile(fileName string) error {

	f, fe := os.Create(fileName)
	defer f.Close()
	
	if fe != nil {
		return errors.New("could not open file")
	}

	text, e := p.String(p.ini)
	if e != nil {
		return e
	}
	f.WriteString(text)
	return nil
}

// look for open and close square brackets around section name 
func isValidSectionName(line string) bool {
	if (len(line) > 0) && (line[0] == '[') && (strings.Count(line, "]") == 1) &&
		(strings.Count(line, "[") == 1) && (line[len(line)-1] == ']') {
		return true
	}
	return false
	
}

// assign value to specific section and key 
func (p *Parser) Set(section, key, value string) error{
	_, ok := p.ini[section]
	
	if !ok { 
		e := p.setSection(section)
		if e != nil {
			return e
		}
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
func (p *Parser) Get(section string, key string) (string, error) {
	_, ok := p.ini[section]
	if !ok {
		return "", errors.New("section does not exist") 
	}
	_, ok = p.ini[section][key]
	if !ok {
		return "", errors.New("key does not exist")
	}

	return p.ini[section][key], nil
}

// get contents of a section
func (p *Parser) GetSection(section string) (map[string]string, error) {
	data, ok := p.ini[section]

	if !ok {
		return nil , errors.New("section does not exist")
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
	return errors.New("cannot create section")
}

func (p *Parser) parse(content string) error {
	
	scanner := bufio.NewScanner(strings.NewReader(content))

	p.ini = make(map[string]map[string]string)

	key := "" 
	value := "" 
	section := "" 
	
	sectionFound := false

	for scanner.Scan() {
		line := scanner.Text()

		items := strings.Split(line, " ")

		// check if line is a comment
		if items[0] == ";" {
			continue
		} else if isValidSectionName(items[0]) == true { // look for section 
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
				return errors.New("syntax incorrect")
			}
		} else if sectionFound == true { // key and values can be retrieved now
			if (strings.Count(line, "=") > 1) || (strings.Count(line, "=") < 1) {
				return errors.New("number of equal signs not equal to 1")
			}
			keyValuePair := strings.Split(line, "=")
			if len(keyValuePair) == 2 {

				key = keyValuePair[0]
				value = keyValuePair[1]
				p.Set(section, key, value)
			} else { // no equal sign
				return errors.New("key value pair incorrect")
			}
		} else { 
			return errors.New("failed to parse the line")
		}

	}
	return nil
}


