package storage

// Field defines all fields for an item.
type Field struct {
	// Key is a unique identifier for the field
	Key string
	// Name of the field, does not have to be unique
	Name string
	// Type of the field must match with the defined types
	// inside of the Data type.
	Type string
	// The name of a group the field is used.
	Group string
	// Items are ordered inside the group.
	Order int
	// If not nil that values are used for a dropdown list
	Select *[]string
}

func (f *Field) ID() string {
	return f.Key
}

type SortFields []Field

func (sf SortFields) Len() int           { return len(sf) }
func (sf SortFields) Swap(i, j int)      { sf[i], sf[j] = sf[j], sf[i] }
func (sf SortFields) Less(i, j int) bool { return sf[i].Order < sf[j].Order }
