package hotel



// embedded type

type Hotel struct {
    Name     string
    Capacity uint8
    Rooms    uint8
    Smoking  bool
    Country
}

type Country struct {
    Name        string
    CapitalCity string
}