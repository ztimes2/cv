package cv

type Config struct {
	FirstName         string
	LastName          string
	Occupation        string
	CurrentLocation   Location
	PhoneNumber       string
	Email             string
	Summary           string
	Links             Links
	EmploymentHistory []Employment
	EducationHistory  []Education
	Skills            []Skill
}

type Location struct {
	City    string
	Country string
}

type Links struct {
	PersonalWebsiteURL string
	GithubURL          string
	LinkedinURL        string
}

type Date struct {
	Month int
	Year  int
}

type Employment struct {
	JobTitle    string
	Employer    string
	Since       Date
	Until       Date
	Location    Location
	Description string
}

type Education struct {
	School      string
	Degree      string
	Since       Date
	Until       Date
	Location    Location
	Description string
}

type SkillLevel int

const (
	// SkillLevelBeginner indicates that a person can handle the basic features of
	// the program, but can’t do complicated tricks or troubleshoot problems yet.
	SkillLevelBeginner SkillLevel = 0

	// SkillLevelIntermediate indicates that a person can also troubleshoot and do
	// some fancy tricks, but might need to Google some functions or ask in forums
	// from time to time.
	SkillLevelIntermediate

	// SkillLevelProficient indicates that a person is not yet an expert, but can
	// handle advanced functions and troubleshoot problems by examining things on
	// their own. They don't need a manual.
	SkillLevelProficient

	// SkillLevelExpert indicates that a person knows the program like the back of
	// their hand. They know of obscure features, tricks, and weird problems, so
	// much so that other people often come to them for help.
	SkillLevelExpert
)

type Skill struct {
	Name  string
	Level SkillLevel
}

type LanguageLevel int

const (
	// LanguageLevelIntermediate indicates that a person can carry basic conversations
	// in a wide variety of situations, but still makes grammar mistakes, and has
	// limited working proficiency.
	LanguageLevelIntermediate LanguageLevel = 0

	// LanguageLevelAdvanced indicates that a person is skilled enough to carry
	// complex conversations but still puts in the conscious effort when speaking
	// and writing.
	LanguageLevelAdvanced

	// LanguageLevelFluent indicates that a person is full of professional working
	// proficiency and has fluid speech, reading, and writing, but with a less
	// advanced vocabulary than a native speaker.
	LanguageLevelFluent

	// LanguageLevelNativeOrBilingual indicates a full mastery of the language
	// through either upbringing and advanced education.
	LanguageLevelNativeOrBilingual
)

type Language struct {
	Name  string
	Level LanguageLevel
}
