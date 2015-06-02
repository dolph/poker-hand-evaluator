package main

import "fmt"

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

    AceSpades uint32 = Ace | Spades
    KingSpades uint32 = King | Spades
    QueenSpades uint32 = Queen | Spades
    FiveSpades uint32 = Five | Spades

    KingClubs uint32 = King | Clubs
    QueenClubs uint32 = Queen | Clubs
    JackClubs uint32 = Jack | Clubs
    TenClubs uint32 = Ten | Clubs
    NineClubs uint32 = Nine | Clubs
    EightClubs uint32 = Eight | Clubs
    SevenClubs uint32 = Seven | Clubs
    SixClubs uint32 = Six | Clubs
    FiveClubs uint32 = Five | Clubs
    FourClubs uint32 = Four | Clubs
    ThreeClubs uint32 = Three | Clubs
    TwoClubs uint32 = Two | Clubs
    AceClubs uint32 = Ace | Clubs

    AceDiamonds uint32 = Ace | Diamonds

    AceHearts uint32 = Ace | Hearts
)

var strings = map[uint32]string{
    AceSpades: "As",
    KingSpades: "Ks",
    QueenSpades: "Qs",
    FiveSpades: "5s",

    KingClubs: "Kc",
    QueenClubs: "Qc",
    JackClubs: "Jc",
    TenClubs: "Tc",
    NineClubs: "9c",
    EightClubs: "8c",
    SevenClubs: "7c",
    SixClubs: "6c",
    FiveClubs: "5c",
    FourClubs: "4c",
    ThreeClubs: "3c",
    TwoClubs: "2c",
    AceClubs: "Ac",

    AceDiamonds: "Ad",

    AceHearts: "Ah",
}

var suits = map[uint32]string{
    Spades: "spades",
    Clubs: "clubs",
    Diamonds: "diamonds",
    Hearts: "hearts",
    0: "none",
}

