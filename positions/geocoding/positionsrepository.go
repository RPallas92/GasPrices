package geocoding

import "rpallas/oil/prices/positions"

type PositionsRepository struct {
	Datasource *PositionsDatasource
}

func (repository *PositionsRepository) GetPositions(originCity, destinationCity string) ([]positions.Position, error) {
	originChannel := make(chan PositionResult)
	destinationChannel := make(chan PositionResult)

	go repository.getPosition(originCity, originChannel)
	go repository.getPosition(destinationCity, destinationChannel)

	originPositionResult, destinationPositionResult := <-originChannel, <-destinationChannel

	if originPositionResult.Error != nil {
		return []positions.Position{}, originPositionResult.Error
	}
	if destinationPositionResult.Error != nil {
		return []positions.Position{}, destinationPositionResult.Error
	}

	return []positions.Position{originPositionResult.Position, destinationPositionResult.Position}, nil
}

func (repository *PositionsRepository) getPosition(city string, channel chan PositionResult) {
	positionResponse, err := repository.Datasource.getPosition(city)
	if err != nil {
		channel <- PositionResult{
			Position: positions.Position{},
			Error:    err,
		}
	}
	position, err := positionResponse.ToPosition()
	if err != nil {
		channel <- PositionResult{
			Position: positions.Position{},
			Error:    err,
		}
	}

	channel <- PositionResult{
		Position: position,
		Error:    nil,
	}
}

type PositionResult struct {
	Position positions.Position
	Error    error
}
