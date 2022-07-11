package task21

type russiaAdapter struct {
	russianHuman *russia
}

func (r *russiaAdapter) insertHumanStats(){
	showStats(r.russianHuman.conertToUSA())
}
