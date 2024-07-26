package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
    tests := []struct {
        input    int
        expected string
    }{
        {15, "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n14\nFizzBuzz\n"},
				{10, "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n"},
        {5, "1\n2\nFizz\n4\nBuzz\n"},
        {3, "1\n2\nFizz\n"},
        {1, "1\n"},
    }

    for _, test := range tests {
        t.Run(fmt.Sprintf("input_%d", test.input), func(t *testing.T) {
            // Redirect stdout to capture the output
            old := os.Stdout
            r, w, _ := os.Pipe()
            os.Stdout = w

            fizzBuzz(test.input)

            // Restore stdout
            w.Close()
            os.Stdout = old

            var buf bytes.Buffer
            io.Copy(&buf, r)
            output := buf.String()

            if output != test.expected {
                t.Errorf("For input %d, expected:\n%s\nbut got:\n%s", test.input, test.expected, output)
            }
        })
    }
}
