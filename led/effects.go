package led

import (
)

type Effect interface {
    Update(dt float32)
    Draw(screen DrawBuffer)
    Position() Vec
    Done() bool
}

type Effects struct {
    list []Effect
}

func NewEffects() *Effects {
    e := &Effects {
        list: make([]Effect, 0, 0),
    }
    return e
}

func (e *Effects) Count() int {
    return len(e.list)
}

func (e *Effects) Add(effect Effect) {
    for idx, other := range e.list {
        if other.Position() == effect.Position() {
            e.list[idx] = effect
            return
        }
    }

    e.list = append(e.list, effect)
}

func (e *Effects) Update(dt float32) {
    alive := make([]Effect, 0, len(e.list))
    for _, effect := range e.list {
        effect.Update(dt)
        if !effect.Done() {
            alive = append(alive, effect)
        }
    }
    e.list = alive
}

func (e *Effects) Draw(screen DrawBuffer) {
    screen.Clear(Color{0,0,0,1})
    for _, effect := range e.list {
        effect.Draw(screen)
    }
    screen.Flip()
}

type Spark struct {
    Color     Color
    Lifetime  float32
    Falloff   float32
    Intensity float32

    life      float32
    intensity float32
    position  Vec
}

func NewSpark(pos Vec, color Color, intensity, lifetime, falloff float32) *Spark {
    s := &Spark {
        Color: color,
        Intensity: intensity,
        Lifetime: lifetime,
        Falloff: falloff,

        life: 0.0,
        position: pos,
        intensity: intensity,
    }
    return s
}

func (s *Spark) Done() bool {
    return s.intensity < 0.05
    return s.life > s.Lifetime
}

func (s *Spark) Update(dt float32) {
    s.life += dt
    s.intensity *= 1.0 - dt * s.Falloff
}

func (s *Spark) Draw(screen DrawBuffer) {
    screen.Set(s.Position(), s.Color.Scale(s.intensity))
}

func (s *Spark) Position() Vec {
    return s.position
}
