//go:build !debug

package main

// debugEnabled is false in ordinary builds — this file is used whenever
// the custom "debug" build tag is NOT passed.
const debugEnabled = false
