CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "UserAccount"(
    ID SERIAL NOT NULL,
    Username varchar(50) UNIQUE NOT NULL,
    Password varchar(72) NOT NULL,
    Email varchar(120) NOT NULL, 
    Uuid varchar(22) UNIQUE NOT NULL,
    Active boolean NOT NULL DEFAULT FALSE,
    CreatedAt timestamp NOT NULL DEFAULT NOW(),
    UpdatedAt timestamp,
    PRIMARY KEY (ID)
);

CREATE TABLE "UserPost"(
       ID Serial NOT NULL,
       Owner int,
       Title varchar(120) NOT NULL,
       Body varchar(10000),
       CreatedAt timestamp NOT NULL DEFAULT NOW(),
       UpdatedAt timestamp,
       PRIMARY KEY (ID),
       FOREIGN KEY (Owner) REFERENCES "UserAccount"(ID)
);

CREATE TABLE "Comment"(
       ID Serial NOT NULL,
       Owner int,
       Body varchar(5000),
       CreatedAt timestamp NOT NULL DEFAULT NOW(),
       UpdatedAt timestamp,
       PRIMARY KEY (ID),
       FOREIGN KEY (Owner) REFERENCES "UserAccount"(ID)
);

CREATE TABLE "CommentOnPost"(
       ID Serial NOT NULL,
       CommentId int,
       PostId int,
       PRIMARY KEY (ID),
       FOREIGN KEY (CommentId) REFERENCES "Comment"(ID),
       FOREIGN KEY (PostId) REFERENCES "UserPost"(ID)
);

CREATE TYPE PointType AS ENUM ('plus', 'minus', 'empty');

CREATE TABLE "PointsOnPost"(
       ID Serial NOT NULL,
       UserId int,
       PostId int,
       Point PointType,
       FOREIGN KEY (UserId) REFERENCES "UserAccount"(ID),
       FOREIGN KEY (PostId) REFERENCES "UserPost"(ID)
);

CREATE TABLE "PointsOnComment"(
  ID Serial NOT NULL,
  UserId int,
  CommentId int,
  Point PointType,
  PRIMARY KEY (ID),
  FOREIGN KEY (UserId) REFERENCES "UserAccount"(ID),
  FOREIGN KEY (CommentId) REFERENCES "Comment"(ID)
);

CREATE TABLE "ReportOnComment" (
  ID Serial NOT NULL,
  ReportingUser int NOT NULL,
  CommentReported int NOT NULL,
  Valid boolean DEFAULT FALSE,
  Details varchar(350) NOT NULL,
  PRIMARY KEY (ID),
  FOREIGN KEY (ReportingUser) REFERENCES "UserAccount"(ID),
  FOREIGN KEY (CommentReported) REFERENCES "Comment"(ID)
);

CREATE TABLE "ReportOnPost"(
  ID Serial NOT NULL,
  ReportingUser int NOT NULL,
  PostReported int NOT NULL,
  Valid boolean DEFAULT FALSE,
  Details varchar(350) NOT NULL,
  PRIMARY KEY (ID),
  FOREIGN KEY (ReportingUser) REFERENCES "UserAccount"(ID),
  FOREIGN KEY (PostReported) REFERENCES "UserPost"(ID)
);

