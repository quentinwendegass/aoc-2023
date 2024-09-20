package main

import (
	"aoc2023/utils"
	"slices"
	"strings"
)

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	moduleMap := convertInput(string(data))

	pulseQueue := utils.CreateQueue[Pulse]()

	lowCount := 0
	highCount := 0

	for i := 0; i < 1000; i++ {
		lowCount++
		pulseQueue.Push(moduleMap["broadcast"].process(Pulse{pulseType: LowPulse})...)
		for pulseQueue.Len() > 0 {
			pulse := pulseQueue.Pop()
			if pulse.pulseType == LowPulse {
				lowCount++
			} else {
				highCount++
			}

			pulses := moduleMap[pulse.dst].process(pulse)
			pulseQueue.Push(pulses...)
		}
	}

	return lowCount * highCount
}

// This is basically a best effort implementation and assumes that the modules before rx are having a certain pattern.
// The problem description doesn't really clarify this, but what I've seen online it seems that most (all) inputs are following this pattern.
func partTwo(data []byte) int {
	moduleMap := convertInput(string(data))

	pulseQueue := utils.CreateQueue[Pulse]()

	moduleToRxKey := moduleMap["rx"].getInputs()[0]
	moduleToRx := moduleMap[moduleToRxKey]

	if _, ok := moduleToRx.(*ConjunctionModule); !ok {
		panic("Not implemented. Not sure if input data even allows that.")
	}

	inputsToConjunction := moduleToRx.getInputs()
	iterationsForInputs := make([]int, 0)

	buttonPressCount := 0

ButtonPress:
	for {
		buttonPressCount++
		pulseQueue.Push(moduleMap["broadcast"].process(Pulse{pulseType: LowPulse})...)
		for pulseQueue.Len() > 0 {
			pulse := pulseQueue.Pop()

			if idx := slices.Index(inputsToConjunction, pulse.src); pulse.dst == moduleToRxKey && pulse.pulseType == HighPulse && idx != -1 {
				inputsToConjunction = utils.RemoveIndex(inputsToConjunction, idx)
				iterationsForInputs = append(iterationsForInputs, buttonPressCount)
			}

			if len(inputsToConjunction) == 0 {
				break ButtonPress
			}

			pulses := moduleMap[pulse.dst].process(pulse)
			pulseQueue.Push(pulses...)
		}
	}

	return utils.Multiply(iterationsForInputs...)
}

type PulseType int

const (
	LowPulse  PulseType = 0
	HighPulse PulseType = 1
)

type ModuleKey string

type Pulse struct {
	dst       ModuleKey
	src       ModuleKey
	pulseType PulseType
}

type FlipFlopModule struct {
	name    ModuleKey
	state   bool
	inputs  []ModuleKey
	outputs []ModuleKey
}

type ConjunctionModule struct {
	name        ModuleKey
	pulseStates map[ModuleKey]PulseType
	inputs      []ModuleKey
	outputs     []ModuleKey
}

type BroadcastModule struct {
	name    ModuleKey
	outputs []ModuleKey
}

type OutputModule struct {
	name   ModuleKey
	inputs []ModuleKey
}

type Module interface {
	process(pulse Pulse) []Pulse
	addInput(input ModuleKey)
	addOutput(output ModuleKey)
	getInputs() []ModuleKey
}

func createFlipFlopModule(name ModuleKey) *FlipFlopModule {
	inputs := make([]ModuleKey, 0)
	outputs := make([]ModuleKey, 0)

	return &FlipFlopModule{inputs: inputs, outputs: outputs, state: false, name: name}
}

func createConjunctionModule(name ModuleKey) *ConjunctionModule {
	inputs := make([]ModuleKey, 0)
	outputs := make([]ModuleKey, 0)
	pulseStates := make(map[ModuleKey]PulseType)

	return &ConjunctionModule{inputs: inputs, outputs: outputs, pulseStates: pulseStates, name: name}
}

func createOutputModule(name ModuleKey) *OutputModule {
	inputs := make([]ModuleKey, 0)

	return &OutputModule{inputs: inputs, name: name}
}

