CREATE TABLE IF NOT EXISTS ANNOTATIONS (
    id               SERIAL   PRIMARY KEY,
    content          TEXT     NOT NULL,
    x                INTEGER  NOT NULL CHECK (x <= x2 AND x >= 0),
    x2               INTEGER  NOT NULL CHECK (x2 >= x AND x2 >= 0),
    y                INTEGER  NOT NULL CHECK (y <= y2 AND y >= 0),
    y2               INTEGER  NOT NULL CHECK (y2 >= y AND y2 >= 0),
    photo_id         INTEGER,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (content, x, x2, y, y2, photo_id)
);

CREATE TABLE IF NOT EXISTS PHOTOS (
    id               SERIAL   PRIMARY KEY,
    content          TEXT     NOT NULL,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE ANNOTATIONS
ADD CONSTRAINT fk_photo_id
FOREIGN KEY (photo_id)
REFERENCES PHOTOS(id)
ON DELETE CASCADE;
