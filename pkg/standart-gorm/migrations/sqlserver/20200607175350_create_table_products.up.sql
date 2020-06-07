BEGIN;

CREATE TABLE [products]
(
    [id] varchar(255) NOT NULL,
    [name] varchar(50) NOT NULL,
    [created_at] varchar(50) NOT NULL,
    [updated_at] varchar(50) NOT NULL,
    PRIMARY KEY ([id])
);
CREATE UNIQUE INDEX [UK_products_name] ON [products]([name]);

END;