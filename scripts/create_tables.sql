CREATE TABLE IF NOT EXISTS ANNOTATIONS (
    id               SERIAL   PRIMARY KEY,
    content          TEXT     NOT NULL,
    x                INTEGER  NOT NULL CHECK (x <= x2 AND x >= 0),
    x2               INTEGER  NOT NULL CHECK (x2 >= x AND x2 >= 0),
    y                INTEGER  NOT NULL CHECK (y <= y2 AND y >= 0),
    y2               INTEGER  NOT NULL CHECK (y2 >= y AND y2 >= 0),
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS PHOTOS (
    id               SERIAL   PRIMARY KEY,
    content          TEXT     NOT NULL,
    is_additional    BOOLEAN  NOT NULL DEFAULT FALSE,
    annotation_id    INTEGER  DEFAULT -1,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE ANNOTATIONS
ADD CONSTRAINT photo_id
FOREIGN KEY (id)
REFERENCES PHOTOS(id)
ON DELETE CASCADE;

ALTER TABLE PHOTOS
ADD CONSTRAINT annotation_id
FOREIGN KEY (id)
REFERENCES ANNOTATIONS(id)
ON DELETE CASCADE;
