package main

import "fmt"
import "os"
import "time"

import "github.com/dolph/poker-hand-evaluator/handeval"

const (
    KingSpades uint32 = handeval.King | handeval.Spades
    QueenSpades uint32 = handeval.Queen | handeval.Spades
    JackSpades uint32 = handeval.Jack | handeval.Spades
    TenSpades uint32 = handeval.Ten | handeval.Spades
    NineSpades uint32 = handeval.Nine | handeval.Spades
    EightSpades uint32 = handeval.Eight | handeval.Spades
    SevenSpades uint32 = handeval.Seven | handeval.Spades
    SixSpades uint32 = handeval.Six | handeval.Spades
    FiveSpades uint32 = handeval.Five | handeval.Spades
    FourSpades uint32 = handeval.Four | handeval.Spades
    ThreeSpades uint32 = handeval.Three | handeval.Spades
    TwoSpades uint32 = handeval.Two | handeval.Spades
    AceSpades uint32 = handeval.Ace | handeval.Spades

    KingClubs uint32 = handeval.King | handeval.Clubs
    QueenClubs uint32 = handeval.Queen | handeval.Clubs
    JackClubs uint32 = handeval.Jack | handeval.Clubs
    TenClubs uint32 = handeval.Ten | handeval.Clubs
    NineClubs uint32 = handeval.Nine | handeval.Clubs
    EightClubs uint32 = handeval.Eight | handeval.Clubs
    SevenClubs uint32 = handeval.Seven | handeval.Clubs
    SixClubs uint32 = handeval.Six | handeval.Clubs
    FiveClubs uint32 = handeval.Five | handeval.Clubs
    FourClubs uint32 = handeval.Four | handeval.Clubs
    ThreeClubs uint32 = handeval.Three | handeval.Clubs
    TwoClubs uint32 = handeval.Two | handeval.Clubs
    AceClubs uint32 = handeval.Ace | handeval.Clubs

    KingDiamonds uint32 = handeval.King | handeval.Diamonds
    QueenDiamonds uint32 = handeval.Queen | handeval.Diamonds
    JackDiamonds uint32 = handeval.Jack | handeval.Diamonds
    TenDiamonds uint32 = handeval.Ten | handeval.Diamonds
    NineDiamonds uint32 = handeval.Nine | handeval.Diamonds
    EightDiamonds uint32 = handeval.Eight | handeval.Diamonds
    SevenDiamonds uint32 = handeval.Seven | handeval.Diamonds
    SixDiamonds uint32 = handeval.Six | handeval.Diamonds
    FiveDiamonds uint32 = handeval.Five | handeval.Diamonds
    FourDiamonds uint32 = handeval.Four | handeval.Diamonds
    ThreeDiamonds uint32 = handeval.Three | handeval.Diamonds
    TwoDiamonds uint32 = handeval.Two | handeval.Diamonds
    AceDiamonds uint32 = handeval.Ace | handeval.Diamonds

    KingHearts uint32 = handeval.King | handeval.Hearts
    QueenHearts uint32 = handeval.Queen | handeval.Hearts
    JackHearts uint32 = handeval.Jack | handeval.Hearts
    TenHearts uint32 = handeval.Ten | handeval.Hearts
    NineHearts uint32 = handeval.Nine | handeval.Hearts
    EightHearts uint32 = handeval.Eight | handeval.Hearts
    SevenHearts uint32 = handeval.Seven | handeval.Hearts
    SixHearts uint32 = handeval.Six | handeval.Hearts
    FiveHearts uint32 = handeval.Five | handeval.Hearts
    FourHearts uint32 = handeval.Four | handeval.Hearts
    ThreeHearts uint32 = handeval.Three | handeval.Hearts
    TwoHearts uint32 = handeval.Two | handeval.Hearts
    AceHearts uint32 = handeval.Ace | handeval.Hearts
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
    handeval.Spades: "spades",
    handeval.Clubs: "clubs",
    handeval.Diamonds: "diamonds",
    handeval.Hearts: "hearts",
    0: "none",
}

var values = map[uint32]string{
    handeval.King: "king",
    handeval.Queen: "queen",
    handeval.Jack: "jack",
    handeval.Ten: "ten",
    handeval.Nine: "nine",
    handeval.Eight: "eight",
    handeval.Seven: "seven",
    handeval.Six: "six",
    handeval.Five: "five",
    handeval.Four: "four",
    handeval.Three: "three",
    handeval.Two: "two",
    handeval.Ace: "ace",
    0: "none",
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
    return suits[handeval.Suit & card]
}

func value_to_string(card uint32) string {
    return values[(handeval.Face | handeval.Value) & card]
}

func string_to_card(s string) uint32 {
    var card uint32
    switch s[0] {
    case 'K':
        card = handeval.King
    case 'k':
        card = handeval.King
    case 'Q':
        card = handeval.Queen
    case 'q':
        card = handeval.Queen
    case 'J':
        card = handeval.Jack
    case 'j':
        card = handeval.Jack
    case 'T':
        card = handeval.Ten
    case 't':
        card = handeval.Ten
    case '9':
        card = handeval.Nine
    case '8':
        card = handeval.Eight
    case '7':
        card = handeval.Seven
    case '6':
        card = handeval.Six
    case '5':
        card = handeval.Five
    case '4':
        card = handeval.Four
    case '3':
        card = handeval.Three
    case '2':
        card = handeval.Two
    case 'A':
        card = handeval.Ace
    case 'a':
        card = handeval.Ace
    default:
        panic("Unrecognized face value character (must be one of: KQJT98765432A)")
    }
    switch s[1] {
    case 's':
        card |= handeval.Spades
    case 'c':
        card |= handeval.Clubs
    case 'd':
        card |= handeval.Diamonds
    case 'h':
        card |= handeval.Hearts
    default:
        panic("Unrecognized suit character (must be one of: scdh)")
    }

    return card
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
                        handeval.Evaluate(cards)
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
    evaluation := handeval.Evaluate(hand)
    fmt.Println(evaluation)
    fmt.Println(handeval.ToString(evaluation))
}
