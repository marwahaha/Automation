package Automation

type PRule struct {
	from string
	with string
	to   string
	pop  string
	push string
}

func NewPRule(from, with, pop, to, push string) PRule {
	return PRule{from: from, with: with, to: to, pop: pop, push: push}
}

func NewPRules(from, with, pop, to, push string) []PRule {
	rules := make([]PRule, 0, len(with))
	for _, char := range with {
		rules = append(rules, NewPRule(from, string(char), pop, to, push))
	}
	return rules
}

func (p PRule) String() string {
	return "(" + p.from + ")-" + p.with + "-|*" + p.pop + "*|" + "(" + p.to + ")<" + p.push + ">"
}

func (p PRule) GetFrom() string {
	return p.from
}

func (p PRule) GetWith() string {
	return p.with
}

func (p PRule) GetTo() string {
	return p.to
}

func (p PRule) GetPush() string {
	return p.push
}

func (p PRule) GetPop() string {
	return p.pop
}

func (p PRule) SetFrom(from string) {
	p.from = from
}

func (p PRule) SetWith(with string) {
	if len(with) <= 1 {
		p.with = with
	}
}

func (p PRule) SetTo(to string) {
	p.to = to
}

func (p PRule) SetPush(push string) {
	p.push = push
}
func (p PRule) SetPop(pop string) {
	p.pop = pop
}
