package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/image/draw"

	"github.com/benhili/advent-of-code-2024/utils"
)

type Robot struct {
	position image.Point
	velocity image.Point
}

func main() {
	part, file, start := utils.Setup("input.txt")
	robots := parse(string(file))
	if part == 1 {
		ans := part1(robots)
		fmt.Println("Output:", ans)
	} else if part == 2 {
		ans := part2(robots, false)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(robots, true)
		fmt.Println("Output:", ans)
	}
	utils.Finish(start)
}

func unsafeAtoi(str string) int {
	num, _ := strconv.Atoi(str)

	return num
}

func parse(file string) []Robot {
	var robots []Robot

	for _, line := range strings.Split(file, "\n") {
		r, _ := regexp.Compile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
		matches := r.FindAllStringSubmatch(line, -1)
		var newRobot Robot
		newRobot.position = image.Point{unsafeAtoi(matches[0][1]), unsafeAtoi(matches[0][2])}
		newRobot.velocity = image.Point{unsafeAtoi(matches[0][3]), unsafeAtoi(matches[0][4])}
		robots = append(robots, newRobot)

	}
	fmt.Println("parsing complete")
	return robots
}

const WIDTH = 101
const HEIGHT = 103
const SECONDS = 100
const MIDDLE_HEIGHT = (HEIGHT - 1) / 2
const MIDDLE_WIDTH = (WIDTH - 1) / 2

func safetyFactor(robots []Robot) int {
	tl := 0
	tr := 0
	bl := 0
	br := 0

	for _, robot := range robots {
		if robot.position.Y < MIDDLE_HEIGHT {
			// top
			if robot.position.X < MIDDLE_WIDTH {
				tl++
			} else if robot.position.X > MIDDLE_WIDTH {
				tr++
			}

		} else if robot.position.Y > MIDDLE_HEIGHT {
			// bottom
			if robot.position.X < MIDDLE_WIDTH {
				bl++
			} else if robot.position.X > MIDDLE_WIDTH {
				br++
			}
		}

	}

	fmt.Println(tl, tr, bl, br)

	return tl * tr * bl * br
}

func part1(robots []Robot) int {
	for range SECONDS {
		for i := 0; i < len(robots); i++ {
			robot := &robots[i]
			robot.position = robot.position.Add(robot.velocity)
			// horizontal wrap
			if robot.position.X >= WIDTH {
				robot.position.X = robot.position.X - WIDTH
			} else if robot.position.X < 0 {
				robot.position.X = WIDTH + robot.position.X
			}

			// vertical wrap
			if robot.position.Y >= HEIGHT {
				robot.position.Y = robot.position.Y - HEIGHT
			} else if robot.position.Y < 0 {
				robot.position.Y = HEIGHT + robot.position.Y
			}

		}

	}

	return safetyFactor(robots)
}

func part2(robots []Robot, visualise bool) int {
	count := 0
	for {
		count++
		robotOnSpace := make(map[image.Point]bool)
		collision := false
		for i := 0; i < len(robots); i++ {
			robot := &robots[i]
			robot.position = robot.position.Add(robot.velocity)
			// horizontal wrap
			if robot.position.X >= WIDTH {
				robot.position.X = robot.position.X - WIDTH
			} else if robot.position.X < 0 {
				robot.position.X = WIDTH + robot.position.X
			}

			// vertical wrap
			if robot.position.Y >= HEIGHT {
				robot.position.Y = robot.position.Y - HEIGHT
			} else if robot.position.Y < 0 {
				robot.position.Y = HEIGHT + robot.position.Y
			}
			if robotOnSpace[robot.position] {
				collision = true
			}
			robotOnSpace[robot.position] = true
		}
		if !collision {
			break
		}
	}

	if visualise {
		printRobots(robots)
	}
	return count
}

func printRobots(robots []Robot) {

	// 2. Create a blank RGBA image with a fixed width and height
	width, height := 101, 103
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 3. Fill the image with a white background (optional)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.White)
		}
	}

	// 4. Draw the points on the image (use black color for this example)
	red := color.RGBA{255, 0, 0, 255}
	green := color.RGBA{55, 139, 41, 255}
	for _, robot := range robots {
		if robot.position.In(img.Bounds()) { // Ensure point is within image bounds
			var pointColor color.RGBA
			if robot.position.Y%3 == 0 {
				pointColor = red
			} else {
				pointColor = green
			}
			img.Set(robot.position.X, robot.position.Y, pointColor)
		}
	}

	scale := 10
	newWidth, newHeight := width*scale, height*scale
	scaledImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	// Use NearestNeighbor to make it pixelated or ApproxBiLinear for smooth resizing
	draw.NearestNeighbor.Scale(scaledImg, scaledImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	// 5. Save the image as a PNG file
	outputFile, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, scaledImg)
	if err != nil {
		panic(err)
	}

	println("Image saved as output.png")
}
