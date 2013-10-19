package gova

func init() {
	Validators.Add("pattern", new(PatternValidator))
	Validators.Add("email", new(EmailValidator))
	Validators.Add("length", new(LengthValidator))
}
