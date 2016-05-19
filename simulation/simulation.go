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
	Simulatables        []Simulatable
	simulationTime      time.Time
	ticker              *time.Ticker
	sendCycleFinished   chan<- time.Time
	CycleFinishedTicker <-chan time.Time
}

func NewSimulation() *Simulation {
	s := new(Simulation)
	s.Simulatables = make([]Simulatable, 0)
	outputChan := make(chan time.Time, 1) // Similar to time.Timer, buffer 1 tick
	s.sendCycleFinished = outputChan
	s.CycleFinishedTicker = outputChan
	return s
}

func (s *Simulation) AddSimulatable(thing Simulatable) {
	s.Simulatables = append(s.Simulatables, thing)
}

func (s *Simulation) StartTime() {
	s.simulationTime = time.Now()
	s.ticker = time.NewTicker(time.Second / SIMULATION_CYCLES_PER_SECOND)
	go func() {
		for tick := range s.ticker.C {
			now := time.Now()
			timeDifference := now.Sub(s.simulationTime).Seconds()
			s.simulationTime = now
			for _, simObj := range s.Simulatables {
				simObj.Move(timeDifference)
			}
			s.sendCycleFinished <- tick
		}
	}()
}

func (s *Simulation) StopTime() {
	s.ticker.Stop()
}
