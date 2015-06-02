package handeval

// 0x0000000F: Suit bits
// 0x000FFFF0: Face value bits
// 0xFFF00000: Unused bits

const (
    Suit uint32 =  0x00000F // Suit mask.
    Face uint32 =  0x0000F0 // Face value bit mask.
    Value uint32 = 0xFFFF00 // Bit value bit mask.

    Spades uint32 = 0x8
    Clubs uint32 = 0x4
    Diamonds uint32 = 0x2
    Hearts uint32 = 0x1

    King uint32 = 0x1000d0
    Queen uint32 = 0x0800c0
    Jack uint32 = 0x0400b0
    Ten uint32 = 0x0200a0
    Nine uint32 = 0x010090
    Eight uint32 = 0x008080
    Seven uint32 = 0x004070
    Six uint32 = 0x002060
    Five uint32 = 0x001050
    Four uint32 = 0x000840
    Three uint32 = 0x000430
    Two uint32 = 0x000220
    Ace uint32 = 0x2001f0 // Both high and low.
)

func flush(cards [5]uint32) uint32 {
    // if all suits are the same, return the suit
    return Suit & cards[0] & cards[1] & cards[2] & cards[3] & cards[4]
}

func swap(cards [5]uint32, a int, b int) [5]uint32 {
    tmp := cards[a]
    cards[a] = cards[b]
    cards[b] = tmp
    return cards
}

// Reference: http://stackoverflow.com/q/1534748/
func Sort(cards [5]uint32) [5]uint32 {
    if (cards[0] < cards[1]) {
        cards = swap(cards, 0, 1)
    }
    if (cards[2] < cards[3]) {
        cards = swap(cards, 2, 3)
    }
    if (cards[0] < cards[2]) {
        cards = swap(cards, 0, 2)
        cards = swap(cards, 1, 3)
    }

    if (cards[4] < cards[2]) {
        if (cards[4] < cards[3]) {
            if (cards[1] < cards[3]) {
                if (cards[1] < cards[4]) {
                    return [5]uint32{cards[0], cards[2], cards[3], cards[4], cards[1]}
                } else {
                    return [5]uint32{cards[0], cards[2], cards[3], cards[1], cards[4]}
                }
            } else {
                if (cards[1] > cards[2]) {
                    return [5]uint32{cards[0], cards[1], cards[2], cards[3], cards[4]}
                } else {
                    return [5]uint32{cards[0], cards[2], cards[1], cards[3], cards[4]}
                }
            }
        } else {
            if (cards[1] < cards[4]) {
                if (cards[1] < cards[3]) {
                    return [5]uint32{cards[0], cards[2], cards[4], cards[3], cards[1]}
                } else {
                    return [5]uint32{cards[0], cards[2], cards[4], cards[1], cards[3]}
                }
            } else {
                if (cards[1] > cards[2]) {
                    return [5]uint32{cards[0], cards[1], cards[2], cards[4], cards[3]}
                } else {
                    return [5]uint32{cards[0], cards[2], cards[1], cards[4], cards[3]}
                }
            }
        }
    } else {
        if (cards[4] > cards[0]) {
            if (cards[1] < cards[2]) {
                if (cards[1] < cards[3]) {
                    return [5]uint32{cards[4], cards[0], cards[2], cards[3], cards[1]}
                } else {
                    return [5]uint32{cards[4], cards[0], cards[2], cards[1], cards[3]}
                }
            } else {
                return [5]uint32{cards[4], cards[0], cards[1], cards[2], cards[3]}
            }
        } else {
            if (cards[1] < cards[2]) {
                if (cards[1] < cards[3]) {
                    return [5]uint32{cards[0], cards[4], cards[2], cards[3], cards[1]}
                } else {
                    return [5]uint32{cards[0], cards[4], cards[2], cards[1], cards[3]}
                }
            } else {
                if (cards[1] > cards[4]) {
                    return [5]uint32{cards[0], cards[1], cards[4], cards[2], cards[3]}
                } else {
                    return [5]uint32{cards[0], cards[4], cards[1], cards[2], cards[3]}
                }
            }
        }
    }
    return cards
}