var values = map[uint32]string{
    King: "king",
    Queen: "queen",
    Jack: "jack",
    Ten: "ten",
    Nine: "nine",
    Eight: "eight",
    Seven: "seven",
    Six: "six",
    Five: "five",
    Four: "four",
    Three: "three",
    Two: "two",
    Ace: "ace",
    0: "none",
}

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
func sort(cards [5]uint32) [5]uint32 {
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

func card_to_string(card uint32) string {
    return strings[card]
}

func cards_to_string(cards [5]uint32) string {
    s := ""
    for i := 0; i < len(cards); i++ {
        s += strings[cards[i]]
        s += " "
    }
    return s
}

func suit_to_string(card uint32) string {
    return suits[Suit & card]
}

func value_to_string(card uint32) string {
    return values[(Face | Value) & card]
}

func evaluate(cards [5]uint32) string {
    fmt.Println("Cards:", cards_to_string(cards))

    // Sort cards by face value from largest to smallest.
    cards = sort(cards)
    // fmt.Println("Sorted:", cards_to_string(cards))

    // Check for flush.
    flush_suit := flush(cards)

    // Check for straight.
    straight_high_card := straight(cards)

    // 0xFF00 0000
    if flush_suit != 0 && straight_high_card != 0 && straight_high_card == Ace {
        fmt.Println((Face & straight_high_card | flush_suit) << 24)
        return "Royal straight flush, " + value_to_string(straight_high_card) + " high"
    }

    // 0xFF00 0000
    if flush_suit != 0 && straight_high_card != 0 {
        fmt.Println((Face & straight_high_card | flush_suit) << 24)
        return "Straight flush, " + value_to_string(straight_high_card) + " high"
    }

    four_of_a_kind_card := four_of_a_kind(cards)
    // 0xF000 0000
    if four_of_a_kind_card != 0 {
        fmt.Println((Face & four_of_a_kind_card) << 24)
        return "Four of a kind, " + value_to_string(four_of_a_kind_card) + "s"
    }

    // Check for three of a kind.
    three_of_a_kind_card := three_of_a_kind(cards)

    // Check for pair.
    high_pair_card := high_pair(cards)
    low_pair_card := low_pair(cards)

    // 0xFF0 0000
    if three_of_a_kind_card != 0 && high_pair_card != 0 && Value & three_of_a_kind_card != Value & high_pair_card {
        fmt.Println((Face & three_of_a_kind_card) << 20 | (Face & high_pair_card) << 16)
        return "Full house, " + value_to_string(three_of_a_kind_card) + "s full of " + value_to_string(high_pair_card) + "s"
    }
    if three_of_a_kind_card != 0 && low_pair_card != 0 && Value & three_of_a_kind_card != Value & low_pair_card {
        fmt.Println((Face & three_of_a_kind_card) << 20 | (Face & low_pair_card) << 16)
        return "Full house, " + value_to_string(three_of_a_kind_card) + "s full of " + value_to_string(low_pair_card) + "s"
    }

    // 0xF0 0000
    if flush_suit != 0 {
        fmt.Println((Face & cards[0]) << 16)
        return "Flush, " + value_to_string(cards[0]) + " high"
    }

    // 0xF 0000
    if straight_high_card != 0 {
        fmt.Println((Face & straight_high_card) << 12)
        return "Straight, " + value_to_string(straight_high_card) + " high"
    }

    // 0xF000
    if three_of_a_kind_card != 0 {
        fmt.Println((Face & three_of_a_kind_card) << 8)
        return "Three of a kind, " + value_to_string(three_of_a_kind_card) + "s"
    }

    // 0xFF0
    if high_pair_card != 0 && low_pair_card != 0 {
        fmt.Println(((Face & high_pair_card) << 4) | (Face & low_pair_card))
        return "Two pair, " + value_to_string(high_pair_card) + "s and " + value_to_string(low_pair_card) + "s"
    }

    // 0xF0
    if high_pair_card != 0 {
        fmt.Println((Face & high_pair_card))
        return "Pair of " + value_to_string(high_pair_card) + "s"
    }
    if low_pair_card != 0 {
        fmt.Println((Face & low_pair_card))
        return "Pair of " + value_to_string(low_pair_card) + "s"
    }

    // 0xF
    fmt.Println((Face & cards[0]) >> 4)
    return "High card " + value_to_string(cards[0])
}

func main() {
    fmt.Println(evaluate([5]uint32{AceClubs, AceSpades, AceDiamonds, QueenSpades, AceHearts}))
    fmt.Println("Four of a kind, aces")
    fmt.Println()

    fmt.Println(evaluate([5]uint32{SixClubs, TwoClubs, ThreeClubs, FourClubs, FiveClubs}))
    fmt.Println("Straight flush, six high")
    fmt.Println()

    fmt.Println(evaluate([5]uint32{AceClubs, TwoClubs, ThreeClubs, FourClubs, FiveClubs}))
    fmt.Println("Straight flush, five high")
    fmt.Println()

    fmt.Println(evaluate([5]uint32{AceClubs, AceSpades, AceDiamonds, KingClubs, KingSpades}))
    fmt.Println("Full house, aces full of kings")
    fmt.Println()

    fmt.Println(evaluate([5]uint32{AceSpades, TenClubs, KingSpades, JackClubs, QueenSpades}))
    fmt.Println("Straight, ace high")
    fmt.Println()

    fmt.Println(evaluate([5]uint32{NineClubs, TenClubs, KingSpades, JackClubs, QueenSpades}))
    fmt.Println("Straight, king high")
    fmt.Println()

    fmt.Println(evaluate([5]uint32{AceSpades, TwoClubs, ThreeClubs, FourClubs, FiveClubs}))
    fmt.Println("Straight, five high")
    fmt.Println()

    fmt.Println(evaluate([5]uint32{KingClubs, AceClubs, AceSpades, QueenSpades, KingSpades}))
    fmt.Println("Two pair, aces and kings")
    fmt.Println()

    fmt.Println(evaluate([5]uint32{KingClubs, AceClubs, EightClubs, QueenSpades, AceSpades}))
    fmt.Println("Pair of aces")
    fmt.Println()

    fmt.Println(evaluate([5]uint32{KingClubs, AceHearts, EightClubs, QueenSpades, AceSpades}))
    fmt.Println("Pair of aces")
    fmt.Println()

    fmt.Println(evaluate([5]uint32{KingClubs, AceClubs, EightClubs, QueenSpades, KingSpades}))
    fmt.Println("Pair of kings")
    fmt.Println()

    fmt.Println(evaluate([5]uint32{AceClubs, QueenSpades, TenClubs, EightClubs, SixClubs}))
    fmt.Println("High card ace")
    fmt.Println()

    fmt.Println(evaluate([5]uint32{SevenClubs, FiveSpades, FourClubs, ThreeClubs, TwoClubs}))
    fmt.Println("High card seven")
    fmt.Println()
}
