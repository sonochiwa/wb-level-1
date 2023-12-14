package exercise_21

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

// LocationProvider - целевой интерфейс в вашем коде
type LocationProvider interface {
	GetLocation() string
}

// GeoLibrary - сторонняя библиотека с несовместимым интерфейсом
type GeoLibrary struct {
	Longitude float64
	Latitude  float64
}

// GeoLibraryAdapter - адаптер для интеграции сторонней библиотеки
type GeoLibraryAdapter struct {
	GeoLibrary *GeoLibrary
}

// GetLocation - метод адаптера для удовлетворения интерфейса LocationProvider
func (adapter *GeoLibraryAdapter) GetLocation() string {
	return fmt.Sprintf("Latitude: %f, Longitude: %f", adapter.GeoLibrary.Latitude, adapter.GeoLibrary.Longitude)
}

func Run() {
	// Создаем объект сторонней библиотеки
	geoLibrary := &GeoLibrary{
		Longitude: 73.366669,
		Latitude:  54.983334,
	}

	// Создаем адаптер и передаем ему объект сторонней библиотеки
	adapter := &GeoLibraryAdapter{
		GeoLibrary: geoLibrary,
	}

	// Используем целевой интерфейс через адаптер
	location := adapter.GetLocation()
	fmt.Println("Location:", location)
}
