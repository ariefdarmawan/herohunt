package herohunt

type AreaTypeEnum string

const (
	AreaTypeRoad     AreaTypeEnum = "Road"
	AreaTypeRiver    AreaTypeEnum = "River"
	AreaTypeJungle   AreaTypeEnum = "Jungle"
	AreaTypeRock     AreaTypeEnum = "Rock"
	AreaTypeMountain AreaTypeEnum = "Mountain"
)

type GameMap struct {
	Width  int
	Length int
}

type MapPoint struct {
	X, Y     int
	AreaType AreaTypeEnum
}
