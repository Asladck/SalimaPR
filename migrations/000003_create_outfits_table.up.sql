CREATE TABLE IF NOT EXISTS outfits (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    top_id UUID REFERENCES clothes(id) ON DELETE SET NULL,
    bottom_id UUID REFERENCES clothes(id) ON DELETE SET NULL,
    shoes_id UUID REFERENCES clothes(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT NOW()
    );

CREATE INDEX IF NOT EXISTS idx_outfits_user_id ON outfits(user_id);
