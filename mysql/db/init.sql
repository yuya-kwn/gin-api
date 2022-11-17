CREATE DATABASE test_db;
USE test_db;

CREATE TABLE test_json(jdoc_json)(
    ID integer,
    アーティスト varchar(30),
    代表曲 varchar(30),
    primary key (ID)
);

[
    {
        "ID":"1",
        "アーティスト":"LEX",
        "代表曲":"「何でも言っちゃって」"
    },
    {
        "ID":"2",
        "アーティスト":"OZworld"
        "代表曲":"「Peter Son」"
    },
    {
        "ID":"3",
        "アーティスト":"Candee",
        "代表曲":"「ASOBI」"
    },
    {
        "ID":"4",
        "アーティスト":"¥ellow Bucks",
        "代表曲":"「Yessir」"
    },
    {
        "ID":"5",
        "アーティスト":"HONEST a.k.a 湊",
        "代表曲":"「ARIDAMIKAN」"
    },
]