func createBroadcastModule(name ModuleKey) *BroadcastModule {
	outputs := make([]ModuleKey, 0)

	return &BroadcastModule{outputs: outputs, name: name}
}

func (module *FlipFlopModule) process(pulse Pulse) []Pulse {
	if pulse.pulseType != LowPulse {
		return []Pulse{}
	}

	var pulseType PulseType
	if !module.state {
		pulseType = HighPulse
	} else {
		pulseType = LowPulse
	}

	module.state = !module.state

	return gatherPulsesToSent(module.outputs, module.name, pulseType)
}

func (module *FlipFlopModule) addInput(input ModuleKey) {
	module.inputs = append(module.inputs, input)
}

func (module *FlipFlopModule) getInputs() []ModuleKey {
	return module.inputs
}

func (module *FlipFlopModule) addOutput(output ModuleKey) {
	module.outputs = append(module.outputs, output)
}

func (module *ConjunctionModule) process(pulse Pulse) []Pulse {
	module.pulseStates[pulse.src] = pulse.pulseType

	pulseType := LowPulse
	for _, state := range module.pulseStates {
		if state == LowPulse {
			pulseType = HighPulse
			break
		}
	}

	return gatherPulsesToSent(module.outputs, module.name, pulseType)
}

func (module *ConjunctionModule) addInput(input ModuleKey) {
	module.inputs = append(module.inputs, input)
	module.pulseStates[input] = LowPulse
}

func (module *ConjunctionModule) getInputs() []ModuleKey {
	return module.inputs
}

func (module *ConjunctionModule) addOutput(output ModuleKey) {
	module.outputs = append(module.outputs, output)
}

func (module *BroadcastModule) process(pulse Pulse) []Pulse {
	return gatherPulsesToSent(module.outputs, module.name, pulse.pulseType)
}

func (module *BroadcastModule) addInput(input ModuleKey) {
	panic("Broadcast module cannot have inputs")
}

func (module *BroadcastModule) getInputs() []ModuleKey {
	panic("Broadcast module cannot have inputs")
}

func (module *BroadcastModule) addOutput(output ModuleKey) {
	module.outputs = append(module.outputs, output)
}

func (module *OutputModule) process(pulse Pulse) []Pulse {
	return []Pulse{}
}

func (module *OutputModule) addInput(input ModuleKey) {
	module.inputs = append(module.inputs, input)
}

func (module *OutputModule) getInputs() []ModuleKey {
	return module.inputs
}

func (module *OutputModule) addOutput(output ModuleKey) {
	panic("Output module cannot have outputs")
}

func gatherPulsesToSent(dst []ModuleKey, src ModuleKey, pulseType PulseType) []Pulse {
	pulsesToSent := make([]Pulse, len(dst))
	for i, moduleKey := range dst {
		pulsesToSent[i] = Pulse{dst: moduleKey, src: src, pulseType: pulseType}
	}

	return pulsesToSent
}

func convertInput(data string) map[ModuleKey]Module {
	lines := strings.Split(data, "\n")

	moduleMap := make(map[ModuleKey]Module)

	for _, line := range lines {
		module := strings.Split(line, " ")[0]

		if module[0] == '%' {
			moduleMap[ModuleKey(module[1:])] = createFlipFlopModule(ModuleKey(module[1:]))
		} else if module[0] == '&' {
			moduleMap[ModuleKey(module[1:])] = createConjunctionModule(ModuleKey(module[1:]))
		} else {
			moduleMap["broadcast"] = createBroadcastModule("broadcast")
		}
	}

	for _, line := range lines {
		moduleName := ModuleKey(strings.Split(line, " ")[0][1:])
		// Amazing code
		if moduleName == "roadcaster" {
			moduleName = "broadcast"
		}

		moduleConnections := strings.Split(strings.Split(line, "-> ")[1], ", ")

		for _, connection := range moduleConnections {
			moduleMap[moduleName].addOutput(ModuleKey(connection))
			_, ok := moduleMap[ModuleKey(connection)]

			if !ok {
				moduleMap[ModuleKey(connection)] = createOutputModule(ModuleKey(connection))
			}

			moduleMap[ModuleKey(connection)].addInput(ModuleKey(moduleName))
		}
	}

	return moduleMap
}
