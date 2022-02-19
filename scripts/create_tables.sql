CREATE TABLE IF NOT EXISTS PHOTOS (
    ID               SERIAL   PRIMARY KEY,
    content          TEXT     NOT NULL,
    annotation_id    INTEGER  DEFAULT 0,
    is_additional    BOOLEAN  NOT NULL DEFAULT FALSE,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS ANNOTATIONS (
    ID               SERIAL   PRIMARY KEY,
    photo_id         INTEGER  REFERENCES PHOTOS(ID) NOT NULL,
    content          TEXT     NOT NULL,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
