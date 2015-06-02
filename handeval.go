package main

import "fmt"
import "os"
import "time"

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

    KingSpades uint32 = King | Spades
    QueenSpades uint32 = Queen | Spades
    JackSpades uint32 = Jack | Spades
    TenSpades uint32 = Ten | Spades
    NineSpades uint32 = Nine | Spades
    EightSpades uint32 = Eight | Spades
    SevenSpades uint32 = Seven | Spades
    SixSpades uint32 = Six | Spades
    FiveSpades uint32 = Five | Spades
    FourSpades uint32 = Four | Spades
    ThreeSpades uint32 = Three | Spades
    TwoSpades uint32 = Two | Spades
    AceSpades uint32 = Ace | Spades

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

    KingDiamonds uint32 = King | Diamonds
    QueenDiamonds uint32 = Queen | Diamonds
    JackDiamonds uint32 = Jack | Diamonds
    TenDiamonds uint32 = Ten | Diamonds
    NineDiamonds uint32 = Nine | Diamonds
    EightDiamonds uint32 = Eight | Diamonds
    SevenDiamonds uint32 = Seven | Diamonds
    SixDiamonds uint32 = Six | Diamonds
    FiveDiamonds uint32 = Five | Diamonds
    FourDiamonds uint32 = Four | Diamonds
    ThreeDiamonds uint32 = Three | Diamonds
    TwoDiamonds uint32 = Two | Diamonds
    AceDiamonds uint32 = Ace | Diamonds

    KingHearts uint32 = King | Hearts
    QueenHearts uint32 = Queen | Hearts
    JackHearts uint32 = Jack | Hearts
    TenHearts uint32 = Ten | Hearts
    NineHearts uint32 = Nine | Hearts
    EightHearts uint32 = Eight | Hearts
    SevenHearts uint32 = Seven | Hearts
    SixHearts uint32 = Six | Hearts
    FiveHearts uint32 = Five | Hearts
    FourHearts uint32 = Four | Hearts
    ThreeHearts uint32 = Three | Hearts
    TwoHearts uint32 = Two | Hearts
    AceHearts uint32 = Ace | Hearts
)

var Deck = [52]uint32{
    KingSpades, QueenSpades, JackSpades, TenSpades, NineSpades, EightSpades,
    SevenSpades, SixSpades, FiveSpades, FourSpades, ThreeSpades, TwoSpades,
    AceSpades, KingClubs, QueenClubs, JackClubs, TenClubs, NineClubs,
    EightClubs, SevenClubs, SixClubs, FiveClubs, FourClubs, ThreeClubs,
    TwoClubs, AceClubs, KingDiamonds, QueenDiamonds, JackDiamonds, TenDiamonds,
    NineDiamonds, EightDiamonds, SevenDiamonds, SixDiamonds, FiveDiamonds,
    FourDiamonds, ThreeDiamonds, TwoDiamonds, AceDiamonds, KingHearts,
    QueenHearts, JackHearts, TenHearts, NineHearts, EightHearts, SevenHearts,
    SixHearts, FiveHearts, FourHearts, ThreeHearts, TwoHearts, AceHearts,
}

