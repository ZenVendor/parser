package parser

import (
	"fmt"
	"slices"
	"strings"
)

type Parser struct {
	Verb   int
	Args   []int
	Kwargs map[int]interface{}
}

func NewParser(defaultVerb int, defaultArgs []int, defaultKwargs map[int]interface{}) (Parser, error) {
	var p = Parser{
		Verb:   0,
		Args:   make([]int, 0, 2),
		Kwargs: map[int]interface{}{},
	}
	for _, arg := range defaultArgs {
		if !slices.Contains(verbValidArgMap()[defaultVerb], arg) {
			err := fmt.Errorf("Command %d does not support argument %d", defaultVerb, arg)
			return p, err
		}
	}
	for key := range defaultKwargs {
		if !slices.Contains(verbValidArgMap()[defaultVerb], key) {
			err := fmt.Errorf("Command %d does not support argument %d", defaultVerb, key)
			return p, err
		}
	}
	p.Verb = defaultVerb
	p.Args = defaultArgs
	p.Kwargs = defaultKwargs
	return p, nil
}

func (p *Parser) Print() {
	fmt.Printf("Command: %d\n", p.Verb)
	fmt.Printf("Switches:")
	for _, s := range p.Args {
		fmt.Printf(" %d", s)
	}
	fmt.Printf("\nParameters:\n")
	for key, value := range p.Kwargs {
		fmt.Printf("%d: %s\n", key, value)
	}
}

func (p *Parser) ToString() string {
	str := fmt.Sprintf("%d -", p.Verb)
	for _, s := range p.Args {
		str = fmt.Sprintf("%s %d", str, s)
	}
	for key, value := range p.Kwargs {
		str = fmt.Sprintf("%s %d:%s", str, key, value)
	}
	return str
}

func (p *Parser) Parse(args []string) error {
	if len(args) == 0 {
		if p.Verb == INVALID_ARG {
			err := fmt.Errorf("No arguments provided")
			return err
		}
		return nil
	}
	p.Args = make([]int, 0, 2)
	p.Kwargs = map[int]interface{}{}

	p.Verb = verbMap()[args[0]]
	if p.Verb == INVALID_ARG {
		err := fmt.Errorf("Invalid command \"%s\"", args[0])
		return err
	}

	argsStart := 1
	if verbValKey := verbValueMap()[p.Verb]; verbValKey != INVALID_ARG {
		if len(args) < 2 {
			err := fmt.Errorf("Command \"%s\" requires a value", args[0])
			return err
		}
		verbVal, err := kwargValidatorMap()[verbValKey](args[1])
		if err != nil {
			return err
		}
		p.Kwargs[verbValKey] = verbVal
		argsStart = 2
	}

	for _, arg := range args[argsStart:] {
		kwarg := strings.SplitN(arg, "=", 2)
		if len(kwarg) == 1 {
			sw := argMap()[kwarg[0]]
			if sw == INVALID_ARG {
				err := fmt.Errorf("Invalid argument %s", arg)
				return err
			}
			if !slices.Contains(verbValidArgMap()[p.Verb], sw) {
				err := fmt.Errorf("Command %s does not support argument %s", args[0], arg)
				return err
			}
			p.Args = append(p.Args, sw)
		} else {
			key := kwargMap()[kwarg[0]]
			if key == INVALID_ARG {
				err := fmt.Errorf("Invalid argument %s", arg)
				return err
			}
			value := kwarg[1]
			if !slices.Contains(verbValidArgMap()[p.Verb], key) {
				err := fmt.Errorf("Command %s does not support argument %s", args[0], arg)
				return err
			}
			val, err := kwargValidatorMap()[key](value)
			if err != nil {
				return err
			}
			p.Kwargs[key] = val
		}
	}

	return nil
}
