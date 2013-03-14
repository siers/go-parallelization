package main

import (
	"fmt"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"os"
	"runtime"
)

const CPUS = 4

func main() {
	runtime.GOMAXPROCS(CPUS)
	var err error
	if err = glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		return
	}

	defer glfw.Terminate()

	if err = glfw.OpenWindow(256, 256, 8, 8, 8, 0, 0, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		return
	}

	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle("Simple GLFW window")
	glfw.SetWindowSizeCallback(onResize)
	glfw.SetWindowCloseCallback(onClose)
	glfw.SetKeyCallback(onKey)
	glfw.SetCharCallback(onChar)

	// Start loop
	running := true
	for running {
		// OpenGL rendering goes here.
		glfw.SwapBuffers()
		running = glfw.Key(glfw.KeyEsc) == 0 && glfw.WindowParam(glfw.Opened) == 1
	}
}

func onClose() int {
	fmt.Println("closed")
	return 1 // return 0 to keep window open.
}

func onKey(key, state int) {
	fmt.Printf("key: %d, %d\n", key, state)
}

func onChar(key, state int) {
	fmt.Printf("char: %d, %d\n", key, state)
}

func onResize(w, h int) {
	gl.DrawBuffer(gl.FRONT_AND_BACK)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Viewport(0, 0, w, h)
	gl.Ortho(0, float64(w), float64(h), 0, -1.0, 1.0)
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	draw(w, h)
	fmt.Printf("resized: %dx%d\n", w, h)
}

func draw(w, h int) {
	zoom[0] = float32(w) * 0.35
	zoom[1] = float32(h) * 0.55
	cs := Calculate(w, h)

	gl.Begin(gl.POINTS)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			color := float32(<-cs[i]) / float32(iters)
			gl.Color3f(0, 0, color)
			gl.Vertex2f(float32(i), float32(j))
		}
	}
	gl.End()
}
