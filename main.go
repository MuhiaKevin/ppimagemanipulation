package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PPMImage struct {
    Width, Height, MaxColorValue int
    Pixels                        [][][3]byte // [R, G, B] values for each pixel
}

func readPPMFile(filename string) (*PPMImage, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // Read header
    scanner.Scan()
    header := scanner.Text()
    if header != "P3" {
        return nil, fmt.Errorf("Unsupported PPM format: %s", header)
    }

    // Read image dimensions and max color value
    scanner.Scan()
    dimensions := strings.Fields(scanner.Text())
    width, _ := strconv.Atoi(dimensions[0])
    height, _ := strconv.Atoi(dimensions[1])

    scanner.Scan()
    maxColorValue, _ := strconv.Atoi(scanner.Text())

    // Read pixel data
    pixels := make([][][3]byte, height)
    for i := 0; i < height; i++ {
        scanner.Scan()
        row := strings.Fields(scanner.Text())
        pixels[i] = make([][3]byte, width)
        
        fmt.Println(len(row))

        for j := 0; j < width; j++ {
            r, _ := strconv.Atoi(row[j*3])
            g, _ := strconv.Atoi(row[j*3+1])
            b, _ := strconv.Atoi(row[j*3+2])
            pixels[i][j] = [3]byte{byte(r), byte(g), byte(b)}
        }
    }

    return &PPMImage{width, height, maxColorValue, pixels}, nil
}

func flipHorizontal(image *PPMImage) {
    for i := 0; i < image.Height; i++ {
        // Find the middle of the row
        mid := image.Width / 2
        for j := 0; j < mid; j++ {
            // Swap pixels from start and end of the row
            image.Pixels[i][j], image.Pixels[i][image.Width-1-j] = image.Pixels[i][image.Width-1-j], image.Pixels[i][j]
        }
    }
}

func writePPMFile(filename string, image *PPMImage) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := bufio.NewWriter(file)

    // Write header
    writer.WriteString("P3\n")
    fmt.Fprintf(writer, "%d %d\n", image.Width, image.Height)
    fmt.Fprintf(writer, "%d\n", image.MaxColorValue)

    // Write pixel data
    for i := 0; i < image.Height; i++ {
        for j := 0; j < image.Width; j++ {
            fmt.Fprintf(writer, "%d %d %d ", image.Pixels[i][j][0], image.Pixels[i][j][1], image.Pixels[i][j][2])
        }
        writer.WriteString("\n")
    }

    writer.Flush()
    return nil
}




func grayscale(image *PPMImage) {
    for i := 0; i < image.Height; i++ {
        for j := 0; j < image.Width; j++ {
            // Calculate grayscale value as average of R, G, and B channels
            gray := (int(image.Pixels[i][j][0]) + int(image.Pixels[i][j][1]) + int(image.Pixels[i][j][2])) / 3
            // Set R, G, and B channels to grayscale value
            image.Pixels[i][j][0] = byte(gray)
            image.Pixels[i][j][1] = byte(gray)
            image.Pixels[i][j][2] = byte(gray)
        }
    }
}



func flipVertical(image *PPMImage) {
    // Find the middle row
    mid := image.Height / 2
    for i := 0; i < mid; i++ {
        // Swap rows from top and bottom of the image
        image.Pixels[i], image.Pixels[image.Height-1-i] = image.Pixels[image.Height-1-i], image.Pixels[i]
    }
}









func invertPixels(image *PPMImage) {
    for i := 0; i < image.Height; i++ {
        for j := 0; j < image.Width; j++ {
            image.Pixels[i][j][0] = byte(image.MaxColorValue) - image.Pixels[i][j][0]
            image.Pixels[i][j][1] = byte(image.MaxColorValue) - image.Pixels[i][j][1]
            image.Pixels[i][j][2] = byte(image.MaxColorValue) - image.Pixels[i][j][2]
        }
    }
}



func flattenColor(image *PPMImage, color int) {
    if color < 0 || color > 2 {
        fmt.Println("Invalid color index. It must be 0 (for Red), 1 (for Green), or 2 (for Blue).")
        return
    }

    for i := 0; i < image.Height; i++ {
        for j := 0; j < image.Width; j++ {
            // Set the specified color channel to zero
            image.Pixels[i][j][color] = 0
        }
    }
}


func extremeContrast(image *PPMImage) {
    mid := image.MaxColorValue / 2

    for i := 0; i < image.Height; i++ {
        for j := 0; j < image.Width; j++ {
            // For each color channel, if the pixel value is below the midpoint, set it to 0; otherwise, set it to MaxColorValue
            for k := 0; k < 3; k++ {
                if int(image.Pixels[i][j][k]) < mid {
                    image.Pixels[i][j][k] = 0
                } else {
                    image.Pixels[i][j][k] = byte(image.MaxColorValue)
                }
            }
        }
    }
}




func main() {
    ppmFile := "./files/small.ppm"
    ppmImage, err := readPPMFile(ppmFile)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }


    fmt.Printf("PPM Image: Width=%d, Height=%d, MaxColorValue=%d\n", ppmImage.Width, ppmImage.Height, ppmImage.MaxColorValue)

    fmt.Println(ppmImage.Pixels)
    grayscale(ppmImage)
    fmt.Println("grayscalled: ", ppmImage.Pixels)

    writePPMFile("grayscale.ppm", ppmImage)

}