func straight(cards [5]uint32) uint32 {
    if (Value & cards[0] >> 4) & (Value & cards[1] >> 3) & (Value & cards[2] >> 2) & (Value & cards[3] >> 1) & (Value & cards[4]) != 0 {
        return cards[0]
    }

    // edge case: five high, with an Ace being (mis)sorted by it's high value
    if (Value & cards[1] >> 4) & (Value & cards[2] >> 3) & (Value & cards[3] >> 2) & (Value & cards[4] >> 1) & Value & cards[0] != 0 {
        return cards[1]
    }

    return 0
}

func four_of_a_kind(cards [5]uint32) uint32 {
    if (Value & cards[0]) & (Value & cards[1]) & (Value & cards[2]) & (Value & cards[3]) != 0 {
        return cards[1]
    }

    if (Value & cards[1]) & (Value & cards[2]) & (Value & cards[3]) & (Value & cards[4]) != 0 {
        return cards[1]
    }

    return 0
}

func three_of_a_kind(cards [5]uint32) uint32 {
    if (Value & cards[0]) & (Value & cards[1]) & (Value & cards[2]) != 0 {
        return cards[2]
    }

    if (Value & cards[1]) & (Value & cards[2]) & (Value & cards[3]) != 0 {
        return cards[2]
    }

    if (Value & cards[2]) & (Value & cards[3]) & (Value & cards[4]) != 0 {
        return cards[2]
    }

    return 0
}

func high_pair(cards [5]uint32) uint32 {
    if (Value & cards[0]) & (Value & cards[1]) != 0 || (Value & cards[1]) & (Value & cards[2]) != 0 {
        return cards[1]
    }

    return 0
}

func low_pair(cards [5]uint32) uint32 {
    if (Value & cards[2]) & (Value & cards[3]) != 0 || (Value & cards[3]) & (Value & cards[4]) != 0 {
        return cards[3]
    }

    return 0
}

func Evaluate(cards [5]uint32) uint32 {
    // Sort cards by face value from largest to smallest.
    cards = Sort(cards)

    // Check for flush.
    flush_suit := flush(cards)

    // Check for straight.
    straight_high_card := straight(cards)

    if flush_suit != 0 && straight_high_card != 0 && straight_high_card == Ace {
        // Royal flush
        // fmt.Println("Royal straight flush, " + value_to_string(straight_high_card) + " high")
        // 0xFF00 0000
        return (Face & straight_high_card | flush_suit) << 24
    }

    if flush_suit != 0 && straight_high_card != 0 {
        // Straight flush
        // 0xFF00 0000
        // fmt.Println("Straight flush, " + value_to_string(straight_high_card) + " high")
        return (Face & straight_high_card | flush_suit) << 24
    }

    four_of_a_kind_card := four_of_a_kind(cards)
    if four_of_a_kind_card != 0 {
        // Four of a kind
        // 0xF000 0000
        // fmt.Println("Four of a kind, " + value_to_string(four_of_a_kind_card) + "s")
        return (Face & four_of_a_kind_card) << 24
    }

    // Check for three of a kind.
    three_of_a_kind_card := three_of_a_kind(cards)

    // Check for pair.
    high_pair_card := high_pair(cards)
    low_pair_card := low_pair(cards)

    if three_of_a_kind_card != 0 && high_pair_card != 0 && Value & three_of_a_kind_card != Value & high_pair_card {
        // Full house
        // 0xFF0 0000
        // fmt.Println("Full house, " + value_to_string(three_of_a_kind_card) + "s full of " + value_to_string(high_pair_card) + "s")
        return (Face & three_of_a_kind_card) << 20 | (Face & high_pair_card) << 16
    }
    if three_of_a_kind_card != 0 && low_pair_card != 0 && Value & three_of_a_kind_card != Value & low_pair_card {
        // Full house
        // 0xFF0 0000
        // fmt.Println("Full house, " + value_to_string(three_of_a_kind_card) + "s full of " + value_to_string(low_pair_card) + "s")
        return (Face & three_of_a_kind_card) << 20 | (Face & low_pair_card) << 16
    }

    if flush_suit != 0 {
        // Flush
        // 0xF0 0000
        // fmt.Println("Flush, " + value_to_string(cards[0]) + " high")
        return (Face & cards[0]) << 16
    }

    if straight_high_card != 0 {
        // Straight
        // 0xF 0000
        // fmt.Println("Straight, " + value_to_string(straight_high_card) + " high")
        return (Face & straight_high_card) << 12
    }

    if three_of_a_kind_card != 0 {
        // Three of a kind
        // 0xF000
        // fmt.Println("Three of a kind, " + value_to_string(three_of_a_kind_card) + "s")
        return (Face & three_of_a_kind_card) << 8
    }

    if high_pair_card != 0 && low_pair_card != 0 {
        // Two pair
        // 0xFF0
        // fmt.Println("Two pair, " + value_to_string(high_pair_card) + "s and " + value_to_string(low_pair_card) + "s")
        return ((Face & high_pair_card) << 4) | (Face & low_pair_card)
    }

    if high_pair_card != 0 {
        // Pair
        // 0xF0
        // fmt.Println("Pair of " + value_to_string(high_pair_card) + "s")
        return Face & high_pair_card
    }
    if low_pair_card != 0 {
        // Pair
        // 0xF0
        // fmt.Println("Pair of " + value_to_string(low_pair_card) + "s")
        return Face & low_pair_card
    }

    // High card
    // 0xF
    // fmt.Println("High card " + value_to_string(cards[0]))
    return (Face & cards[0]) >> 4
}

