package main

import "fmt"

// 0xF0000000: Suit bits
// 0x0000FFFF: Face value bits
// 0x0FFF0000: Unused bits

const (
    Suit uint32 = 0xF0000
    Spades uint32 = 0x80000
    Clubs uint32 = 0x40000
    Diamonds uint32 = 0x20000
    Hearts uint32 = 0x10000

    Value uint32 = 0x0FFFF
    King uint32 = 0x1000
    Queen uint32 = 0x0800
    Jack uint32 = 0x0400
    Ten uint32 = 0x0200
    Nine uint32 = 0x0100
    Eight uint32 = 0x0080
    Seven uint32 = 0x0040
    Six uint32 = 0x0020
    Five uint32 = 0x0010
    Four uint32 = 0x0008
    Three uint32 = 0x0004
    Two uint32 = 0x0002
    Ace uint32 = 0x2001 // Both high and low.

    AceSpades uint32 = Ace | Spades
    KingSpades uint32 = King | Spades
    QueenSpades uint32 = Queen | Spades

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

    RoyalStraightFlush // 4 values
    StraightFlush
    FourOfAKind // 13 values
    FullHouse // 156 (13 * 12) values
    Flush // 13 values
    Straight //
    ThreeOfAKind // 13 values
    TwoPair // 156 (13 * 12) values
    OnePair // 13 values
    HighCard // 13 values

)

var strings = map[uint32]string{
    AceSpades: "As",
    KingSpades: "Ks",
    QueenSpades: "Qs",
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
    fmt.Println("Swapping", a, "&", b)

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
        return Value & cards[0]
    }

    // edge case: five high, with an Ace being (mis)sorted by it's high value
    if (Value & cards[1] >> 4) & (Value & cards[2] >> 3) & (Value & cards[3] >> 2) & (Value & cards[4] >> 1) & Value & cards[0] != 0 {
        return Value & cards[1]
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
    return values[Value & card]
}

func evaluate(cards [5]uint32) {
    fmt.Println("Cards:", cards_to_string(cards))

    // Sort cards by face value from largest to smallest.
    cards = sort(cards)
    fmt.Println("Sorted:", cards_to_string(cards))

    // Check for flush.
    flush_suit := flush(cards)
    fmt.Println("Flush suit:", suit_to_string(flush_suit))

    // Check for straight.
    straight_high_card := straight(cards)
    fmt.Println("Straight high card:", value_to_string(straight_high_card))

    if flush_suit != 0 && straight_high_card != 0 && straight_high_card == Ace {
        fmt.Println("Royal straight flush,", value_to_string(straight_high_card), "high")
        return
    }

    if flush_suit != 0 && straight_high_card != 0 {
        fmt.Println("Straight flush,", value_to_string(straight_high_card), "high")
        return
    }

    // Four of a kind? return
    // Check for three of a kind.
    // Check for pair.
    // Full house? return
    // Flush? return

    if flush_suit != 0 && straight_high_card != 0 {
        fmt.Println("Flush,", value_to_string(cards[0]), "high")
        return
    }

    if straight_high_card != 0 {
        fmt.Println("Straight,", value_to_string(straight_high_card), "high")
        return
    }

    // Three of a kind? return
    // Two pair? return
    // One pair? return
    // return high card
    fmt.Println("High card,", value_to_string(cards[0]), "high")
}

func main() {
    fmt.Println("Straight, king high")
    evaluate([5]uint32{NineClubs, TenClubs, KingSpades, JackClubs, QueenSpades})
    fmt.Println()

    fmt.Println("Straight, ace high")
    evaluate([5]uint32{AceSpades, TenClubs, KingSpades, JackClubs, QueenSpades})
    fmt.Println()

    fmt.Println("Straight, five high")
    evaluate([5]uint32{AceSpades, TwoClubs, ThreeClubs, FourClubs, FiveClubs})
    fmt.Println()

    fmt.Println("Straight flush, six high")
    evaluate([5]uint32{SixClubs, TwoClubs, ThreeClubs, FourClubs, FiveClubs})
    fmt.Println()
}
