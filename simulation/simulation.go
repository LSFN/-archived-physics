package simulation

import (
	"time"

	"github.com/LSFN/physics/movement"
)

const (
	SIMULATION_CYCLES_PER_SECOND = 60
)

type Simulatable interface {
	movement.Positionable
	movement.Mover
}

type Simulation struct {
	simulatables   []Simulatable
	simulationTime time.Time
	ticker         time.Ticker
}

func NewSimulation() *Simulation {
	s := new(Simulation)
	s.simulatables = make([]Simulatable, 0)
	return s
}

func (s *Simulation) AddSimulatable(thing Simulatable) {
	s.simulatables = append(s.simulatables, thing)
}

func (s *Simulation) StartTime() {
	s.simulationTime = time.Now()
	s.ticker = time.NewTicker(time.Second / SIMULATION_CYCLES_PER_SECOND)
	go func() {
		for tick := range s.ticker.C {
			now := time.Now()
			timeDifference := now.Sub(s.simulationTime).Seconds()
			s.simulationTime = now
			for simObj := range s.simulatables {
				simObj.Move(timeDifference)
			}
		}
	}()
}

func (s *Simulation) StopTime() {
	s.ticker.Stop()
}
