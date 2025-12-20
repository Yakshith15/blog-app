CREATE INDEX idx_blogs_author_id_id
ON blogs(author_id, id);

CREATE INDEX idx_blogs_created_at_desc
ON blogs(created_at DESC);
