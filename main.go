package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rl.InitWindow(int32(rl.GetMonitorWidth(rl.GetCurrentMonitor())), int32(rl.GetMonitorHeight(rl.GetCurrentMonitor())), "learning GO - 7-Segment Display")
	rl.ToggleFullscreen()
	defer rl.CloseWindow()

	segments := []bool{true, true, true, true, true, false, true}
	segmentsnumber := 0

	segmentcor := []rl.Vector2{
		rl.NewVector2(float32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))/2-150, float32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))/2-225),
		rl.NewVector2(float32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))/2-150, float32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))/2+25),
		rl.NewVector2(float32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))/2+100, float32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))/2-225),
		rl.NewVector2(float32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))/2+100, float32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))/2+25),
		rl.NewVector2(float32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))/2-100, float32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))/2-275),
		rl.NewVector2(float32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))/2-100, float32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))/2-25),
		rl.NewVector2(float32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))/2-100, float32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))/2+225),
	}

	segnumbers := [][]bool{
		{true, true, true, true, true, false, true},
		{false, false, true, true, false, false, false},
		{false, true, true, false, true, true, true},
		{false, false, true, true, true, true, true},
		{true, false, true, true, false, true, false},
		{true, false, false, true, true, true, true},
		{true, true, false, true, true, true, true},
		{false, false, true, true, true, false, false},
		{true, true, true, true, true, true, true},
		{true, false, true, true, true, true, true},
	}

	rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawText("Use mouse wheel to cycle between numbers 0-9", 20, 20, 50, rl.DarkGray)

		for i := range segments {
			if segments[i] {
				if i >= 4 {
					rl.DrawRectangleRounded(rl.NewRectangle(segmentcor[i].X, segmentcor[i].Y, 200, 50), 1, 1, rl.Green)
				} else {
					rl.DrawRectangleRounded(rl.NewRectangle(segmentcor[i].X, segmentcor[i].Y, 50, 200), 1, 1, rl.Green)
				}

			} else {
				if i >= 4 {
					rl.DrawRectangleRounded(rl.NewRectangle(segmentcor[i].X, segmentcor[i].Y, 200, 50), 1, 1, rl.DarkGray)
				} else {
					rl.DrawRectangleRounded(rl.NewRectangle(segmentcor[i].X, segmentcor[i].Y, 50, 200), 1, 1, rl.DarkGray)
				}

			}
		}

		if rl.GetMouseWheelMove() > 0 && segmentsnumber < 9 {
			segmentsnumber++
			segments = segnumbers[segmentsnumber]
		}

		if rl.GetMouseWheelMove() < 0 && segmentsnumber > 0 {
			segmentsnumber--
			segments = segnumbers[segmentsnumber]
		}

		rl.EndDrawing()
	}
}
