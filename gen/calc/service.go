// Code generated by goa v3.0.2, DO NOT EDIT.
//
// calc service
//
// Command:
// $ goa gen goa-ddb/design

package calc

import (
	"context"
	calcviews "goa-ddb/gen/calc/views"
)

// The calc service performs operations on numbers
type Service interface {
	// Add implements add.
	Add(context.Context, *AddPayload) (res int, err error)
	// Addresume implements addresume.
	Addresume(context.Context, []*Resume) (res []int, err error)
	// List implements list.
	List(context.Context) (res StoredResumeCollection, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "calc"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"add", "addresume", "list"}

// AddPayload is the payload type of the calc service add method.
type AddPayload struct {
	// Left operand
	A int
	// Right operand
	B int
}

// StoredResumeCollection is the result type of the calc service list method.
type StoredResumeCollection []*StoredResume

type Resume struct {
	// Name in the resume
	Name string
	// Experience section in the resume
	Experience []*Experience
	// Education section in the resume
	Education []*Education
}

type Experience struct {
	// Name of the company
	Company string
	// Name of the role in the company
	Role string
	// Duration (in years) in the company
	Duration int
}

type Education struct {
	// Name of the institution
	Institution string
	// Major name
	Major string
}

type StoredResume struct {
	// ID of the resume
	ID int
	// Time when resume was created
	CreatedAt string
	// Name in the resume
	Name string
	// Experience section in the resume
	Experience []*Experience
	// Education section in the resume
	Education []*Education
}

// NewStoredResumeCollection initializes result type StoredResumeCollection
// from viewed result type StoredResumeCollection.
func NewStoredResumeCollection(vres calcviews.StoredResumeCollection) StoredResumeCollection {
	var res StoredResumeCollection
	switch vres.View {
	case "default", "":
		res = newStoredResumeCollection(vres.Projected)
	}
	return res
}

// NewViewedStoredResumeCollection initializes viewed result type
// StoredResumeCollection from result type StoredResumeCollection using the
// given view.
func NewViewedStoredResumeCollection(res StoredResumeCollection, view string) calcviews.StoredResumeCollection {
	var vres calcviews.StoredResumeCollection
	switch view {
	case "default", "":
		p := newStoredResumeCollectionView(res)
		vres = calcviews.StoredResumeCollection{p, "default"}
	}
	return vres
}

// newStoredResumeCollection converts projected type StoredResumeCollection to
// service type StoredResumeCollection.
func newStoredResumeCollection(vres calcviews.StoredResumeCollectionView) StoredResumeCollection {
	res := make(StoredResumeCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredResume(n)
	}
	return res
}

// newStoredResumeCollectionView projects result type StoredResumeCollection to
// projected type StoredResumeCollectionView using the "default" view.
func newStoredResumeCollectionView(res StoredResumeCollection) calcviews.StoredResumeCollectionView {
	vres := make(calcviews.StoredResumeCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredResumeView(n)
	}
	return vres
}

// newStoredResume converts projected type StoredResume to service type
// StoredResume.
func newStoredResume(vres *calcviews.StoredResumeView) *StoredResume {
	res := &StoredResume{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.CreatedAt != nil {
		res.CreatedAt = *vres.CreatedAt
	}
	if vres.Experience != nil {
		res.Experience = make([]*Experience, len(vres.Experience))
		for i, val := range vres.Experience {
			res.Experience[i] = transformCalcviewsExperienceViewToExperience(val)
		}
	}
	if vres.Education != nil {
		res.Education = make([]*Education, len(vres.Education))
		for i, val := range vres.Education {
			res.Education[i] = transformCalcviewsEducationViewToEducation(val)
		}
	}
	return res
}

// newStoredResumeView projects result type StoredResume to projected type
// StoredResumeView using the "default" view.
func newStoredResumeView(res *StoredResume) *calcviews.StoredResumeView {
	vres := &calcviews.StoredResumeView{
		ID:        &res.ID,
		CreatedAt: &res.CreatedAt,
		Name:      &res.Name,
	}
	if res.Experience != nil {
		vres.Experience = make([]*calcviews.ExperienceView, len(res.Experience))
		for i, val := range res.Experience {
			vres.Experience[i] = transformExperienceToCalcviewsExperienceView(val)
		}
	}
	if res.Education != nil {
		vres.Education = make([]*calcviews.EducationView, len(res.Education))
		for i, val := range res.Education {
			vres.Education[i] = transformEducationToCalcviewsEducationView(val)
		}
	}
	return vres
}

// transformCalcviewsExperienceViewToExperience builds a value of type
// *Experience from a value of type *calcviews.ExperienceView.
func transformCalcviewsExperienceViewToExperience(v *calcviews.ExperienceView) *Experience {
	if v == nil {
		return nil
	}
	res := &Experience{
		Company:  *v.Company,
		Role:     *v.Role,
		Duration: *v.Duration,
	}

	return res
}

// transformCalcviewsEducationViewToEducation builds a value of type *Education
// from a value of type *calcviews.EducationView.
func transformCalcviewsEducationViewToEducation(v *calcviews.EducationView) *Education {
	if v == nil {
		return nil
	}
	res := &Education{
		Institution: *v.Institution,
		Major:       *v.Major,
	}

	return res
}

// transformExperienceToCalcviewsExperienceView builds a value of type
// *calcviews.ExperienceView from a value of type *Experience.
func transformExperienceToCalcviewsExperienceView(v *Experience) *calcviews.ExperienceView {
	res := &calcviews.ExperienceView{
		Company:  &v.Company,
		Role:     &v.Role,
		Duration: &v.Duration,
	}

	return res
}

// transformEducationToCalcviewsEducationView builds a value of type
// *calcviews.EducationView from a value of type *Education.
func transformEducationToCalcviewsEducationView(v *Education) *calcviews.EducationView {
	res := &calcviews.EducationView{
		Institution: &v.Institution,
		Major:       &v.Major,
	}

	return res
}
