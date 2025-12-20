CREATE TABLE blogs (
    id UUID PRIMARY KEY,
    author_id UUID NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_blogs_author_id ON blogs(author_id);
CREATE INDEX idx_blogs_created_at ON blogs(created_at);
