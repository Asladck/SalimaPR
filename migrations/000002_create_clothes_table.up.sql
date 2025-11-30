CREATE TABLE IF NOT EXISTS clothes (
                                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    name TEXT NOT NULL,
    category TEXT NOT NULL, -- top / bottom / shoes / accessory
    color TEXT,
    season TEXT,
    material TEXT,

    image_url TEXT, -- ссылка на фото в S3 / MinIO

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
    );

-- Индексы
CREATE INDEX IF NOT EXISTS idx_clothes_user_id ON clothes(user_id);
CREATE INDEX IF NOT EXISTS idx_clothes_category ON clothes(category);
