package main

import "bufio"
import "os"
import "fmt"
import "strings"
import "errors"


type Parser struct {
	ini map[string]map[string]string
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) getDataFromString(data string) (e error) {
	e = p.parse(data)
	return e
}

func (p *Parser) getDataFromFile(fileName string) (e error) {
	file, e := os.ReadFile(fileName)
	if e != nil {
		return e
	}
	data := string(file)
	e = p.parse(data)
	return e
}


func (p *Parser) saveToFile(fileName string, parser map[string]map[string]string) (err error) {

	f, fe := os.Create(fileName)
	defer f.Close()
	if fe != nil {
		return errors.New("Could not open file")
	}
	for k := range parser {
		_, e := f.WriteString(k + "\n")
		for key, value := range parser[k] {
			f.WriteString(key + " = " + value + "\n")
		}
		f.WriteString("\n")
		if e != nil {
			return errors.New("Encountered problem while writing to file")
		}

	}
	return nil
}

func checkSectionName(line string) bool {
	if (len(line) > 0) && (line[0] == '[') && (strings.Count(line, "]") == 1) &&
		(strings.Count(line, "[") == 1) && (line[len(line)-1] == ']') {
		return true
	}
	return false
	
}

func (p *Parser) setValues(section, key, value string) error{
	_, ok := p.ini[section]
	
	if !ok {
		return errors.New("Cannot add value, section or key not found")
	}
	p.ini[section][key] = value
	return nil
	
}

func (p *Parser) getSectionNames() []string {

	keys := make([]string, 0, len(p.ini))
	for k := range p.ini {
		keys = append(keys, k)
	}
	return keys
}

func (p *Parser) getValue(section string, key string) (string, error) {
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

func (p *Parser) getSection(section string) (map[string]string, error) {
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

		if items[0] == ";" {
			continue
		} else if checkSectionName(items[0]) == true {
			x := string(items[0])
			_ = p.setSection(x[1:len(x)-1])
			section = x[1:len(x)-1]
			sectionFound = true
		} else if len(items) == 1 {
			if items[0] == "" {
				continue
			} else {
				return errors.New("Syntax incorrect")
			}
		} else if sectionFound == true {
			keyValuePair := strings.Split(line, "=")

			if len(keyValuePair) == 2 {

				key = keyValuePair[0]
				value = keyValuePair[1]
				p.setValues(section, key, value)
			} else {
				return errors.New("Key value pair incorrect")
			}

		} else {
			return errors.New("Syntax incorrect")
		}

	}
	p.saveToFile("output.txt", p.ini)
	return nil
}


func main() {
	parser := Parser{}
	parser.getDataFromFile("input.ini")
	
	fmt.Println(parser.ini)
}
