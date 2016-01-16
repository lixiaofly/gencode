package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGencodeSize(t *testing.T) {
	p := Group{
		Name: "test",
		Members: []Person{
			{
				Name:   "John",
				Age:    21,
				Height: 5.9,
			},
			{
				Name:   "Tom",
				Age:    23,
				Height: 5.8,
			},
			{
				Name:   "Alan",
				Age:    24,
				Height: 6,
			},
		},
	}
	buf := bytes.NewBuffer(nil)
	p.Serialize(buf)
	fmt.Printf("Gencode encoded size: %v\n", len(buf.Bytes()))
}

func BenchmarkGencodeSerialize(b *testing.B) {
	p := Group{
		Name: "test",
		Members: []Person{
			{
				Name:   "John",
				Age:    21,
				Height: 5.9,
			},
			{
				Name:   "Tom",
				Age:    23,
				Height: 5.8,
			},
			{
				Name:   "Alan",
				Age:    24,
				Height: 6,
			},
		},
	}
	buf := bytes.NewBuffer(nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Serialize(buf)
		buf.Reset()
	}
}

func BenchmarkGencodeDeserialize(b *testing.B) {
	p := Group{
		Name: "test",
		Members: []Person{
			{
				Name:   "John",
				Age:    21,
				Height: 5.9,
			},
			{
				Name:   "Tom",
				Age:    23,
				Height: 5.8,
			},
			{
				Name:   "Alan",
				Age:    24,
				Height: 6,
			},
		},
	}
	buf := bytes.NewBuffer(nil)
	p.Serialize(buf)
	rbuf := bytes.NewReader(buf.Bytes())
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Deserialize(rbuf)
		rbuf.Seek(0, 0)
	}
}

func BenchmarkFixedGencodeSerialize(b *testing.B) {
	p := Fixed{
		A: -5,
		B: 6,
		C: 6.7,
		D: 12.65,
	}
	buf := bytes.NewBuffer(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Serialize(buf)
		buf.Reset()
	}
}

func BenchmarkFixedGencodeDeserialize(b *testing.B) {
	p := Fixed{
		A: -5,
		B: 6,
		C: 6.7,
		D: 12.65,
	}
	buf := bytes.NewBuffer(nil)
	p.Serialize(buf)
	rbuf := bytes.NewReader(buf.Bytes())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Deserialize(rbuf)
		rbuf.Seek(0, 0)
	}
}