var strings = map[uint32]string{
    KingSpades: "Ks",
    QueenSpades: "Qs",
    JackSpades: "Js",
    TenSpades: "Ts",
    NineSpades: "9s",
    EightSpades: "8s",
    SevenSpades: "7s",
    SixSpades: "6s",
    FiveSpades: "5s",
    FourSpades: "4s",
    ThreeSpades: "3s",
    TwoSpades: "2s",
    AceSpades: "As",

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

    KingDiamonds: "Kd",
    QueenDiamonds: "Qd",
    JackDiamonds: "Jd",
    TenDiamonds: "Td",
    NineDiamonds: "9d",
    EightDiamonds: "8d",
    SevenDiamonds: "7d",
    SixDiamonds: "6d",
    FiveDiamonds: "5d",
    FourDiamonds: "4d",
    ThreeDiamonds: "3d",
    TwoDiamonds: "2d",
    AceDiamonds: "Ad",

    KingHearts: "Kh",
    QueenHearts: "Qh",
    JackHearts: "Jh",
    TenHearts: "Th",
    NineHearts: "9h",
    EightHearts: "8h",
    SevenHearts: "7h",
    SixHearts: "6h",
    FiveHearts: "5h",
    FourHearts: "4h",
    ThreeHearts: "3h",
    TwoHearts: "2h",
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

func evaluate(cards [5]uint32) uint32 {
    // Sort cards by face value from largest to smallest.
    cards = sort(cards)

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

func string_to_card(s string) uint32 {
    var card uint32
    switch s[0] {
    case 'K':
        card = King
    case 'k':
        card = King
    case 'Q':
        card = Queen
    case 'q':
        card = Queen
    case 'J':
        card = Jack
    case 'j':
        card = Jack
    case 'T':
        card = Ten
    case 't':
        card = Ten
    case '9':
        card = Nine
    case '8':
        card = Eight
    case '7':
        card = Seven
    case '6':
        card = Six
    case '5':
        card = Five
    case '4':
        card = Four
    case '3':
        card = Three
    case '2':
        card = Two
    case 'A':
        card = Ace
    case 'a':
        card = Ace
    default:
        panic("Unrecognized face value character (must be one of: KQJT98765432A)")
    }
    switch s[1] {
    case 's':
        card |= Spades
    case 'c':
        card |= Clubs
    case 'd':
        card |= Diamonds
    case 'h':
        card |= Hearts
    default:
        panic("Unrecognized suit character (must be one of: scdh)")
    }

    return card
}

func eval_to_string(evaluation uint32) string {
    if 0x0000000F & evaluation != 0 {
        // (Face & cards[0]) >> 4
        return "High card"
    }
    if 0x000000F0 & evaluation != 0 && 0xFFFFFF0F & evaluation == 0 {
        // Face & pair_card
        return "Pair"
    }
    if 0x00000FF0 & evaluation != 0 {
        // ((Face & high_pair_card) << 4) | (Face & low_pair_card)
        return "Two pair"
    }
    if 0x0000F000 & evaluation != 0 {
        // (Face & three_of_a_kind_card) << 8
        return "Three of a kind"
    }
    if 0x000F0000 & evaluation != 0 {
        // (Face & straight_high_card) << 12
        return "Straight"
    }
    if 0x00F00000 & evaluation != 0 && 0xFF0FFFFF & evaluation == 0 {
        // (Face & cards[0]) << 16
        return "Flush"
    }
    if 0x0FF00000 & evaluation != 0 && 0xF00FFFFF & evaluation == 0 {
        // (Face & three_of_a_kind_card) << 20 | (Face & high_pair_card) << 16
        return "Full house"
    }
    if 0xF0000000 & evaluation != 0 && 0x0FFFFFFF & evaluation == 0 {
        // (Face & four_of_a_kind_card) << 24
        return "Four of a kind"
    }
    if 0xFF000000 & evaluation != 0 {
        // (Face & straight_high_card | flush_suit) << 24
        return "Royal or straight flush"
    }

    panic("Unable to decode evaluation: " + string(evaluation))
}

func benchmark() {
    var cards [5]uint32
    start := time.Now()
    for a := 0; a < 52 - 4; a++ {
        for b := a + 1; b < 52 - 3; b++ {
            for c := b + 1; c < 52 - 2; c++ {
                for d := c + 1; d < 52 - 1; d++ {
                    for e := d + 1; e < 52; e++ {
                        cards = [5]uint32{Deck[a], Deck[b], Deck[c], Deck[d], Deck[e]}
                        evaluate(cards)
                    }
                }
            }
        }
    }
    end := time.Now()
    fmt.Println(end.Sub(start))
}

func main() {
    if len(os.Args[1:]) == 0 {
        benchmark()
        return
    }

    var hand [5]uint32
    for i := 1; i <= len(os.Args[1:]); i++ {
        hand[i - 1] = string_to_card(os.Args[i])
    }

    fmt.Println(cards_to_string(hand))
    evaluation := evaluate(hand)
    fmt.Println(evaluation)
    fmt.Println(eval_to_string(evaluation))
}
