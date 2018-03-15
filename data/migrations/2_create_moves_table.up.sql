CREATE TABLE Moves (
    MoveID int NOT NULL AUTO_INCREMENT,
    MatchID int NOT NULL,
    Player int NOT NULL,
    MoveX int NOT NULL DEFAULT 0,
    MoveY int NOT NULL DEFAULT 0,
    CreatedAt timestamp NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (MoveID, MatchID)
);