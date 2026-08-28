//go:build debug

package main

// debugEnabled is true only in builds/runs that pass -tags debug.
// See debug_off.go for the complementary, default definition.
const debugEnabled = true
