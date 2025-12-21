CREATE TABLE comments (
    id UUID PRIMARY KEY,
    blogId UUID NOT NULL,
    authorId UUID NOT NULL,
    content TEXT NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL,
    updatedAt TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_comments_blogId_createdAt
ON comments (blogId, createdAt);
