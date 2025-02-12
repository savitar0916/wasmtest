package main

import (
	"fmt"
	"os"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	err := loadWasmFunc("module.wasm")
	if err != nil {
		panic(err)
	}
}

func loadWasmFunc(fileName string) error {
	wasmBytes, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		return err
	}

	wasiEnv, _ := wasmer.NewWasiStateBuilder("wasi-program").
		// Choose according to your actual situation
		// Argument("--foo").
		// Environment("ABC", "DEF").
		// MapDirectory("./", ".").
		Finalize()
	// Instantiates the module
	// importObject := wasmer.NewImportObject()
	importObject, err := wasiEnv.GenerateImportObject(store, module)
	if err != nil {
		return err
	}

	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		return err
	}
	// start, err := instance.Exports.GetWasiStartFunction()
	// if err != nil {
	// 	return err
	// }
	// start()
	Add, err := instance.Exports.GetFunction("Add")
	if err != nil {
		return err
	}
	// Gets the `sum` exported function from the WebAssembly instance.
	// add, _ := instance.Exports.GetFunction("add")

	result, _ := Add(1, 2)
	fmt.Println(result)
	return nil
}
