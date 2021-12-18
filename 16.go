package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Packet struct {
	version    int
	typeID     int
	subPackets []Packet
	value      int
}

func Sixteen() {
	file, err := os.Open("./input-16.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	hexMap := map[rune][]int{
		'0': {0, 0, 0, 0},
		'1': {0, 0, 0, 1},
		'2': {0, 0, 1, 0},
		'3': {0, 0, 1, 1},
		'4': {0, 1, 0, 0},
		'5': {0, 1, 0, 1},
		'6': {0, 1, 1, 0},
		'7': {0, 1, 1, 1},
		'8': {1, 0, 0, 0},
		'9': {1, 0, 0, 1},
		'A': {1, 0, 1, 0},
		'B': {1, 0, 1, 1},
		'C': {1, 1, 0, 0},
		'D': {1, 1, 0, 1},
		'E': {1, 1, 1, 0},
		'F': {1, 1, 1, 1},
	}
	bits := make([]int, 0)
	for _, char := range line {
		bits = append(bits, hexMap[char]...)
	}

	bitIndex := 0
	packet := parsePackage(bits, &bitIndex)

	fmt.Printf("Evaluated %d\n", evaluatePacket(packet))
}

func evaluatePacket(packet Packet) int {
	if packet.typeID == 0 {
		sum := 0
		for _, subPacket := range packet.subPackets {
			sum += evaluatePacket(subPacket)
		}
		return sum
	}
	if packet.typeID == 1 {
		product := 1
		for _, subPacket := range packet.subPackets {
			product *= evaluatePacket(subPacket)
		}
		return product
	}
	if packet.typeID == 2 {
		minimum := math.MaxInt
		for _, subPacket := range packet.subPackets {
			value := evaluatePacket(subPacket)
			if value < minimum {
				minimum = value
			}
		}
		return minimum
	}
	if packet.typeID == 3 {
		maximum := math.MinInt
		for _, subPacket := range packet.subPackets {
			value := evaluatePacket(subPacket)
			if value > maximum {
				maximum = value
			}
		}
		return maximum
	}
	if packet.typeID == 5 {
		a := evaluatePacket(packet.subPackets[0])
		b := evaluatePacket(packet.subPackets[1])
		if a > b {
			return 1
		}
		return 0
	}
	if packet.typeID == 6 {
		a := evaluatePacket(packet.subPackets[0])
		b := evaluatePacket(packet.subPackets[1])
		if a < b {
			return 1
		}
		return 0
	}
	if packet.typeID == 7 {
		a := evaluatePacket(packet.subPackets[0])
		b := evaluatePacket(packet.subPackets[1])
		if a == b {
			return 1
		}
		return 0
	}
	return packet.value
}

func sumVersionNumbers(packet Packet) int {
	versionSum := packet.version
	for _, subPacket := range packet.subPackets {
		versionSum += sumVersionNumbers(subPacket)
	}
	return versionSum
}

func parsePackage(bits []int, bitIndex *int) Packet {

	packet := Packet{}
	version := bitsToDecimal(bits[*bitIndex : *bitIndex+3])
	typeID := bitsToDecimal(bits[*bitIndex+3 : *bitIndex+6])
	packet.version = version
	packet.typeID = typeID

	*bitIndex += 6
	if typeID == 4 {
		literalValue := readLiteralValue(bits, bitIndex)
		packet.value = literalValue
	} else {
		lengthID := bits[*bitIndex]

		subPackets := make([]Packet, 0)
		if lengthID == 0 {
			length := bitsToDecimal(bits[*bitIndex+1 : *bitIndex+16])
			packetParseEnd := *bitIndex + 16 + length
			*bitIndex += 16
			for *bitIndex < packetParseEnd {
				subPacket := parsePackage(bits, bitIndex)
				subPackets = append(subPackets, subPacket)
			}

		}
		if lengthID == 1 {
			length := bitsToDecimal(bits[*bitIndex+1 : *bitIndex+12])
			*bitIndex += 12
			for i := 0; i < length; i++ {
				subPacket := parsePackage(bits, bitIndex)
				subPackets = append(subPackets, subPacket)
			}
		}
		packet.subPackets = subPackets
	}
	return packet
}

func readLiteralValue(bits []int, bitIndex *int) int {
	literalValue := make([]int, 0)
	for true {
		group := bits[*bitIndex : *bitIndex+5]

		literalValue = append(literalValue, group[1:]...)

		*bitIndex += 5
		if group[0] == 0 {
			break
		}
	}
	return bitsToDecimal(literalValue)

}

func bitsToDecimal(bits []int) int {
	sum := 0
	pow := 2 << (len(bits) - 2)
	for _, bit := range bits {
		sum += bit * pow
		pow = pow >> 1
	}
	return sum
}
