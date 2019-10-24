package asm

type Executable struct {
    Main byte
    Sections []Section
}

type SectionType string

const (
    SectionText SectionType = "text"
    SectionData SectionType = "data"
)

type Section struct {
    Type SectionType
    Lines []Line
}

type Line struct {
    Tag string
    Directive string
    Args []string
    Comment string
}
