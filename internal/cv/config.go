package cv

type Config struct {
	FirstName         string       `yaml:"first_name"`
	LastName          string       `yaml:"last_name"`
	Occupation        string       `yaml:"occupation"`
	Location          Location     `yaml:"location"`
	PhoneNumber       string       `yaml:"phone_number"`
	Email             string       `yaml:"email"`
	Summary           string       `yaml:"summary"`
	Links             Links        `yaml:"links"`
	EmploymentHistory []Employment `yaml:"employment"`
	EducationHistory  []Education  `yaml:"education"`
	Skills            []Skill      `yaml:"skills"`
	Languages         []Language   `yaml:"languages"`
}

type Location struct {
	City    string `yaml:"city"`
	Country string `yaml:"country"`
}

type Links struct {
	PersonalWebsiteURL string `yaml:"personal_website"`
	GithubURL          string `yaml:"github"`
	LinkedinURL        string `yaml:"linkedin"`
}

type Date struct {
	Month int `yaml:"month"`
	Year  int `yaml:"year"`
}

type Employment struct {
	JobTitle    string   `yaml:"job_title"`
	Employer    string   `yaml:"employer"`
	Since       Date     `yaml:"since"`
	Until       Date     `yaml:"until"`
	Location    Location `yaml:"location"`
	Description string   `yaml:"description"`
}

type Education struct {
	School      string   `yaml:"school"`
	Degree      string   `yaml:"degree"`
	Since       Date     `yaml:"since"`
	Until       Date     `yaml:"until"`
	Location    Location `yaml:"location"`
	Description string   `yaml:"description"`
}

type SkillLevel int

const (
	SkillLevelBeginner LanguageLevel = 1

	// SkillLevelBeginner indicates that a person can handle the basic features of
	// the program, but can’t do complicated tricks or troubleshoot problems yet.
	SkillLevelIntermediate

	// SkillLevelIntermediate indicates that a person can also troubleshoot and do
	// some fancy tricks, but might need to Google some functions or ask in forums
	// from time to time.
	SkillLevelAdvanced

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
	Name  string     `yaml:"name"`
	Level SkillLevel `yaml:"level"`
}

type LanguageLevel int

const (
	LanguageLevelBeginner LanguageLevel = 1

	// LanguageLevelIntermediate indicates that a person can carry basic conversations
	// in a wide variety of situations, but still makes grammar mistakes, and has
	// limited working proficiency.
	LanguageLevelIntermediate

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
	Name  string        `yaml:"name"`
	Level LanguageLevel `yaml:"level"`
}
