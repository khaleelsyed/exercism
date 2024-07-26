from enum import Enum, auto


class Planet(Enum):
    MERCURY = auto()
    VENUS = auto()
    EARTH = auto()
    MARS = auto()
    JUPITER = auto()
    SATURN = auto()
    URANUS = auto()
    NEPTUNE = auto()


class SpaceAge:

    SECONDS_PER_EARTH_YEAR = 31557600

    RELATIVE_ORBITAL_PERIODS_EARTH_YEARS: dict[Planet, float] = {
        Planet.MERCURY: 0.2408467,
        Planet.VENUS: 0.61519726,
        Planet.EARTH: 1.0,
        Planet.MARS: 1.8808158,
        Planet.JUPITER: 11.862615,
        Planet.SATURN: 29.447498,
        Planet.URANUS: 84.016846,
        Planet.NEPTUNE: 164.79132,
    }

    def __init__(self, seconds: int) -> None:
        self._earth_years = seconds/self.SECONDS_PER_EARTH_YEAR

    def _get_planet_age(self, planet: Planet) -> float:
        return round((self._earth_years / self.RELATIVE_ORBITAL_PERIODS_EARTH_YEARS[planet]), 2)
    
    def on_mercury(self) -> float:
        return self._get_planet_age(Planet.MERCURY)
    
    def on_venus(self) -> float:
        return self._get_planet_age(Planet.VENUS)

    def on_earth(self) -> float:
        return self._get_planet_age(Planet.EARTH)

    def on_mars(self) -> float:
        return self._get_planet_age(Planet.MARS)
    
    def on_jupiter(self) -> float:
        return self._get_planet_age(Planet.JUPITER)
    
    def on_saturn(self) -> float:
        return self._get_planet_age(Planet.SATURN)

    def on_uranus(self) -> float:
        return self._get_planet_age(Planet.URANUS)
    
    def on_neptune(self) -> float:
        return self._get_planet_age(Planet.NEPTUNE)
