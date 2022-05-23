package value_object

import "time"

// Explicitness provides clearness for the outside world in cases
// where original types from Golang do not support particular behavior
// or supported behavior is not intuitive

type Birthday time.Time

func (b Birthday) IsYoungerThan(other time.Time) bool {
	return time.Time(b).After(other)
}
func (b Birthday) IsAdult() bool {
	return time.Time(b).AddDate(18, 0, 0).Before(time.Now())
}

const (
	Freelancer  = iota
	LLC         = 1
	Corporation = 2
)

type LegalForm int

func (l LegalForm) IsIndividual() bool {
	return l == Freelancer
}
func (l LegalForm) HasLimitedResponsability() bool {
	return l == LLC || l == Corporation
}

// Sometimes Value Object does not need to be explicitly defined as part of any
// other Entity or Value Object. Still, we can define a Value Object as a HELPER OBJECT
// that provides clearness for later usage in code. In case with dealing with
// a Customer who can be a Person or a Company. Depending on the type of the Customer
// we have different flows in the application
type Customer struct {
	ID        string
	Name      string
	LegalForm LegalForm
	Date      time.Time
}

func (c Customer) toPerson() Person {
	return Person{
		FullName: c.Name,
		Birthday: Birthday(c.Date),
	}
}
func (c Customer) toCompany() Company {
	return Company{
		Name:         c.Name,
		CreationDate: c.Date,
	}
}

type Person struct {
	FullName string
	Birthday Birthday
}
type Company struct {
	Name         string
	CreationDate time.Time
}

// Whenever we notice that some particular smaller group of fields
// interacts with each other constantly, but they are inside some larger group
// this is already a sign that we should group hem into Value Object and use it
// like that inside our larger group (which now become smaller)
