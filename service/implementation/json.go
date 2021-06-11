package planterepository

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"

	"github.com/inventoryServices/domain"
	"gitlab.com/constraintAutomaton/plante_maman/infrastructure"
)

type jsonRepository struct {
	Inventory map[string]domain.Product   `json:"inventory"`
	Category  map[string][]domain.Product `json:"category"`
	FilePath  string                      `json:"-"`
}

func NewJsonRepository(filePath string) (infrastructure.PlanteRepository, error) {
	repo := jsonRepository{
		Feature:   map[string]domain.Plante{},
		Inventory: map[string]domain.Plante{},
		FilePath:  filePath,
	}
	err := repo.loadFile()
	if err != nil {
		return nil, fmt.Errorf(infrastructure.WAS_NOT_ABLE_TO_INSTANTIATE_REPO, err)
	}
	return &repo, nil
}
func (j *jsonRepository) GetFeaturePlantes(n uint) ([]domain.Plante, error) {
	plantes := make([]domain.Plante, n)
	var i uint = 0
	for _, plante := range j.Feature {
		plantes[i] = plante
		if i == n {
			return plantes, nil
		}
	}
	return []domain.Plante{}, nil
}

func (j *jsonRepository) GetPlantByName(name string) (domain.Plante, error) {
	plante, exist := j.Inventory[name]
	if exist {
		return plante, nil
	}
	return domain.Plante{}, fmt.Errorf(infrastructure.DONT_EXIST, name)
}

func (j *jsonRepository) AddPlants(plantes ...*domain.Plante) error {
	alreadyExist := make([]domain.Plante, 0, len(plantes))
	for _, plante := range plantes {
		if _, exist := j.Inventory[plante.Name]; !exist {
			j.Inventory[plante.Name] = *plante
		} else {
			alreadyExist = append(alreadyExist, *plante)
		}
	}
	if len(alreadyExist) != 0 {
		return fmt.Errorf(infrastructure.ALREADY_EXIST, alreadyExist)
	}
	err := j.updateFile()
	if err != nil {
		return fmt.Errorf(infrastructure.WAS_NOT_ABLE_TO_UPDATE_PERSISTANCE)
	}
	return nil
}

func (j *jsonRepository) SetToFeature(plantes ...*domain.Plante) error {
	alreadyInFeature := make([]domain.Plante, 0, len(plantes))
	for _, plante := range plantes {
		if _, exist := j.Feature[plante.Name]; !exist {
			j.Feature[plante.Name] = *plante
		} else {
			alreadyInFeature = append(alreadyInFeature, *plante)
		}
	}
	if len(alreadyInFeature) != 0 {
		return fmt.Errorf(infrastructure.ALREADY_FEATURE, alreadyInFeature)
	}
	return nil
}

func (j *jsonRepository) loadFile() error {
	data, err := ioutil.ReadFile(j.FilePath)
	if err != nil {
		err = ioutil.WriteFile(j.FilePath, []byte("{}"), fs.ModeDevice)
		if err != nil {
			return err
		}
	}
	if len(data) != 0 {
		err = json.Unmarshal(data, j)
		return err
	}
	return nil
}

func (j *jsonRepository) updateFile() error {
	b, err := json.Marshal(j)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(j.FilePath, b, fs.ModeDevice)
	return err
}

func (j *jsonRepository) Len() (uint, error) {
	return uint(len(j.Inventory)), nil
}

func (j *jsonRepository) Clear() error {
	j.Feature = map[string]domain.Plante{}
	j.Inventory = map[string]domain.Plante{}

	err := ioutil.WriteFile(j.FilePath, []byte("{}"), fs.ModeDevice)
	return err
}
