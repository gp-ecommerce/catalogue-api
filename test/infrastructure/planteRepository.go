package infrastructure_test

import (
	"testing"

	"gitlab.com/constraintAutomaton/plante_maman/domain"
	"gitlab.com/constraintAutomaton/plante_maman/infrastructure"
)

func ProvidePlanteRepoTest() []func(infrastructure.PlanteRepository) (func(*testing.T), string) {
	return []func(infrastructure.PlanteRepository) (func(*testing.T), string){
		testAddAPlante,
		testAddPlanteMultipleTime,
		testAddMultiplePlantes,
	}
}
func testAddAPlante(repo infrastructure.PlanteRepository) (func(*testing.T), string) {
	return func(t *testing.T) {
		name := "foo"
		a_plant := domain.Plante{Name: name}
		repo.AddPlants(&a_plant)
		if l, _ := repo.Len(); l != 1 {
			t.Fatalf("the len should be 1")
		}
		if p, _ := repo.GetPlantByName(name); p != a_plant {
			t.Fatalf(`The plante are not consistant the plant added is:'%v'
			The returned plante is '%v'`, a_plant, p)
		}
	}, "TestAddAPlante"

}

func testAddPlanteMultipleTime(repo infrastructure.PlanteRepository) (func(*testing.T), string) {
	return func(t *testing.T) {
		plantes := []domain.Plante{
			{Name: "foo"},
			{Name: "bar"},
			{Name: "tar"},
		}

		for i, plante := range plantes {
			repo.AddPlants(&plante)
			if l, _ := repo.Len(); l != uint(i+1) {
				t.Fatalf("the len should be %v", i)
			}
			if p, _ := repo.GetPlantByName(plante.Name); p != plante {
				t.Fatalf(`The plante are not consistant the plant added is:'%v'
			The returned plante is '%v'`, plante, p)
			}
		}
	}, "TestAddPlanteMultipleTime"

}

func testAddMultiplePlantes(repo infrastructure.PlanteRepository) (func(*testing.T), string) {
	return func(t *testing.T) {
		plantes := []*domain.Plante{
			{Name: "foo"},
			{Name: "bar"},
			{Name: "tar"},
		}

		repo.AddPlants(plantes...)

		if l, _ := repo.Len(); l != uint(len(plantes)) {
			t.Fatalf("the len should be %v", len(plantes))
		}
		for _, plante := range plantes {
			if p, _ := repo.GetPlantByName(plante.Name); p != *plante {
				t.Fatalf(`The plante are not consistant the plant added is:'%v'
			The returned plante is '%v'`, *plante, p)
			}
		}

	}, "TestAddMultiplePlantes"

}
