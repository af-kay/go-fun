package interfaces

import (
	"encoding/binary"
	"io"
	"math/rand"
)

type generator struct {
	rnd rand.Source // Генератор случайных чисел. Вообще rand.Rand уже реализует интерфейс io.Reader, но для примера мы реализуем его самостоятельно.
}

// New — обратите внимание, что мы возвращаем generator, присвоенный интерфейсу io.Reader, сама структура generator неэкспортируемая.
// Мы скрыли внутри пакета все детали.
func New(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

// Read — реализация io.Reader
func (g *generator) Read(bytes []byte) (n int, err error) { // error — это тип ошибки, подробнее мы рассмотрим его в следующем разделе.
	rndInt := g.rnd.Int63()
	binary.LittleEndian.PutUint64(bytes, uint64(rndInt))

	return 64 / 8, nil
}
