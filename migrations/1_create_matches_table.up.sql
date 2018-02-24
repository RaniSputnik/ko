CREATE TABLE Matches (
    MatchID int NOT NULL AUTO_INCREMENT,
    Owner int,
    Opponent int,
    BoardSize int,
    PRIMARY KEY (MatchID)
);