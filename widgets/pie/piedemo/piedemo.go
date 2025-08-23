package main

import (
	"context"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/pie"
	
)

func main() {
	t, err := tcell.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())

	pieWidget, err := pie.New()
	if err != nil {
		panic(err)
	}

	// Set initial values for the pie chart.
	values := []int{30, 20, 50}
	colors := []cell.Color{cell.ColorRed, cell.ColorGreen, cell.ColorBlue}
	if err := pieWidget.Values(values, pie.WithColors(colors)); err != nil {
		panic(err)
	}

	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle("PRESS Q TO QUIT"),
		container.PlaceWidget(pieWidget),
	)
	if err != nil {
		panic(err)
	}

	// Quitter function to handle 'q' key press.
	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	// Update the pie chart values periodically.
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				// Update values dynamically.
				values = []int{values[2], values[0], values[1]}
				if err := pieWidget.Values(values, pie.WithColors(colors)); err != nil {
					panic(err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(1*time.Second)); err != nil {
		panic(err)
	}
}