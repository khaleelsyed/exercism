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
        """
        Calculate the age (in years) relative to planets orbital lengths.

        Args:
            seconds: Number of seconds since birth.
        """
        self._earth_years = seconds/self.SECONDS_PER_EARTH_YEAR
        self._relative_age: dict[Planet, float] = {
            planet: round(self._earth_years / orbit_time_relative_to_earth, 2)
            for planet, orbit_time_relative_to_earth in self.RELATIVE_ORBITAL_PERIODS_EARTH_YEARS.items()
        }
    
    def on_mercury(self) -> float:
        """Relative age in years on mercury."""
        return self._relative_age[Planet.MERCURY]
    
    def on_venus(self) -> float:
        """Relative age in years on venus."""
        return self._relative_age[Planet.VENUS]

    def on_earth(self) -> float:
        """Relative age in years on earth."""
        return self._relative_age[Planet.EARTH]

    def on_mars(self) -> float:
        """Relative age in years on mars."""
        return self._relative_age[Planet.MARS]
    
    def on_jupiter(self) -> float:
        """Relative age in years on jupiter."""
        return self._relative_age[Planet.JUPITER]
    
    def on_saturn(self) -> float:
        """Relative age in years on saturn."""
        return self._relative_age[Planet.SATURN]

    def on_uranus(self) -> float:
        """Relative age in years on uranus."""
        return self._relative_age[Planet.URANUS]
    
    def on_neptune(self) -> float:
        """Relative age in years on neptune."""
        return self._relative_age[Planet.NEPTUNE]
