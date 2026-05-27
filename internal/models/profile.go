package models

type Profile struct {
	Name            string          `json:"name"`
	Title           string          `json:"title"`
	Subtitle        string          `json:"subtitle"`
	Bio             string          `json:"bio"`
	EducationInfo   EducationInfo   `json:"education_info"`
	TechTags        []string        `json:"tech_tags"`
	Stats           []Stat          `json:"stats"`
	Highlights      []Highlight     `json:"highlights"`
	Location        string          `json:"location"`
	SkillCategories []SkillCategory `json:"skill_categories"`
	Socials         []Social        `json:"socials"`
}

type EducationInfo struct {
	Degree      string `json:"degree"`
	Institution string `json:"institution"`
}

type Stat struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type Highlight struct {
	Title   string `json:"title"`
	Value   string `json:"value"`
	Subtext string `json:"subtext"`
}

// SkillCategory is a group of related skills (e.g. "Front-End Development")
type SkillCategory struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Icon        string      `json:"icon"` // react-icon component name e.g. "FaCode"
	ColSpan     int         `json:"col_span"` // 1, 2, or 3
	Skills      []SkillItem `json:"skills"`
}

// SkillItem is an individual technology inside a category
type SkillItem struct {
	Name string `json:"name"`
	Icon string `json:"icon"` // react-icon name e.g. "SiReact"
}

type Social struct {
	Platform string `json:"platform"`
	URL      string `json:"url"`
	Icon     string `json:"icon"`
}
