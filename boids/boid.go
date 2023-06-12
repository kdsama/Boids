package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var (
	rwLock = sync.RWMutex{}
	// wLock  = sync.WM
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

// create all the dots , and put them on the map
func createBoid(bid int) {
	b := Boid{position: Vector2D{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
		// velocity between 1 and -1
		velocity: Vector2D{x: (rand.Float64() * 2) - 1, y: (rand.Float64() * 2) - 1},
		id:       bid,
	}
	boids[bid] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = bid
	go b.start()
}

func (b *Boid) start() {

	for {
		b.moveOne()
		time.Sleep(1 * time.Millisecond)
	}
}

func (b *Boid) moveOne() {
	acceleration := b.newAcceleration()
	rwLock.Lock()
	b.velocity = b.velocity.Add(acceleration)
	b.velocity = b.velocity.Limit(-1, 1)
	fmt.Println(b.position)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)

	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	next := b.position.Add(b.velocity)
	if next.x >= screenWidth || next.x < 0 {
		b.velocity = Vector2D{-b.velocity.x, b.velocity.y}
	}
	if next.y >= screenHeight || next.y < 0 {
		b.velocity = Vector2D{-b.velocity.x, -b.velocity.y}
	}
	rwLock.Unlock()
}

// alignment cohesion and separation are the key things we use in this boid simulation

func (b *Boid) newAcceleration() Vector2D {

	// whats the size of box we gonna consider here ?

	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)

	avgVel, avgPos, separation := Vector2D{0, 0}, Vector2D{0, 0}, Vector2D{0, 0}
	count := 0.0
	rwLock.RLock()
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if others := boidMap[int(i)][int(j)]; others != -1 && others != b.id {
				if dist := boids[others].position.Distance(b.position); dist < viewRadius {
					count++
					avgVel = avgVel.Add(boids[others].velocity)
					avgPos = avgPos.Add(boids[others].position)
					b.position = b.position.Subtract(boids[others].position)
					x := b.position.DivisionV(dist)
					separation = separation.Add(x)
				}
			}
		}
	}
	rwLock.RUnlock()
	acc := Vector2D{b.borderBounce(b.position.x, screenWidth), b.borderBounce(b.position.y, screenHeight)}
	if count > 0 {
		avgVel, avgPos = avgVel.DivisionV(count), avgPos.DivisionV(count)
		accAlignment := avgVel.Subtract(b.velocity)
		accAlignment = accAlignment.MultiplyV(adjRate)
		accCohesion := avgVel.Subtract(b.position)
		accCohesion = accCohesion.MultiplyV(adjRate)
		accelSep := separation.MultiplyV(adjRate)
		acc = acc.Add(accAlignment)
		acc = acc.Add(accCohesion)
		acc = acc.Add(accelSep)

	}
	return acc
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {

	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}
