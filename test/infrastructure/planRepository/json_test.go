package json_test

import (
	"testing"

	planterepository "gitlab.com/constraintAutomaton/plante_maman/infrastructure/planteRepository"
	infrastructure_test "gitlab.com/constraintAutomaton/plante_maman/test/infrastructure"
)

func TestJsonRepository(t *testing.T) {
	for _, testCase := range infrastructure_test.ProvidePlanteRepoTest() {
		repo, err := planterepository.NewJsonRepository("./json_repo.json")
		if err != nil {
			t.Fatalf("was not able to load the repository; '%v'", err)
		}
		testFunction, name := testCase(repo)
		t.Run(name, testFunction)
		if err := repo.Clear(); err != nil {
			t.Fatalf("was not able to clear to repository. Because error '%v'", err)
		}
	}
}