var face_values = map[uint32]string{
    Face & King: "king",
    Face & Queen: "queen",
    Face & Jack: "jack",
    Face & Ten: "ten",
    Face & Nine: "nine",
    Face & Eight: "eight",
    Face & Seven: "seven",
    Face & Six: "six",
    Face & Five: "five",
    Face & Four: "four",
    Face & Three: "three",
    Face & Two: "two",
    Face & Ace: "ace",
    0: "none",
}

var face_values_plural = map[uint32]string{
    Face & King: "kings",
    Face & Queen: "queens",
    Face & Jack: "jacks",
    Face & Ten: "tens",
    Face & Nine: "nines",
    Face & Eight: "eights",
    Face & Seven: "sevens",
    Face & Six: "sixes",
    Face & Five: "fives",
    Face & Four: "fours",
    Face & Three: "threes",
    Face & Two: "twos",
    Face & Ace: "aces",
    0: "none",
}

func ToString(evaluation uint32) string {
    if 0x0000000F & evaluation != 0 {
        return "High card, " + face_values[evaluation << 4]
    }
    if 0x000000F0 & evaluation != 0 && 0xFFFFFF0F & evaluation == 0 {
        // Face & pair_card
        return "Pair of " + face_values_plural[evaluation]
    }
    if 0x00000FF0 & evaluation != 0 {
        // ((Face & high_pair_card) << 4) | (Face & low_pair_card)
        return "Two pair, " + face_values_plural[0x00000F00 & evaluation >> 4] + " and " + face_values_plural[0x000000F0 & evaluation]
    }
    if 0x0000F000 & evaluation != 0 {
        // (Face & three_of_a_kind_card) << 8
        return "Three of a kind, " + face_values_plural[evaluation >> 8]
    }
    if 0x000F0000 & evaluation != 0 {
        // (Face & straight_high_card) << 12
        return "Straight, " + face_values[evaluation >> 12] + " high"
    }
    if 0x00F00000 & evaluation != 0 && 0xFF0FFFFF & evaluation == 0 {
        // (Face & cards[0]) << 16
        return "Flush, " + face_values[evaluation >> 16] + " high"
    }
    if 0x0FF00000 & evaluation != 0 && 0xF00FFFFF & evaluation == 0 {
        // (Face & three_of_a_kind_card) << 20 | (Face & high_pair_card) << 16
        return "Full house, " + face_values_plural[0x0F000000 & evaluation >> 20] + " full of " + face_values_plural[0x00F00000 & evaluation >> 16]
    }
    if 0xF0000000 & evaluation != 0 && 0x0FFFFFFF & evaluation == 0 {
        // (Face & four_of_a_kind_card) << 24
        return "Four of a kind, " + face_values_plural[evaluation >> 24]
    }
    if 0xFF000000 & evaluation != 0 {
        // (Face & straight_high_card | flush_suit) << 24
        if 0xF0000000 & evaluation >> 24 == Face & Ace {
            return "Royal flush"
        } else {
            return "Straight flush, " + face_values[0xF0000000 & evaluation >> 24] + " high"
        }
    }

    panic("Unable to decode evaluation: " + string(evaluation))
}
